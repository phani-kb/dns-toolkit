package top

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	entrySources     map[string]map[uint16]struct{}
	expectedResults  [][]common.EntryCountPair
	testSummaries    []common.TopSummary
	expectedFiltered []common.TopSummary
	testEntries      []common.EntryCountPair
}

func getTestData() testData {
	entrySources := map[string]map[uint16]struct{}{
		"entry1": {0: {}, 1: {}, 2: {}},                      // 3 sources
		"entry2": {0: {}, 1: {}, 2: {}, 3: {}},               // 4 sources
		"entry3": {0: {}, 1: {}},                             // 2 sources
		"entry4": {0: {}, 1: {}, 2: {}, 3: {}, 4: {}},        // 5 sources
		"entry5": {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}}, // 6 sources
	}

	expectedResults := [][]common.EntryCountPair{
		{
			{Entry: "entry5", Count: 6},
			{Entry: "entry4", Count: 5},
		},
		{
			{Entry: "entry5", Count: 6},
			{Entry: "entry4", Count: 5},
			{Entry: "entry2", Count: 4},
			{Entry: "entry1", Count: 3},
		},
		{},
	}

	testSummaries := []common.TopSummary{
		{
			GenericSourceType: "domain",
			ListType:          "blocklist",
			Count:             0, // Empty should be filtered out
			TopEntries:        []common.EntryCountPair{},
		},
		{
			GenericSourceType: "domain",
			ListType:          "allowlist",
			Count:             5,
			TopEntries: []common.EntryCountPair{
				{Entry: "test1", Count: 5},
				{Entry: "test2", Count: 4},
			},
		},
		{
			GenericSourceType: "ipv4",
			ListType:          "blocklist",
			Count:             3,
			TopEntries: []common.EntryCountPair{
				{Entry: "192.0.2.1", Count: 5},
				{Entry: "192.0.2.2", Count: 4},
				{Entry: "192.0.2.3", Count: 3},
			},
		},
	}

	expectedFiltered := []common.TopSummary{
		{
			GenericSourceType: "domain",
			ListType:          "allowlist",
			Count:             5,
			TopEntries: []common.EntryCountPair{
				{Entry: "test1", Count: 5},
				{Entry: "test2", Count: 4},
			},
		},
		{
			GenericSourceType: "ipv4",
			ListType:          "blocklist",
			Count:             3,
			TopEntries: []common.EntryCountPair{
				{Entry: "192.0.2.1", Count: 5},
				{Entry: "192.0.2.2", Count: 4},
				{Entry: "192.0.2.3", Count: 3},
			},
		},
	}

	testEntries := []common.EntryCountPair{
		{Entry: "test1.com", Count: 5},
		{Entry: "test2.com", Count: 4},
		{Entry: "test3.com", Count: 3},
	}

	return testData{
		entrySources:     entrySources,
		expectedResults:  expectedResults,
		testSummaries:    testSummaries,
		expectedFiltered: expectedFiltered,
		testEntries:      testEntries,
	}
}

func setup(t *testing.T) (*multilog.Logger, string, string) {
	logger, _ := multilog.NewTestLogger(t)

	testDir, err := os.MkdirTemp("", "top_test_*")
	require.NoError(t, err)

	topDir := filepath.Join(testDir, "top")
	summaryDir := filepath.Join(testDir, "summary")

	err = os.MkdirAll(topDir, 0755)
	require.NoError(t, err)

	err = os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	return logger, topDir, summaryDir
}

func cleanup(t *testing.T, testDir string) {
	if testDir != "" {
		err := os.RemoveAll(testDir)
		if err != nil {
			t.Logf("Failed to clean up test directory: %v", err)
		}
	}
}

func createTestProcessedFiles(t *testing.T, testDir, gst, listType string, numSources int) []common.ProcessedFile {
	var files []common.ProcessedFile

	for i := 0; i < numSources; i++ {
		filename := filepath.Join(testDir, fmt.Sprintf("source_%d.txt", i))

		f, err := os.Create(filename)
		require.NoError(t, err)

		// Each file has some unique entries and some shared entries
		for j := i; j < i+10; j++ {
			_, err := fmt.Fprintf(f, "entry_%d\n", j)
			require.NoError(t, err)
		}

		err = f.Close()
		require.NoError(t, err)

		files = append(files, common.ProcessedFile{
			Name:              fmt.Sprintf("source_%d", i),
			GenericSourceType: gst,
			ListType:          listType,
			Filepath:          filename,
			Valid:             true,
		})
	}

	return files
}

