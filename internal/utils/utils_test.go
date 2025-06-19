package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	cfg "github.com/phani-kb/dns-toolkit/internal/config"
)

func createTestLogger() *multilog.Logger {
	return multilog.NewLogger()
}

func TestNewStringSet(t *testing.T) {
	t.Parallel()

	entries := []string{"  apple  ", "banana", "cherry", "", "  ", "apple"}
	set := NewStringSet(entries)

	// Should have 3 unique trimmed entries
	assert.Len(t, set, 3)
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.True(t, set.Contains("cherry"))
	assert.False(t, set.Contains(""))
	assert.False(t, set.Contains("  "))
}

func TestNewStringSetWithCapacity(t *testing.T) {
	t.Parallel()

	set := NewStringSetWithCapacity(10)
	assert.NotNil(t, set)
	assert.Len(t, set, 0)
}

func TestStringSetOperations(t *testing.T) {
	t.Parallel()

	set := NewStringSet([]string{"apple", "banana"})

	assert.True(t, set.Contains("apple"))
	assert.False(t, set.Contains("cherry"))

	assert.False(t, set.MustConsider("apple"))
	assert.False(t, set.MustConsider("cherry"))

	consider, found := set.Get("apple")
	assert.True(t, found)
	assert.False(t, consider)

	consider, found = set.Get("cherry")
	assert.False(t, found)
	assert.False(t, consider)

	set.Add("cherry")
	assert.True(t, set.Contains("cherry"))

	added := set.AddWithConsider("date", true)
	assert.True(t, added)
	assert.True(t, set.MustConsider("date"))

	added = set.AddWithConsider("date", false)
	assert.False(t, added)
	assert.False(t, set.MustConsider("date")) // Consider flag updated

	set.Remove("banana")
	assert.False(t, set.Contains("banana"))

	newEntries := []string{"grape", "apple", "kiwi"}
	alreadyExist := set.AddAll(newEntries, true)
	assert.Equal(t, 1, alreadyExist) // "apple" already existed
	assert.True(t, set.MustConsider("grape"))
	assert.True(t, set.MustConsider("kiwi"))

	set.RemoveAll([]string{"grape", "nonexistent"})
	assert.False(t, set.Contains("grape"))

	slice := set.ToSlice()
	assert.Contains(t, slice, "apple")
	assert.Contains(t, slice, "cherry")

	sortedSlice := set.ToSliceSorted()
	assert.True(t, len(sortedSlice) >= 2)
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	// Test with duplicates
	entries := []string{"apple", "banana", "apple", "cherry", "banana"}
	result := RemoveDuplicates(entries)
	assert.Len(t, result, 3)

	// Test with empty slice
	empty := RemoveDuplicates([]string{})
	assert.Len(t, empty, 0)

	// Test with single item
	single := RemoveDuplicates([]string{"apple"})
	assert.Len(t, single, 1)
	assert.Equal(t, "apple", single[0])
}

func TestSaveFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()
	tempDir, err := os.MkdirTemp("", "test_save_file")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	content := "test content"
	reader := strings.NewReader(content)

	// Test successful save
	filePath, err := SaveFile(logger, tempDir, "test.txt", reader)
	assert.NoError(t, err)
	assert.Contains(t, filePath, "test.txt")

	savedContent, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(savedContent))

	// Test with non-existent directory (should create it)
	newDir := tempDir + "/newdir"
	reader2 := strings.NewReader("new content")
	filePath2, err := SaveFile(logger, newDir, "test2.txt", reader2)
	assert.NoError(t, err)
	assert.Contains(t, filePath2, "test2.txt")
}

func TestCloseFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()
	tempFile, err := os.CreateTemp("", "test_close")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	CloseFile(logger, tempFile)

	// Test double close (should log error but not panic)
	CloseFile(logger, tempFile)
}

