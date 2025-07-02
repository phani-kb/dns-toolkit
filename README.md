# DNS Toolkit

`dns-toolkit` is a Go-based command-line utility designed to download, process, consolidate, and analyze DNS blocklists and allowlists from multiple sources. It resolves domain names to IP addresses, performs reverse lookups, detects overlaps, and generates output lists tailored for use with DNS sinkholes or firewalls.

## Published Outputs

The DNS Toolkit automatically processes and publishes ready-to-use blocklist/allowlist files daily via GitHub Actions:

### Generated Output Files

Processed and formatted blocklist/allowlist files are published to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch. These files are ready for direct use with DNS sinkholes and network security tools:

- **Template-formatted domain and IP blocklists/allowlists** - Compatible with Pi-hole, pfBlockerNG, AdGuard Home, Unbound, and other DNS filtering solutions
- **Consolidated lists organized by size groups** (mini, lite, normal, big) - Choose the appropriate size for your needs
- **Category-based lists** (advertising, malware, privacy, etc.) - Target specific threat categories
- **Top entries lists** based on source frequency - High-confidence entries appearing across multiple sources

**Usage Examples:**

- Pi-hole: Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your blocklists
- pfBlockerNG: Import raw URLs from the output branch
- AdGuard Home: Use as custom DNS filters

### Summary Archives

