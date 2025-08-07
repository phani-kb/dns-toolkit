package utils

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/phani-kb/multilog"
)

// FindProjectRoot walks up the directory tree to find the project root (containing go.mod)
// If startDir is empty, it uses the current working directory
func FindProjectRoot(startDir string) (string, error) {
	if startDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		startDir = wd
	}

	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from starting directory: %s", startDir)
		}
		dir = parent
	}
}

type StringSet map[string]bool

func NewStringSet(entries []string) StringSet {
	set := make(StringSet, len(entries))
	for _, entry := range entries {
		if entry = strings.TrimSpace(entry); entry != "" {
			set[entry] = false
		}
	}
	return set
}

func NewStringSetWithCapacity(capacity int) StringSet {
	return make(StringSet, capacity)
}

func (s StringSet) Contains(str string) bool {
	_, found := s[str]
	return found
}

func (s StringSet) MustConsider(str string) bool {
	consider, found := s[str]
	return found && consider
}

func (s StringSet) Get(str string) (bool, bool) {
	consider, found := s[str]
	return consider, found
}

func (s StringSet) Add(str string) {
	s[str] = false
}

// AddWithConsider adds a string to the set with the specified consider flag.
// It returns true if the entry was added (not found in the set).
// If the entry is already in the set, it updates the consider flag.
func (s StringSet) AddWithConsider(str string, consider bool) bool {
	_, found := s[str]
	s[str] = consider
	return !found
}

func (s StringSet) Remove(str string) {
	delete(s, str)
}

// AddAll adds all entries to the set, optionally considering existing entries.
// It returns the number of entries that were already in the set.
func (s StringSet) AddAll(entries []string, consider bool) int {
	alreadyExist := 0
	for _, entry := range entries {
		if !s.AddWithConsider(entry, consider) {
			alreadyExist++
		}
	}
	return alreadyExist
}

func (s StringSet) RemoveAll(entries []string) {
	for _, entry := range entries {
		s.Remove(entry)
	}
}

func (s StringSet) ToSlice() []string {
	result := make([]string, 0, len(s))
	for entry := range s {
		result = append(result, entry)
	}
	return result
}

func (s StringSet) ToSliceSorted() []string {
	result := s.ToSlice()
	slices.Sort(result)
	return result
}

func RemoveDuplicates(entries []string) []string {
	if len(entries) == 0 {
		return entries
	}

	set := NewStringSet(entries)

	return set.ToSlice()
}

// SaveFile saves the content from the reader to the specified destination folder and file name.
// It creates the destination folder if it doesn't exist.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - destFolder: Target directory where the file will be saved
//   - fileName: Name of the file to create
//   - reader: Source of content to be saved
//
// Returns:
//   - The absolute path of the saved file or an empty string on error
//   - An error object if the operation failed, nil on success
func SaveFile(logger *multilog.Logger, destFolder, fileName string, reader io.Reader) (string, error) {
	// Ensure the destination folder exists
	if _, err := os.Stat(destFolder); os.IsNotExist(err) {
		err := os.MkdirAll(destFolder, os.ModePerm)
		if err != nil {
			logger.Errorf("Destination folder creation error: %v (folder: %s)", err, destFolder)
			return "", err
		}
	}

	// Save file
	filePath := filepath.Join(destFolder, fileName)
	if _, err := os.Stat(filePath); err == nil {
		logger.Debugf("File already exists: %s", filePath)
	}

	out, err := os.Create(filePath)
	if err != nil {
		logger.Errorf("Creating file error: %v (file: %s)", err, filePath)
		return "", err
	}
	defer CloseFile(logger, out)

	_, err = io.Copy(out, reader)
	if err != nil {
		logger.Errorf("Saving file error: %v (file: %s)", err, filePath)
		return "", err
	}

	return filePath, nil
}

// CloseFile safely closes the given file and logs an error if it fails.
//
// Parameters:
//   - logger: Logger for recording errors
//   - file: The file handle to close
func CloseFile(logger *multilog.Logger, file *os.File) {
	err := file.Close()
	if err != nil {
		logger.Errorf("Closing file error: %v (file: %s)", err, file.Name())
	}
}

