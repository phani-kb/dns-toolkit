package constants

const (
	AppName        = "dns-toolkit"
	AppVersion     = "1.0.0"
	AppDescription = "A toolkit for DNS data processing and analysis."
)

// Default directories for various operations
var (
	DownloadDir               = "data/download"
	ProcessedDir              = "data/processed"
	ConsolidatedDir           = "data/consolidated"
	ConsolidatedGroupsDir     = "data/consolidated_groups"
	ConsolidatedCategoriesDir = "data/consolidated_categories"
	SummaryDir                = "data"
)

// Folders - Map of folder names to their respective directories
var Folders = map[string]string{
	"download":     DownloadDir,
	"processed":    ProcessedDir,
	"consolidated": ConsolidatedDir,
}

// DefaultSummaryFiles - Map of summary types to their default file names
var DefaultSummaryFiles = map[string]string{
	"download":            "download_summary.json",
	"processed":           "processed_summary.json",
	"consolidated":        "consolidated_summary.json",
	"consolidated_groups": "consolidated_groups_summary.json",
}

const (
	ListTypeBlocklist = "blocklist"
	ListTypeAllowlist = "allowlist"
)

var ListTypes = []string{
	ListTypeBlocklist,
	ListTypeAllowlist,
}

var ListTypeMap = map[string]string{
	ListTypeBlocklist: "BL",
	ListTypeAllowlist: "AL",
}

const (
	TimestampFormat            = "20060102_150405"
	DefaultRetryDelayInSeconds = 10
)

// CommentPrefixes Comment prefixes used in various formats
var CommentPrefixes = []string{
	"#",
	"!",
	";",
	"//",
	"--",
	"%",
	"rem ", // windows batch file, space is important
	"[",
	"<!--",
	"-->",
	"/*",
	"*",
	"*/",
	"<",
	">",
}
