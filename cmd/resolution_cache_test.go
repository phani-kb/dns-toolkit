package cmd

import (
	"os"
	"path/filepath"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestGetCachedResolutionSets(t *testing.T) {
	logger, _ := multilog.NewTestLogger(t)

	cleanup, testDataDir := setupTestEnvironmentForCmdTests(t)
	defer cleanup()

	bl := filepath.Join(testDataDir, "cache_bl.txt")
	if err := os.WriteFile(bl, []byte("bad.example.org\n"), 0o644); err != nil {
		t.Fatalf("failed to write block list: %v", err)
	}

	al := filepath.Join(testDataDir, "cache_al.txt")
	if err := os.WriteFile(al, []byte("good.example.org\n"), 0o644); err != nil {
		t.Fatalf("failed to write allow list: %v", err)
	}

	blChecksum := u.CalculateChecksum(logger, bl, "")
	alChecksum := u.CalculateChecksum(logger, al, "")

	processed := []c.ProcessedFile{
		{
			Name:              "bl-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeBlocklist,
			Filepath:          bl,
			Checksum:          blChecksum,
			Valid:             true,
		},
		{
			Name:              "al-src",
			GenericSourceType: constants.SourceTypeDomain,
			ListType:          constants.ListTypeAllowlist,
			Filepath:          al,
			Checksum:          alChecksum,
			Valid:             true,
		},
	}

	// first call
	allow1, block1, conflicts1, _, _, _ := GetCachedResolutionSets(logger, processed)
	assert.NotNil(t, allow1)
	assert.NotNil(t, block1)
	assert.NotNil(t, conflicts1)

	newAllowEntry := "allow.com"
	if allow1[constants.SourceTypeDomain] == nil {
		allow1[constants.SourceTypeDomain] = u.NewStringSet([]string{})
	}
	allow1[constants.SourceTypeDomain].Add(newAllowEntry)

	// second call
	allow2, _, conflicts2, _, _, _ := GetCachedResolutionSets(logger, processed)
	assert.Contains(t, allow2[constants.SourceTypeDomain].ToSlice(), newAllowEntry)
	assert.Equal(t, conflicts1, conflicts2)

	processed[0].Checksum = "changed-checksum"
	allow3, _, conflicts3, _, _, _ := GetCachedResolutionSets(logger, processed)
	if allow3[constants.SourceTypeDomain] != nil {
		assert.NotContains(t, allow3[constants.SourceTypeDomain].ToSlice(), newAllowEntry)
	}
	assert.NotNil(t, conflicts3)
}
