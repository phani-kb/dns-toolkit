package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsolidateCategoriesCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, consolidateCategoriesCmd)
	assert.Equal(t, "categories", consolidateCategoriesCmd.Use)
	assert.Contains(t, consolidateCategoriesCmd.Short, "category-based")
	assert.NotNil(t, consolidateCategoriesCmd.Run)
}