// CalculateChecksum calculates the checksum of the specified file using the specified algorithm.
// Supports MD5 and SHA256 algorithms. If the algorithm is empty, it defaults to MD5.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - filePath: Path to the file to calculate the checksum for
//   - algo: Algorithm to use ("md5" or "sha256")
//
// Returns:
//   - A hex string representation of the checksum or an empty string on error
func CalculateChecksum(logger *multilog.Logger, filePath string, algo string) string {
	if algo == "" {
		algo = constants.DefaultHashAlgorithm
	}
	file, err := os.Open(filePath)
	if err != nil {
		logger.Errorf("Opening file error: %v (file: %s)", err, filePath)
		return ""
	}
	defer CloseFile(logger, file)
	var h hash.Hash
	switch algo {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	default:
		logger.Errorf("Unsupported algorithm: %s", algo)
		return ""
	}

	if _, err := io.Copy(h, file); err != nil {
		logger.Errorf("Calculating checksum error: %v (file: %s)", err, filePath)
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

// CalculateChecksumFromContent calculates the checksum of the provided content using the specified algorithm.
// Supports MD5 and SHA256 algorithms. If the algorithm is empty, it defaults to MD5.
//
// Parameters:
//   - content: The byte slice to calculate the checksum for
//   - algo: Algorithm to use ("md5" or "sha256")
//
// Returns:
//   - A hex string representation of the checksum
func CalculateChecksumFromContent(content []byte, algo string) string {
	if algo == "" {
		algo = constants.DefaultHashAlgorithm
	}

	var h hash.Hash
	switch algo {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	default:
		return ""
	}

	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}

// IsComment determines if a line is a comment or an empty line.
// A line is considered a comment if it's empty or starts with any of the common comment prefixes.
//
// Parameters:
//   - line: The string to check
//
// Returns:
//   - true if the line is a comment or empty, false otherwise
func IsComment(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	if trimmedLine == "" {
		return true
	}

	for _, prefix := range constants.CommentPrefixes {
		if strings.HasPrefix(trimmedLine, prefix) {
			return true
		}
	}
	return false
}

// IsIP checks if a string is a valid IP address (IPv4 or IPv6).
//
// Parameters:
//   - line: The string to check
//
// Returns:
//   - true if the string is a valid IP address, false otherwise
func IsIP(line string) bool {
	return net.ParseIP(line) != nil
}

// IsIPv6 checks if a string is a valid IPv6 address.
//
// Parameters:
//   - line: The string to check
//
// Returns:
//   - true if the string is a valid IPv6 address, false otherwise
func IsIPv6(line string) bool {
	ip := net.ParseIP(line)
	return ip != nil && ip.To4() == nil
}

// IsCIDR checks if a string is a valid CIDR notation address.
//
// Parameters:
//   - line: The string to check
//
// Returns:
//   - true if the string is a valid CIDR address, false otherwise
func IsCIDR(line string) bool {
	ipAddr, ipNet, err := net.ParseCIDR(line)
	if err != nil {
		return false
	}
	return ipAddr != nil && ipNet != nil
}

// GetTimestamp returns the current time formatted according to the application's standard timestamp format.
//
// Returns:
//   - A string representation of the current time
func GetTimestamp() string {
	return time.Now().Format(constants.TimestampFormat)
}

// copyFile copies a file from src to dst
func copyFile(logger *multilog.Logger, src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer CloseFile(logger, sourceFile)

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer CloseFile(logger, destFile)

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func StringInSlice(str string, slice []string) bool {
	return NewStringSet(slice).Contains(str)
}

// FindOverlap finds the overlap between two files and returns the overlapping content.
// It reads both files, ignores comments, and identifies common lines.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - file1: Path to the first file
//   - file2: Path to the second file
//
// Returns:
//   - A slice of strings representing the overlapping content
//   - The number of non-comment lines in file1
//   - The number of non-comment lines in file2
//
// Optimized FindOverlap: uses bufio.Scanner and fast comment detection for performance
func FindOverlap(logger *multilog.Logger, file1, file2 string) ([]string, int, int) {
	set := make(map[string]struct{})
	count1 := 0

	f1, err := os.Open(file1)
	if err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, file1)
		return nil, 0, 0
	}
	defer CloseFile(logger, f1)

	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		if !isCommentFast(line) {
			set[line] = struct{}{}
			count1++
		}
	}
	if scanErr := scanner1.Err(); scanErr != nil {
		logger.Errorf("Reading file error: %v (file: %s)", scanErr, file1)
		return nil, 0, 0
	}

	var overlap []string
	count2 := 0
	f2, err := os.Open(file2)
	if err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, file2)
		return nil, 0, 0
	}
	defer CloseFile(logger, f2)

	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		if !isCommentFast(line) {
			if _, ok := set[line]; ok {
				overlap = append(overlap, line)
			}
			count2++
		}
	}
	if err := scanner2.Err(); err != nil {
		logger.Errorf("Reading file error: %v (file: %s)", err, file2)
		return nil, 0, 0
	}

	return overlap, count1, count2
}

