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

## Published Outputs

**Ready-to-use blocklist files are published daily to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch:**

- Domain and IP blocklists/allowlists compatible with Pi-hole, pfBlockerNG, AdGuard Home
- Lists organized by size (mini, lite, normal, big) and category (ads, malware, privacy)
- Top entries based on source frequency for high-confidence blocking

**Usage:** Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your DNS filtering tool.

**[View Detailed Overlap Analysis →](https://github.com/phani-kb/dns-toolkit/blob/output/overlap.md)** - Comprehensive analysis showing how entries are shared across different DNS sources.

> 🔍 **Overlap report**
>
> **Why it matters:** the overlap report helps you spot redundant or conflicting sources.
>
> - **Unique Entries = 0** → source is fully covered by same-list sources (low value-add).
> - **Conflicts > 0** → entries appear in different list types (e.g., blocklist vs allowlist); investigate mismatches.
> - **Overlap % (table)** → shown relative to the *target*; high values mean the target is largely covered by this source.
>
> Run `dns-toolkit overlap` or open `overlap.md` to explore details.
> 
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
| **Last Updated** | 2025-08-24 19:04:18 UTC | Statistics generation time |

<!-- STATS_END -->

<!-- CREDITS_START -->
## Credits

This project is made possible by the following blocklist and allowlist sources:

Legend: S = Status, C/U/X = Count / Unique / Conflicts

<details>
<summary><strong>📄 sources_domain_al.json</strong> (20 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| AdGuardSDNSFilter_exclusions | ✅ | others | - | - |
| AdGuardTeam_HttpsExclusions_android | ✅ | mobile | <span style="white-space:nowrap">97 / 68 / 9</span> | - |
| AdGuardTeam_HttpsExclusions_banks | ✅ | finance | <span style="white-space:nowrap">3971 / 3922 / 14</span> | - |
| AdGuardTeam_HttpsExclusions_firefox | ✅ | browser | <span style="white-space:nowrap">18 / 10 / 0</span> | - |
| AdGuardTeam_HttpsExclusions_issues | ✅ | issues | <span style="white-space:nowrap">68 / 60 / 3</span> | - |
| AdGuardTeam_HttpsExclusions_mac | ✅ | mac | <span style="white-space:nowrap">11 / 4 / 0</span> | - |
| AdGuardTeam_HttpsExclusions_sensitive | ✅ | others | <span style="white-space:nowrap">164 / 133 / 12</span> | - |
| AdGuardTeam_HttpsExclusions_windows | ✅ | windows | <span style="white-space:nowrap">7 / 6 / 0</span> | - |
| anudeepND_Allowlist | ❌ | others | - | Last updated on 2021-12-01. This list is no longer maintained. |
| BlahDNS_whitelist | ✅ | others | <span style="white-space:nowrap">773 / 0 / 481</span> | - |
| China_CDN_Whitelist | ❌ | others | - | - |
| DandelionSprout_AdGuardHome_Whitelist | ✅ | others | <span style="white-space:nowrap">285 / 40 / 0</span> | - |
| Dogino_Discord_Official | ✅ | discord | <span style="white-space:nowrap">43 / 0 / 7</span> | - |
| fabriziosalmi_allowlist | ✅ | others | <span style="white-space:nowrap">2256 / 557 / 647</span> | - |
| Freekers_Whitelist | ❌ | others | - | No update since 2019 |
| Notracking_Hosts_whitelist | ✅ | others | <span style="white-space:nowrap">1979 / 0 / 1293</span> | Huge list, use with caution |
| ShadowWhisperer_Allowlist | ✅ | others | <span style="white-space:nowrap">653 / 219 / 219</span> | - |
| ShadowWhisperer_Whitelist | ✅ | others | - | - |
| T145_allowlist-domains | ❌ | others | - | Huge list, use with caution |
| TogoFire_AD_Settings_whitelist | ✅ | others | <span style="white-space:nowrap">1764 / 1519 / 0</span> | Huge list, use with caution |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (99 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| 1Hosts (Lite) | ✅ | ads, trackers | <span style="white-space:nowrap">128566 / 0 / 516</span> | 100% covered by other sources |
| abpvn_hosts | ✅ | ads | <span style="white-space:nowrap">1071 / 954 / 0</span> | - |
| Adaway | ✅ | ads | <span style="white-space:nowrap">6540 / 0 / 271</span> | >99% overlap with StevenBlack Fake Gambling list |
| AdBlockID | ✅ | ads | <span style="white-space:nowrap">3847 / 3814 / 0</span> | - |
| AdGuard Base filter | ✅ | ads, trackers | <span style="white-space:nowrap">96814 / 0 / 0</span> | - |
| AdGuard CNAME Mail Trackers | ✅ | trackers | <span style="white-space:nowrap">32718 / 32668 / 2</span> | - |
| AdGuard CNAME Trackers | ✅ | trackers | <span style="white-space:nowrap">84544 / 59167 / 46</span> | - |
| AdGuard DNS filter | ✅ | ads, trackers | <span style="white-space:nowrap">181 / 0 / 194</span> | - |
| AntiAdBlockFilters | ✅ | annoyance | <span style="white-space:nowrap">1710 / 1705 / 0</span> | - |
| anudeepND_adservers | ❌ | ads | - | No update since 2023-01-16 |
| bigdargon_hostsVN | ✅ | ads | <span style="white-space:nowrap">18967 / 0 / 435</span> | - |
| Blocklists UT1 Cryptojacking | ✅ | cryptocurrency | <span style="white-space:nowrap">16291 / 14974 / 25</span> | - |
| Blocklists UT1 Malware | ✅ | malware | <span style="white-space:nowrap">225392 / 0 / 15</span> | >80% overlap with phishing_army |
| Blocklists UT1 Publicite | ✅ | ads | <span style="white-space:nowrap">4270 / 0 / 254</span> | 100% covered by other sources |
| Blocklists UT1 Shortener | ✅ | url_shorteners | <span style="white-space:nowrap">4518 / 0 / 56</span> | - |
| Boutetnico_URL_Shorteners | ✅ | url_shorteners | <span style="white-space:nowrap">418 / 191 / 50</span> | - |
| Cameleon | ❌ | ads | - | No update since 2018-03-17 |
| CF Torrent Trackers | ✅ | torrent_trackers | - | - |
| CJX Annoyance | ✅ | annoyance | <span style="white-space:nowrap">6 / 4 / 2</span> | - |
| CybercrimeTracker_All | ✅ | botnet, malicious, malware | <span style="white-space:nowrap">2864 / 1743 / 0</span> | - |
| CybercrimeTracker_CCAM | ❌ | botnet, malicious, malware | - | No regular updates |
| CybercrimeTracker_CCPMGate | ✅ | botnet, malicious, malware | <span style="white-space:nowrap">103 / 34 / 0</span> | - |
| cyberhost_malware-blocklist | ✅ | malware | <span style="white-space:nowrap">17423 / 16 / 6</span> | - |
| Dan Pollock's List | ✅ | ads, malware, trackers | <span style="white-space:nowrap">11806 / 0 / 113</span> | >95% overlap with StevenBlack Fake Gambling list |
| DandelionSprout-Anti-Malware-List | ✅ | malware | <span style="white-space:nowrap">32640 / 32629 / 0</span> | - |
| Easy Privacy | ✅ | privacy, trackers | <span style="white-space:nowrap">655 / 0 / 949</span> | - |
| EasyList | ✅ | ads | <span style="white-space:nowrap">53907 / 0 / 0</span> | 100% covered by other sources |
| fabriziosalmi_blocklists | ❌ | malicious, threat | - | Huge list, >3 million entries |
| FadeMind_2o7Net | ❌ | ads, privacy, trackers | - | No update since 2023-11-30 |
| FakeWebshopListHUN | ✅ | fake, phishing, scam, threat | <span style="white-space:nowrap">8210 / 4735 / 2</span> | - |
| Frogeye-firstparty-trackers | ✅ | trackers | <span style="white-space:nowrap">33314 / 10215 / 53</span> | - |
| GetAdmiral Domains Filter List | ✅ | ads, annoyance | <span style="white-space:nowrap">2695 / 0 / 0</span> | - |
| GlobalAntiScamOrg-blocklist-domains | ✅ | scam | <span style="white-space:nowrap">11065 / 7341 / 3</span> | - |
| HaGeZi Amazon Tracker | ✅ | privacy, trackers | <span style="white-space:nowrap">615 / 0 / 38</span> | >98% overlap with HaGeZi Pro |
| HaGeZi Apple Tracker | ✅ | privacy, trackers | <span style="white-space:nowrap">290 / 0 / 14</span> | >80% overlap with HaGeZi Pro |
| HaGeZi DNS TIF Mini | ✅ | malicious, threat | <span style="white-space:nowrap">120736 / 2252 / 3</span> | 100% covered by other sources |
| HaGeZi Encrypted DNS Servers | ✅ | doh | <span style="white-space:nowrap">1437 / 246 / 11</span> | - |
| HaGeZi Gambling Only Domains | ✅ | gambling | <span style="white-space:nowrap">180321 / 174125 / 11</span> | Huge list and gambling-specific focus |
| HaGeZi Microsoft Tracker | ✅ | privacy, trackers | <span style="white-space:nowrap">971 / 0 / 36</span> | >75% overlap with HaGeZi Pro |
| HaGeZi Most Abused TLDs | ✅ | spam | <span style="white-space:nowrap">425 / 423 / 0</span> | - |
| HaGeZi Normal | ❌ | ads, malware, trackers | - | 100% overlap with HaGeZi Pro |
| HaGeZi Pro | ✅ | ads, malware, phishing, trackers | <span style="white-space:nowrap">407308 / 6434 / 500</span> | - |
| HaGeZi Xiaomi Tracker | ✅ | privacy, trackers | <span style="white-space:nowrap">475 / 0 / 15</span> | >95% overlap with HaGeZi Pro |
| Hestat_Minerchk | ❌ | cryptocurrency | - | No update since 2018 |
| hkamran80_smarttv | ✅ | smarttv | <span style="white-space:nowrap">293 / 0 / 31</span> | - |
| Hostsfile | ❌ | ads | - | No update since 2018-04-20 |
| hufilter | ✅ | ads | <span style="white-space:nowrap">100 / 0 / 5</span> | >90% overlap with HaGeZi Pro |
| iam-py-test_my-filters-001-antitypo | ✅ | fake | <span style="white-space:nowrap">824 / 822 / 0</span> | - |
| jarelllama_Scam-Blocklist | ✅ | scam | <span style="white-space:nowrap">457737 / 410290 / 21</span> | Disabled due to very large size (457K entries) - scam-specific focus |
| kadantiscam | ✅ | kad | <span style="white-space:nowrap">228972 / 0 / 6</span> | peer-to-peer network protocol |
| Korlabs_UrlShortener | ✅ | url_shorteners | <span style="white-space:nowrap">237 / 0 / 44</span> | - |
| Malicious URL Blocklist (URLHaus) | ✅ | ads | <span style="white-space:nowrap">2109 / 0 / 0</span> | 100% covered by other sources |
| Maltrail_StaticTrails | ✅ | malware, threat | <span style="white-space:nowrap">203541 / 171748 / 5</span> | - |
| malware-filter_phishing-filter | ✅ | malware, phishing | <span style="white-space:nowrap">25777 / 0 / 0</span> | - |
| OISD Blocklist Big | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | <span style="white-space:nowrap">201252 / 0 / 145</span> | Huge list |
| OISD Blocklist NSFW Small | ✅ | adult | <span style="white-space:nowrap">15634 / 0 / 31</span> | - |
| OISD Blocklist Small | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | <span style="white-space:nowrap">44379 / 0 / 95</span> | - |
| OpenPhish_Feed | ✅ | phishing | <span style="white-space:nowrap">260 / 72 / 3</span> | - |
| Peter Lowe's Blocklist | ✅ | ads | <span style="white-space:nowrap">3421 / 0 / 214</span> | 100% covered by other sources |
| pexcn Torrent Trackers | ✅ | torrent_trackers | - | - |
| phishing_army | ✅ | phishing | <span style="white-space:nowrap">121615 / 0 / 4</span> | - |
| Policeman_SimpleDomainsBlocklist | ❌ | malicious | - | Archived on 2021-12-26 |
| PuppyScams | ✅ | fake, scam | <span style="white-space:nowrap">102 / 85 / 0</span> | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| quidsup_notrack-annoyance | ✅ | annoyance | <span style="white-space:nowrap">475 / 0 / 6</span> | >90% overlap with HaGeZi Pro |
| quidsup_notrack-malware | ✅ | malware | <span style="white-space:nowrap">150 / 0 / 4</span> | - |
| quidsup_notrack-tracker | ✅ | trackers | <span style="white-space:nowrap">15690 / 0 / 366</span> | - |
| RedDragonWebDesign_block-everything | ✅ | ads, malicious, trackers | <span style="white-space:nowrap">652 / 648 / 0</span> | - |
| RPiList_specials-malware | ✅ | malware | <span style="white-space:nowrap">610524 / 314861 / 0</span> | Huge list |
| RPiList_specials-phishing | ✅ | phishing | <span style="white-space:nowrap">788316 / 479805 / 0</span> | Huge list |
| ShadowWhisperer's Dating List | ✅ | dating | - | - |
| ShadowWhisperer_BlockLists Ads | ✅ | ads | <span style="white-space:nowrap">23265 / 0 / 160</span> | - |
| ShadowWhisperer_BlockLists Adult | ✅ | adult | <span style="white-space:nowrap">277870 / 212268 / 27</span> | Huge list and adult-specific focus |
| ShadowWhisperer_BlockLists Malware | ✅ | malware | <span style="white-space:nowrap">52025 / 2795 / 14</span> | - |
| ShadowWhisperer_BlockLists Scam | ✅ | scam | <span style="white-space:nowrap">11082 / 7607 / 4</span> | - |
| ShadowWhisperer_UrlShortener | ✅ | url_shorteners | <span style="white-space:nowrap">5656 / 1053 / 8</span> | - |
| Sinfonietta_Adult | ✅ | adult | <span style="white-space:nowrap">58949 / 0 / 45</span> | - |
| Sinfonietta_Gambling | ✅ | gambling | <span style="white-space:nowrap">2639 / 0 / 6</span> | - |
| Sinfonietta_Social | ✅ | social | <span style="white-space:nowrap">3242 / 2752 / 178</span> | - |
| Spam404 | ✅ | spam | <span style="white-space:nowrap">8141 / 5669 / 8</span> | - |
| Stamparm_Blackbook | ✅ | malicious, threat | <span style="white-space:nowrap">18145 / 0 / 4</span> | >95% overlap with Blocklists UT1 Malware |
| StevenBlack_Adhoc_list | ❌ | ads, malware, trackers | - | 100% overlap with StevenBlack Fake Gambling list |
| StevenBlack_Fake_Gambling_Porn | ✅ | ads, adult, fake, fakenews, gambling | <span style="white-space:nowrap">317834 / 0 / 640</span> | - |
| T145_black-mirror | ❌ | malicious, threat | - | Huge list, >8 million entries |
| Torrent Trackers | ✅ | torrent_trackers | <span style="white-space:nowrap">485 / 455 / 2</span> | - |
| Ukrainian Ad Filter | ✅ | ads | <span style="white-space:nowrap">1449 / 1246 / 0</span> | - |
| Ukrainian Annoyance Filter | ✅ | annoyance | - | - |
| Ukrainian Privacy Filter | ✅ | privacy, trackers | <span style="white-space:nowrap">345 / 51 / 1</span> | - |
| Ukrainian Security Filter | ✅ | malicious, threat | <span style="white-space:nowrap">1736 / 1174 / 0</span> | - |
| UncheckyAds | ❌ | ads, privacy, trackers | - | No update since 2021 |
| URLHaus (Abuse.ch) | ✅ | malware | - | - |
| USOM-Blocklists-domains | ✅ | malicious, threat | <span style="white-space:nowrap">403814 / 350805 / 40</span> | Huge list |
| Viriback_Dump | ✅ | malware | <span style="white-space:nowrap">4584 / 0 / 0</span> | - |
| WaLLy3K | ✅ | ads | <span style="white-space:nowrap">350 / 0 / 45</span> | - |
| WindowsSpyBlocker_Hosts_spy | ❌ | privacy, trackers | - | No update since 2022-05-16 |
| Winhelp2002 | ❌ | ads | - | No update since 2021-03-06 |
| YousList | ✅ | ads | <span style="white-space:nowrap">624 / 0 / 8</span> | - |
| YousList-AdGuard | ✅ | ads | <span style="white-space:nowrap">12 / 0 / 19</span> | - |
| youtube_GoodbyeAds | ✅ | ads | <span style="white-space:nowrap">97645 / 97220 / 35</span> | - |
| Yoyo Adservers-Hosts | ✅ | ads | <span style="white-space:nowrap">3421 / 0 / 214</span> | >95% overlap with StevenBlack Fake Gambling list |

</details>

<details>
<summary><strong>📄 sources_domain_new.json</strong> (1 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| nrd-14day-mini | ❌ | others | - | Huge list with low unique contribution |

</details>

<details>
<summary><strong>📄 sources_domain_top.json</strong> (1 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| tranco | ✅ | topdomains | <span style="white-space:nowrap">1000 / 0 / 1151</span> | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| AlienVault_Reputation | ❌ | malicious, threat | - | Not available anymore. The service has been discontinued. |
| BinaryDefense_Banlist | ✅ | malicious, threat | <span style="white-space:nowrap">3023 / 0 / 0</span> | This is for public use only. |
| Blackhole_Today | ❌ | malicious, threat | - | Download fails frequently due to network instability or potential blocking. |
| BlockListDE_Brute | ✅ | threat | <span style="white-space:nowrap">960 / 0 / 0</span> | >95% overlap with Firehol_level2 |
| BlockListDE_Strong | ✅ | malicious, threat | <span style="white-space:nowrap">262 / 0 / 0</span> | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| Borestad_AbuseIPDB_S100_3d | ✅ | malicious, threat | <span style="white-space:nowrap">80682 / 0 / 0</span> | - |
| BruteforceBlocker | ✅ | threat | <span style="white-space:nowrap">433 / 0 / 0</span> | >95% overlap with EmergingThreats_CompromisedIPs |
| CINSScore_BadGuys_Army | ✅ | malicious, threat | <span style="white-space:nowrap">15000 / 0 / 0</span> | - |
| DoH_IP_blocklists | ✅ | doh | <span style="white-space:nowrap">2576 / 676 / 10</span> | >90% overlap with HaGeZi Encrypted DNS Servers |
| DoH_IP_list | ✅ | doh | <span style="white-space:nowrap">731 / 0 / 0</span> | - |
| DShield | ✅ | malicious, threat | <span style="white-space:nowrap">5120 / 0 / 0</span> | 100% overlap with Firehol_level2/Firehol_level3 |
| EmergingThreats_CompromisedIPs | ✅ | malicious, threat | <span style="white-space:nowrap">428 / 0 / 0</span> | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| ET_fwip | ✅ | malicious, threat | <span style="white-space:nowrap">1607 / 133 / 0</span> | - |
| FabrizioSalmi_DNS | ✅ | dns | <span style="white-space:nowrap">66 / 0 / 0</span> | - |
| Firehol_abusers_30d | ❌ | malicious, threat | - | False positives are common, use with caution. |
| Firehol_BitcoinNodes_1d | ✅ | cryptocurrency | <span style="white-space:nowrap">7190 / 7055 / 0</span> | - |
| Firehol_Botscout_1d | ✅ | malicious, threat | <span style="white-space:nowrap">636 / 467 / 0</span> | - |
| Firehol_CleanTalk | ✅ | malicious, threat | <span style="white-space:nowrap">494 / 415 / 0</span> | - |
| Firehol_CleanTalk_Top20 | ✅ | malicious, threat | <span style="white-space:nowrap">20 / 1 / 0</span> | - |
| Firehol_GPF_Comics | ✅ | malicious, threat | <span style="white-space:nowrap">2380 / 1085 / 0</span> | - |
| Firehol_level1 | ✅ | malicious, threat | <span style="white-space:nowrap">4529 / 3055 / 0</span> | - |
| Firehol_level2 | ✅ | malicious, threat | <span style="white-space:nowrap">14853 / 0 / 0</span> | - |
| Firehol_level3 | ✅ | malicious, threat | <span style="white-space:nowrap">12187 / 0 / 3</span> | - |
| Firehol_SocksProxy_7d | ✅ | anonymizer, privacy, proxy | <span style="white-space:nowrap">2547 / 2279 / 0</span> | - |
| Firehol_SSLProxies_1d | ✅ | anonymizer, privacy, proxy | <span style="white-space:nowrap">299 / 225 / 0</span> | - |
| GlobalAntiScamOrg-blocklist-ips | ✅ | scam | - | - |
| Greensnow | ✅ | malicious, malware, threat | <span style="white-space:nowrap">5491 / 0 / 0</span> | >95% overlap with Firehol_level2 |
| HaGeZi_DoH | ✅ | doh | <span style="white-space:nowrap">1707 / 0 / 0</span> | >90% overlap with DoH_IP_blocklists |
| HaGeZi_TIF | ✅ | malicious, threat | <span style="white-space:nowrap">67314 / 0 / 0</span> | No unique contribution |
| MyIP_MS_Blocklist | ✅ | malicious, threat | - | - |
| Public_DNS4 | ✅ | dns | <span style="white-space:nowrap">62607 / 61672 / 0</span> | - |
| Rutgers_DROP | ✅ | malicious, threat | <span style="white-space:nowrap">1999 / 0 / 0</span> | - |
| Sblam_Blocklist | ✅ | spam | <span style="white-space:nowrap">1765 / 949 / 0</span> | - |
| ScriptzTeam_BadIPS | ✅ | malicious, threat | <span style="white-space:nowrap">2567 / 886 / 0</span> | - |
| Sentinel_Greylist | ✅ | malicious, threat | <span style="white-space:nowrap">8687 / 0 / 1</span> | - |
| spamhaus_drop | ✅ | spam, threat | - | - |
| T145_allowlist-ips | ❌ | others | - | Huge list, use with caution. More than its blocklist counterpart. |
| T145_blocklist | ❌ | malicious, malware, threat | - | Huge list, use with caution. |
| URLHaus_Text | ✅ | malware | <span style="white-space:nowrap">16082 / 0 / 0</span> | - |
| USOM-Blocklists-ips | ✅ | malicious, threat | <span style="white-space:nowrap">12540 / 0 / 0</span> | - |
| Yoyo AdServers-IPList | ✅ | ads | <span style="white-space:nowrap">8950 / 8899 / 0</span> | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (7 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| Local AI Allowlist (Domain) | ✅ | local | <span style="white-space:nowrap">49 / 0 / 51</span> | - |
| Local AI Blocklist (Domain) | ✅ | local | <span style="white-space:nowrap">49 / 0 / 51</span> | - |
| Local Allowlist (AdGuard) | ✅ | local | - | - |
| Local Allowlist (Domain) | ✅ | local | <span style="white-space:nowrap">47 / 6 / 8</span> | - |
| Local Allowlist (ipv4) | ✅ | local | <span style="white-space:nowrap">74 / 53 / 19</span> | - |
| Local Blocklist (AdGuard) | ✅ | local | <span style="white-space:nowrap">7 / 0 / 0</span> | - |
| Local Blocklist (Domain) | ✅ | local | <span style="white-space:nowrap">1 / 0 / 1</span> | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| VXVault_URLList | ✅ | malware | <span style="white-space:nowrap">38 / 0 / 0</span> | >95% overlap with Firehol_level3 |

</details>

<!-- CREDITS_END -->

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
