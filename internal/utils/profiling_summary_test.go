package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartProfiling(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_profiling_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name string
		opts ProfileOptions
	}{
		{
			name: "No profiling enabled",
			opts: ProfileOptions{
				CPUProfile:       false,
				MemProfile:       false,
				GoroutineProfile: false,
				BlockProfile:     false,
			},
		},
		{
			name: "CPU profiling enabled",
			opts: ProfileOptions{
				CPUProfile:      true,
				OutputDir:       tempDir,
				ProfileNameBase: "test_cpu",
			},
		},
		{
			name: "Memory profiling enabled",
			opts: ProfileOptions{
				MemProfile:      true,
				OutputDir:       tempDir,
				ProfileNameBase: "test_mem",
			},
		},
		{
			name: "Block profiling enabled",
			opts: ProfileOptions{
				BlockProfile:     true,
				BlockProfileRate: 1,
				OutputDir:        tempDir,
				ProfileNameBase:  "test_block",
			},
		},
		{
			name: "All profiling enabled",
			opts: ProfileOptions{
				CPUProfile:       true,
				MemProfile:       true,
				GoroutineProfile: true,
				BlockProfile:     true,
				BlockProfileRate: 1,
				OutputDir:        tempDir,
				ProfileNameBase:  "test_all",
			},
		},
		{
			name: "Default profile name",
			opts: ProfileOptions{
				CPUProfile: true,
				OutputDir:  tempDir,
			},
		},
		{
			name: "Invalid output directory",
			opts: ProfileOptions{
				CPUProfile: true,
				OutputDir:  "/invalid/directory/path",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stopFunc := StartProfiling(logger, tt.opts)
			assert.NotNil(t, stopFunc)

			time.Sleep(10 * time.Millisecond)

			stopFunc()
		})
	}
}

func TestAnalyzeProfiles(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_analyze_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	err = AnalyzeProfiles(logger, ProfileOptions{OutputDir: "/non/existent/directory"})
	assert.NoError(t, err)

	err = AnalyzeProfiles(logger, ProfileOptions{OutputDir: tempDir})
	assert.NoError(t, err)

	cpuProfile := filepath.Join(tempDir, "cpu.prof")
	memProfile := filepath.Join(tempDir, "mem.prof")

	err = os.WriteFile(cpuProfile, []byte("fake cpu profile data"), 0644)
	require.NoError(t, err)

	err = os.WriteFile(memProfile, []byte("fake memory profile data"), 0644)
	require.NoError(t, err)

	err = AnalyzeProfiles(logger, ProfileOptions{OutputDir: tempDir})
	assert.NoError(t, err)
}

func TestSaveSummary(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_save_summary_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	summaryFile := filepath.Join(tempDir, "test_summary.json")

	summary := c.ProcessedSummary{
		Name: "test-source",
		ValidFiles: []c.ProcessedFile{
			{GenericSourceType: "domain"},
		},
		InvalidFiles: []c.ProcessedFile{
			{GenericSourceType: "ip"},
		},
	}

	lessFunc := func(i, j c.ProcessedSummary) bool {
		return i.Name < j.Name
	}

	count, err := SaveSummary(logger, summary, summaryFile, lessFunc)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	assert.FileExists(t, summaryFile)

	content, err := os.ReadFile(summaryFile)
	require.NoError(t, err)

	var savedSummaries []c.ProcessedSummary
	err = json.Unmarshal(content, &savedSummaries)
	require.NoError(t, err)

	assert.Len(t, savedSummaries, 1)
	assert.Equal(t, "test-source", savedSummaries[0].Name)
}

func TestSaveSummaries(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_save_summaries_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	summaryFile := filepath.Join(tempDir, "test_summaries.json")

	summaries := []c.ProcessedSummary{
		{
			Name: "source-b",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
			},
		},
		{
			Name: "source-a",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "ip"},
			},
		},
	}

	lessFunc := func(i, j c.ProcessedSummary) bool {
		return i.Name < j.Name
	}

	count, err := SaveSummaries(logger, summaries, summaryFile, lessFunc)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	content, err := os.ReadFile(summaryFile)
	require.NoError(t, err)

	var savedSummaries []c.ProcessedSummary
	err = json.Unmarshal(content, &savedSummaries)
	require.NoError(t, err)

	assert.Len(t, savedSummaries, 2)
	assert.Equal(t, "source-a", savedSummaries[0].Name)
	assert.Equal(t, "source-b", savedSummaries[1].Name)

	emptyFile := filepath.Join(tempDir, "empty_summaries.json")
	count, err = SaveSummaries(logger, []c.ProcessedSummary{}, emptyFile, lessFunc)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
	testBackupFile := filepath.Join(tempDir, "backup_test.json")

	count, err = SaveSummaries(logger, summaries, testBackupFile, lessFunc)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	originalModTime := time.Now().Add(-time.Hour)
	err = os.Chtimes(testBackupFile, originalModTime, originalModTime)
	require.NoError(t, err)

	count, err = SaveSummaries(logger, summaries, testBackupFile, lessFunc)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	backupPattern := "backup_test_*.json"
	if matches, err := filepath.Glob(filepath.Join("data/backup", backupPattern)); err == nil {
		for _, match := range matches {
			os.Remove(match)
		}
	}

	invalidFile := "/non/existent/dir/summaries.json"
	count, err = SaveSummaries(logger, summaries, invalidFile, lessFunc)
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}

func TestGetFilesFromSummaries(t *testing.T) {
	t.Parallel()

	summaries := []c.ProcessedSummary{
		{
			Name: "source1",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
				{GenericSourceType: "ip"},
			},
			InvalidFiles: []c.ProcessedFile{
				{GenericSourceType: "domain"},
			},
		},
		{
			Name: "source2",
			ValidFiles: []c.ProcessedFile{
				{GenericSourceType: "url"},
			},
		},
	}

	files := GetFilesFromSummaries(summaries, "processed")
	assert.IsType(t, map[string]c.ProcessedSummary{}, files)
}

func TestGetFilePathForSummary(t *testing.T) {
	t.Parallel()

	summary := c.ProcessedSummary{
		Name: "test-source",
	}

	summaryType := "processed"

	filePath := getFilePathForSummary(summary, summaryType)

	assert.IsType(t, "", filePath)
}
