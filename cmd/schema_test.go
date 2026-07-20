package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	cfg "github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/dns-toolkit/internal/constants"
)

func TestGetDBPath_UsesConfiguredPath(t *testing.T) {
	originalConfig := AppConfig
	originalTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE")
	originalTestConfigPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	t.Cleanup(func() {
		AppConfig = originalConfig
		_ = os.Setenv("DNS_TOOLKIT_TEST_MODE", originalTestMode)
		_ = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", originalTestConfigPath)
	})

	AppConfig = &cfg.AppConfig{}
	AppConfig.DNSToolkit.Database.Path = "testdata/custom.db"
	_ = os.Setenv("DNS_TOOLKIT_TEST_MODE", constants.BooleanTrue)
	_ = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", filepath.Join("..", "testdata", "config.yml"))

	got := getDBPath()
	if !filepath.IsAbs(got) {
		t.Fatalf("expected absolute configured db path in test mode, got %q", got)
	}
	if !strings.HasSuffix(filepath.ToSlash(got), "/testdata/custom.db") {
		t.Fatalf("expected configured db path to end with /testdata/custom.db, got %q", got)
	}
}

func TestGetDBPath_UsesDefaultTestPathInTestMode(t *testing.T) {
	originalConfig := AppConfig
	originalTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE")
	originalTestConfigPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	t.Cleanup(func() {
		AppConfig = originalConfig
		_ = os.Setenv("DNS_TOOLKIT_TEST_MODE", originalTestMode)
		_ = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", originalTestConfigPath)
	})

	AppConfig = nil
	_ = os.Setenv("DNS_TOOLKIT_TEST_MODE", constants.BooleanTrue)
	_ = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", filepath.Join("..", "testdata", "config.yml"))

	got := getDBPath()
	if !filepath.IsAbs(got) {
		t.Fatalf("expected absolute test db path, got %q", got)
	}
	if !strings.HasSuffix(filepath.ToSlash(got), "/"+constants.DefaultTestDBPath) {
		t.Fatalf("expected test db path to end with %q, got %q", constants.DefaultTestDBPath, got)
	}
}

func TestGetDBPath_UsesDefaultProdPathOutsideTestMode(t *testing.T) {
	originalConfig := AppConfig
	originalTestMode := os.Getenv("DNS_TOOLKIT_TEST_MODE")
	originalTestConfigPath := os.Getenv("DNS_TOOLKIT_TEST_CONFIG_PATH")
	t.Cleanup(func() {
		AppConfig = originalConfig
		_ = os.Setenv("DNS_TOOLKIT_TEST_MODE", originalTestMode)
		_ = os.Setenv("DNS_TOOLKIT_TEST_CONFIG_PATH", originalTestConfigPath)
	})

	AppConfig = nil
	_ = os.Unsetenv("DNS_TOOLKIT_TEST_MODE")

	got := getDBPath()
	if got != constants.DefaultDBPath {
		t.Fatalf("expected default db path %q, got %q", constants.DefaultDBPath, got)
	}
}
