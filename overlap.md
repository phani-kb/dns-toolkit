# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-10-31 01:24:50 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 162 |
| Total Entries Analyzed | 9.5M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 140 |
| allowlist | 22 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 32 |
| cidr_ipv4 | 3 |
| domain | 87 |
| ipv4 | 40 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 69 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 68.0% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 11.3K | 60.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 33.1K | 55.1% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 315 | 46.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 166 | 45.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 142 | 37.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 590 | 35.6% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 110 | 31.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 34 | 30.9% |
| HaGeZi Pro | blocklist | domain | 214.2K | 58.0K | 27.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 6.5K | 22.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 39.5K | 21.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 40.3K | 17.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 13.1K | 15.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 6.3K | 13.5% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 4.4K | 12.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 15.1K | 7.8% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 147 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 245 | 1.3% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 3.3K | 1.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 26 | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 458 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 144 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 42 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 432 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 462 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 85 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 42 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 739 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| phishing_army | blocklist | domain | 152.6K | 89 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 631 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 229 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 921 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 2 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 23 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 29 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 36 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 51 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.7K | 14.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 6.5K | 7.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 267 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 6 | 5.5% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 433 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 9 | 2.3% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 10 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 523 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 408 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 749 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| HaGeZi Pro | blocklist | domain | 214.2K | 1.7K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 89 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1.0K | 0.5% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 50 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 39 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 93 | Targets: 8 | Unique: 32 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 212 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 3 | 0.2% |
| EasyList | blocklist | adguard | 61.7K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 35 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.1K | 504 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 678 | 0.4% |
| EasyList | blocklist | adguard | 61.7K | 95 | 0.2% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 508 | 0.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 35 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 44 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 129 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 4 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 682 | Targets: 31 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 504 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 119 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 659 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 567 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 55 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 508 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 315 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 129 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 46 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 35 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 32 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 258.8K | Targets: 14 | Unique: 258.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 432 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 9 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 300.0K | Targets: 23 | Unique: 183.6K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 98.5K | 53.3% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2.3K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 585 | 3.8% |
| HaGeZi Pro | blocklist | domain | 214.2K | 7.9K | 3.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 916 | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 823 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 205 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 83 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1.2K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 2 | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 353 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 61 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 43 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 92 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 185.4K | Targets: 26 | Unique: 0 | Conflicts: 201</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 204 | 200 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.7% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 54.9K | 91.3% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 61.7K | 42.0K | 68.2% |
| Easy Privacy | blocklist | adguard | 55.6K | 28.7K | 51.6% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 678 | 50.7% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 59.8K | 25.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 21.6K | 11.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 2.4K | 6.8% |
| abpvn_hosts | blocklist | adguard | 1.0K | 23 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 125 | 0.7% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 12 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 178 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 184.7K | Targets: 68 | Unique: 0 | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 682 | 659 | 96.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.5K | 93.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 54.9K | 91.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.3% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 214.2K | 76.8K | 35.9% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 98.5K | 32.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5.0K | 27.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 59.8K | 25.5% |
| YousList | blocklist | hostname | 625 | 150 | 24.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 6.0K | 21.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39.5K | 19.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 1.7K | 11.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5.4K | 11.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 21.6K | 11.2% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 25 | 7.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2.4K | 6.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 5.0K | 5.8% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 20 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 553 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 125 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 276 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 345 | 0.4% |
| Torrent Trackers | blocklist | domain | 501 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 188 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 194 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 270 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 29 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 10 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 271 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 105 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 128 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 362 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 628 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 4 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 841 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 825 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 71 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 77 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 7 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 204 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 212 | 200 | 94.3% |

</details>

---

### AdGuardTeam_HttpsExclusions_android

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 97 | Targets: 10 | Unique: 71 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 10 | Unique: 4.0K | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 3 | 1.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_firefox

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 18 | Targets: 4 | Unique: 13 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_issues

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 6 | Unique: 59 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_mac

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 11 | Targets: 2 | Unique: 5 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |

</details>

---

