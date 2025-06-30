# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-05-06 11:49:16 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 167 |
| Total Entries Analyzed | 8.3M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 144 |
| allowlist | 23 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 33 |
| cidr_ipv4 | 3 |
| domain | 89 |
| ipv4 | 42 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 186.1K | Targets: 72 | Unique: 0 | Conflicts: 125</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 5.1K | 78.4% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 430 | 68.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 68.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 37.7K | 66.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 291 | 63.5% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 11.0K | 62.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 376 | 54.1% |
| WaLLy3K | blocklist | domain | 350 | 177 | 50.6% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 185 | 46.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 632 | 38.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 173 | 36.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 106 | 36.1% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 38 | 34.5% |
| hufilter | blocklist | hostname | 96 | 33 | 34.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 33.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 48.2K | 29.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 114.8K | 26.2% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 4.7K | 26.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 4.0K | 25.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 6.3K | 23.5% |
| quidsup_notrack-malware | blocklist | domain | 138 | 32 | 23.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2.5K | 20.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 6.6K | 15.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 13.6K | 15.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 11.3K | 14.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 46.0K | 14.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 3.2K | 10.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 402 | 8.1% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.2K | 6.8% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 3.4K | 3.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 71 | 3.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 347 | 2.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 15 | 2.2% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 53 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 53 | 1.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 3 | 1.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 196 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 16 | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 26 | 0.8% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 501 | 0.8% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 602 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 872 | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 19 | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 203 | 0.5% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 8 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 395 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 28 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 1.3K | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 55 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1.9K | 0.2% |
| kadantiscam | blocklist | domain | 48.2K | 101 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 34 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 411 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 444 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 131 | 0.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| phishing_army | blocklist | domain | 144.8K | 73 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 339 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 267 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 40 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 27 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.2K | Targets: 6 | Unique: 1.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 29 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 997 | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 23 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 33 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 49 | Unique: 0 | Conflicts: 82</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 2.7K | 15.4% |
| WaLLy3K | blocklist | domain | 350 | 54 | 15.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 258 | 7.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 6.5K | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1.1K | 6.0% |
| hufilter | blocklist | hostname | 96 | 5 | 5.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 22 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 404 | 3.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 22 | 3.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 14 | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 444 | 2.9% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 5.1K | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 49 | 2.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 9 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 11 | 1.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 407 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 739 | 1.3% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.4K | 1.0% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 4 | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2.3K | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1.0K | 0.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 105 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 10 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 26 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 15 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 50 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 87 | Targets: 4 | Unique: 64 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 197 | 1 | 0.5% |
| EasyList | blocklist | adguard | 75.8K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 2 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 997 | Targets: 12 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 299 | 0.5% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 402 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.2K | 9 | 0.1% |
| Easy Privacy | blocklist | adguard | 54.6K | 38 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 12 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 301 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.2K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 4 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 9 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 3 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 400 | Targets: 29 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 299 | 0.5% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 79 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 12 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 40 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 385 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 185 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 24 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 367 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 10 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 301 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 12 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 6 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 28 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 24 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 98.6K | Targets: 16 | Unique: 98.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 395 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 10 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 5 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 100.1K | Targets: 26 | Unique: 0 | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 69.0K | 41.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 7.1K | 22.6% |
| hufilter | blocklist | hostname | 96 | 20 | 20.8% |
| HaGeZi Pro | blocklist | domain | 438.7K | 19.6K | 4.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 656 | 4.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 3.4K | 1.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 953 | 1.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 205 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 726 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 29 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 122 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 65 | 0.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.3K | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 2 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 260 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 40 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 165.2K | Targets: 69 | Unique: 0 | Conflicts: 73</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 400 | 385 | 96.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.5K | 91.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 51.8K | 91.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 403 | 88.0% |
| hufilter | blocklist | hostname | 96 | 74 | 77.1% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 69.0K | 69.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.2% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 5.0K | 28.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 48.2K | 25.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 6.7K | 25.0% |
| YousList | blocklist | hostname | 625 | 152 | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 138 | 33 | 23.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 985 | 23.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 86.9K | 19.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 5.5K | 17.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 7.4K | 17.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 56.0K | 17.1% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2.3K | 14.8% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2.2K | 12.4% |
| WaLLy3K | blocklist | domain | 350 | 35 | 10.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 30 | 6.4% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 4.8K | 5.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 540 | 4.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 21 | 3.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 11 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 31 | 1.4% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 28 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 28 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 10 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 121 | 0.7% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 349 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 279 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 365 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 12 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 54 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 84 | 0.2% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 11 | 0.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 25 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 190 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 312 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 339 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 780 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 30 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 5 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 191 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 138 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 4 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 97 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 10 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 165.9K | Targets: 23 | Unique: 0 | Conflicts: 186</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 189 | 185 | 97.9% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 51.8K | 91.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 90.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.5% |
| EasyList | blocklist | adguard | 75.8K | 56.0K | 73.9% |
| Easy Privacy | blocklist | adguard | 54.6K | 27.2K | 49.8% |
| AdGuard Base filter | blocklist | adguard | 997 | 402 | 40.3% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 56.0K | 17.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 51 | 13.9% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 2.2K | 12.4% |
| abpvn_hosts | blocklist | adguard | 1.2K | 23 | 2.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 28 | 1.9% |
| AdBlockID | allowlist | adguard | 87 | 1 | 1.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 10 | 0.8% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 121 | 0.7% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 16 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 365 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 2 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 289 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 11 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 850 | 1.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 5 | 1.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 174 | 1.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 1.2K | 0.7% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 826 | 0.3% |
| AdGuard Base filter | blocklist | adguard | 997 | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 54.6K | 73 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 1 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 189 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 197 | 185 | 93.9% |

</details>

---

### AdGuardTeam_HttpsExclusions_android

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 97 | Targets: 11 | Unique: 67 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 5 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 5 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 11 | Unique: 3.9K | Conflicts: 21</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 3 | 1.7% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_firefox

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 18 | Targets: 3 | Unique: 14 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |

</details>

---

