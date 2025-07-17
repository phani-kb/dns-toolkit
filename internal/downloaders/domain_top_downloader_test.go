package downloaders

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDomainTopDownloaderWithRetries(t *testing.T) {
	t.Parallel()

	maxRetries := 3
	downloader := NewDomainTopDownloaderWithRetries(maxRetries)

	assert.NotNil(t, downloader, "Downloader should not be nil")
	assert.Equal(t, "tranco", downloader.Name(), "Downloader name should be 'tranco'")
}

func TestDomainTopDownloader_Name(t *testing.T) {
	t.Parallel()

	downloader := NewDomainTopDownloaderWithRetries(1)
	assert.Equal(t, "tranco", downloader.Name(), "Downloader name should be 'tranco'")
}

func TestDomainTopDownloader_Download(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	testContent := `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com
6,wikipedia.org
7,yahoo.com
8,whatsapp.com`

	// Setup test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(testContent))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	downloader := NewDomainTopDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "tranco_top.csv",
	}

	filePath, fetchSkipped, err := downloader.Download(logger, file, false, nil, config.ApplicationConfig{})

	assert.NoError(t, err, "Download should succeed")
	assert.False(t, fetchSkipped, "Fetch should not be skipped for new download")
	assert.Equal(t, filepath.Join(destDir, "tranco_top.csv"), filePath, "File path should match expected")

	// Verify file exists and has correct content
	content, err := os.ReadFile(filePath)
	require.NoError(t, err)
	assert.Equal(t, testContent, string(content), "Downloaded content should match server content")
}