### AdGuardTeam_HttpsExclusions_sensitive

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 182 | Targets: 12 | Unique: 155 | Conflicts: 15</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 11 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 1 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 91.8% |
| Easy Privacy | blocklist | adguard | 55.6K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 2.1K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 60 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 18 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 2 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 56 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.0K | 56.4% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 625 | 198 | 31.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 205 | 12.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 12 | 10.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 7.6K | 8.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 55 | 8.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 4.7K | 7.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.7% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1.8K | 6.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 21 | 6.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 18 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.5K | 4.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 7.1K | 3.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 5.0K | 2.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 6.4K | 2.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 937 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 996 | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 54 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Torrent Trackers | blocklist | domain | 501 | 1 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 18 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 43 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 10 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 28 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 142 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 35 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 12 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 820 | Targets: 21 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 291 | 18 | 6.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 103 | 2.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 248 | 1.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 174 | 1.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 25 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 495 | 1.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 7 | 1.3% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 121 | 1.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 290 | 1.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 34 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 5 | 0.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 163 | 0.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 502 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 152 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 62 | 0.3% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 6 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 2 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 710 | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 18.8K | 571 | 3.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 23 | 1.9% |
| Greensnow | blocklist | ipv4 | 3.5K | 66 | 1.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 203 | 1.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 10 | 0.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 70 | 0.7% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 2 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 6 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 51 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 107 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 202 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 16 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 1 | 0.2% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 12 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 31 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 38 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 291 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 116 | 6.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 153 | 4.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 66 | 2.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 18 | 2.2% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 228 | 1.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 4 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 230 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 271 | 0.5% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 2 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 46 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 2 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 7 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 17 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 9 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 23 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 9 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 2 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 39 | Unique: 10.0K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 76 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 45 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 190 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 22 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 137 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 190 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 34 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 53 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 8 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 17 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 264 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 40 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 126 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 50 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 152.6K | 96.8K | 63.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 253 | 56.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 67.8K | 35.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 69.9K | 29.8% |
| kadantiscam | blocklist | domain | 42.9K | 4.4K | 10.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 5.9K | 6.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1.0K | 4.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 3.1K | 3.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| HaGeZi Pro | blocklist | domain | 214.2K | 3.8K | 1.8% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 583 | 1.8% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.8K | 1.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 19.5K | 1.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 6.3K | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 354 | 1.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 3 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 196 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 16 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 25 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 470 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 105 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 43 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 12 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 32 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 2 | 0.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 65 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| phishing_army | blocklist | domain | 152.6K | 35 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 55 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 15 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 11 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 57.7K | Targets: 34 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 291 | 271 | 93.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.2K | 68.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 2.4K | 68.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.5K | 63.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 502 | 61.2% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 5.9K | 60.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 712 | 60.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.1K | 60.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 311 | 57.4% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 303 | 56.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 5.4K | 50.7% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 13.9K | 38.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 8.6K | 36.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 6.4K | 34.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 6.5K | 33.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 5.4K | 29.2% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 202 | 28.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 290 | 25.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 95 | 9.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 17 | 5.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 80 | 4 | 5.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 390 | 2.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 2 | 1.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 42 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 244 | 1.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 368 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 12 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 21 | Unique: 208 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 536 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 527 | 97.2% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 508 | 5.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 28 | 1.6% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 2 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 132 | 0.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 26 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 224 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 5 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 303 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 74 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 34 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 23 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 9 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 25 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 5 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 104 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 69 | 69 | 100.0% |
| Torrent Trackers | blocklist | domain | 501 | 103 | 20.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 22 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 9.7K | 6.1K | 63.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 248 | 30.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 5.9K | 24.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 4.4K | 22.8% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 5.9K | 16.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 9.5K | 16.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 1.7K | 16.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 632 | 12.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 2.0K | 10.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 94 | 2.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 9 | 1.7% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 9 | 1.7% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 225 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 14 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 4 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 74 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 1 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 28 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 28 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 31 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 7 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 54 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 48 | Unique: 11.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 1.8K | 5.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 30.8K | 5.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 23 | 5.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 23.4K | 1.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 5 | 1.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 72 | 1.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.3K | 1.5% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 337 | 1.4% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 476 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 168 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.1K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2.0K | 1.1% |
| phishing_army | blocklist | domain | 152.6K | 1.5K | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.1K | 0.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 126 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 267 | 0.6% |
| HaGeZi Pro | blocklist | domain | 214.2K | 958 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 173 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 271 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 458 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 46 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 28 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 188 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 440 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 614 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 35 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 17 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 82 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 13.1K | Targets: 54 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| YousList | blocklist | hostname | 625 | 108 | 17.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 12.9K | 15.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 46 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 423 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.1K | 5.7% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 363 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 8 | 1.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.5K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 582 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 280 | 0.8% |
| HaGeZi Pro | blocklist | domain | 214.2K | 1.4K | 0.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 537 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 553 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 95 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 168 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 635 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 15 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 30 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 691 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 15 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 20 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 83 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 36 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 9 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 14.0K | Targets: 7 | Unique: 14.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.3K | 11 | 0.8% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 445 | 2 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 1 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 1 | 0.0% |

</details>

---

### DandelionSprout_AdGuardHome_Whitelist

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 285 | Targets: 1 | Unique: 40 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| TogoFire_AD_Settings_whitelist | allowlist | adguard | 1.8K | 245 | 13.9% |

</details>

---

### DanMeUK_TorExitNodes

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 163 | 14.5% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 23 | 3.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 2 | 1.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 712 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 188 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 17 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 4 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 47 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 59 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 50 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 5 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 2 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 16 | 0.0% |

</details>

---

