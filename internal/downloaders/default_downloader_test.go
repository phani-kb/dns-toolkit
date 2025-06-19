package downloaders

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper functions for tests
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

// Create a test downloader with a minimal retry delay and timeout settings
func newTestDownloader(maxRetries int) *DefaultDownloader {
	return NewDefaultDownloaderForTesting(maxRetries, 1*time.Millisecond)
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

	sourceContent := "test content"
	sourceFilename := "source.txt"
	sourcePath := createTestFile(t, testDir, sourceFilename, sourceContent)
	sourceURL := fmt.Sprintf("file://%s", sourcePath)

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

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

func TestDefaultDownloader_PostDownloadProcess(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	d := newTestDownloader(1)

	// Default implementation returns nil
	err := d.PostDownloadProcess(logger, "test_path", 10)
	assert.NoError(t, err)
}

func TestDefaultDownloader_RetryLogic(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

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

	shouldDownload := d.ShouldDownload(logger, summaryPath, file)
	assert.True(t, shouldDownload, "File should download if force flag is set")
}

func TestAdditionalShouldDownload(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	_ = createTestFile(t, testDir, "existing.txt", "test content")

	d := NewDefaultDownloaderWithRetries(1)

	file := c.DownloadFile{
		Folder:   testDir,
		Filename: "existing.txt",
	}
	shouldDownload := d.ShouldDownload(logger, "summary.json", file)
	assert.False(t, shouldDownload, "Should not download existing file")

	file = c.DownloadFile{
		Folder:   testDir,
		Filename: "nonexisting.txt",
	}
	shouldDownload = d.ShouldDownload(logger, "summary.json", file)
	assert.True(t, shouldDownload, "Should download non-existing file")
}

func TestCanSkipDownload(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	summaryDir := filepath.Join(testDir, "summary")
	err := os.MkdirAll(summaryDir, 0755)
	require.NoError(t, err)

	origSummaryDir := constants.SummaryDir
	constants.SummaryDir = summaryDir
	defer func() {
		constants.SummaryDir = origSummaryDir
	}()

	constants.DefaultSummaryFiles = map[string]string{
		"download": "download_summary.json",
	}

	testFile := createTestFile(t, testDir, "test.txt", "test content")
	fileInfo, err := os.Stat(testFile)
	require.NoError(t, err)
	fileSize := fileInfo.Size()
	modTime := fileInfo.ModTime()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.Header().Set("Content-Length", "12") // Length of "test content"
			w.Header().Set("Last-Modified", modTime.Format(http.TimeFormat))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      ts.URL,
		Folder:   testDir,
		Filename: "test.txt",
	}

	_, err = url.Parse(file.URL)
	require.NoError(t, err)

	mockDir := filepath.Join(testDir, "mock")
	err = os.MkdirAll(mockDir, 0755)
	require.NoError(t, err)

	mockFilePath := filepath.Join(mockDir, "test.txt")
	err = os.WriteFile(mockFilePath, []byte("test content"), 0644)
	require.NoError(t, err)

	file.Folder = mockDir
	canSkip := d.canSkipDownload(logger, &http.Client{}, "test-agent", file, fileSize, modTime)
	assert.True(t, canSkip, "Should skip download for existing files")

	badServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer badServer.Close()

	file.URL = badServer.URL
	canSkip = d.canSkipDownload(logger, &http.Client{}, "test-agent", file, fileSize, modTime)
	assert.True(t, canSkip, "Should skip download since ShouldDownload returns false for existing files")
}

func TestCanSkipDownloadWithBadRequest(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Setup server that always returns an error
	badServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer badServer.Close()

	d := NewDefaultDownloaderWithRetries(1)

	// Test with non-existent file - can never skip
	nonExistentFile := c.DownloadFile{
		URL:      badServer.URL,
		Folder:   testDir,
		Filename: "does_not_exist.txt",
	}

	canSkip := d.canSkipDownload(logger, &http.Client{}, "test-agent", nonExistentFile, 0, time.Time{})
	assert.False(t, canSkip, "Should not skip download when file doesn't exist")
}

