package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopEntriesCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, topEntriesCmd)
	assert.Equal(t, "top", topEntriesCmd.Use)
	assert.Contains(t, topEntriesCmd.Short, "top entry")
	assert.NotNil(t, topEntriesCmd.Run)
}

func TestTopEntriesFlags(t *testing.T) {
	t.Parallel()

	flags := topEntriesCmd.Flags()

	minSourcesFlag := flags.Lookup("min-sources")
	assert.NotNil(t, minSourcesFlag)

	maxEntriesFlag := flags.Lookup("max-entries")
	assert.NotNil(t, maxEntriesFlag)

	cpuProfileFlag := flags.Lookup("cpu-profile")
	assert.NotNil(t, cpuProfileFlag)

	memProfileFlag := flags.Lookup("mem-profile")
	assert.NotNil(t, memProfileFlag)

	goroutineProfileFlag := flags.Lookup("goroutine-profile")
	assert.NotNil(t, goroutineProfileFlag)

	blockProfileFlag := flags.Lookup("block-profile")
	assert.NotNil(t, blockProfileFlag)

	profileDirFlag := flags.Lookup("profile-dir")
	assert.NotNil(t, profileDirFlag)
}
