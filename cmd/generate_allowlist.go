package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/phani-kb/multilog"
	"github.com/spf13/cobra"
)

var overwrite bool

var generateAllowlistCmd = &cobra.Command{
	Use:   "allowlist",
	Short: "Generate allowlist files from source URL domains",
	Long:  "Generate allowlist files including: domains, adguard rules, and ipv4 addresses resolved from domains.",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("DNS_TOOLKIT_TEST_MODE") == "true" {
			return
		}

		logger := Logger
		logger.Info("Starting allowlist generation...")
		generateAllowlist(logger)
		logger.Info("Finished allowlist generation.")
	},
}

const (
	startSource = "### START sources"
	endSource   = "### END sources"
)

func adgFormat(s string) string {
	if s != "" {
		return "@@||" + s + "^"
	}
	return ""
}

func strFormat(s string) string {
	return s
}

func generateAllowlist(logger *multilog.Logger) {
	domainsFile := constants.AllowlistFilesMap[constants.SourceTypeDomain]
	ipv4File := constants.AllowlistFilesMap[constants.SourceTypeIpv4]
	adguardFile := constants.AllowlistFilesMap[constants.SourceTypeAdguard]
	customALDomainsFile := constants.CustomAllowlistFilesMap[constants.SourceTypeDomain]

	logger.Info("Extracting domains from source URLs...")
	sourceDomains := extractSourceDomains()
	customDomains := loadCustomElements(logger, customALDomainsFile)
	allDomains := combineDomains(customDomains, sourceDomains)

	if err := writeAllowlistWithStructure(logger, domainsFile, customDomains, sourceDomains, strFormat); err != nil {
		logger.Error("Failed to write domains file", "error", err)
		return
	}
	logger.Info("Domains written to file", "file", domainsFile, "count", len(allDomains))

	if err := generateAdGuardRules(logger, adguardFile, customDomains, sourceDomains); err != nil {
		logger.Error("Failed to generate AdGuard format", "error", err)
		return
	}
	logger.Info("AdGuard format written to file", "file", adguardFile)

	customIPv4File := constants.CustomAllowlistFilesMap[constants.SourceTypeIpv4]
	customIPv4 := loadCustomElements(logger, customIPv4File)
	if err := generateIPv4Addresses(logger, ipv4File, customIPv4, customDomains, sourceDomains); err != nil {
		logger.Error("Failed to generate IPv4 addresses", "error", err)
		return
	}
	logger.Info("IPv4 addresses written to file", "file", ipv4File)

	backupExistingFiles(logger, overwrite)
}

func backupExistingFiles(logger *multilog.Logger, overwrite bool) {
	if !overwrite {
		return
	}

	logger.Info("Backing up existing allowlist files...")

	timestamp := time.Now().Format(constants.TimestampFormat)
	backupDir := filepath.Join("data", "backup", "allowlist_backup_"+timestamp)

	if err := os.MkdirAll(backupDir, 0755); err != nil {
		logger.Error("Failed to create backup directory", "error", err)
		return
	}

	filesToBackup := []string{
		constants.CustomAllowlistFilesMap[constants.SourceTypeDomain],
		constants.CustomAllowlistFilesMap[constants.SourceTypeIpv4],
		constants.CustomAllowlistFilesMap[constants.SourceTypeAdguard],
	}

	for _, file := range filesToBackup {
		if _, err := os.Stat(file); err == nil {
			backupFile := filepath.Join(backupDir, filepath.Base(file))
			if err := utils.CopyFile(logger, file, backupFile); err != nil {
				logger.Warn("Failed to backup file", "file", file, "error", err)
			} else {
				logger.Info("Backed up file", "file", file, "backup", backupFile)
			}
		}
	}

	logger.Info("Backup completed", "directory", backupDir)

	sourceTargetMap := make(map[string]string)
	for sourceType, source := range constants.AllowlistFilesMap {
		sourceTargetMap[source] = constants.CustomAllowlistFilesMap[sourceType]
	}

	for source, target := range sourceTargetMap {
		if _, err := os.Stat(source); err == nil {
			if err := utils.CopyFile(logger, source, target); err != nil {
				logger.Warn("Failed to copy file", "source", source, "target", target, "error", err)
			} else {
				logger.Info("Copied file", "source", source, "target", target)
			}
		}
	}
}

func extractSourceDomains() []string {
	domains := make(map[string]bool)

	for _, sourcesConfig := range SourcesConfigs {
		for _, source := range sourcesConfig.Sources {
			if source.Disabled || strings.HasPrefix(source.URL, "file://") || source.URL == "" {
				continue
			}

			domain := extractDomainFromURL(source.URL)
			if domain != "" {
				domains[domain] = true
			}
		}
	}

	result := make([]string, 0, len(domains))
	for domain := range domains {
		result = append(result, domain)
	}
	sort.Strings(result)

	return result
}

