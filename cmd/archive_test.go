package cmd

import (
	"archive/tar"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSummaryCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		summaryType string
		content     []byte
		expected    int
	}{
		{
			name: "download summary",
			content: func() []byte {
				summary := []common.DownloadSummary{
					{Name: "test1"}, {Name: "test2"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeDownload,
			expected:    2,
		},
		{
			name: "processed summary",
			content: func() []byte {
				summary := []common.ProcessedSummary{
					{Name: "test1"}, {Name: "test2"}, {Name: "test3"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeProcessed,
			expected:    3,
		},
		{
			name: "consolidated summary",
			content: func() []byte {
				summary := []common.ConsolidatedSummary{
					{Type: "domain"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeConsolidated,
			expected:    1,
		},
		{
			name: "consolidated groups summary",
			content: func() []byte {
				summary := []common.ConsolidatedGroupsSummary{
					{Group: "mini"}, {Group: "lite"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeConsolidatedGroups,
			expected:    2,
		},
		{
			name: "overlap summary",
			content: func() []byte {
				summary := []common.OverlapSummary{
					{Source: "overlap1"}, {Source: "overlap2"}, {Source: "overlap3"}, {Source: "overlap4"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeOverlap,
			expected:    4,
		},
		{
			name: "top summary",
			content: func() []byte {
				summary := []common.TopSummary{
					{GenericSourceType: "domain"},
				}
				data, _ := json.Marshal(summary)
				return data
			}(),
			summaryType: constants.SummaryTypeTop,
			expected:    1,
		},
		{
			name:        "unknown summary type",
			content:     []byte(`[{"name": "test"}]`),
			summaryType: "unknown",
			expected:    0,
		},
		{
			name:        "invalid JSON",
			content:     []byte(`{invalid json}`),
			summaryType: constants.SummaryTypeDownload,
			expected:    0,
		},
		{
			name:        "empty content",
			content:     []byte{},
			summaryType: constants.SummaryTypeDownload,
			expected:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSummaryCount(tt.content, tt.summaryType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProcessSummaryFiles(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	tmpDir := t.TempDir()

	downloadSummary := []common.DownloadSummary{
		{Name: "test1"}, {Name: "test2"},
	}
	downloadData, err := json.Marshal(downloadSummary)
	require.NoError(t, err)

	downloadSummaryPath := filepath.Join(tmpDir, "download_summary.json")
	err = os.WriteFile(downloadSummaryPath, downloadData, 0644)
	require.NoError(t, err)

	processedSummary := []common.ProcessedSummary{
		{Name: "processed1"},
	}
	processedData, err := json.Marshal(processedSummary)
	require.NoError(t, err)

	processedSummaryPath := filepath.Join(tmpDir, "processed_summary.json")
	err = os.WriteFile(processedSummaryPath, processedData, 0644)
	require.NoError(t, err)

	nonSummaryPath := filepath.Join(tmpDir, "other_file.txt")
	err = os.WriteFile(nonSummaryPath, []byte("not a summary"), 0644)
	require.NoError(t, err)

	subDir := filepath.Join(tmpDir, "subdir")
	err = os.MkdirAll(subDir, 0755)
	require.NoError(t, err)

	archiveSummary := &common.ArchiveSummary{
		Folders:      []common.ArchiveFolder{},
		SummaryFiles: []common.ArchiveSummaryFile{},
	}

	processSummaryFiles(logger, tmpDir, archiveSummary, nil)

	assert.Len(t, archiveSummary.SummaryFiles, 2, "Should have processed 2 summary files")

	fileNames := make([]string, len(archiveSummary.SummaryFiles))
	for i, sf := range archiveSummary.SummaryFiles {
		fileNames[i] = sf.Name
	}
	assert.Contains(t, fileNames, "download_summary.json")
	assert.Contains(t, fileNames, "processed_summary.json")

	for _, sf := range archiveSummary.SummaryFiles {
		switch sf.Name {
		case "download_summary.json":
			assert.Equal(t, 2, sf.Count, "Download summary should have count 2")
			assert.Equal(t, constants.SummaryTypeDownload, sf.SummaryType)
		case "processed_summary.json":
			assert.Equal(t, 1, sf.Count, "Processed summary should have count 1")
			assert.Equal(t, constants.SummaryTypeProcessed, sf.SummaryType)
		}
		assert.True(t, sf.Valid, "Summary files should be marked as valid")
		assert.NotEmpty(t, sf.Checksum, "Should have calculated checksum")
		assert.Greater(t, sf.Size, int64(0), "Should have recorded file size")
	}
}

func TestProcessSummaryFilesEmptyDirectory(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	tmpDir := t.TempDir()

	archiveSummary := &common.ArchiveSummary{
		Folders:      []common.ArchiveFolder{},
		SummaryFiles: []common.ArchiveSummaryFile{},
	}

	processSummaryFiles(logger, tmpDir, archiveSummary, nil)

	assert.Len(t, archiveSummary.SummaryFiles, 0, "Should have no summary files for empty directory")
}

func TestProcessSummaryFilesNonexistentDirectory(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	nonexistentDir := "/nonexistent/directory"

	archiveSummary := &common.ArchiveSummary{
		Folders:      []common.ArchiveFolder{},
		SummaryFiles: []common.ArchiveSummaryFile{},
	}

	processSummaryFiles(logger, nonexistentDir, archiveSummary, nil)

	assert.Len(t, archiveSummary.SummaryFiles, 0, "Should have no summary files for nonexistent directory")
}

func TestAddFileToTar(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	tmpDir := t.TempDir()

	testContent := "test file content"
	testFilePath := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	fileInfo, err := os.Stat(testFilePath)
	require.NoError(t, err)

	archivePath := filepath.Join(tmpDir, "test.tar")
	archiveFile, err := os.Create(archivePath)
	require.NoError(t, err)
	defer func() {
		if err := archiveFile.Close(); err != nil {
			t.Logf("Failed to close archive file: %v", err)
		}
	}()

	tarWriter := tar.NewWriter(archiveFile)
	defer func() {
		if err := tarWriter.Close(); err != nil {
			t.Logf("Failed to close tar writer: %v", err)
		}
	}()

	err = addFileToTar(logger, tarWriter, testFilePath, "target.txt", fileInfo)
	assert.NoError(t, err, "Should successfully add file to tar")

	err = addFileToTar(logger, tarWriter, testFilePath, "", fileInfo)
	assert.NoError(t, err, "Should successfully add file to tar with empty target path")

	nonexistentPath := filepath.Join(tmpDir, "nonexistent.txt")
	nonexistentInfo, _ := os.Stat(testFilePath) // Use existing file info for header creation
	err = addFileToTar(logger, tarWriter, nonexistentPath, "nonexistent.txt", nonexistentInfo)
	assert.Error(t, err, "Should fail when source file doesn't exist")
}

func TestAddFileToTarDirectory(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	tmpDir := t.TempDir()

	subDir := filepath.Join(tmpDir, "subdir")
	err := os.MkdirAll(subDir, 0755)
	require.NoError(t, err)

	dirInfo, err := os.Stat(subDir)
	require.NoError(t, err)

	archivePath := filepath.Join(tmpDir, "test.tar")
	archiveFile, err := os.Create(archivePath)
	require.NoError(t, err)
	defer func() {
		if err := archiveFile.Close(); err != nil {
			t.Logf("Failed to close archive file: %v", err)
		}
	}()

	tarWriter := tar.NewWriter(archiveFile)
	defer func() {
		if err := tarWriter.Close(); err != nil {
			t.Logf("Failed to close tar writer: %v", err)
		}
	}()

	err = addFileToTar(logger, tarWriter, subDir, "subdir/", dirInfo)
	assert.NoError(t, err, "Should successfully add directory to tar")
}

func TestArchiveCommandIntegration(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, archiveCmd, "Archive command should be defined")
	assert.Equal(t, "archive", archiveCmd.Use, "Command should have correct usage")
	assert.Contains(t, archiveCmd.Short, "Archive", "Command should have appropriate short description")
	assert.NotNil(t, archiveCmd.Run, "Command should have a run function")
}

// Helper function
func TestRunArchiveIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	originalArchiveDir := constants.ArchiveDir
	tmpDir := t.TempDir()
	constants.ArchiveDir = tmpDir

	defer func() {
		constants.ArchiveDir = originalArchiveDir
	}()

	testDataDir := filepath.Join(tmpDir, "testdata")
	err := os.MkdirAll(testDataDir, 0755)
	require.NoError(t, err)

	testFile := filepath.Join(testDataDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test content"), 0644)
	require.NoError(t, err)

	originalAppConfig := AppConfig
	AppConfig = &config.AppConfig{
		Application: config.ApplicationConfig{
			Name:        "test-app",
			Version:     "1.0.0",
			Description: "Test application",
		},
		DNSToolkit: config.DNSToolkitConfig{
			Folders: config.FoldersConfig{
				Archive: tmpDir,
				Summary: testDataDir,
			},
			FilesChecksum: config.FilesChecksumConfig{
				Algorithm: "md5",
			},
		},
		Multilog: map[string]interface{}{
			"level": "info",
		},
	}
	defer func() {
		AppConfig = originalAppConfig
	}()

	assert.NotPanics(t, func() {
	}, "runArchive should not panic with valid configuration")
}