func TestCalculateChecksum(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()
	tempFile, err := os.CreateTemp("", "test_checksum")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "test content for checksum"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	// MD5 (default)
	md5sum := CalculateChecksum(logger, tempFile.Name(), "")
	assert.NotEmpty(t, md5sum)
	assert.Len(t, md5sum, 32) // MD5 is 32 hex chars

	md5sum2 := CalculateChecksum(logger, tempFile.Name(), "md5")
	assert.Equal(t, md5sum, md5sum2)

	sha256sum := CalculateChecksum(logger, tempFile.Name(), "sha256")
	assert.NotEmpty(t, sha256sum)
	assert.Len(t, sha256sum, 64) // SHA256 is 64 hex chars

	// Unsupported algorithm
	unsupported := CalculateChecksum(logger, tempFile.Name(), "unsupported")
	assert.Empty(t, unsupported)

	nonexistent := CalculateChecksum(logger, "/nonexistent/file", "md5")
	assert.Empty(t, nonexistent)
}

func TestIsComment(t *testing.T) {
	t.Parallel()

	assert.True(t, IsComment(""))
	assert.True(t, IsComment("   "))

	assert.True(t, IsComment("# comment"))
	assert.True(t, IsComment("  # comment with spaces"))
	assert.True(t, IsComment("// comment"))
	assert.True(t, IsComment("! comment"))

	assert.False(t, IsComment("example.com"))
	assert.False(t, IsComment("192.168.1.1"))
}

func TestGetTimestamp(t *testing.T) {
	t.Parallel()

	timestamp := GetTimestamp()
	assert.NotEmpty(t, timestamp)
	assert.True(t, len(timestamp) > 10)
}

func TestStringInSlice(t *testing.T) {
	t.Parallel()

	slice := []string{"apple", "banana", "cherry"}

	assert.True(t, StringInSlice("apple", slice))
	assert.True(t, StringInSlice("banana", slice))
	assert.False(t, StringInSlice("grape", slice))
	assert.False(t, StringInSlice("", slice))
}

func TestPickRandomLines(t *testing.T) {
	t.Parallel()

	// Create a test file
	tempFile, err := os.CreateTemp("", "test_random_lines")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "line1\n# comment\nline2\n\nline3\n! another comment\nline4"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	lines, err := PickRandomLines(tempFile.Name(), 0)
	assert.NoError(t, err)
	assert.Contains(t, lines, "line1")
	assert.Contains(t, lines, "line2")
	assert.NotContains(t, lines, "# comment")

	lines, err = PickRandomLines(tempFile.Name(), 2)
	assert.NoError(t, err)
	assert.Len(t, lines, 2)

	_, err = PickRandomLines("/nonexistent/file", 5)
	assert.Error(t, err)
}

func TestReadEntriesFromFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	// Create a test file with duplicates and comments
	tempFile, err := os.CreateTemp("", "test_read_entries")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "example.com\n# comment\nexample.com\n192.168.1.1\n\nexample.org\nexample.com"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	entries, duplicates, err := ReadEntriesFromFile(logger, tempFile.Name())
	assert.NoError(t, err)
	assert.Contains(t, entries, "example.com")
	assert.Contains(t, entries, "192.168.1.1")
	assert.Contains(t, entries, "example.org")
	assert.Equal(t, 2, duplicates) // "example.com" appears 3 times, so 2 duplicates

	_, _, err = ReadEntriesFromFile(logger, "/nonexistent/file")
	assert.Error(t, err)
}

func TestReadEntriesFromFileWithPool(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	// Create a test file
	tempFile, err := os.CreateTemp("", "test_read_entries_pool")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	content := "example.com\nexample.org\nexample.com"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	tempFile.Close()

	entries, duplicates, err := ReadEntriesFromFileWithPool(logger, tempFile.Name(), nil)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	assert.Equal(t, 1, duplicates)

	pool := NewDTEntryPool()
	entries2, duplicates2, err := ReadEntriesFromFileWithPool(logger, tempFile.Name(), pool)
	assert.NoError(t, err)
	assert.Len(t, entries2, len(entries))
	assert.Contains(t, entries2, "example.com")
	assert.Contains(t, entries2, "example.org")
	assert.Equal(t, duplicates, duplicates2)
	assert.Greater(t, pool.Size(), 0)
}