### AdGuardTeam_HttpsExclusions_issues

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 5 | Unique: 61 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 2 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_mac

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 11 | Targets: 3 | Unique: 4 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_sensitive

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 180 | Targets: 11 | Unique: 154 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 2 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 11 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 1 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 4 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 75.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 1 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.6K | Targets: 57 | Unique: 0 | Conflicts: 85</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.0K | 57.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.9% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 625 | 199 | 31.8% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 87 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 101 | 14.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 15 | 13.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 212 | 12.8% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2.0K | 11.3% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 40 | 10.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 28 | 9.5% |
| quidsup_notrack-malware | blocklist | domain | 138 | 13 | 9.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 42 | 8.9% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 1.1K | 8.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 7.6K | 8.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 4.7K | 8.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.2K | 8.0% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1.6K | 6.2% |
| hufilter | blocklist | hostname | 96 | 6 | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 11.0K | 5.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 5.0K | 3.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 11.4K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 883 | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 6.4K | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 13 | 1.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 34 | 1.5% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 32 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 32 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 154 | 0.2% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 68 | 0.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 34 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 138 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 20 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 19 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 40 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 32 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 28 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 9 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 20 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 97 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 144 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 8 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.1K | Targets: 25 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 97 | 24.9% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 94 | 23.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 994 | 19.4% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 8.1K | 13.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 1.3K | 12.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.7K | 11.2% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 1.2K | 10.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 2.2K | 9.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 1.2K | 7.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 1.1K | 6.8% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 46 | 6.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 68 | 4.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 3.8K | 4.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 82 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 322 | 2.2% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 43 | 2.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 656 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 45 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 34 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 5 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 20 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 20 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 27 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 3.5K | 554 | 15.6% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1.8K | 12.5% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 36 | 9.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 31 | 7.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 621 | 3.9% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 1.5K | 2.5% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 17 | 2.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 14 | 1.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 94 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 8 | 0.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 116 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 501 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 59 | 0.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 14 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 43 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 7 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 12 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 3 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 28 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 10 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 4 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 764 | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 78 | 5.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 96 | 3.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 115 | 3.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 5 | 1.3% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 187 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 3 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 17 | 0.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 46 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 251 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 236 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 14 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 41 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 21 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 28 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 17 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 40 | Unique: 15.0K | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 138 | 3 | 2.2% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 24 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 45 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 19 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 201 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 85 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 54 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 42 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 262 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 43 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 57 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 55 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 76 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 167 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 329.8K | Targets: 50 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.5K | 96.7% |
| phishing_army | blocklist | domain | 144.8K | 108.4K | 74.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 801 | 63.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 43.3% |
| kadantiscam | blocklist | domain | 48.2K | 16.7K | 34.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 78.3K | 24.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 18.1K | 20.3% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 90.4K | 20.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 4.8K | 13.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 995 | 8.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 10.8K | 8.1% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 3.9K | 6.9% |
| quidsup_notrack-malware | blocklist | domain | 138 | 8 | 5.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 17.3K | 5.3% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1.1K | 3.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 29.7K | 2.9% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 11.6K | 2.5% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.7K | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 2 | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1.3K | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 244 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 57 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 32 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 15 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 6 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 3 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 9 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 138 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 168 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 25 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 10 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 16 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 16 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 31 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 18 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 56 | Unique: 0 | Conflicts: 91</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.9K | 52.9% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 182 | 11.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1.9K | 10.6% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 48 | 10.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 40 | 5.8% |
| hufilter | blocklist | hostname | 96 | 5 | 5.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 514 | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| WaLLy3K | blocklist | domain | 350 | 13 | 3.7% |
| quidsup_notrack-malware | blocklist | domain | 138 | 5 | 3.6% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 493 | 3.2% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 568 | 3.2% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.6K | 2.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 2.2K | 2.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 670 | 2.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 8 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 4 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1.4K | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 17 | 0.7% |
| HaGeZi Pro | blocklist | domain | 438.7K | 2.5K | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 985 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.9K | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 30 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 42 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 4 | 0.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 24 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 18 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 18 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 11 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 26 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 5 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 5 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.5K | Targets: 31 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 4.1K | 71.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 93 | 39.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 43 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 25 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 74 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 25 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 21 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 12 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 6 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 15 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 6 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 30 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 26 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 86.1K | Targets: 36 | Unique: 22.4K | Conflicts: 77</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 3.8K | 47.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 598 | 45.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.4K | 42.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 563 | 40.9% |
| Greensnow | blocklist | ipv4 | 3.5K | 1.4K | 40.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.0K | 38.3% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 4.2K | 34.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 5.4K | 33.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 3.4K | 32.7% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 4.7K | 32.0% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 236 | 30.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 6.6K | 29.4% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 16.3K | 28.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 501 | 25.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 3.1K | 18.9% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 145 | 15.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 47 | 12.1% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 42 | 10.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 241 | 9.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 153 | 6.8% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 15 | 3.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 9 | 2.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 46 | 1 | 2.2% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 259 | 1.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 3.3K | 1.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 204 | 3 | 1.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 27 | 1.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 161 | 1.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 76 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 76 | 0.7% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 16 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 7 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.9K | 2 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 22 | Unique: 209 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 237 | 56 | 23.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 6 | 0.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 12 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 16 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 10 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 13 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 399 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 372 | 95.4% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 380 | 3.1% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 36 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 16 | 1.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 94 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 388 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 106 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 88 | 0.6% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 3 | 0.4% |
| Greensnow | blocklist | ipv4 | 3.5K | 13 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 27 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 12 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 42 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 7 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 127 | Targets: 6 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 68 | 95.8% |
| Torrent Trackers | blocklist | domain | 619 | 126 | 20.4% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 3 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 23 | Unique: 0 | Conflicts: 29</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.1K | 7.2K | 59.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 3.9K | 24.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 1.7K | 20.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 4.0K | 17.9% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 8.3K | 14.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 1.4K | 13.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 6.4K | 7.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 276 | 5.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 665 | 4.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 9 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 8 | 2.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 72 | 2.0% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 216 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 19 | 1.4% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 15 | 0.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 1.0K | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 9 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 29 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 29 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 35 | 0.2% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 10 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 8 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.2K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 997 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 10 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 4 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 60 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 37.0K | Targets: 45 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 3.6K | 6.2% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 7 | 2.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 8.4K | 2.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 16.4K | 1.6% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 4.8K | 1.5% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 5.6K | 1.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 10 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 921 | 0.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 33 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 176 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.3K | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| phishing_army | blocklist | domain | 144.8K | 462 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 140 | 0.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 28 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 24 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 203 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 106 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 75 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 12 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 19 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 545 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 3 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 12 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 50 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 410 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 84 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 11 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 11 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 5 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 12.3K | Targets: 56 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| YousList | blocklist | hostname | 625 | 108 | 17.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 12.2K | 13.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 429 | 12.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1.1K | 6.0% |
| WaLLy3K | blocklist | domain | 350 | 20 | 5.7% |
| hufilter | blocklist | hostname | 96 | 5 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 20 | 2.9% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 466 | 2.6% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 6 | 1.5% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 2.5K | 1.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 349 | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 197 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 9 | 1.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 581 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2.7K | 0.8% |
| HaGeZi Pro | blocklist | domain | 438.7K | 3.2K | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 995 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 540 | 0.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 48 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 90 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 133 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 122 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 33 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 4 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 689 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 11 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 24 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 42 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 65 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 24 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 39 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 85 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 17 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 34 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 14.2K | Targets: 5 | Unique: 14.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 997 | 9 | 0.9% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 831 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 442 | 2 | 0.5% |
| EasyList | blocklist | adguard | 75.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 5 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.3K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 98 | 10.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 10 | 3.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 14 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 598 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 115 | 0.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 23 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 307 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 52 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 9 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 27 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 33 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 12 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 12 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 3 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 36 | 0.0% |

</details>

---

