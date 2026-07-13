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

var consolidateGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Generate different sized consolidated lists (mini, lite, normal, big)",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Generating sized consolidated lists...")
		ctx := context.Background()

		database, consolidatedRepo, dbErr := openConsolidatedRepo(ctx, Logger, "group")
		if dbErr != nil {
			Logger.Errorf("Failed to open consolidated repo: %v", dbErr)
			return
		}
		defer database.CloseLogError(Logger)

		entriesRepo := db.NewEntriesRepo(database)

		// get unique groups directly from DB
		groups, groupsErr := entriesRepo.GetUniqueGroups(ctx)
		if groupsErr != nil {
			Logger.Errorf("Failed to get groups: %v", groupsErr)
			return
		}
		if len(groups) == 0 {
			Logger.Errorf("No groups found in database")
			return
		}

		// get generic source types from DB
		genericSourceTypes, typesErr := entriesRepo.GetGenericSourceTypes(ctx)
		if typesErr != nil {
			Logger.Errorf("Failed to get source types: %v", typesErr)
			return
		}

		Logger.Infof("Found %d unique groups: %v", len(groups), groups)
		Logger.Infof("Found %d generic source types: %v", len(genericSourceTypes), genericSourceTypes)

		var persistMu sync.Mutex

		// Process each size group and create consolidated lists
		for _, group := range groups {
			processGroupConsolidation(
				ctx,
				Logger,
				group,
				genericSourceTypes,
				entriesRepo,
				consolidatedRepo,
				&persistMu,
			)
		}

		Logger.Infof("Groups consolidation complete")
	},
}

// processGroupConsolidation processes consolidation for a specific group using DB entries
func processGroupConsolidation(
	ctx context.Context,
	logger *multilog.Logger,
	group string,
	genericSourceTypes []string,
	entriesRepo *db.EntriesRepo,
	consolidatedRepo *db.ConsolidatedRepo,
	persistMu *sync.Mutex,
) {
	logger.Infof("Processing group: %s", group)

	// process allowlists first
	allowlistEntriesByType := make(map[string]u.StringSet)
	for _, gst := range genericSourceTypes {
		allowEntries, err := entriesRepo.GetEntriesByGroup(ctx, group, gst, constants.ListTypeAllowlist)
		if err != nil {
			logger.Errorf("Failed to get allowlist entries for group %s, type %s: %v", group, gst, err)
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
				"%s allowlist [Group %s]: %d sources, %d entries",
				gst,
				group,
				len(sourceNames),
				entrySet.Size(),
			)

			// persist allowlist entries
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo, persistMu,
				entrySet, gst,
				constants.ListTypeAllowlist, "group", group, "", true,
			); err != nil {
				logger.Errorf("Failed to persist allowlist entries for group %s, type %s: %v", group, gst, err)
			}
		} else {
			allowlistEntriesByType[gst] = u.NewStringSet([]string{})
		}
	}

	// process blocklists, filtering with allowlist entries
	for _, gst := range genericSourceTypes {
		blockEntries, err := entriesRepo.GetEntriesByGroup(ctx, group, gst, constants.ListTypeBlocklist)
		if err != nil {
			logger.Errorf("Failed to get blocklist entries for group %s, type %s: %v", group, gst, err)
			continue
		}

		if len(blockEntries) == 0 {
			continue
		}

		// build entry set and track sources
		entrySet := u.NewStringSet([]string{})
		sourceNames := make(map[string]struct{})
		for _, e := range blockEntries {
			entrySet.AddWithConsider(e.Entry, e.MustConsider)
			sourceNames[e.SourceName] = struct{}{}
		}

		originalCount := entrySet.Size()

		// filter with allowlist entries
		allowlistEntries := allowlistEntriesByType[gst]
		filteredSet, ignoredSet := filterEntriesWithAllowlist(entrySet, allowlistEntries)

		if filteredSet.Size() > 0 {
			if ignoredSet.Size() > 0 {
				logger.Infof(
					"%s blocklist [Group %s]: %d sources, %d total -> %d final (%d filtered)",
					gst,
					group,
					len(sourceNames),
					originalCount,
					filteredSet.Size(),
					ignoredSet.Size(),
				)
			} else {
				logger.Infof(
					"%s blocklist [Group %s]: %d sources, %d total -> %d final",
					gst,
					group,
					len(sourceNames),
					originalCount,
					filteredSet.Size(),
				)
			}

			// persist blocklist entries
			if err := persistConsolidatedEntries(
				ctx, logger, consolidatedRepo, persistMu,
				filteredSet, gst,
				constants.ListTypeBlocklist, "group", group, "", true,
			); err != nil {
				logger.Errorf("Failed to persist blocklist entries for group %s, type %s: %v", group, gst, err)
			}
		}
	}
}
