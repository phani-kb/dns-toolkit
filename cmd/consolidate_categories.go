package cmd

import (
	"context"
	"sync"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var consolidateCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Generate category-based consolidated lists (ads, malware, privacy, etc)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating category-based consolidated lists (DB-based)...")
		ctx := context.Background()

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "category")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		entriesRepo := db.NewEntriesRepo(database)

		categories, catErr := entriesRepo.GetUniqueCategoriesFromDB(ctx)
		if catErr != nil {
			Logger.Errorf("Failed to get categories from database: %v", catErr)
			return
		}
		if len(categories) == 0 {
			Logger.Errorf("No categories found in database")
			return
		}

		genericSourceTypes, typesErr := entriesRepo.GetGenericSourceTypesFromDB(ctx)
		if typesErr != nil {
			Logger.Errorf("Failed to get source types from database: %v", typesErr)
			return
		}

		Logger.Infof("Found %d unique categories: %v", len(categories), categories)
		Logger.Infof("Found %d generic source types: %v", len(genericSourceTypes), genericSourceTypes)

		var persistMu sync.Mutex

		for _, category := range categories {
			processCategoryConsolidationFromDB(
				ctx,
				Logger,
				category,
				genericSourceTypes,
				entriesRepo,
				consolidatedRepo,
				&persistMu,
			)
		}

		Logger.Infof("Categories consolidation complete")
	},
}

// processCategoryConsolidationFromDB processes consolidation for a specific category using DB entries
func processCategoryConsolidationFromDB(
	ctx context.Context,
	logger *multilog.Logger,
	category string,
	genericSourceTypes []string,
	entriesRepo *db.EntriesRepo,
	consolidatedRepo *db.ConsolidatedRepo,
	persistMu *sync.Mutex,
) {
	logger.Infof("Processing category: %s", category)

	// process allowlists first
	allowlistEntriesByType := make(map[string]u.StringSet)
	for _, gst := range genericSourceTypes {
		allowEntries, err := entriesRepo.GetEntriesByCategory(ctx, category, gst, constants.ListTypeAllowlist)
		if err != nil {
			logger.Errorf("Failed to get allowlist entries for category %s, type %s: %v", category, gst, err)
			continue
		}

		if len(allowEntries) > 0 {
			entrySet := u.NewStringSet([]string{})
			sourceNames := make(map[string]struct{})
			for _, e := range allowEntries {
				entrySet.AddWithConsider(e.Entry, e.MustConsider)
				sourceNames[e.SourceName] = struct{}{}
			}
			allowlistEntriesByType[gst] = entrySet

			logger.Infof(
				"%s allowlist [Category %s]: %d sources, %d entries",
				gst,
				category,
				len(sourceNames),
				entrySet.Size(),
			)

			// persist allowlist entries
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo, persistMu,
				entrySet, gst,
				constants.ListTypeAllowlist, "category", "", category, true,
			); err != nil {
				logger.Errorf("Failed to persist allowlist entries for category %s, type %s: %v", category, gst, err)
			}
		} else {
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
		}
	}

	// process blocklists, filtering with allowlist entries
	for _, gst := range genericSourceTypes {
		blockEntries, err := entriesRepo.GetEntriesByCategory(ctx, category, gst, constants.ListTypeBlocklist)
		if err != nil {
			logger.Errorf("Failed to get blocklist entries for category %s, type %s: %v", category, gst, err)
			continue
		}

		if len(blockEntries) == 0 {
			continue
		}

		// Build entry set and track sources
		entrySet := u.NewStringSet([]string{})
		sourceNames := make(map[string]struct{})
		for _, e := range blockEntries {
			entrySet.AddWithConsider(e.Entry, e.MustConsider)
			sourceNames[e.SourceName] = struct{}{}
		}

		originalCount := entrySet.Size()

		// Filter with allowlist entries
		allowlistEntries := allowlistEntriesByType[gst]
		filteredSet, ignoredSet := filterEntriesWithAllowlist(entrySet, allowlistEntries)

		if filteredSet.Size() > 0 {
			if ignoredSet.Size() > 0 {
				logger.Infof(
					"%s blocklist [Category %s]: %d sources, %d total -> %d final (%d filtered)",
					gst,
					category,
					len(sourceNames),
					originalCount,
					filteredSet.Size(),
					ignoredSet.Size(),
				)
			} else {
				logger.Infof(
					"%s blocklist [Category %s]: %d sources, %d total -> %d final",
					gst,
					category,
					len(sourceNames),
					originalCount,
					filteredSet.Size(),
				)
			}

			// Persist blocklist entries
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo, persistMu,
				filteredSet, gst,
				constants.ListTypeBlocklist, "category", "", category, true,
			); err != nil {
				logger.Errorf("Failed to persist blocklist entries for category %s, type %s: %v", category, gst, err)
			}
		}
	}
}

// filterEntriesWithAllowlist filters blocklist entries using allowlist entries
func filterEntriesWithAllowlist(blockSet, allowSet u.StringSet) (u.StringSet, u.StringSet) {
	filteredSet := u.NewStringSet([]string{})
	ignoredSet := u.NewStringSet([]string{})

	for entry := range blockSet {
		mustConsiderBlock, _ := blockSet.Get(entry)
		mustConsiderAllow, existsInAllow := allowSet.Get(entry)

		if existsInAllow {
			// entry is in allowlist
			if mustConsiderBlock && !mustConsiderAllow {
				// block entry has higher priority
				filteredSet.AddWithConsider(entry, mustConsiderBlock)
			} else {
				// allow entry wins, ignore from blocklist
				ignoredSet.Add(entry)
			}
		} else {
			// entry not in allowlist, keep it
			filteredSet.AddWithConsider(entry, mustConsiderBlock)
		}
	}

	return filteredSet, ignoredSet
}