// Precompute fast lookup tables for comment prefixes
var (
	singleCharCommentPrefixes map[byte]struct{}
	twoCharCommentPrefixes    map[string]struct{}
	longCommentPrefixes       []string
	commentPrefixInit         sync.Once
)

func initCommentPrefixTables() {
	singleCharCommentPrefixes = make(map[byte]struct{})
	twoCharCommentPrefixes = make(map[string]struct{})
	for _, prefix := range constants.CommentPrefixes {
		if len(prefix) == 1 {
			singleCharCommentPrefixes[prefix[0]] = struct{}{}
		} else if len(prefix) == 2 {
			twoCharCommentPrefixes[prefix] = struct{}{}
		} else if len(prefix) > 2 {
			longCommentPrefixes = append(longCommentPrefixes, prefix)
		}
	}
}

// isCommentFast checks for all comment prefixes in constants.CommentPrefixes after a fast path
func isCommentFast(line string) bool {
	commentPrefixInit.Do(initCommentPrefixTables)
	// Fast path: skip leading spaces, check for empty or comment prefix
	i := 0
	for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
		i++
	}
	if i == len(line) {
		return true // empty or all whitespace
	}
	ch := line[i:]
	if len(ch) == 0 {
		return true
	}
	// Single-char prefix
	if _, ok := singleCharCommentPrefixes[ch[0]]; ok {
		return true
	}
	// Two-char prefix
	if len(ch) > 1 {
		if _, ok := twoCharCommentPrefixes[ch[:2]]; ok {
			return true
		}
	}
	// Long prefix fallback
	for _, prefix := range longCommentPrefixes {
		if strings.HasPrefix(ch, prefix) {
			return true
		}
	}
	return false
}

// IsDomain checks if a string is a valid domain name.
// It uses the domain regex pattern and also ensures the string is not an IP address.
//
// Parameters:
//   - domain: The string to check
//
// Returns:
//   - true if the string is a valid domain name, false otherwise
func IsDomain(domain string) bool {
	if (constants.SourceTypeRegexMap[constants.SourceTypeDomain].MatchString(domain) && !IsIP(domain)) ||
		isPunycodeEncoded(domain) {
		return true
	}
	return false
}

func isPunycodeEncoded(domain string) bool {
	return strings.HasPrefix(domain, constants.PunycodePrefix)
}

