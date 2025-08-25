package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAllowlist(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	projectRoot, err := utils.FindProjectRoot("")
	if err != nil {
		t.Fatalf("failed to find project root: %v", err)
	}

	testDataDir := filepath.Join(projectRoot, "testdata")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	if err := os.Chdir(testDataDir); err != nil {
		t.Fatalf("failed to chdir to testdata: %v", err)
	}

	defer func() {
		_ = os.Chdir(cwd)
	}()

	generateAllowlist(logger)
}

func TestBackupExistingFiles(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	backupExistingFiles(logger, true)
}

func TestExtractSourceDomains(t *testing.T) {
	domains := extractSourceDomains()
	if domains == nil {
		t.Error("Expected non-nil slice")
	}
}

func TestLoadCustomElements(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	_ = loadCustomElements(logger, "test_custom.txt")
}

func TestWriteAllowlistWithStructure(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	file := "test_allowlist.txt"
	defer os.Remove(file)
	err := writeAllowlistWithStructure(
		logger,
		file,
		[]string{"a.com"},
		[]string{"b.com"},
		func(s string) string { return s },
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGenerateAdGuardRules(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	file := "test_adguard.txt"
	defer os.Remove(file)
	err := generateAdGuardRules(logger, file, []string{"a.com"}, []string{"b.com"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGenerateIPv4Addresses(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	file := "test_ipv4.txt"
	defer os.Remove(file)
	err := generateIPv4Addresses(logger, file, []string{"1.2.3.4"}, []string{"a.com"}, []string{"b.com"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetResolvedIPs(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ips := getResolvedIPs(logger, []string{"localhost"})
	if ips == nil {
		t.Error("Expected non-nil slice")
	}
}

func TestResolveDomainIPv4(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)
	ips, failedDomains := utils.ResolveDomainsToIPv4(logger, []string{"localhost"})
	if ips == nil {
		t.Error("Expected non-nil slice")
	}
	if len(failedDomains) > 0 {
		t.Error("Expected non-nil slice")
	}
}

func TestExtractDomainFromURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://example.com", "example.com"},
		{"http://sub.example.com", "sub.example.com"},
		{"https://example.com/path", "example.com"},
		{"https://example.com:8080", "example.com"},
		{"https://example.com/path?param=value", "example.com"},
		{"example.com", "example.com"},
		{"", ""},
	}

	for _, test := range tests {
		result := extractDomainFromURL(test.input)
		assert.Equal(t, test.expected, result, "Input: %s", test.input)
	}
}

func TestCombineDomains(t *testing.T) {
	custom := []string{"custom1.com", "custom2.com"}
	source := []string{"source1.com", "custom1.com", "source2.com"}

	result := combineDomains(custom, source)

	assert.Len(t, result, 4)
	assert.Equal(t, "custom1.com", result[0])
	assert.Equal(t, "custom2.com", result[1])
	assert.Contains(t, result, "source1.com")
	assert.Contains(t, result, "source2.com")
}

func TestStrFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"example.com", "example.com"},
		{"sub.example.com", "sub.example.com"},
		{"", ""},
	}

	for _, test := range tests {
		result := strFormat(test.input)
		assert.Equal(t, test.expected, result, "Input: %s", test.input)
	}
}

func TestAdgFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"example.com", "@@||example.com^"},
		{"sub.example.com", "@@||sub.example.com^"},
		{"", ""},
	}

	for _, test := range tests {
		result := adgFormat(test.input)
		assert.Equal(t, test.expected, result, "Input: %s", test.input)
	}
}
