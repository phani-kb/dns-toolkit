package constants

import (
	"regexp"
	"time"
)

const (
	AppName        = "dns-toolkit"
	AppVersion     = "1.0.0"
	AppDescription = "A toolkit for DNS data processing and analysis."
)

// SummaryTypes - Enum-like constants for summary types
const (
	SummaryTypeDownload               = "download"
	SummaryTypeProcessed              = "processed"
	SummaryTypeConsolidated           = "consolidated"
	SummaryTypeConsolidatedGroups     = "consolidated_groups"
	SummaryTypeConsolidatedCategories = "consolidated_categories"
	SummaryTypeOverlap                = "overlap"
	SummaryTypeOverlapDetailed        = "overlap_detailed"
	SummaryTypeTop                    = "top"
	SummaryTypeArchive                = "archive"
	SummaryTypeUnknown                = "unknown"
	SummaryTypeOutput                 = "output"
)

// Default directories for various operations
var (
	DownloadDir               = "data/download"
	ProcessedDir              = "data/processed"
	ConsolidatedDir           = "data/consolidated"
	ConsolidatedGroupsDir     = "data/consolidated_groups"
	ConsolidatedCategoriesDir = "data/consolidated_categories"
	SummaryDir                = "data"
	OverlapDir                = "data/overlap"
	TopDir                    = "data/top"
	ArchiveDir                = "data/archive"
	BackupDir                 = "data/backup"
	OutputDir                 = "data/output"
)

// Folders - Map of folder names to their respective directories
var Folders = map[string]string{
	"download":                DownloadDir,
	"processed":               ProcessedDir,
	"consolidated":            ConsolidatedDir,
	"consolidated_groups":     ConsolidatedGroupsDir,
	"consolidated_categories": ConsolidatedCategoriesDir,
	"summary":                 SummaryDir,
	"overlap":                 OverlapDir,
	"top":                     TopDir,
}

// DefaultSummaryFiles - Map of summary types to their default file names
var DefaultSummaryFiles = map[string]string{
	"download":                "download_summary.json",
	"processed":               "processed_summary.json",
	"consolidated":            "consolidated_summary.json",
	"consolidated_groups":     "consolidated_groups_summary.json",
	"consolidated_categories": "consolidated_categories_summary.json",
	"overlap_detailed":        "overlap_detailed_summary.json",
	"overlap":                 "overlap_summary.json",
	"top":                     "top_summary.json",
	"archive":                 "archive_summary.json",
}

const (
	FrequencyDaily   = "daily"
	FrequencyWeekly  = "weekly"
	FrequencyMonthly = "monthly"

	CategoryAdult          = "adult"
	CategoryMalware        = "malware"
	CategoryAds            = "ads"
	CategoryFamily         = "family"
	CategoryOthers         = "others"
	CategoryDns            = "dns"
	CategoryDoh            = "doh"
	CategorySpam           = "spam"
	CategoryScam           = "scam"
	CategoryPhishing       = "phishing"
	CategoryCryptocurrency = "cryptocurrency"
	CategoryTrackers       = "trackers"
	CategorySocial         = "social"
	CategoryAnnoyance      = "annoyance"
	CategoryFake           = "fake"
	CategoryFakeNews       = "fakenews"
	CategoryGambling       = "gambling"
	CategoryThreat         = "threat"
	CategoryPrivacy        = "privacy"
	CategorySecurity       = "security"
	CategoryMalicious      = "malicious"
	CategoryAnonymizer     = "anonymizer"
	CategoryTopDomains     = "topdomains"
	CategoryNewDomains     = "newdomains"
	CategoryMobile         = "mobile"
	CategoryProxy          = "proxy"
	CategoryTrojan         = "trojan"
	CategoryRansomware     = "ransomware"
	CategoryBotnet         = "botnet"
	CategoryExploit        = "exploit"

	GroupMini   = "mini"
	GroupLite   = "lite"
	GroupNormal = "normal"
	GroupBig    = "big"
)

const (
	SourceTypeIpv4         = "ipv4"
	SourceTypeIpv6         = "ipv6"
	SourceTypeCidrIpv4     = "cidr_ipv4"
	SourceTypeDomain       = "domain"
	SourceTypeAdguard      = "adguard"
	SourceTypeIpv4Hostname = "ipv4_hostname"
	SourceTypeMixed        = "mixed"
	SourceTypeHostname     = "hostname"
	SourceTypeUnknown      = "unknown"

	ListTypeBlocklist = "blocklist"
	ListTypeAllowlist = "allowlist"
)

