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
| **Last Updated** | 2025-08-24 03:57:35 UTC | Statistics generation time |

<!-- STATS_END -->

<!-- CREDITS_START -->
## Source Credits

This project is made possible by the following blocklist and allowlist sources:

Legend: S = Status, C/U/X = Count / Unique / Conflicts

<details>
<summary><strong>📄 sources_domain_al.json</strong> (20 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| AdGuardSDNSFilter_exclusions | ✅ | others | - | - |
| AdGuardTeam_HttpsExclusions_android | ✅ | mobile | 97 / 68 / 9 | - |
| AdGuardTeam_HttpsExclusions_banks | ✅ | finance | 3971 / 3922 / 14 | - |
| AdGuardTeam_HttpsExclusions_firefox | ✅ | browser | 18 / 10 / 0 | - |
| AdGuardTeam_HttpsExclusions_issues | ✅ | issues | 68 / 60 / 3 | - |
| AdGuardTeam_HttpsExclusions_mac | ✅ | mac | 11 / 4 / 0 | - |
| AdGuardTeam_HttpsExclusions_sensitive | ✅ | others | 164 / 133 / 12 | - |
| AdGuardTeam_HttpsExclusions_windows | ✅ | windows | 7 / 6 / 0 | - |
| BlahDNS_whitelist | ✅ | others | 773 / 0 / 481 | - |
| China_CDN_Whitelist | ❌ | others | - | - |
| DandelionSprout_AdGuardHome_Whitelist | ✅ | others | 285 / 40 / 0 | - |
| Dogino_Discord_Official | ✅ | discord | 43 / 0 / 7 | - |
| Freekers_Whitelist | ❌ | others | - | No update since 2019 |
| Notracking_Hosts_whitelist | ✅ | others | 1979 / 0 / 1293 | Huge list, use with caution |
| ShadowWhisperer_Allowlist | ✅ | others | 653 / 219 / 219 | - |
| ShadowWhisperer_Whitelist | ✅ | others | - | - |
| T145_allowlist-domains | ❌ | others | - | Huge list, use with caution |
| TogoFire_AD_Settings_whitelist | ✅ | others | 1764 / 1519 / 0 | Huge list, use with caution |
| anudeepND_Allowlist | ❌ | others | - | Last updated on 2021-12-01. This list is no longer maintained. |
| fabriziosalmi_allowlist | ✅ | others | 2256 / 557 / 647 | - |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (99 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| 1Hosts (Lite) | ✅ | ads, trackers | 128566 / 0 / 516 | 100% covered by other sources |
| AdBlockID | ✅ | ads | 3847 / 3814 / 0 | - |
| AdGuard Base filter | ✅ | ads, trackers | 96814 / 0 / 0 | - |
| AdGuard CNAME Mail Trackers | ✅ | trackers | 32718 / 32668 / 2 | - |
| AdGuard CNAME Trackers | ✅ | trackers | 84544 / 59167 / 46 | - |
| AdGuard DNS filter | ✅ | ads, trackers | 181 / 0 / 194 | - |
| Adaway | ✅ | ads | 6540 / 0 / 271 | >99% overlap with StevenBlack Fake Gambling list |
| AntiAdBlockFilters | ✅ | annoyance | 1710 / 1705 / 0 | - |
| Blocklists UT1 Cryptojacking | ✅ | cryptocurrency | 16291 / 14975 / 25 | - |
| Blocklists UT1 Malware | ✅ | malware | 223679 / 0 / 16 | >80% overlap with phishing_army |
| Blocklists UT1 Publicite | ✅ | ads | 4270 / 0 / 254 | 100% covered by other sources |
| Blocklists UT1 Shortener | ✅ | url_shorteners | 4518 / 0 / 56 | - |
| Boutetnico_URL_Shorteners | ✅ | url_shorteners | 418 / 191 / 50 | - |
| CF Torrent Trackers | ✅ | torrent_trackers | - | - |
| CJX Annoyance | ✅ | annoyance | 6 / 4 / 2 | - |
| Cameleon | ❌ | ads | - | No update since 2018-03-17 |
| CybercrimeTracker_All | ✅ | botnet, malicious, malware | 2864 / 1743 / 0 | - |
| CybercrimeTracker_CCAM | ❌ | botnet, malicious, malware | - | No regular updates |
| CybercrimeTracker_CCPMGate | ✅ | botnet, malicious, malware | 103 / 34 / 0 | - |
| Dan Pollock's List | ✅ | ads, malware, trackers | 11806 / 0 / 113 | >95% overlap with StevenBlack Fake Gambling list |
| DandelionSprout-Anti-Malware-List | ✅ | malware | 32640 / 32629 / 0 | - |
| Easy Privacy | ✅ | privacy, trackers | 655 / 0 / 949 | - |
| EasyList | ✅ | ads | 53907 / 0 / 0 | 100% covered by other sources |
| FadeMind_2o7Net | ❌ | ads, privacy, trackers | - | No update since 2023-11-30 |
| FakeWebshopListHUN | ✅ | fake, phishing, scam, threat | 8210 / 4735 / 2 | - |
| Frogeye-firstparty-trackers | ✅ | trackers | 33314 / 10215 / 53 | - |
| GetAdmiral Domains Filter List | ✅ | ads, annoyance | 2695 / 0 / 0 | - |
| GlobalAntiScamOrg-blocklist-domains | ✅ | scam | 11061 / 7337 / 3 | - |
| HaGeZi Amazon Tracker | ✅ | privacy, trackers | 615 / 0 / 38 | >98% overlap with HaGeZi Pro |
| HaGeZi Apple Tracker | ✅ | privacy, trackers | 290 / 0 / 14 | >80% overlap with HaGeZi Pro |
| HaGeZi DNS TIF Mini | ✅ | malicious, threat | 120736 / 3785 / 3 | 100% covered by other sources |
| HaGeZi Encrypted DNS Servers | ✅ | doh | 1437 / 246 / 11 | - |
| HaGeZi Gambling Only Domains | ✅ | gambling | 180321 / 174125 / 11 | Huge list and gambling-specific focus |
| HaGeZi Microsoft Tracker | ✅ | privacy, trackers | 971 / 0 / 36 | >75% overlap with HaGeZi Pro |
| HaGeZi Most Abused TLDs | ✅ | spam | 425 / 423 / 0 | - |
| HaGeZi Normal | ❌ | ads, malware, trackers | - | 100% overlap with HaGeZi Pro |
| HaGeZi Pro | ✅ | ads, malware, phishing, trackers | 407308 / 8273 / 500 | - |
| HaGeZi Xiaomi Tracker | ✅ | privacy, trackers | 475 / 0 / 15 | >95% overlap with HaGeZi Pro |
| Hestat_Minerchk | ❌ | cryptocurrency | - | No update since 2018 |
| Hostsfile | ❌ | ads | - | No update since 2018-04-20 |
| Korlabs_UrlShortener | ✅ | url_shorteners | 237 / 0 / 44 | - |
| Malicious URL Blocklist (URLHaus) | ✅ | ads | 2109 / 0 / 0 | 100% covered by other sources |
| Maltrail_StaticTrails | ✅ | malware, threat | 203541 / 171748 / 5 | - |
| OISD Blocklist Big | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | 201252 / 0 / 145 | Huge list |
| OISD Blocklist NSFW Small | ✅ | adult | 15634 / 0 / 31 | - |
| OISD Blocklist Small | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | 44379 / 0 / 95 | - |
| OpenPhish_Feed | ✅ | phishing | 260 / 135 / 3 | - |
| Peter Lowe's Blocklist | ✅ | ads | 3421 / 0 / 214 | 100% covered by other sources |
| Policeman_SimpleDomainsBlocklist | ❌ | malicious | - | Archived on 2021-12-26 |
| PuppyScams | ✅ | fake, scam | 102 / 85 / 0 | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| RPiList_specials-malware | ✅ | malware | 610524 / 314861 / 0 | Huge list |
| RPiList_specials-phishing | ✅ | phishing | 788316 / 479805 / 0 | Huge list |
| RedDragonWebDesign_block-everything | ✅ | ads, malicious, trackers | 652 / 648 / 0 | - |
| ShadowWhisperer's Dating List | ✅ | dating | - | - |
| ShadowWhisperer_BlockLists Ads | ✅ | ads | 23265 / 0 / 160 | - |
| ShadowWhisperer_BlockLists Adult | ✅ | adult | 277844 / 212242 / 27 | Huge list and adult-specific focus |
| ShadowWhisperer_BlockLists Malware | ✅ | malware | 52007 / 2790 / 14 | - |
| ShadowWhisperer_BlockLists Scam | ✅ | scam | 11082 / 7607 / 4 | - |
| ShadowWhisperer_UrlShortener | ✅ | url_shorteners | 5639 / 1037 / 8 | - |
| Sinfonietta_Adult | ✅ | adult | 58949 / 0 / 45 | - |
| Sinfonietta_Gambling | ✅ | gambling | 2639 / 0 / 6 | - |
| Sinfonietta_Social | ✅ | social | 3242 / 2752 / 178 | - |
| Spam404 | ✅ | spam | 8141 / 5669 / 8 | - |
| Stamparm_Blackbook | ✅ | malicious, threat | 18145 / 0 / 4 | >95% overlap with Blocklists UT1 Malware |
| StevenBlack_Adhoc_list | ❌ | ads, malware, trackers | - | 100% overlap with StevenBlack Fake Gambling list |
| StevenBlack_Fake_Gambling_Porn | ✅ | ads, adult, fake, fakenews, gambling | 314461 / 0 / 640 | - |
| T145_black-mirror | ❌ | malicious, threat | - | Huge list, >8 million entries |
| Torrent Trackers | ✅ | torrent_trackers | 485 / 455 / 2 | - |
| URLHaus (Abuse.ch) | ✅ | malware | - | - |
| USOM-Blocklists-domains | ✅ | malicious, threat | 403814 / 350800 / 40 | Huge list |
| Ukrainian Ad Filter | ✅ | ads | 1449 / 1246 / 0 | - |
| Ukrainian Annoyance Filter | ✅ | annoyance | - | - |
| Ukrainian Privacy Filter | ✅ | privacy, trackers | 345 / 51 / 1 | - |
| Ukrainian Security Filter | ✅ | malicious, threat | 1736 / 1174 / 0 | - |
| UncheckyAds | ❌ | ads, privacy, trackers | - | No update since 2021 |
| Viriback_Dump | ✅ | malware | 4584 / 0 / 0 | - |
| WaLLy3K | ✅ | ads | 350 / 0 / 45 | - |
| WindowsSpyBlocker_Hosts_spy | ❌ | privacy, trackers | - | No update since 2022-05-16 |
| Winhelp2002 | ❌ | ads | - | No update since 2021-03-06 |
| YousList | ✅ | ads | 624 / 0 / 8 | - |
| YousList-AdGuard | ✅ | ads | 12 / 0 / 19 | - |
| Yoyo Adservers-Hosts | ✅ | ads | 3421 / 0 / 214 | >95% overlap with StevenBlack Fake Gambling list |
| abpvn_hosts | ✅ | ads | 1071 / 954 / 0 | - |
| anudeepND_adservers | ❌ | ads | - | No update since 2023-01-16 |
| bigdargon_hostsVN | ✅ | ads | 18967 / 0 / 435 | - |
| cyberhost_malware-blocklist | ✅ | malware | 17423 / 22 / 6 | - |
| fabriziosalmi_blocklists | ❌ | malicious, threat | - | Huge list, >3 million entries |
| hkamran80_smarttv | ✅ | smarttv | 293 / 0 / 31 | - |
| hufilter | ✅ | ads | 100 / 0 / 5 | >90% overlap with HaGeZi Pro |
| iam-py-test_my-filters-001-antitypo | ✅ | fake | 824 / 822 / 0 | - |
| jarelllama_Scam-Blocklist | ✅ | scam | 457737 / 410232 / 21 | Disabled due to very large size (457K entries) - scam-specific focus |
| kadantiscam | ✅ | kad | 228972 / 0 / 6 | peer-to-peer network protocol |
| malware-filter_phishing-filter | ✅ | malware, phishing | 25777 / 0 / 0 | - |
| pexcn Torrent Trackers | ✅ | torrent_trackers | - | - |
| phishing_army | ✅ | phishing | 121615 / 0 / 4 | - |
| quidsup_notrack-annoyance | ✅ | annoyance | 475 / 0 / 6 | >90% overlap with HaGeZi Pro |
| quidsup_notrack-malware | ✅ | malware | 150 / 0 / 4 | - |
| quidsup_notrack-tracker | ✅ | trackers | 15690 / 0 / 366 | - |
| youtube_GoodbyeAds | ✅ | ads | 97645 / 97220 / 35 | - |

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
| tranco | ✅ | topdomains | 1000 / 0 / 1152 | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| AlienVault_Reputation | ❌ | malicious, threat | - | Not available anymore. The service has been discontinued. |
| BinaryDefense_Banlist | ✅ | malicious, threat | 3023 / 0 / 0 | This is for public use only. |
| Blackhole_Today | ❌ | malicious, threat | - | Download fails frequently due to network instability or potential blocking. |
| BlockListDE_Brute | ✅ | threat | 960 / 0 / 0 | >95% overlap with Firehol_level2 |
| BlockListDE_Strong | ✅ | malicious, threat | 262 / 0 / 0 | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| Borestad_AbuseIPDB_S100_3d | ✅ | malicious, threat | 80682 / 0 / 0 | - |
| BruteforceBlocker | ✅ | threat | 433 / 0 / 0 | >95% overlap with EmergingThreats_CompromisedIPs |
| CINSScore_BadGuys_Army | ✅ | malicious, threat | 15000 / 0 / 0 | - |
| DShield | ✅ | malicious, threat | 5120 / 0 / 0 | 100% overlap with Firehol_level2/Firehol_level3 |
| DoH_IP_blocklists | ✅ | doh | 2576 / 676 / 10 | >90% overlap with HaGeZi Encrypted DNS Servers |
| DoH_IP_list | ✅ | doh | 731 / 0 / 0 | - |
| ET_fwip | ✅ | malicious, threat | 1607 / 133 / 0 | - |
| EmergingThreats_CompromisedIPs | ✅ | malicious, threat | 428 / 0 / 0 | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| FabrizioSalmi_DNS | ✅ | dns | 66 / 0 / 0 | - |
| Firehol_BitcoinNodes_1d | ✅ | cryptocurrency | 7190 / 7055 / 0 | - |
| Firehol_Botscout_1d | ✅ | malicious, threat | 636 / 467 / 0 | - |
| Firehol_CleanTalk | ✅ | malicious, threat | 494 / 415 / 0 | - |
| Firehol_CleanTalk_Top20 | ✅ | malicious, threat | 20 / 1 / 0 | - |
| Firehol_GPF_Comics | ✅ | malicious, threat | 2380 / 1085 / 0 | - |
| Firehol_SSLProxies_1d | ✅ | anonymizer, privacy, proxy | 299 / 225 / 0 | - |
| Firehol_SocksProxy_7d | ✅ | anonymizer, privacy, proxy | 2547 / 2279 / 0 | - |
| Firehol_abusers_30d | ❌ | malicious, threat | - | False positives are common, use with caution. |
| Firehol_level1 | ✅ | malicious, threat | 4529 / 3055 / 0 | - |
| Firehol_level2 | ✅ | malicious, threat | 14853 / 0 / 0 | - |
| Firehol_level3 | ✅ | malicious, threat | 12187 / 0 / 3 | - |
| GlobalAntiScamOrg-blocklist-ips | ✅ | scam | - | - |
| Greensnow | ✅ | malicious, malware, threat | 5491 / 0 / 0 | >95% overlap with Firehol_level2 |
| HaGeZi_DoH | ✅ | doh | 1707 / 0 / 0 | >90% overlap with DoH_IP_blocklists |
| HaGeZi_TIF | ✅ | malicious, threat | 67314 / 0 / 0 | No unique contribution |
| MyIP_MS_Blocklist | ✅ | malicious, threat | - | - |
| Public_DNS4 | ✅ | dns | 62607 / 61672 / 0 | - |
| Rutgers_DROP | ✅ | malicious, threat | 1999 / 0 / 0 | - |
| Sblam_Blocklist | ✅ | spam | 1765 / 949 / 0 | - |
| ScriptzTeam_BadIPS | ✅ | malicious, threat | 2567 / 886 / 0 | - |
| Sentinel_Greylist | ✅ | malicious, threat | 8687 / 0 / 1 | - |
| T145_allowlist-ips | ❌ | others | - | Huge list, use with caution. More than its blocklist counterpart. |
| T145_blocklist | ❌ | malicious, malware, threat | - | Huge list, use with caution. |
| URLHaus_Text | ✅ | malware | 16082 / 0 / 0 | - |
| USOM-Blocklists-ips | ✅ | malicious, threat | 12540 / 0 / 0 | - |
| Yoyo AdServers-IPList | ✅ | ads | 8950 / 8899 / 0 | - |
| spamhaus_drop | ✅ | spam, threat | - | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (7 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| Local AI Allowlist (Domain) | ✅ | local | 49 / 0 / 51 | - |
| Local AI Blocklist (Domain) | ✅ | local | 49 / 0 / 51 | - |
| Local Allowlist (AdGuard) | ✅ | local | - | - |
| Local Allowlist (Domain) | ✅ | local | 47 / 6 / 8 | - |
| Local Allowlist (ipv4) | ✅ | local | 76 / 55 / 19 | - |
| Local Blocklist (AdGuard) | ✅ | local | 7 / 0 / 0 | - |
| Local Blocklist (Domain) | ✅ | local | 1 / 0 / 1 | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | S | Categories | C/U/X | Notes |
|------|---|------------|-------|-------|
| VXVault_URLList | ✅ | malware | 38 / 0 / 0 | >95% overlap with Firehol_level3 |

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
