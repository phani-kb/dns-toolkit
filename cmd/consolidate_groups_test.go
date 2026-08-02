package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsolidateGroupsCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, consolidateGroupsCmd)
	assert.Equal(t, "groups", consolidateGroupsCmd.Use)
	assert.Contains(t, consolidateGroupsCmd.Short, "sized consolidated lists")
	assert.NotNil(t, consolidateGroupsCmd.Run)
}