### Dogino_Discord_Official

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 4 | Unique: 8 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 14 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 7 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 7 | 0.2% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 9 | Unique: 319 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 95.9% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 80 | 10.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 91 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 1 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1.0K | 30.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 886 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 7 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 78 | 5.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 80 | 4.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 9 | Unique: 13.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 3.0K | 89.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 54 | 0.2% |
| EasyList | blocklist | adguard | 61.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 33 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 8 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 14 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 39 | Unique: 11.4K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3.0K | 89.0% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 886 | 78.7% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 54 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 135 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 10 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 15 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 63 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 33 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 539 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 8 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 5 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 4.9K | 26.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 5.1K | 21.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 103 | 12.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 859 | 8.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 2.6K | 7.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.1K | 5.3% |
| Greensnow | blocklist | ipv4 | 3.5K | 164 | 4.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 822 | 4.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 632 | 4.2% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 7 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 16 | 2.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 23 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 5 | 0.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 5 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 47 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 3 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 38 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.6K | Targets: 20 | Unique: 13.6K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.5% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 28.7K | 15.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.8% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 2.7K | 4.4% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 44 | 3.3% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 6.3K | 2.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 341 | 1.0% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 129 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 3 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 12 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 61.7K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.1K | 35.2K | 58.6% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 42.0K | 22.7% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 35.7K | 15.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 17.0K | 8.8% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 95 | 7.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 657 | 1.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 64 | 0.3% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 43 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 12 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 542 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 536 | 527 | 98.3% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 511 | 5.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 32 | 1.9% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 4 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 7 | 0.9% |
| Greensnow | blocklist | ipv4 | 3.5K | 25 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 128 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 232 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 311 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 73 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 34 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 27 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 24 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 10 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 7 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.8K | 100.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 34.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 5 | 0.0% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 41 | Unique: 1.2K | Conflicts: 213</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 14 | 32.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 13 | 31.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 500 | 123 | 24.6% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 56 | 7.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 3 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 42 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 7 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.8% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 19 | Unique: 4.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 8 | 8.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 38 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 51 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 15 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 23 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 47 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 23 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 110 | Targets: 8 | Unique: 95 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 3 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 2 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 2 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 9 | Unique: 470 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 2 | 1.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 9 | Unique: 2 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 1 | 0.1% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 982 | Targets: 22 | Unique: 660 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 60 | 2.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 2 | 1.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 6 | 0.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 95 | 0.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 6 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 20 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 16 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 13 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 43 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 9 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 4 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 8 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.7K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.1% |
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.1% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 18.8K | Targets: 31 | Unique: 0 | Conflicts: 814</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 4.9K | 95.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 3.0K | 85.5% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 571 | 80.4% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 228 | 78.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.3K | 73.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 8.5K | 44.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 9.8K | 41.0% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 132 | 24.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 128 | 23.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 163 | 19.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 1.8K | 16.9% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 5.6K | 15.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2.0K | 13.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 6.4K | 11.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 814 | 7.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 814 | 7.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 63 | 5.6% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 109 | 4.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 50 | 4.2% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 326 | 3.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 13 | 1.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 3 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 1 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 76 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 30 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 89 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.7K | Targets: 28 | Unique: 0 | Conflicts: 44</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 508 | 94.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 511 | 94.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 9.8K | 52.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.1K | 40.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 290 | 35.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 3.8K | 19.5% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 6.7K | 18.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 1.8K | 17.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 8.6K | 14.8% |
| Greensnow | blocklist | ipv4 | 3.5K | 232 | 6.7% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 17 | 5.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 31 | 4.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 40 | 2.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 20 | 2.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 326 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 128 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 104 | 0.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 41 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 41 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 177 | 0.1% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 13 | Unique: 2.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 72 | 21.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 60 | 6.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 10 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 1 | 0.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 7 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 11 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 6 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 335 | Targets: 10 | Unique: 223 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 2 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 2 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.8K | Targets: 19 | Unique: 5.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 314 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 89 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1.7K | 0.9% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 214.2K | 1.6K | 0.8% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 2.3K | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 18 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 348 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 54 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 131 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 15 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 117 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.7% |
| Easy Privacy | blocklist | adguard | 55.6K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 65 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 26 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 24 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.7K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 397 | 11.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 501 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 205 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1.5K | 0.8% |
| HaGeZi Pro | blocklist | domain | 214.2K | 1.6K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 401 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 590 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 188 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 65 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 26 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 24 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 4 | 0.0% |

</details>

---

