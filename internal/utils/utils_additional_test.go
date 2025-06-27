package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopyFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_copy_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	srcFile := filepath.Join(tempDir, "source.txt")
	dstFile := filepath.Join(tempDir, "destination.txt")

	testContent := "This is test content for file copying"
	err = os.WriteFile(srcFile, []byte(testContent), 0644)
	require.NoError(t, err)

	err = copyFile(logger, srcFile, dstFile)
	assert.NoError(t, err)

	copiedContent, err := os.ReadFile(dstFile)
	require.NoError(t, err)
	assert.Equal(t, testContent, string(copiedContent))

	err = copyFile(logger, "/non/existent/file.txt", dstFile)
	assert.Error(t, err)

	err = copyFile(logger, srcFile, "/non/existent/dir/file.txt")
	assert.Error(t, err)
}

func TestFindOverlap(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_overlap_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")

	content1 := `# Comment in file 1
domain1.com
domain2.com
domain3.com
unique1.com`

	content2 := `# Comment in file 2
domain2.com
domain3.com
domain4.com
unique2.com`

	err = os.WriteFile(file1, []byte(content1), 0644)
	require.NoError(t, err)

	err = os.WriteFile(file2, []byte(content2), 0644)
	require.NoError(t, err)

	overlap, count1, count2 := FindOverlap(logger, file1, file2)

	// domain2.com and domain3.com overlapping
	assert.Len(t, overlap, 2)
	assert.Contains(t, overlap, "domain2.com")
	assert.Contains(t, overlap, "domain3.com")

	assert.Equal(t, 4, count1) // 4 non-comment lines in file1
	assert.Equal(t, 4, count2) // 4 non-comment lines in file2

	overlap, count1, count2 = FindOverlap(logger, "/non/existent/file.txt", file2)
	assert.Nil(t, overlap)
	assert.Equal(t, 0, count1)
	assert.Equal(t, 0, count2)

	emptyFile1 := filepath.Join(tempDir, "empty1.txt")
	emptyFile2 := filepath.Join(tempDir, "empty2.txt")
	err = os.WriteFile(emptyFile1, []byte(""), 0644)
	require.NoError(t, err)
	err = os.WriteFile(emptyFile2, []byte(""), 0644)
	require.NoError(t, err)

	overlap, count1, count2 = FindOverlap(logger, emptyFile1, emptyFile2)
	assert.Len(t, overlap, 0)
	assert.Equal(t, 0, count1)
	assert.Equal(t, 0, count2)
}

func TestInitCommentPrefixTablesAndIsCommentFast(t *testing.T) {
	t.Parallel()

	testLines := []string{
		"# Hash comment",
		"// Slash comment",
		"; Semicolon comment",
		"! Exclamation comment",
		"regular line",
		"#comment",
	}

	expectedResults := []bool{
		true,  // # Hash comment
		true,  // // Slash comment
		true,  // ; Semicolon comment
		true,  // ! Exclamation comment
		false, // regular line
		true,  // #comment
	}

	for i, line := range testLines {
		result := isCommentFast(line)
		assert.Equal(t, expectedResults[i], result, "Line should be detected correctly: %s", line)
	}
}

func TestIsCommentFast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line     string
		expected bool
	}{
		{"# This is a comment", true},
		{"// This is a comment", true},
		{"; This is a comment", true},
		{"! This is a comment", true},
		{"domain.com", false},
		{"", true}, // Empty lines are treated as comments
		{"  # Indented comment", true},
		{"#comment", true},
		{"//comment", true},
		{";comment", true},
		{"!comment", true},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			result := isCommentFast(tt.line)
			assert.Equal(t, tt.expected, result, "Line: %q", tt.line)
		})
	}
}

func TestIsDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected bool
	}{
		{"example.com", true},
		{"sub.example.com", true},
		{"test-domain.org", true},
		{"domain123.net", true},
		{"xn--example.com", true}, // IDN domain
		{"192.168.1.1", false},    // IP address
		{"not_a_domain", false},
		{"", false},
		{".com", false},
		{"example.", false},
		{"example..com", false},
		{"-example.com", false},
		{"example-.com", false},
		{"exam ple.com", false}, // Space in domain
		{"very-long-subdomain-name-that-exceeds-normal-length.example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsDomain(tt.input)
			assert.Equal(t, tt.expected, result, "Input: %q", tt.input)
		})
	}
}

func TestWriteEntriesToFile(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_write_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	outputFile := filepath.Join(tempDir, "output.txt")

	entries := []string{
		"domain1.com",
		"domain2.com",
		"domain3.com",
	}

	err = WriteEntriesToFile(logger, outputFile, entries)
	assert.NoError(t, err)

	content, err := os.ReadFile(outputFile)
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	assert.Len(t, lines, 3)
	assert.Equal(t, "domain1.com", lines[0])
	assert.Equal(t, "domain2.com", lines[1])
	assert.Equal(t, "domain3.com", lines[2])

	emptyFile := filepath.Join(tempDir, "empty.txt")
	err = WriteEntriesToFile(logger, emptyFile, []string{})
	assert.NoError(t, err)

	_, err = os.ReadFile(emptyFile)
	assert.Error(t, err) // File should not exist

	err = WriteEntriesToFile(logger, "/non/existent/dir/file.txt", entries)
	assert.Error(t, err)
}

func TestIsSkipIP(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tests := []struct {
		ip       string
		expected bool
	}{
		{"127.0.0.1", false},
		{"127.1.2.3", false},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.1.1", true},
		{"169.254.1.1", false},
		{"224.0.0.1", false},
		{"255.255.255.255", true},
		{"0.0.0.0", false},
		{"8.8.8.8", true},
		{"1.1.1.1", true},
		{"invalid-ip", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			result := IsSkipIP(logger, tt.ip)
			assert.Equal(t, tt.expected, result, "IP: %q", tt.ip)
		})
	}
}

func TestParsePercent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected float64
	}{
		{"50", 50.0},
		{"100", 100.0},
		{"0", 0.0},
		{"75.5", 75.5},
		{"invalid", 0.0},
		{"-10", -10.0},
		{"150", 150.0},
		{"", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParsePercent(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractTarGzErrorCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_extract_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	err = extractTarGz(logger, "/non/existent/file.tar.gz", tempDir)
	assert.Error(t, err)

	invalidFile := filepath.Join(tempDir, "invalid.tar.gz")
	err = os.WriteFile(invalidFile, []byte("not a tar.gz file"), 0644)
	require.NoError(t, err)

	err = extractTarGz(logger, invalidFile, tempDir)
	assert.Error(t, err)

	validTarGz := filepath.Join(tempDir, "test.tar.gz")
	err = createMinimalTarGz(validTarGz)
	require.NoError(t, err)

	err = extractTarGz(logger, validTarGz, "/non/existent/output/dir")
	assert.Error(t, err)
}

func TestExtractZipErrorCases(t *testing.T) {
	t.Parallel()

	logger := createTestLogger()

	tempDir, err := os.MkdirTemp("", "test_extract_zip_*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	err = extractZip(logger, "/non/existent/file.zip", tempDir)
	assert.Error(t, err)

	invalidFile := filepath.Join(tempDir, "invalid.zip")
	err = os.WriteFile(invalidFile, []byte("not a zip file"), 0644)
	require.NoError(t, err)

	err = extractZip(logger, invalidFile, tempDir)
	assert.Error(t, err)
}

func createMinimalTarGz(filename string) error {
	return os.WriteFile(filename, []byte{}, 0644)
}
