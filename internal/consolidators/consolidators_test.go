package consolidators_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/consolidators"
	"github.com/phani-kb/dns-toolkit/internal/consolidators/mocks"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConsolidatorRegistry tests all registry-related functionality
func TestConsolidatorRegistry(t *testing.T) {
	registry := consolidators.NewConsolidatorRegistry()
	assert.NotNil(t, registry, "Registry should not be nil")

	mockConsolidator := mocks.NewConsolidator(t)
	sourceType := "test-source"
	listType := "blocklist"
	mockConsolidator.On("GetSourceType").Return(sourceType).Maybe()
	mockConsolidator.On("GetListType").Return(listType).Maybe()

	registry.RegisterConsolidator(sourceType, listType, mockConsolidator)

	retrievedConsolidator, exists := registry.GetConsolidator(sourceType, listType)
	assert.True(t, exists, "Consolidator should exist in registry")
	assert.Equal(
		t,
		mockConsolidator,
		retrievedConsolidator,
		"Retrieved consolidator should match registered consolidator",
	)

	_, exists = registry.GetConsolidator("nonexistent", listType)
	assert.False(t, exists, "Non-existent consolidator should not be found")

	consolidatorMap := registry.ListConsolidators()
	assert.Equal(t, 1, len(consolidatorMap), "Registry should contain 1 consolidator")

	key := sourceType + ":" + listType

	listedConsolidator, exists := consolidatorMap[key]
	assert.True(t, exists, "Consolidator should exist in the map with the correct key")
	assert.Equal(
		t,
		mockConsolidator,
		listedConsolidator,
		"Listed consolidator should match registered consolidator",
	)
}

// TestGlobalHelperFunctions tests the global helper functions
func TestGlobalHelperFunctions(t *testing.T) {
	originalRegistry := consolidators.Consolidators
	consolidators.Consolidators = consolidators.NewConsolidatorRegistry()
	defer func() {
		consolidators.Consolidators = originalRegistry
	}()

	sourceType := "test-helper-source"
	listType := "blocklist"

	consolidatorFactory := func(srcType, lstType string) consolidators.Consolidator {
		mockConsolidator := mocks.NewConsolidator(t)
		mockConsolidator.On("GetSourceType").Return(srcType)
		mockConsolidator.On("GetListType").Return(lstType)
		return mockConsolidator
	}

	consolidators.RegisterConsolidator(sourceType, listType, consolidatorFactory)

	consolidator, exists := consolidators.Consolidators.GetConsolidator(sourceType, listType)
	assert.True(t, exists, "Consolidator should exist in global registry")
	assert.Equal(t, sourceType, consolidator.GetSourceType(), "Consolidator should have correct source type")
	assert.Equal(t, listType, consolidator.GetListType(), "Consolidator should have correct list type")

	newSourceType := "test-types-source"
	listTypes := []string{"blocklist", "allowlist", "custom-list"}

	consolidators.Consolidators = consolidators.NewConsolidatorRegistry()

	consolidators.RegisterConsolidatorTypes(newSourceType, listTypes, consolidatorFactory)

	for _, lt := range listTypes {
		cons, exists := consolidators.Consolidators.GetConsolidator(newSourceType, lt)
		assert.True(t, exists, "Consolidator should exist for "+lt)
		assert.Equal(t, newSourceType, cons.GetSourceType(), "Consolidator should have correct source type")
		assert.Equal(t, lt, cons.GetListType(), "Consolidator should have correct list type")
	}
}

// TestBaseConsolidator tests the BaseConsolidator functionality
func TestBaseConsolidator(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"
	bc := consolidators.NewBaseConsolidator(sourceType, listType)

	assert.Equal(t, sourceType, bc.GetSourceType())
	assert.Equal(t, listType, bc.GetListType())

	validFile := common.ProcessedFile{
		GenericSourceType: sourceType,
		ListType:          listType,
	}
	invalidFile := common.ProcessedFile{
		GenericSourceType: "other-source",
		ListType:          listType,
	}

	assert.True(t, bc.IsValid(validFile), "Should validate correct source and list type")
	assert.False(t, bc.IsValid(invalidFile), "Should not validate incorrect source type")
}

