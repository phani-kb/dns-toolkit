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