Processing summaries and metadata are archived in the [`summaries`](https://github.com/phani-kb/dns-toolkit/tree/summaries) branch, organized by month (01-12) with 1-year retention. These include:

- Download summaries with source status and timestamps
- Processing summaries with validation results
- Consolidation summaries with entry counts and checksums
- Overlap analysis results
- Top entries analysis data

The automated pipeline runs daily at 14:00 UTC and processes all enabled sources, ensuring fresh and up-to-date blocklist data.

## Features

- **Multi-source Downloads**: Download from multiple DNS blocklist/allowlist sources with retry mechanisms
- **Content Processing**: Extract and validate domains, IPv4/IPv6 addresses, CIDR blocks from various formats
- **Consolidation**: Merge lists by type, size groups, or categories with deduplication
- **Overlap Analysis**: Find overlapping entries between different sources
- **Top Entries Analysis**: Identify most common entries across sources
- **Search Functionality**: Search for specific domains/IPs across processed files with DNS resolution

## Installation

```bash
# Clone the repository
git clone https://github.com/phani-kb/dns-toolkit.git
cd dns-toolkit

# Build the binary
go build -o bin/dns-toolkit main.go

# Or use the build script
./scripts/build.sh
```

## Configuration

The toolkit uses YAML configuration files located in the `configs/` directory. The main configuration file `config.yml` defines:

- Application settings (name, version, description)
- Directory paths for different output folders
- Source filters and processing options

Sources are defined in separate JSON files located in `data/config/sources*.json` that specify:

- Download URLs and frequencies
- Source types (AdGuard, domain, IPv4, IPv6, CIDR, etc.)
- List types (blocklist, allowlist)
- Processing groups and categories

## Commands

### Core Pipeline Commands

#### `download`

Downloads enabled sources from configured URLs with retry mechanisms and concurrent processing.

```bash
dns-toolkit download
```

**Features:**

- Configurable retry logic with exponential backoff
- Concurrent downloads with worker pools
- Certificate verification options
- File integrity checking with checksums
- Download summary generation

#### `process`

Processes downloaded files to extract and validate entries by type.

```bash
dns-toolkit process
```

**Features:**

- Multi-format content extraction (domains, IPs, CIDR blocks)
- Custom processors for complex formats (AdGuard, CSV, HTML)
- Validation using regex patterns
- Separation of valid/invalid entries

#### `consolidate`

Consolidates processed files by type with allowlist filtering.

```bash
# Consolidate all types
dns-toolkit consolidate [all]

# Consolidate by size groups (mini, lite, normal, big)
dns-toolkit consolidate groups

# Consolidate by categories (advertising, malware, privacy, etc.)
dns-toolkit consolidate categories
```

**Options:**

- `--ignore-allowlist`: Skip allowlist filtering during consolidation
- `--include-invalid`: Include invalid entries in consolidation
- `--calculate-checksum`: Calculate checksums for consolidated files

**Features:**

- Allowlist-based filtering to remove false positives
- Deduplication across sources
- Size-based grouping (mini/lite/normal/big)
- Category-based grouping (advertising/malware/privacy/etc.)
- Ignored entries tracking

### Analysis Commands

#### `overlap`

Analyzes overlapping entries between different sources with performance profiling.

```bash
dns-toolkit overlap [--cpu-profile] [--mem-profile] [--profile-dir /path/to/profiles]
```

**Profiling Options:**

- `--cpu-profile`: Enable CPU profiling
- `--mem-profile`: Enable memory profiling  
- `--goroutine-profile`: Enable goroutine profiling
- `--block-profile`: Enable block profiling
- `--profile-dir`: Directory to store profile files

**Features:**

- Concurrent overlap analysis with worker pools
- Memory-efficient processing for large datasets
- Performance profiling and analysis
- Compact summary generation

#### `top`

Identifies top entries appearing across multiple sources.

```bash
dns-toolkit top [--min-sources 3] [--max-entries 1000] [--cpu-profile]
```

**Options:**

- `--min-sources, -m`: Minimum sources required (3-12, default: processes multiple ranges)
- `--max-entries, -x`: Maximum entries to output (default: unlimited)
- Profiling options: `--cpu-profile`, `--mem-profile`, `--goroutine-profile`, `--block-profile`, `--profile-dir`

**Features:**

- Configurable source count thresholds
- Concurrent processing with worker pools
- Performance profiling capabilities
- Multiple minimum source ranges processing

#### `search`

Searches for domains or IP addresses across processed and consolidated files.

```bash
dns-toolkit search [domain-or-ip] [options]
```

**Options:**

- `--exact, -e`: Perform exact match instead of substring match
- `--processed, -p`: Search in processed files (default: true)
- `--consolidated, -c`: Search in consolidated files (default: true)
- `--dns, -d`: Perform DNS lookup for domains to find associated IPs
- `--cname, -n`: Perform CNAME lookup and search for CNAME records (default: true)

**Features:**

- Domain and IP address search
- DNS resolution integration
- CNAME record following

### Output Commands

#### `generate output`

Generates formatted output files with template headers.

```bash
dns-toolkit generate output [--include-ignored] [--delete-folders]
```

**Options:**

- `--include-ignored, -i`: Include ignored files in output
- `--delete-folders, -d`: Delete source folders after generation

**Features:**

- Template-based output generation
- Static and dynamic template support
- Summary file copying and archiving

#### `archive`

Creates compressed archives of processed data with metadata.

```bash
dns-toolkit archive
```

**Features:**

- Compressed (.tgz) archive creation
- Metadata and checksums for all files
- Archive summary generation

### Utility Commands

#### `sts` (Source Type Summary)

Analyzes and reports source type distribution and validation.

```bash
dns-toolkit sts
```

**Features:**

- Source type detection and validation
- Mismatch reporting between expected and actual types
- Statistics on source type distribution

#### `validate-sources`

Validates the configuration files and source definitions.

```bash
dns-toolkit validate-sources
```

**Features:**

- Configuration file validation
- Schema validation

#### `version`

Displays the current version information.

```bash
dns-toolkit version
```

## Typical Workflow

1. **Download Sources**: `dns-toolkit download`
2. **Process Content**: `dns-toolkit process`
3. **Consolidate Lists**: `dns-toolkit consolidate all`
4. **Generate Output**: `dns-toolkit generate output`
5. **Archive Data**: `dns-toolkit archive`

## Advanced Usage

### Performance Profiling

Enable profiling for resource-intensive operations:

```bash
# Profile overlap analysis
dns-toolkit overlap --cpu-profile --mem-profile --profile-dir ./profiles

# Profile top entries analysis  
dns-toolkit top --min-sources 5 --cpu-profile --profile-dir ./profiles
```

### Search Operations

```bash
# Search for a domain with DNS resolution
dns-toolkit search example.com --dns

# Search for an IP with exact matching
dns-toolkit search 1.2.3.4 --exact

# Search only in consolidated files
dns-toolkit search malware.com --processed=false
```

## Output Structure

The toolkit generates organized output in several directories:

- `data/download/`: Raw downloaded files
- `data/processed/`: Extracted and validated entries  
- `data/consolidated/`: Merged and deduplicated lists
- `data/output/`: Template-formatted final outputs
- `data/summary/`: JSON summaries of all operations
- `data/archive/`: Compressed historical data
- `logs/`: Application logs

## Configuration Management

### Main Configuration (`configs/config.yml`)

```yaml
application:
  name: dns-toolkit
  version: 0.1.0
  description: A toolkit for DNS data processing and analysis.

dns_toolkit:
  max_retries: 3
  
  folders:
    download: "data/download"
    processed: "data/processed"
    consolidated: "data/consolidated"
    # ... other folder configurations
    
  files_checksum:
    enabled: true
    algorithm: "sha256"
```

### Source Configuration

Sources are defined in JSON files with structure:

```json
{
  "sources": [
    {
      "name": "example-blocklist",
      "url": "https://example.com/blocklist.txt",
      "categories": "advertising, tracking",
      "disabled": false,
      "license": "GPL-3.0",
      "types": [
        {
          "name": "domain",
          "list_types": [
            {
              "name": "blocklist",
              "groups": "normal, big",
              "must_consider": true
            }
          ]
        }
      ],
      "website": "https://example.com"
    }
  ]
}
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
./scripts/coverage.sh

# Run specific test package
go test ./cmd/...
```

### Building from Source

```bash
# Build for current platform
go build -o bin/dns-toolkit main.go

# Build for multiple platforms
./scripts/build.sh
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes with tests
4. Submit a pull request

## License

This project is licensed under the terms specified in the LICENSE file.
