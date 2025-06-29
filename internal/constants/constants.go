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
	ProfilesDir               = "data/profiles"
	OutputGroupsDir           = OutputDir + "/groups"
	OutputCategoriesDir       = OutputDir + "/categories"
	OutputIgnoredDir          = OutputDir + "/ignored"
	OutputTopDir              = OutputDir + "/top"
	OutputSummariesDir        = OutputDir + "/summaries"
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
	"archive":                 ArchiveDir,
	"backup":                  BackupDir,
	"output":                  OutputDir,
	"profiles":                ProfilesDir,
	"output_ignored":          OutputIgnoredDir,
	"output_groups":           OutputGroupsDir,
	"output_categories":       OutputCategoriesDir,
	"output_top":              OutputTopDir,
	"output_summaries":        OutputSummariesDir,
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
	SourceTypeIpv4                       = "ipv4"
	SourceTypeIpv4RangeExpand            = "ipv4_range_expand"
	SourceTypeIpv6                       = "ipv6"
	SourceTypeCidrIpv4                   = "cidr_ipv4"
	SourceTypeIpv4CidrExpand             = "ipv4_cidr_expand"
	SourceTypeDomain                     = "domain"
	SourceTypeDomainComment              = "domain_comment"
	SourceTypeDomainFinder               = "domain_finder"
	SourceTypeAdguard                    = "adguard"
	SourceTypeIpv4Hostname               = "ipv4_hostname"
	SourceTypeMixed                      = "mixed"
	SourceTypeHostname                   = "hostname"
	SourceTypeUnknown                    = "unknown"
	SourceTypeDomainAdguard              = "domain_adguard"
	SourceTypeDomainCsvHttpUrlFind       = "domain_csv_http_url_find"
	SourceTypeDomainCustomCsvBlackbook   = "domain_custom_csv_blackbook"
	SourceTypeDomainCustomCsvMaltrail    = "domain_custom_csv_maltrail"
	SourceTypeDomainCustomHtmlCcam       = "domain_custom_html_ccam"
	SourceTypeDomainHttpUrl              = "domain_http_url"
	SourceTypeDomainUrl                  = "domain_url"
	SourceTypeDomainWithCommentSuffix    = "domain_with_comment_suffix"
	SourceTypeIpv4CustomHtmlCcam         = "ipv4_custom_html_ccam"
	SourceTypeIpv4Find                   = "ipv4_find"
	SourceTypeIpv4HttpUrl                = "ipv4_http_url"
	SourceTypeIpv4Url                    = "ipv4_url"
	SourceTypeIpv6Find                   = "ipv6_find"
	SourceTypeIpv6Htaccess               = "ipv6_htaccess"
	SourceTypeTopDomains                 = "domain_top"
	SourceTypeDomainCustomHtmlPuppyScams = "domain_custom_html_puppyscams"

	ListTypeBlocklist = "blocklist"
	ListTypeAllowlist = "allowlist"
)