### Dogino_Discord_Official

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 4 | Unique: 0 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 8 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 26 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 7 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 7 | 0.2% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.9K | Targets: 10 | Unique: 202 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 98.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 87 | 11.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 95 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 13 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2K | Targets: 9 | Unique: 94 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1.1K | 31.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 7 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 8 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 85 | 5.9% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 87 | 4.5% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 1 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 5.1K | 31.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 5.1K | 22.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 994 | 12.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 680 | 6.6% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 1.7K | 2.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 429 | 2.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 2.0K | 2.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 8 | 2.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 276 | 1.8% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 14 | 1.8% |
| Greensnow | blocklist | ipv4 | 3.5K | 57 | 1.6% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 6 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 15 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 12 | 0.6% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 1.0K | 0.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 67 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 31 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 1 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 1 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 54.6K | Targets: 17 | Unique: 17.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 90.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 165 | 45.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 27.2K | 16.4% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 73 | 5.5% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 2.6K | 4.5% |
| AdGuard Base filter | blocklist | adguard | 997 | 38 | 3.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 521 | 2.9% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 5.3K | 1.6% |
| abpvn_hosts | blocklist | adguard | 1.2K | 2 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 2 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 5 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 7 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 75.8K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 35.2K | 62.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 56.0K | 33.8% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 35.3K | 10.8% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 52 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 402 | 2.2% |
| AdGuard Base filter | blocklist | adguard | 997 | 9 | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 68 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 301 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 2 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 4 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 4 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.2K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 5 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 72 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 390 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 399 | 372 | 93.2% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 362 | 3.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 31 | 1.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 97 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 17 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 389 | 0.7% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 5 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 84 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 79 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 15 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 29 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 13 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 47 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.6K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.6K | 1.6K | 99.6% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.5K | 33.1% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 2.3K | Targets: 65 | Unique: 934 | Conflicts: 1.0K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 26 | 60.5% |
| local_source_domain_allowlist | allowlist | domain | 43 | 16 | 37.2% |
| tranco | allowlist | domain_top | 500 | 171 | 34.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 76 | 10.9% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 8 | 7.3% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 6 | 6.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 32 | 4.6% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 7 | 3.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 12 | 2.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 3 | 1.7% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 43 | 1.3% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 43 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 177 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 49 | 0.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 7 | 0.6% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 2 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 17 | 0.4% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 114 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 33 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 34 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 6 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 9 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 17 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 84 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 72 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 31 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 10 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 8 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 10 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 71 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 8 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 13 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 9 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 25 | 1.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 4 | 0.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 19 | Unique: 4.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 96 | 8 | 8.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 38 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 23 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 17 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 18 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 18 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 4 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 35 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |

</details>

---

### Firehol_BitcoinNodes_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 7.9K | Targets: 4 | Unique: 7.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 15 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 332 | Targets: 12 | Unique: 269 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 46 | 1 | 2.2% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 14 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 10 | 0.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 7 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 4 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 9 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 9 | Unique: 464 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 15 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 5 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 2 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 7 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 6 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 2 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.3K | Targets: 26 | Unique: 1.7K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 92 | 4.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 46 | 2 | 4.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 7 | 2.1% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 17 | 1.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 204 | 2 | 1.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 9 | 0.7% |
| Greensnow | blocklist | ipv4 | 3.5K | 15 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 7 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 25 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 17 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 99 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 153 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 41 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 22 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 5 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 6 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 9 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 4 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.5K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.5K | 90.2% |
| spamhaus_drop | blocklist | cidr_ipv4 | 1.6K | 1.5K | 90.1% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 14.7K | Targets: 31 | Unique: 0 | Conflicts: 1.4K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 1.8K | 92.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 3.2K | 91.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 749 | 54.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 7.0K | 31.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 4.8K | 30.3% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 106 | 26.6% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 187 | 24.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 84 | 21.5% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 10.8K | 18.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 1.2K | 14.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 1.4K | 13.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 1.4K | 13.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 1.0K | 9.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 142 | 5.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 4.7K | 5.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 665 | 4.4% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 322 | 2.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 33 | 2.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 1 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 22 | 1.0% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 9 | 1.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 1.8K | 0.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 68 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 37 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 1 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 4 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.1K | Targets: 28 | Unique: 0 | Conflicts: 81</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 40 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 380 | 95.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 362 | 92.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.2K | 48.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 7.0K | 43.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 2.2K | 27.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 2.3K | 14.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 1.3K | 12.7% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 7.2K | 12.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 6.6K | 7.6% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Greensnow | blocklist | ipv4 | 3.5K | 152 | 4.3% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 28 | 3.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 44 | 3.2% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 59 | 2.9% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 322 | 2.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 27 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 41 | 1.8% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2.8K | 1.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 155 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 7 | 0.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 78 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 78 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 94 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 15 | Unique: 1.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 204 | 29 | 14.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 92 | 4.1% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 14 | 1.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 4 | 1.2% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 3 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 27 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 1 | 0.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 4 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 8 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 28 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 1 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 204 | Targets: 4 | Unique: 168 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 29 | 1.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 2 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 31.5K | Targets: 26 | Unique: 9.0K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 7.1K | 7.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 537 | 3.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 5.5K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 3.2K | 1.7% |
| Adaway | blocklist | hostname | 6.5K | 105 | 1.6% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.5K | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 30 | 0.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 30 | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 542 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 4 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 11 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 68 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 71 | 0.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 215 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 147 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 42 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 370 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 24 | 0.1% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.7K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 458 | 386 | 84.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 394 | 11.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 182 | 4.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 494 | 1.9% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 212 | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1.5K | 0.9% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.6K | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 398 | 0.4% |
| HaGeZi Pro | blocklist | domain | 438.7K | 1.6K | 0.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 203 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 632 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 22 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 80 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 54.6K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 1.5K | 0.9% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 1.6K | 0.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 22 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 80 | 0.1% |
| EasyList | blocklist | adguard | 75.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 1 | 0.0% |

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
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 13 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 2 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.5K | Targets: 30 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 554 | 27.6% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 3.2K | 22.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 237 | 17.2% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 115 | 15.1% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 3.3K | 5.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 15 | 3.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 89 | 3.5% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 13 | 3.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 499 | 3.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 233 | 2.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 23 | 1.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 1.4K | 1.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 1 | 1.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 226 | 1.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 57 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 82 | 1.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 152 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 15 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 55 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 72 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 4 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 1 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 9 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 9 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 18 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 41 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 695 | Targets: 20 | Unique: 0 | Conflicts: 34</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 101 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 22 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 8 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 680 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 376 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 20 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 24 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 66 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 52 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 21 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 10 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 110 | Targets: 14 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | domain | 17.9K | 86 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 15 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 9 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 12 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 27 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 65 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 12 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 38 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 133.4K | Targets: 15 | Unique: 92.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 1.5K | 10.0% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 9.5K | 6.4% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 16.3K | 5.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 12.2K | 1.2% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 356 | 0.6% |
| EasyList | blocklist | adguard | 75.8K | 301 | 0.4% |
| AdGuard Base filter | blocklist | adguard | 997 | 3 | 0.3% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 365 | 0.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 15 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 14 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 54.6K | 2 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 133.4K | Targets: 53 | Unique: 8.2K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 662 | 52.1% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 6.5K | 22.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 24 | 9.9% |
| kadantiscam | blocklist | domain | 48.2K | 4.2K | 8.7% |
| phishing_army | blocklist | domain | 144.8K | 10.6K | 7.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 5.9% |
| HaGeZi Pro | blocklist | domain | 438.7K | 24.7K | 5.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 16.3K | 5.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 4.4K | 4.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 15.3K | 4.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 190 | 3.8% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 10.8K | 3.3% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 921 | 2.5% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1.2K | 2.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 18.8K | 1.9% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 78 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.7K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 328 | 0.8% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 3 | 0.8% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 2.5K | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 356 | 0.6% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 872 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 35 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 718 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 39 | 0.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 5 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 43 | 0.3% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 23 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 28 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 365 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 32 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 8 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 15 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 14 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 20 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 100 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 34 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 1 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.4K | Targets: 6 | Unique: 3.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-malware | blocklist | adguard | 977.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 38 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.4K | Targets: 14 | Unique: 2.2K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1.1K | 90.3% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 38 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 60 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 204.9K | Targets: 42 | Unique: 196.0K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1.1K | 43.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 99 | 5.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.4K | 3.9% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 439 | 2.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 718 | 0.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.1K | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 817 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 8 | 0.2% |
| kadantiscam | blocklist | domain | 48.2K | 83 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 307 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 131 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 16 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 17 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 168 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 56 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 18 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 97 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 5 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 69 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 6 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 10 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 100 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 19 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 5 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 17 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 11 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.5K | Targets: 18 | Unique: 0 | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 438.7K | 14.3K | 3.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 12 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 48 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 71 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 10 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 347 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 34 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 38 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 21 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 77 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 47 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 25 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 27 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 64 | 0.0% |

