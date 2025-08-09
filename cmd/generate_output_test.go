package cmd

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"text/template"

	"github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateOutputCmdRun(t *testing.T) {
	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	require.NoError(t, err)
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

	generateOutputCmd.Run(generateOutputCmd, []string{})
}

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
	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Store original values
	origOutputDir := constants.OutputDir
	origOutputIgnoredDir := constants.OutputIgnoredDir
	origOutputSummariesDir := constants.OutputSummariesDir
	origArchiveDir := constants.ArchiveDir
	origIncludeIgnored := includeIgnored
	defer func() {
		constants.OutputDir = origOutputDir
		constants.OutputIgnoredDir = origOutputIgnoredDir
		constants.OutputSummariesDir = origOutputSummariesDir
		constants.ArchiveDir = origArchiveDir
		includeIgnored = origIncludeIgnored
	}()

	constants.OutputDir = filepath.Join(tempDir, "output")
	constants.OutputIgnoredDir = filepath.Join(tempDir, "ignored")
	constants.OutputSummariesDir = filepath.Join(tempDir, "summaries")
	constants.ArchiveDir = filepath.Join(tempDir, "archive")

	includeIgnored = false
	err = prepareDirectories()
	assert.NoError(t, err)

	assert.DirExists(t, constants.OutputDir)
	assert.DirExists(t, constants.OutputSummariesDir)
	assert.DirExists(t, constants.ArchiveDir)
	assert.NoDirExists(t, constants.OutputIgnoredDir) // Should not exist when includeIgnored is false

	includeIgnored = true
	err = prepareDirectories()
	assert.NoError(t, err)

	assert.DirExists(t, constants.OutputIgnoredDir)
}

func TestCopySummaryFile(t *testing.T) {
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "dns-toolkit-test")
	assert.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

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

	err := os.Setenv("DNS_TOOLKIT_TEST_MODE", "true")
	assert.NoError(t, err)
	configPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	if configPath == "" {
		t.Skip("DNS_TOOLKIT_TEST_CONFIG_PATH is not set, skipping test")
		return
	}
	defer func() {
		err := os.Unsetenv("DNS_TOOLKIT_TEST_MODE")
		if err != nil {
			t.Logf("Failed to unset DNS_TOOLKIT_TEST_MODE: %v", err)
		}
	}()

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
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

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
		"Files:\n  - test file 1\n  - test file 2",
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
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove temp directory: %v", err)
		}
	}(tempDir)

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
	filesInvolved := map[string][]common.FileInfo{
		inputFile: {},
	}

	processRegularFiles(dynTmpl, staticTemplate, "testtype", typeFiles, fileCount, filesInvolved)

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
	origIncludeIgnored := includeIgnored
	defer func() { includeIgnored = origIncludeIgnored }()

	tempDir, err := os.MkdirTemp("", "ignored-files-test")
	assert.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Logf("Failed to remove temp directory: %v", err)
		}
	}(tempDir)

	origIgnoredDir := constants.OutputIgnoredDir
	constants.OutputIgnoredDir = tempDir
	defer func() { constants.OutputIgnoredDir = origIgnoredDir }()

	dynTmpl, err := template.New("dynamic").Parse("Header: {{.FileName}} - {{.Description}} - {{.Count}}")
	assert.NoError(t, err)
	staticTemplate := []byte("STATIC HEADER")

	includeIgnored = false
	ignoredFilesCount := map[string]int{
		"some_file.txt": 3,
	}
	processIgnoredFiles(dynTmpl, staticTemplate, "testtype", ignoredFilesCount)

	includeIgnored = true
	emptyIgnored := map[string]int{}
	processIgnoredFiles(dynTmpl, staticTemplate, "testtype", emptyIgnored)

	includeIgnored = true
	ignoredFile := filepath.Join(tempDir, "ignored.txt")
	ignoredContent := "ignored data"
	err = os.WriteFile(ignoredFile, []byte(ignoredContent), 0644)
	assert.NoError(t, err)

	ignoredFilesCount = map[string]int{
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

	nonExistentIgnored := map[string]int{
		"/nonexistent/path/file.txt": 5,
	}
	processIgnoredFiles(dynTmpl, staticTemplate, "testtype", nonExistentIgnored)
}