// TestBaseConsolidatorConsolidate tests the Consolidate functionality
func TestBaseConsolidatorConsolidate(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"
	bc := consolidators.NewBaseConsolidator(sourceType, listType)
	logger := multilog.NewLogger()

	tempDir := t.TempDir()
	testFile1 := filepath.Join(tempDir, "test1.txt")
	require.NoError(t, os.WriteFile(testFile1, []byte("domain1.com\ndomain2.com\n"), 0644))

	testFile2 := filepath.Join(tempDir, "test2.txt")
	require.NoError(t, os.WriteFile(testFile2, []byte("domain3.com\ndomain4.com\n"), 0644))

	processedFiles := []common.ProcessedFile{
		{
			GenericSourceType: sourceType,
			ListType:          listType,
			Filepath:          testFile1,
			NumberOfEntries:   2,
			Name:              "test1",
		},
		{
			GenericSourceType: sourceType,
			ListType:          listType,
			Filepath:          testFile2,
			NumberOfEntries:   2,
			Name:              "test2",
		},
		{
			GenericSourceType: "other-source",
			ListType:          listType,
			Filepath:          testFile1,
			NumberOfEntries:   2,
			Name:              "invalid",
		},
	}

	consolidatedSet, fileInfos := bc.Consolidate(logger, processedFiles)

	assert.Equal(t, 4, len(consolidatedSet), "Should consolidate entries from both valid files")
	assert.Equal(t, 2, len(fileInfos), "Should have info for two valid files")

	invalidEntryCountFile := common.ProcessedFile{
		GenericSourceType: sourceType,
		ListType:          listType,
		Filepath:          testFile1,
		NumberOfEntries:   5, // Mismatch with actual count (2)
		Name:              "mismatch",
	}

	_, fileInfosMismatch := bc.Consolidate(logger, []common.ProcessedFile{invalidEntryCountFile})
	assert.Equal(t, 0, len(fileInfosMismatch), "Should skip files with entry count mismatch")
}

// TestBaseConsolidatorFilterEntries tests the FilterEntries functionality
func TestBaseConsolidatorFilterEntries(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"
	bc := consolidators.NewBaseConsolidator(sourceType, listType)
	logger := multilog.NewLogger()

	entrySet := utils.NewStringSet([]string{"domain1.com", "domain2.com", "domain3.com", "domain4.com"})
	entrySet.AddWithConsider("domain5.com", true)

	filterSet := utils.NewStringSet([]string{"domain2.com", "domain4.com"})
	filterSet.AddWithConsider("domain2.com", true)

	filteredSet, ignoredSet := bc.FilterEntries(logger, entrySet, filterSet)

	assert.Equal(t, 3, len(filteredSet), "Should have 3 entries after filtering")
	assert.Equal(t, 2, len(ignoredSet), "Should have 2 ignored entries")

	assert.True(t, filteredSet.Contains("domain1.com"), "domain1.com should be in filtered set")
	assert.True(t, filteredSet.Contains("domain3.com"), "domain3.com should be in filtered set")
	assert.True(t, filteredSet.Contains("domain5.com"), "domain5.com should be in filtered set (must consider)")
	assert.True(t, ignoredSet.Contains("domain2.com"), "domain2.com should be ignored")
	assert.True(t, ignoredSet.Contains("domain4.com"), "domain4.com should be ignored")

	emptySet := utils.NewStringSet([]string{})
	emptyFiltered, emptyIgnored := bc.FilterEntries(logger, emptySet, filterSet)
	assert.Equal(t, 0, len(emptyFiltered), "Empty entry set should result in empty filtered set")
	assert.Equal(t, 0, len(emptyIgnored), "Empty entry set should result in empty ignored set")
}

// TestBaseConsolidatorSaveEntries tests the SaveEntries functionality
func TestBaseConsolidatorSaveEntries(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"
	bc := consolidators.NewBaseConsolidator(sourceType, listType)
	logger := multilog.NewLogger()

	entrySet := utils.NewStringSet([]string{"domain1.com", "domain2.com", "domain3.com"})

	tempDir := t.TempDir()
	outFile := filepath.Join(tempDir, "output.txt")

	err := bc.SaveEntries(logger, entrySet, outFile)
	assert.NoError(t, err, "SaveEntries should not return an error")

	content, err := os.ReadFile(outFile)
	assert.NoError(t, err, "Should be able to read the output file")

	fileContent := string(content)
	assert.Contains(t, fileContent, "domain1.com")
	assert.Contains(t, fileContent, "domain2.com")
	assert.Contains(t, fileContent, "domain3.com")
}

