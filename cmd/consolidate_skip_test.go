package cmd

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/stretchr/testify/assert"
)

// TestSkipGeneralConsolidationFeature demonstrates the skip_general_consolidation functionality
func TestSkipGeneralConsolidationFeature(t *testing.T) {
	sourcesConfig := config.SourcesConfig{
		Sources: []config.Source{
			{
				Name:                     "regular-source",
				URL:                      "http://example.com/regular.txt",
				SkipGeneralConsolidation: false,
				Disabled:                 false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
			{
				Name:                     "large-source",
				URL:                      "http://example.com/large.txt",
				SkipGeneralConsolidation: true,
				Disabled:                 false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
		},
	}

	appConfig := config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			MaxWorkers: 4,
		},
	}

	sourceConfigs := []config.SourcesConfig{sourcesConfig}

	t.Run("General consolidation excludes skip_general_consolidation sources", func(t *testing.T) {
		generalSources := sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, generalSources, 1)
		assert.Equal(t, "regular-source", generalSources[0].Name)

		assert.True(t, config.IsEnabledSourceForConsolidation("regular-source", sourceConfigs, appConfig, "general"))
		assert.False(t, config.IsEnabledSourceForConsolidation("large-source", sourceConfigs, appConfig, "general"))
	})

	t.Run("Groups consolidation includes skip_general_consolidation sources", func(t *testing.T) {
		groupsSources := sourcesConfig.GetSourcesForGroupsConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, groupsSources, 2)

		sourceNames := make(map[string]bool)
		for _, s := range groupsSources {
			sourceNames[s.Name] = true
		}
		assert.True(t, sourceNames["regular-source"])
		assert.True(t, sourceNames["large-source"])

		assert.True(t, config.IsEnabledSourceForConsolidation("regular-source", sourceConfigs, appConfig, "groups"))
		assert.True(t, config.IsEnabledSourceForConsolidation("large-source", sourceConfigs, appConfig, "groups"))
	})

	t.Run("Categories consolidation includes skip_general_consolidation sources", func(t *testing.T) {
		categoriesSources := sourcesConfig.GetSourcesForCategoriesConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, categoriesSources, 2)

		assert.True(t, config.IsEnabledSourceForConsolidation("regular-source", sourceConfigs, appConfig, "categories"))
		assert.True(t, config.IsEnabledSourceForConsolidation("large-source", sourceConfigs, appConfig, "categories"))
	})

	t.Run("Source methods work correctly", func(t *testing.T) {
		regularSource := sourcesConfig.Sources[0]
		largeSource := sourcesConfig.Sources[1]

		assert.True(t, regularSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, regularSource.ShouldIncludeInGroupsConsolidation())

		assert.False(t, largeSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, largeSource.ShouldIncludeInGroupsConsolidation())
	})
}

func TestSkipGroupsConsolidationFlag(t *testing.T) {
	sourcesConfig := config.SourcesConfig{
		Sources: []config.Source{
			{
				Name:                    "normal-source",
				URL:                     "http://example.com/normal.txt",
				SkipGroupsConsolidation: false,
				Disabled:                false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
			{
				Name:                    "no-groups-source",
				URL:                     "http://example.com/no-groups.txt",
				SkipGroupsConsolidation: true,
				Disabled:                false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
		},
	}

	appConfig := config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			MaxWorkers: 1,
		},
	}

	sourceConfigs := []config.SourcesConfig{sourcesConfig}

	t.Run("Groups consolidation excludes skip_groups_consolidation sources", func(t *testing.T) {
		groupsSources := sourcesConfig.GetSourcesForGroupsConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, groupsSources, 1)
		assert.Equal(t, "normal-source", groupsSources[0].Name)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "groups"))
		assert.False(t, config.IsEnabledSourceForConsolidation("no-groups-source", sourceConfigs, appConfig, "groups"))
	})

	t.Run("General consolidation includes skip_groups_consolidation sources", func(t *testing.T) {
		generalSources := sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, generalSources, 2)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "general"))
		assert.True(t, config.IsEnabledSourceForConsolidation("no-groups-source", sourceConfigs, appConfig, "general"))
	})

	t.Run("Categories consolidation includes skip_groups_consolidation sources", func(t *testing.T) {
		categoriesSources := sourcesConfig.GetSourcesForCategoriesConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, categoriesSources, 2)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "categories"))
		assert.True(
			t,
			config.IsEnabledSourceForConsolidation("no-groups-source", sourceConfigs, appConfig, "categories"),
		)
	})

	t.Run("Source funcs", func(t *testing.T) {
		normalSource := sourcesConfig.Sources[0]
		noGroupsSource := sourcesConfig.Sources[1]

		assert.True(t, normalSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, normalSource.ShouldIncludeInGroupsConsolidation())
		assert.True(t, normalSource.ShouldIncludeInCategoriesConsolidation())

		assert.True(t, noGroupsSource.ShouldIncludeInGeneralConsolidation())
		assert.False(t, noGroupsSource.ShouldIncludeInGroupsConsolidation())
		assert.True(t, noGroupsSource.ShouldIncludeInCategoriesConsolidation())
	})
}