func TestProcessFilesForSummaryType(t *testing.T) {
	tests := []struct {
		wantType    map[string]string
		wantCount   map[string]int
		wantIgnored map[string]int
		name        string
		summaryType string
		summaryData string
	}{
		{
			name:        "consolidated summary",
			summaryType: "consolidated",
			summaryData: `[{"filepath":"file1.txt","list_type":"blocklist","count":10,"ignored_entries_count":2,"ignored_filepath":"file1_ignored.txt"}]`,
			wantType:    map[string]string{"file1.txt": "blocklist"},
			wantCount:   map[string]int{"file1.txt": 10},
			wantIgnored: map[string]int{"file1_ignored.txt": 2},
		},
		{
			name:        "consolidated_groups summary",
			summaryType: "consolidated_groups",
			summaryData: `[{"filepath":"file2.txt","list_type":"allowlist","count":5,"ignored_entries_count":0,"ignored_filepath":""}]`,
			wantType:    map[string]string{"file2.txt": "allowlist"},
			wantCount:   map[string]int{"file2.txt": 5},
			wantIgnored: map[string]int{},
		},
		{
			name:        "consolidated_categories summary",
			summaryType: "consolidated_categories",
			summaryData: `[{"filepath":"file3.txt","list_type":"blocklist","count":7,"ignored_entries_count":1,"ignored_filepath":"file3_ignored.txt"}]`,
			wantType:    map[string]string{"file3.txt": "blocklist"},
			wantCount:   map[string]int{"file3.txt": 7},
			wantIgnored: map[string]int{"file3_ignored.txt": 1},
		},
		{
			name:        "top summary",
			summaryType: "top",
			summaryData: `[{"filepath":"file4.txt","list_type":"blocklist","count":3}]`,
			wantType:    map[string]string{"file4.txt": "blocklist"},
			wantCount:   map[string]int{"file4.txt": 3},
			wantIgnored: map[string]int{},
		},
		{
			name:        "unknown summary type",
			summaryType: "unknown",
			summaryData: `[]`,
			wantType:    map[string]string{},
			wantCount:   map[string]int{},
			wantIgnored: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotCount, gotFilesInvolved, gotIgnored := processFilesForSummaryType(
				tt.summaryType,
				[]byte(tt.summaryData),
			)
			assert.Equal(t, tt.wantType, gotType)
			assert.Equal(t, tt.wantCount, gotCount)
			assert.Equal(t, tt.wantIgnored, gotIgnored)
			assert.NotNil(t, gotFilesInvolved)
		})
	}
}

func TestCopySummaryFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "copy-summary-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	origOutputSummariesDir := constants.OutputSummariesDir
	origSummaryDir := constants.SummaryDir
	defer func() {
		constants.OutputSummariesDir = origOutputSummariesDir
		constants.SummaryDir = origSummaryDir
	}()

	summariesDir := filepath.Join(tempDir, "summaries")
	outputDir := filepath.Join(tempDir, "output")
	constants.OutputSummariesDir = outputDir
	constants.SummaryDir = summariesDir

	require.NoError(t, os.MkdirAll(summariesDir, 0755))
	require.NoError(t, os.MkdirAll(outputDir, 0755))

	testFiles := []string{
		constants.DefaultSummaryFiles[constants.SummaryTypeDownload],
		constants.DefaultSummaryFiles[constants.SummaryTypeProcessed],
	}
	testContent := `{"test": "data"}`
	for _, filename := range testFiles {
		filePath := filepath.Join(summariesDir, filename)
		require.NoError(t, os.WriteFile(filePath, []byte(testContent), 0644))
	}

	processedFiles := map[string]string{
		constants.SummaryTypeDownload: filepath.Join(
			summariesDir,
			constants.DefaultSummaryFiles[constants.SummaryTypeDownload],
		),
		constants.SummaryTypeProcessed: filepath.Join(
			summariesDir,
			constants.DefaultSummaryFiles[constants.SummaryTypeProcessed],
		),
	}
	copySummaryFiles(processedFiles, outputDir)

	for _, filename := range testFiles {
		destPath := filepath.Join(outputDir, filename)
		assert.FileExists(t, destPath)

		content, err := os.ReadFile(destPath)
		require.NoError(t, err)
		assert.Equal(t, `{"test": "data"}`, string(content))
	}

	nonExistentFiles := map[string]string{
		"missing.json": "/nonexistent/path/missing.json",
	}
	copySummaryFiles(nonExistentFiles, outputDir)
}

func TestArchiveSummaryFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "archive-summary-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Store original values
	origArchiveDir := constants.ArchiveDir
	origSummaryDir := constants.SummaryDir
	defer func() {
		constants.ArchiveDir = origArchiveDir
		constants.SummaryDir = origSummaryDir
	}()

	summariesDir := filepath.Join(tempDir, "summaries")
	archiveDir := filepath.Join(tempDir, "archive")
	constants.ArchiveDir = archiveDir
	constants.SummaryDir = summariesDir

	require.NoError(t, os.MkdirAll(summariesDir, 0755))
	require.NoError(t, os.MkdirAll(archiveDir, 0755))

	testSummaryTypes := []string{"download", "processed"}
	processedFiles := make(map[string]string)

	for _, summaryType := range testSummaryTypes {
		filename := constants.DefaultSummaryFiles[summaryType]
		filePath := filepath.Join(summariesDir, filename)
		require.NoError(t, os.WriteFile(filePath, []byte(`{"archive": "test"}`), 0644))
		processedFiles[summaryType] = filePath
	}

	archiveSummaryFiles(processedFiles)

	archiveFiles, err := os.ReadDir(archiveDir)
	require.NoError(t, err)

	foundArchiveFiles := 0
	for _, file := range archiveFiles {
		if strings.Contains(file.Name(), "download") || strings.Contains(file.Name(), "processed") {
			foundArchiveFiles++
		}
	}
	assert.Greater(t, foundArchiveFiles, 0)
}

