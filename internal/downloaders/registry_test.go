package downloaders_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/downloaders"
	"github.com/phani-kb/dns-toolkit/internal/downloaders/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRegisterDownloader tests successful registration of a downloader
func TestRegisterDownloader(t *testing.T) {
	downloaderName := "test-downloader-" + t.Name()

	mockDownloader := mocks.NewDownloader(t)
	mockDownloader.On("Name").Return(downloaderName)

	err := downloaders.RegisterDownloader(mockDownloader)
	require.NoError(t, err)

	// Verify the downloader can be retrieved
	retrieved, exists := downloaders.GetDownloader(downloaderName)
	assert.True(t, exists)
	assert.Equal(t, mockDownloader, retrieved)
}

// TestRegisterDuplicateDownloader tests that registering a downloader with the same name fails
func TestRegisterDuplicateDownloader(t *testing.T) {
	downloaderName := "duplicate-name-" + t.Name()

	mockDownloader1 := mocks.NewDownloader(t)
	mockDownloader1.On("Name").Return(downloaderName)

	mockDownloader2 := mocks.NewDownloader(t)
	mockDownloader2.On("Name").Return(downloaderName)

	err := downloaders.RegisterDownloader(mockDownloader1)
	require.NoError(t, err)

	err = downloaders.RegisterDownloader(mockDownloader2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

// TestGetDownloader tests retrieving an existing downloader
func TestGetDownloader(t *testing.T) {
	downloaderName := "get-test-downloader-" + t.Name()

	mockDownloader := mocks.NewDownloader(t)
	mockDownloader.On("Name").Return(downloaderName)

	err := downloaders.RegisterDownloader(mockDownloader)
	require.NoError(t, err)

	retrieved, exists := downloaders.GetDownloader(downloaderName)
	assert.True(t, exists)
	assert.Equal(t, mockDownloader, retrieved)
}

// TestGetNonExistentDownloader tests retrieving a non-existent downloader
func TestGetNonExistentDownloader(t *testing.T) {
	_, exists := downloaders.GetDownloader("non-existent-downloader")
	assert.False(t, exists)
}
