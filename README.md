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

## Published Outputs

**Ready-to-use blocklist files are published daily to the [`output`](https://github.com/phani-kb/dns-toolkit/tree/output) branch:**

- Domain and IP blocklists/allowlists compatible with Pi-hole, pfBlockerNG, AdGuard Home
- Lists organized by size (mini, lite, normal, big) and category (ads, malware, privacy)
- Top entries based on source frequency for high-confidence blocking

**Usage:** Add `https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/[filename]` to your DNS filtering tool.

**[View Detailed Overlap Analysis](https://github.com/phani-kb/dns-toolkit/blob/output/overlap.md)** Comprehensive analysis showing how entries are shared across different DNS sources.

**[Conflicts Report (allowlist vs blocklist)](https://github.com/phani-kb/dns-toolkit/blob/output/conflicts.md)** A daily-generated markdown report listing entries found in both allowlists and blocklists, including the source(s) where they were found.

> 🔍 **Overlap Analysis**
>
> **Why it matters:** the overlap analysis helps to identify redundant or conflicting sources.
>
> - **Unique Entries = 0** → source is fully covered by same-list sources (low value-add).
> - **Conflicts > 0** → entries appear in different list types (e.g., blocklist vs. allowlist); investigate mismatches.
> - **Overlap % (table)** → shown relative to the *target*; high values mean the target is largely covered by this source.
>
> Run `dns-toolkit overlap` or open `overlap.md` to explore details.
> 
**Processing summaries and metadata are archived in the [`summaries`](https://github.com/phani-kb/dns-toolkit/tree/summaries) branch with 1-year retention.**

<!-- BRANCH_SIZES_START -->
## Branch Sizes

**Note:** The repo size badge above only reflects the default branch (`main`).

- **Output branch size:** 385.78 MB
- **Summaries branch size:** 1.30 MB

<!-- BRANCH_SIZES_END -->

---

<!-- STATS_START -->
## Source Statistics

*Automatically generated statistics from source configuration files*

| Metric | Count | Details |
|--------|-------|---------|
| **Total Sources** | 171 | 145 enabled, 26 disabled |
| **Blocklist Sources** | 146 | Sources providing blocking rules |
| **Allowlist Sources** | 37 | Sources providing exception rules |
| **Categories** | 39 | ads, adult, ai, annoyance, anonymizer, botnet, browser, cryptocurrency, dating, discord, dns, doh, fake, fakenews, finance, gambling, issues, kad, local, mac, malicious, malware, mobile, others, phishing, privacy, proxy, ransomware, scam, smarttv, social, spam, spyware, threat, topdomains, torrent_trackers, trackers, url_shorteners, windows |
| **Source Types** | 27 | adguard, cidr_ipv4, domain, domain_adguard, domain_comment, domain_csv_http_url_find, domain_custom_csv_blackbook, domain_custom_csv_maltrail, domain_custom_html_ccam, domain_custom_html_puppyscams, domain_http_url, domain_top, domain_url, domain_with_comment_suffix, hostname, ipv4, ipv4_cidr_expand, ipv4_csv_http_url_find, ipv4_custom_html_ccam, ipv4_find, ipv4_from_domain, ipv4_http_url, ipv4_range_expand, ipv4_url, ipv6, ipv6_find, ipv6_htaccess |
| **Geographic Coverage** | 21 countries | CN, CZ, DE, ES, FI, FR, HU, ID, IL, IT, KR, LV, MY, NL, RO, RU, SA, SK, UA, US, VN |
| **Last Updated** | 2025-08-27 20:27:37 UTC | Statistics generation time |

<!-- STATS_END -->

<!-- CREDITS_START -->
## Credits

This project is made possible by the following blocklist and allowlist sources:

Legend: S = Status, C/U/X = Count / Unique / Conflicts

<details>
<summary><strong>📄 sources_domain_al.json</strong> (19 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| AdGuardSDNSFilter_exclusions | ✅ | others | - | - |
| AdGuardTeam_HttpsExclusions_android | ✅ | mobile | 97 / 75 / 10 | - |
| AdGuardTeam_HttpsExclusions_banks | ✅ | finance | 3971 / 3938 / 15 | - |
| AdGuardTeam_HttpsExclusions_firefox | ✅ | browser | 18 / 13 / 0 | - |
| AdGuardTeam_HttpsExclusions_issues | ✅ | issues | 68 / 60 / 3 | - |
| AdGuardTeam_HttpsExclusions_mac | ✅ | mac | 11 / 5 / 0 | - |
| AdGuardTeam_HttpsExclusions_sensitive | ✅ | others | 169 / 142 / 12 | - |
| AdGuardTeam_HttpsExclusions_windows | ✅ | windows | 7 / 6 / 0 | - |
| anudeepND_Allowlist | ❌ | others | - | Last updated on 2021-12-01. This list is no longer maintained. |
| BlahDNS_whitelist | ❌ | others | - | Too many conflicts with other sources |
| China_CDN_Whitelist | ❌ | others | - | - |
| DandelionSprout_AdGuardHome_Whitelist | ✅ | others | 285 / 40 / 0 | - |
| Dogino_Discord_Official | ✅ | discord | 43 / 0 / 7 | - |
| fabriziosalmi_allowlist | ✅ | others | 2256 / 1080 / 759 | - |
| Freekers_Whitelist | ❌ | others | - | No update since 2019 |
| Notracking_Hosts_whitelist | ❌ | others | - | Archived by the owner on Aug 8, 2023 |
| ShadowWhisperer_Allowlist | ✅ | others | 654 / 317 / 246 | - |
| T145_allowlist-domains | ❌ | others | - | Huge list, use with caution |
| TogoFire_AD_Settings_whitelist | ✅ | others | 1764 / 1519 / 0 | Huge list, use with caution |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (101 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| 1Hosts (Lite) | ✅ | ads, trackers | 128566 / 0 / 184 | 100% covered by other sources |
| abpvn_hosts | ✅ | ads | 1061 / 945 / 0 | - |
| Adaway | ✅ | ads | 6540 / 0 / 111 | >99% overlap with StevenBlack Fake Gambling list |
| AdBlockID | ✅ | ads | 3848 / 3815 / 0 | - |
| AdGuard Base filter | ✅ | ads, trackers | 98051 / 0 / 0 | - |
| AdGuard CNAME Mail Trackers | ✅ | trackers | 37639 / 37587 / 0 | - |
| AdGuard CNAME Trackers | ✅ | trackers | 75301 / 49354 / 10 | - |
| AdGuard DNS filter | ✅ | ads, trackers | 132411 / 0 / 0 | - |
| AdGuard Spyware Filter - Mobile | ✅ | ads, mobile, spyware | 1078 / 0 / 0 | - |
| AntiAdBlockFilters | ✅ | annoyance | 1710 / 1705 / 0 | - |
| anudeepND_adservers | ❌ | ads | - | No update since 2023-01-16 |
| bigdargon_hostsVN | ✅ | ads | 18974 / 0 / 153 | - |
| Blocklists UT1 Cryptojacking | ✅ | cryptocurrency | 16291 / 14977 / 8 | - |
| Blocklists UT1 Malware | ✅ | malware | 225392 / 0 / 4 | >80% overlap with phishing_army |
| Blocklists UT1 Publicite | ✅ | ads | 4270 / 0 / 129 | 100% covered by other sources |
| Blocklists UT1 Shortener | ✅ | url_shorteners | 4518 / 0 / 21 | - |
| Boutetnico_URL_Shorteners | ✅ | url_shorteners | 418 / 216 / 24 | - |
| Cameleon | ❌ | ads | - | No update since 2018-03-17 |
| CF Torrent Trackers | ✅ | torrent_trackers | - | - |
| CJX Annoyance | ✅ | annoyance | 1800 / 1722 / 0 | - |
| CybercrimeTracker_All | ✅ | botnet, malicious, malware | 2864 / 1742 / 0 | - |
| CybercrimeTracker_CCAM | ❌ | botnet, malicious, malware | - | No regular updates |
| CybercrimeTracker_CCPMGate | ✅ | botnet, malicious, malware | 103 / 34 / 0 | - |
| cyberhost_malware-blocklist | ✅ | malware | 17426 / 36 / 3 | - |
| Dan Pollock's List | ✅ | ads, malware, trackers | 11806 / 0 / 32 | >95% overlap with StevenBlack Fake Gambling list |
| DandelionSprout-Anti-Malware-List | ✅ | malware | 32640 / 32629 / 0 | - |
| Easy Privacy | ✅ | privacy, trackers | 53302 / 16088 / 1 | - |
| EasyList | ✅ | ads | 55080 / 0 / 0 | 100% covered by other sources |
| fabriziosalmi_blocklists | ❌ | malicious, threat | - | Huge list, >3 million entries |
| FadeMind_2o7Net | ❌ | ads, privacy, trackers | - | No update since 2023-11-30 |
| FakeWebshopListHUN | ✅ | fake, phishing, scam, threat | 8210 / 4734 / 1 | - |
| Frogeye-firstparty-trackers | ✅ | trackers | 33252 / 9876 / 11 | - |
| GetAdmiral Domains Filter List | ✅ | ads, annoyance | 1770 / 0 / 0 | - |
| GlobalAntiScamOrg-blocklist-domains | ✅ | scam | 11065 / 7358 / 2 | - |
| HaGeZi Amazon Tracker | ✅ | privacy, trackers | 615 / 0 / 34 | >98% overlap with HaGeZi Pro |
| HaGeZi Apple Tracker | ✅ | privacy, trackers | 290 / 0 / 7 | >80% overlap with HaGeZi Pro |
| HaGeZi DNS TIF Mini | ✅ | malicious, threat | 120519 / 1270 / 1 | 100% covered by other sources |
| HaGeZi Encrypted DNS Servers | ✅ | doh | 1438 / 260 / 9 | - |
| HaGeZi Gambling Only Domains | ✅ | gambling | 185977 / 179602 / 5 | Huge list and gambling-specific focus |
| HaGeZi Microsoft Tracker | ✅ | privacy, trackers | 971 / 0 / 12 | >75% overlap with HaGeZi Pro |
| HaGeZi Most Abused TLDs | ✅ | spam | 429 / 427 / 0 | - |
| HaGeZi Normal | ❌ | ads, malware, trackers | - | 100% overlap with HaGeZi Pro |
| HaGeZi Pro | ✅ | ads, malware, phishing, trackers | 409765 / 0 / 191 | - |
| HaGeZi Xiaomi Tracker | ✅ | privacy, trackers | 477 / 0 / 0 | >95% overlap with HaGeZi Pro |
| Hestat_Minerchk | ❌ | cryptocurrency | - | No update since 2018 |
| hkamran80_smarttv | ✅ | smarttv | 293 / 0 / 15 | - |
| Hostsfile | ❌ | ads | - | No update since 2018-04-20 |
| hufilter | ✅ | ads | 100 / 0 / 3 | >90% overlap with HaGeZi Pro |
| iam-py-test_my-filters-001-antitypo | ✅ | fake | 824 / 822 / 0 | - |
| jarelllama_Scam-Blocklist | ✅ | scam | 457737 / 410604 / 11 | Disabled due to very large size (457K entries) - scam-specific focus |
| kadantiscam | ✅ | kad | 230104 / 0 / 2 | peer-to-peer network protocol |
| Korlabs_UrlShortener | ✅ | url_shorteners | 237 / 0 / 17 | - |
| Malicious URL Blocklist (URLHaus) | ✅ | ads | 2147 / 0 / 0 | 100% covered by other sources |
| Maltrail_StaticTrails | ✅ | malware, threat | 203918 / 172277 / 5 | - |
| malware-filter_phishing-filter | ✅ | malware, phishing | 25711 / 0 / 0 | - |
| OISD Blocklist Big | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | 201817 / 0 / 77 | Huge list |
| OISD Blocklist NSFW Small | ✅ | adult | 19923 / 0 / 14 | - |
| OISD Blocklist Small | ✅ | ads, cryptocurrency, malware, phishing, ransomware, trackers | 44797 / 0 / 51 | - |
| OpenPhish_Feed | ✅ | phishing | 282 / 0 / 0 | - |
| Peter Lowe's Blocklist | ✅ | ads | 3421 / 0 / 98 | 100% covered by other sources |
| pexcn Torrent Trackers | ✅ | torrent_trackers | - | - |
| phishing_army | ✅ | phishing | 122462 / 0 / 2 | - |
| Policeman_SimpleDomainsBlocklist | ❌ | malicious | - | Archived on 2021-12-26 |
| PuppyScams | ✅ | fake, scam | 102 / 85 / 0 | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| quidsup_notrack-annoyance | ✅ | annoyance | 475 / 0 / 1 | >90% overlap with HaGeZi Pro |
| quidsup_notrack-malware | ✅ | malware | 150 / 0 / 1 | - |
| quidsup_notrack-tracker | ✅ | trackers | 15679 / 0 / 163 | - |
| RedDragonWebDesign_block-everything | ✅ | ads, malicious, trackers | 652 / 648 / 0 | - |
| RPiList_specials-malware | ✅ | malware | 641579 / 344051 / 0 | Huge list |
| RPiList_specials-phishing | ✅ | phishing | 796259 / 482523 / 0 | Huge list |
| ShadowWhisperer's Dating List | ✅ | dating | - | - |
| ShadowWhisperer_BlockLists Ads | ✅ | ads | 23277 / 0 / 76 | - |
| ShadowWhisperer_BlockLists Adult | ✅ | adult | 283254 / 214945 / 14 | Huge list and adult-specific focus |
| ShadowWhisperer_BlockLists Malware | ✅ | malware | 52055 / 3075 / 4 | - |
| ShadowWhisperer_BlockLists Scam | ✅ | scam | 11083 / 7621 / 0 | - |
| ShadowWhisperer_UrlShortener | ✅ | url_shorteners | 5657 / 1056 / 2 | - |
| Sinfonietta_Adult | ✅ | adult | 58949 / 0 / 20 | - |
| Sinfonietta_Gambling | ✅ | gambling | 2639 / 0 / 1 | - |
| Sinfonietta_Social | ✅ | social | 3242 / 2786 / 106 | - |
| Spam404 | ✅ | spam | 8141 / 5655 / 1 | - |
| Stamparm_Blackbook | ✅ | malicious, threat | 18145 / 0 / 0 | >95% overlap with Blocklists UT1 Malware |
| StevenBlack_Adhoc_list | ❌ | ads, malware, trackers | - | 100% overlap with StevenBlack Fake Gambling list |
| StevenBlack_Fake_Gambling_Porn | ✅ | ads, adult, fake, fakenews, gambling | 317834 / 0 / 236 | - |
| T145_black-mirror | ❌ | malicious, threat | - | Huge list, >8 million entries |
| Torrent Trackers | ✅ | torrent_trackers | 485 / 459 / 0 | - |
| Ukrainian Ad Filter | ✅ | ads | 1449 / 1246 / 0 | - |
| Ukrainian Annoyance Filter | ✅ | annoyance | - | - |
| Ukrainian Privacy Filter | ✅ | privacy, trackers | 345 / 48 / 1 | - |
| Ukrainian Security Filter | ✅ | malicious, threat | 1736 / 1165 / 0 | - |
| UncheckyAds | ❌ | ads, privacy, trackers | - | No update since 2021 |
| URLHaus (Abuse.ch) | ✅ | malware | - | - |
| USOM-Blocklists-domains | ✅ | malicious, threat | 404591 / 350568 / 18 | Huge list |
| Viriback_Dump | ✅ | malware | 4594 / 0 / 0 | - |
| WaLLy3K | ✅ | ads | 350 / 0 / 12 | - |
| Warui_Adhosts | ✅ | ads | 75775 / 0 / 246 | Huge list |
| WindowsSpyBlocker_Hosts_spy | ❌ | privacy, trackers | - | No update since 2022-05-16 |
| Winhelp2002 | ❌ | ads | - | No update since 2021-03-06 |
| YousList | ✅ | ads | 624 / 0 / 3 | - |
| YousList-AdGuard | ✅ | ads | 7360 / 7184 / 0 | - |
| youtube_GoodbyeAds | ✅ | ads | 97645 / 97165 / 15 | - |
| Yoyo Adservers-Hosts | ✅ | ads | 3421 / 0 / 98 | >95% overlap with StevenBlack Fake Gambling list |

</details>

<details>
<summary><strong>📄 sources_domain_new.json</strong> (1 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| nrd-14day-mini | ❌ | others | - | Huge list with low unique contribution |

</details>

<details>
<summary><strong>📄 sources_domain_top.json</strong> (1 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| tranco | ✅ | topdomains | 1000 / 0 / 1253 | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| AlienVault_Reputation | ❌ | malicious, threat | - | Not available anymore. The service has been discontinued. |
| BinaryDefense_Banlist | ✅ | malicious, threat | 3023 / 0 / 0 | This is for public use only. |
| Blackhole_Today | ❌ | malicious, threat | - | Download fails frequently due to network instability or potential blocking. |
| BlockListDE_Brute | ✅ | threat | 659 / 0 / 0 | >95% overlap with Firehol_level2 |
| BlockListDE_Strong | ✅ | malicious, threat | 249 / 0 / 0 | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| Borestad_AbuseIPDB_S100_3d | ✅ | malicious, threat | 80682 / 0 / 0 | - |
| BruteforceBlocker | ✅ | threat | 404 / 0 / 0 | >95% overlap with EmergingThreats_CompromisedIPs |
| CINSScore_BadGuys_Army | ✅ | malicious, threat | 15000 / 0 / 0 | - |
| DoH_IP_blocklists | ✅ | doh | 2565 / 660 / 10 | >90% overlap with HaGeZi Encrypted DNS Servers |
| DoH_IP_list | ✅ | doh | 731 / 0 / 0 | - |
| DShield | ✅ | malicious, threat | 5120 / 0 / 0 | 100% overlap with Firehol_level2/Firehol_level3 |
| EmergingThreats_CompromisedIPs | ✅ | malicious, threat | 409 / 0 / 0 | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| ET_fwip | ✅ | malicious, threat | 1597 / 125 / 0 | - |
| FabrizioSalmi_DNS | ✅ | dns | 66 / 0 / 0 | - |
| Firehol_abusers_30d | ❌ | malicious, threat | - | False positives are common, use with caution. |
| Firehol_BitcoinNodes_1d | ✅ | cryptocurrency | 7249 / 7119 / 0 | - |
| Firehol_Botscout_1d | ✅ | malicious, threat | 718 / 621 / 0 | - |
| Firehol_CleanTalk | ✅ | malicious, threat | 494 / 413 / 0 | - |
| Firehol_CleanTalk_Top20 | ✅ | malicious, threat | 20 / 1 / 0 | - |
| Firehol_GPF_Comics | ✅ | malicious, threat | 2434 / 1084 / 0 | - |
| Firehol_level1 | ✅ | malicious, threat | 4523 / 3051 / 0 | - |
| Firehol_level2 | ✅ | malicious, threat | 14252 / 0 / 0 | - |
| Firehol_level3 | ✅ | malicious, threat | 20222 / 0 / 3 | - |
| Firehol_SocksProxy_7d | ✅ | anonymizer, privacy, proxy | 2547 / 2298 / 0 | - |
| Firehol_SSLProxies_1d | ✅ | anonymizer, privacy, proxy | 299 / 223 / 0 | - |
| GlobalAntiScamOrg-blocklist-ips | ✅ | scam | - | - |
| Greensnow | ✅ | malicious, malware, threat | 5863 / 0 / 0 | >95% overlap with Firehol_level2 |
| HaGeZi_DoH | ✅ | doh | 1705 / 0 / 0 | >90% overlap with DoH_IP_blocklists |
| HaGeZi_TIF | ✅ | malicious, threat | 65060 / 0 / 0 | No unique contribution |
| MyIP_MS_Blocklist | ✅ | malicious, threat | - | - |
| Public_DNS4 | ✅ | dns | 62607 / 61665 / 0 | - |
| Rutgers_DROP | ✅ | malicious, threat | 1968 / 0 / 0 | - |
| Sblam_Blocklist | ✅ | spam | 1809 / 1038 / 0 | - |
| ScriptzTeam_BadIPS | ✅ | malicious, threat | 2567 / 935 / 0 | - |
| Sentinel_Greylist | ✅ | malicious, threat | 9443 / 0 / 1 | - |
| spamhaus_drop | ✅ | spam, threat | - | - |
| T145_allowlist-ips | ❌ | others | - | Huge list, use with caution. More than its blocklist counterpart. |
| T145_blocklist | ❌ | malicious, malware, threat | - | Huge list, use with caution. |
| URLHaus_Text | ✅ | malware | 15748 / 0 / 0 | - |
| USOM-Blocklists-ips | ✅ | malicious, threat | 12563 / 0 / 0 | - |
| Yoyo AdServers-IPList | ✅ | ads | 8950 / 8898 / 0 | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (7 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
| Local AI Allowlist (Domain) | ✅ | ai, local | 49 / 0 / 51 | - |
| Local AI Blocklist (Domain) | ✅ | ai, local | 49 / 0 / 51 | - |
| Local Allowlist (AdGuard) | ✅ | local | - | - |
| Local Allowlist (Domain) | ✅ | local | 48 / 18 / 8 | - |
| Local Allowlist (ipv4) | ✅ | local | 76 / 55 / 19 | - |
| Local Blocklist (AdGuard) | ✅ | local | 7 / 0 / 0 | - |
| Local Blocklist (Domain) | ✅ | local | 1 / 0 / 0 | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | S | Categories |         C / U / X        | Notes |
|------|---|------------|--------------------------|-------|
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
4. **Validate commit messages** before pushing. Commit messages must reference a GitHub issue.
5. Submit a pull request

## Issues

If you encounter a bug, have a feature request, or want to suggest an improvement, please open an issue on the [GitHub Issues](https://github.com/phani-kb/dns-toolkit/issues) page.

## License

This project is licensed under the terms specified in the LICENSE file.