var (
	ValidTypes = map[string]bool{
		SourceTypeIpv4:         true,
		SourceTypeIpv6:         true,
		SourceTypeDomain:       true,
		SourceTypeAdguard:      true,
		SourceTypeIpv4Hostname: true,
		SourceTypeHostname:     true,
		SourceTypeUnknown:      true,
	}
	ValidListTypes = map[string]bool{
		ListTypeBlocklist: true,
		ListTypeAllowlist: true,
	}
)

var ListTypes = []string{
	ListTypeBlocklist,
	ListTypeAllowlist,
}

var ListTypeMap = map[string]string{
	ListTypeBlocklist: "BL",
	ListTypeAllowlist: "AL",
}

var SourceTypeRegexMap = map[string]*regexp.Regexp{
	SourceTypeIpv4:     regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`),
	SourceTypeIpv6:     regexp.MustCompile(`\b([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\b`),
	SourceTypeCidrIpv4: regexp.MustCompile(`\b\d{1,3}(\.\d{1,3}){3}/\d{1,2}\b`),
	SourceTypeDomain:   regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`),
	SourceTypeIpv4Hostname: regexp.MustCompile(
		`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\s+([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`,
	),
}
var (
	GenericSourceTypes = []string{
		SourceTypeIpv4,
		SourceTypeIpv6,
		SourceTypeCidrIpv4,
		SourceTypeDomain,
		SourceTypeAdguard,
	}

	GenericSourceTypeAliases = map[string]string{
		SourceTypeHostname: SourceTypeDomain,
	}
)

var (
	ValidFrequencies = map[string]bool{
		FrequencyDaily:   true,
		FrequencyWeekly:  true,
		FrequencyMonthly: true,
	}
	ValidCategories = map[string]bool{
		CategoryAdult:          true,
		CategoryMalware:        true,
		CategoryAds:            true,
		CategoryFamily:         true,
		CategoryOthers:         true,
		CategorySocial:         true,
		CategoryFake:           true,
		CategoryFakeNews:       true,
		CategoryGambling:       true,
		CategoryPhishing:       true,
		CategoryCryptocurrency: true,
		CategorySpam:           true,
		CategoryScam:           true,
		CategoryDns:            true,
		CategoryDoh:            true,
		CategoryTrackers:       true,
		CategoryAnnoyance:      true,
		CategoryThreat:         true,
		CategoryPrivacy:        true,
		CategorySecurity:       true,
		CategoryMalicious:      true,
		CategoryAnonymizer:     true,
		CategoryTopDomains:     true,
		CategoryNewDomains:     true,
		CategoryMobile:         true,
		CategoryProxy:          true,
		CategoryTrojan:         true,
		CategoryRansomware:     true,
		CategoryBotnet:         true,
		CategoryExploit:        true,
	}

	ValidGroups = map[string]bool{
		GroupMini:   true,
		GroupLite:   true,
		GroupNormal: true,
		GroupBig:    true,
	}

	DefaultGroup = GroupBig

	GroupIdMap = map[string]int{
		GroupMini:   0,
		GroupLite:   1,
		GroupNormal: 2,
		GroupBig:    3,
	}

	SizeGroups = []string{
		GroupMini,
		GroupLite,
		GroupNormal,
		GroupBig,
	}
)

var UrlRegex = regexp.MustCompile(`^(https?|ftp|file)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`)

var ArchiveExtensions = []string{".zip", ".tar.gz"}

const (
	SearchProcessedFile    = "processed"
	SearchConsolidatedFile = "consolidated"
)

const (
	MaxSampleLinesToCategorize    = 100
	TimestampFormat               = "20060102_150405"
	ArchiveFileTimestampFormat    = "20060102_150405"
	BackupFileTimestampFormat     = "20060102_150405"
	DownloadInterval              = 2000 * time.Millisecond
	DefaultHashAlgorithm          = "md5"
	DefaultMaxRetries             = 3
	DefaultRetryDelayInSeconds    = 10
	DefaultClientTimeoutInSeconds = 30
	EntryAverageCharLength        = 30
	EntryMinCharLength            = 6
	MaxPreallocEntries            = 10_000_000
	MinPreallocEntries            = 10_000
	MinFilesForParallelProcessing = 10
	MaxEntryLength                = 255
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
