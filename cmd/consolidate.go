package cmd

import (
	"context"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var (
	ignoreAllowlist             bool
	includeInvalid              bool
	calculateChecksum           bool
	skipConsolidatedSummary     bool
	generateConflictsReport     bool
	emitResolvedLists           bool
	applyResolvedToConsolidated bool
	forceConsolidate            bool
)

var consolidateCmd = &cobra.Command{
	Use:   "consolidate",
	Short: "Consolidate processed files",
	Run: func(cmd *cobra.Command, args []string) {
		consolidateAllCmd.Run(cmd, args)
	},
}

var consolidateAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Consolidate all processed files",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Consolidating all processed entries...")
		ctx := context.Background()

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "general", forceConsolidate)
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		if database == nil {
			return
		}
		defer database.CloseLogError(Logger)

		entriesRepo := db.NewEntriesRepo(database)

		// get generic source types
		genericSourceTypes, typesErr := entriesRepo.GetGenericSourceTypes(ctx)
		if typesErr != nil {
			Logger.Errorf("Failed to get source types: %v", typesErr)
			return
		}
		if len(genericSourceTypes) == 0 {
			Logger.Errorf("No source types found in database")
			return
		}

		Logger.Infof("Found %d generic source types: %v", len(genericSourceTypes), genericSourceTypes)

		typesToProcess, typeFingerprints := filterTypesForConsolidation(
			ctx, Logger, consolidatedRepo, "general", genericSourceTypes, forceConsolidate,
		)

		if len(typesToProcess) == 0 {
			Logger.Infof("All source types unchanged, nothing to consolidate")
			saveConsolidationFingerprints(Logger, consolidatedRepo, "general", typesToProcess, typeFingerprints)
			Logger.Infof("Consolidation complete")
			return
		}

		Logger.Infof("Processing %d changed source types: %v", len(typesToProcess), typesToProcess)

		// build resolution sets for conflict resolution
		resResult, resErr := GetResolutionSets(Logger, database)
		var allowByType, blockByType map[string]u.StringSet
		if resErr != nil {
			Logger.Warnf("Failed to build resolution sets (continuing without): %v", resErr)
			allowByType = make(map[string]u.StringSet)
			blockByType = make(map[string]u.StringSet)
		} else {
			allowByType = resResult.AllowByType
			blockByType = resResult.BlockByType
		}

		// phase 1: process allowlists
		Logger.Infof("Processing allowlists...")
		allowlistEntriesByType := make(map[string]u.StringSet)
		processAllowlists(
			ctx,
			Logger,
			entriesRepo,
			consolidatedRepo,
			typesToProcess,
			blockByType,
			allowlistEntriesByType,
		)

		// build allowlist filter for blocklists
		allowFilterByType := make(map[string]u.StringSet)
		for _, gst := range typesToProcess {
			if aset, ok := allowByType[gst]; ok && aset != nil && aset.Size() > 0 {
				allowFilterByType[gst] = aset
			} else if existing, ok := allowlistEntriesByType[gst]; ok && existing != nil {
				allowFilterByType[gst] = existing
			} else {
				allowFilterByType[gst] = u.NewStringSet([]string{})
			}
		}

		// phase 2: blocklist consolidation
		Logger.Infof("Processing blocklists...")

		var resolvedAllow []db.ResolvedAllowEntry
		for gst, set := range allowFilterByType {
			for entry := range set {
				mustConsider, _ := set.Get(entry)
				resolvedAllow = append(resolvedAllow, db.ResolvedAllowEntry{
					GenericSourceType: gst,
					Entry:             entry,
					MustConsider:      mustConsider,
				})
			}
		}
		if err := consolidatedRepo.LoadResolvedAllowSet(ctx, resolvedAllow); err != nil {
			Logger.Errorf("Failed to load resolved allow set: %v", err)
			return
		}

		if err := consolidatedRepo.DropConsolidatedIndexes(ctx); err != nil {
			Logger.Warnf("Failed to drop consolidated indexes: %v", err)
		}

		for _, gst := range typesToProcess {
			result, err := consolidatedRepo.ConsolidateBlocklistGeneral(ctx, gst, true)
			if err != nil {
				Logger.Errorf("Failed to consolidate blocklist for %s: %v", gst, err)
				continue
			}

			if result.FinalCount > 0 {
				filtered := result.OriginalCount - result.FinalCount
				if filtered > 0 {
					Logger.Infof(
						"%s blocklist: %d sources, %d total -> %d final (%d filtered)",
						gst, result.SourceCount, result.OriginalCount, result.FinalCount, filtered,
					)
				} else {
					Logger.Infof(
						"%s blocklist: %d sources, %d total -> %d final",
						gst, result.SourceCount, result.OriginalCount, result.FinalCount,
					)
				}
			}

			// handle invalid entries if requested
			if includeInvalid {
				invResult, invErr := consolidatedRepo.ConsolidateBlocklistGeneral(ctx, gst, false)
				if invErr != nil {
					Logger.Errorf("Failed to consolidate invalid blocklist for %s: %v", gst, invErr)
					continue
				}
				if invResult.FinalCount > 0 {
					Logger.Infof("%s invalid blocklist: %d entries", gst, invResult.FinalCount)
				}
			}
		}

		if err := consolidatedRepo.CreateConsolidatedIndexes(ctx); err != nil {
			Logger.Warnf("Failed to recreate consolidated indexes: %v", err)
		}

		if generateConflictsReport {
			manager := NewConsolidationManager(Logger, database)
			if err := manager.GenerateConflictReportFromResult(resResult); err != nil {
				Logger.Errorf("Failed to generate conflicts report: %v", err)
			}
		}

		saveConsolidationFingerprints(Logger, consolidatedRepo, "general", typesToProcess, typeFingerprints)

		Logger.Infof("Consolidation complete")
	},
}