func TestGetTopNEntries(t *testing.T) {
	data := getTestData()

	testCases := []struct {
		name     string
		minSrc   int
		maxEnt   int
		expected []common.EntryCountPair
	}{
		{"top_2", 3, 2, data.expectedResults[0]},
		{"top_10", 3, 10, data.expectedResults[1]},
		{"empty", 10, 10, data.expectedResults[2]},
	}

	_, topDir, summaryDir := setup(t)
	defer cleanup(t, filepath.Dir(topDir))

	service := NewDefaultService(topDir, summaryDir)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := service.GetTopNEntries(data.entrySources, tc.minSrc, tc.maxEnt)

			assert.Equal(t, len(tc.expected), len(results))

			for i, expected := range tc.expected {
				if i < len(results) {
					assert.Equal(t, expected.Entry, results[i].Entry)
					assert.Equal(t, expected.Count, results[i].Count)
				}
			}
		})
	}
}

func TestGetTopNEntries_DefaultImpl(t *testing.T) {
	_, topDir, summaryDir := setup(t)
	defer cleanup(t, filepath.Dir(topDir))

	service := NewDefaultService(topDir, summaryDir)

	testData := getTestData()

	entrySources := testData.entrySources

	{
		results := service.GetTopNEntries(entrySources, 3, 2)

		assert.Len(t, results, 2)
		assert.Equal(t, "entry5", results[0].Entry) // Most sources (6)
		assert.Equal(t, 6, results[0].Count)
		assert.Equal(t, "entry4", results[1].Entry) // Second most sources (5)
		assert.Equal(t, 5, results[1].Count)
	}

	{
		results := service.GetTopNEntries(entrySources, 3, 10)

		assert.Len(t, results, 4)
		// Should be sorted in descending order by count
		assert.Equal(t, "entry5", results[0].Entry) // 6 sources
		assert.Equal(t, "entry4", results[1].Entry) // 5 sources
		assert.Equal(t, "entry2", results[2].Entry) // 4 sources
		assert.Equal(t, "entry1", results[3].Entry) // 3 sources
	}

	{
		results := service.GetTopNEntries(entrySources, 10, 10)
		assert.Empty(t, results)
	}
}

func TestFilterTopSummaries(t *testing.T) {
	data := getTestData()

	_, topDir, summaryDir := setup(t)
	defer cleanup(t, filepath.Dir(topDir))

	service := NewDefaultService(topDir, summaryDir)
	filtered := service.FilterTopSummaries(data.testSummaries)

	assert.Len(t, filtered, 2)
	assert.Equal(t, "allowlist", filtered[0].ListType)
	assert.Equal(t, "blocklist", filtered[1].ListType)
	assert.Equal(t, 5, filtered[0].Count)
	assert.Equal(t, 3, filtered[1].Count)
}

func TestSaveTopEntries(t *testing.T) {
	data := getTestData()
	logger, topDir, summaryDir := setup(t)
	defer cleanup(t, filepath.Dir(topDir))

	service := NewDefaultService(topDir, summaryDir)

	filePath, err := service.SaveTopEntries(logger, "domain", "blocklist", 3, data.testEntries)
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	assert.NoError(t, err)

	content, err := os.ReadFile(filePath)
	require.NoError(t, err)

	expectedContent := "test1.com\ntest2.com\ntest3.com\n"
	assert.Equal(t, expectedContent, string(content))

	filePath, err = service.SaveTopEntries(logger, "domain", "blocklist", 3, []common.EntryCountPair{})
	assert.NoError(t, err)
	assert.Empty(t, filePath)
}

func TestSaveTopSummaries(t *testing.T) {
	logger, topDir, summaryDir := setup(t)
	defer cleanup(t, filepath.Dir(topDir))

	summaries := []common.TopSummary{
		{
			GenericSourceType: "domain",
			ListType:          "blocklist",
			MinSources:        3,
			Count:             2,
			Filepath:          filepath.Join(topDir, "top_domain_blocklist_min3.txt"),
			TopEntries: []common.EntryCountPair{
				{Entry: "test1.com", Count: 5},
				{Entry: "test2.com", Count: 4},
			},
		},
	}

	expectedCount := 1

	service := NewDefaultService(topDir, summaryDir)

	count, err := service.SaveTopSummaries(logger, summaries)
	require.NoError(t, err)
	assert.Equal(t, expectedCount, count)

	summaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles["top"])
	_, err = os.Stat(summaryFile)
	assert.NoError(t, err)
}