</details>

---

### HaGeZi Most Abused TLDs

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 442 | Targets: 1 | Unique: 440 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.2K | 2 | 0.0% |

</details>

---

### HaGeZi Pro

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 438.7K | Targets: 68 | Unique: 0 | Conflicts: 121</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 56.6K | 99.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 99.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 465 | 98.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 14.3K | 98.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 680 | 97.8% |
| hufilter | blocklist | hostname | 96 | 90 | 93.8% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 367 | 91.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.2K | 90.5% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 400 | 87.3% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| Adaway | blocklist | hostname | 6.5K | 4.4K | 67.7% |
| YousList | blocklist | hostname | 625 | 411 | 65.8% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 11.4K | 64.3% |
| quidsup_notrack-malware | blocklist | domain | 138 | 88 | 63.8% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 114.8K | 61.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.5K | 59.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 65 | 59.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 86.9K | 52.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 154 | 52.4% |
| WaLLy3K | blocklist | domain | 350 | 157 | 44.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 9.3K | 34.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 13.9K | 32.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 4.9K | 27.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 84.5K | 25.9% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 3.2K | 25.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 19.4K | 21.7% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 19.6K | 19.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 3.0K | 19.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 24.7K | 18.5% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 311 | 17.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 13.4K | 17.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 4.5K | 14.2% |
| kadantiscam | blocklist | domain | 48.2K | 4.9K | 10.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 22 | 9.1% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1.1K | 4.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 72 | 3.2% |
| phishing_army | blocklist | domain | 144.8K | 4.5K | 3.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 125 | 2.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 30 | 2.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 406 | 2.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 26 | 2.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 50 | 1.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 60 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 11 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 50 | 1.5% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 545 | 1.5% |
| Spam404 | blocklist | domain | 8.1K | 126 | 1.5% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 50 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 885 | 1.5% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1.1K | 1.4% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 4.7K | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 222 | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 167 | 1.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 61 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 7 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 1.5K | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 817 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 3.5K | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 181 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 872 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 427 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 17 | 0.2% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 527 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 11 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 62 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 3 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 470 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 59 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 42 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 173 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 465 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 29 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 13 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 30 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 23 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 36 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 163 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 10 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1.4K | 74.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 85 | 11.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 97 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 3.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 13 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 58.1K | Targets: 34 | Unique: 0 | Conflicts: 141</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 8.1K | 100.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 389 | 99.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 15.6K | 98.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 16.6K | 98.3% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 388 | 97.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 3.3K | 92.4% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 10.8K | 73.7% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 1.5K | 72.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 904 | 65.6% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 7.2K | 59.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8.3K | 55.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 4.3K | 41.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.7K | 33.2% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 251 | 32.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 6.6K | 29.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 307 | 23.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 16.3K | 18.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 7 | 17.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 2.6K | 16.3% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 127 | 13.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 176 | 6.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 99 | 4.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 12 | 3.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 226 | 1.6% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 141 | 1.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 141 | 1.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 5 | 1.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 2.1K | 0.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 15 | 0.7% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 22 | 0.4% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 10 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 11</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 4 | 0.9% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 23 | 0.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 4 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 29 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 28 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 19 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 9 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 53 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 106 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 109 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 13 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 154 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 21 | 0.0% |

</details>

---

### hufilter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 96 | Targets: 24 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 89 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 20 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 90 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 33 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 74 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 87 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 6 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 831 | Targets: 2 | Unique: 826 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.2K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 59 | Unique: 422.0K | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 959 | 13.2% |
| quidsup_notrack-malware | blocklist | domain | 138 | 14 | 10.1% |
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 7 | 6.9% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1.8K | 6.2% |
| hufilter | blocklist | hostname | 96 | 4 | 4.2% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 11.6K | 3.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 3.7K | 2.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 996 | 2.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 117 | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 4 | 1.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 74 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 4.5K | 1.4% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 410 | 1.1% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4.6K | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 27 | 0.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 2.3K | 0.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 7 | 0.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 97 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 5.1K | 0.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 226 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 39 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 3 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 339 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 51 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 339 | 0.2% |
| phishing_army | blocklist | domain | 144.8K | 263 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 195 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 34 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 872 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 19 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 50 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 46 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 307 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 199 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 26 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 50 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 22 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 48.2K | Targets: 44 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 39.4K | 44.1% |
| phishing_army | blocklist | domain | 144.8K | 20.8K | 14.4% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 16.7K | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 15.3K | 4.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 8 | 3.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 4.2K | 3.1% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 796 | 2.8% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.9K | 1.1% |
| Spam404 | blocklist | domain | 8.1K | 25 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 33 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 2 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 18 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 12 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 69 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 50 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 14 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 15 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 8 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 101 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 19 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 83 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 120 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 136 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 13 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 30 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 16 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 14 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 7 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 5 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 23 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 237 | Targets: 28 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 56 | 13.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 93 | 2.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 60 | 1.0% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 28 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 8 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 17 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 4 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 165.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 3 | 0.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 49 | Targets: 2 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | ipv4_from_domain | 49 | 49 | 100.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 2 | 0.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 24 | Targets: 5 | Unique: 0 | Conflicts: 26</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |

</details>

---

### local_ai_blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 49 | Targets: 2 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | ipv4_from_domain | 49 | 49 | 100.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 2 | 0.0% |

</details>

---

### local_ai_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 24 | Targets: 5 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |

</details>

---

### local_domain_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 6 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 6 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 6 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 10 | Unique: 0 | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1 | 0.0% |

</details>

---

### local_mobile_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 4 | Targets: 1 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 4 | 0.0% |

</details>

---

### local_social_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1 | Targets: 4 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### local_source_domain_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 2 | Unique: 25 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 16 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |

</details>

---

