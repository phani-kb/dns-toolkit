package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsolidateCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, consolidateCmd)
	assert.Equal(t, "consolidate", consolidateCmd.Use)
	assert.Contains(t, consolidateCmd.Short, "Consolidate")
	assert.NotNil(t, consolidateCmd.Run)
}

func TestConsolidateAllCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, consolidateAllCmd)
	assert.Equal(t, "all", consolidateAllCmd.Use)
	assert.Contains(t, consolidateAllCmd.Short, "Consolidate all")
	assert.NotNil(t, consolidateAllCmd.Run)
}

func TestConsolidateFlags(t *testing.T) {
	t.Parallel()

	oldIgnoreAllowlist := ignoreAllowlist
	oldIncludeInvalid := includeInvalid
	oldCalculateChecksum := calculateChecksum
	oldSkipConsolidatedSummary := skipConsolidatedSummary

	ignoreAllowlist = true
	includeInvalid = true
	calculateChecksum = true
	skipConsolidatedSummary = true

	assert.True(t, ignoreAllowlist)
	assert.True(t, includeInvalid)
	assert.True(t, calculateChecksum)
	assert.True(t, skipConsolidatedSummary)

	ignoreAllowlist = oldIgnoreAllowlist
	includeInvalid = oldIncludeInvalid
	calculateChecksum = oldCalculateChecksum
	skipConsolidatedSummary = oldSkipConsolidatedSummary
}
