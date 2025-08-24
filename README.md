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
| **Last Updated** | 2025-08-24 02:53:33 UTC | Statistics generation time |

<!-- STATS_END -->

<!-- CREDITS_START -->
## Source Credits

This project is made possible by the following blocklist and allowlist sources:

<details>
<summary><strong>📄 sources_domain_al.json</strong> (20 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| AdGuardSDNSFilter_exclusions | ✅ Enabled | others | - | - |
| AdGuardTeam_HttpsExclusions_android | ✅ Enabled | mobile | 68/9 | - |
| AdGuardTeam_HttpsExclusions_banks | ✅ Enabled | finance | 3922/14 | - |
| AdGuardTeam_HttpsExclusions_firefox | ✅ Enabled | browser | 10/0 | - |
| AdGuardTeam_HttpsExclusions_issues | ✅ Enabled | issues | 60/3 | - |
| AdGuardTeam_HttpsExclusions_mac | ✅ Enabled | mac | 4/0 | - |
| AdGuardTeam_HttpsExclusions_sensitive | ✅ Enabled | others | 133/12 | - |
| AdGuardTeam_HttpsExclusions_windows | ✅ Enabled | windows | 6/0 | - |
| BlahDNS_whitelist | ✅ Enabled | others | 0/481 | - |
| China_CDN_Whitelist | ❌ Disabled | others | - | - |
| DandelionSprout_AdGuardHome_Whitelist | ✅ Enabled | others | 40/0 | - |
| Dogino_Discord_Official | ✅ Enabled | discord | 0/7 | - |
| Freekers_Whitelist | ❌ Disabled | others | - | No update since 2019 |
| Notracking_Hosts_whitelist | ✅ Enabled | others | 0/1293 | Huge list, use with caution |
| ShadowWhisperer_Allowlist | ✅ Enabled | others | 219/219 | - |
| ShadowWhisperer_Whitelist | ✅ Enabled | others | - | - |
| T145_allowlist-domains | ❌ Disabled | others | - | Huge list, use with caution |
| TogoFire_AD_Settings_whitelist | ✅ Enabled | others | 1519/0 | Huge list, use with caution |
| anudeepND_Allowlist | ❌ Disabled | others | - | Last updated on 2021-12-01. This list is no longer maintained. |
| fabriziosalmi_allowlist | ✅ Enabled | others | 557/647 | - |

</details>

<details>
<summary><strong>📄 sources_domain_bl.json</strong> (99 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| 1Hosts (Lite) | ✅ Enabled | ads, trackers | 0/516 | 100% covered by other sources |
| AdBlockID | ✅ Enabled | ads | 3814/0 | - |
| AdGuard Base filter | ✅ Enabled | ads, trackers | 0/0 | - |
| AdGuard CNAME Mail Trackers | ✅ Enabled | trackers | 32668/2 | - |
| AdGuard CNAME Trackers | ✅ Enabled | trackers | 59167/46 | - |
| AdGuard DNS filter | ✅ Enabled | ads, trackers | 0/194 | - |
| Adaway | ✅ Enabled | ads | 0/271 | >99% overlap with StevenBlack Fake Gambling list |
| AntiAdBlockFilters | ✅ Enabled | annoyance | 1705/0 | - |
| Blocklists UT1 Cryptojacking | ✅ Enabled | cryptocurrency | 14975/25 | - |
| Blocklists UT1 Malware | ✅ Enabled | malware | 0/16 | >80% overlap with phishing_army |
| Blocklists UT1 Publicite | ✅ Enabled | ads | 0/254 | 100% covered by other sources |
| Blocklists UT1 Shortener | ✅ Enabled | url_shorteners | 0/56 | - |
| Boutetnico_URL_Shorteners | ✅ Enabled | url_shorteners | 191/50 | - |
| CF Torrent Trackers | ✅ Enabled | torrent_trackers | - | - |
| CJX Annoyance | ✅ Enabled | annoyance | 4/2 | - |
| Cameleon | ❌ Disabled | ads | - | No update since 2018-03-17 |
| CybercrimeTracker_All | ✅ Enabled | botnet, malicious, malware | 1743/0 | - |
| CybercrimeTracker_CCAM | ❌ Disabled | botnet, malicious, malware | - | No regular updates |
| CybercrimeTracker_CCPMGate | ✅ Enabled | botnet, malicious, malware | 34/0 | - |
| Dan Pollock's List | ✅ Enabled | ads, malware, trackers | 0/113 | >95% overlap with StevenBlack Fake Gambling list |
| DandelionSprout-Anti-Malware-List | ✅ Enabled | malware | 32629/0 | - |
| Easy Privacy | ✅ Enabled | privacy, trackers | 0/949 | - |
| EasyList | ✅ Enabled | ads | 0/0 | 100% covered by other sources |
| FadeMind_2o7Net | ❌ Disabled | ads, privacy, trackers | - | No update since 2023-11-30 |
| FakeWebshopListHUN | ✅ Enabled | fake, phishing, scam, threat | 4735/2 | - |
| Frogeye-firstparty-trackers | ✅ Enabled | trackers | 10215/53 | - |
| GetAdmiral Domains Filter List | ✅ Enabled | ads, annoyance | 0/0 | - |
| GlobalAntiScamOrg-blocklist-domains | ✅ Enabled | scam | 7337/3 | - |
| HaGeZi Amazon Tracker | ✅ Enabled | privacy, trackers | 0/38 | >98% overlap with HaGeZi Pro |
| HaGeZi Apple Tracker | ✅ Enabled | privacy, trackers | 0/14 | >80% overlap with HaGeZi Pro |
| HaGeZi DNS TIF Mini | ✅ Enabled | malicious, threat | 3785/3 | 100% covered by other sources |
| HaGeZi Encrypted DNS Servers | ✅ Enabled | doh | 246/11 | - |
| HaGeZi Gambling Only Domains | ✅ Enabled | gambling | 174125/11 | Huge list and gambling-specific focus |
| HaGeZi Microsoft Tracker | ✅ Enabled | privacy, trackers | 0/36 | >75% overlap with HaGeZi Pro |
| HaGeZi Most Abused TLDs | ✅ Enabled | spam | 423/0 | - |
| HaGeZi Normal | ❌ Disabled | ads, malware, trackers | - | 100% overlap with HaGeZi Pro |
| HaGeZi Pro | ✅ Enabled | ads, malware, phishing, trackers | 8273/500 | - |
| HaGeZi Xiaomi Tracker | ✅ Enabled | privacy, trackers | 0/15 | >95% overlap with HaGeZi Pro |
| Hestat_Minerchk | ❌ Disabled | cryptocurrency | - | No update since 2018 |
| Hostsfile | ❌ Disabled | ads | - | No update since 2018-04-20 |
| Korlabs_UrlShortener | ✅ Enabled | url_shorteners | 0/44 | - |
| Malicious URL Blocklist (URLHaus) | ✅ Enabled | ads | 0/0 | 100% covered by other sources |
| Maltrail_StaticTrails | ✅ Enabled | malware, threat | 171748/5 | - |
| OISD Blocklist Big | ✅ Enabled | ads, cryptocurrency, malware, phishing, ransomware, trackers | 0/145 | Huge list |
| OISD Blocklist NSFW Small | ✅ Enabled | adult | 0/31 | - |
| OISD Blocklist Small | ✅ Enabled | ads, cryptocurrency, malware, phishing, ransomware, trackers | 0/95 | - |
| OpenPhish_Feed | ✅ Enabled | phishing | 135/3 | - |
| Peter Lowe's Blocklist | ✅ Enabled | ads | 0/214 | 100% covered by other sources |
| Policeman_SimpleDomainsBlocklist | ❌ Disabled | malicious | - | Archived on 2021-12-26 |
| PuppyScams | ✅ Enabled | fake, scam | 85/0 | List of top 100 pet scams is not being shared anymore, https://puppyscams.org/top-100-pet-scams |
| RPiList_specials-malware | ✅ Enabled | malware | 314861/0 | Huge list |
| RPiList_specials-phishing | ✅ Enabled | phishing | 479805/0 | Huge list |
| RedDragonWebDesign_block-everything | ✅ Enabled | ads, malicious, trackers | 648/0 | - |
| ShadowWhisperer's Dating List | ✅ Enabled | dating | - | - |
| ShadowWhisperer_BlockLists Ads | ✅ Enabled | ads | 0/160 | - |
| ShadowWhisperer_BlockLists Adult | ✅ Enabled | adult | 212242/27 | Huge list and adult-specific focus |
| ShadowWhisperer_BlockLists Malware | ✅ Enabled | malware | 2790/14 | - |
| ShadowWhisperer_BlockLists Scam | ✅ Enabled | scam | 7607/4 | - |
| ShadowWhisperer_UrlShortener | ✅ Enabled | url_shorteners | 1037/8 | - |
| Sinfonietta_Adult | ✅ Enabled | adult | 0/45 | - |
| Sinfonietta_Gambling | ✅ Enabled | gambling | 0/6 | - |
| Sinfonietta_Social | ✅ Enabled | social | 2752/178 | - |
| Spam404 | ✅ Enabled | spam | 5669/8 | - |
| Stamparm_Blackbook | ✅ Enabled | malicious, threat | 0/4 | >95% overlap with Blocklists UT1 Malware |
| StevenBlack_Adhoc_list | ❌ Disabled | ads, malware, trackers | - | 100% overlap with StevenBlack Fake Gambling list |
| StevenBlack_Fake_Gambling_Porn | ✅ Enabled | ads, adult, fake, fakenews, gambling | 0/640 | - |
| T145_black-mirror | ❌ Disabled | malicious, threat | - | Huge list, >8 million entries |
| Torrent Trackers | ✅ Enabled | torrent_trackers | 455/2 | - |
| URLHaus (Abuse.ch) | ✅ Enabled | malware | - | - |
| USOM-Blocklists-domains | ✅ Enabled | malicious, threat | 350800/40 | Huge list |
| Ukrainian Ad Filter | ✅ Enabled | ads | 1246/0 | - |
| Ukrainian Annoyance Filter | ✅ Enabled | annoyance | - | - |
| Ukrainian Privacy Filter | ✅ Enabled | privacy, trackers | 51/1 | - |
| Ukrainian Security Filter | ✅ Enabled | malicious, threat | 1174/0 | - |
| UncheckyAds | ❌ Disabled | ads, privacy, trackers | - | No update since 2021 |
| Viriback_Dump | ✅ Enabled | malware | 0/0 | - |
| WaLLy3K | ✅ Enabled | ads | 0/45 | - |
| WindowsSpyBlocker_Hosts_spy | ❌ Disabled | privacy, trackers | - | No update since 2022-05-16 |
| Winhelp2002 | ❌ Disabled | ads | - | No update since 2021-03-06 |
| YousList | ✅ Enabled | ads | 0/8 | - |
| YousList-AdGuard | ✅ Enabled | ads | 0/19 | - |
| Yoyo Adservers-Hosts | ✅ Enabled | ads | 0/214 | >95% overlap with StevenBlack Fake Gambling list |
| abpvn_hosts | ✅ Enabled | ads | 954/0 | - |
| anudeepND_adservers | ❌ Disabled | ads | - | No update since 2023-01-16 |
| bigdargon_hostsVN | ✅ Enabled | ads | 0/435 | - |
| cyberhost_malware-blocklist | ✅ Enabled | malware | 22/6 | - |
| fabriziosalmi_blocklists | ❌ Disabled | malicious, threat | - | Huge list, >3 million entries |
| hkamran80_smarttv | ✅ Enabled | smarttv | 0/31 | - |
| hufilter | ✅ Enabled | ads | 0/5 | >90% overlap with HaGeZi Pro |
| iam-py-test_my-filters-001-antitypo | ✅ Enabled | fake | 822/0 | - |
| jarelllama_Scam-Blocklist | ✅ Enabled | scam | 410232/21 | Disabled due to very large size (457K entries) - scam-specific focus |
| kadantiscam | ✅ Enabled | kad | 0/6 | peer-to-peer network protocol |
| malware-filter_phishing-filter | ✅ Enabled | malware, phishing | 0/0 | - |
| pexcn Torrent Trackers | ✅ Enabled | torrent_trackers | - | - |
| phishing_army | ✅ Enabled | phishing | 0/4 | - |
| quidsup_notrack-annoyance | ✅ Enabled | annoyance | 0/6 | >90% overlap with HaGeZi Pro |
| quidsup_notrack-malware | ✅ Enabled | malware | 0/4 | - |
| quidsup_notrack-tracker | ✅ Enabled | trackers | 0/366 | - |
| youtube_GoodbyeAds | ✅ Enabled | ads | 97220/35 | - |

</details>

<details>
<summary><strong>📄 sources_domain_new.json</strong> (1 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| nrd-14day-mini | ❌ Disabled | others | - | Huge list with low unique contribution |

</details>

<details>
<summary><strong>📄 sources_domain_top.json</strong> (1 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| tranco | ✅ Enabled | topdomains | 0/1152 | - |

</details>

<details>
<summary><strong>📄 sources_ip.json</strong> (41 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| AlienVault_Reputation | ❌ Disabled | malicious, threat | - | Not available anymore. The service has been discontinued. |
| BinaryDefense_Banlist | ✅ Enabled | malicious, threat | 0/0 | This is for public use only. |
| Blackhole_Today | ❌ Disabled | malicious, threat | - | Download fails frequently due to network instability or potential blocking. |
| BlockListDE_Brute | ✅ Enabled | threat | 0/0 | >95% overlap with Firehol_level2 |
| BlockListDE_Strong | ✅ Enabled | malicious, threat | 0/0 | >95% overlap with Borestad_AbuseIPDB_S100_3d |
| Borestad_AbuseIPDB_S100_3d | ✅ Enabled | malicious, threat | 0/0 | - |
| BruteforceBlocker | ✅ Enabled | threat | 0/0 | >95% overlap with EmergingThreats_CompromisedIPs |
| CINSScore_BadGuys_Army | ✅ Enabled | malicious, threat | 0/0 | - |
| DShield | ✅ Enabled | malicious, threat | 0/0 | 100% overlap with Firehol_level2/Firehol_level3 |
| DoH_IP_blocklists | ✅ Enabled | doh | 676/10 | >90% overlap with HaGeZi Encrypted DNS Servers |
| DoH_IP_list | ✅ Enabled | doh | 0/0 | - |
| ET_fwip | ✅ Enabled | malicious, threat | 133/0 | - |
| EmergingThreats_CompromisedIPs | ✅ Enabled | malicious, threat | 0/0 | >95% overlap with Firehol_level3,  and Borestad_AbuseIPDB_S100_3d |
| FabrizioSalmi_DNS | ✅ Enabled | dns | 0/0 | - |
| Firehol_BitcoinNodes_1d | ✅ Enabled | cryptocurrency | 7055/0 | - |
| Firehol_Botscout_1d | ✅ Enabled | malicious, threat | 467/0 | - |
| Firehol_CleanTalk | ✅ Enabled | malicious, threat | 415/0 | - |
| Firehol_CleanTalk_Top20 | ✅ Enabled | malicious, threat | 1/0 | - |
| Firehol_GPF_Comics | ✅ Enabled | malicious, threat | 1085/0 | - |
| Firehol_SSLProxies_1d | ✅ Enabled | anonymizer, privacy, proxy | 225/0 | - |
| Firehol_SocksProxy_7d | ✅ Enabled | anonymizer, privacy, proxy | 0/0 | - |
| Firehol_abusers_30d | ❌ Disabled | malicious, threat | - | False positives are common, use with caution. |
| Firehol_level1 | ✅ Enabled | malicious, threat | 3055/0 | - |
| Firehol_level2 | ✅ Enabled | malicious, threat | 0/0 | - |
| Firehol_level3 | ✅ Enabled | malicious, threat | 0/3 | - |
| GlobalAntiScamOrg-blocklist-ips | ✅ Enabled | scam | - | - |
| Greensnow | ✅ Enabled | malicious, malware, threat | 0/0 | >95% overlap with Firehol_level2 |
| HaGeZi_DoH | ✅ Enabled | doh | 0/0 | >90% overlap with DoH_IP_blocklists |
| HaGeZi_TIF | ✅ Enabled | malicious, threat | 0/0 | No unique contribution |
| MyIP_MS_Blocklist | ✅ Enabled | malicious, threat | - | - |
| Public_DNS4 | ✅ Enabled | dns | 61672/0 | - |
| Rutgers_DROP | ✅ Enabled | malicious, threat | 0/0 | - |
| Sblam_Blocklist | ✅ Enabled | spam | 949/0 | - |
| ScriptzTeam_BadIPS | ✅ Enabled | malicious, threat | 886/0 | - |
| Sentinel_Greylist | ✅ Enabled | malicious, threat | 0/1 | - |
| T145_allowlist-ips | ❌ Disabled | others | - | Huge list, use with caution. More than its blocklist counterpart. |
| T145_blocklist | ❌ Disabled | malicious, malware, threat | - | Huge list, use with caution. |
| URLHaus_Text | ✅ Enabled | malware | 0/0 | - |
| USOM-Blocklists-ips | ✅ Enabled | malicious, threat | 0/0 | - |
| Yoyo AdServers-IPList | ✅ Enabled | ads | 8899/0 | - |
| spamhaus_drop | ✅ Enabled | spam, threat | - | - |

</details>

<details>
<summary><strong>📄 sources_local.json</strong> (7 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| Local AI Allowlist (Domain) | ✅ Enabled | local | 0/51 | - |
| Local AI Blocklist (Domain) | ✅ Enabled | local | 0/51 | - |
| Local Allowlist (AdGuard) | ✅ Enabled | local | - | - |
| Local Allowlist (Domain) | ✅ Enabled | local | 6/8 | - |
| Local Allowlist (ipv4) | ✅ Enabled | local | 55/19 | - |
| Local Blocklist (AdGuard) | ✅ Enabled | local | 0/0 | - |
| Local Blocklist (Domain) | ✅ Enabled | local | 0/1 | - |

</details>

<details>
<summary><strong>📄 sources_mis.json</strong> (1 sources)</summary>

| Name | Status | Categories | Unique/Conflicts | Notes |
|------|--------|------------|------------------|-------|
| VXVault_URLList | ✅ Enabled | malware | 0/0 | >95% overlap with Firehol_level3 |

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