// PickRandomLines reads a file and returns a specified number of random lines from it.
// If maxLines is 0, it returns all the lines from the file (excluding comments).
//
// Parameters:
//   - filePath: Path to the file to read
//   - maxLines: Maximum number of lines to return (0 for all)
//
// Returns:
//   - A slice of strings with the selected lines
//   - An error object if reading fails, nil on success
func PickRandomLines(filePath string, maxLines int) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	filteredLines := slices.DeleteFunc(slices.Clone(lines), IsComment)

	if maxLines == 0 || len(filteredLines) <= maxLines {
		return filteredLines, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedLines := make([]string, maxLines)
	for i := 0; i < maxLines; i++ {
		selectedLines[i] = filteredLines[r.Intn(len(filteredLines))]
	}

	return selectedLines, nil
}

// ReadEntriesFromFile reads entries from a file and returns them as a slice of strings.
// Comments are ignored, and duplicate entries are removed.
//
// Parameters:
//   - filepath: Path to the file to read
//
// Returns:
//   - A slice of strings containing the unique entries
//   - Number of duplicate entries found
//   - An error object if reading fails, nil on success
func ReadEntriesFromFile(logger *multilog.Logger, filepath string) ([]string, int, error) {
	return ReadEntriesFromFileWithPool(logger, filepath, nil)
}

// ReadEntriesFromFileWithPool reads entries from a file using a string intern pool.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - filepath: Path to the file to read
//   - pool: Optional string intern pool for memory optimization (maybe nil)
//
// Returns:
//   - A slice of strings containing the unique entries
//   - Number of duplicate entries found
//   - An error object if reading fails, nil on success
func ReadEntriesFromFileWithPool(logger *multilog.Logger, filepath string, pool *DTEntryPool) ([]string, int, error) {
	// Get file info for pre-allocation optimization
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, 0, err
	}

	// Estimate number of entries based on average line length
	estimatedEntries := CapPreallocEntries(int(fileInfo.Size()/constants.EntryAverageCharLength + 1))

	duplicateCount := 0
	file, err := os.Open(filepath)
	if err != nil {
		return nil, duplicateCount, err
	}
	defer CloseFile(logger, file)

	// Pre-allocate the set with estimated capacity to reduce reallocations
	set := make(StringSet, estimatedEntries)
	scanner := bufio.NewScanner(file)

	const maxScannerBuffer = 4 * 1024 * 1024
	buf := make([]byte, maxScannerBuffer)
	scanner.Buffer(buf, maxScannerBuffer)

	lineCount := 0
	noCommentCount := 0

	// Process the file line by line
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		if !IsComment(line) {
			noCommentCount++

			// Intern the string if a pool is provided
			if pool != nil {
				line = pool.Intern(line)
			}

			if !set.AddWithConsider(line, false) {
				duplicateCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, duplicateCount, err
	}

	// Convert the set to a slice
	result := make([]string, 0, len(set))
	for entry := range set {
		result = append(result, entry)
	}

	return result, duplicateCount, nil
}