### GlobalAntiScamOrg-blocklist-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.2K | Targets: 19 | Unique: 7.5K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.5K | Targets: 25 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 291 | 153 | 52.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 581 | 33.6% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 3.0K | 15.8% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 66 | 9.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 2.3K | 6.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 26 | 4.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 25 | 4.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 34 | 4.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 2.4K | 4.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 375 | 3.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 85 | 3.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 164 | 3.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 430 | 2.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 308 | 1.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 17 | 1.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 232 | 1.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 60 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 94 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 6 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 37 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 369 | Targets: 18 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 91 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 214.2K | 334 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 49 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 166 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 22 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 20 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 36 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 6 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 110 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 21 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 70 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 23 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 9 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 193.1K | Targets: 52 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 300 | 67.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 141.6K | 60.2% |
| phishing_army | blocklist | domain | 152.6K | 72.2K | 47.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 22.6K | 37.5% |
| kadantiscam | blocklist | domain | 42.9K | 14.2K | 33.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 63.7K | 29.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 67.8K | 28.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 9.9K | 21.3% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 4.8K | 19.6% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 129 | 18.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 14.8K | 17.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.5% |
| Spam404 | blocklist | domain | 8.1K | 1.1K | 13.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 21.6K | 11.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 15.1K | 7.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 373 | 7.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.5K | 7.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 996 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| YousList | blocklist | hostname | 625 | 29 | 4.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 537 | 4.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 291 | 3.9% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.2K | 3.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 736 | 2.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 2.0K | 2.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 71 | 2.0% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 24 | 1.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 137 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 3 | 1.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 130 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 9 | 0.7% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 206 | 0.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 2 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 3.1K | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.7K | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 47 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 39 | 0.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.0K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.8K | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.6K | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 99 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 116 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 9 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 193 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 92 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 4 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 193.1K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 235.0K | 141.6K | 60.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 930 | 50.3% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 71.9K | 47.4% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 22.6K | 37.5% |
| EasyList | blocklist | adguard | 61.7K | 17.0K | 27.6% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 78.0K | 11.7% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 21.6K | 11.6% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 129 | 9.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1.2K | 3.4% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 24 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 9 | 0.7% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 18 | 0.7% |
| abpvn_hosts | blocklist | adguard | 1.0K | 7 | 0.7% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 7 | 0.5% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.6K | 129 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.4K | Targets: 5 | Unique: 352 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.3% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.4K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.3% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 41 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 624.6K | Targets: 43 | Unique: 604.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2.7K | 7.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 3.9K | 4.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 1.6K | 0.8% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 194 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1.4K | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.1K | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 440 | 0.5% |
| HaGeZi Pro | blocklist | domain | 214.2K | 928 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 3 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 12 | 0.3% |
| kadantiscam | blocklist | domain | 42.9K | 117 | 0.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1.4K | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 470 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 35 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 44 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 229 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.5K | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 6 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 128 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 10 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 24 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 84 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 34 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 46 | 0.1% |
| phishing_army | blocklist | domain | 152.6K | 167 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 10 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 30 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 383 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 46 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 10 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 331 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 24 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 66 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 40 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 20 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 25 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 57 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 10 | 0.0% |

</details>

---

### HaGeZi Most Abused TLDs

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 445 | Targets: 1 | Unique: 443 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 2 | 0.0% |

</details>

---

### HaGeZi Pro

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 214.2K | Targets: 69 | Unique: 0 | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 347 | 99.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 98.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 57.1K | 95.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 334 | 90.5% |
| hufilter | blocklist | hostname | 92 | 80 | 87.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 331 | 86.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 300 | 85.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.5% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 567 | 83.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 344 | 77.1% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 70 | 63.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 77 | 61.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 76.8K | 41.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.1K | 38.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 111 | 37.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 87.1K | 37.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 63.7K | 33.0% |
| YousList | blocklist | hostname | 625 | 200 | 32.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 8.7K | 30.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 58.0K | 28.6% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 25.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 11.6K | 24.9% |
| WaLLy3K | blocklist | domain | 351 | 82 | 23.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.7K | 17.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 13.7K | 15.9% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 4.8K | 13.7% |
| kadantiscam | blocklist | domain | 42.9K | 5.9K | 13.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 1.6K | 11.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.4K | 10.5% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 2.4K | 10.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| phishing_army | blocklist | domain | 152.6K | 8.6K | 5.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 248 | 3.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 135 | 2.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 465 | 2.6% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 7.9K | 2.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 64 | 1.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 64 | 1.7% |
| Spam404 | blocklist | domain | 8.1K | 141 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.8K | 1.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 261 | 1.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 41 | 1.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 27 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 958 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 531 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 628 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 135 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 2 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 51 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.7K | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 128 | 0.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 4 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.5K | 0.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1.8K | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.2K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 486 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 928 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 41 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 350 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 214.2K | 347 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 87 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 16 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 9 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 69.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 78 | 10.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 94 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 12 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 36.1K | Targets: 35 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 12.9K | 79.9% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 230 | 79.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.2K | 69.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 2.3K | 66.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 495 | 60.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.6K | 50.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 232 | 42.8% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 224 | 41.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.9K | 39.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 3.5K | 33.4% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 5.6K | 29.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 6.7K | 28.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 4.9K | 26.3% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 2.5K | 25.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 13.9K | 24.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 4.4K | 22.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 215 | 19.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 188 | 15.9% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 107 | 15.1% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 117 | 4.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 43 | 4.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 320 | 2.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 2 | 1.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 3 | 0.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 14 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 35 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 35 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 538 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 6 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 23 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 111 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 53 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 20 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 108 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 17 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 13 | 0.0% |

</details>

---

### hufilter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 92 | Targets: 24 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 84 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 80 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 70 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 12 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 16 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 82 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 1 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 56 | Unique: 420.5K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 995 | 13.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1.4K | 5.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 9 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2.7K | 1.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2.7K | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.9K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| phishing_army | blocklist | domain | 152.6K | 1.3K | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 614 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 98 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.5K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.2K | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 2.4K | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 182 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 121 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 2.1K | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 180 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 54 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 270 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 20 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 42 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 211 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 12 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 42.9K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 36.4K | 42.3% |
| phishing_army | blocklist | domain | 152.6K | 18.7K | 12.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 14.2K | 7.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 12.9K | 5.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 5.9K | 2.7% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 556 | 2.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4.4K | 1.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 2 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 30 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 12 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 65 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 68 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 117 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 24 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 47 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 21 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 29 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 85 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 63 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 42 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 5 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 30 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 175 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 126 | 2.1% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 2 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 47 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 17 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 29 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 4 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 24 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 185.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 3 | 0.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 49 | Targets: 1 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | ipv4_from_domain | 49 | 49 | 100.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 24 | Targets: 5 | Unique: 0 | Conflicts: 26</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### local_ai_blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 49 | Targets: 1 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | ipv4_from_domain | 49 | 49 | 100.0% |