func TestCanSkipDownloadWithErrors(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	d := NewDefaultDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      errorServer.URL,
		Folder:   testDir,
		Filename: "error_test.txt",
	}

	// Create test file to have something to compare against
	testFile := createTestFile(t, testDir, "error_test.txt", "test content")
	fileInfo, err := os.Stat(testFile)
	require.NoError(t, err)

	client := &http.Client{Timeout: 1 * time.Second}
	canSkip := d.canSkipDownload(logger, client, "test-agent", file, fileInfo.Size(), fileInfo.ModTime())
	assert.True(t, canSkip, "Should skip download when HEAD request fails but file exists")

	// Test with non-existent file
	nonExistentFile := c.DownloadFile{
		URL:      errorServer.URL,
		Folder:   testDir,
		Filename: "non_existent.txt",
	}
	canSkip = d.canSkipDownload(logger, client, "test-agent", nonExistentFile, 0, time.Time{})
	assert.False(t, canSkip, "Should not skip download when HEAD fails and file does not exist")
}

func TestHandleArchiveFile(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	regularFile := c.DownloadFile{
		Folder:    testDir,
		Filename:  "regular.txt",
		IsArchive: false,
	}
	filePath := createTestFile(t, testDir, "regular.txt", "test content")

	d := NewDefaultDownloaderWithRetries(1)
	err := d.handleArchiveFile(logger, regularFile, filePath)
	assert.NoError(t, err, "Should handle non-archive files without error")

	archiveFile := c.DownloadFile{
		Folder:    testDir,
		Filename:  "archive.zip",
		IsArchive: true,
		Targets:   []c.DownloadTarget{{SourceFile: "source.txt", TargetFolder: "target"}},
	}
	archivePath := createTestFile(t, testDir, "archive.zip", "fake archive")

	err = d.handleArchiveFile(logger, archiveFile, archivePath)
	assert.Error(t, err, "Should fail with fake archive file")
}

func TestHandleArchiveFileComprehensive(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	// Test with empty targets - but should still error on invalid archive format
	archiveFileWithNoTargets := c.DownloadFile{
		Folder:    testDir,
		Filename:  "empty_targets.zip",
		IsArchive: true,
		Targets:   []c.DownloadTarget{},
	}
	archivePath := createTestFile(t, testDir, "empty_targets.zip", "fake archive")

	d := NewDefaultDownloaderWithRetries(1)
	err := d.handleArchiveFile(logger, archiveFileWithNoTargets, archivePath)
	assert.Error(t, err, "Should error with invalid archive format")

	// Test with archive flag false but targets existing
	nonArchiveWithTargets := c.DownloadFile{
		Folder:    testDir,
		Filename:  "not_archive.txt",
		IsArchive: false,
		Targets:   []c.DownloadTarget{{SourceFile: "source.txt", TargetFolder: "target"}},
	}
	nonArchivePath := createTestFile(t, testDir, "not_archive.txt", "not an archive")

	err = d.handleArchiveFile(logger, nonArchiveWithTargets, nonArchivePath)
	assert.NoError(t, err, "Should handle non-archive file with targets without error")
}

func TestCreateHTTPClient(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	d := NewDefaultDownloaderWithRetries(1)

	parsedURL, err := url.Parse("https://example.com")
	require.NoError(t, err)

	client := d.createHTTPClient(logger, false, nil, parsedURL)
	assert.NotNil(t, client, "Should create a default HTTP client")

	clientWithSkip := d.createHTTPClient(logger, true, nil, parsedURL)
	assert.NotNil(t, clientWithSkip, "Should create a client with TLS verification skipped")

	clientWithHosts := d.createHTTPClient(logger, false, []string{"example.com"}, parsedURL)
	assert.NotNil(t, clientWithHosts, "Should create a client with skipCertHosts")

	clientWithBoth := d.createHTTPClient(logger, true, []string{"example.com"}, parsedURL)
	assert.NotNil(t, clientWithBoth, "Should create a client with both skip options")
}

// ERROR HANDLING AND RETRY TESTS

