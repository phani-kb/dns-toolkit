package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
)

func createTestLogger(t *testing.T) *multilog.Logger {
	logger, _ := multilog.NewTestLogger(t)
	return logger
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

	empty := RemoveDuplicates([]string{})
	assert.Len(t, empty, 0)

	single := RemoveDuplicates([]string{"apple"})
	assert.Len(t, single, 1)
	assert.Equal(t, "apple", single[0])
}

func TestSaveFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tempDir, err := os.MkdirTemp("", "test_save_file")
	require.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove temp directory: %v", err)
		}
	}(tempDir)

	content := "test content"
	reader := strings.NewReader(content)

	filePath, err := SaveFile(logger, tempDir, "test.txt", reader)
	assert.NoError(t, err)
	assert.Contains(t, filePath, "test.txt")

	savedContent, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(savedContent))

	newDir := tempDir + "/newdir"
	reader2 := strings.NewReader("new content")
	filePath2, err := SaveFile(logger, newDir, "test2.txt", reader2)
	assert.NoError(t, err)
	assert.Contains(t, filePath2, "test2.txt")
}

func TestCloseFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tempFile, err := os.CreateTemp("", "test_close")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	CloseFile(logger, tempFile)

	CloseFile(logger, tempFile)
}

func TestCalculateChecksum(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tempFile, err := os.CreateTemp("", "test_checksum")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	content := "test content for checksum"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)

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

func TestCalculateChecksumFromContent(t *testing.T) {
	t.Parallel()

	content := []byte("test content for checksum")

	md5sum := CalculateChecksumFromContent(content, "")
	assert.NotEmpty(t, md5sum)
	assert.Len(t, md5sum, 32) // MD5 is 32 hex chars

	md5sum2 := CalculateChecksumFromContent(content, "md5")
	assert.Equal(t, md5sum, md5sum2)

	sha256sum := CalculateChecksumFromContent(content, "sha256")
	assert.NotEmpty(t, sha256sum)
	assert.Len(t, sha256sum, 64) // SHA256 is 64 hex chars

	unsupported := CalculateChecksumFromContent(content, "unsupported")
	assert.Empty(t, unsupported)

	empty := CalculateChecksumFromContent([]byte{}, "md5")
	assert.NotEmpty(t, empty) // Even empty content produces a hash
}

func TestGetArchiveExtension(t *testing.T) {
	t.Parallel()

	assert.Equal(t, ".tar.gz", GetArchiveExtension("file.tar.gz"))
	assert.Equal(t, ".zip", GetArchiveExtension("archive.zip"))
	assert.Equal(t, ".tar.gz", GetArchiveExtension("/path/to/file.tar.gz"))
	assert.Equal(t, "", GetArchiveExtension("file.txt"))
	assert.Equal(t, "", GetArchiveExtension("file.tar"))
	assert.Equal(t, "", GetArchiveExtension(""))
}

func TestExtractArchive(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	err := ExtractArchive(logger, "file.rar", "/tmp")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported archive format")

	err = ExtractArchive(logger, "nonexistent.tar.gz", "/tmp")
	assert.Error(t, err)
}

func TestCloseBody(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_close_body")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	CloseBody(logger, tempFile)
	CloseBody(logger, tempFile)
}

func TestShouldDownloadSource(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	result := ShouldDownloadSource(logger, "/nonexistent/file", "test-source")
	assert.True(t, result) // Should return true when file doesn't exist (no previous download)

	tempFile, err := os.CreateTemp("", "test_summary")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	writeSummary := func(lastDownloadTimestamp, frequency string) {
		summaries := []c.DownloadSummary{
			{
				Name:                  "test-source",
				LastDownloadTimestamp: lastDownloadTimestamp,
				Frequency:             frequency,
			},
		}
		data, err := json.Marshal(summaries)
		require.NoError(t, err)
		err = os.WriteFile(tempFile.Name(), data, 0644)
		require.NoError(t, err)
	}

	writeSummary("", constants.FrequencyDaily)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download when timestamp is empty

	writeSummary("0001-01-01T00:00:00Z", constants.FrequencyDaily)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download when timestamp is zero

	writeSummary("invalid-timestamp", constants.FrequencyDaily)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.False(t, result) // Should return false when timestamp is invalid

	recentTime := time.Now().Add(-1 * time.Hour).Format(constants.TimestampFormat)
	writeSummary(recentTime, constants.FrequencyDaily)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.False(t, result) // Should not download when recent

	oldTime := time.Now().Add(-25 * time.Hour).Format(constants.TimestampFormat)
	writeSummary(oldTime, constants.FrequencyDaily)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download when old enough

	weekOldTime := time.Now().Add(-8 * 24 * time.Hour).Format(constants.TimestampFormat)
	writeSummary(weekOldTime, constants.FrequencyWeekly)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download when week has passed

	monthOldTime := time.Now().Add(-31 * 24 * time.Hour).Format(constants.TimestampFormat)
	writeSummary(monthOldTime, constants.FrequencyMonthly)
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download when month has passed

	writeSummary(oldTime, "unknown")
	result = ShouldDownloadSource(logger, tempFile.Name(), "test-source")
	assert.True(t, result) // Should download with unknown frequency defaulting to daily
}

