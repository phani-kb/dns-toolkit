package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlapCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, overlapCmd)
	assert.Equal(t, "overlap", overlapCmd.Use)
	assert.Contains(t, overlapCmd.Short, "overlap")
	assert.NotNil(t, overlapCmd.Run)
}

func TestOverlapCommandFlags(t *testing.T) {
	t.Parallel()

	flags := overlapCmd.Flags()

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
