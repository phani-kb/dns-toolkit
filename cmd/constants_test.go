package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
)

// TestConstantsSummaryDirSetProperly tests that the constants.SummaryDir is set correctly
func TestConstantsSummaryDirSetProperly(t *testing.T) {
	InitForTesting()

	assert.True(t, strings.HasSuffix(constants.SummaryDir, "testdata/summary"),
		"constants.SummaryDir should end with testdata/summary from test config, got: %s", constants.SummaryDir)
	assert.NotEqual(t, "data", constants.SummaryDir,
		"constants.SummaryDir should not be the default 'data'")
	assert.True(t, filepath.IsAbs(constants.SummaryDir),
		"constants.SummaryDir should be an absolute path in test mode")
}