func TestSkipCategoriesConsolidationFlag(t *testing.T) {
	sourcesConfig := config.SourcesConfig{
		Sources: []config.Source{
			{
				Name:                        "normal-source",
				URL:                         "http://example.com/normal.txt",
				SkipCategoriesConsolidation: false,
				Disabled:                    false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
			{
				Name:                        "no-categories-source",
				URL:                         "http://example.com/no-categories.txt",
				SkipCategoriesConsolidation: true,
				Disabled:                    false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
		},
	}

	appConfig := config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			MaxWorkers: 1,
		},
	}

	sourceConfigs := []config.SourcesConfig{sourcesConfig}

	t.Run("Categories consolidation excludes skip_categories_consolidation sources", func(t *testing.T) {
		categoriesSources := sourcesConfig.GetSourcesForCategoriesConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, categoriesSources, 1)
		assert.Equal(t, "normal-source", categoriesSources[0].Name)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "categories"))
		assert.False(
			t,
			config.IsEnabledSourceForConsolidation("no-categories-source", sourceConfigs, appConfig, "categories"),
		)
	})

	t.Run("General consolidation includes skip_categories_consolidation sources", func(t *testing.T) {
		generalSources := sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, generalSources, 2)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "general"))
		assert.True(
			t,
			config.IsEnabledSourceForConsolidation("no-categories-source", sourceConfigs, appConfig, "general"),
		)
	})

	t.Run("Groups consolidation includes skip_categories_consolidation sources", func(t *testing.T) {
		groupsSources := sourcesConfig.GetSourcesForGroupsConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, groupsSources, 2)

		assert.True(t, config.IsEnabledSourceForConsolidation("normal-source", sourceConfigs, appConfig, "groups"))
		assert.True(
			t,
			config.IsEnabledSourceForConsolidation("no-categories-source", sourceConfigs, appConfig, "groups"),
		)
	})

	t.Run("Source funcs", func(t *testing.T) {
		normalSource := sourcesConfig.Sources[0]
		noCategoriesSource := sourcesConfig.Sources[1]

		assert.True(t, normalSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, normalSource.ShouldIncludeInGroupsConsolidation())
		assert.True(t, normalSource.ShouldIncludeInCategoriesConsolidation())

		assert.True(t, noCategoriesSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, noCategoriesSource.ShouldIncludeInGroupsConsolidation())
		assert.False(t, noCategoriesSource.ShouldIncludeInCategoriesConsolidation())
	})
}