func TestDeleteFilesAndFoldersAfterGeneration(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "delete-test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temporary directory: %v", err)
		}
	}()

	// Store original values
	origSummaryDir := constants.SummaryDir
	origProcessedDir := constants.ProcessedDir
	origDeleteFolders := deleteFolders
	defer func() {
		constants.SummaryDir = origSummaryDir
		constants.ProcessedDir = origProcessedDir
		constants.SummaryTypesDirMap[constants.SummaryTypeProcessed] = origProcessedDir
		deleteFolders = origDeleteFolders
	}()

	summaryDir := filepath.Join(tempDir, "summary")
	processedDir := filepath.Join(tempDir, "processed")
	constants.SummaryDir = summaryDir
	constants.ProcessedDir = processedDir
	constants.SummaryTypesDirMap[constants.SummaryTypeProcessed] = processedDir
	deleteFolders = true // Enable deletion

	require.NoError(t, os.MkdirAll(summaryDir, 0755))
	require.NoError(t, os.MkdirAll(processedDir, 0755))

	processedSummaryFile := filepath.Join(summaryDir, constants.DefaultSummaryFiles[constants.SummaryTypeProcessed])
	require.NoError(t, os.WriteFile(processedSummaryFile, []byte(`{"test": "data"}`), 0644))

	testFile := filepath.Join(processedDir, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("test content"), 0644))

	assert.FileExists(t, processedSummaryFile)
	assert.FileExists(t, testFile)

	deleteFilesAndFoldersAfterGeneration()

	assert.NoFileExists(t, processedSummaryFile)
	assert.NoDirExists(t, processedDir)

	deleteFolders = false

	require.NoError(t, os.MkdirAll(processedDir, 0755))
	require.NoError(t, os.WriteFile(processedSummaryFile, []byte(`{"test": "data"}`), 0644))
	require.NoError(t, os.WriteFile(testFile, []byte("test content"), 0644))

	deleteFilesAndFoldersAfterGeneration()

	assert.FileExists(t, processedSummaryFile)
	assert.FileExists(t, testFile)
}

func TestParseFileInfoFromString(t *testing.T) {
	cases := []struct {
		input   string
		want    common.FileInfo
		wantErr bool
	}{
		{
			input: "foo.txt [domain] [/path/foo.txt] [42]",
			want: common.FileInfo{
				Name:         "foo.txt",
				SourceType:   "domain",
				Filepath:     "/path/foo.txt",
				Count:        42,
				MustConsider: false,
			},
			wantErr: false,
		},
		{
			input: "bar.txt [ipv4] [/bar.txt] [7] [must consider]",
			want: common.FileInfo{
				Name:         "bar.txt",
				SourceType:   "ipv4",
				Filepath:     "/bar.txt",
				Count:        7,
				MustConsider: true,
			},
			wantErr: false,
		},
		{
			input:   "invalid format",
			want:    common.FileInfo{},
			wantErr: true,
		},
	}
	for _, tc := range cases {
		got, err := parseFileInfoFromString(tc.input)
		if tc.wantErr && err == nil {
			t.Errorf("expected error for input %q", tc.input)
		}
		if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
			t.Errorf("parseFileInfoFromString(%q) = %+v, want %+v", tc.input, got, tc.want)
		}
	}
}

func TestParseFilesFromConsolidatedSummary(t *testing.T) {
	summary := common.ConsolidatedSummary{
		Files: []string{
			"foo.txt [domain] [/foo.txt] [10]",
			"bar.txt [ipv4] [/bar.txt] [5] [must consider]",
		},
	}
	want := []common.FileInfo{
		{Name: "foo.txt", SourceType: "domain", Filepath: "/foo.txt", Count: 10, MustConsider: false},
		{Name: "bar.txt", SourceType: "ipv4", Filepath: "/bar.txt", Count: 5, MustConsider: true},
	}
	got := parseFilesFromConsolidatedSummary(summary)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseFilesFromConsolidatedSummary() = %+v, want %+v", got, want)
	}
}

func TestGenerateFilesList(t *testing.T) {
	files := []common.FileInfo{
		{Name: "foo", SourceType: "domain", Count: 10, MustConsider: false},
		{Name: "bar", SourceType: "ipv4_cidr_expand", Count: 5, MustConsider: true},
	}
	got := generateFilesList("", "domain", "", files)
	want := "# This domain list was consolidated from 2 source file(s):\n#   - foo domain          : 10\n#   - bar ipv4_cidr_expand: 5  [must consider]"
	if got != want {
		t.Errorf("generateFilesList() = %q, want %q", got, want)
	}
	if generateFilesList("", "domain", "", nil) != "" {
		t.Errorf("generateFilesList() with nil files should return empty string")
	}
}
