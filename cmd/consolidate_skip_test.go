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