### local_source_ipv4_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 62 | Targets: 3 | Unique: 57 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 15.2K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-malware | blocklist | adguard | 977.3K | 14.0K | 1.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 1.5K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 2.4K | 0.7% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 3 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.0M | Targets: 62 | Unique: 684.6K | Conflicts: 30</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 37.2K | 64.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 689 | 54.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 145.6K | 44.4% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 16.4K | 44.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 138 | 35 | 25.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 18.8K | 14.1% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 29.7K | 9.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 3.4K | 8.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 35.0K | 7.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 11 | 6.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 689 | 5.6% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.4K | 3.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 9.3K | 2.9% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.9K | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 262 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 293 | 1.6% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 6 | 1.5% |
| Spam404 | blocklist | domain | 8.1K | 109 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 281 | 1.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.1K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 587 | 1.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1.9K | 1.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 3.5K | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 144 | 0.8% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 780 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 25 | 0.5% |
| phishing_army | blocklist | domain | 144.8K | 724 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 25 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 3 | 0.4% |
| kadantiscam | blocklist | domain | 48.2K | 136 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 35 | 0.2% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 49 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 110 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 32 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 27 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 100 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 226.2K | Targets: 39 | Unique: 201.9K | Conflicts: 48</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 4.3K | 87.1% |
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 23 | 57.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 5.1K | 35.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.0K | 20.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 2.8K | 12.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 1.8K | 11.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 656 | 8.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.0K | 6.9% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 757 | 4.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 3.3K | 3.9% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 2.1K | 3.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 340 | 3.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 36 | 2.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 440 | 2.6% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| Greensnow | blocklist | ipv4 | 3.5K | 41 | 1.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 204 | 2 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 13 | 0.9% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 13 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 8 | 0.9% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 95 | 0.8% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 13 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 48 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 2 | 0.5% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 2 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 73 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 10 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 4 | 0.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 8 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 9 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 47 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 47 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 1 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.9K | 15 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 45 | 0.1% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.5K | Targets: 34 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 144.8K | 12.4K | 8.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 20 | 8.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 23.2K | 7.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 6.5K | 4.9% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 8 | 3.4% |
| kadantiscam | blocklist | domain | 48.2K | 796 | 1.7% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 176 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 25 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.8K | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 21 | 0.4% |
| HaGeZi Pro | blocklist | domain | 438.7K | 1.1K | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 1.1K | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 145 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 7 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 68 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 19 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 49 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 34 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 51 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 12 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 326.7K | Targets: 70 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 99.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 56.1K | 98.6% |
| hufilter | blocklist | hostname | 96 | 87 | 90.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 1.1K | 89.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 398 | 86.9% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 23.2K | 81.3% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 301 | 75.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 68.9% |
| phishing_army | blocklist | domain | 144.8K | 74.4K | 51.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 45.5% |
| YousList | blocklist | hostname | 625 | 272 | 43.5% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 109 | 37.1% |
| quidsup_notrack-malware | blocklist | domain | 138 | 51 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 6.4K | 36.5% |
| Adaway | blocklist | hostname | 6.5K | 2.3K | 35.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 163 | 34.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 56.0K | 33.9% |
| kadantiscam | blocklist | domain | 48.2K | 15.3K | 31.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 13.1K | 30.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 24.8K | 27.8% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 46.0K | 24.7% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 27 | 24.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 6.5K | 24.2% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 78.3K | 23.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 16.6K | 22.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2.7K | 21.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 3.8K | 21.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 347 | 20.0% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.6% |
| HaGeZi Pro | blocklist | domain | 438.7K | 84.5K | 19.3% |
| WaLLy3K | blocklist | domain | 350 | 56 | 16.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 33 | 13.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 16.3K | 12.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 521 | 10.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.5K | 9.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 36 | 5.2% |
| tranco | allowlist | domain_top | 500 | 18 | 3.6% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1.3K | 3.5% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1.5K | 2.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 114 | 1.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 18 | 1.4% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 1.3K | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 201 | 1.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 370 | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 38 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 31 | 1.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 3.3K | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 7 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.5K | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 31 | 1.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 9.3K | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 406 | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 19 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 108 | 0.7% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 477 | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 2.1K | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 21 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1.1K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 35 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 64 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 18 | 0.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 3 | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 257 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 326.7K | Targets: 25 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 56.1K | 98.6% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 826 | 62.6% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 72.2K | 48.5% |
| EasyList | blocklist | adguard | 75.8K | 35.3K | 46.6% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 56.0K | 33.8% |
| AdGuard Base filter | blocklist | adguard | 997 | 301 | 30.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 3.8K | 21.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 66 | 18.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 2.4K | 15.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 16.3K | 12.2% |
| Easy Privacy | blocklist | adguard | 54.6K | 5.3K | 9.7% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 80.1K | 8.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 60 | 3.3% |
| abpvn_hosts | blocklist | adguard | 1.2K | 33 | 2.8% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 37 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 18 | 1.4% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 38 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 108 | 0.7% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 831 | 1 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.2K | 5 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.5K | Targets: 39 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 6.1K | 10.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 6.7K | 8.7% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 11.6K | 5.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 23 | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 138 | 1 | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 21 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| Torrent Trackers | blocklist | domain | 619 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 45 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 24 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 76 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 121 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 11 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 222 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 20 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 196 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 37 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 108 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 3 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 14 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 19 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 24 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 10 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 11 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.5K | Targets: 8 | Unique: 16.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 23 | 1.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 35 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 76 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 121 | 0.1% |
| EasyList | blocklist | adguard | 75.8K | 68 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 108 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 56.8K | Targets: 63 | Unique: 0 | Conflicts: 26</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 96 | 89 | 92.7% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 299 | 74.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 46.6% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 37.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 51.8K | 31.4% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 4.7K | 26.7% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 37.7K | 20.2% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| quidsup_notrack-malware | blocklist | domain | 138 | 25 | 18.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 7.5K | 17.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 56.1K | 17.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 4.0K | 15.2% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 56.6K | 12.9% |
| Adaway | blocklist | hostname | 6.5K | 739 | 11.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1.4K | 7.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.0K | 6.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 29 | 6.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 26 | 5.7% |
| WaLLy3K | blocklist | domain | 350 | 20 | 5.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 80 | 4.8% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 581 | 4.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.6K | 4.1% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 11 | 1.6% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 953 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 5 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 85 | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 147 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 76 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 219 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 265 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 356 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 2 | 0.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 27 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 11 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 75 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 1 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 25 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 587 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 50 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 73 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 226 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 66 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 56 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 161 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.8K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 850 | 64.4% |
| EasyList | blocklist | adguard | 75.8K | 35.2K | 46.5% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 51.8K | 31.2% |
| AdGuard Base filter | blocklist | adguard | 997 | 299 | 30.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 56.1K | 17.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 41 | 11.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 1.4K | 7.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 80 | 4.8% |
| Easy Privacy | blocklist | adguard | 54.6K | 2.6K | 4.7% |
| abpvn_hosts | blocklist | adguard | 1.2K | 29 | 2.5% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 28 | 1.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 76 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 356 | 0.3% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 24 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 2 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 3 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 1 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 4 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 100 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 1 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 242 | Targets: 11 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 20 | 0.1% |
| phishing_army | blocklist | domain | 144.8K | 124 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 22 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 33 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 24 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 3 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 71 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 68 | 53.5% |
| Torrent Trackers | blocklist | domain | 619 | 70 | 11.3% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.9K | Targets: 73 | Unique: 0 | Conflicts: 333</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 86 | 78.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 667 | 18.9% |
| Adaway | blocklist | hostname | 6.5K | 1.1K | 16.5% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| YousList | blocklist | hostname | 625 | 95 | 15.2% |
| WaLLy3K | blocklist | domain | 350 | 50 | 14.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 568 | 13.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 59 | 12.6% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 28 | 11.8% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 2.0K | 11.5% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 47 | 10.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 29 | 9.9% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 1 | 9.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 59 | 8.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.2K | 8.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 177 | 7.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1.6K | 6.0% |
| hufilter | blocklist | hostname | 96 | 5 | 5.2% |
| quidsup_notrack-malware | blocklist | domain | 138 | 7 | 5.1% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 4 | 4.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 466 | 3.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 24 | 3.5% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 12 | 3.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.2K | 2.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 76 | 2.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.4K | 2.5% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 4.7K | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 2.2K | 2.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 60 | 1.9% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 60 | 1.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2.2K | 1.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 22 | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 3.8K | 1.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.9K | 1.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 43 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 32 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 8 | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 71 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 35 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 71 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 7 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 439 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 19 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 57 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 122 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 80 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 64 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 15 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 60 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 293 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 29 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 7 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 25 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 16 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 7 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 17.9K | Targets: 21 | Unique: 9.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 5 | 71.4% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 174 | 13.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 18 | 4.9% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 1.4K | 2.5% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 22 | 1.3% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 2.2K | 1.3% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 3.8K | 1.2% |
| AdGuard Base filter | blocklist | adguard | 997 | 12 | 1.2% |
| Easy Privacy | blocklist | adguard | 54.6K | 521 | 1.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 9 | 0.6% |
| CJX Annoyance | blocklist | adguard | 1.8K | 10 | 0.6% |
| EasyList | blocklist | adguard | 75.8K | 402 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 35 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 1 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 12 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 54 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 15 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 144.8K | Targets: 38 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 124 | 51.2% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 12.4K | 43.6% |
| kadantiscam | blocklist | domain | 48.2K | 20.8K | 43.2% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 108.4K | 32.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 74.4K | 22.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 17.7K | 19.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 10.6K | 8.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 11 | 4.6% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 462 | 1.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 4.5K | 1.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 30 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 25 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 756 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 724 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 7 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 263 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 12 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 32 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 32 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 2 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 73 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 83 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 9 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 21 | Unique: 61.6K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 97 | 6.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 95 | 4.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 28 | 1.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 31 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 4 | 0.1% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 1 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 2 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 10 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 45 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 7 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 6 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 1 | 0.0% |

