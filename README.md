# DNS Toolkit

[![Go CI Workflow](https://github.com/phani-kb/dns-toolkit/actions/workflows/ci.yml/badge.svg)](https://github.com/phani-kb/dns-toolkit/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/phani-kb/dns-toolkit)](https://goreportcard.com/report/github.com/phani-kb/dns-toolkit)
[![codecov](https://codecov.io/gh/phani-kb/dns-toolkit/branch/main/graph/badge.svg)](https://codecov.io/gh/phani-kb/dns-toolkit)
[![GoDoc](https://godoc.org/github.com/phani-kb/dns-toolkit?status.svg)](https://godoc.org/github.com/phani-kb/dns-toolkit)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/doc/go1.23)

A Go-based command-line utility for downloading, processing, resolving, and consolidating DNS blocklists and allowlists from multiple sources. Performs DNS-to-IP resolution, reverse lookups, and overlap detection. Generates ready-to-use lists for DNS sinkholes, Pi-hole, AdGuard Home, and other network security tools.

<!-- STATS_START -->
## Source Statistics

*Automatically generated statistics from source configuration files*

| Metric | Count | Details |
|--------|-------|---------|
| **Total Sources** | 134 | 91 enabled, 43 disabled |
| **Blocklist Sources** | 126 | Sources providing blocking rules |
| **Allowlist Sources** | 20 | Sources providing exception rules |
| **Categories** | 24 | ads, adult, annoyance, anonymizer, botnet, dns, doh, exploit, fakenews, gambling, malicious, malware, mobile, others, phishing, privacy, proxy, ransomware, scam, security, spam, threat, trackers, trojan |
| **Source Types** | 25 | adguard, cidr_ipv4, domain, domain_adguard, domain_comment, domain_csv_http_url_find, domain_custom_csv_blackbook, domain_custom_csv_maltrail, domain_custom_html_ccam, domain_custom_html_puppyscams, domain_http_url, domain_top, domain_url, domain_with_comment_suffix, hostname, ipv4, ipv4_cidr_expand, ipv4_custom_html_ccam, ipv4_find, ipv4_http_url, ipv4_range_expand, ipv4_url, ipv6, ipv6_find, ipv6_htaccess |
| **Geographic Coverage** | 21 countries | CN, CZ, DE, ES, FI, FR, HU, ID, IL, IT, KR, LV, MY, NL, RO, RU, SA, SK, UA, US, VN |
| **Last Updated** | 2025-07-05 03:08:24 UTC | Statistics generation time |

<!-- STATS_END -->

## Published Outputs

**Ready-to-use blocklist files are published daily to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch:**

- Domain and IP blocklists/allowlists compatible with Pi-hole, pfBlockerNG, AdGuard Home
- Lists organized by size (mini, lite, normal, big) and category (advertising, malware, privacy)
- Top entries based on source frequency for high-confidence blocking

**Usage:** Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your DNS filtering tool.

**Processing summaries and metadata are archived in the [`summaries`](https://github.com/phani-kb/dns-toolkit/tree/summaries) branch with 1-year retention.**

## Installation

```bash
git clone https://github.com/phani-kb/dns-toolkit.git
cd dns-toolkit
go build -o bin/dns-toolkit main.go
```

## Quick Start

```bash
# Download and process all sources
dns-toolkit download
dns-toolkit process
dns-toolkit consolidate
dns-toolkit generate output

# Search for a domain
dns-toolkit search example.com

# Analyze overlaps between sources
dns-toolkit overlap

# Find top entries across sources
dns-toolkit top
```

## Key Commands

```text
DNS Toolkit

Usage:
  dns-toolkit [command]

Available Commands:
  archive          Archive DNS toolkit data
  consolidate      Consolidate processed files
  download         Download enabled sources
  generate         Generate different types of outputs
  help             Help about any command
  overlap          Find overlap between source files
  process          Process downloaded files
  search           Search for a domain or IP in the processed files
  sts              Prints the source types summary
  top              Find top entry(s) in each generic source type
  validate-sources Validate the sources configuration
  version          Print the version number of DNS Toolkit

Flags:
  -h, --help   help for dns-toolkit

Use "dns-toolkit [command] --help" for more information about a command.
```

## Configuration

Sources are configured in `data/config/sources*.json` files. Each source specifies:

- Download URL and frequency
- Source type (domain, IPv4, IPv6, AdGuard, etc.)
- Categories (advertising, malware, privacy, etc.)
- License and website information

## Output Structure

```text
data/output/
├── *_blocklist.txt            # Blocklists for various source types (adguard, domain, ipv4, ipv6, cidr)
├── *_allowlist.txt            # Allowlists for various source types (adguard, domain, etc.)
├── categories/                # Lists by category (ads, malware, privacy, etc.)
├── groups/                    # Lists by size (mini, lite, normal, big)
├── top/                       # Top entries based on source frequency
├── ignored/                   # Entries filtered by allowlists
└── summaries/                 # Processing metadata and statistics
```

## Credits

This project is made possible by the following blocklist and allowlist sources:

1Hosts (Lite), abpvn_hosts, Adaway, AdBlockID, AdGuard Base filter, AdGuard DNS filter, AlienVault_Reputation, AntiAdBlockFilters, anudeepND_adservers, anudeepND_Allowlist, bigdargon_hostsVN, BinaryDefense_Banlist, Blackhole_Today, BlockListDE_Brute, BlockListDE_Strong, Blocklists UT1 Cryptojacking, Blocklists UT1 Malware, Blocklists UT1 Publicite, Borestad_AbuseIPDB, BruteforceBlocker, Cameleon, CINSScore_BadGuys_Army, CJX Annoyance, CybercrimeTracker_All, CybercrimeTracker_CCAM, CybercrimeTracker_CCPMGate, cyberhost_malware-blocklist, Dan Pollock's List, DandelionSprout-Anti-Malware-List, DoH_IP_blocklists, DoH_IP_list, DShield, Easy Privacy, EasyList, EmergingThreats_CompromisedIPs, ET_fwip, fabriziosalmi_allowlist, fabriziosalmi_blocklists, FabrizioSalmi_DNS, FadeMind_2o7Net, FakeWebshopListHUN, Firehol_abusers_30d, Firehol_BitcoinNodes_1d, Firehol_Botscout_1d, Firehol_GPF_Comics, Firehol_level1, Firehol_level2, Firehol_level3, Firehol_SocksProxy_7d, Firehol_SSLProxies_1d, Frogeye trackers, GetAdmiral Domains Filter List, GlobalAntiScamOrg-blocklist, Greensnow, HaGeZi Amazon Tracker, HaGeZi Apple Tracker, HaGeZi DNS TIF Mini, HaGeZi_DoH, HaGeZi Encrypted DNS Servers, HaGeZi Gambling Only Domains, HaGeZi Microsoft Tracker, HaGeZi Most Abused TLDs, HaGeZi Normal, HaGeZi Pro, HaGeZi's Pro Blocklist, HaGeZi_TIF, HaGeZi Xiaomi Tracker, Hestat_Minerchk, Hostsfile, hufilter, iam-py-test_my-filters-001-antitypo, jarelllama_Scam-Blocklist, kadantiscam, Malicious URL Blocklist (URLHaus), Maltrail_StaticTrails, malware-filter_phishing-filter, MyIP_MS_Blocklist, nrd-14day-mini, OISD Blocklist Big, OpenPhish_Feed, Peter Lowe's Blocklist, phishing_army, Policeman_SimpleDomainsBlocklist, Public_DNS4, PuppyScams, quidsup_notrack-annoyance, quidsup_notrack-malware, quidsup_notrack-tracker, RedDragonWebDesign_block-everything, RPiList_specials-malware, RPiList_specials-phishing, Rutgers_DROP, Sblam_Blocklist, ScriptzTeam_BadIPS, Sentinel_Greylist, ShadowWhisperer_Allowlist, ShadowWhisperer_BlockLists Ads, ShadowWhisperer_BlockLists Adult, ShadowWhisperer_BlockLists Malware, ShadowWhisperer_BlockLists Scam, ShadowWhisperer's Dating List, Spam404, spamhaus_drop, Stamparm_Blackbook, StevenBlack_Adhoc_list, StevenBlack_Fake_Gambling_Porn, T145_allowlist, T145_black-mirror, T145_blocklist, transco, Ukrainian Ad Filter, Ukrainian Annoyance Filter, Ukrainian Privacy Filter, Ukrainian Security Filter, UncheckyAds, URLHaus (Abuse.ch), URLHaus_Text, USOM-Blocklists, Viriback_Dump, VXVault_URLList, WaLLy3K, WindowsSpyBlocker_Hosts_spy, Winhelp2002, YousList, YousList-AdGuard, youtube_GoodbyeAds, Yoyo AdServers, Yoyo Adservers.

**Note:** Detailed information about each source including URLs, licenses, categories, and website links can be found in the configuration files located in `data/config/sources*.json`.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes with tests
4. Submit a pull request

## License

This project is licensed under the terms specified in the LICENSE file.