// WriteEntriesToFile writes the given entries to the specified file.
// The entries are sorted before writing, and each entry is written on a separate line.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - filepath: Path to the file to write
//   - entries: Slice of strings to write
//
// Returns:
//   - An error object if writing fails, nil on success
func WriteEntriesToFile(logger *multilog.Logger, filepath string, entries []string) error {
	if len(entries) == 0 {
		return nil
	}

	sorted := slices.Clone(entries)
	slices.Sort(sorted)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, file)

	writer := bufio.NewWriter(file)
	for _, entry := range sorted {
		_, err := writer.WriteString(entry + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// GetMapKeys returns the keys of a map as a slice.
// This is a generic function that works with any map that has comparable keys.
//
// Type Parameters:
//   - K: The type of map keys (must be comparable)
//   - V: The type of map values
//
// Parameters:
//   - m: The map to extract keys from
//
// Returns:
//   - A slice containing all the keys in the map
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ExpandIpv4Range expands a range of IPv4 addresses into individual addresses.
// The expected format is "startIP-endIP", e.g., "194.180.49.0-194.180.49.255".
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - ipRange: A string containing the IP range in the format "startIP-endIP"
//
// Returns:
//   - A slice of strings with all individual IP addresses in the range (inclusive)
func ExpandIpv4Range(logger *multilog.Logger, ipRange string) []string {
	ips := make([]string, 0, 256)

	// Split the range into start and end IPs
	parts := strings.Split(ipRange, "-")
	if len(parts) != 2 {
		logger.Errorf("Invalid IP range format: %s", ipRange)
		return ips
	}

	startIP := net.ParseIP(strings.TrimSpace(parts[0]))
	endIP := net.ParseIP(strings.TrimSpace(parts[1]))

	if startIP == nil || endIP == nil {
		logger.Errorf("Invalid IP range format: %s", ipRange)
		return ips
	}

	startIP = startIP.To4()
	endIP = endIP.To4()

	if startIP == nil || endIP == nil {
		logger.Errorf("Invalid IP range format: %s", ipRange)
		return ips
	}

	startNum := uint32(startIP[0])<<24 | uint32(startIP[1])<<16 | uint32(startIP[2])<<8 | uint32(startIP[3])
	endNum := uint32(endIP[0])<<24 | uint32(endIP[1])<<16 | uint32(endIP[2])<<8 | uint32(endIP[3])

	if startNum > endNum {
		logger.Errorf("Invalid IP range format: %s", ipRange)
		return ips
	}

	// Generate all IPs in the range (inclusive)
	for i := startNum; i <= endNum; i++ {
		ip := net.IPv4(
			byte(i>>24),
			byte(i>>16),
			byte(i>>8),
			byte(i),
		)
		ips = append(ips, ip.String())
	}

	return ips
}

// ExpandIpv4Cidr expands a CIDR notation IPv4 address into individual addresses.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - cidr: A string containing an IPv4 address in CIDR notation
//
// Returns:
//   - A slice of strings with all individual IP addresses in the CIDR range
func ExpandIpv4Cidr(logger *multilog.Logger, cidr string) ([]string, error) {
	var ips []string
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		logger.Errorf("Invalid CIDR format: %s. Error: %v", cidr, err)
		return nil, fmt.Errorf("invalid CIDR format: %w", err)
	}
	ip := ipNet.IP.To4()
	if ip == nil {
		logger.Errorf("Invalid IPv4 address in CIDR (or IPv6 given): %s", cidr)
		return nil, fmt.Errorf("invalid IPv4 address or non-IPv4 CIDR: %s", cidr)
	}
	maskSize, bits := ipNet.Mask.Size()
	if bits != 32 {
		logger.Errorf("Not an IPv4 mask in CIDR: %s (bits=%d)", cidr, bits)
		return nil, fmt.Errorf("CIDR is not for IPv4: %s", cidr)
	}
	hostBits := bits - maskSize
	if hostBits > 24 {
		logger.Errorf("CIDR range too large to expand: %s", cidr)
		return nil, fmt.Errorf("CIDR range too large to expand: %s", cidr)
	}
	numIPs := uint64(1) << hostBits // 2^hostBits
	ips = make([]string, 0, numIPs)
	startIPInt := uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
	mask := uint32(0xFFFFFFFF) << uint(hostBits)
	startIPInt = startIPInt & mask
	for i := uint32(0); i < uint32(numIPs); i++ {
		currentIPVal := startIPInt + i
		ipAddr := net.IPv4(
			byte((currentIPVal)>>24), // Extract the first byte
			byte((currentIPVal)>>16), // Extract the second byte
			byte((currentIPVal)>>8),  // Extract the third byte
			byte(currentIPVal&0xFF),  // Extract the fourth byte
		)
		ips = append(ips, ipAddr.String())
	}

	return ips, nil
}

// IsArchive checks if a file is an archive based on its extension.
func IsArchive(filePath string) bool {
	for _, ext := range constants.ArchiveExtensions {
		if strings.HasSuffix(filePath, ext) {
			return true
		}
	}
	return false
}

func GetArchiveExtension(uri string) string {
	for _, ext := range constants.ArchiveExtensions {
		if strings.HasSuffix(uri, ext) {
			return ext
		}
	}
	return ""
}

// ExtractArchive extracts the contents of an archive file (either .tar.gz or .zip) to the specified destination folder.
//
// Parameters:
//   - archivePath: Path to the archive file
//   - destFolder: Destination directory where the contents will be extracted
//
// Returns:
//   - An error object if the extraction fails, nil on success
func ExtractArchive(logger *multilog.Logger, archivePath, destFolder string) error {
	for _, ext := range constants.ArchiveExtensions {
		if strings.HasSuffix(archivePath, ext) {
			switch ext {
			case ".tar.gz":
				return extractTarGz(logger, archivePath, destFolder)
			case ".zip":
				return extractZip(logger, archivePath, destFolder)
			}
		}
	}
	return fmt.Errorf("unsupported archive format: %s", archivePath)
}

// extractTarGz extracts a .tar.gz archive to the specified destination folder.
func extractTarGz(logger *multilog.Logger, archivePath, destFolder string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, file)

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := gzipReader.Close(); err != nil {
			logger.Errorf("Closing gzip reader error: %v", err)
		}
	}()

	tarReader := tar.NewReader(gzipReader)

	// Get the base name of the archive without extension for potential subfolder creation
	baseName := filepath.Base(archivePath)
	baseName = strings.TrimSuffix(baseName, ".tar.gz") // Remove .tar.gz suffix

	for {
		head, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Determine the correct path for extraction
		// Ensure the file path does not contain directory traversal elements
		if strings.Contains(head.Name, "..") {
			return fmt.Errorf("invalid file path in archive: %s (contains '..')", head.Name)
		}
		if err := validateArchiveFilePath(head.Name); err != nil {
			return err
		}

		// First try the default path in the destination folder
		filePath := filepath.Join(destFolder, head.Name)

		if !isWithinDirectory(destFolder, filePath) {
			return fmt.Errorf("invalid file path in archive: %s", head.Name)
		}

		// Create the necessary directories
		if head.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		} else if head.Typeflag == tar.TypeReg {
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}

			// Create and write the file
			outFile, err := os.Create(filePath)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				CloseFile(logger, outFile)
				return err
			}
			CloseFile(logger, outFile)

			// For files with no directory structure in the archive (just filenames),
			// also create a copy in a subdirectory named after the archive base name
			// This helps with archives that contain files directly at the root
			if !strings.Contains(head.Name, "/") && !strings.Contains(head.Name, "\\") {
				// Ensure the base name does not contain directory traversal elements
				if strings.Contains(baseName, "..") {
					return fmt.Errorf("invalid archive base name: %s (contains '..')", baseName)
				}
				if err := validateArchiveFilePath(baseName); err != nil {
					return fmt.Errorf("invalid archive base name: %v", err)
				}

				subDirPath := filepath.Join(destFolder, baseName)

				if !isWithinDirectory(destFolder, subDirPath) {
					return fmt.Errorf("invalid subdirectory path: %s", baseName)
				}

				if err := os.MkdirAll(subDirPath, os.ModePerm); err != nil {
					return err
				}

				subFilePath := filepath.Join(subDirPath, head.Name)

				if !isWithinDirectory(destFolder, subFilePath) {
					return fmt.Errorf("invalid sub file path: %s", head.Name)
				}

				subFile, err := os.Create(subFilePath)
				if err != nil {
					return err
				}

				// Reopen the original file to copy its contents
				srcFile, err := os.Open(filePath)
				if err != nil {
					CloseFile(logger, subFile)
					return err
				}

				if _, err := io.Copy(subFile, srcFile); err != nil {
					CloseFile(logger, subFile)
					CloseFile(logger, srcFile)
					return err
				}

				CloseFile(logger, subFile)
				CloseFile(logger, srcFile)
			}
		}
	}

	return nil
}