</details>

---

### PuppyScams

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 102 | Targets: 2 | Unique: 92 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 3 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 458 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 386 | 23.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 163 | 4.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 87 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 47 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 56 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 291 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 167 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 403 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 400 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 398 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 26 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 138 | Targets: 29 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 71 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 8 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 20 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 32 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 51 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 33 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 88 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 25 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.6K | Targets: 54 | Unique: 0 | Conflicts: 87</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 593 | 16.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 493 | 11.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| tranco | allowlist | domain_top | 500 | 40 | 8.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1.2K | 7.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1.2K | 7.0% |
| Adaway | blocklist | hostname | 6.5K | 444 | 6.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 958 | 3.6% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 96 | 3 | 3.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 13 | 2.8% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 10 | 2.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 4.0K | 2.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.0K | 1.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 537 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 11 | 1.6% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 197 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 10 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2.3K | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.3K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.0K | 1.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 25 | 0.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 1 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 25 | 0.8% |
| HaGeZi Pro | blocklist | domain | 438.7K | 3.0K | 0.7% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 656 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 263 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.5K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 38 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| Torrent Trackers | blocklist | domain | 619 | 2 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 39 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 26 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 23 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 4 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 35 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 1 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 75.8K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 977.3K | Targets: 15 | Unique: 759.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 14.0K | 92.6% |
| RPiList_specials-phishing | blocklist | adguard | 148.8K | 110.9K | 74.5% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 80.1K | 24.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 12.2K | 9.1% |
| AdGuard Base filter | blocklist | adguard | 997 | 4 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 54 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 100 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.3K | 3 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 289 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 15 | 0.1% |
| EasyList | blocklist | adguard | 75.8K | 72 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.4K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 7 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 148.8K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 326.7K | 72.2K | 22.1% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 110.9K | 11.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 9.5K | 7.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 12 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 15.2K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 11 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 4 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 22 | Unique: 0 | Conflicts: 105</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 764 | 78 | 10.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 237 | 6.7% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 749 | 5.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 17 | 4.4% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 16 | 4.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 904 | 1.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 118 | 1.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 105 | 1.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 105 | 1.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 68 | 0.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 116 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 563 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 87 | 0.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 44 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 25 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 3 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 19 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 13 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 3 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 923 | Targets: 19 | Unique: 443 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 98 | 7.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 14 | 4.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 17 | 0.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 14 | 0.7% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 8 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 127 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 145 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 7 | 0.1% |
| Greensnow | blocklist | ipv4 | 3.5K | 4 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 8 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 5 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 5 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 6 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 2 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 17 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 764 | 96 | 12.6% |
| Greensnow | blocklist | ipv4 | 3.5K | 89 | 2.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 16 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 142 | 1.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 14 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 176 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 241 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 16 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 17 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 10.9K | Targets: 22 | Unique: 0 | Conflicts: 13.6K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 10.9K | 100.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 10.9K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1.4K | 9.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 105 | 7.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 534 | 3.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 32 | 1.7% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 78 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 40 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 58 | 0.4% |
| Greensnow | blocklist | ipv4 | 3.5K | 9 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 4 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 29 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 141 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 20 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 76 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 47 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.4K | Targets: 28 | Unique: 0 | Conflicts: 46</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_mobile_allowlist | allowlist | ipv4_from_domain | 4 | 4 | 100.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 1.3K | 15.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 680 | 13.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.4K | 9.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 118 | 8.6% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 1.0K | 8.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 29 | 7.4% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 4.3K | 7.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 1.1K | 6.9% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 27 | 6.8% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 1.0K | 6.8% |
| Greensnow | blocklist | ipv4 | 3.5K | 233 | 6.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 1.3K | 5.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 804 | 5.0% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 94 | 4.7% |
| local_ai_blocklist | blocklist | ipv4_from_domain | 49 | 2 | 4.1% |
| local_ai_allowlist | allowlist | ipv4_from_domain | 49 | 2 | 4.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 3.4K | 3.9% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 21 | 2.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 325 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 17 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 5 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 40 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 40 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 5 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 29 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 340 | 0.2% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 16.5K | 23 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 18 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 10 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 1 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3K | Targets: 18 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.2K | 12 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 12 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 23 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 26 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 16 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 10 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 13 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 18 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 695 | Targets: 40 | Unique: 264 | Conflicts: 339</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 7 | 1 | 14.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 76 | 3.3% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 2 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 40 | 0.9% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 17 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 17 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 59 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 13 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 11 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 7 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 11 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 11 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 31 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 26.7K | Targets: 55 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 494 | 29.7% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 973 | 27.6% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 79 | 19.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 670 | 15.7% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 56 | 12.2% |
| hufilter | blocklist | hostname | 96 | 11 | 11.5% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1.6K | 9.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1.6K | 9.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 4.0K | 7.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 19 | 6.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 958 | 6.2% |
| Adaway | blocklist | hostname | 6.5K | 407 | 6.2% |
| quidsup_notrack-malware | blocklist | domain | 138 | 7 | 5.1% |
| tranco | allowlist | domain_top | 500 | 24 | 4.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 6.7K | 4.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 6.3K | 3.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 349 | 2.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| HaGeZi Pro | blocklist | domain | 438.7K | 9.3K | 2.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.9K | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 6.5K | 2.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 17 | 0.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 3 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 45 | 0.3% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 11 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 11 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 160 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 122 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 205 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 24 | 0.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 21 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 51 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 281 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 20 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 16 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 22 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 9 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 214.5K | Targets: 34 | Unique: 160.1K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 11.6K | 70.7% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 19.2K | 31.6% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 21.5K | 28.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 60 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 100 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 73 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 267 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 257 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 15 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 9 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 190 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 50 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 104 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 427 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 199 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 17 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 82 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 23 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 31 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 110 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 23 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 11 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 42.9K | Targets: 47 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 138 | 71 | 51.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 7.5K | 13.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 24 | 6.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 883 | 5.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 7.4K | 4.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 13.1K | 4.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 6.6K | 3.6% |
| HaGeZi Pro | blocklist | domain | 438.7K | 13.9K | 3.2% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 263 | 1.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 59 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 6 | 1.3% |
| hufilter | blocklist | hostname | 96 | 1 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 90 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 314 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 140 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 3.4K | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 57 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 45 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 328 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 15 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 996 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 162 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 244 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 24 | 0.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 1 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 49 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 207 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 43 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 69 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 6 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 142 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 19 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 32 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 12 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.2K | Targets: 29 | Unique: 4.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 3 | 2.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 959 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 78 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 21 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 12 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 11 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 13 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 114 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 12 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 15 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 61 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 28 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 5.8K | Targets: 26 | Unique: 1.3K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4.1K | 90.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 60 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 32 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 21 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 18 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 8 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 25 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 18 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 17 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 25 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 9 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 2 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 60.9K | Targets: 46 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.7K | 60.9K | 79.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 6.1K | 37.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 19.2K | 9.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 2 | 2.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 3 | 2.4% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 65 | 1.8% |
| Torrent Trackers | blocklist | domain | 619 | 10 | 1.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 122 | 1.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 13 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 138 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 138 | 1 | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 615 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 122 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 64 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 219 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 501 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 279 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 12 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 885 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 1 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 23 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 43 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 406 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 5 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 32 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 27 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 11 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 46 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 6 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.6K | Targets: 18 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 2.6K | 3.0% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 11 | 0.6% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1.1K | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 76 | 0.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 19 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 50 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 19 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 4 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.2K | Targets: 35 | Unique: 0 | Conflicts: 101</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Social | blocklist | hostname | 3.2K | 3.2K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 17 | 2.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 43 | 1.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 12 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 23 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 60 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 25 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 32 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 11 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 29 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 53 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 50 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 28 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 23 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 31 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 33 | Unique: 6.0K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 138 | 2 | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.6K | 0.5% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 20 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 55 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 25 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 4 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 109 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 126 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.6K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.6K | 99.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.5K | 32.9% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.1% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 17.5K | 5.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 4.3K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1.1K | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1.2K | 0.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 5.0K | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 4 | 0.3% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 956 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 382 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 406 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 28 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 26 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 19 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 1 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 89.3K | Targets: 69 | Unique: 0 | Conflicts: 153</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2.6K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.8% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 12.2K | 99.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.5K | 98.0% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 48.2K | 39.4K | 81.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 51.4% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 7.6K | 43.3% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 167 | 36.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 393 | 30.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.8K | 27.4% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 398 | 23.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| quidsup_notrack-malware | blocklist | domain | 138 | 20 | 14.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| hufilter | blocklist | hostname | 96 | 12 | 12.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 899 | 12.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2.2K | 12.2% |
| phishing_army | blocklist | domain | 144.8K | 17.7K | 12.2% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 12 | 10.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.3K | 8.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 36 | 7.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 24.8K | 7.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 52 | 7.5% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 13.6K | 7.3% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1.9K | 7.1% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 28 | 7.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3.6K | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 18.1K | 5.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 31 | 4.5% |
| HaGeZi Pro | blocklist | domain | 438.7K | 19.4K | 4.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 84 | 3.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 4.4K | 3.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 4.8K | 2.9% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 542 | 1.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 3.4K | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 615 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 29 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 29 | 0.9% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 522 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 703 | 0.9% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 726 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 55 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 314 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 145 | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 77 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 16 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 106 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 42 | 0.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 887 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 37 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1.9K | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 18 | 0.2% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 246 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 26 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 8 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 195 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 104 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.7K | Targets: 49 | Unique: 0 | Conflicts: 15</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 60.9K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 6.7K | 40.7% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 21.5K | 10.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 2 | 2.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 3 | 2.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 73 | 2.1% |
| hufilter | blocklist | hostname | 96 | 2 | 2.1% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 619 | 10 | 1.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 133 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 154 | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 703 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 138 | 1 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 13 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 160 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 265 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 80 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 602 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 26 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 1.1K | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 349 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 13 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 477 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 49 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 48.2K | 13 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 32 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 34 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 6 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 6 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 18 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 6 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.2K | Targets: 35 | Unique: 0 | Conflicts: 101</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 3.2K | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 17 | 2.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 43 | 1.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 12 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 23 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 60 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 25 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 32 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 50 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 29 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 11 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 28 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 53 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 31 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 23 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 2 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 57.4K | Targets: 27 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 900 | 70.8% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 3.6K | 9.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 18.4K | 5.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 37.2K | 3.7% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 3.9K | 1.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 1.2K | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 522 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.5K | 0.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 15 | 0.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 10 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 19 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 504 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 27 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 181 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 23 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 32 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 10 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 328.1K | Targets: 47 | Unique: 97.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 566 | 44.5% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 18.4K | 32.2% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 8.4K | 22.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 145.6K | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 15.3K | 11.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 426 | 8.6% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 17.3K | 5.3% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 14.1K | 3.1% |
| quidsup_notrack-malware | blocklist | domain | 138 | 4 | 2.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 382 | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 242 | 3 | 1.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 887 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 3.3K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 3 | 0.8% |
| phishing_army | blocklist | domain | 144.8K | 756 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 207 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.3K | 0.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 161 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 15 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 34 | 0.3% |
| HaGeZi Pro | blocklist | domain | 438.7K | 1.5K | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 18 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 48.2K | 120 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 51 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 411 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 312 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 9 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 22 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 9 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 11 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 16 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 15 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 23 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 69 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 82 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.9K | Targets: 28 | Unique: 0 | Conflicts: 58</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 14.7K | 4.8K | 32.7% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 621 | 31.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 15.6K | 26.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3.9K | 26.1% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 88 | 22.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 79 | 20.3% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 2.2K | 18.2% |
| Greensnow | blocklist | ipv4 | 3.5K | 499 | 14.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 1.1K | 13.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 1.1K | 10.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 2.3K | 10.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 115 | 8.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 116 | 8.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 429 | 8.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 5.4K | 6.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 896 | 5.5% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 17 | 2.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 203 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 25 | 1.1% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 7 | 0.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 58 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 58 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 757 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 1 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 35 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 619 | Targets: 12 | Unique: 393 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 126 | 99.2% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 70 | 98.6% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 2 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 10 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 2 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 45 | Unique: 0 | Conflicts: 599</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 8 | 18.6% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_blocklist | blocklist | domain | 24 | 2 | 8.3% |
| local_ai_allowlist | allowlist | domain | 24 | 2 | 8.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 171 | 7.5% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 43 | 2 | 4.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 96 | 2 | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 10 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 34 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 27 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 27 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 32 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 180 | 1 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 82 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 22 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 40 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 38 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 1 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 1 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 39 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 24 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 11 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.4K | 3 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 18 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 30 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 37 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 37 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 36 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 16 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 9 | 0.1% |
| EasyList | blocklist | adguard | 75.8K | 52 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 133.4K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 37 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 28 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 977.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 28 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 3 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 367 | Targets: 10 | Unique: 16 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 54.6K | 165 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| Easy Privacy | allowlist | adguard | 817 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 41 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 18 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 51 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 66 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 2 | 0.0% |

