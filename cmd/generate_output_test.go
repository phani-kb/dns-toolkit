package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/phani-kb/dns-toolkit/internal/constants"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		summaryType string
		fileName    string
		format      string
		listType    string
		expected    string
	}{
		{
			name:        "blocklist description",
			summaryType: "domain",
			fileName:    "domain_test_blocklist.txt",
			format:      "txt",
			listType:    "blocklist",
			expected:    "Txt Domain blocklist",
		},
		{
			name:        "allowlist description",
			summaryType: "ipv4",
			fileName:    "ipv4_test_allowlist.txt",
			format:      "txt",
			listType:    "allowlist",
			expected:    "Txt IPv4 allowlist",
		},
		{
			name:        "unknown list type",
			summaryType: "domain",
			fileName:    "general_test.txt",
			format:      "txt",
			listType:    "unknown",
			expected:    "Txt General unknown",
		},
		{
			name:        "empty parameters",
			summaryType: "",
			fileName:    "",
			format:      "",
			listType:    "",
			expected:    " General unknown",
		},
		{
			name:        "adguard format",
			summaryType: "adguard",
			fileName:    "adguard_test_blocklist.txt",
			format:      "adguard_rules",
			listType:    "blocklist",
			expected:    "Adguard Rules AdGuard blocklist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateDescription(tt.summaryType, tt.fileName, tt.format, tt.listType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPrepareDirectories(t *testing.T) {
	t.Parallel()

	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Change to the temp directory
	oldWd, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(oldWd)

	err = os.Chdir(tempDir)
	assert.NoError(t, err)

	includeIgnored = false
	_ = prepareDirectories()
}

func TestCopySummaryFile(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	srcPath := filepath.Join(tempDir, "source.txt")
	dstPath := filepath.Join(tempDir, "dest.txt")

	content := "test content"
	err = os.WriteFile(srcPath, []byte(content), 0644)
	assert.NoError(t, err)

	err = copySummaryFile(srcPath, dstPath)
	assert.NoError(t, err)

	assert.FileExists(t, dstPath)
	destContent, err := os.ReadFile(dstPath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(destContent))

	nonExistentSrc := filepath.Join(tempDir, "nonexistent.txt")
	err = copySummaryFile(nonExistentSrc, dstPath)
	assert.Error(t, err)
}

func TestLoadTemplates(t *testing.T) {
	t.Parallel()

	os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	defer os.Unsetenv("DNS_TOOLKIT_TEST_MODE")

	tmpl, staticTemplate, err := loadTemplates()

	if err != nil {
		assert.Contains(t, err.Error(), "template")
	} else {
		assert.NotNil(t, tmpl)
		assert.NotNil(t, staticTemplate)
		assert.Greater(t, len(staticTemplate), 0)
	}
}

func TestCreateOutputFromFile(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "output-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	inputFile := filepath.Join(tempDir, "input.txt")
	inputContent := "data line 1\ndata line 2"
	err = os.WriteFile(inputFile, []byte(inputContent), 0644)
	assert.NoError(t, err)

	dynTmpl, err := template.New("dynamic").
		Parse("Header: {{.FileName}} - {{.Description}} - {{.Count}} - {{.LastUpdated}}")
	assert.NoError(t, err)

	staticTemplate := []byte("STATIC HEADER")

	outputFile := filepath.Join(tempDir, "output.txt")

	err = createOutputFromFile(
		dynTmpl,
		staticTemplate,
		inputFile,
		"input.txt",
		"Test Description",
		42,
		outputFile,
	)
	assert.NoError(t, err)

	assert.FileExists(t, outputFile)

	outContent, err := os.ReadFile(outputFile)
	assert.NoError(t, err)
	contentStr := string(outContent)

	assert.Contains(t, contentStr, "Header: input.txt - Test Description - 42")
	assert.Contains(t, contentStr, "STATIC HEADER")
	assert.Contains(t, contentStr, inputContent)
	assert.True(t, strings.Contains(contentStr, "data line 1"))
}

func TestProcessRegularFiles(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "regular-files-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	origOutputDir := constants.SummaryTypesOutputDirMap["testtype"]
	constants.SummaryTypesOutputDirMap["testtype"] = tempDir
	defer func() { constants.SummaryTypesOutputDirMap["testtype"] = origOutputDir }()

	inputFile := filepath.Join(tempDir, "input.txt")
	inputContent := "sample data"
	err = os.WriteFile(inputFile, []byte(inputContent), 0644)
	assert.NoError(t, err)

	dynTmpl, err := template.New("dynamic").Parse("Header: {{.FileName}} - {{.Description}} - {{.Count}}")
	assert.NoError(t, err)

	staticTemplate := []byte("STATIC HEADER")

	typeFiles := map[string]string{
		inputFile: "blocklist",
	}
	fileCount := map[string]int{
		inputFile: 7,
	}

	processRegularFiles(dynTmpl, staticTemplate, "testtype", typeFiles, fileCount)

	outputFile := filepath.Join(tempDir, "input.txt")
	assert.FileExists(t, outputFile)

	outContent, err := os.ReadFile(outputFile)
	assert.NoError(t, err)
	contentStr := string(outContent)

	assert.Contains(t, contentStr, "Header: input.txt")
	assert.Contains(t, contentStr, "STATIC HEADER")
	assert.Contains(t, contentStr, inputContent)
}

func TestProcessIgnoredFiles(t *testing.T) {
	t.Parallel()

	includeIgnored = true
	defer func() { includeIgnored = false }()

	tempDir, err := os.MkdirTemp("", "ignored-files-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	origIgnoredDir := constants.OutputIgnoredDir
	constants.OutputIgnoredDir = tempDir
	defer func() { constants.OutputIgnoredDir = origIgnoredDir }()

	ignoredFile := filepath.Join(tempDir, "ignored.txt")
	ignoredContent := "ignored data"
	err = os.WriteFile(ignoredFile, []byte(ignoredContent), 0644)
	assert.NoError(t, err)

	dynTmpl, err := template.New("dynamic").Parse("Header: {{.FileName}} - {{.Description}} - {{.Count}}")
	assert.NoError(t, err)

	staticTemplate := []byte("STATIC HEADER")

	ignoredFilesCount := map[string]int{
		ignoredFile: 3,
	}

	processIgnoredFiles(dynTmpl, staticTemplate, "testtype", ignoredFilesCount)

	outputFile := filepath.Join(tempDir, "ignored.txt")
	assert.FileExists(t, outputFile)

	outContent, err := os.ReadFile(outputFile)
	assert.NoError(t, err)
	contentStr := string(outContent)

	assert.Contains(t, contentStr, "Header: ignored.txt")
	assert.Contains(t, contentStr, "STATIC HEADER")
	assert.Contains(t, contentStr, ignoredContent)
}