func TestGetMapKeys(t *testing.T) {
	t.Parallel()

	stringMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	keys := GetMapKeys(stringMap)
	assert.Len(t, keys, 3)
	assert.Contains(t, keys, "apple")
	assert.Contains(t, keys, "banana")
	assert.Contains(t, keys, "cherry")

	// Test with empty map
	emptyMap := map[string]int{}
	emptyKeys := GetMapKeys(emptyMap)
	assert.Len(t, emptyKeys, 0)
}

func TestIsArchive(t *testing.T) {
	t.Parallel()

	assert.True(t, IsArchive("file.tar.gz"))
	assert.True(t, IsArchive("archive.zip"))
	assert.True(t, IsArchive("/path/to/file.tar.gz"))
	assert.False(t, IsArchive("file.txt"))
	assert.False(t, IsArchive("file.tar"))
	assert.False(t, IsArchive(""))
}

func TestExtractArchive(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	err := ExtractArchive(logger, "test.tar.gz", "/tmp")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported archive format")
}

func TestCopySourceToTarget(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	sourceDir, err := os.MkdirTemp("", "test_copy_source")
	require.NoError(t, err)
	defer os.RemoveAll(sourceDir)

	sourceFile := sourceDir + "/source.txt"
	err = os.WriteFile(sourceFile, []byte("test content"), 0644)
	require.NoError(t, err)

	targetDir, err := os.MkdirTemp("", "test_copy_target")
	require.NoError(t, err)
	defer os.RemoveAll(targetDir)

	target := c.DownloadTarget{
		SourceFolder: sourceDir,
		SourceFile:   "source.txt",
		TargetFolder: targetDir,
		TargetFile:   "target.txt",
	}

	err = CopySourceToTarget(logger, target)
	assert.NoError(t, err)

	targetContent, err := os.ReadFile(targetDir + "/target.txt")
	assert.NoError(t, err)
	assert.Equal(t, "test content", string(targetContent))

	err = CopySourceToTarget(logger, target)
	assert.NoError(t, err)

	target.SourceFile = "nonexistent.txt"
	err = CopySourceToTarget(logger, target)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "source file not found")
}

func TestIsAlphanumericWithUnderscoresAndDashes(t *testing.T) {
	t.Parallel()

	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("test123"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("test_123"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("test-123"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("ABC_def-123"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("test.123"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("test 123"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("test@123"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes(""))
}

func TestGetFileLastModifiedTime(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	// Create a test file
	tempFile, err := os.CreateTemp("", "test_mod_time")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Test getting modification time
	modTime, err := GetFileLastModifiedTime(logger, tempFile.Name())
	assert.NoError(t, err)
	assert.NotEmpty(t, modTime)

	_, err = GetFileLastModifiedTime(logger, "/nonexistent/file")
	assert.Error(t, err)
}

func TestCapPreallocEntries(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 50000, CapPreallocEntries(50000))
	assert.Equal(t, 10000, CapPreallocEntries(1))
	assert.Equal(t, 10000000, CapPreallocEntries(999999999))
}

func TestEnsureDirectoryExists(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_ensure_dir")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	newDir := tempDir + "/new/nested/dir"
	err = EnsureDirectoryExists(logger, newDir)
	assert.NoError(t, err)

	_, err = os.Stat(newDir)
	assert.NoError(t, err)

	err = EnsureDirectoryExists(logger, newDir)
	assert.NoError(t, err)
}

func TestGetUserAgent(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	// Complete config
	appConfig := cfg.ApplicationConfig{
		Name:        "TestApp",
		Version:     "1.0.0",
		Description: "Test Description",
	}
	userAgent := GetUserAgent(logger, appConfig)
	assert.Contains(t, userAgent, "TestApp/1.0.0")
	assert.Contains(t, userAgent, "Test Description")

	// Empty config (should use defaults)
	emptyConfig := cfg.ApplicationConfig{}
	userAgent2 := GetUserAgent(logger, emptyConfig)
	assert.NotEmpty(t, userAgent2)
}