</details>

---

### Ukrainian Security Filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.7K | Targets: 16 | Unique: 871 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 11 | 0.4% |
| HaGeZi Pro | blocklist | domain | 438.7K | 311 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 347 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 13 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 8 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 99 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 46 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3K | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 900 | 1.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 662 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 393 | 0.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 1.1K | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 801 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 566 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 689 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 30 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 10 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 5 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 75.9K | Targets: 1 | Unique: 75.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 5 | 5.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.8K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 16.6K | 28.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 4 | 10.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 325 | 3.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 67 | 1.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 203 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 4 | 1.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 4 | 1.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 155 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 45 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 78 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 68 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 259 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 13 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 440 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 2 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 19 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 16 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 3 | 0.1% |
| Greensnow | blocklist | ipv4 | 3.5K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 6 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 452.2K | Targets: 56 | Unique: 292.8K | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 90.4K | 27.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 813 | 16.5% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 5.6K | 15.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 17 | 7.2% |
| quidsup_notrack-malware | blocklist | domain | 138 | 8 | 5.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 956 | 5.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 14.1K | 4.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 35.0K | 3.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 9 | 2.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 2.5K | 1.9% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 4 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.6K | 1.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 504 | 0.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 680 | 0.9% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 127 | 1 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 5 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 85 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2.1K | 0.6% |
| YousList | blocklist | hostname | 625 | 4 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 26 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 76 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 1.3K | 5 | 0.4% |
| Torrent Trackers | blocklist | domain | 619 | 2 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 246 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 142 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 17 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 444 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 68 | 0.2% |
| HaGeZi Pro | blocklist | domain | 438.7K | 527 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 191 | 0.1% |
| phishing_army | blocklist | domain | 144.8K | 83 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 25 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 50 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 20 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 6 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 5 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 4 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 11 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 6 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 14 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 9 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 14.5K | Targets: 36 | Unique: 8.2K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 3 | 7.5% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 283 | 5.7% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 52 | 3.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 46 | 1 | 2.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 5.1K | 2.2% |
| Sblam_Blocklist | blocklist | ipv4 | 923 | 6 | 0.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 31 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 332 | 2 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 78 | 0.5% |
| BruteforceBlocker | blocklist | ipv4_find | 399 | 2 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 390 | 2 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.5K | 18 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 8.1K | 34 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 226 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 94 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.4K | 29 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 42 | 0.3% |
| BlockListDE_Strong | blocklist | ipv4 | 764 | 2 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 6 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 161 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 35 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 2.0K | 4 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 37 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.7K | 27 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 35 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.4K | 3 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 10.9K | 3 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 10.9K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.9K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 2.1K | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 4.0K | 0.4% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 813 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 521 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 402 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 190 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 426 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 33 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 125 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.0K | Targets: 14 | Unique: 280 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 40 | 3 | 7.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 283 | 2.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 4.3K | 1.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.3K | 3 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 13 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 22 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 7 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 86.1K | 16 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.9K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.1K | 5 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 96 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 75.9K | 5 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 40 | Targets: 6 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.1K | 40 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.0K | 3 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 23 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 350 | Targets: 33 | Unique: 0 | Conflicts: 11</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| YousList | blocklist | hostname | 625 | 9 | 1.4% |
| hufilter | blocklist | hostname | 96 | 1 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 138 | 1 | 0.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 19 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 85 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 3 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 41 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 50 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 53 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 177 | 0.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 2 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 85 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 56 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 12 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 157 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 20 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 35 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 66 | Unique: 0 | Conflicts: 184</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.7K | 75.8% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3.0K | 69.9% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 6.2K | 35.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 3.6K | 29.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 20.8K | 23.3% |
| WaLLy3K | blocklist | domain | 350 | 81 | 23.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 67 | 14.6% |
| hufilter | blocklist | hostname | 96 | 14 | 14.6% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 2.2K | 12.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 203 | 12.2% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 12 | 10.9% |
| quidsup_notrack-malware | blocklist | domain | 138 | 14 | 10.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 66 | 9.5% |
| tranco | allowlist | domain_top | 500 | 39 | 7.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 1.0K | 6.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 1.8K | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 11.3K | 6.1% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 24 | 6.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3.1K | 5.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 16.6K | 5.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 114 | 5.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 23 | 4.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 27 | 3.9% |
| HaGeZi Pro | blocklist | domain | 438.7K | 13.4K | 3.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 3.2K | 1.9% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 876 | 1.4% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 38 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 38 | 1.2% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 958 | 1.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 215 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 162 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 47 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 260 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 2.4K | 0.2% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 680 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.3K | 2 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 4 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 17 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 24 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 57.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 34 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 50 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 51 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 16 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 16 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 4 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 4 | 0.0% |