// TestGenericConsolidator tests the GenericConsolidator functionality
func TestGenericConsolidator(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"

	testConsolidateFunc := func(logger *multilog.Logger, processedFiles []common.ProcessedFile) (utils.StringSet, []common.FileInfo) {
		result := utils.NewStringSet([]string{"test1.com", "test2.com"})
		fileInfos := []common.FileInfo{
			{
				Name:     "test",
				Filepath: "test/path",
			},
		}
		return result, fileInfos
	}

	gc := consolidators.NewGenericConsolidator(sourceType, listType, testConsolidateFunc)

	assert.Equal(t, sourceType, gc.GetSourceType())
	assert.Equal(t, listType, gc.GetListType())

	validFile := common.ProcessedFile{
		GenericSourceType: sourceType,
		ListType:          listType,
	}
	assert.True(t, gc.IsValid(validFile), "Should validate correct file")

	logger := multilog.NewLogger()
	result, fileInfos := gc.Consolidate(logger, []common.ProcessedFile{})

	assert.Equal(t, 2, len(result), "Should have entries from custom function")
	assert.Equal(t, 1, len(fileInfos), "Should have file info from custom function")
}

// TestRegisterGenericConsolidator tests the generic consolidator registration
func TestRegisterGenericConsolidator(t *testing.T) {
	originalRegistry := consolidators.Consolidators
	consolidators.Consolidators = consolidators.NewConsolidatorRegistry()
	defer func() {
		consolidators.Consolidators = originalRegistry
	}()

	sourceType := "test-generic-source"
	listType := "blocklist"

	testConsolidateFunc := func(logger *multilog.Logger, processedFiles []common.ProcessedFile) (utils.StringSet, []common.FileInfo) {
		return utils.NewStringSet([]string{}), []common.FileInfo{}
	}

	consolidators.RegisterGenericConsolidator(sourceType, listType, testConsolidateFunc)

	consolidator, exists := consolidators.Consolidators.GetConsolidator(sourceType, listType)
	assert.True(t, exists, "Generic consolidator should be registered")
	assert.Equal(t, sourceType, consolidator.GetSourceType())
	assert.Equal(t, listType, consolidator.GetListType())

	sourceType = "test-generic-multi"
	listTypes := []string{"blocklist", "allowlist"}

	consolidators.RegisterGenericConsolidatorTypes(sourceType, listTypes, testConsolidateFunc)

	for _, lt := range listTypes {
		cons, exists := consolidators.Consolidators.GetConsolidator(sourceType, lt)
		assert.True(t, exists, "Generic consolidator should exist for "+lt)
		assert.Equal(t, sourceType, cons.GetSourceType())
		assert.Equal(t, lt, cons.GetListType())
	}
}

// TestCommonConsolidator tests the CommonConsolidator functionality
func TestCommonConsolidator(t *testing.T) {
	sourceType := constants.SourceTypeDomain
	listType := constants.ListTypeBlocklist

	cc := consolidators.NewCommonConsolidator(sourceType, listType)

	assert.Equal(t, sourceType, cc.GetSourceType())
	assert.Equal(t, listType, cc.GetListType())

	originalRegistry := consolidators.Consolidators
	consolidators.Consolidators = consolidators.NewConsolidatorRegistry()
	defer func() {
		consolidators.Consolidators = originalRegistry
	}()

	consolidators.InitForTesting()

	commonSourceTypes := []string{
		constants.SourceTypeIpv4,
		constants.SourceTypeIpv6,
		constants.SourceTypeCidrIpv4,
		constants.SourceTypeDomain,
	}

	commonListTypes := []string{
		constants.ListTypeAllowlist,
		constants.ListTypeBlocklist,
	}

	for _, st := range commonSourceTypes {
		for _, lt := range commonListTypes {
			_, exists := consolidators.Consolidators.GetConsolidator(st, lt)
			assert.True(t, exists, "Common consolidator should be registered for "+st+":"+lt)
		}
	}
}