</details>

---

### local_ai_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 24 | Targets: 5 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### local_domain_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7 | Targets: 21 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 9 | Unique: 0 | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### local_social_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1 | Targets: 4 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### local_source_domain_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 42 | Targets: 2 | Unique: 27 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 13 | 0.8% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |

</details>

---

### local_source_ipv4_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 62 | Targets: 2 | Unique: 58 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 9.7K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 456 | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 930 | 0.5% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 1.3K | 0.5% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 1.8K | 0.3% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 3 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3M | Targets: 60 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.1M | 98.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 272.6K | 49.5% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 12.1K | 37.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 23.4K | 24.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 42 | 9.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 19.5K | 8.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 3.7K | 8.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 11 | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2.0K | 5.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 691 | 5.3% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 539 | 3.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 5.8K | 3.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 6.4K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 214.2K | 4.5K | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1.2K | 1.4% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 8 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 288 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 402 | 0.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 26 | 0.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.2K | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 90 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| phishing_army | blocklist | domain | 152.6K | 499 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 628 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.5K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 125 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 7 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 18 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 37 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 63 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 219.8K | Targets: 37 | Unique: 207.8K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 5.4K | 34.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 38 | 5.4% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 9 | 3.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 276 | 1.7% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 538 | 1.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 8 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 16 | 1.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 29 | 1.2% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 6 | 1.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 12 | 0.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 177 | 0.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 38 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 6 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 7 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 8 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 139 | 0.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 11 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 61 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 89 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 17 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 69 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 32 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 5 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 31 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 8 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 8 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 48 | 0.1% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1M | Targets: 48 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.1M | 80.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 222.5K | 40.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 16.3K | 16.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 33 | 7.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2.2K | 6.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 847 | 2.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 4.0K | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 4.3K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 985 | 1.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2.2K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.5K | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 631 | 0.3% |
| phishing_army | blocklist | domain | 152.6K | 384 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 85 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.2K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 106 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 362 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 56 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 11 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 69 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 9 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 24.4K | Targets: 32 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 152.6K | 20.4K | 13.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 18.0K | 7.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 15 | 5.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 17 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 4.8K | 2.5% |
| kadantiscam | blocklist | domain | 42.9K | 556 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 367 | 1.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2.4K | 1.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 21 | 0.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.0K | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 337 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 118 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 32 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 85 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 12 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 106 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 24 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 90 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 235.0K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.1K | 59.4K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 74.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 141.6K | 73.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1.3K | 68.2% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 825 | 61.7% |
| EasyList | blocklist | adguard | 61.7K | 35.7K | 57.8% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 79.5K | 52.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 508 | 38.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 59.8K | 32.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 66 | 17.9% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 79.2K | 11.9% |
| Easy Privacy | blocklist | adguard | 55.6K | 6.3K | 11.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 3.9K | 11.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 36 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 54 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 124 | 0.7% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 33 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 235.0K | Targets: 64 | Unique: 0 | Conflicts: 21</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 59.4K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 98.7% |
| hufilter | blocklist | hostname | 92 | 82 | 89.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 382 | 85.7% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 508 | 74.5% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 18.0K | 73.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 141.6K | 73.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.7% |
| phishing_army | blocklist | domain | 152.6K | 80.7K | 52.9% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 268 | 42.9% |
| HaGeZi Pro | blocklist | domain | 214.2K | 87.1K | 40.7% |
| WaLLy3K | blocklist | domain | 351 | 135 | 38.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.4K | 34.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 59.8K | 32.4% |
| kadantiscam | blocklist | domain | 42.9K | 12.9K | 30.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 69.9K | 29.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 22.1K | 25.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 87 | 24.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 10.5K | 22.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 23 | 20.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.8K | 20.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 5.9K | 20.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 40.3K | 19.8% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 57 | 14.9% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 557 | 11.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 3.9K | 11.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 36 | 9.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 7 | 2.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 2.1K | 2.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 85 | 1.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 117 | 0.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 124 | 0.7% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 217 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.9K | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 395 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 463 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.4K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.3K | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 1.2K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 2.4K | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 9 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.4K | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 33 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 305 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 4 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.9K | 9.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.4K | 8.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 12.5K | 5.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 27 | 2.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 21 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 47 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 58 | 0.2% |
| Torrent Trackers | blocklist | domain | 501 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 245 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 76 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 124 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 125 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 261 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 21 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 46 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 22 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 18 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 40 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 10 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 22 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 18.5K | Targets: 10 | Unique: 18.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 27 | 2.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 58 | 0.2% |
| EasyList | blocklist | adguard | 61.7K | 64 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 125 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 124 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 76 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 14 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 60.1K | Targets: 63 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 84 | 91.3% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 504 | 73.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.6% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 54.9K | 29.7% |
| HaGeZi Pro | blocklist | domain | 214.2K | 57.1K | 26.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 59.4K | 25.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4.7K | 25.3% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33.1K | 16.3% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5.7K | 12.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 22.6K | 11.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 3.3K | 11.6% |
| Adaway | blocklist | hostname | 6.5K | 749 | 11.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 25 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 22 | 6.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 917 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 582 | 4.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.5K | 4.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 3.6K | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 65 | 3.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 76 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 215 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 60 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 76 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 261 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 916 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 173 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 10 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 24 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 156 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 75 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 84 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 43 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 106 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 402 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 182 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 60.1K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 841 | 62.9% |
| EasyList | blocklist | adguard | 61.7K | 35.2K | 57.1% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 504 | 37.7% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 54.9K | 29.6% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 59.4K | 25.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 22.6K | 11.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| Easy Privacy | blocklist | adguard | 55.6K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1.5K | 4.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 65 | 3.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 29 | 2.8% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 76 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 3 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 3 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 67 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 291 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| phishing_army | blocklist | domain | 152.6K | 280 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 15 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 8 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 5 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 69 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 104 | 69 | 66.3% |
| Torrent Trackers | blocklist | domain | 501 | 68 | 13.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 34.9K | Targets: 24 | Unique: 22.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 456 | 24.7% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 35 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 1.5K | 2.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 3.9K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 26 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 2.4K | 1.3% |
| EasyList | blocklist | adguard | 61.7K | 657 | 1.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.6K | 341 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 1.2K | 0.6% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 538 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 58 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 54 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 1.6K | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 34.9K | Targets: 76 | Unique: 0 | Conflicts: 166</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 698 | 19.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 75 | 16.8% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 47 | 9.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.5K | 8.3% |
| Adaway | blocklist | hostname | 6.5K | 523 | 8.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 2.0K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 35 | 5.1% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 8 | 2.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1.5K | 2.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 4.8K | 2.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 280 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 564 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1.6K | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 3.9K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 26 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 367 | 1.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 65 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2.4K | 1.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 55 | 0.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 4.3K | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 1.2K | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 476 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 192 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 2.7K | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| phishing_army | blocklist | domain | 152.6K | 538 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 58 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 54 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.0K | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 65 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 354 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 22 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 847 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 7 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 61 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 121 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 77 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 152.6K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 280 | 96.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 20.4K | 83.5% |
| kadantiscam | blocklist | domain | 42.9K | 18.7K | 43.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 96.8K | 40.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 72.2K | 37.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 80.7K | 34.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 16.4K | 19.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 24 | 4.8% |
| HaGeZi Pro | blocklist | domain | 214.2K | 8.6K | 4.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 538 | 1.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1.5K | 1.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 35 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.3K | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 1 | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 43 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 512 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 12 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 12 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 89 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 499 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 167 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 384 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 19 | Unique: 61.7K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 94 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 91 | 4.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 48 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 6 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 12 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 290 | 17.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 76 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 46 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 296 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 145 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 47 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 300 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 298 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 27 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 59 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 13 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 28 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 35 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 77 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 47 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 22 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 20 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 28 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 8 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 55 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 577 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 34 | 8.9% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 8 | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 433 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.3K | 3.7% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 906 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 314 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 12 | 1.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 6 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 917 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1.2K | 1.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 2.7K | 1.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 991 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2.1K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 104 | 1 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1.4K | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Torrent Trackers | blocklist | domain | 501 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 180 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 585 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 4 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 130 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 22 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 25 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 1 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 61.7K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 666.4K | Targets: 15 | Unique: 381.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1.8K | 99.5% |
| RPiList_specials-phishing | blocklist | adguard | 151.7K | 123.5K | 81.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 78.0K | 40.4% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 79.2K | 33.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1.6K | 4.7% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 14 | 0.1% |
| EasyList | blocklist | adguard | 61.7K | 43 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 178 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 14 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 67 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 3 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 151.7K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 71.9K | 37.3% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 79.5K | 33.9% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 123.5K | 18.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 538 | 1.5% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 4 | 0.2% |
| EasyList | blocklist | adguard | 61.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 12 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.7K | Targets: 24 | Unique: 0 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 291 | 116 | 39.9% |
| Greensnow | blocklist | ipv4 | 3.5K | 581 | 16.7% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 1.3K | 6.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 32 | 5.9% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 28 | 5.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 381 | 3.6% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 1.2K | 3.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 25 | 3.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1.2K | 2.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 25 | 1.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 199 | 1.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 117 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 23 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 31 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 40 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 14 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 19 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 14 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 19 | Unique: 272 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 163 | 13.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 110 | 3 | 2.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 14 | 1.4% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 10 | 1.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 4 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 215 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 290 | 0.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 10 | 0.4% |
| Greensnow | blocklist | ipv4 | 3.5K | 14 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 63 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 31 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 4 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 20 | Unique: 2.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 291 | 66 | 22.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 85 | 2.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 25 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 109 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 2 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 2 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 117 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 8 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 1 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 7 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 8 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 1 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.5K | Targets: 20 | Unique: 0 | Conflicts: 12.7K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 814 | 4.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 33 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 14 | 0.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 86 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 4 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 41 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 28 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 19 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 22 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 35 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 2 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.6K | Targets: 27 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 381 | 22.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 174 | 21.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 859 | 16.8% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 46 | 15.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.7K | 11.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 375 | 10.8% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 70 | 9.9% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 3.5K | 9.8% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 1.8K | 9.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.4K | 9.3% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 908 | 9.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 1.7K | 9.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 1.8K | 7.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 1.3K | 6.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 34 | 6.3% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 34 | 6.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 205 | 1.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 5 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 71 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 19 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 19 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 32 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.4K | Targets: 18 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 27 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 29 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 15 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 27 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 10 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 718 | Targets: 38 | Unique: 378 | Conflicts: 267</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 7 | 1 | 14.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 56 | 3.3% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 7 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 2 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 2 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 27 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 31 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 11 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.6K | Targets: 55 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 501 | 30.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 974 | 27.4% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 119 | 17.4% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 87 | 13.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.8K | 9.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 24 | 6.3% |
| Adaway | blocklist | hostname | 6.5K | 408 | 6.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 22 | 6.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 906 | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2.0K | 5.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3.3K | 5.5% |
| HaGeZi Pro | blocklist | domain | 214.2K | 8.7K | 4.1% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 6.0K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.5K | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 363 | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 5.9K | 2.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 2.0K | 2.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 736 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 47 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 126 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 166 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 205 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 15 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 17 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 11 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 36 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 34 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 54 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 288 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 236.1K | Targets: 33 | Unique: 179.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 12.5K | 67.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.6K | 28.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 462 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 486 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 77 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 193 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 194 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 82 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 107 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 10 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 75 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 305 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 156 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 125 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 69 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 21 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 211 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 29 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 30 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 6 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.7K | Targets: 44 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 5.7K | 9.4% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 46 | 6.7% |
| HaGeZi Pro | blocklist | domain | 214.2K | 11.6K | 5.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 9.9K | 5.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 937 | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 10.5K | 4.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.3K | 3.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 5.4K | 2.9% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 55 | 1.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 180 | 1.2% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 95 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 192 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 308 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3.7K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 267 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 68 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 196 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 371 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 46 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 12 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 12 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.5K | Targets: 31 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 995 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 291 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 147 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 248 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 7 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 85 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 13 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 6 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 25 | Unique: 1.4K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 2 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 55 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 16 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 7 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 24 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 49 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 9 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 26 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 5.9K | 31.6% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 19.4K | 8.2% |
| pexcn Torrent Trackers | blocklist | domain_url | 69 | 2 | 2.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 104 | 2 | 1.9% |
| Torrent Trackers | blocklist | domain | 501 | 9 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 615 | 0.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 215 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 126 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 12 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 78 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 531 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 395 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 99 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 276 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 44 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 37 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 8 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 39 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 17 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 2.7K | 3.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.2K | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 27 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 34 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 15 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 32 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 64 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 35 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 29 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 31 | Unique: 4.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 1.6K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 1.1K | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 20 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 41 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 22 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 141 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 52 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 110 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 18 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 5 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.8K | 98.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.6% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 27 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 17.6K | 7.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 4.3K | 1.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 2.6K | 1.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.0K | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 2 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 465 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 389 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 126 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 20 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 10 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 86.0K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 12.9K | 98.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 98.4% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 42.9K | 36.4K | 85.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.6K | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 401 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 66 | 17.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 49 | 13.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 899 | 12.1% |
| phishing_army | blocklist | domain | 152.6K | 16.4K | 10.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 22.1K | 9.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 14.8K | 7.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 2.0K | 6.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.1K | 6.5% |
| HaGeZi Pro | blocklist | domain | 214.2K | 13.7K | 6.4% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3.6K | 6.0% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 32 | 4.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 16 | 4.6% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.6K | 4.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 5.0K | 2.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 5.9K | 2.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 348 | 2.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 308 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 3.9K | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 118 | 0.5% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 823 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 271 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 40 | 0.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 1 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 40 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 7 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 985 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 107 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 11 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 9 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 158 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 180 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 47 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 6.4K | 34.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 21.6K | 9.1% |
| pexcn Torrent Trackers | blocklist | domain_url | 69 | 2 | 2.9% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 104 | 2 | 1.9% |
| Torrent Trackers | blocklist | domain | 501 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 703 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 157 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 166 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 261 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 95 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 214.2K | 628 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 345 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 13 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 463 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 116 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 50 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 37 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 43 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 20 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 48 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 34 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 15 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 38 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 32 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 17 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 64 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 29 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 32.0K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 34 | 7.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 24.4K | 4.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1.8K | 1.9% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 564 | 1.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12.1K | 0.9% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 2 | 0.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 8 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.2K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 583 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 32 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 217 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 128 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 206 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 144 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 11 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 43 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 194 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 550.4K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 24.4K | 76.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 446 | 236 | 52.9% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 30.8K | 31.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 272.6K | 20.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 222.5K | 20.5% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 4.3K | 12.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6.3K | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 3.1K | 1.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.4K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 6 | 0.9% |
| HaGeZi Pro | blocklist | domain | 214.2K | 1.8K | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 371 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 24 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 106 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 739 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| phishing_army | blocklist | domain | 152.6K | 512 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 156 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 158 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 13 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1.4K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 156 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 47 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 36 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 21 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 271 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 23 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 19.2K | Targets: 28 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 18.8K | 8.5K | 45.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.4K | 29.2% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 203 | 28.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 152 | 18.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 1.7K | 16.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 822 | 16.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 3.8K | 15.8% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 1.5K | 15.6% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 74 | 13.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 73 | 13.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 430 | 12.4% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 4.4K | 12.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 199 | 11.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 6.5K | 11.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 1.9K | 10.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 59 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 31 | 2.8% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 7 | 2.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 8 | 0.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 53 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 7 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 17 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 139 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### TogoFire_AD_Settings_whitelist

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 1.8K | Targets: 1 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout_AdGuardHome_Whitelist | allowlist | adguard | 285 | 245 | 86.0% |

