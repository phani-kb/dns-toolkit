package downloaders

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func setupTestLogger() *multilog.Logger {
	return multilog.NewLogger()
}

func setupTestDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "downloader_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	return tempDir
}

func createTestFile(t *testing.T, dir, filename, content string) string {
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	return path
}

// Create a test downloader with a short retry delay
func newTestDownloader(maxRetries int) *DefaultDownloader {
	return NewDefaultDownloaderForTesting(maxRetries, 10*time.Millisecond)
}

func TestDefaultDownloaderName(t *testing.T) {
	t.Parallel()
	d := NewDefaultDownloaderWithRetries(1)
	assert.Equal(t, defaultDownloaderName, d.Name())
	assert.Equal(t, defaultDownloaderName, DefaultDownloaderName())
}

func TestDefaultDownloader_CopyLocalFile(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Create a source file
	sourceContent := "test content"
	sourceFilename := "source.txt"
	sourcePath := createTestFile(t, testDir, sourceFilename, sourceContent)
	sourceURL := fmt.Sprintf("file://%s", sourcePath)

	// Create the destination directory
	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	// Test copy local file
	d := newTestDownloader(1)
	destPath, exists, err := d.copyLocalFile(logger, sourceURL, destDir, "destination.txt")

	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, filepath.Join(destDir, "destination.txt"), destPath)

	content, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Equal(t, sourceContent, string(content))

	destPath, exists, err = d.copyLocalFile(logger, sourceURL, destDir, "destination.txt")
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, filepath.Join(destDir, "destination.txt"), destPath)
}

func TestDefaultDownloader_Download_LocalFile(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	sourceContent := "test content"
	sourceFilename := "source.txt"
	sourcePath := createTestFile(t, testDir, sourceFilename, sourceContent)
	sourceURL := fmt.Sprintf("file://%s", sourcePath)

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	d := newTestDownloader(3)
	file := c.DownloadFile{
		URL:      sourceURL,
		Folder:   destDir,
		Filename: "destination.txt",
	}

	destPath, exists, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, filepath.Join(destDir, "destination.txt"), destPath)

	content, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Equal(t, sourceContent, string(content))
}

func TestDefaultDownloader_Download_RemoteFile(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Setup test server
	content := "test file content"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Last-Modified", time.Now().Format(time.RFC1123))
		_, err := fmt.Fprintln(w, content)
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	d := newTestDownloader(3)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "remote.txt",
	}

	destPath, exists, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, filepath.Join(destDir, "remote.txt"), destPath)

	downloadedContent, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Contains(t, string(downloadedContent), content)
}

func TestDefaultDownloader_Download_HTTPError(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Setup test server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	d := newTestDownloader(1) // Use fast test downloader
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "error.txt",
	}

	_, _, err = d.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.Error(t, err)
	var httpErr *HTTPStatusError
	assert.ErrorAs(t, err, &httpErr)
	assert.Equal(t, http.StatusNotFound, httpErr.StatusCode)
}

func TestDefaultDownloader_PostDownloadProcess(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	d := newTestDownloader(1) // Use fast test downloader

	// Default implementation returns nil
	err := d.PostDownloadProcess(logger, "test_path", 10)
	assert.NoError(t, err)
}

func TestDefaultDownloader_RetryLogic(t *testing.T) {
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Setup test server that fails first then succeeds
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		if _, err := fmt.Fprintln(w, "success after retry"); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	// Test download with retry
	d := newTestDownloader(2)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "retry.txt",
	}

	destPath, exists, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Equal(t, filepath.Join(destDir, "retry.txt"), destPath)
	assert.Equal(t, 2, attemptCount)

	downloadedContent, err := os.ReadFile(destPath)
	assert.NoError(t, err)
	assert.Contains(t, string(downloadedContent), "success after retry")
}

func TestShouldDownload(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	d := newTestDownloader(1)

	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	summaryPath := createTestFile(t, testDir, "summary.json", `{
		"entries": {
			"test_file.txt": {
				"lastModified": "2025-01-01T00:00:00Z",
				"contentLength": 100
			}
		}
	}`)

	file := c.DownloadFile{
		URL:      "http://example.com/test_file.txt",
		Filename: "test_file.txt",
	}

	// Test the ShouldDownload logic
	shouldDownload := d.ShouldDownload(logger, summaryPath, file)
	assert.True(t, shouldDownload, "File should download if force flag is set")
}