</details>

---

### YousList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 625 | Targets: 33 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 199 | 1.1% |
| hufilter | blocklist | hostname | 96 | 1 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 108 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 23 | 0.7% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 2 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 95 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 3 | 0.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 86 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 94 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 430 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 272 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 152 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 438.7K | 411 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 17 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 7 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 367 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 24 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 17.9K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 165.9K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 54.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 75.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 326.7K | 69 | 0.0% |

</details>

---

### youtube_GoodbyeAds

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 23 | Unique: 97.1K | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| hufilter | blocklist | hostname | 96 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 50 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 29 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 74 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 8 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 438.7K | 62 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 40 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 45 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.5K | Targets: 60 | Unique: 0 | Conflicts: 54</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.6% |
| quidsup_notrack-annoyance | blocklist | domain | 458 | 163 | 35.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 394 | 23.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 17.6K | 2.0K | 11.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| quidsup_notrack-malware | blocklist | domain | 138 | 9 | 6.5% |
| WaLLy3K | blocklist | domain | 350 | 19 | 5.4% |
| hufilter | blocklist | hostname | 96 | 5 | 5.2% |
| Adaway | blocklist | hostname | 6.5K | 258 | 3.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.5K | 3.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.6K | 593 | 3.8% |
| ph00lt0_blocklist | blocklist | domain | 17.9K | 667 | 3.7% |
| YousList | blocklist | hostname | 625 | 23 | 3.7% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 4 | 3.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 26.7K | 973 | 3.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.7K | 3.5% |
| Dan Pollock's List | blocklist | hostname | 12.3K | 429 | 3.5% |
| AdGuard Base filter | blocklist | domain_adguard | 400 | 12 | 3.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.6K | 2.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 695 | 11 | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 186.1K | 2.4K | 1.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 470 | 5 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 165.2K | 1.6K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 326.7K | 2.4K | 0.7% |
| HaGeZi Pro | blocklist | domain | 438.7K | 3.2K | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 695 | 3 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 42.9K | 59 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 16.5K | 11 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.7K | 73 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 31.5K | 30 | 0.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 14.5K | 10 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 60.9K | 65 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 452.2K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 28.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.5K | 3 | 0.0% |
| phishing_army | blocklist | domain | 144.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 98.6K | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 328.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 204.9K | 8 | 0.0% |
| kadantiscam | blocklist | domain | 48.2K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 27 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 37.0K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.0M | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 329.8K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 133.4K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 100.1K | 29 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.9K | Targets: 4 | Unique: 8.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.9K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 58.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 226.2K | 48 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 14.5K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