// processAllowlists processes allowlists
func processAllowlists(
	ctx context.Context,
	logger *multilog.Logger,
	entriesRepo *db.EntriesRepo,
	consolidatedRepo *db.ConsolidatedRepo,
	genericSourceTypes []string,
	resolvedBlockByType map[string]u.StringSet,
	allowlistEntriesByType map[string]u.StringSet,
) {
	for _, gst := range genericSourceTypes {
		// get all allowlist entries
		allowEntries, err := entriesRepo.GetEntriesForGeneralConsolidation(ctx, gst, constants.ListTypeAllowlist)
		if err != nil {
			logger.Errorf("Failed to get allowlist entries for %s: %v", gst, err)
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
			continue
		}

		if len(allowEntries) == 0 {
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
			continue
		}

		// build entry set, tracking must_consider
		entrySet := u.NewStringSet([]string{})
		mustConsiderSet := u.NewStringSet([]string{})
		for _, e := range allowEntries {
			entrySet.AddWithConsider(e.Entry, e.MustConsider)
			if e.MustConsider {
				mustConsiderSet.AddWithConsider(e.Entry, true)
			}
		}

		// get resolved blocklist for filtering
		var resolvedBlocklist u.StringSet
		if bset, ok := resolvedBlockByType[gst]; ok && bset != nil {
			resolvedBlocklist = bset
			logger.Infof("Resolved blocklist for %s: %d entries", gst, bset.Size())
		} else {
			resolvedBlocklist = u.NewStringSet([]string{})
		}

		// filter: remove entries that are in resolved blocklist (unless must_consider)
		finalAllowlist := u.NewStringSet([]string{})
		removedByResolution := 0

		for entry := range entrySet {
			mustConsider, _ := entrySet.Get(entry)
			inResolvedBlock := resolvedBlocklist.Contains(entry)
			isMustConsider := mustConsiderSet.Contains(entry)

			if isMustConsider {
				finalAllowlist.AddWithConsider(entry, true)
			} else if !inResolvedBlock {
				finalAllowlist.AddWithConsider(entry, mustConsider)
			} else {
				removedByResolution++
			}
		}

		logger.Infof(
			"Final allowlist for %s: consolidated=%d removed_by_resolution=%d must_consider=%d final=%d",
			gst, entrySet.Size(), removedByResolution, mustConsiderSet.Size(), finalAllowlist.Size(),
		)

		// persist allowlist entries
		if finalAllowlist.Size() > 0 {
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo,
				finalAllowlist, gst,
				constants.ListTypeAllowlist, "general", "", "", true,
			); err != nil {
				logger.Errorf("Failed to persist allowlist entries for %s: %v", gst, err)
			}
		}

		allowlistEntriesByType[gst] = finalAllowlist

		// handle invalid allowlist entries
		if includeInvalid {
			invalidEntries, invErr := entriesRepo.GetInvalidEntriesForGeneralConsolidation(
				ctx,
				gst,
				constants.ListTypeAllowlist,
			)
			if invErr != nil {
				logger.Errorf("Failed to get invalid allowlist entries for %s: %v", gst, invErr)
				continue
			}

			if len(invalidEntries) > 0 {
				invalidSet := u.NewStringSet([]string{})
				for _, e := range invalidEntries {
					invalidSet.AddWithConsider(e.Entry, e.MustConsider)
				}

				if err := persistConsolidatedEntries(
					ctx, logger, consolidatedRepo,
					invalidSet, gst,
					constants.ListTypeAllowlist, "general", "", "", false,
				); err != nil {
					logger.Errorf("Failed to persist invalid allowlist entries for %s: %v", gst, err)
				}
			}
		}
	}
	logger.Debugf("Finished processing allowlists")
}

func init() {
	consolidateCmd.PersistentFlags().
		BoolVar(&forceConsolidate, "force", false, "Force consolidation even if nothing has changed since last run")
	consolidateCmd.PersistentFlags().
		BoolVar(&ignoreAllowlist, "ignore-allowlist", false, "Ignore allowlist during consolidation where applicable")
	consolidateCmd.PersistentFlags().
		BoolVar(&includeInvalid, "include-invalid", false, "Include invalid entry(s) during consolidation")
	consolidateCmd.PersistentFlags().
		BoolVar(&calculateChecksum, "calculate-checksum", false, "Calculate checksum on the consolidated files")
	// nolint:lll
	consolidateCmd.PersistentFlags().
		BoolVar(&generateConflictsReport, "gen-conflicts", false, "Generate a conflict report, allowlist vs. blocklist")
	consolidateCmd.PersistentFlags().
		BoolVar(&emitResolvedLists, "emit-resolved-lists", false, "Emit allowlist and blocklist when resolving conflicts")
	// nolint:lll
	consolidateCmd.PersistentFlags().
		BoolVar(&applyResolvedToConsolidated, "apply-resolved-to-consolidated", true, "Apply resolved allow sets to consolidated output files (opt-in)")
	consolidateCategoriesCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the consolidated summary file")
	// nolint:lll
	consolidateGroupsCmd.PersistentFlags().
		BoolVar(&skipConsolidatedSummary, "skip-consolidated-summary", false, "Skip creating the regular consolidated summary file")
	consolidateCmd.AddCommand(consolidateAllCmd)
	consolidateCmd.AddCommand(consolidateGroupsCmd)
	consolidateCmd.AddCommand(consolidateCategoriesCmd)
}
