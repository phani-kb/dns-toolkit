package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/stretchr/testify/assert"
)

func TestBaseProcessor(t *testing.T) {
	sourceType := "test-source"
	listType := "blocklist"

	bp := processors.NewBaseProcessor(sourceType, listType)

	assert.Equal(t, sourceType, bp.GetSourceType(), "GetSourceType should return the source type")
	assert.Equal(t, listType, bp.GetListType(), "GetListType should return the list type")
}