func TestDefaultDownloader_Download_HTTPError(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	assert.NoError(t, err)

	d := newTestDownloader(1)
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

func TestDownloadEdgeCases_Merged(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	downloadDir := filepath.Join(testDir, "download")
	err := os.MkdirAll(downloadDir, 0755)
	assert.NoError(t, err)
	t.Run("MalformedURL", func(t *testing.T) {
		t.Parallel()
		d := NewDefaultDownloaderWithRetries(1)

		file := c.DownloadFile{
			URL:      "http://a b.com", // URL with space is invalid
			Folder:   downloadDir,
			Filename: "malformed_url.txt",
		}

		_, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.Error(t, err, "Download with malformed URL should fail")
	})

	t.Run("InvalidURL", func(t *testing.T) {
		t.Parallel()
		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      "://invalid-url",
			Folder:   downloadDir,
			Filename: "invalid_url.txt",
		}

		_, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.Error(t, err, "Download with invalid URL should fail")
	})

	t.Run("NonExistentLocalFile", func(t *testing.T) {
		t.Parallel()
		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      "file:///non/existent/file.txt",
			Folder:   downloadDir,
			Filename: "non_existent.txt",
		}

		_, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.Error(t, err, "Download with non-existent local file should fail")
	})

	t.Run("NonExistentDestinationFolder", func(t *testing.T) {
		// Don't run this in parallel as it has issues with shared testDir
		// Create a local test dir for this test
		localTestDir, err := os.MkdirTemp("", "nonexistent_folder_test")
		assert.NoError(t, err)
		defer os.RemoveAll(localTestDir)

		sourceContent := "test content"
		sourceFilename := "source.txt"
		sourcePath := createTestFile(t, localTestDir, sourceFilename, sourceContent)
		sourceURL := fmt.Sprintf("file://%s", sourcePath)

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      sourceURL,
			Folder:   filepath.Join(localTestDir, "nonexistent"),
			Filename: "destination.txt",
		}

		_, _, err = d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.Error(t, err, "Download to non-existent folder should fail")
	})
	t.Run("HeadRequestError", func(t *testing.T) {
		errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}))
		defer errorServer.Close()

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      errorServer.URL,
			Folder:   downloadDir,
			Filename: "error_response.txt",
		}

		_, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.Error(t, err, "Download should fail with HTTP error")

		var httpErr *HTTPStatusError
		if assert.ErrorAs(t, err, &httpErr) {
			assert.Equal(t, http.StatusInternalServerError, httpErr.StatusCode, "Should return correct status code")
		}
	})

	t.Run("Redirect", func(t *testing.T) {
		t.Parallel()
		finalServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, err := fmt.Fprintln(w, "final content"); err != nil {
				fmt.Println("Error writing response:", err)
			}
		}))
		defer finalServer.Close()

		redirectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, finalServer.URL, http.StatusFound)
		}))
		defer redirectServer.Close()

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      redirectServer.URL,
			Folder:   downloadDir,
			Filename: "redirect.txt",
		}

		filePath, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.NoError(t, err)

		content, err := os.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Contains(t, string(content), "final content")
	})

	t.Run("LargeFile", func(t *testing.T) {
		t.Parallel()
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "102400") // 100KB instead of 1MB
			w.Header().Set("Content-Type", "application/octet-stream")

			// Write 100KB of data in chunks
			chunk := make([]byte, 1024)
			for i := range chunk {
				chunk[i] = 'A'
			}

			for i := 0; i < 100; i++ { // 100 chunks instead of 1024
				_, _ = w.Write(chunk)
			}
		}))
		defer server.Close()

		d := NewDefaultDownloaderWithRetries(1)
		file := c.DownloadFile{
			URL:      server.URL,
			Folder:   downloadDir,
			Filename: "large_file.bin",
		}

		filePath, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
		assert.NoError(t, err)

		fileInfo, err := os.Stat(filePath)
		assert.NoError(t, err)
		assert.Equal(t, int64(102400), fileInfo.Size())
	})
}

func TestDownloaderConcurrency_Merged(t *testing.T) {
	t.Parallel()
	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer os.RemoveAll(testDir)

	downloadDir := filepath.Join(testDir, "download")
	err := os.MkdirAll(downloadDir, 0755)
	assert.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "content for request %s", r.URL.Path); err != nil {
			fmt.Println("Error writing response:", err)
		}
	}))
	defer server.Close()

	// Run multiple downloads concurrently
	numDownloads := 3
	results := make(chan error, numDownloads)

	for i := 0; i < numDownloads; i++ {
		go func(index int) {
			d := NewDefaultDownloaderWithRetries(1)
			file := c.DownloadFile{
				URL:      fmt.Sprintf("%s/file%d", server.URL, index),
				Folder:   downloadDir,
				Filename: fmt.Sprintf("concurrent_%d.txt", index),
			}

			_, _, err := d.Download(logger, file, false, nil, config.ApplicationConfig{})
			results <- err
		}(i)
	}

	for i := 0; i < numDownloads; i++ {
		err := <-results
		assert.NoError(t, err, "Concurrent download should not fail")
	}
}