func extractDomainFromURL(url string) string {
	re := regexp.MustCompile(constants.DomainHttpUrlRegex)
	matches := re.FindStringSubmatch(url)
	if len(matches) > 1 {
		domain := matches[1]
		domain = strings.TrimRight(domain, ".")
		if utils.IsDomain(domain) {
			return domain
		}
	}
	return ""
}

func loadCustomElements(logger *multilog.Logger, filename string) []string {
	domains := make([]string, 0)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		logger.Debug("Custom domains file does not exist", "file", filename)
		return domains
	}

	file, err := os.Open(filename)
	if err != nil {
		logger.Warn("Failed to open custom domains file", "file", filename, "error", err)
		return domains
	}
	defer utils.CloseFile(logger, file)

	scanner := bufio.NewScanner(file)
	inSourceDomains := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == startSource {
			break
		}

		if !utils.IsComment(line) && !inSourceDomains {
			domains = append(domains, line)
		}
	}
	sort.Strings(domains)

	logger.Debug("Loaded custom domains", "count", len(domains))
	return domains
}

func combineDomains(customDomains, sourceDomains []string) []string {
	allDomainsMap := make(map[string]bool)

	for _, domain := range customDomains {
		allDomainsMap[domain] = true
	}

	for _, domain := range sourceDomains {
		if !allDomainsMap[domain] {
			allDomainsMap[domain] = true
		}
	}

	result := make([]string, 0, len(allDomainsMap))
	for domain := range allDomainsMap {
		result = append(result, domain)
	}
	sort.Strings(result)

	return result
}

func writeAllowlistWithStructure(
	logger *multilog.Logger,
	filename string,
	customElements, sourceElements []string,
	formatFn func(string) string,
) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer utils.CloseFile(logger, file)

	writer := bufio.NewWriter(file)
	for _, elem := range customElements {
		if _, err := writer.WriteString(strFormat(elem) + "\n"); err != nil { // write existing elements as is
			return fmt.Errorf("failed to write custom element: %w", err)
		}
	}

	if len(sourceElements) > 0 {
		if _, err := writer.WriteString(startSource + "\n"); err != nil {
			return fmt.Errorf("failed to write separator: %w", err)
		}
		for _, elem := range sourceElements {
			if _, err := writer.WriteString(formatFn(elem) + "\n"); err != nil {
				return fmt.Errorf("failed to write source element: %w", err)
			}
		}
		if _, err := writer.WriteString(endSource + "\n"); err != nil {
			return fmt.Errorf("failed to write end marker: %w", err)
		}
	}
	return writer.Flush()
}

func generateAdGuardRules(logger *multilog.Logger, filename string, customDomains, sourceDomains []string) error {
	customAdguardRules := loadCustomElements(logger, constants.CustomAllowlistFilesMap[constants.SourceTypeAdguard])
	logger.Infof("%v existing custom AdGuard rules", len(customAdguardRules))

	rulesSet := utils.NewStringSet(customAdguardRules)

	for _, domain := range customDomains {
		adgRule := adgFormat(domain)
		rulesSet.Add(adgRule)
	}

	return writeAllowlistWithStructure(logger, filename, rulesSet.ToSliceSorted(), sourceDomains, adgFormat)
}

func generateIPv4Addresses(
	logger *multilog.Logger,
	filename string,
	customIPs, customDomains, sourceDomains []string,
) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	resolvedCustomIPs := getResolvedIPs(logger, customDomains)

	logger.Info("Existing custom IPs", "count", len(customIPs))
	customIPs = append(customIPs, resolvedCustomIPs...)
	customIPs = utils.RemoveDuplicates(customIPs)

	resolvedIPs := getResolvedIPs(logger, sourceDomains)
	uniqueIPs := utils.RemoveDuplicates(resolvedIPs)

	return writeAllowlistWithStructure(logger, filename, customIPs, uniqueIPs, strFormat)
}

func getResolvedIPs(logger *multilog.Logger, domains []string) []string {
	totalDomains := len(domains)
	logger.Infof("Resolving IPv4 addresses for %v domains...", totalDomains)

	resolvedIPs, failedDomains := utils.ResolveDomainsToIPv4(logger, domains)

	logger.Infof("Resolved IPv4 addresses count: %v", len(resolvedIPs))
	if len(failedDomains) > 0 {
		logger.Warnf("Failed to resolve %v domains", len(failedDomains))
		for _, domain := range failedDomains {
			logger.Warnf("Domain: %s", domain)
		}
	}

	return resolvedIPs
}

func init() {
	generateCmd.AddCommand(generateAllowlistCmd)
	generateAllowlistCmd.Flags().BoolVar(&overwrite, "overwrite", false, "Overwrite existing allowlist files")
}