func TestAllSkipConsolidationFlags(t *testing.T) {
	sourcesConfig := config.SourcesConfig{
		Sources: []config.Source{
			{
				Name:                        "normal-source",
				URL:                         "http://example.com/normal.txt",
				SkipGeneralConsolidation:    false,
				SkipGroupsConsolidation:     false,
				SkipCategoriesConsolidation: false,
				Disabled:                    false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
			{
				Name:                        "skip-all-source",
				URL:                         "http://example.com/skip-all.txt",
				SkipGeneralConsolidation:    true,
				SkipGroupsConsolidation:     true,
				SkipCategoriesConsolidation: true,
				Disabled:                    false,
				Types: []c.SourceType{
					{Name: "domain"},
				},
			},
		},
	}

	appConfig := config.AppConfig{
		DNSToolkit: config.DNSToolkitConfig{
			MaxWorkers: 4,
		},
	}

	sourceConfigs := []config.SourcesConfig{sourcesConfig}

	t.Run("Source that skips all consolidation types", func(t *testing.T) {
		generalSources := sourcesConfig.GetSourcesForGeneralConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, generalSources, 1)
		assert.Equal(t, "normal-source", generalSources[0].Name)

		groupsSources := sourcesConfig.GetSourcesForGroupsConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, groupsSources, 1)
		assert.Equal(t, "normal-source", groupsSources[0].Name)

		categoriesSources := sourcesConfig.GetSourcesForCategoriesConsolidation(appConfig.DNSToolkit.SourceFilters)
		assert.Len(t, categoriesSources, 1)
		assert.Equal(t, "normal-source", categoriesSources[0].Name)

		assert.False(t, config.IsEnabledSourceForConsolidation("skip-all-source", sourceConfigs, appConfig, "general"))
		assert.False(t, config.IsEnabledSourceForConsolidation("skip-all-source", sourceConfigs, appConfig, "groups"))
		assert.False(
			t,
			config.IsEnabledSourceForConsolidation("skip-all-source", sourceConfigs, appConfig, "categories"),
		)
	})

	t.Run("Source skip methods work correctly", func(t *testing.T) {
		normalSource := sourcesConfig.Sources[0]
		skipAllSource := sourcesConfig.Sources[1]

		assert.True(t, normalSource.ShouldIncludeInGeneralConsolidation())
		assert.True(t, normalSource.ShouldIncludeInGroupsConsolidation())
		assert.True(t, normalSource.ShouldIncludeInCategoriesConsolidation())

		assert.False(t, skipAllSource.ShouldIncludeInGeneralConsolidation())
		assert.False(t, skipAllSource.ShouldIncludeInGroupsConsolidation())
		assert.False(t, skipAllSource.ShouldIncludeInCategoriesConsolidation())
	})
}

func TestProcessedFileSkipFlags(t *testing.T) {
	processedFiles := []c.ProcessedFile{
		{
			Name:                        "normal-file",
			GenericSourceType:           "domain",
			ListType:                    "blocklist",
			Valid:                       true,
			SkipGeneralConsolidation:    false,
			SkipGroupsConsolidation:     false,
			SkipCategoriesConsolidation: false,
		},
		{
			Name:                        "skip-general-file",
			GenericSourceType:           "domain",
			ListType:                    "blocklist",
			Valid:                       true,
			SkipGeneralConsolidation:    true,
			SkipGroupsConsolidation:     false,
			SkipCategoriesConsolidation: false,
		},
		{
			Name:                        "skip-groups-file",
			GenericSourceType:           "domain",
			ListType:                    "blocklist",
			Valid:                       true,
			SkipGeneralConsolidation:    false,
			SkipGroupsConsolidation:     true,
			SkipCategoriesConsolidation: false,
		},
		{
			Name:                        "skip-categories-file",
			GenericSourceType:           "domain",
			ListType:                    "blocklist",
			Valid:                       true,
			SkipGeneralConsolidation:    false,
			SkipGroupsConsolidation:     false,
			SkipCategoriesConsolidation: true,
		},
	}

	t.Run("ProcessedFile skip flags work correctly", func(t *testing.T) {
		normalFile := processedFiles[0]
		skipGeneralFile := processedFiles[1]
		skipGroupsFile := processedFiles[2]
		skipCategoriesFile := processedFiles[3]

		assert.False(t, normalFile.SkipGeneralConsolidation)
		assert.False(t, normalFile.SkipGroupsConsolidation)
		assert.False(t, normalFile.SkipCategoriesConsolidation)

		assert.True(t, skipGeneralFile.SkipGeneralConsolidation)
		assert.False(t, skipGeneralFile.SkipGroupsConsolidation)
		assert.False(t, skipGeneralFile.SkipCategoriesConsolidation)

		assert.False(t, skipGroupsFile.SkipGeneralConsolidation)
		assert.True(t, skipGroupsFile.SkipGroupsConsolidation)
		assert.False(t, skipGroupsFile.SkipCategoriesConsolidation)

		assert.False(t, skipCategoriesFile.SkipGeneralConsolidation)
		assert.False(t, skipCategoriesFile.SkipGroupsConsolidation)
		assert.True(t, skipCategoriesFile.SkipCategoriesConsolidation)
	})
}