func TestLogMemStats(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	assert.NotPanics(t, func() {
		LogMemStats(logger, "test-prefix")
	})
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

	tempFile, err := os.CreateTemp("", "test_random_lines")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	content := "line1\n# comment\nline2\n\nline3\n! another comment\nline4"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)

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

	logger := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_read_entries")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	content := "example.com\n# comment\nexample.com\n192.168.1.1\n\nexample.org\nexample.com"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)

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

	logger := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_read_entries_pool")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	content := "example.com\nexample.org\nexample.com"
	_, err = tempFile.WriteString(content)
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)

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

func TestCopySourceToTarget(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	sourceDir, err := os.MkdirTemp("", "test_copy_source")
	require.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove test directory: %v", err)
		}
	}(sourceDir)

	sourceFile := sourceDir + "/source.txt"
	err = os.WriteFile(sourceFile, []byte("test content"), 0644)
	require.NoError(t, err)

	targetDir, err := os.MkdirTemp("", "test_copy_target")
	require.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove target directory: %v", err)
		}
	}(targetDir)

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

func TestCopySourceToTargetComprehensive(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tmpDir := t.TempDir()

	sourceFolder := filepath.Join(tmpDir, "source")
	err := os.MkdirAll(sourceFolder, 0755)
	require.NoError(t, err)

	sourceFile := "source.txt"
	sourceContent := "test source content"
	sourceFilePath := filepath.Join(sourceFolder, sourceFile)
	err = os.WriteFile(sourceFilePath, []byte(sourceContent), 0644)
	require.NoError(t, err)

	target := c.DownloadTarget{
		SourceFolder: sourceFolder,
		SourceFile:   "nonexistent.txt",
		TargetFolder: filepath.Join(tmpDir, "target"),
		TargetFile:   "target.txt",
	}

	err = CopySourceToTarget(logger, target)
	assert.Error(t, err)

	targetFolder := filepath.Join(tmpDir, "target", "nested")
	target = c.DownloadTarget{
		SourceFolder: sourceFolder,
		SourceFile:   sourceFile,
		TargetFolder: targetFolder,
		TargetFile:   "target.txt",
	}

	err = CopySourceToTarget(logger, target)
	assert.NoError(t, err)

	targetContent, err := os.ReadFile(filepath.Join(targetFolder, "target.txt"))
	assert.NoError(t, err)
	assert.Equal(t, sourceContent, string(targetContent))
}

func TestSaveFileErrorCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tmpDir := t.TempDir()

	if runtime.GOOS != "windows" {
		reader := strings.NewReader("test content")
		_, err := SaveFile(logger, "/proc/invalid/path", "test.txt", reader)
		assert.Error(t, err)
	}

	reader := strings.NewReader("test content")
	_, err := SaveFile(logger, tmpDir, "", reader)
	assert.Error(t, err)

	dirPath := filepath.Join(tmpDir, "directory")
	err = os.Mkdir(dirPath, 0755)
	require.NoError(t, err)

	reader = strings.NewReader("test content")
	_, err = SaveFile(logger, tmpDir, "directory", reader)
	assert.Error(t, err)
}

