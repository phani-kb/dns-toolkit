package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestGenericProcessor(t *testing.T) {
	t.Parallel()
	logger := multilog.NewLogger()
	sourceType := "test-source"
	listType := "blocklist"

	processFunc := func(logger *multilog.Logger, content string) ([]string, []string) {
		valid := []string{"example.com", "example.org"}
		invalid := []string{"invalid..domain"}
		return valid, invalid
	}

	gp := processors.NewGenericProcessor(sourceType, listType, processFunc)

	assert.Equal(t, sourceType, gp.GetSourceType(), "GetSourceType should return the correct source type")
	assert.Equal(t, listType, gp.GetListType(), "GetListType should return the correct list type")

	valid, invalid := gp.Process(logger, "dummy content")
	assert.Equal(t, []string{"example.com", "example.org"}, valid, "Valid entries should match expected")
	assert.Equal(t, []string{"invalid..domain"}, invalid, "Invalid entries should match expected")
}

func TestRegisterGenericProcessor(t *testing.T) {
	t.Parallel()
	oldProcessors := processors.Processors
	processors.Processors = processors.NewProcessorRegistry()
	defer func() {
		processors.Processors = oldProcessors
	}()

	sourceType := "test-generic-source"
	processFunc := func(logger *multilog.Logger, content string) ([]string, []string) {
		return []string{"valid.domain"}, []string{"invalid..domain"}
	}

	// Test with both list types (default behavior)
	processors.RegisterGenericProcessor(sourceType, processFunc, false, false)

	processor, exists := processors.Processors.GetProcessor(sourceType, constants.ListTypeBlocklist)
	assert.True(t, exists, "Processor should exist for blocklist")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Processor should have correct source type")
	assert.Equal(t, constants.ListTypeBlocklist, processor.GetListType(), "Processor should have correct list type")

	processor, exists = processors.Processors.GetProcessor(sourceType, constants.ListTypeAllowlist)
	assert.True(t, exists, "Processor should exist for allowlist")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Processor should have correct source type")
	assert.Equal(t, constants.ListTypeAllowlist, processor.GetListType(), "Processor should have correct list type")

	// Test blocklist only
	sourceType = "blocklist-only-source"
	processors.RegisterGenericProcessor(sourceType, processFunc, true, false)

	_, exists = processors.Processors.GetProcessor(sourceType, constants.ListTypeBlocklist)
	assert.True(t, exists, "Processor should exist for blocklist")

	_, exists = processors.Processors.GetProcessor(sourceType, constants.ListTypeAllowlist)
	assert.False(t, exists, "Processor should not exist for allowlist")

	// Test allowlist only
	sourceType = "allowlist-only-source"
	processors.RegisterGenericProcessor(sourceType, processFunc, false, true)

	_, exists = processors.Processors.GetProcessor(sourceType, constants.ListTypeBlocklist)
	assert.False(t, exists, "Processor should not exist for blocklist")

	_, exists = processors.Processors.GetProcessor(sourceType, constants.ListTypeAllowlist)
	assert.True(t, exists, "Processor should exist for allowlist")
}
