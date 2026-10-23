package downloaders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandURLSecrets(t *testing.T) {
	t.Setenv("TEST_URL_KEY", "abc123")

	got, err := expandURLSecrets("https://example.com/exports/${TEST_URL_KEY}/hostfile.txt")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/exports/abc123/hostfile.txt", got)

	got, err = expandURLSecrets("https://example.com/hostfile.txt")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/hostfile.txt", got)

	_, err = expandURLSecrets("https://example.com/exports/${TEST_URL_MISSING}/hostfile.txt")
	assert.ErrorContains(t, err, "TEST_URL_MISSING")
}