func TestCalculateChecksumFromContentAllAlgorithms(t *testing.T) {
	t.Parallel()

	testContent := []byte("test content for all algorithms")

	md5Result := CalculateChecksumFromContent(testContent, "md5")
	sha256Result := CalculateChecksumFromContent(testContent, "sha256")
	defaultResult := CalculateChecksumFromContent(testContent, "")
	unsupportedResult := CalculateChecksumFromContent(testContent, "unsupported")

	assert.NotEmpty(t, md5Result)
	assert.NotEmpty(t, sha256Result)
	assert.NotEmpty(t, defaultResult)
	assert.Empty(t, unsupportedResult) // Should be empty for unsupported algorithm

	assert.Equal(t, md5Result, defaultResult)   // Default should be MD5
	assert.NotEqual(t, md5Result, sha256Result) // Different algorithms should give different results

	emptyMD5 := CalculateChecksumFromContent([]byte{}, "md5")
	assert.NotEmpty(t, emptyMD5) // Even empty content has a hash
	assert.Len(t, emptyMD5, 32)  // MD5 is always 32 hex characters
}

func TestAllRemainingEdgeCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	assert.NotPanics(t, func() {
		LogMemStats(logger, "")
		LogMemStats(logger, "test")
		LogMemStats(logger, "very-long-prefix-string-that-might-cause-issues")
	})

	assert.True(t, len(constants.ArchiveExtensions) > 0)
	for _, ext := range constants.ArchiveExtensions {
		assert.True(t, strings.HasPrefix(ext, "."))
	}

	commentTestCases := []string{
		"\r\n",     // Windows line ending
		"\r",       // Mac classic line ending
		"   \t   ", // Mixed whitespace
		"#",        // Just hash
		"//",       // Just double slash
		"!",        // Just exclamation
		"# \t \n",  // Comment with trailing whitespace
	}

	for _, testCase := range commentTestCases {
		result := IsComment(testCase)
		assert.True(t, result, "Should be considered comment: %q", testCase)
	}

	nonCommentCases := []string{
		"a",                    // Single character
		"example.com#comment",  // Domain with hash
		"192.168.1.1//comment", // IP with double slash
	}

	for _, testCase := range nonCommentCases {
		result := IsComment(testCase)
		assert.False(t, result, "Should not be considered comment: %q", testCase)
	}
}

func TestEntryPoolInternBoundaryConditions(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	minLengthString := "123456" // Exactly 6 characters
	interned := pool.Intern(minLengthString)
	assert.Equal(t, minLengthString, interned)
	assert.Equal(t, 1, pool.Size()) // Should be added to pool

	maxLengthString := strings.Repeat("a", 255) // Exactly 255 characters
	interned = pool.Intern(maxLengthString)
	assert.Equal(t, maxLengthString, interned)
	assert.Equal(t, 2, pool.Size()) // Should be added to pool

	belowMinString := "12345" // 5 characters
	interned = pool.Intern(belowMinString)
	assert.Equal(t, belowMinString, interned)
	assert.Equal(t, 2, pool.Size()) // Should not be added to pool

	aboveMaxString := strings.Repeat("a", 256) // 256 characters
	interned = pool.Intern(aboveMaxString)
	assert.Equal(t, aboveMaxString, interned)
	assert.Equal(t, 2, pool.Size()) // Should not be added to pool
}

func TestSummaryUtilsErrorHandling(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_summary_permission")
	require.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Failed to remove temp file: %v", err)
		}
	}()

	validJSON := `[{"name": "test", "lastDownloadTimestamp": "20230101_120000"}]`
	err = os.WriteFile(tempFile.Name(), []byte(validJSON), 0644)
	require.NoError(t, err)

	summary, err := GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "test", summary.Name)

	summary, err = GetLastSummary[c.DownloadSummary](logger, tempFile.Name(), "nonexistent")
	assert.NoError(t, err)
	assert.Equal(t, "", summary.Name) // Should be zero value
}

func TestCapPreallocEntriesExactBoundaries(t *testing.T) {
	t.Parallel()

	small1 := CapPreallocEntries(1)
	small2 := CapPreallocEntries(10)
	small3 := CapPreallocEntries(100)

	assert.Greater(t, small1, 0)
	assert.Greater(t, small2, 0)
	assert.Greater(t, small3, 0)

	// Very large values should be capped to maximum
	large1 := CapPreallocEntries(1000000000)
	large2 := CapPreallocEntries(2000000000)

	assert.Greater(t, large1, 0)
	assert.Greater(t, large2, 0)
	assert.Equal(t, large1, large2) // Should both be capped to same maximum
}

func TestExtractArchiveErrorCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	err := ExtractArchive(logger, "/non/existent/file.zip", t.TempDir())
	assert.Error(t, err)

	tmpDir := t.TempDir()
	invalidPath := filepath.Join(tmpDir, "invalid.xyz")
	err = os.WriteFile(invalidPath, []byte("not an archive"), 0644)
	require.NoError(t, err)

	err = ExtractArchive(logger, invalidPath, t.TempDir())
	assert.Error(t, err)

	invalidZip := filepath.Join(tmpDir, "invalid.zip")
	err = os.WriteFile(invalidZip, []byte("not a zip file"), 0644)
	require.NoError(t, err)

	err = ExtractArchive(logger, invalidZip, t.TempDir())
	assert.Error(t, err)

	invalidTarGz := filepath.Join(tmpDir, "invalid.tar.gz")
	err = os.WriteFile(invalidTarGz, []byte("not a tar.gz file"), 0644)
	require.NoError(t, err)

	err = ExtractArchive(logger, invalidTarGz, t.TempDir())
	assert.Error(t, err)
}

func TestIsAlphanumericWithUnderscoresAndDashes(t *testing.T) {
	t.Parallel()

	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("valid"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("valid-name"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("valid_name"))
	assert.True(t, IsAlphanumericWithUnderscoresAndDashes("valid-name_123"))

	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("invalid!"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("invalid space"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes("invalid.dot"))
	assert.False(t, IsAlphanumericWithUnderscoresAndDashes(""))
}

func TestGetFileLastModifiedTime(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test-modified-time.txt")
	err := os.WriteFile(filePath, []byte("test content"), 0644)
	require.NoError(t, err)

	modTime, err := GetFileLastModifiedTime(logger, filePath)
	assert.NoError(t, err)
	assert.NotEmpty(t, modTime)

	nonExistentTime, err := GetFileLastModifiedTime(logger, filepath.Join(tmpDir, "non-existent.txt"))
	assert.Error(t, err)
	assert.Empty(t, nonExistentTime)
}

func TestEnsureDirectoryExists(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)
	tmpDir := t.TempDir()

	newDirPath := filepath.Join(tmpDir, "new-dir")
	err := EnsureDirectoryExists(logger, newDirPath)
	assert.NoError(t, err)

	dirInfo, err := os.Stat(newDirPath)
	require.NoError(t, err)
	assert.True(t, dirInfo.IsDir())

	err = EnsureDirectoryExists(logger, newDirPath) // Should be a no-op
	assert.NoError(t, err)

	if runtime.GOOS != "windows" {
		err = EnsureDirectoryExists(logger, "/proc/invalid/path")
		assert.Error(t, err)
	}
}

func TestGetUserAgent(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	userAgent := GetUserAgent(logger, "dns-toolkit-test", "v1.0.0", "Test Application")
	assert.Contains(t, userAgent, "dns-toolkit-test")
	assert.Contains(t, userAgent, "v1.0.0")
	assert.Contains(t, userAgent, "Test Application")

	defaultUserAgent := GetUserAgent(logger, "", "", "")
	assert.Contains(t, defaultUserAgent, constants.AppName)
	assert.Contains(t, defaultUserAgent, constants.AppVersion)
}

func TestCapPreallocEntriesEdgeCases(t *testing.T) {
	t.Parallel()

	result := CapPreallocEntries(0)
	assert.Equal(t, constants.MinPreallocEntries, result)

	result = CapPreallocEntries(-1)
	assert.Equal(t, constants.MinPreallocEntries, result)

	result = CapPreallocEntries(1000000000)
	assert.Equal(t, constants.MaxPreallocEntries, result)

	normalValue := constants.MinPreallocEntries + 5000
	result = CapPreallocEntries(normalValue)
	assert.Equal(t, normalValue, result)
}

func TestExtractArchiveTarGz(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tmpDir := t.TempDir()

	tarGzPath := filepath.Join(tmpDir, "test.tar.gz")

	testContent := "test content for extraction"
	testDir := filepath.Join(tmpDir, "source")
	err := os.MkdirAll(testDir, 0755)
	require.NoError(t, err)

	testFilePath := filepath.Join(testDir, "testfile.txt")
	err = os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	nestedDir := filepath.Join(testDir, "nested")
	err = os.MkdirAll(nestedDir, 0755)
	require.NoError(t, err)

	nestedFilePath := filepath.Join(nestedDir, "nestedfile.txt")
	err = os.WriteFile(nestedFilePath, []byte("nested content"), 0644)
	require.NoError(t, err)

	cmd := fmt.Sprintf("cd %s && tar -czf %s source/", tmpDir, tarGzPath)
	_, err = exec.Command("bash", "-c", cmd).Output()
	require.NoError(t, err)

	outputDir := filepath.Join(tmpDir, "output")
	err = ExtractArchive(logger, tarGzPath, outputDir)
	assert.NoError(t, err)

	extractedContent, err := os.ReadFile(filepath.Join(outputDir, "source", "testfile.txt"))
	assert.NoError(t, err)
	assert.Equal(t, testContent, string(extractedContent))

	nestedExtractedContent, err := os.ReadFile(filepath.Join(outputDir, "source", "nested", "nestedfile.txt"))
	assert.NoError(t, err)
	assert.Equal(t, "nested content", string(nestedExtractedContent))
}

func TestExtractArchiveZip(t *testing.T) {
	t.Parallel()

	logger := createTestLogger(t)

	tmpDir := t.TempDir()

	zipPath := filepath.Join(tmpDir, "test.zip")

	testContent := "test content for extraction"
	testDir := filepath.Join(tmpDir, "source")
	err := os.MkdirAll(testDir, 0755)
	require.NoError(t, err)

	testFilePath := filepath.Join(testDir, "testfile.txt")
	err = os.WriteFile(testFilePath, []byte(testContent), 0644)
	require.NoError(t, err)

	nestedDir := filepath.Join(testDir, "nested")
	err = os.MkdirAll(nestedDir, 0755)
	require.NoError(t, err)

	nestedFilePath := filepath.Join(nestedDir, "nestedfile.txt")
	err = os.WriteFile(nestedFilePath, []byte("nested content"), 0644)
	require.NoError(t, err)

	cmd := fmt.Sprintf("cd %s && zip -r %s source/", tmpDir, zipPath)
	_, err = exec.Command("bash", "-c", cmd).Output()
	require.NoError(t, err)

	outputDir := filepath.Join(tmpDir, "output")
	err = ExtractArchive(logger, zipPath, outputDir)
	assert.NoError(t, err)

	extractedContent, err := os.ReadFile(filepath.Join(outputDir, "source", "testfile.txt"))
	assert.NoError(t, err)
	assert.Equal(t, testContent, string(extractedContent))

	nestedExtractedContent, err := os.ReadFile(filepath.Join(outputDir, "source", "nested", "nestedfile.txt"))
	assert.NoError(t, err)
	assert.Equal(t, "nested content", string(nestedExtractedContent))
}

func TestFindProjectRoot(t *testing.T) {
	projectRoot, err := FindProjectRoot("")
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(projectRoot, "dns-toolkit"),
		"Project root should end with 'dns-toolkit', got: %s", projectRoot)

	goModPath := filepath.Join(projectRoot, "go.mod")
	_, err = os.Stat(goModPath)
	assert.NoError(t, err, "go.mod should exist in project root")

	wd, err := os.Getwd()
	require.NoError(t, err)

	projectRoot2, err := FindProjectRoot(wd)
	assert.NoError(t, err)
	assert.Equal(t, projectRoot, projectRoot2, "Should find same project root")

	_, err = FindProjectRoot("/nonexistent/path/that/does/not/exist")
	assert.Error(t, err)
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid IPv4 - basic", "192.168.1.1", true},
		{"Valid IPv4 - localhost", "127.0.0.1", true},
		{"Valid IPv4 - all zeros", "0.0.0.0", true},
		{"Valid IPv4 - all 255s", "255.255.255.255", true},
		{"Valid IPv4 - edge case", "10.0.0.1", true},

		{"Valid IPv6 - localhost", "::1", true},
		{"Valid IPv6 - full notation", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"Valid IPv6 - compressed", "2001:db8:85a3::8a2e:370:7334", true},
		{"Valid IPv6 - all zeros", "::", true},
		{"Valid IPv6 - mixed case", "2001:DB8:85A3::8A2E:370:7334", true},

		{"Invalid - empty string", "", false},
		{"Invalid - domain name", "example.com", false},
		{"Invalid - IPv4 with port", "192.168.1.1:8080", false},
		{"Invalid - IPv4 out of range", "256.1.1.1", false},
		{"Invalid - IPv4 negative", "-1.1.1.1", false},
		{"Invalid - IPv4 incomplete", "192.168.1", false},
		{"Invalid - IPv4 too many octets", "192.168.1.1.1", false},
		{"Invalid - IPv6 with port", "[::1]:8080", false},
		{"Invalid - IPv6 invalid format", "2001:0db8:85a3::8a2e::7334", false},
		{"Invalid - text", "not-an-ip", false},
		{"Invalid - CIDR notation", "192.168.1.0/24", false},
		{"Invalid - IPv6 CIDR", "2001:db8::/32", false},
		{"Invalid - spaces", " 192.168.1.1 ", false},
		{"Invalid - alphanumeric", "192.168.abc.1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIP(tt.input)
			if result != tt.expected {
				t.Errorf("IsIP(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsCIDR(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid IPv4 CIDR - /24", "192.168.1.0/24", true},
		{"Valid IPv4 CIDR - /32", "192.168.1.1/32", true},
		{"Valid IPv4 CIDR - /0", "0.0.0.0/0", true},
		{"Valid IPv4 CIDR - /16", "10.0.0.0/16", true},
		{"Valid IPv4 CIDR - /8", "172.16.0.0/8", true},
		{"Valid IPv4 CIDR - localhost", "127.0.0.1/32", true},

		{"Valid IPv6 CIDR - /64", "2001:db8::/64", true},
		{"Valid IPv6 CIDR - /128", "::1/128", true},
		{"Valid IPv6 CIDR - /0", "::/0", true},
		{"Valid IPv6 CIDR - /32", "2001:db8:85a3::/32", true},
		{"Valid IPv6 CIDR - full notation", "2001:0db8:85a3:0000:0000:8a2e:0370:7334/128", true},

		{"Invalid - empty string", "", false},
		{"Invalid - just IP without prefix", "192.168.1.1", false},
		{"Invalid - IPv6 without prefix", "::1", false},
		{"Invalid - invalid IP", "256.1.1.1/24", false},
		{"Invalid - negative prefix", "192.168.1.0/-1", false},
		{"Invalid - IPv4 prefix too large", "192.168.1.0/33", false},
		{"Invalid - IPv6 prefix too large", "::1/129", false},
		{"Invalid - no IP before slash", "/24", false},
		{"Invalid - no prefix after slash", "192.168.1.0/", false},
		{"Invalid - multiple slashes", "192.168.1.0/24/8", false},
		{"Invalid - text prefix", "192.168.1.0/abc", false},
		{"Invalid - domain with CIDR", "example.com/24", false},
		{"Invalid - IPv4 with port and CIDR", "192.168.1.1:8080/24", false},
		{"Invalid - IPv6 with brackets", "[::1]/128", false},
		{"Invalid - spaces", " 192.168.1.0/24 ", false},
		{"Invalid - decimal prefix", "192.168.1.0/24.5", false},
		{"Invalid - invalid IPv6", "2001:0db8:85a3::8a2e::7334/64", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCIDR(tt.input)
			if result != tt.expected {
				t.Errorf("IsCIDR(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExpandIpv4Range(t *testing.T) {
	logger := multilog.NewLogger()

	tests := []struct {
		name    string
		ipRange string
		want    []string
	}{
		{
			name:    "Valid range",
			ipRange: "192.168.1.1-192.168.1.3",
			want:    []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
		},
		{
			name:    "Single IP",
			ipRange: "10.0.0.1-10.0.0.1",
			want:    []string{"10.0.0.1"},
		},
		{
			name:    "Invalid format",
			ipRange: "192.168.1.1/24",
			want:    []string{},
		},
		{
			name:    "Start > End",
			ipRange: "192.168.1.10-192.168.1.1",
			want:    []string{},
		},
		{
			name:    "Non-IPv4",
			ipRange: "abcd-efgh",
			want:    []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandIpv4Range(logger, tt.ipRange)
			if len(got) != len(tt.want) {
				t.Errorf("ExpandIpv4Range(%q) got %d IPs, want %d", tt.ipRange, len(got), len(tt.want))
			}
			for i := range got {
				if i >= len(tt.want) || got[i] != tt.want[i] {
					t.Errorf("ExpandIpv4Range(%q)[%d] = %q, want %q", tt.ipRange, i, got[i], tt.want[i])
				}
			}
		})
	}
}
