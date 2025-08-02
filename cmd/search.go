package cmd

import (
	"bufio"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	u "github.com/phani-kb/dns-toolkit/internal/utils"
	"github.com/spf13/cobra"
)

var (
	exactMatch         bool
	searchProcessed    bool
	searchConsolidated bool
	performDNSLookup   bool
	performCNAMELookup bool
	searchAguard       bool
)

var searchCmd = &cobra.Command{
	Use:   "search [domain or IP]",
	Short: "Search for a domain or IP in the processed files",
	Long:  `Search for a given domain among the valid processed domain files and report the sources in which it was found. Also looks for its IP address among the valid IPv4 processed files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(args[0])
		Logger.Infof("Searching for: %s", query)

		isIP := net.ParseIP(query) != nil

		// Collect IP addresses and CNAMEs based on a query
		ipAddresses, cnames := collectQueryData(query, isIP)

		// Search for the query in files
		domainResults, ipResults, cnameResults := searchInAllFiles(query, isIP, ipAddresses, cnames)

		// Display search results
		displaySearchResults(query, isIP, domainResults, ipResults, cnameResults, ipAddresses, cnames)
	},
}

// collectQueryData collects IP addresses and CNAME records based on the query
func collectQueryData(query string, isIP bool) (u.StringSet, []string) {
	ipAddresses := u.NewStringSet(nil)
	var cnames []string

	if isIP {
		ipAddresses.Add(query)
	} else {
		if performCNAMELookup {
			cname, err := net.LookupCNAME(query)
			if err != nil {
				Logger.Warnf("Could not lookup CNAME for domain '%s': %v", query, err)
			} else if cname != query && cname != query+"." {
				cnames = append(cnames, strings.TrimSuffix(cname, "."))
				Logger.Infof("Domain %s has CNAME: %s", query, strings.TrimSuffix(cname, "."))
			}
		}

		if performDNSLookup {
			lookupDomains := []string{query}
			lookupDomains = append(lookupDomains, cnames...)

			for _, domain := range lookupDomains {
				resolveDomainToIPs(domain, ipAddresses, net.LookupIP)
			}
		}
	}

	return ipAddresses, cnames
}

// resolveDomainToIPs resolves a domain name to its IP addresses
func resolveDomainToIPs(domain string, ipAddresses u.StringSet, lookupIPFunc func(string) ([]net.IP, error)) {
	ips, err := lookupIPFunc(domain)
	if err != nil {
		Logger.Warnf("Could not resolve IP for domain '%s': %v", domain, err)
		return
	}

	for _, ip := range ips {
		// Only include IPv4 addresses
		if ip.To4() != nil {
			ipString := ip.String()
			// Skip specific IPs like 0.0.0.0, 127.0.0.1, etc.
			if u.IsSkipIP(Logger, ipString) {
				ipAddresses.Add(ipString)
				Logger.Infof("Domain %s resolved to IP: %s", domain, ipString)
			} else {
				Logger.Debugf("Skipping special IP for domain %s: %s", domain, ipString)
			}
		}
	}
}

// searchInAllFiles searches for the query in all relevant files
func searchInAllFiles(
	query string,
	isIP bool,
	ipAddresses u.StringSet,
	cnames []string,
) (map[string][]string, map[string][]string, map[string][]string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	domainResults := make(map[string][]string)
	ipResults := make(map[string][]string)
	cnameResults := make(map[string][]string)

	// Search in processed files
	if searchProcessed {
		wg.Add(1)
		go func() {
			defer wg.Done()
			searchInFileType(query, isIP, ipAddresses, cnames, constants.SearchProcessedFile,
				&mu, domainResults, ipResults, cnameResults)
		}()
	}

	// Search in consolidated files
	if searchConsolidated {
		wg.Add(1)
		go func() {
			defer wg.Done()
			searchInFileType(query, isIP, ipAddresses, cnames, constants.SearchConsolidatedFile,
				&mu, domainResults, ipResults, cnameResults)
		}()
	}

	wg.Wait()
	return domainResults, ipResults, cnameResults
}

// searchInFileType searches for the query in files of a specific type
func searchInFileType(query string, isIP bool, ipAddresses u.StringSet, cnames []string,
	fileType string, mu *sync.Mutex, domainResults, ipResults, cnameResults map[string][]string) {

	// Search for domain if a query is not an IP
	if !isIP {
		results := searchInFiles(query, constants.SourceTypeDomain, fileType, exactMatch)
		mu.Lock()
		mergeSearchResults(domainResults, results)
		mu.Unlock()

		// Search for CNAME records
		for _, cname := range cnames {
			results := searchInFiles(cname, constants.SourceTypeDomain, fileType, exactMatch)
			mu.Lock()
			mergeSearchResults(cnameResults, results)
			mu.Unlock()
		}
	}

	// Search for IP addresses
	for _, ip := range ipAddresses.ToSlice() {
		results := searchInFiles(ip, constants.SourceTypeIpv4, fileType, exactMatch)
		mu.Lock()
		mergeSearchResults(ipResults, results)
		mu.Unlock()
	}
}

// mergeSearchResults merges search results from different sources
func mergeSearchResults(target, source map[string][]string) {
	for k, v := range source {
		target[k] = append(target[k], v...)
	}
}

// displaySearchResults shows the search results in a formatted way
func displaySearchResults(query string, isIP bool, domainResults, ipResults, cnameResults map[string][]string,
	ipAddresses u.StringSet, cnames []string) {

	// Display domain results
	if !isIP {
		displayDomainResults(query, domainResults)
		displayCnameResults(cnames, cnameResults)
	}

	// Display IP results
	displayIpResults(ipAddresses, ipResults)

	// Show a "no matches" message if nothing was found
	if !isIP && len(ipAddresses) == 0 && len(domainResults) == 0 && len(cnameResults) == 0 {
		Logger.Infof("No matches found for '%s'", query)
	}
}

// displayDomainResults shows the results for domain searches
func displayDomainResults(query string, domainResults map[string][]string) {
	if len(domainResults) > 0 {
		Logger.Infof("Domain '%s' found in %d sources:", query, len(domainResults))
		for source, files := range domainResults {
			Logger.Infof("Source: %s", source)
			for _, file := range files {
				Logger.Infof("  - %s", file)
			}
			Logger.Infof("Total files: %d", len(files))
		}
	} else {
		Logger.Infof("Domain '%s' not found in any sources", query)
	}
}

// displayCnameResults shows the results for CNAME searches
func displayCnameResults(cnames []string, cnameResults map[string][]string) {
	if len(cnames) > 0 && len(cnameResults) > 0 {
		for _, cname := range cnames {
			sourceCount := 0
			for range cnameResults {
				sourceCount++
			}
			Logger.Infof("CNAME '%s' found in %d sources:", cname, sourceCount)
			for source, files := range cnameResults {
				Logger.Infof("Source: %s", source)
				for _, file := range files {
					Logger.Infof("  - %s", file)
				}
				Logger.Infof("Total files: %d", len(files))
			}
		}
	}
}

// displayIpResults shows the results for IP searches
func displayIpResults(ipAddresses u.StringSet, ipResults map[string][]string) {
	ipList := ipAddresses.ToSlice()
	if len(ipList) > 0 {
		for _, ip := range ipList {
			// Process results for this IP
			ipSourceResults := processIpSearchResults(ip, ipResults)

			// Display results
			if len(ipSourceResults) > 0 {
				Logger.Infof("IP '%s' found in %d sources:", ip, len(ipSourceResults))
				for source, files := range ipSourceResults {
					Logger.Infof("Source: %s", source)
					for _, file := range files {
						Logger.Infof("  - %s", file)
					}
					Logger.Infof("Total files: %d", len(files))
				}
			} else {
				Logger.Infof("IP '%s' not found in any sources", ip)
			}
		}
	}
}

// processIpSearchResults processes the search results for an IP address
func processIpSearchResults(ip string, ipResults map[string][]string) map[string][]string {
	ipSourceFiles := make(map[string]u.StringSet)

	for source, files := range ipResults {
		// Initialize a StringSet for this source if needed
		if _, exists := ipSourceFiles[source]; !exists {
			ipSourceFiles[source] = u.NewStringSet(nil)
		}

		for _, file := range files {
			found, err := entryContains(ip, file, exactMatch)
			if err != nil {
				Logger.Errorf("Error checking file %s for IP %s: %v", file, ip, err)
				continue
			}
			if found {
				ipSourceFiles[source].Add(file)
			}
		}
	}

	// Convert results back to the expected format
	ipSourceResults := make(map[string][]string)
	for source, fileSet := range ipSourceFiles {
		if len(fileSet) > 0 {
			ipSourceResults[source] = fileSet.ToSlice()
		}
	}

	return ipSourceResults
}

// searchInFiles searches for the query in files of the specified type in the given directory
func searchInFiles(query string, sourceType string, searchFileType string, exactMatch bool) map[string][]string {
	results := make(map[string][]string)

	var allFiles []string
	switch searchFileType {
	case constants.SearchProcessedFile:
		summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["processed"])
		if files, err := getProcessedFiles(sourceType, summaryFile); err == nil {
			allFiles = append(allFiles, files...)
		} else {
			Logger.Errorf("Error getting processed files: %v", err)
		}
	case constants.SearchConsolidatedFile:
		summaryFile := filepath.Join(constants.SummaryDir, constants.DefaultSummaryFiles["consolidated"])
		if files, err := getConsolidatedFiles(sourceType, summaryFile); err == nil {
			allFiles = append(allFiles, files...)
		} else {
			Logger.Errorf("Error getting consolidated files: %v", err)
		}
	}

	// Use a StringSet to track found files more efficiently
	foundFiles := u.NewStringSet(nil)

	for _, file := range allFiles {
		// Search in file
		found, err := entryContains(query, file, exactMatch)
		if err != nil {
			Logger.Errorf("Error searching in file %s: %v", file, err)
			continue
		}

		if found {
			foundFiles.Add(file)
		}
	}

	// Only add results if we found something
	if len(foundFiles) > 0 {
		results[sourceType] = foundFiles.ToSlice()
	}

	return results
}

// getProcessedFiles retrieves the list of processed files from the summary file
func getProcessedFiles(sourceType string, summaryFile string) ([]string, error) {
	return u.GetSummaryFiles(
		Logger,
		sourceType,
		summaryFile,
		func(summary c.ProcessedSummary, sourceType string) []string {
			// Use StringSet to efficiently collect unique files
			fileSet := u.NewStringSet(nil)
			for _, validFile := range summary.ValidFiles {
				if validFile.GenericSourceType == sourceType {
					filePath := validFile.Filepath
					fileSet.Add(filePath)
				}
			}
			return fileSet.ToSlice()
		},
	)
}

// getConsolidatedFiles retrieves the list of consolidated files from the summary file
func getConsolidatedFiles(sourceType string, summaryFile string) ([]string, error) {
	return u.GetSummaryFiles(
		Logger,
		sourceType,
		summaryFile,
		func(summary c.ConsolidatedSummary, sourceType string) []string {
			if summary.Type == sourceType {
				return []string{summary.Filepath}
			}
			return nil
		},
	)
}

// entryContains checks if the file contains the query string
func entryContains(query, filePath string, exactMatch bool) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer u.CloseFile(Logger, file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if u.IsComment(line) {
			continue
		}

		if exactMatch {
			if line == query {
				return true, nil
			}
		} else {
			if strings.Contains(line, query) {
				return true, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func init() {

	searchCmd.Flags().BoolVarP(&exactMatch, "exact", "e", false, "Perform exact match instead of substring match")
	searchCmd.Flags().BoolVarP(&searchProcessed, "processed", "p", true, "Search in processed files")
	searchCmd.Flags().BoolVarP(&searchConsolidated, "consolidated", "c", true, "Search in consolidated files")
	searchCmd.Flags().
		BoolVarP(&performDNSLookup, "dns", "d", false, "Perform DNS lookup for domain names to find associated IPs")
	searchCmd.Flags().
		BoolVarP(&performCNAMELookup, "cname", "n", true, "Perform CNAME lookup for domain names and search for CNAME records")
	searchCmd.Flags().
		BoolVarP(&searchAguard, "adguard", "g", false, "Search in AdGuard files")
}