// extractZip extracts a .zip archive to the specified destination folder.
func extractZip(logger *multilog.Logger, archivePath, destFolder string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			logger.Errorf("Closing zip reader error: %v", err)
		}
	}()

	for _, f := range r.File {
		if err := extractZipFile(logger, f, destFolder); err != nil {
			return err
		}
	}

	return nil
}

// extractZipFile extracts a single file from the zip archive
func extractZipFile(logger *multilog.Logger, f *zip.File, destFolder string) error {
	if err := validateArchiveFilePath(f.Name); err != nil {
		return err
	}

	// Explicitly check for directory traversal elements in the file name
	if strings.Contains(f.Name, "..") {
		return fmt.Errorf("invalid file path in archive: %s", f.Name)
	}

	filePath := filepath.Join(destFolder, f.Name)

	if !isWithinDirectory(destFolder, filePath) {
		return fmt.Errorf("invalid file path in archive: %s", f.Name)
	}

	if f.FileInfo().IsDir() {
		return os.MkdirAll(filePath, os.ModePerm)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, outFile)

	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer CloseBody(logger, rc)

	_, err = io.Copy(outFile, rc)
	return err
}

// copySourceToTargetInternal is the internal implementation for copying files
func copySourceToTargetInternal(logger *multilog.Logger, target c.DownloadTarget, forceOverwrite bool) error {
	sourceFilepath := filepath.Join(target.SourceFolder, target.SourceFile)
	if _, err := os.Stat(sourceFilepath); os.IsNotExist(err) {
		return fmt.Errorf("source file not found: %s", sourceFilepath)
	}
	if _, err := os.Stat(target.TargetFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(target.TargetFolder, os.ModePerm); err != nil {
			return err
		}
	}
	targetFilepath := filepath.Join(target.TargetFolder, target.TargetFile)

	// Check if target exists and we shouldn't overwrite
	if !forceOverwrite {
		if _, err := os.Stat(targetFilepath); err == nil {
			logger.Debugf("Target file already exists: %s", targetFilepath)
			return nil
		}
	}

	// copy a source file to a target file
	sourceFile, err := os.Open(sourceFilepath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, sourceFile)

	targetFile, err := os.Create(targetFilepath)
	if err != nil {
		return err
	}
	defer CloseFile(logger, targetFile)

	if _, err := io.Copy(targetFile, sourceFile); err != nil {
		return err
	}

	if forceOverwrite {
		logger.Debugf("Force copied %s to %s", sourceFilepath, targetFilepath)
	}
	return nil
}

