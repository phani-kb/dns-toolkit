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
- Categories (advertising, malware, privacy, etc.)
- License and website information

[View and edit source configuration files in `data/config/`](https://github.com/phani-kb/dns-toolkit/tree/main/data/config)

> **To add, modify, or review sources, always refer to the files in `data/config/`.**

## Special Note on Top Domains (tranco-list.eu)

Top domains sourced from the tranco-list.eu list (`domain_top` type) are treated as an allowlist. You can configure the `count_to_consider` value in the relevant config file (`data/config/sources_domain_top.json`) to increase the number of top domains included in the allowlist. This is useful for fine-tuning the strictness or permissiveness of your DNS filtering setup.

## Published Outputs

**Ready-to-use blocklist files are published daily to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch:**

- Domain and IP blocklists/allowlists compatible with Pi-hole, pfBlockerNG, AdGuard Home
- Lists organized by size (mini, lite, normal, big) and category (advertising, malware, privacy)
- Top entries based on source frequency for high-confidence blocking

**Usage:** Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your DNS filtering tool.

**[View Detailed Overlap Analysis →](https://github.com/phani-kb/dns-toolkit/blob/output/overlap.md)** - Comprehensive analysis showing how entries are shared across different DNS sources.

**Processing summaries and metadata are archived in the [`summaries`](https://github.com/phani-kb/dns-toolkit/tree/summaries) branch with 1-year retention.**

<!-- BRANCH_SIZES_START -->
## Branch Sizes

**Note:** The repo size badge above only reflects the default branch (`main`).

- **Output branch size:** 290.10 MB
- **Summaries branch size:** 1.02 MB

<!-- BRANCH_SIZES_END -->

---

<!-- STATS_START -->
## Source Statistics

*Automatically generated statistics from source configuration files*

| Metric | Count | Details |
|--------|-------|---------|
| **Total Sources** | 169 | 119 enabled, 50 disabled |
| **Blocklist Sources** | 144 | Sources providing blocking rules |
| **Allowlist Sources** | 37 | Sources providing exception rules |
| **Categories** | 35 | ads, adult, annoyance, anonymizer, botnet, browser, cryptocurrency, discord, dns, doh, fake, fakenews, finance, gambling, issues, local, mac, malicious, malware, mobile, others, phishing, privacy, proxy, ransomware, scam, smarttv, social, spam, threat, topdomains, torrent_trackers, trackers, url_shorteners, windows |
| **Source Types** | 26 | adguard, cidr_ipv4, domain, domain_adguard, domain_comment, domain_csv_http_url_find, domain_custom_csv_blackbook, domain_custom_csv_maltrail, domain_custom_html_ccam, domain_custom_html_puppyscams, domain_http_url, domain_top, domain_url, domain_with_comment_suffix, hostname, ipv4, ipv4_cidr_expand, ipv4_csv_http_url_find, ipv4_custom_html_ccam, ipv4_find, ipv4_http_url, ipv4_range_expand, ipv4_url, ipv6, ipv6_find, ipv6_htaccess |
| **Geographic Coverage** | 21 countries | CN, CZ, DE, ES, FI, FR, HU, ID, IL, IT, KR, LV, MY, NL, RO, RU, SA, SK, UA, US, VN |
| **Last Updated** | 2025-08-23 01:06:05 UTC | Statistics generation time |

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

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [AdGuardSDNSFilter_exclusions](https://raw.githubusercontent.com/AdguardTeam/AdGuardSDNSFilter/master/Filters/exclusions.txt) | ✅ Enabled | others | - |
| [AdGuardTeam_HttpsExclusions_android](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/android.txt) | ✅ Enabled | mobile | - |
| [AdGuardTeam_HttpsExclusions_banks](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/banks.txt) | ✅ Enabled | finance | - |
| [AdGuardTeam_HttpsExclusions_firefox](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/firefox.txt) | ✅ Enabled | browser | - |
| [AdGuardTeam_HttpsExclusions_issues](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/issues.txt) | ✅ Enabled | issues | - |
| [AdGuardTeam_HttpsExclusions_mac](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/mac.txt) | ✅ Enabled | mac | - |
| [AdGuardTeam_HttpsExclusions_sensitive](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/sensitive.txt) | ✅ Enabled | others | - |
| [AdGuardTeam_HttpsExclusions_windows](https://raw.githubusercontent.com/AdguardTeam/HttpsExclusions/master/exclusions/windows.txt) | ✅ Enabled | windows | - |
| [BlahDNS_whitelist](https://raw.githubusercontent.com/ookangzheng/blahdns/master/hosts/whitelist.txt) | ✅ Enabled | others | - |
| [China_CDN_Whitelist](https://raw.githubusercontent.com/mawenjian/china-cdn-domain-whitelist/master/china-cdn-domain-whitelist.txt) | ❌ Disabled | others | - |
| [DandelionSprout_AdGuardHome_Whitelist](https://raw.githubusercontent.com/DandelionSprout/AdGuard-Home-Whitelist/master/whitelist.txt) | ✅ Enabled | others | - |
| [Dogino_Discord_Official](https://raw.githubusercontent.com/Dogino/Discord-Phishing-URLs/main/official-domains.txt) | ✅ Enabled | discord | - |
| [Freekers_Whitelist](https://raw.githubusercontent.com/freekers/whitelist/master/domains/whitelist.txt) | ❌ Disabled | others | No update since 2019 |
| [Notracking_Hosts_whitelist](https://raw.githubusercontent.com/notracking/hosts-blocklists-scripts/master/hostnames.whitelist.txt) | ❌ Disabled | others | Huge list use with caution |
| [ShadowWhisperer_Allowlist](https://raw.githubusercontent.com/ShadowWhisperer/BlockLists/refs/heads/master/Whitelists/Whitelist) | ✅ Enabled | others | - |
| [ShadowWhisperer_Whitelist](https://raw.githubusercontent.com/ShadowWhisperer/BlockLists/master/Whitelists/Whitelist) | ✅ Enabled | others | - |
| [T145_allowlist-domains](https://github.com/T145/black-mirror/releases/download/latest/ALLOW_DOMAIN.txt) | ❌ Disabled | others | Too many allowlist domains, use with caution. |
| [TogoFire_AD_Settings_whitelist](https://raw.githubusercontent.com/TogoFire-Home/AD-Settings/main/Filters/whitelist.txt) | ❌ Disabled | others | Huge list use with caution |
| [anudeepND_Allowlist](https://raw.githubusercontent.com/anudeepND/whitelist/refs/heads/master/domains/whitelist.txt) | ❌ Disabled | others | Last updated on 2021-12-01. This list is no longer maintained. |
| [fabriziosalmi_allowlist](https://raw.githubusercontent.com/fabriziosalmi/blacklists/refs/heads/main/whitelist.txt) | ✅ Enabled | others | - |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (100 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [1Hosts (Lite)](https://o0.pages.dev/Lite/domains.txt) | ❌ Disabled | ads, trackers | Covered by other sources |
| [AdBlockID](https://cdn.jsdelivr.net/gh/realodix/AdBlockID@master/dist/adblockid.adfl.txt) | ✅ Enabled | ads | - |
| [AdGuard Base filter](https://filters.adtidy.org/extension/firefox/filters/2.txt) | ✅ Enabled | ads, trackers | - |
| [AdGuard CNAME Mail Trackers](https://raw.githubusercontent.com/AdguardTeam/cname-trackers/master/data/combined_disguised_mail_trackers_justdomains.txt) | ✅ Enabled | trackers | - |
| [AdGuard CNAME Trackers](https://raw.githubusercontent.com/AdguardTeam/cname-trackers/master/data/combined_disguised_trackers_justdomains.txt) | ✅ Enabled | trackers | - |
| [AdGuard DNS filter](https://adguardteam.github.io/HostlistsRegistry/assets/filter_1.txt) | ✅ Enabled | ads, trackers | - |
| [Adaway](https://adaway.org/hosts.txt) | ❌ Disabled | ads | >99% overlap with StevenBlack Fake Gambling list |
| [AntiAdBlockFilters](https://easylist-downloads.adblockplus.org/antiadblockfilters.txt) | ✅ Enabled | annoyance | - |
| [Blocklists UT1 Cryptojacking](https://github.com/olbat/ut1-blacklists/raw/refs/heads/master/blacklists/cryptojacking/domains) | ✅ Enabled | cryptocurrency | - |
| [Blocklists UT1 Malware](https://github.com/olbat/ut1-blacklists/raw/refs/heads/master/blacklists/malware/domains) | ✅ Enabled | malware | >80% overlap with phishing_army |
| [Blocklists UT1 Publicite](https://dsi.ut-capitole.fr/blacklists/download/publicite.tar.gz) | ❌ Disabled | ads | Covered by other sources |
| [Blocklists UT1 Shortener](https://raw.githubusercontent.com/olbat/ut1-blacklists/refs/heads/master/blacklists/shortener/domains) | ✅ Enabled | url_shorteners | - |
| [Boutetnico_URL_Shorteners](https://raw.githubusercontent.com/boutetnico/url-shorteners/master/list.txt) | ✅ Enabled | url_shorteners | - |
| [CF Torrent Trackers](https://cf.trackerslist.com/all.txt) | ✅ Enabled | torrent_trackers | - |
| [CJX Annoyance](https://raw.githubusercontent.com/cjx82630/cjxlist/refs/heads/master/cjx-annoyance.txt) | ✅ Enabled | annoyance | - |
| [Cameleon](https://sysctl.org/cameleon/hosts) | ❌ Disabled | ads | No update since 2018-03-17 |
| [CybercrimeTracker_All](https://cybercrime-tracker.net/all.php) | ✅ Enabled | botnet, malicious, malware | - |
| [CybercrimeTracker_CCAM](https://cybercrime-tracker.net/ccam.php) | ❌ Disabled | botnet, malicious, malware | No regular updates |
| [CybercrimeTracker_CCPMGate](https://cybercrime-tracker.net/ccpmgate.php) | ✅ Enabled | botnet, malicious, malware | - |
| [Dan Pollock's List](https://someonewhocares.org/hosts/hosts) | ❌ Disabled | ads, malware, trackers | >95% overlap with StevenBlack Fake Gambling list |
| [DandelionSprout-Anti-Malware-List](https://raw.githubusercontent.com/DandelionSprout/adfilt/refs/heads/master/Dandelion%20Sprout's%20Anti-Malware%20List.txt) | ✅ Enabled | malware | - |
| [Easy Privacy](https://easylist.to/easylist/easyprivacy.txt) | ✅ Enabled | privacy, trackers | - |
| [EasyList](https://easylist.to/easylist/easylist.txt) | ❌ Disabled | ads | Covered by other sources |
| [FadeMind_2o7Net](https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.2o7Net/hosts) | ❌ Disabled | ads, privacy, trackers | No update since 2023-11-30 |
| [FakeWebshopListHUN](https://raw.githubusercontent.com/FakesiteListHUN/FakeWebshopListHUN/refs/heads/main/fakewebshoplist) | ✅ Enabled | fake, phishing, scam, threat | - |
| [Frogeye-firstparty-trackers](https://hostfiles.frogeye.fr/firstparty-trackers-hosts.txt) | ✅ Enabled | trackers | - |
| [GetAdmiral Domains Filter List](https://raw.githubusercontent.com/LanikSJ/ubo-filters/refs/heads/main/filters/getadmiral-domains.txt) | ✅ Enabled | ads, annoyance | - |
| [GlobalAntiScamOrg-blocklist-domains](https://raw.githubusercontent.com/elliotwutingfeng/GlobalAntiScamOrg-blocklist/refs/heads/main/global-anti-scam-org-scam-urls-pihole.txt) | ✅ Enabled | scam | To be reviewed |
| [HaGeZi Amazon Tracker](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/domains/native.amazon.txt) | ❌ Disabled | privacy, trackers | >98% overlap with HaGeZi Pro |
| [HaGeZi Apple Tracker](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/domains/native.apple.txt) | ❌ Disabled | privacy, trackers | >80% overlap with HaGeZi Pro |
| [HaGeZi DNS TIF Mini](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/adblock/tif.mini.txt) | ❌ Disabled | malicious, threat | Covered by other sources |
| [HaGeZi Encrypted DNS Servers](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/adblock/doh.txt) | ✅ Enabled | doh | - |
| [HaGeZi Gambling Only Domains](https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/wildcard/gambling-onlydomains.txt) | ✅ Enabled | gambling | Skipped due to large size (202K entries) and gambling-specific focus |
| [HaGeZi Microsoft Tracker](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/domains/native.winoffice.txt) | ❌ Disabled | privacy, trackers | >75% overlap with HaGeZi Pro |
| [HaGeZi Most Abused TLDs](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/adblock/spam-tlds-ublock.txt) | ✅ Enabled | spam | - |
| [HaGeZi Normal](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/hosts/multi.txt) | ❌ Disabled | ads, malware, trackers | 100% overlap with HaGeZi Pro |
| [HaGeZi Pro](https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/domains/pro.txt) | ✅ Enabled | ads, malware, phishing, trackers | - |
| [HaGeZi Xiaomi Tracker](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/domains/native.xiaomi.txt) | ❌ Disabled | privacy, trackers | >95% overlap with HaGeZi Pro |
| [HaGeZi's Pro Blocklist](https://adguardteam.github.io/HostlistsRegistry/assets/filter_48.txt) | ❌ Disabled | ads | 100% overlap with HaGeZi Pro |
| [Hestat_Minerchk](https://raw.githubusercontent.com/Hestat/minerchk/master/hostslist.txt) | ❌ Disabled | cryptocurrency | No update since 2018 |
| [Hostsfile](https://www.hostsfile.org/Downloads/hosts.txt) | ❌ Disabled | ads | No update since 2018-04-20 |
| [Korlabs_UrlShortener](https://raw.githubusercontent.com/korlabsio/urlshortener/refs/heads/main/names.txt) | ✅ Enabled | url_shorteners | - |
| [Malicious URL Blocklist (URLHaus)](https://adguardteam.github.io/HostlistsRegistry/assets/filter_11.txt) | ❌ Disabled | ads | Covered by other sources |
| [Maltrail_StaticTrails](https://raw.githubusercontent.com/stamparm/aux/master/maltrail-static-trails.txt) | ✅ Enabled | malware, threat | - |
| [OISD Blocklist Big](https://big.oisd.nl/) | ❌ Disabled | ads, malware, trackers | - |
| [OISD Blocklist NSFW Small](https://nsfw-small.oisd.nl) | ✅ Enabled | adult | - |
| [OISD Blocklist Small](https://small.oisd.nl/) | ✅ Enabled | ads, cryptocurrency, malware, phishing, ransomware, trackers | - |
| [OpenPhish_Feed](https://openphish.com/feed.txt) | ✅ Enabled | phishing | - |
| [Peter Lowe's Blocklist](https://pgl.yoyo.org/adservers/serverlist.php?hostformat=nohtml) | ❌ Disabled | ads | Covered by other sources |
| [Policeman_SimpleDomainsBlocklist](https://raw.githubusercontent.com/futpib/policeman-rulesets/master/examples/simple_domains_blacklist.txt) | ❌ Disabled | malicious | Archived on 2021-12-26 |
| [PuppyScams](https://web.archive.org/web/20250403122633/https://puppyscams.org/top-100-pet-scams/) | ✅ Enabled | fake, scam | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| [RPiList_specials-malware](https://raw.githubusercontent.com/RPiList/specials/master/Blocklisten/malware) | ❌ Disabled | malware | Huge list |
| [RPiList_specials-phishing](https://raw.githubusercontent.com/RPiList/specials/master/Blocklisten/Phishing-Angriffe) | ✅ Enabled | phishing | Huge list |
| [RedDragonWebDesign_block-everything](https://raw.githubusercontent.com/RedDragonWebDesign/block-everything/refs/heads/master/block-everything.txt) | ✅ Enabled | ads, malicious, trackers | - |
| [ShadowWhisperer's Dating List](https://adguardteam.github.io/HostlistsRegistry/assets/filter_57.txt) | ✅ Enabled | ads | - |
| [ShadowWhisperer_BlockLists Ads](https://github.com/ShadowWhisperer/BlockLists/raw/master/Lists/Ads) | ✅ Enabled | ads | - |
| [ShadowWhisperer_BlockLists Adult](https://github.com/ShadowWhisperer/BlockLists/raw/master/Lists/Adult) | ✅ Enabled | adult | Disabled due to large size and adult content focus |
| [ShadowWhisperer_BlockLists Malware](https://github.com/ShadowWhisperer/BlockLists/raw/master/Lists/Malware) | ✅ Enabled | malware | - |
| [ShadowWhisperer_BlockLists Scam](https://github.com/ShadowWhisperer/BlockLists/raw/master/Lists/Scam) | ✅ Enabled | scam | - |
| [ShadowWhisperer_UrlShortener](https://raw.githubusercontent.com/ShadowWhisperer/BlockLists/refs/heads/master/Lists/UrlShortener) | ✅ Enabled | url_shorteners | - |
| [Sinfonietta_Adult](https://raw.githubusercontent.com/Sinfonietta/hostfiles/refs/heads/master/pornography-hosts) | ✅ Enabled | adult | - |
| [Sinfonietta_Gambling](https://raw.githubusercontent.com/Sinfonietta/hostfiles/refs/heads/master/gambling-hosts) | ✅ Enabled | gambling | - |
| [Sinfonietta_Social](https://raw.githubusercontent.com/Sinfonietta/hostfiles/refs/heads/master/social-hosts) | ✅ Enabled | social | - |
| [Spam404](https://raw.githubusercontent.com/Spam404/lists/master/main-blacklist.txt) | ✅ Enabled | spam | - |
| [Stamparm_Blackbook](https://raw.githubusercontent.com/stamparm/blackbook/master/blackbook.csv) | ✅ Enabled | malicious, threat | >95% overlap with Blocklists UT1 Malware |
| [StevenBlack_Adhoc_list](https://raw.githubusercontent.com/StevenBlack/hosts/master/data/StevenBlack/hosts) | ❌ Disabled | ads, malware, trackers | 100% overlap with StevenBlack Fake Gambling list |
| [StevenBlack_Fake_Gambling_Porn](https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-porn/hosts) | ✅ Enabled | ads, adult, fake, fakenews, gambling | - |
| [T145_black-mirror](https://github.com/T145/black-mirror/releases/download/latest/BLOCK_DOMAIN.txt) | ❌ Disabled | malicious, threat | Huge list |
| [Torrent Trackers](https://raw.githubusercontent.com/im-sm/Pi-hole-Torrent-Blocklist/main/all-torrent-trackers.txt) | ✅ Enabled | torrent_trackers | - |
| [URLHaus (Abuse.ch)](https://urlhaus.abuse.ch/downloads/hostfile) | ✅ Enabled | malware | - |
| [USOM-Blocklists-domains](https://raw.githubusercontent.com/elliotwutingfeng/USOM-Blocklists/refs/heads/main/urls_pihole.txt) | ✅ Enabled | malicious, threat | Huge list |
| [Ukrainian Ad Filter](https://raw.githubusercontent.com/ukrainianfilters/lists/refs/heads/main/ads/ads.txt) | ✅ Enabled | ads | - |
| [Ukrainian Annoyance Filter](https://raw.githubusercontent.com/ukrainianfilters/lists/refs/heads/main/annoyances/annoyances.txt) | ✅ Enabled | annoyance | - |
| [Ukrainian Privacy Filter](https://raw.githubusercontent.com/ukrainianfilters/lists/refs/heads/main/privacy/privacy.txt) | ✅ Enabled | privacy, trackers | - |
| [Ukrainian Security Filter](https://raw.githubusercontent.com/braveinnovators/ukrainian-security-filter/main/lists/domains.txt) | ✅ Enabled | malicious, threat | - |
| [UncheckyAds](https://raw.githubusercontent.com/FadeMind/hosts.extras/master/UncheckyAds/hosts) | ❌ Disabled | ads, privacy, trackers | No update since 2021 |
| [Viriback_Dump](https://tracker.viriback.com/dump.php) | ✅ Enabled | malware | - |
| [WaLLy3K](https://v.firebog.net/hosts/static/w3kbl.txt) | ✅ Enabled | ads | - |
| [WindowsSpyBlocker_Hosts_spy](https://raw.githubusercontent.com/crazy-max/WindowsSpyBlocker/master/data/hosts/spy.txt) | ❌ Disabled | privacy, trackers | No update since 2022-05-16 |
| [Winhelp2002](https://winhelp2002.mvps.org/hosts.txt) | ❌ Disabled | ads | No update since 2021-03-06 |
| [YousList](https://raw.githubusercontent.com/yous/YousList/master/hosts.txt) | ✅ Enabled | ads | - |
| [YousList-AdGuard](https://raw.githubusercontent.com/yous/YousList/refs/heads/master/youslist.txt) | ✅ Enabled | ads | - |
| [Yoyo Adservers-Hosts](https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts&showintro=0&mimetype=plaintext) | ❌ Disabled | ads | >95% overlap with StevenBlack Fake Gambling list |
| [abpvn_hosts](https://raw.githubusercontent.com/abpvn/abpvn/refs/heads/master/filter/abpvn.txt) | ✅ Enabled | ads | - |
| [anudeepND_adservers](https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt) | ❌ Disabled | ads | No update since 2023-01-16 |
| [bigdargon_hostsVN](https://raw.githubusercontent.com/bigdargon/hostsVN/refs/heads/master/hosts) | ✅ Enabled | ads | - |
| [cyberhost_malware-blocklist](https://lists.cyberhost.uk/malware.txt) | ✅ Enabled | malware | - |
| [fabriziosalmi_blocklists](https://github.com/fabriziosalmi/blacklists/releases/download/latest/blacklist.txt) | ❌ Disabled | malicious, threat | Huge list |
| [hkamran80_smarttv](https://raw.githubusercontent.com/hkamran80/blocklists/refs/heads/main/smart-tv.txt) | ✅ Enabled | smarttv | - |
| [hufilter](https://cdn.jsdelivr.net/gh/hufilter/hufilter@gh-pages/hufilter-hosts.txt) | ✅ Enabled | ads | >90% overlap with HaGeZi Pro |
| [iam-py-test_my-filters-001-antitypo](https://raw.githubusercontent.com/iam-py-test/my_filters_001/main/antitypo.txt) | ✅ Enabled | fake | - |
| [jarelllama_Scam-Blocklist](https://raw.githubusercontent.com/jarelllama/Scam-Blocklist/main/lists/wildcard_domains/scams.txt) | ✅ Enabled | scam | Disabled due to very large size (457K entries) - scam-specific focus |
| [kadantiscam](https://raw.githubusercontent.com/FiltersHeroes/KADhosts/master/KADomains.txt) | ❌ Disabled | ads | >90% overlap with phishing_army |
| [malware-filter_phishing-filter](https://malware-filter.gitlab.io/malware-filter/phishing-filter-hosts.txt) | ✅ Enabled | malware, phishing | - |
| [pexcn Torrent Trackers](https://raw.githubusercontent.com/pexcn/daily/gh-pages/trackerlist/trackerlist-best.txt) | ✅ Enabled | torrent_trackers | - |
| [phishing_army](https://phishing.army/download/phishing_army_blocklist.txt) | ❌ Disabled | phishing | Covered by other sources |
| [quidsup_notrack-annoyance](https://quidsup.net/notrack/blocklist.php?download=annoyancedomains) | ❌ Disabled | annoyance | >90% overlap with HaGeZi Pro |
| [quidsup_notrack-malware](https://quidsup.net/notrack/blocklist.php?download=malwaredomains) | ✅ Enabled | malware | - |
| [quidsup_notrack-tracker](https://quidsup.net/notrack/blocklist.php?download=trackersdomains) | ✅ Enabled | trackers | - |
| [youtube_GoodbyeAds](https://raw.githubusercontent.com/jerryn70/GoodbyeAds/master/Extension/GoodbyeAds-YouTube-AdBlock.txt) | ✅ Enabled | ads | - |

</details>

<details>
<summary><strong>📄 sources_domain_new.json</strong> (1 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [nrd-14day-mini](https://raw.githubusercontent.com/xRuffKez/NRD/refs/heads/main/lists/14-day-mini/domains-only/nrd-14day-mini.txt) | ❌ Disabled | others | Huge list with low unique contribution |

</details>

<details>
<summary><strong>📄 sources_domain_top.json</strong> (1 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [tranco](https://tranco-list.eu/top-1m.csv.zip) | ✅ Enabled | topdomains | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [AlienVault_Reputation](https://reputation.alienvault.com/reputation.generic) | ❌ Disabled | malicious, threat | Not available anymore. The service has been discontinued. |
| [BinaryDefense_Banlist](https://www.binarydefense.com/banlist.txt) | ✅ Enabled | malicious, threat | This is for public use only. |
| [Blackhole_Today](https://blackhole.s-e-r-v-e-r.pw/blackhole-today) | ❌ Disabled | malicious, threat | Download fails frequently due to network instability or potential blocking. |
| [BlockListDE_Brute](https://lists.blocklist.de/lists/bruteforcelogin.txt) | ✅ Enabled | threat | >95% overlap with Firehol_level2 |
| [BlockListDE_Strong](https://lists.blocklist.de/lists/strongips.txt) | ✅ Enabled | malicious, threat | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| [Borestad_AbuseIPDB_S100_3d](https://raw.githubusercontent.com/borestad/blocklist-abuseipdb/refs/heads/main/abuseipdb-s100-3d.ipv4) | ✅ Enabled | malicious, threat | - |
| [BruteforceBlocker](http://danger.rulez.sk/projects/bruteforceblocker/blist.php) | ❌ Disabled | threat | >95% overlap with EmergingThreats_CompromisedIPs |
| [CINSScore_BadGuys_Army](http://cinsscore.com/list/ci-badguys.txt) | ✅ Enabled | malicious, threat | - |
| [DShield](https://opendbl.net/lists/dshield.list) | ❌ Disabled | malicious, threat | 100% overlap with Firehol_level2/Firehol_level3 |
| [DoH_IP_blocklists](https://raw.githubusercontent.com/dibdot/DoH-IP-blocklists/refs/heads/master/doh-ipv4.txt) | ✅ Enabled | doh | >90% overlap with HaGeZi Encrypted DNS Servers |
| [DoH_IP_list](https://raw.githubusercontent.com/oneoffdallas/dohservers/master/iplist.txt) | ✅ Enabled | doh | - |
| [ET_fwip](https://rules.emergingthreats.net/fwrules/emerging-Block-IPs.txt) | ✅ Enabled | malicious, threat | - |
| [EmergingThreats_CompromisedIPs](https://rules.emergingthreats.net/open/suricata/rules/compromised-ips.txt) | ✅ Enabled | malicious, threat | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| [FabrizioSalmi_DNS](https://raw.githubusercontent.com/fabriziosalmi/blacklists/refs/heads/main/dns_servers.txt) | ✅ Enabled | dns | - |
| [Firehol_BitcoinNodes_1d](https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/bitcoin_nodes_1d.ipset) | ✅ Enabled | cryptocurrency | - |
| [Firehol_Botscout_1d](https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/botscout_1d.ipset) | ✅ Enabled | malicious, threat | - |
| [Firehol_CleanTalk](https://iplists.firehol.org/files/cleantalk.ipset) | ✅ Enabled | malicious, threat | - |
| [Firehol_CleanTalk_Top20](https://iplists.firehol.org/files/cleantalk_top20.ipset) | ✅ Enabled | malicious, threat | - |
| [Firehol_GPF_Comics](https://iplists.firehol.org/files/gpf_comics.ipset) | ✅ Enabled | malicious, threat | - |
| [Firehol_SSLProxies_1d](https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/sslproxies_1d.ipset) | ✅ Enabled | anonymizer, privacy, proxy | - |
| [Firehol_SocksProxy_7d](https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/socks_proxy_7d.ipset) | ✅ Enabled | anonymizer, privacy, proxy | - |
| [Firehol_abusers_30d](https://iplists.firehol.org/files/firehol_abusers_30d.netset) | ❌ Disabled | malicious, threat | False positives are common, use with caution. |
| [Firehol_level1](https://iplists.firehol.org/files/firehol_level1.netset) | ✅ Enabled | malicious, threat | - |
| [Firehol_level2](https://iplists.firehol.org/files/firehol_level2.netset) | ✅ Enabled | malicious, threat | - |
| [Firehol_level3](https://iplists.firehol.org/files/firehol_level3.netset) | ✅ Enabled | malicious, threat | - |
| [GlobalAntiScamOrg-blocklist-ips](https://raw.githubusercontent.com/elliotwutingfeng/GlobalAntiScamOrg-blocklist/refs/heads/main/global-anti-scam-org-scam-ips.txt) | ✅ Enabled | scam | - |
| [Greensnow](https://blocklist.greensnow.co/greensnow.txt) | ✅ Enabled | malicious, malware, threat | >95% overlap with Firehol_level2 |
| [HaGeZi_DoH](https://raw.githubusercontent.com/hagezi/dns-blocklists/refs/heads/main/ips/doh.txt) | ❌ Disabled | doh | >90% overlap with DoH_IP_blocklists |
| [HaGeZi_TIF](https://raw.githubusercontent.com/hagezi/dns-blocklists/main/ips/tif.txt) | ❌ Disabled | malicious, threat | No unique contribution |
| [MyIP_MS_Blocklist](https://myip.ms/files/blacklist/htaccess/latest_blacklist.txt) | ✅ Enabled | malicious, threat | - |
| [Public_DNS4](https://public-dns.info/nameservers.txt) | ✅ Enabled | dns | - |
| [Rutgers_DROP](https://report.cs.rutgers.edu/DROP/attackers) | ✅ Enabled | malicious, threat | - |
| [Sblam_Blocklist](https://sblam.com/blacklist.txt) | ✅ Enabled | spam | - |
| [ScriptzTeam_BadIPS](https://raw.githubusercontent.com/scriptzteam/badIPS/main/ips.txt) | ✅ Enabled | malicious, threat | - |
| [Sentinel_Greylist](https://view.sentinel.turris.cz/greylist-data/greylist-latest.csv) | ✅ Enabled | malicious, threat | - |
| [T145_allowlist-ips](https://github.com/T145/black-mirror/releases/download/latest/ALLOW_IPV4.txt) | ❌ Disabled | others | Too many allowlist IPs, use with caution. |
| [T145_blocklist](https://github.com/T145/black-mirror/releases/download/latest/BLOCK_IPV4.txt) | ❌ Disabled | malicious, malware, threat | Too many blocklist IPs, use with caution. |
| [URLHaus_Text](https://urlhaus.abuse.ch/downloads/text/) | ✅ Enabled | malware | - |
| [USOM-Blocklists-ips](https://raw.githubusercontent.com/elliotwutingfeng/USOM-Blocklists/refs/heads/main/ips.txt) | ✅ Enabled | malicious, threat | - |
| [Yoyo AdServers-IPList](https://pgl.yoyo.org/adservers/iplist.php?showintro=1&mimetype=plaintext) | ✅ Enabled | ads | - |
| [spamhaus_drop](http://www.spamhaus.org/drop/drop.txt) | ✅ Enabled | spam, threat | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (5 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| Local Allowlist (AdGuard) | ✅ Enabled | local | - |
| Local Allowlist (Domain) | ✅ Enabled | local | - |
| Local Allowlist (ipv4) | ✅ Enabled | local | - |
| Local Blocklist (AdGuard) | ✅ Enabled | local | - |
| Local Blocklist (Domain) | ✅ Enabled | local | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | Status | Categories | Notes |
|------|--------|------------|-------|
| [VXVault_URLList](https://vxvault.net/URL_List.php) | ❌ Disabled | malware | >95% overlap with Firehol_level3 |

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
