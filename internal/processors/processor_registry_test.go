package processors_test

import (
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/dns-toolkit/internal/processors/mocks"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestProcessorRegistry(t *testing.T) {
	registry := processors.NewProcessorRegistry()

	mockProcessor := mocks.NewProcessor(t)
	sourceType := "test-source"
	listType := "blocklist"

	mockProcessor.On("GetSourceType").Maybe().Return(sourceType)
	mockProcessor.On("GetListType").Maybe().Return(listType)

	registry.RegisterProcessor(sourceType, listType, mockProcessor)

	retrievedProcessor, exists := registry.GetProcessor(sourceType, listType)
	assert.True(t, exists, "Processor should exist in registry")
	assert.Equal(t, mockProcessor, retrievedProcessor, "Retrieved processor should match registered processor")

	_, exists = registry.GetProcessor("nonexistent", listType)
	assert.False(t, exists, "Non-existent processor should not be found")

	processorMap := registry.ListProcessors()
	assert.Equal(t, 1, len(processorMap), "Registry should contain 1 processor")

	key := sourceType + ":" + listType
	listedProcessor, exists := processorMap[key]
	assert.True(t, exists, "Processor should exist in the map with the correct key")
	assert.Equal(t, mockProcessor, listedProcessor, "Listed processor should match registered processor")
}

func TestRegisterProcessorHelper(t *testing.T) {
	// Reset the global Processors registry for testing
	oldProcessors := processors.Processors
	processors.Processors = processors.NewProcessorRegistry()
	defer func() {
		processors.Processors = oldProcessors
	}()

	sourceType := "test-helper-source"

	processorFactory := func(srcType, lstType string) processors.Processor {
		mockProcessor := mocks.NewProcessor(t)
		mockProcessor.On("GetSourceType").Return(srcType)
		mockProcessor.On("GetListType").Return(lstType)
		return mockProcessor
	}

	processors.RegisterProcessor(sourceType, processorFactory)

	processor, exists := processors.Processors.GetProcessor(sourceType, constants.ListTypeBlocklist)
	assert.True(t, exists, "Processor should exist for blocklist")
	assert.Equal(t, sourceType, processor.GetSourceType(), "Processor should have correct source type")
	assert.Equal(t, constants.ListTypeBlocklist, processor.GetListType(), "Processor should have blocklist type")
}

func TestRegisterProcessorTypes(t *testing.T) {
	// Reset the global Processors registry for testing
	oldProcessors := processors.Processors
	processors.Processors = processors.NewProcessorRegistry()
	defer func() {
		processors.Processors = oldProcessors
	}()

	sourceType := "test-types-source"
	listTypes := []string{constants.ListTypeBlocklist, constants.ListTypeAllowlist, "custom-list"}

	processorFactory := func(srcType, lstType string) processors.Processor {
		mockProcessor := mocks.NewProcessor(t)
		mockProcessor.On("GetSourceType").Return(srcType)
		mockProcessor.On("GetListType").Return(lstType)
		return mockProcessor
	}

	processors.RegisterProcessorTypes(sourceType, listTypes, processorFactory)

	for _, listType := range listTypes {
		processor, exists := processors.Processors.GetProcessor(sourceType, listType)
		assert.True(t, exists, "Processor should exist for "+listType)
		assert.Equal(t, sourceType, processor.GetSourceType(), "Processor should have correct source type")
		assert.Equal(t, listType, processor.GetListType(), "Processor should have correct list type")
	}
}

func TestProcessor_Process(t *testing.T) {
	t.Helper()
	mockProcessor := &mocks.Processor{}
	mockProcessor.Test(t)

	logger := &multilog.Logger{} // Empty logger for testing
	content := "example.com\ninvalid..domain\nexample.org"
	validEntries := []string{"example.com", "example.org"}
	invalidEntries := []string{"invalid..domain"}

	mockProcessor.On("Process", logger, content).Return(validEntries, invalidEntries)
	mockProcessor.On("GetSourceType").Maybe().Return("test-source")
	mockProcessor.On("GetListType").Maybe().Return("test-list")

	valid, invalid := mockProcessor.Process(logger, content)

	assert.Equal(t, validEntries, valid, "Valid entries should match expected")
	assert.Equal(t, invalidEntries, invalid, "Invalid entries should match expected")
}
