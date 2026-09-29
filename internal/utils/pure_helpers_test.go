package utils

import (
	"strings"
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestSplitAndSortCSV(t *testing.T) {
	assert.Nil(t, SplitAndSortCSV(""))
	assert.Equal(t, []string{"a", "b", "c"}, SplitAndSortCSV("c, a, b"))
	assert.Equal(t, []string{"a"}, SplitAndSortCSV(" a , , "))
}

func TestValidateArchiveFilePath(t *testing.T) {
	assert.Error(t, validateArchiveFilePath("/abs/path"))
	assert.Error(t, validateArchiveFilePath("../up"))
	assert.Error(t, validateArchiveFilePath("../../up"))
	assert.NoError(t, validateArchiveFilePath("relative/file.txt"))
}

func TestGetFilePathForSummaryConsolidatedAndTop(t *testing.T) {
	consolidated := c.ConsolidatedSummary{Filepath: "/tmp/consolidated.json"}
	assert.Equal(t, "/tmp/consolidated.json", getFilePathForSummary(consolidated, constants.SummaryTypeConsolidated))
	assert.Equal(
		t,
		"/tmp/consolidated.json",
		getFilePathForSummary(consolidated, constants.SummaryTypeConsolidatedGroups),
	)

	top := c.TopSummary{Filepath: "/tmp/top.json"}
	assert.Equal(t, "/tmp/top.json", getFilePathForSummary(top, constants.SummaryTypeTop))

	assert.Equal(t, "", getFilePathForSummary(consolidated, "unknown"))
}

func TestIsValidASCIILabel(t *testing.T) {
	assert.True(t, isValidASCIILabel("example"))
	assert.True(t, isValidASCIILabel("valid-123"))
	assert.False(t, isValidASCIILabel(""))
	assert.False(t, isValidASCIILabel("-leading"))
	assert.False(t, isValidASCIILabel("trailing-"))
	assert.False(t, isValidASCIILabel("has.dot"))
	assert.False(t, isValidASCIILabel(strings.Repeat("a", 64)))
}

func TestContainsLetter(t *testing.T) {
	assert.True(t, containsLetter("abc"))
	assert.True(t, containsLetter("123a"))
	assert.False(t, containsLetter("123"))
	assert.False(t, containsLetter(""))
}