</details>

---

### Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 501 | Targets: 8 | Unique: 307 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 104 | 103 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 69 | 68 | 98.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 44 | Unique: 0 | Conflicts: 528</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 123 | 7.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 2 | 4.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 19 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 82 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 37 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 36 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 19 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 27 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 32 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 14 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 11 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 31 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 61.7K | 51 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 666.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 30 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 11 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 11 | Unique: 24 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.6K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 852 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 43 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 53 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 66 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 7 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 446 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | domain | 34.9K | 75 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 382 | 0.2% |
| HaGeZi Pro | blocklist | domain | 214.2K | 344 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 300 | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 34 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 253 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 33 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 42 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 23 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 236 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.2K | Targets: 25 | Unique: 2.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 12.9K | 35.7% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 205 | 1.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 10 | 1.8% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 9 | 1.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 96 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 4 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 128 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 244 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 1 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 2 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 9 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 2 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 30 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 18 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 13 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 17 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 276 | 0.1% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.9K | Targets: 35 | Unique: 8.6K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 282 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 47 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 291 | 9 | 3.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5.4K | 2.5% |
| BlockListDE_Brute | blocklist | ipv4 | 710 | 13 | 1.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 982 | 16 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 542 | 7 | 1.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 13 | 1.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 37 | 1.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 19 | 1.1% |
| BruteforceBlocker | blocklist | ipv4_find | 536 | 5 | 0.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 47 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 320 | 0.9% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 79 | 0.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 820 | 6 | 0.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.6K | 71 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 390 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 96 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 74 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 69 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 104 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 76 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 53 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.1K | 0.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 3.9K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 557 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 373 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 442 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 72 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 135 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 8 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 14 | Unique: 355 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 282 | 1.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.6K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.8K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 14 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.7K | 3 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 19.2K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 8 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.8K | 2 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 9.7K | 45 | 0.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.1K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 23 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 351 | Targets: 32 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| YousList | blocklist | hostname | 625 | 9 | 1.4% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 18 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 83 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 54 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 85 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 135 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 19 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 20 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 82 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 35 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 65 | Unique: 0 | Conflicts: 90</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.5% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.2K | 33.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 20.7K | 24.1% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 188 | 11.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 40 | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 15.8K | 6.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 991 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 1.8K | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 3.1K | 5.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 1.4K | 4.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 25 | 3.7% |
| HaGeZi Pro | blocklist | domain | 214.2K | 6.8K | 3.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 5.5K | 2.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 131 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 161 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.5K | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 977 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 353 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 15 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 41 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 44 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 51 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 23 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### YousList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 625 | Targets: 33 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 9 | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 198 | 1.1% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 3 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 241 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 87 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 94 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 214.2K | 200 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 150 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 42 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 268 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 17 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 29 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| AdGuard DNS filter | blocklist | adguard | 185.4K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 61.7K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 193.1K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 235.0K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.1K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.9K | 2 | 0.0% |

</details>

---

### youtube_GoodbyeAds

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 22 | Unique: 97.2K | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| WaLLy3K | blocklist | domain | 351 | 7 | 2.0% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 50 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 52 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 74 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 8 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 45 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| HaGeZi Pro | blocklist | domain | 214.2K | 41 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 61 | Unique: 0 | Conflicts: 43</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 397 | 23.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.0K | 10.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 18 | 5.1% |
| Adaway | blocklist | hostname | 6.5K | 267 | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 86.0K | 3.5K | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 577 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 4 | 3.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.6K | 974 | 3.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 423 | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.1K | 1.6K | 2.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.9K | 698 | 2.0% |
| AdGuard Base filter | blocklist | domain_adguard | 682 | 13 | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 5 | 1.4% |
| HaGeZi Pro | blocklist | domain | 214.2K | 3.0K | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 235.0K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.7K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 55 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 18 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.4K | 13 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 624.6K | 12 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.1K | 5 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 300.0K | 30 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 193.1K | 71 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.4K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 258.8K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.6K | 7 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 7 | 0.0% |
| phishing_army | blocklist | domain | 152.6K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 32.0K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.9K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

