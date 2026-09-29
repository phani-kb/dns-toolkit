package db

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaChecksum_UsesRawSchemaFileContent(t *testing.T) {
	expected := fmt.Sprintf("%x", sha256.Sum256([]byte(schemaSQL)))
	assert.Equal(t, expected, SchemaChecksum())
}

func TestSchemaTableNamesMatchConstants(t *testing.T) {
	const createTablePrefix = "create table if not exists "
	tablesFromSchema := make([]string, 0)
	for _, line := range strings.Split(schemaSQL, "\n") {
		trimmedLower := strings.ToLower(strings.TrimSpace(line))
		if !strings.HasPrefix(trimmedLower, createTablePrefix) {
			continue
		}

		remainder := strings.TrimPrefix(trimmedLower, createTablePrefix)
		name := remainder
		if idx := strings.IndexAny(remainder, " (\t\r\n"); idx >= 0 {
			name = remainder[:idx]
		}
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		tablesFromSchema = append(tablesFromSchema, name)
	}

	tablesFromConstants := make([]string, 0)
	for _, name := range schemaTableNamesFromConstants() {
		tablesFromConstants = append(tablesFromConstants, strings.ToLower(name))
	}

	assert.ElementsMatch(t, tablesFromConstants, tablesFromSchema,
		"table constants and schema.sql tables are out of sync")
}