func TestFindTopEntries(t *testing.T) {
	logger, topDir, summaryDir := setup(t)
	testDir := filepath.Dir(topDir)
	defer cleanup(t, testDir)

	gst := "domain"
	listType := "blocklist"
	processedFiles := createTestProcessedFiles(t, testDir, gst, listType, 5)
	stringPool := utils.NewDTEntryPool()

	validateSummary := func(t *testing.T, summary common.TopSummary) {
		assert.Equal(t, gst, summary.GenericSourceType)
		assert.Equal(t, listType, summary.ListType)
		assert.Equal(t, 2, summary.MinSources)
		assert.NotEmpty(t, summary.Filepath)
		assert.Greater(t, summary.Count, 0)
		assert.Equal(t, summary.Count, len(summary.TopEntries))

		_, err := os.Stat(summary.Filepath)
		assert.NoError(t, err)
	}

	service := NewDefaultService(topDir, summaryDir)

	summary, err := service.FindTopEntries(
		logger,
		gst,
		listType,
		processedFiles,
		2,  // minSources
		10, // maxEntries
		stringPool,
	)

	require.NoError(t, err)
	validateSummary(t, summary)
}

func TestFindTopEntries_DefaultImpl(t *testing.T) {
	logger, topDir, summaryDir := setup(t)
	testDir := filepath.Dir(topDir)
	defer cleanup(t, testDir)

	service := NewDefaultService(topDir, summaryDir)

	gst := "domain"
	listType := "blocklist"
	processedFiles := createTestProcessedFiles(t, testDir, gst, listType, 5)

	stringPool := utils.NewDTEntryPool()

	summary, err := service.FindTopEntries(
		logger,
		gst,
		listType,
		processedFiles,
		2,  // minSources
		10, // maxEntries
		stringPool,
	)

	require.NoError(t, err)

	assert.Equal(t, gst, summary.GenericSourceType)
	assert.Equal(t, listType, summary.ListType)
	assert.Equal(t, 2, summary.MinSources)
	assert.NotEmpty(t, summary.Filepath)
	assert.Greater(t, summary.Count, 0)
	assert.Equal(t, summary.Count, len(summary.TopEntries))

	_, err = os.Stat(summary.Filepath)
	assert.NoError(t, err)
}

func TestProcessTopEntries(t *testing.T) {
	logger, topDir, summaryDir := setup(t)
	testDir := filepath.Dir(topDir)
	defer cleanup(t, testDir)

	gsts := []string{"domain", "ipv4"}
	minSources := []int{1, 2}

	files := append(
		createTestProcessedFiles(t, testDir, "domain", "blocklist", 3),
		createTestProcessedFiles(t, testDir, "ipv4", "blocklist", 2)...,
	)

	validateSummaries := func(t *testing.T, summaries []common.TopSummary) {
		assert.NotEmpty(t, summaries)
		foundDomain := false
		foundIPv4 := false
		for _, summary := range summaries {
			if summary.GenericSourceType == "domain" {
				foundDomain = true
			}
			if summary.GenericSourceType == "ipv4" {
				foundIPv4 = true
			}
			assert.NotEmpty(t, summary.Filepath)
			assert.Greater(t, summary.Count, 0)
			assert.NotEmpty(t, summary.TopEntries)
			validMinSource := false
			for _, ms := range minSources {
				if summary.MinSources == ms {
					validMinSource = true
					break
				}
			}
			assert.True(t, validMinSource)

			_, err := os.Stat(summary.Filepath)
			assert.NoError(t, err)
		}
		assert.True(t, foundDomain)
		assert.True(t, foundIPv4)
	}

	service := NewDefaultService(topDir, summaryDir)

	summaries, err := service.ProcessTopEntries(
		logger,
		gsts,
		files,
		minSources,
		5, // maxEntries
		2, // maxWorkers
	)

	require.NoError(t, err)
	validateSummaries(t, summaries)

	summaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles["top"])
	_, err = os.Stat(summaryFile)
	assert.NoError(t, err)
}