func CopySourceToTarget(logger *multilog.Logger, target c.DownloadTarget) error {
	return copySourceToTargetInternal(logger, target, false)
}

func ForceCopySourceToTarget(logger *multilog.Logger, target c.DownloadTarget) error {
	return copySourceToTargetInternal(logger, target, true)
}

func CloseBody(logger *multilog.Logger, body io.Closer) {
	if err := body.Close(); err != nil {
		logger.Errorf("Closing body error: %v", err)
	}
}

func IsSkipIP(_ *multilog.Logger, ip string) bool {
	if ip != "0.0.0.0" && !strings.HasPrefix(ip, "127.") &&
		!strings.HasPrefix(ip, "169.254.") && !strings.HasPrefix(ip, "224.") {
		return true
	}
	return false
}

func ShouldDownloadSource(logger *multilog.Logger, summaryFile string, sourceName string) bool {
	summary, err := GetLastSummary[c.DownloadSummary](logger, summaryFile, sourceName)
	if err != nil {
		logger.Errorf("Getting last download summary error: %v", err)
		return false
	}

	lastDownload := summary.LastDownloadTimestamp
	if lastDownload == "" || lastDownload == "0001-01-01T00:00:00Z" {
		return true
	}

	lastDownloadTime, err := time.Parse(constants.TimestampFormat, lastDownload)
	if err != nil {
		logger.Errorf("Parsing last download timestamp error: %v", err)
		return false
	}
	now := time.Now()
	switch summary.Frequency {
	case constants.FrequencyDaily:
		return now.Sub(lastDownloadTime) >= 24*time.Hour
	case constants.FrequencyWeekly:
		return now.Sub(lastDownloadTime) >= 7*24*time.Hour
	case constants.FrequencyMonthly:
		return now.Sub(lastDownloadTime) >= 30*24*time.Hour
	default:
		return now.Sub(lastDownloadTime) >= 24*time.Hour
	}
}

