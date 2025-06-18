package downloaders

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test the error types
func TestHTTPStatusError(t *testing.T) {
	t.Parallel()

	err := &HTTPStatusError{
		StatusCode: http.StatusNotFound,
		Status:     "Not Found",
		URL:        "http://example.com/test",
	}

	expectedMsg := "HTTP request returned non-success status: 404 Not Found for http://example.com/test"
	assert.Equal(t, expectedMsg, err.Error())
}

func TestCertVerificationError(t *testing.T) {
	t.Parallel()

	originalErr := assert.AnError
	err := &CertVerificationError{
		Host: "example.com",
		Err:  originalErr,
	}

	expectedMsg := "Certificate verification failed for host example.com: assert.AnError general error for testing"
	assert.Equal(t, expectedMsg, err.Error())
}

func TestDefaultDownloaderNameFunction(t *testing.T) {
	t.Parallel()

	name := DefaultDownloaderName()
	assert.Equal(t, "default", name)
}

func TestContainsHost(t *testing.T) {
	t.Parallel()

	hosts := []string{"example.com", "test.org", "localhost"}

	assert.True(t, containsHost(hosts, "example.com"))
	assert.True(t, containsHost(hosts, "test.org"))
	assert.False(t, containsHost(hosts, "notfound.com"))
	assert.False(t, containsHost(nil, "example.com"))
	assert.False(t, containsHost([]string{}, "example.com"))
}

func TestNewDefaultDownloaderForTesting(t *testing.T) {
	t.Parallel()

	d := NewDefaultDownloaderForTesting(3, 50*time.Millisecond)
	assert.NotNil(t, d)
	assert.Equal(t, "default", d.Name())

	assert.Equal(t, 50*time.Millisecond, d.retryDelay)
	assert.Equal(t, 3, d.maxRetries)
}

func TestDownloaderInterfaces(t *testing.T) {
	t.Parallel()

	// Test if downloader implements the Downloader interface
	var _ Downloader = (*DefaultDownloader)(nil)

	d := NewDefaultDownloaderWithRetries(1)
	err := d.PostDownloadProcess(nil, "test_path", 100)
	assert.NoError(t, err)
}
