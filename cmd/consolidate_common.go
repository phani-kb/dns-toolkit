package cmd

import (
	"context"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/db"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
)

// getFileStrings converts FileInfo slice to string slice
func getFileStrings(fileInfos []c.FileInfo) []string {
	fileStrings := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		fileStrings = append(fileStrings, fileInfo.GetString())
	}
	return fileStrings
}

// ConsolidationParams holds parameters for consolidation functions
type ConsolidationParams struct {
	GenericSourceType string
	ListType          string
	Identifier        string // group or category
	OutputDir         string
	IdentifierField   string // "Group" or "Category"
}

// ProcessingConfig holds configuration for processing consolidation
type ProcessingConfig struct {
	GetFilesFunc       func([]c.ProcessedFile, string) []c.ProcessedFile
	ConsolidateFunc    func(*multilog.Logger, string, string, string, u.StringSet, []c.ProcessedFile) (u.StringSet, c.ConsolidatedSummary) // nolint:lll
	AllowFilterByType  map[string]u.StringSet
	ConsolidatedRepo   *db.ConsolidatedRepo
	DBCtx              context.Context
	PersistMu          *sync.Mutex
	Identifier         string
	IdentifierField    string
	ProcessedFiles     []c.ProcessedFile
	GenericSourceTypes []string
}