// IsAlphanumericWithUnderscoresAndDashes checks if a string contains only alphanumeric characters,
// underscores, and dashes.
func IsAlphanumericWithUnderscoresAndDashes(s string) bool {
	match, err := regexp.MatchString("^[a-zA-Z0-9_-]+$", s)
	if err != nil {
		return false
	}
	return match
}

// GetFileLastModifiedTime retrieves the last modified time of a file and formats it according to the
// application's standard timestamp format.
func GetFileLastModifiedTime(logger *multilog.Logger, filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		logger.Errorf("Getting file info error: %v", err)
		return "", err
	}
	return fileInfo.ModTime().Format(constants.TimestampFormat), nil
}

func LogMemStats(logger *multilog.Logger, prefix string) {
	logger.Perff(prefix)
}

// CapPreallocEntries limits the estimated entries to avoid excessive or insufficient allocation.
func CapPreallocEntries(estimated int) int {
	switch {
	case estimated > constants.MaxPreallocEntries:
		return constants.MaxPreallocEntries
	case estimated < constants.MinPreallocEntries:
		return constants.MinPreallocEntries
	default:
		return estimated
	}
}

// ParsePercent converts a percentage string to a float64 for comparison
func ParsePercent(percentStr string) float64 {
	// Try to parse the percentage string
	percent, err := strconv.ParseFloat(percentStr, 64)
	if err != nil {
		// Return 0 if parsing fails
		return 0.0
	}
	return percent
}

// validateArchiveFilePath checks if an archive file path is safe from directory traversal attacks
func validateArchiveFilePath(filePath string) error {
	if filepath.IsAbs(filePath) {
		return fmt.Errorf("absolute paths not allowed in archive: %s", filePath)
	}

	cleanPath := filepath.Clean(filePath)

	if strings.HasPrefix(cleanPath, "..") || strings.Contains(cleanPath, "/../") {
		return fmt.Errorf("path traversal detected in archive: %s", filePath)
	}

	return nil
}

// isWithinDirectory checks if the target path is within the base directory to prevent directory traversal attacks
func isWithinDirectory(baseDir, targetPath string) bool {
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return false
	}
	absTargetPath, err := filepath.Abs(targetPath)
	if err != nil {
		return false
	}

	absBaseDir = filepath.Clean(absBaseDir)
	absTargetPath = filepath.Clean(absTargetPath)

	if !strings.HasSuffix(absBaseDir, string(filepath.Separator)) {
		absBaseDir += string(filepath.Separator)
	}

	return strings.HasPrefix(absTargetPath+string(filepath.Separator), absBaseDir)
}

// EnsureDirectoryExists creates a directory if it doesn't exist
func EnsureDirectoryExists(logger *multilog.Logger, dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Infof("Creating directory: %s", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			logger.Errorf("Error creating directory %s: %v", dir, err)
			return err
		}
	}
	return nil
}

// GetUserAgent constructs a User-Agent string for HTTP requests based on application information.
// If the application name or version is not provided, defaults will be used.
//
// Parameters:
//   - logger: Logger for recording operations and errors
//   - appName: The name of the application
//   - appVersion: The version of the application
//   - appDescription: Optional description of the application
//
// Returns:
//   - A formatted User-Agent string
func GetUserAgent(logger *multilog.Logger, appName string, appVersion string, appDescription string) string {
	if appName == "" {
		appName = constants.AppName
	}
	if appVersion == "" {
		appVersion = constants.AppVersion
	}

	osInfo := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	userAgent := fmt.Sprintf("%s/%s (%s)", appName, appVersion, osInfo)

	if appDescription == "" {
		appDescription = constants.AppDescription
	}

	userAgent = fmt.Sprintf("%s; %s", userAgent, appDescription)

	logger.Debugf("User agent: %s", userAgent)
	return userAgent
}
