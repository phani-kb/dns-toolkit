# DNS Toolkit

[![Go CI Workflow](https://github.com/phani-kb/dns-toolkit/actions/workflows/ci.yml/badge.svg)](https://github.com/phani-kb/dns-toolkit/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/phani-kb/dns-toolkit)](https://goreportcard.com/report/github.com/phani-kb/dns-toolkit)
[![codecov](https://codecov.io/gh/phani-kb/dns-toolkit/branch/main/graph/badge.svg)](https://codecov.io/gh/phani-kb/dns-toolkit)
[![GoDoc](https://godoc.org/github.com/phani-kb/dns-toolkit?status.svg)](https://godoc.org/github.com/phani-kb/dns-toolkit)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org/doc/go1.24)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/phani-kb/dns-toolkit)](https://github.com/phani-kb/dns-toolkit/graphs/commit-activity)
[![GitHub repo size](https://img.shields.io/github/repo-size/phani-kb/dns-toolkit)](https://github.com/phani-kb/dns-toolkit)

A Go-based command-line utility for downloading, processing, resolving, and consolidating DNS blocklists and allowlists from multiple sources. Performs DNS-to-IP resolution, reverse lookups, and overlap detection. Generates ready-to-use lists for DNS sinkholes, Pi-hole, AdGuard Home, and other network security tools.

---

## Pipeline Flow

```mermaid
---
config:
  theme: mc
---
flowchart LR
    B["Process<br>via Processors"] --> C["Consolidate<br>Group<br>Categorize<br>Top Entries<br>Overlaps"]
    C --> D["Blocklist/Allowlist Files"] & E["Main README"] & F["Overlap Analysis"] & G["Summary Report"] & H["Statistics Report"] & n2["Format Conversion<br>via Converters"] & n4["Graph Generation<br>via Graphers"]
    D --> I["Archive"]
    E --> I
    F --> I
    H --> I
    G --> I
    n1["Data files"] --> B
    A["Download<br>via Downloaders"] --> n1
    n2 --> n3["Converted Files"]
    n4 --> n5["Graph Images"]
    n3 --> I
    n5 --> I
    B@{ shape: procs}
    C@{ shape: h-cyl}
    D@{ shape: docs}
    n2@{ shape: h-cyl}
    n4@{ shape: paper-tape}
    n1@{ shape: docs}
    n3@{ shape: docs}
    style C stroke-width:4px,stroke-dasharray: 0
    style n2 stroke-width:2px,stroke-dasharray: 2
```

---

## Allowlist Generation Flow

```mermaid
---
config:
  theme: mc
---
flowchart LR
 subgraph horizontal[" "]
    direction LR
        OutputIPv4["allowlist_ipv4.txt"]
  end
    ReadSources["Read Sources"] --> ExtractDomains["Extract Domains from Source URLs"]
    ExtractDomains --> LoadCustom["Load Custom Files"]
    LoadCustom --> Combine["Combine & Deduplicate"]
    Combine --> OutputDomains["allowlist_domains.txt"] & OutputAdGuard["allowlist_adg.txt<br>(@@||domain^ format)"] & ResolveDNS["Resolve to IPv4"]
    ResolveDNS --> OutputIPv4

    OutputIPv4@{ shape: doc}
    ReadSources@{ shape: procs}
    ExtractDomains@{ shape: procs}
    OutputDomains@{ shape: doc}
    OutputAdGuard@{ shape: doc}
    ResolveDNS@{ shape: proc}
    style Combine stroke-width:2px,stroke-dasharray: 2
    style horizontal fill:transparent,stroke:transparent

```

---

## Source Configuration (Important!)

Sources are configured in `data/config/sources*.json` files. Each source specifies:

- Download URL and frequency
- Source type (domain, IPv4, IPv6, AdGuard, etc.)
- Categories (ads, malware, privacy, etc.)
- License and website information

[View and edit source configuration files in `data/config/`](https://github.com/phani-kb/dns-toolkit/tree/main/data/config)

> **To add, modify, or review sources, always refer to the files in `data/config/`.**

## Special Note on Top Domains (tranco-list.eu)

Top domains sourced from the tranco-list.eu list (`domain_top` type) are treated as an allowlist. You can configure the `count_to_consider` value in the relevant config file (`data/config/sources_domain_top.json`) to increase the number of top domains included in the allowlist. This is useful for fine-tuning the strictness or permissiveness of your DNS filtering setup.

## Published Outputs

**Ready-to-use blocklist files are published daily to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch:**

- Domain and IP blocklists/allowlists compatible with Pi-hole, pfBlockerNG, AdGuard Home
- Lists organized by size (mini, lite, normal, big) and category (ads, malware, privacy)
- Top entries based on source frequency for high-confidence blocking

**Usage:** Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your DNS filtering tool.

**[View Detailed Overlap Analysis →](https://github.com/phani-kb/dns-toolkit/blob/output/overlap.md)** - Comprehensive analysis showing how entries are shared across different DNS sources.

**Processing summaries and metadata are archived in the [`summaries`](https://github.com/phani-kb/dns-toolkit/tree/summaries) branch with 1-year retention.**

<!-- BRANCH_SIZES_START -->
## Branch Sizes

**Note:** The repo size badge above only reflects the default branch (`main`).

- **Output branch size:** 294.00 MB
- **Summaries branch size:** 1.07 MB

<!-- BRANCH_SIZES_END -->

---

<!-- STATS_START -->
## Source Statistics

*Automatically generated statistics from source configuration files*

| Metric | Count | Details |
|--------|-------|---------|
| **Total Sources** | 170 | 146 enabled, 24 disabled |
| **Blocklist Sources** | 144 | Sources providing blocking rules |
| **Allowlist Sources** | 38 | Sources providing exception rules |
| **Categories** | 37 | ads, adult, annoyance, anonymizer, botnet, browser, cryptocurrency, dating, discord, dns, doh, fake, fakenews, finance, gambling, issues, kad, local, mac, malicious, malware, mobile, others, phishing, privacy, proxy, ransomware, scam, smarttv, social, spam, threat, topdomains, torrent_trackers, trackers, url_shorteners, windows |
| **Source Types** | 27 | adguard, cidr_ipv4, domain, domain_adguard, domain_comment, domain_csv_http_url_find, domain_custom_csv_blackbook, domain_custom_csv_maltrail, domain_custom_html_ccam, domain_custom_html_puppyscams, domain_http_url, domain_top, domain_url, domain_with_comment_suffix, hostname, ipv4, ipv4_cidr_expand, ipv4_csv_http_url_find, ipv4_custom_html_ccam, ipv4_find, ipv4_from_domain, ipv4_http_url, ipv4_range_expand, ipv4_url, ipv6, ipv6_find, ipv6_htaccess |
| **Geographic Coverage** | 21 countries | CN, CZ, DE, ES, FI, FR, HU, ID, IL, IT, KR, LV, MY, NL, RO, RU, SA, SK, UA, US, VN |
| **Last Updated** | 2025-08-24 00:27:55 UTC | Statistics generation time |

<!-- STATS_END -->

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
<!-- CREDITS_START -->
## Source Credits

This project is made possible by the following blocklist and allowlist sources:

<details>
<summary><strong>📄 sources_domain_al.json</strong> (20 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| AdGuardSDNSFilter_exclusions | ✅ Enabled | others | [AL] | - |
| AdGuardTeam_HttpsExclusions_android | ✅ Enabled | mobile | [AL] | - |
| AdGuardTeam_HttpsExclusions_banks | ✅ Enabled | finance | [AL] | - |
| AdGuardTeam_HttpsExclusions_firefox | ✅ Enabled | browser | [AL] | - |
| AdGuardTeam_HttpsExclusions_issues | ✅ Enabled | issues | [AL] | - |
| AdGuardTeam_HttpsExclusions_mac | ✅ Enabled | mac | [AL] | - |
| AdGuardTeam_HttpsExclusions_sensitive | ✅ Enabled | others | [AL] | - |
| AdGuardTeam_HttpsExclusions_windows | ✅ Enabled | windows | [AL] | - |
| BlahDNS_whitelist | ✅ Enabled | others | [AL] | - |
| China_CDN_Whitelist | ❌ Disabled | others | [AL] | - |
| DandelionSprout_AdGuardHome_Whitelist | ✅ Enabled | others | [AL] | - |
| Dogino_Discord_Official | ✅ Enabled | discord | [AL] | - |
| Freekers_Whitelist | ❌ Disabled | others | [AL] | No update since 2019 |
| Notracking_Hosts_whitelist | ✅ Enabled | others | [AL] | Huge list, use with caution |
| ShadowWhisperer_Allowlist | ✅ Enabled | others | [AL] | - |
| ShadowWhisperer_Whitelist | ✅ Enabled | others | [AL] | - |
| T145_allowlist-domains | ❌ Disabled | others | [AL] | Huge list, use with caution |
| TogoFire_AD_Settings_whitelist | ✅ Enabled | others | [AL] | Huge list, use with caution |
| anudeepND_Allowlist | ❌ Disabled | others | [AL] | Last updated on 2021-12-01. This list is no longer maintained. |
| fabriziosalmi_allowlist | ✅ Enabled | others | [AL] | - |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (99 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| 1Hosts (Lite) | ✅ Enabled | ads, trackers | [BL] | 100% covered by other sources |
| AdBlockID | ✅ Enabled | ads | [AL BL] | - |
| AdGuard Base filter | ✅ Enabled | ads, trackers | [AL BL] | - |
| AdGuard CNAME Mail Trackers | ✅ Enabled | trackers | [BL] | - |
| AdGuard CNAME Trackers | ✅ Enabled | trackers | [BL] | - |
| AdGuard DNS filter | ✅ Enabled | ads, trackers | [AL BL] | - |
| Adaway | ✅ Enabled | ads | [BL] | >99% overlap with StevenBlack Fake Gambling list |
| AntiAdBlockFilters | ✅ Enabled | annoyance | [AL BL] | - |
| Blocklists UT1 Cryptojacking | ✅ Enabled | cryptocurrency | [BL] | - |
| Blocklists UT1 Malware | ✅ Enabled | malware | [BL] | >80% overlap with phishing_army |
| Blocklists UT1 Publicite | ✅ Enabled | ads | [BL] | 100% covered by other sources |
| Blocklists UT1 Shortener | ✅ Enabled | url_shorteners | [BL] | - |
| Boutetnico_URL_Shorteners | ✅ Enabled | url_shorteners | [BL] | - |
| CF Torrent Trackers | ✅ Enabled | torrent_trackers | [BL] | - |
| CJX Annoyance | ✅ Enabled | annoyance | [AL BL] | - |
| Cameleon | ❌ Disabled | ads | [BL] | No update since 2018-03-17 |
| CybercrimeTracker_All | ✅ Enabled | botnet, malicious, malware | [BL] | - |
| CybercrimeTracker_CCAM | ❌ Disabled | botnet, malicious, malware | [BL] | No regular updates |
| CybercrimeTracker_CCPMGate | ✅ Enabled | botnet, malicious, malware | [BL] | - |
| Dan Pollock's List | ✅ Enabled | ads, malware, trackers | [BL] | >95% overlap with StevenBlack Fake Gambling list |
| DandelionSprout-Anti-Malware-List | ✅ Enabled | malware | [AL BL] | - |
| Easy Privacy | ✅ Enabled | privacy, trackers | [AL BL] | - |
| EasyList | ✅ Enabled | ads | [BL] | 100% covered by other sources |
| FadeMind_2o7Net | ❌ Disabled | ads, privacy, trackers | [BL] | No update since 2023-11-30 |
| FakeWebshopListHUN | ✅ Enabled | fake, phishing, scam, threat | [BL] | - |
| Frogeye-firstparty-trackers | ✅ Enabled | trackers | [BL] | - |
| GetAdmiral Domains Filter List | ✅ Enabled | ads, annoyance | [AL BL] | - |
| GlobalAntiScamOrg-blocklist-domains | ✅ Enabled | scam | [BL] | - |
| HaGeZi Amazon Tracker | ✅ Enabled | privacy, trackers | [BL] | >98% overlap with HaGeZi Pro |
| HaGeZi Apple Tracker | ✅ Enabled | privacy, trackers | [BL] | >80% overlap with HaGeZi Pro |
| HaGeZi DNS TIF Mini | ✅ Enabled | malicious, threat | [BL] | 100% covered by other sources |
| HaGeZi Encrypted DNS Servers | ✅ Enabled | doh | [BL] | - |
| HaGeZi Gambling Only Domains | ✅ Enabled | gambling | [BL] | Huge list and gambling-specific focus |
| HaGeZi Microsoft Tracker | ✅ Enabled | privacy, trackers | [BL] | >75% overlap with HaGeZi Pro |
| HaGeZi Most Abused TLDs | ✅ Enabled | spam | [BL] | - |
| HaGeZi Normal | ❌ Disabled | ads, malware, trackers | [BL] | 100% overlap with HaGeZi Pro |
| HaGeZi Pro | ✅ Enabled | ads, malware, phishing, trackers | [BL] | - |
| HaGeZi Xiaomi Tracker | ✅ Enabled | privacy, trackers | [BL] | >95% overlap with HaGeZi Pro |
| Hestat_Minerchk | ❌ Disabled | cryptocurrency | [BL] | No update since 2018 |
| Hostsfile | ❌ Disabled | ads | [BL] | No update since 2018-04-20 |
| Korlabs_UrlShortener | ✅ Enabled | url_shorteners | [BL] | - |
| Malicious URL Blocklist (URLHaus) | ✅ Enabled | ads | [BL] | 100% covered by other sources |
| Maltrail_StaticTrails | ✅ Enabled | malware, threat | [BL] | - |
| OISD Blocklist Big | ✅ Enabled | ads, cryptocurrency, malware, phishing, ransomware, trackers | [BL] | Huge list |
| OISD Blocklist NSFW Small | ✅ Enabled | adult | [BL] | - |
| OISD Blocklist Small | ✅ Enabled | ads, cryptocurrency, malware, phishing, ransomware, trackers | [BL] | - |
| OpenPhish_Feed | ✅ Enabled | phishing | [BL] | - |
| Peter Lowe's Blocklist | ✅ Enabled | ads | [BL] | 100% covered by other sources |
| Policeman_SimpleDomainsBlocklist | ❌ Disabled | malicious | [BL] | Archived on 2021-12-26 |
| PuppyScams | ✅ Enabled | fake, scam | [BL] | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| RPiList_specials-malware | ✅ Enabled | malware | [BL] | Huge list |
| RPiList_specials-phishing | ✅ Enabled | phishing | [BL] | Huge list |
| RedDragonWebDesign_block-everything | ✅ Enabled | ads, malicious, trackers | [AL BL] | - |
| ShadowWhisperer's Dating List | ✅ Enabled | dating | [BL] | - |
| ShadowWhisperer_BlockLists Ads | ✅ Enabled | ads | [BL] | - |
| ShadowWhisperer_BlockLists Adult | ✅ Enabled | adult | [BL] | Huge list and adult-specific focus |
| ShadowWhisperer_BlockLists Malware | ✅ Enabled | malware | [BL] | - |
| ShadowWhisperer_BlockLists Scam | ✅ Enabled | scam | [BL] | - |
| ShadowWhisperer_UrlShortener | ✅ Enabled | url_shorteners | [BL] | - |
| Sinfonietta_Adult | ✅ Enabled | adult | [BL] | - |
| Sinfonietta_Gambling | ✅ Enabled | gambling | [BL] | - |
| Sinfonietta_Social | ✅ Enabled | social | [BL] | - |
| Spam404 | ✅ Enabled | spam | [BL] | - |
| Stamparm_Blackbook | ✅ Enabled | malicious, threat | [BL] | >95% overlap with Blocklists UT1 Malware |
| StevenBlack_Adhoc_list | ❌ Disabled | ads, malware, trackers | [BL] | 100% overlap with StevenBlack Fake Gambling list |
| StevenBlack_Fake_Gambling_Porn | ✅ Enabled | ads, adult, fake, fakenews, gambling | [BL] | - |
| T145_black-mirror | ❌ Disabled | malicious, threat | [BL] | Huge list, >8 million entries |
| Torrent Trackers | ✅ Enabled | torrent_trackers | [BL] | - |
| URLHaus (Abuse.ch) | ✅ Enabled | malware | [BL] | - |
| USOM-Blocklists-domains | ✅ Enabled | malicious, threat | [BL] | Huge list |
| Ukrainian Ad Filter | ✅ Enabled | ads | [BL] | - |
| Ukrainian Annoyance Filter | ✅ Enabled | annoyance | [BL] | - |
| Ukrainian Privacy Filter | ✅ Enabled | privacy, trackers | [AL BL] | - |
| Ukrainian Security Filter | ✅ Enabled | malicious, threat | [BL] | - |
| UncheckyAds | ❌ Disabled | ads, privacy, trackers | [BL] | No update since 2021 |
| Viriback_Dump | ✅ Enabled | malware | [BL] | - |
| WaLLy3K | ✅ Enabled | ads | [BL] | - |
| WindowsSpyBlocker_Hosts_spy | ❌ Disabled | privacy, trackers | [BL] | No update since 2022-05-16 |
| Winhelp2002 | ❌ Disabled | ads | [BL] | No update since 2021-03-06 |
| YousList | ✅ Enabled | ads | [BL] | - |
| YousList-AdGuard | ✅ Enabled | ads | [AL BL] | - |
| Yoyo Adservers-Hosts | ✅ Enabled | ads | [BL] | >95% overlap with StevenBlack Fake Gambling list |
| abpvn_hosts | ✅ Enabled | ads | [AL BL] | - |
| anudeepND_adservers | ❌ Disabled | ads | [BL] | No update since 2023-01-16 |
| bigdargon_hostsVN | ✅ Enabled | ads | [BL] | - |
| cyberhost_malware-blocklist | ✅ Enabled | malware | [BL] | - |
| fabriziosalmi_blocklists | ❌ Disabled | malicious, threat | [BL] | Huge list, >3 million entries |
| hkamran80_smarttv | ✅ Enabled | smarttv | [BL] | - |
| hufilter | ✅ Enabled | ads | [BL] | >90% overlap with HaGeZi Pro |
| iam-py-test_my-filters-001-antitypo | ✅ Enabled | fake | [BL] | - |
| jarelllama_Scam-Blocklist | ✅ Enabled | scam | [BL] | Disabled due to very large size (457K entries) - scam-specific focus |
| kadantiscam | ✅ Enabled | kad | [BL] | peer-to-peer network protocol |
| malware-filter_phishing-filter | ✅ Enabled | malware, phishing | [BL] | - |
| pexcn Torrent Trackers | ✅ Enabled | torrent_trackers | [BL] | - |
| phishing_army | ✅ Enabled | phishing | [BL] | - |
| quidsup_notrack-annoyance | ✅ Enabled | annoyance | [BL] | >90% overlap with HaGeZi Pro |
| quidsup_notrack-malware | ✅ Enabled | malware | [BL] | - |
| quidsup_notrack-tracker | ✅ Enabled | trackers | [BL] | - |
| youtube_GoodbyeAds | ✅ Enabled | ads | [BL] | - |

</details>

<details>
<summary><strong>📄 sources_domain_new.json</strong> (1 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| nrd-14day-mini | ❌ Disabled | others | [BL] | Huge list with low unique contribution |

</details>

<details>
<summary><strong>📄 sources_domain_top.json</strong> (1 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| tranco | ✅ Enabled | topdomains | [AL] | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| AlienVault_Reputation | ❌ Disabled | malicious, threat | [BL] | Not available anymore. The service has been discontinued. |
| BinaryDefense_Banlist | ✅ Enabled | malicious, threat | [BL] | This is for public use only. |
| Blackhole_Today | ❌ Disabled | malicious, threat | [BL] | Download fails frequently due to network instability or potential blocking. |
| BlockListDE_Brute | ✅ Enabled | threat | [BL] | >95% overlap with Firehol_level2 |
| BlockListDE_Strong | ✅ Enabled | malicious, threat | [BL] | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| Borestad_AbuseIPDB_S100_3d | ✅ Enabled | malicious, threat | [BL] | - |
| BruteforceBlocker | ✅ Enabled | threat | [BL] | >95% overlap with EmergingThreats_CompromisedIPs |
| CINSScore_BadGuys_Army | ✅ Enabled | malicious, threat | [BL] | - |
| DShield | ✅ Enabled | malicious, threat | [BL] | 100% overlap with Firehol_level2/Firehol_level3 |
| DoH_IP_blocklists | ✅ Enabled | doh | [BL] | >90% overlap with HaGeZi Encrypted DNS Servers |
| DoH_IP_list | ✅ Enabled | doh | [BL] | - |
| ET_fwip | ✅ Enabled | malicious, threat | [BL] | - |
| EmergingThreats_CompromisedIPs | ✅ Enabled | malicious, threat | [BL] | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| FabrizioSalmi_DNS | ✅ Enabled | dns | [BL] | - |
| Firehol_BitcoinNodes_1d | ✅ Enabled | cryptocurrency | [BL] | - |
| Firehol_Botscout_1d | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_CleanTalk | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_CleanTalk_Top20 | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_GPF_Comics | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_SSLProxies_1d | ✅ Enabled | anonymizer, privacy, proxy | [BL] | - |
| Firehol_SocksProxy_7d | ✅ Enabled | anonymizer, privacy, proxy | [BL] | - |
| Firehol_abusers_30d | ❌ Disabled | malicious, threat | [BL] | False positives are common, use with caution. |
| Firehol_level1 | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_level2 | ✅ Enabled | malicious, threat | [BL] | - |
| Firehol_level3 | ✅ Enabled | malicious, threat | [BL] | - |
| GlobalAntiScamOrg-blocklist-ips | ✅ Enabled | scam | [BL] | - |
| Greensnow | ✅ Enabled | malicious, malware, threat | [BL] | >95% overlap with Firehol_level2 |
| HaGeZi_DoH | ✅ Enabled | doh | [BL] | >90% overlap with DoH_IP_blocklists |
| HaGeZi_TIF | ✅ Enabled | malicious, threat | [BL] | No unique contribution |
| MyIP_MS_Blocklist | ✅ Enabled | malicious, threat | [BL] | - |
| Public_DNS4 | ✅ Enabled | dns | [BL] | - |
| Rutgers_DROP | ✅ Enabled | malicious, threat | [BL] | - |
| Sblam_Blocklist | ✅ Enabled | spam | [BL] | - |
| ScriptzTeam_BadIPS | ✅ Enabled | malicious, threat | [BL] | - |
| Sentinel_Greylist | ✅ Enabled | malicious, threat | [BL] | - |
| T145_allowlist-ips | ❌ Disabled | others | [AL] | Huge list, use with caution. More than its blocklist counterpart. |
| T145_blocklist | ❌ Disabled | malicious, malware, threat | [BL] | Huge list, use with caution. |
| URLHaus_Text | ✅ Enabled | malware | [BL] | - |
| USOM-Blocklists-ips | ✅ Enabled | malicious, threat | [BL] | - |
| Yoyo AdServers-IPList | ✅ Enabled | ads | [BL] | - |
| spamhaus_drop | ✅ Enabled | spam, threat | [BL] | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (7 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| Local AI Allowlist (Domain) | ✅ Enabled | local | [AL] | - |
| Local AI Blocklist (Domain) | ✅ Enabled | local | [BL] | - |
| Local Allowlist (AdGuard) | ✅ Enabled | local | [AL] | - |
| Local Allowlist (Domain) | ✅ Enabled | local | [AL] | - |
| Local Allowlist (ipv4) | ✅ Enabled | local | [AL] | - |
| Local Blocklist (AdGuard) | ✅ Enabled | local | [BL] | - |
| Local Blocklist (Domain) | ✅ Enabled | local | [BL] | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | Status | Categories | AL/BL | Notes |
|------|--------|------------|-------|-------|
| VXVault_URLList | ✅ Enabled | malware | [BL] | >95% overlap with Firehol_level3 |

</details>

<!-- CREDITS_END -->

**Note:** Detailed information about each source can be found in the configuration files located in `data/config/sources*.json`.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes with tests
4. **Validate commit messages** before pushing. All commit message must reference a GitHub issue.
5. Submit a pull request

## Issues

If you encounter a bug, have a feature request, or want to suggest an improvement, please open an issue in the [GitHub Issues](https://github.com/phani-kb/dns-toolkit/issues) page.

## License

This project is licensed under the terms specified in the LICENSE file.
