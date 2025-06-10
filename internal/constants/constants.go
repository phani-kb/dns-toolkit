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
)

// Folders - Map of folder names to their respective directories
var Folders = map[string]string{
	"download":                DownloadDir,
	"processed":               ProcessedDir,
	"consolidated":            ConsolidatedDir,
	"consolidated_groups":     ConsolidatedGroupsDir,
	"consolidated_categories": ConsolidatedCategoriesDir,
	"summary":                 SummaryDir,
}

// DefaultSummaryFiles - Map of summary types to their default file names
var DefaultSummaryFiles = map[string]string{
	"download":                "download_summary.json",
	"processed":               "processed_summary.json",
	"consolidated":            "consolidated_summary.json",
	"consolidated_groups":     "consolidated_groups_summary.json",
	"consolidated_categories": "consolidated_categories_summary.json",
}

const (
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

const (
	MaxSampleLinesToCategorize = 100
	TimestampFormat            = "20060102_150405"
	DownloadInterval           = 2000 * time.Millisecond
	DefaultHashAlgorithm       = "md5"
	DefaultMaxRetries          = 3
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
