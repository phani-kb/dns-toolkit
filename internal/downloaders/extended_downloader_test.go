package downloaders

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanSkipDownload_WithModifiedHeader(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	testFile := createTestFile(t, testDir, "existing.txt", "test content")
	fileInfo, err := os.Stat(testFile)
	require.NoError(t, err)
	fileSize := fileInfo.Size()

	// Use a time in the past for local file
	pastTime := time.Now().Add(-24 * time.Hour)
	require.NoError(t, os.Chtimes(testFile, pastTime, pastTime))

	// Get updated file info after changing the timestamp
	fileInfo, err = os.Stat(testFile)
	require.NoError(t, err)
	localModTime := fileInfo.ModTime()

	// Server returns older last-modified time (should skip)
	serverModTime := pastTime.Add(-time.Hour)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
			w.Header().Set("Last-Modified", serverModTime.Format(http.TimeFormat))
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   testDir,
		Filename: "existing.txt",
	}

	// Setup the summary directory
	summaryDir := filepath.Join(testDir, "summary")
	err = os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	// Save original and restore after test
	origSummaryDir := constants.SummaryDir
	constants.SummaryDir = summaryDir
	defer func() {
		constants.SummaryDir = origSummaryDir
	}()

	constants.DefaultSummaryFiles = map[string]string{
		"download": "download_summary.json",
	}

	client := &http.Client{}
	canSkip := d.CanSkipDownloadForTest(logger, client, "test-agent", file, fileSize, localModTime)
	assert.True(t, canSkip, "Should skip download when server has older modification time")
}

func TestCanSkipDownload_NoLastModifiedHeader(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	testFile := createTestFile(t, testDir, "existing.txt", "test content")
	fileInfo, err := os.Stat(testFile)
	require.NoError(t, err)
	fileSize := fileInfo.Size()
	localModTime := fileInfo.ModTime()

	// Server returns only content length, no last-modified
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   testDir,
		Filename: "existing.txt",
	}

	// Setup the summary directory
	summaryDir := filepath.Join(testDir, "summary")
	err = os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	// Save original and restore after test
	origSummaryDir := constants.SummaryDir
	constants.SummaryDir = summaryDir
	defer func() {
		constants.SummaryDir = origSummaryDir
	}()

	constants.DefaultSummaryFiles = map[string]string{
		"download": "download_summary.json",
	}

	client := &http.Client{}
	canSkip := d.CanSkipDownloadForTest(logger, client, "test-agent", file, fileSize, localModTime)

	// Should skip because file sizes match and no Last-Modified header
	assert.True(t, canSkip, "Should skip download when sizes match and no Last-Modified header")
}

func TestCanSkipDownload_DifferentFileSize(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	testFile := createTestFile(t, testDir, "existing.txt", "test content")
	fileInfo, err := os.Stat(testFile)
	require.NoError(t, err)
	localFileSize := fileInfo.Size()
	localModTime := fileInfo.ModTime()

	// Server returns different content length
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", localFileSize+100)) // Different size
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   testDir,
		Filename: "existing.txt",
	}

	// Setup the summary directory
	summaryDir := filepath.Join(testDir, "summary")
	err = os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	// Save original and restore after test
	origSummaryDir := constants.SummaryDir
	constants.SummaryDir = summaryDir
	defer func() {
		constants.SummaryDir = origSummaryDir
	}()

	constants.DefaultSummaryFiles = map[string]string{
		"download": "download_summary.json",
	}

	client := &http.Client{}
	canSkip := d.CanSkipDownloadForTest(logger, client, "test-agent", file, localFileSize, localModTime)

	// Should not skip because file sizes don't match
	assert.False(t, canSkip, "Should not skip download when file sizes don't match")
}

func TestDownloadFile_ForceFlagSet(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	contentBefore := "old content"
	contentAfter := "new content"
	testFile := createTestFile(t, testDir, "force_test.txt", contentBefore)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, contentAfter)
		if err != nil {
			t.Logf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   testDir,
		Filename: "force_test.txt",
	}

	// Remove the file before force download to ensure overwrite
	if err := os.Remove(testFile); err != nil && !os.IsNotExist(err) {
		t.Logf("Failed to remove test file: %v", err)
	}

	filePath, exists, err := d.Download(logger, file, true, nil, config.ApplicationConfig{})
	assert.NoError(t, err)
	assert.False(t, exists, "Should indicate file was downloaded")
	assert.Equal(t, filepath.Join(testDir, "force_test.txt"), filePath)

	content, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, contentAfter, string(content), "File content should be updated when force=true")
}

// Test downloading with various connection issues
func TestDownloadFile_ConnectionErrors(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	// Test with a connection-refused scenario (invalid port)
	d := NewDefaultDownloaderForTesting(1, 10*time.Millisecond)
	file := c.DownloadFile{
		URL:      "http://localhost:1", // Using port 1 which should be unavailable
		Folder:   destDir,
		Filename: "connection_error.txt",
	}

	_, _, err = d.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.Error(t, err, "Should return error for connection failures")
	assert.Contains(t, err.Error(), "connection", "Error should mention connection issue")
}

// Test skipped with zero size files, closed connections, etc.
func TestDownloadFile_EdgeCases(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(testDir)

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	t.Run("ZeroSizeFile", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "0")
			// Return empty content
		}))
		defer server.Close()

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      server.URL,
			Folder:   destDir,
			Filename: "empty.txt",
		}

		filePath, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.NoError(t, err, "Should successfully download zero-size file")

		info, err := os.Stat(filePath)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), info.Size(), "File should be zero bytes")
	})

	t.Run("EmptyDirectory", func(t *testing.T) {
		t.Parallel()
		emptyDir := filepath.Join(testDir, "empty_dir")
		err := os.MkdirAll(emptyDir, 0755)
		assert.NoError(t, err)

		contentStr := "test content"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(contentStr))
			if err != nil {
				t.Logf("Failed to write response: %v", err)
			}
		}))
		defer server.Close()

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      server.URL,
			Folder:   emptyDir,
			Filename: "file_in_empty_dir.txt",
		}

		filePath, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.NoError(t, err, "Should download to empty directory")

		content, err := os.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, contentStr, string(content), "File content should match")
	})
}
