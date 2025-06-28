package consolidators

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAdguardConsolidator(t *testing.T) {
	ac := NewAdguardConsolidator("adguard", "blocklist")
	assert.NotNil(t, ac)
	assert.Equal(t, "adguard", ac.GetSourceType())
	assert.Equal(t, "blocklist", ac.GetListType())
}

func TestAdguardConsolidator_FilterEntries(t *testing.T) {
	ac := NewAdguardConsolidator("adguard", "blocklist")
	logger := multilog.NewLogger()

	entrySet := u.NewStringSet([]string{"domain1.com", "domain2.com", "domain3.com"})
	filterSet := u.NewStringSet([]string{
		"@@domain2.com$important",
		"@@domain3.com",
		"domain4.com",
	})

	filtered, ignored := ac.FilterEntries(logger, entrySet, filterSet)

	assert.True(t, filtered.Contains("domain1.com"))
	assert.False(t, filtered.Contains("domain2.com"))
	assert.False(t, filtered.Contains("domain3.com"))
	assert.True(t, ignored.Contains("domain2.com"))
	assert.True(t, ignored.Contains("domain3.com"))
	assert.False(t, ignored.Contains("domain4.com"))
}

func TestAdguardConsolidator_Consolidate(t *testing.T) {
	ac := NewAdguardConsolidator("adguard", "blocklist")
	logger := multilog.NewLogger()

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("a.com\nb.com\n"), 0644))

	files := []c.ProcessedFile{
		{
			GenericSourceType: "adguard",
			ListType:          "blocklist",
			Filepath:          testFile,
			NumberOfEntries:   2,
			Name:              "test",
		},
	}
	set, infos := ac.Consolidate(logger, files)
	assert.Equal(t, 2, len(set))
	assert.Equal(t, 1, len(infos))
}

func TestAdguardConsolidator_SaveEntries(t *testing.T) {
	ac := NewAdguardConsolidator("adguard", "blocklist")
	logger := multilog.NewLogger()

	entrySet := u.NewStringSet([]string{"a.com", "b.com"})
	tempDir := t.TempDir()
	outFile := filepath.Join(tempDir, "out.txt")

	err := ac.SaveEntries(logger, entrySet, outFile)
	assert.NoError(t, err)

	content, err := os.ReadFile(outFile)
	assert.NoError(t, err)
	assert.Contains(t, string(content), "a.com")
	assert.Contains(t, string(content), "b.com")
}

func TestAdguardConsolidator_IsValid(t *testing.T) {
	ac := NewAdguardConsolidator("adguard", "blocklist")
	valid := c.ProcessedFile{
		GenericSourceType: "adguard",
		ListType:          "blocklist",
	}
	invalid := c.ProcessedFile{
		GenericSourceType: "other",
		ListType:          "blocklist",
	}
	assert.True(t, ac.IsValid(valid))
	assert.False(t, ac.IsValid(invalid))
}
