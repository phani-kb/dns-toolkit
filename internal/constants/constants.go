package constants

// Default directories for various operations
var (
	DownloadDir = "data/download"
	SummaryDir  = "data"
)

// Folders - Map of folder names to their respective directories
var Folders = map[string]string{
	"download": DownloadDir,
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
	TimestampFormat = "20060102_150405"
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