var (
	ValidTypes = map[string]bool{
		SourceTypeIpv4:                       true,
		SourceTypeIpv4RangeExpand:            true,
		SourceTypeIpv6:                       true,
		SourceTypeCidrIpv4:                   true,
		SourceTypeIpv4CidrExpand:             true,
		SourceTypeDomain:                     true,
		SourceTypeDomainComment:              true,
		SourceTypeDomainFinder:               true,
		SourceTypeAdguard:                    true,
		SourceTypeIpv4Hostname:               true,
		SourceTypeMixed:                      true,
		SourceTypeHostname:                   true,
		SourceTypeUnknown:                    true,
		SourceTypeDomainAdguard:              true,
		SourceTypeDomainCsvHttpUrlFind:       true,
		SourceTypeDomainCustomCsvBlackbook:   true,
		SourceTypeDomainCustomCsvMaltrail:    true,
		SourceTypeDomainCustomHtmlCcam:       true,
		SourceTypeDomainHttpUrl:              true,
		SourceTypeDomainUrl:                  true,
		SourceTypeDomainWithCommentSuffix:    true,
		SourceTypeIpv4CustomHtmlCcam:         true,
		SourceTypeIpv4Find:                   true,
		SourceTypeIpv4HttpUrl:                true,
		SourceTypeIpv4Url:                    true,
		SourceTypeIpv6Find:                   true,
		SourceTypeIpv6Htaccess:               true,
		SourceTypeTopDomains:                 true,
		SourceTypeDomainCustomHtmlPuppyScams: true,
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
	SourceTypeDomainFinder: regexp.MustCompile(`([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}`),
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
		SourceTypeHostname:        SourceTypeDomain,
		SourceTypeIpv4RangeExpand: SourceTypeIpv4,
		SourceTypeIpv4CidrExpand:  SourceTypeIpv4,
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
	DefaultBlockProfileRate       = 1000
	MinFilesForParallelProcessing = 10
	MaxEntryLength                = 255
)

var DefaultMinSourcesRange = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

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

const PunycodePrefix = "xn--"

const ContentSeparator = "###\n"

// SummaryTypesMap maps summary type constants to their string values
var SummaryTypesMap = map[string]string{
	SummaryTypeDownload:               SummaryTypeDownload,
	SummaryTypeProcessed:              SummaryTypeProcessed,
	SummaryTypeConsolidated:           SummaryTypeConsolidated,
	SummaryTypeConsolidatedGroups:     SummaryTypeConsolidatedGroups,
	SummaryTypeConsolidatedCategories: SummaryTypeConsolidatedCategories,
	SummaryTypeOverlap:                SummaryTypeOverlap,
	SummaryTypeOverlapDetailed:        SummaryTypeOverlapDetailed,
	SummaryTypeTop:                    SummaryTypeTop,
	SummaryTypeArchive:                SummaryTypeArchive,
	SummaryTypeOutput:                 SummaryTypeOutput,
}

// AllSummaryTypes contains all valid summary types as a slice
var AllSummaryTypes = []string{
	SummaryTypeDownload,
	SummaryTypeProcessed,
	SummaryTypeConsolidated,
	SummaryTypeConsolidatedGroups,
	SummaryTypeConsolidatedCategories,
	SummaryTypeOverlap,
	SummaryTypeOverlapDetailed,
	SummaryTypeTop,
	SummaryTypeArchive,
	SummaryTypeOutput,
}

// FolderToSummaryTypeMap maps folder names to their corresponding summary types
var FolderToSummaryTypeMap = map[string]string{
	"download":                SummaryTypeDownload,
	"processed":               SummaryTypeProcessed,
	"consolidated":            SummaryTypeConsolidated,
	"consolidated_groups":     SummaryTypeConsolidatedGroups,
	"consolidated_categories": SummaryTypeConsolidatedCategories,
	"overlap":                 SummaryTypeOverlap,
	"top":                     SummaryTypeTop,
	"archive":                 SummaryTypeArchive,
	"output":                  SummaryTypeOutput,
	"output_ignored":          SummaryTypeOutput,
	"output_groups":           SummaryTypeConsolidatedGroups,
	"output_categories":       SummaryTypeConsolidatedCategories,
	"output_top":              SummaryTypeTop,
	"output_summaries":        SummaryTypeOutput,
}

// SummaryTypesWithTemplateMap SummaryTypes with template strings
var SummaryTypesWithTemplateMap = map[string]string{
	SummaryTypeConsolidated:           DefaultSummaryFiles[SummaryTypeConsolidated],
	SummaryTypeConsolidatedGroups:     DefaultSummaryFiles[SummaryTypeConsolidatedGroups],
	SummaryTypeConsolidatedCategories: DefaultSummaryFiles[SummaryTypeConsolidatedCategories],
	SummaryTypeTop:                    DefaultSummaryFiles[SummaryTypeTop],
}

// SummaryTypesDirMap maps summary types to their respective directories
var SummaryTypesDirMap = map[string]string{
	SummaryTypeDownload:               DownloadDir,
	SummaryTypeProcessed:              ProcessedDir,
	SummaryTypeOverlap:                OverlapDir,
	SummaryTypeConsolidated:           ConsolidatedDir,
	SummaryTypeConsolidatedGroups:     ConsolidatedGroupsDir,
	SummaryTypeConsolidatedCategories: ConsolidatedCategoriesDir,
	SummaryTypeTop:                    TopDir,
	SummaryTypeArchive:                ArchiveDir,
	SummaryTypeOutput:                 OutputDir,
}

// SummaryTypesOutputDirMap maps summary types to their output directories
var SummaryTypesOutputDirMap = map[string]string{
	SummaryTypeConsolidated:           OutputDir,
	SummaryTypeConsolidatedGroups:     OutputGroupsDir,
	SummaryTypeConsolidatedCategories: OutputCategoriesDir,
	SummaryTypeTop:                    OutputTopDir,
	SummaryTypeOutput:                 OutputDir,
}

var SummaryTypesOutputToSkipMap = map[string]bool{
	SummaryTypeDownload:  true,
	SummaryTypeProcessed: true,
	SummaryTypeOverlap:   true,
	SummaryTypeArchive:   true,
	SummaryTypeOutput:    false,
}

// SummaryTypesOutputSummaryFileMap maps summary types to their output file names
var SummaryTypesOutputSummaryFileMap = map[string]string{
	SummaryTypeDownload:               DefaultSummaryFiles[SummaryTypeDownload],
	SummaryTypeProcessed:              DefaultSummaryFiles[SummaryTypeProcessed],
	SummaryTypeConsolidated:           DefaultSummaryFiles[SummaryTypeConsolidated],
	SummaryTypeConsolidatedGroups:     DefaultSummaryFiles[SummaryTypeConsolidatedGroups],
	SummaryTypeConsolidatedCategories: DefaultSummaryFiles[SummaryTypeConsolidatedCategories],
	SummaryTypeOverlap:                DefaultSummaryFiles[SummaryTypeOverlap],
	SummaryTypeTop:                    DefaultSummaryFiles[SummaryTypeTop],
	SummaryTypeArchive:                DefaultSummaryFiles[SummaryTypeArchive],
}

// SummaryTypesOutputSummaryFileToSkipMap maps summary types to their output file names that should be skipped
var SummaryTypesOutputSummaryFileToSkipMap = map[string]bool{
	SummaryTypeOutput:  true,
	SummaryTypeArchive: true,
}

// ArchiveFoldersToSkipMap contains top level folders that should be skipped during archiving
var ArchiveFoldersToSkipMap = map[string]bool{
	DownloadDir:               true,
	ProcessedDir:              true,
	ConsolidatedDir:           true,
	ConsolidatedGroupsDir:     true,
	ConsolidatedCategoriesDir: true,
	OverlapDir:                true,
	TopDir:                    true,
	ArchiveDir:                true,
	BackupDir:                 true,
	ProfilesDir:               true,
}

var SummaryTypesToDeleteAfterOutputGenerationMap = map[string]bool{
	SummaryTypeDownload:               false,
	SummaryTypeProcessed:              true,
	SummaryTypeConsolidated:           true,
	SummaryTypeConsolidatedGroups:     true,
	SummaryTypeConsolidatedCategories: true,
	SummaryTypeOverlap:                true,
	SummaryTypeTop:                    true,
	SummaryTypeArchive:                false,
}