func TestDomainTopDownloader_PostDownloadProcess(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	// Test cases for PostDownloadProcess
	tests := []struct {
		name            string
		originalContent string
		count           int
		expectedContent string
		expectError     bool
	}{
		{
			name: "process top 5 entries",
			originalContent: `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com
6,wikipedia.org
7,yahoo.com
8,whatsapp.com`,
			count: 5,
			expectedContent: `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com
`,
		},
		{
			name: "process top 3 entries",
			originalContent: `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com`,
			count: 3,
			expectedContent: `1,google.com
2,youtube.com
3,facebook.com
`,
		},
		{
			name: "count larger than available entries",
			originalContent: `1,google.com
2,youtube.com`,
			count: 5,
			expectedContent: `1,google.com
2,youtube.com
`,
		},
		{
			name:            "empty file",
			originalContent: "",
			count:           5,
			expectedContent: "",
		},
		{
			name: "count is zero",
			originalContent: `1,google.com
2,youtube.com
3,facebook.com`,
			count:           0,
			expectedContent: "",
		},
		{
			name:            "single entry",
			originalContent: `1,google.com`,
			count:           1,
			expectedContent: `1,google.com
`,
		},
		{
			name: "entries with extra whitespace",
			originalContent: `  1,google.com  
  2,youtube.com  
  3,facebook.com  `,
			count: 2,
			expectedContent: `  1,google.com  
  2,youtube.com  
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(testDir, fmt.Sprintf("test_%s.csv", strings.ReplaceAll(tt.name, " ", "_")))
			err := os.WriteFile(testFile, []byte(tt.originalContent), 0644)
			require.NoError(t, err)

			// Create downloader and process
			downloader := NewDomainTopDownloaderWithRetries(1)
			err = downloader.PostDownloadProcess(logger, testFile, tt.count)

			if tt.expectError {
				assert.Error(t, err, "PostDownloadProcess should return error for test: %s", tt.name)
				return
			}

			assert.NoError(t, err, "PostDownloadProcess should not return error for test: %s", tt.name)

			// Verify processed content
			processedContent, err := os.ReadFile(testFile)
			require.NoError(t, err)
			assert.Equal(
				t,
				tt.expectedContent,
				string(processedContent),
				"Processed content should match expected for test: %s",
				tt.name,
			)
		})
	}
}

func TestDomainTopDownloader_PostDownloadProcess_FileErrors(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	downloader := NewDomainTopDownloaderWithRetries(1)

	t.Run("non-existent file", func(t *testing.T) {
		err := downloader.PostDownloadProcess(logger, "/non/existent/file.csv", 5)
		assert.Error(t, err, "Should return error for non-existent file")
	})

	t.Run("directory instead of file", func(t *testing.T) {
		testDir := setupTestDir(t)
		defer func() {
			if err := os.RemoveAll(testDir); err != nil {
				t.Logf("Failed to remove test directory: %v", err)
			}
		}()

		err := downloader.PostDownloadProcess(logger, testDir, 5)
		assert.Error(t, err, "Should return error when trying to open directory as file")
	})

	t.Run("readonly file", func(t *testing.T) {
		testDir := setupTestDir(t)
		defer func() {
			if err := os.RemoveAll(testDir); err != nil {
				t.Logf("Failed to remove test directory: %v", err)
			}
		}()

		// Create readonly file
		testFile := filepath.Join(testDir, "readonly.csv")
		err := os.WriteFile(testFile, []byte("1,google.com\n2,youtube.com"), 0444)
		require.NoError(t, err)

		err = downloader.PostDownloadProcess(logger, testFile, 1)
		assert.Error(t, err, "Should return error when trying to write to readonly file")
	})
}

func TestDomainTopDownloader_PostDownloadProcess_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	tests := []struct {
		name            string
		originalContent string
		count           int
		expectedContent string
	}{
		{
			name: "file with only newlines",
			originalContent: `


`,
			count: 3,
			expectedContent: `


`,
		},
		{
			name: "file with mixed content and empty lines",
			originalContent: `1,google.com

2,youtube.com

3,facebook.com`,
			count: 2,
			expectedContent: `1,google.com

`,
		},
		{
			name: "very large count",
			originalContent: `1,google.com
2,youtube.com`,
			count: 1000000,
			expectedContent: `1,google.com
2,youtube.com
`,
		},
		{
			name: "zero count (should result in empty file)",
			originalContent: `1,google.com
2,youtube.com
3,facebook.com`,
			count:           0,
			expectedContent: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join(testDir, fmt.Sprintf("edge_%s.csv", strings.ReplaceAll(tt.name, " ", "_")))
			err := os.WriteFile(testFile, []byte(tt.originalContent), 0644)
			require.NoError(t, err)

			downloader := NewDomainTopDownloaderWithRetries(1)
			err = downloader.PostDownloadProcess(logger, testFile, tt.count)
			assert.NoError(t, err, "PostDownloadProcess should not return error for edge case: %s", tt.name)

			processedContent, err := os.ReadFile(testFile)
			require.NoError(t, err)
			assert.Equal(
				t,
				tt.expectedContent,
				string(processedContent),
				"Processed content should match expected for edge case: %s",
				tt.name,
			)
		})
	}
}

func TestDomainTopDownloader_Integration(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	testContent := `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com
6,wikipedia.org
7,yahoo.com
8,whatsapp.com
9,amazon.com
10,netflix.com
11,linkedin.com
12,reddit.com
13,ebay.com`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(testContent))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	downloader := NewDomainTopDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "tranco_integration.csv",
	}

	// Step 1: Download
	filePath, fetchSkipped, err := downloader.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.NoError(t, err, "Download should succeed")
	assert.False(t, fetchSkipped, "Fetch should not be skipped")

	// Step 2: Post-process to get top 10
	err = downloader.PostDownloadProcess(logger, filePath, 10)
	assert.NoError(t, err, "PostDownloadProcess should succeed")

	// Step 3: Verify final content
	finalContent, err := os.ReadFile(filePath)
	require.NoError(t, err)

	expectedFinalContent := `1,google.com
2,youtube.com
3,facebook.com
4,twitter.com
5,instagram.com
6,wikipedia.org
7,yahoo.com
8,whatsapp.com
9,amazon.com
10,netflix.com
`

	assert.Equal(
		t,
		expectedFinalContent,
		string(finalContent),
		"Final processed content should contain only top 10 entries",
	)
}

func TestDomainTopDownloader_Download_HTTPError(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	// Setup server that returns 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	destDir := filepath.Join(testDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	require.NoError(t, err)

	downloader := NewDomainTopDownloaderWithRetries(1)
	file := c.DownloadFile{
		URL:      server.URL,
		Folder:   destDir,
		Filename: "error_test.csv",
	}

	_, _, err = downloader.Download(logger, file, false, nil, config.ApplicationConfig{})
	assert.Error(t, err, "Download should fail with HTTP error")

	var httpErr *HTTPStatusError
	assert.ErrorAs(t, err, &httpErr, "Error should be HTTPStatusError")
	assert.Equal(t, http.StatusNotFound, httpErr.StatusCode, "Should return correct status code")
}

func TestDomainTopDownloader_PostDownloadProcess_LargeFile(t *testing.T) {
	t.Parallel()

	logger := setupTestLogger()
	testDir := setupTestDir(t)
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}()

	// Create a large file with many entries
	testFile := filepath.Join(testDir, "large_file.csv")
	file, err := os.Create(testFile)
	require.NoError(t, err)

	// Write 1000 entries
	for i := 1; i <= 1000; i++ {
		_, err := fmt.Fprintf(file, "%d,domain%d.com\n", i, i)
		require.NoError(t, err)
	}
	file.Close()

	downloader := NewDomainTopDownloaderWithRetries(1)
	err = downloader.PostDownloadProcess(logger, testFile, 100)
	assert.NoError(t, err, "PostDownloadProcess should succeed for large file")

	processedContent, err := os.ReadFile(testFile)
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(string(processedContent)), "\n")
	assert.Equal(t, 100, len(lines), "Should have exactly 100 lines after processing")

	assert.Equal(t, "1,domain1.com", lines[0], "First entry should be correct")
	assert.Equal(t, "100,domain100.com", lines[99], "Last entry should be correct")
}

func TestDomainTopDownloader_Registration(t *testing.T) {
	t.Parallel()

	downloader := NewDomainTopDownloaderWithRetries(1)
	err := RegisterDownloader(downloader)
	assert.NoError(t, err, "Registration should succeed")

	retrievedDownloader, exists := GetDownloader("tranco")
	assert.True(t, exists, "Downloader should exist after registration")
	assert.Equal(t, downloader.Name(), retrievedDownloader.Name(), "Retrieved downloader should have same name")

	// Test duplicate registration
	err = RegisterDownloader(downloader)
	assert.Error(t, err, "Duplicate registration should fail")
}
