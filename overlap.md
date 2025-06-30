# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-09-05 13:10:12 UTC

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
| Total Entries Analyzed | 6.9M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 140 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 34 |
| cidr_ipv4 | 3 |
| domain | 86 |
| ipv4 | 39 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 68 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 68.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 42.5K | 67.2% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 11.2K | 60.9% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 297 | 52.3% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 167 | 45.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 142 | 36.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 590 | 36.0% |
| hufilter | blocklist | hostname | 94 | 31 | 33.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 32.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 110 | 31.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 34 | 31.5% |
| HaGeZi Pro | blocklist | domain | 225.2K | 70.0K | 31.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 51.8K | 28.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 3.7K | 24.4% |
| quidsup_notrack-malware | blocklist | domain | 123 | 29 | 23.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 6.1K | 22.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2.6K | 19.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 49.3K | 19.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.5K | 17.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 13.2K | 14.9% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 4.3K | 14.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 6.0K | 13.7% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 15 | 2.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 144 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 52 | 1.6% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 3.3K | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 327 | 1.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 1.4K | 0.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 27 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 2 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 431 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 167 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 442 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 444 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 43 | 0.2% |
| kadantiscam | blocklist | domain | 43.0K | 97 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 5 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 46 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 736 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 231 | 0.1% |
| phishing_army | blocklist | domain | 156.0K | 101 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 214 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 993 | Targets: 8 | Unique: 893 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 32 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 24 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 37 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 49 | Unique: 0 | Conflicts: 37</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2.7K | 14.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 259 | 7.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 6.5K | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 6 | 5.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 21 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 123 | 4 | 3.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 434 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 9 | 2.3% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 10 | 1.8% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 520 | 1.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 408 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 745 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 3 | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 2.4K | 0.9% |
| HaGeZi Pro | blocklist | domain | 225.2K | 1.7K | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 89 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1.0K | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 50 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 2 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 93 | Targets: 8 | Unique: 34 | Conflicts: 58</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 206 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 3 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 35 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 9 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.2K | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.3K | 429 | 0.7% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 568 | 0.3% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 431 | 0.2% |
| EasyList | blocklist | adguard | 67.1K | 68 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 43 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 10 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 993 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 33 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 5 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 568 | Targets: 29 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 429 | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 105 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 550 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 486 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 431 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 33 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 38 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 7 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 11 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 297 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 209.7K | Targets: 14 | Unique: 209.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 444 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 3 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 224.8K | Targets: 24 | Unique: 116.7K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 89.7K | 50.0% |
| hufilter | blocklist | hostname | 94 | 18 | 19.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.3K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 599 | 3.9% |
| HaGeZi Pro | blocklist | domain | 225.2K | 8.5K | 3.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 938 | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 813 | 0.9% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 30 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 204 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 84 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.2K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 2 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 335 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 59 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 43 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 179.4K | Targets: 67 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 568 | 550 | 96.8% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 94.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 58.3K | 92.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 94 | 72 | 76.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.3% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 89.7K | 39.9% |
| HaGeZi Pro | blocklist | domain | 225.2K | 81.7K | 36.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5.1K | 27.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 51.8K | 25.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 62.8K | 24.5% |
| YousList | blocklist | hostname | 625 | 152 | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 123 | 28 | 22.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 969 | 22.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 6.1K | 22.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 21 | 19.4% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2.1K | 13.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 6.0K | 13.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.7K | 11.7% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2.3K | 7.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 24 | 6.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 4.9K | 5.6% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 19 | 4.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 549 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 28 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 122 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 275 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 344 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 11 | 0.2% |
| Torrent Trackers | blocklist | domain | 480 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 180 | 0.2% |
| kadantiscam | blocklist | domain | 43.0K | 29 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 221 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 2 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 282 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 189 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 275 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 10 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 8 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 121 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 5 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 107 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 180.1K | Targets: 25 | Unique: 0 | Conflicts: 195</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 198 | 194 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 93.4% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 58.3K | 92.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.3% |
| EasyList | blocklist | adguard | 67.1K | 47.7K | 71.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 28.8K | 52.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 568 | 47.0% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 62.8K | 24.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 2.3K | 7.8% |
| abpvn_hosts | blocklist | adguard | 993 | 24 | 2.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.7K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 122 | 0.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 221 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 10 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 2 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 242 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 8 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 837 | 1.3% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 821 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 71 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 75 | 0.1% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 198 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 206 | 194 | 94.2% |

</details>

---

### AdGuardTeam_HttpsExclusions_android

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 97 | Targets: 11 | Unique: 70 | Conflicts: 17</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 5 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 10 | Unique: 4.0K | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_issues

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 6 | Unique: 60 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
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
<summary>List Type: allowlist | Source Type: domain | Total: 181 | Targets: 11 | Unique: 163 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 1 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.9% |
| Easy Privacy | blocklist | adguard | 55.2K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 2.1K | 0.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 60 | 0.1% |
| EasyList | blocklist | adguard | 67.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 1 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.4K | Targets: 55 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.0K | 57.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.8% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.7% |
| YousList | blocklist | hostname | 625 | 199 | 31.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 206 | 12.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 12 | 11.1% |
| quidsup_notrack-malware | blocklist | domain | 123 | 13 | 10.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 52 | 9.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 7.6K | 8.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.3K | 8.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1.1K | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1.2K | 7.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 4.8K | 7.5% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1.7K | 6.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 21 | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.2K | 5.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 20 | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.5K | 5.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 7.3K | 3.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 5.1K | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 6.5K | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 919 | 2.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 14 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 32 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 52 | 0.4% |
| Torrent Trackers | blocklist | domain | 480 | 1 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 18 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 12 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 53 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 34 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 43 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.3K | Targets: 20 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 355 | 17 | 4.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 203 | 4.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 384 | 3.7% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 18 | 3.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 14 | 2.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 397 | 2.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 494 | 2.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 35 | 1.8% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 1.2K | 1.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 241 | 1.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 777 | 1.5% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 136 | 1.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 43 | 1.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 172 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 114 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 7 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 13 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 2 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 21.6K | 982 | 4.6% |
| Greensnow | blocklist | ipv4 | 4.3K | 171 | 3.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 569 | 2.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 26 | 1.9% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 11 | 1.1% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 569 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 60 | 0.6% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 2 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 87 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 4 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 171 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 17 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 19 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 31 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 355 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 88 | 4.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 86 | 3.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 129 | 3.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 301 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 17 | 1.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 279 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 46 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 2 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 287 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 2 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 40 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 18 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 2 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 3 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 38 | Unique: 10.5K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 123 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 24 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 45 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 79 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 185 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 196 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 22 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 17 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 50 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 4 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 49 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 31 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 53 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 43 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 250.1K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 156.0K | 105.2K | 67.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 243 | 63.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 75.4K | 41.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 79.0K | 30.8% |
| kadantiscam | blocklist | domain | 43.0K | 6.0K | 14.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 7.6K | 8.6% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 995 | 7.7% |
| quidsup_notrack-malware | blocklist | domain | 123 | 7 | 5.7% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1.8K | 4.8% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 3.3K | 4.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 8.6K | 3.8% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1.1K | 2.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 8.2K | 1.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 3 | 1.3% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 172 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 195 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 9 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 16 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 107 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 236 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 12 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 46 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 25 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 231 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 31 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 13 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 15 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 55 | Unique: 0 | Conflicts: 71</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.8K | 51.7% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 48 | 13.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.9K | 10.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 167 | 10.2% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 40 | 5.6% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| quidsup_notrack-malware | blocklist | domain | 123 | 5 | 4.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 514 | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 13 | 3.7% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 473 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 11 | 2.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1.6K | 2.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 672 | 2.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 2.2K | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 569 | 1.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2.0K | 0.9% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 5 | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.9K | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 969 | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 25 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 42 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 5 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 24 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 19 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 12 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 32 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 4.1K | 70.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 174 | 34.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 2 | 0.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 58 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 25 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 55 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 15 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 75 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 48 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 6 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 11 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 51.1K | Targets: 34 | Unique: 0 | Conflicts: 43</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 355 | 279 | 78.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 777 | 58.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 4 | 2 | 50.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 2.0K | 46.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.9K | 45.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.3K | 44.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 4.4K | 41.7% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 4.6K | 36.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 502 | 36.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 640 | 33.1% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 6.9K | 31.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 7.0K | 30.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 5.1K | 24.7% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 16.8K | 22.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 3.8K | 22.6% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 96 | 16.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 94 | 16.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 171 | 15.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 141 | 14.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 105 | 7.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 156 | 6.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 18 | 5.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 2 | 4.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 416 | 2.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 6 | 2.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 9 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 217 | 1.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 30 | 1.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 43 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 43 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 11 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 20 | Unique: 223 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 4 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 9 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 572 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 536 | 95.0% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 549 | 4.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 31 | 1.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 18 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 198 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 2 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 99 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 402 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 39 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 11 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 96 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 10 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 7 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 8 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 109 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 73 | 73 | 100.0% |
| Torrent Trackers | blocklist | domain | 480 | 108 | 22.5% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 21 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.5K | 8.0K | 64.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 4.2K | 20.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 241 | 18.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 3.1K | 13.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6.9K | 13.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1.4K | 12.9% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 7.0K | 9.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 300 | 5.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 880 | 5.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 87 | 2.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 9 | 1.6% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 9 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 285 | 1.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 24 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 13 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 10 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 54 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 1 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 27 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 27 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 5 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 7 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 993 | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 55 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 80.7K | Targets: 46 | Unique: 36.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 4.3K | 9.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 25 | 6.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 24.8K | 4.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 8 | 3.4% |
| quidsup_notrack-malware | blocklist | domain | 123 | 3 | 2.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 3.3K | 1.8% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 503 | 1.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 62 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 3.3K | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 364 | 1.2% |
| phishing_army | blocklist | domain | 156.0K | 1.8K | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 125 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 2.2K | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 253 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 91 | 0.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 80 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| HaGeZi Pro | blocklist | domain | 225.2K | 983 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 169 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 31 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 251 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 7 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 431 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 28 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 13 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 10 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 180 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 567 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 41 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 138 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 30 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 31 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 66 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 11 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 13.0K | Targets: 53 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| YousList | blocklist | hostname | 625 | 108 | 17.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 12.9K | 14.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 421 | 11.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 46 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 123 | 4 | 3.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 9 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 351 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 189 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 7 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 2.5K | 1.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 572 | 0.9% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 278 | 0.9% |
| HaGeZi Pro | blocklist | domain | 225.2K | 1.5K | 0.6% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 995 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 549 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 125 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 90 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 30 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 15 | 0.1% |
| phishing_army | blocklist | domain | 156.0K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 21 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 36 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 84 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 15 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 78 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 10 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 14.0K | Targets: 6 | Unique: 14.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.2K | 10 | 0.8% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 445 | 2 | 0.4% |
| EasyList | blocklist | adguard | 67.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 18 | Unique: 141 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 70 | 7.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 26 | 2.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 4 | 1.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 502 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 470 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 4 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 15 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 54 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 51 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 38 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 5 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### Dogino_Discord_Official

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 4 | Unique: 7 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 8 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 14 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 7 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 7 | 0.2% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 7 | Unique: 332 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1.4K | 97.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 81 | 11.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 92 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 2 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 997 | 26.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 890 | 5.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 6 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 79 | 5.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 81 | 4.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 22 | 0.2% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 17.5K | Targets: 10 | Unique: 14.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 3.0K | 78.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 46 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.2K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 46 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 35 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 10 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 4 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.5K | Targets: 38 | Unique: 13.1K | Conflicts: 13</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 3.0K | 78.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 890 | 77.9% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 5 | 1.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 46 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 166 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 13 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 35 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 46 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 77 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 5 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 5.1K | 30.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 5.1K | 22.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 203 | 15.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 545 | 5.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2.3K | 4.5% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 1.8K | 2.5% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 8 | 2.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 300 | 2.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 360 | 1.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 68 | 1.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 17 | 1.5% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 6 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 15 | 0.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 4 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 41 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 2 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 2 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 8 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.2K | Targets: 21 | Unique: 13.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 93.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 28.8K | 16.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 75 | 5.7% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 2.6K | 4.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 43 | 3.6% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 5.6K | 2.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 336 | 1.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| abpvn_hosts | blocklist | adguard | 993 | 2 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 6 | 0.0% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 2 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 3 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 67.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.3K | 40.5K | 64.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 47.7K | 26.5% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 40.7K | 15.8% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 68 | 5.6% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 649 | 2.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 70 | 0.3% |
| abpvn_hosts | blocklist | adguard | 993 | 2 | 0.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 150 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 3 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 62 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 564 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 572 | 536 | 93.7% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 530 | 4.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 27 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 14 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 174 | 0.8% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 2 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 397 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 85 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 34 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 94 | 0.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 10 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 10 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 6 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 11 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.6% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 5 | 0.0% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 38 | Unique: 1.2K | Conflicts: 205</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_source_domain_allowlist | allowlist | domain | 43 | 14 | 32.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 14 | 32.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 500 | 129 | 25.8% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 56 | 7.9% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 27 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 42 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 11 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 6 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 16 | 0.1% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 17 | Unique: 4.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 94 | 8 | 8.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 38 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 25 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 8 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 52 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 16 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 36 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 256 | Targets: 10 | Unique: 201 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 959 | 32 | 3.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 4 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 2 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 10 | Unique: 475 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 9 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 7 | Unique: 2 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 5 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 21 | Unique: 944 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 67 | 2.5% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 13 | 1.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 2 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 11 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 4 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 35 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 21 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 105 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 80 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 16 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 16 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 11 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.7K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.2% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 90.7% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 21.6K | Targets: 29 | Unique: 0 | Conflicts: 248</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 3.8K | 87.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 982 | 86.7% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 301 | 84.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.3K | 69.0% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 198 | 34.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 7.7K | 34.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 174 | 30.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 397 | 30.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 5.9K | 28.5% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 18.3K | 25.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 2.3K | 21.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6.9K | 13.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 180 | 7.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 880 | 5.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 54 | 3.9% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 441 | 3.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 248 | 2.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 248 | 2.2% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 16 | 1.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 187 | 1.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 4 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 15 | 1.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 1 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 26 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 7 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.5K | Targets: 28 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 549 | 96.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 530 | 94.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8.0K | 53.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 7.7K | 46.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 494 | 37.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 3.1K | 14.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 7.0K | 13.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1.2K | 11.8% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 5.8K | 7.9% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 18 | 5.1% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 64 | 3 | 4.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 140 | 3.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 31 | 2.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 50 | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 35 | 2.6% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 441 | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 91 | 0.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 5 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 44 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 38 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 38 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 2 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 68 | Targets: 12 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 4 | 2 | 50.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 55 | 17.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 67 | 4.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 1 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 4 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 30 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 26 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 19 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 324 | Targets: 7 | Unique: 234 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 55 | 2.0% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 9 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 18 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 4 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.7K | Targets: 19 | Unique: 5.1K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 319 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 89 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2.3K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1.7K | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 225.2K | 1.9K | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 25 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 16 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 345 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 127 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 14 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 21 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 60 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 15 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 108 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 396 | 11.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 167 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 495 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 206 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1.5K | 0.9% |
| HaGeZi Pro | blocklist | domain | 225.2K | 1.6K | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.6K | 0.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 402 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 590 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 187 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 27 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 68 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.7% |
| Easy Privacy | blocklist | adguard | 55.2K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 1.5K | 0.9% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 1.6K | 0.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 68 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 27 | 0.1% |
| EasyList | blocklist | adguard | 67.1K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |

</details>

---

### GlobalAntiScamOrg-blocklist-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.2K | Targets: 18 | Unique: 7.5K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.3K | Targets: 27 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 355 | 129 | 36.3% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 3.8K | 17.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 171 | 15.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 275 | 14.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 585 | 5.6% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 3.6K | 5.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2.0K | 3.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 43 | 3.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 650 | 3.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 65 | 2.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 372 | 2.2% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 11 | 1.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 10 | 1.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 68 | 1.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 15 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 11 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 86 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 140 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 87 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 5 | 0.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 1 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 51 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 8 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 8 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 3 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 369 | Targets: 19 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 91 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 225.2K | 338 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 167 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 49 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 6 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 11 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 108 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 20 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 66 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 23 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 21 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 180.6K | Targets: 48 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 320 | 83.1% |
| phishing_army | blocklist | domain | 156.0K | 85.7K | 54.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 84.5K | 32.9% |
| kadantiscam | blocklist | domain | 43.0K | 13.6K | 31.6% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 75.4K | 30.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 12.3K | 26.8% |
| HaGeZi Pro | blocklist | domain | 225.2K | 47.0K | 20.9% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 7.1K | 18.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 13.8K | 15.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 32 | 13.7% |
| quidsup_notrack-malware | blocklist | domain | 123 | 10 | 8.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 977 | 5.4% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.6K | 5.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 23 | 4.6% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 3.3K | 4.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 202 | 4.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 244 | 3.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 13.2K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 939 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 7 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 48 | 1.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 55 | 0.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 78 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 162 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 36 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 50 | 0.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 6 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 207 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 53 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 67 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 46 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 7 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 34 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 42 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 9 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 221 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 5 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 307 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 18 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 103 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 180.6K | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 86.8K | 54.2% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 84.5K | 32.9% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 1.1K | 23.5% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 81.6K | 9.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1.6K | 5.4% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 7 | 0.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 6 | 0.4% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 207 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 67 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 46 | 0.3% |
| EasyList | blocklist | adguard | 67.1K | 150 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 5 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 221 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 11 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 997 | 87.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 3.0K | 17.1% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 27 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 62 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.8K | Targets: 6 | Unique: 805 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 3.0K | 17.1% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 3 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 466.1K | Targets: 40 | Unique: 453.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1.2K | 44.7% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2.3K | 7.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.9K | 4.4% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 2 | 0.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.1K | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 15 | 0.4% |
| HaGeZi Pro | blocklist | domain | 225.2K | 855 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 12 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 34 | 0.2% |
| kadantiscam | blocklist | domain | 43.0K | 98 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 138 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 307 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 21 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 7 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 77 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 23 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 9 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 630 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 33 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 121 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 43 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 236 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 38 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 214 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 21 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 388 | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.0K | 46 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 10 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 34 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 338 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 24 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 20 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 66 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 24 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 10 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 19 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 54 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 225.2K | Targets: 68 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 342 | 98.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 60.8K | 96.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 338 | 91.6% |
| hufilter | blocklist | hostname | 94 | 82 | 87.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 338 | 87.1% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 486 | 85.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 300 | 85.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.0K | 84.0% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-malware | blocklist | domain | 123 | 80 | 65.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 66 | 61.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.0K | 47.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 81.7K | 45.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.3K | 39.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 114 | 38.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 94.0K | 36.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 70.0K | 34.5% |
| YousList | blocklist | hostname | 625 | 201 | 32.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 8.7K | 31.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 12.6K | 28.6% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 26.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 47.0K | 26.0% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 14.9K | 16.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2.5K | 16.5% |
| kadantiscam | blocklist | domain | 43.0K | 6.6K | 15.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 4.4K | 14.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.9K | 12.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 47 | 12.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1.5K | 11.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| phishing_army | blocklist | domain | 156.0K | 12.0K | 7.7% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 9 | 3.9% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 8.5K | 3.8% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 8.6K | 3.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 243 | 3.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 135 | 2.7% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 989 | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 420 | 2.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| Spam404 | blocklist | domain | 8.1K | 159 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 65 | 1.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 62 | 1.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 185 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 47 | 1.4% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 983 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 31 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 8 | 1.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 250 | 1.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 470 | 1.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 166 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 540 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 630 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 49 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.6K | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 52 | 0.6% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 2.1K | 0.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 473 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 855 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 42 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 346 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 225.2K | 342 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 87 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 24 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 21 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.5K | Targets: 7 | Unique: 0 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 70.4% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 79 | 10.8% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 94 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 2 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 73.2K | Targets: 33 | Unique: 0 | Conflicts: 377</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 13.0K | 97.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 1.2K | 88.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 18.4K | 88.4% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 18.3K | 85.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 3.6K | 83.5% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 287 | 80.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.4K | 74.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 12.0K | 71.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 397 | 70.4% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 402 | 70.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 5.7K | 54.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 569 | 50.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.0K | 46.4% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 5.8K | 46.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.8K | 35.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 470 | 33.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 16.8K | 32.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 5.5K | 24.4% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 122 | 12.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 186 | 7.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 80 | 5.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 517 | 3.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 377 | 3.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 377 | 3.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 9 | 2.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 6 | 2.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 108 | 2.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 26 | 1.0% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 14 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 23 | 0.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 3 | 0.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 30 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 21 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 20 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 9 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 15 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 53 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 114 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 108 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 13 | 0.0% |

</details>

---

### hufilter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 94 | Targets: 24 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 87 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 12 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 85 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 18 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 82 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 72 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 6 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 55 | Unique: 431.6K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 959 | 13.1% |
| quidsup_notrack-malware | blocklist | domain | 123 | 12 | 9.8% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1.7K | 4.5% |
| hufilter | blocklist | hostname | 94 | 4 | 4.3% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 10.8K | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 8 | 3.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 1.0K | 2.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 9 | 2.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 75 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 2.4K | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 3.2K | 1.3% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2.6K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 567 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 98 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 2.4K | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1.4K | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 188 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 34 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 93 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 188 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 275 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 53 | 0.2% |
| phishing_army | blocklist | domain | 156.0K | 276 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 30 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 208 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 20 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 23 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 16 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 43.0K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 39.5K | 44.4% |
| phishing_army | blocklist | domain | 156.0K | 18.2K | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 13.6K | 7.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 12.5K | 4.9% |
| quidsup_notrack-malware | blocklist | domain | 123 | 4 | 3.3% |
| HaGeZi Pro | blocklist | domain | 225.2K | 6.6K | 2.9% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 6.0K | 2.4% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 590 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 3 | 1.3% |
| Spam404 | blocklist | domain | 8.1K | 23 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 57 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 68 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 41 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 8 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 19 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 41 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 62 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 98 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 25 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 29 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 97 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 30 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 174 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 126 | 2.2% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 1 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 46 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 29 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 23 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 18 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 26 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 180.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 3 | 0.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 24 | Targets: 4 | Unique: 0 | Conflicts: 25</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 24 | Targets: 4 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |

</details>

---

### local_domain_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 2 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 10 | Unique: 0 | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |

</details>

---

### local_social_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1 | Targets: 4 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### local_source_domain_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 2 | Unique: 27 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 14 | 0.8% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |

</details>

---

### local_source_ipv4_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 64 | Targets: 1 | Unique: 61 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.5K | 3 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 4.8K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 441 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 1.1K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 1.3K | 0.5% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 4.8K | 0.5% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 2 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 38.1K | Targets: 30 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 142 | 60.9% |
| phishing_army | blocklist | domain | 156.0K | 19.0K | 12.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 27.5K | 10.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 7.1K | 3.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 18 | 3.6% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 451 | 1.5% |
| kadantiscam | blocklist | domain | 43.0K | 590 | 1.4% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 1.8K | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 503 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 25 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.7K | 0.4% |
| HaGeZi Pro | blocklist | domain | 225.2K | 989 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 23 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 134 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 7 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 46 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 13 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 46 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 129 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 256.8K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.3K | 62.6K | 98.9% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.7% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.3% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 821 | 61.9% |
| EasyList | blocklist | adguard | 67.1K | 40.7K | 60.7% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 83.7K | 52.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 84.5K | 46.8% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 431 | 35.6% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 62.8K | 34.9% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 1.3K | 27.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 65 | 17.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 3.9K | 13.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 5.6K | 10.2% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 80.9K | 9.1% |
| abpvn_hosts | blocklist | adguard | 993 | 37 | 3.7% |
| CJX Annoyance | blocklist | adguard | 1.8K | 55 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 17 | 1.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 124 | 0.6% |
| AdBlockID | blocklist | adguard | 3.7K | 9 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 9 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 35 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 256.8K | Targets: 64 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 62.6K | 98.9% |
| hufilter | blocklist | hostname | 94 | 85 | 90.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 323 | 83.9% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 431 | 75.9% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 27.5K | 72.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 67.3% |
| phishing_army | blocklist | domain | 156.0K | 82.9K | 53.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 84.5K | 46.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 44.1% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 267 | 42.7% |
| HaGeZi Pro | blocklist | domain | 225.2K | 94.0K | 41.8% |
| WaLLy3K | blocklist | domain | 351 | 139 | 39.6% |
| quidsup_notrack-malware | blocklist | domain | 123 | 48 | 39.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6.5K | 35.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 62.8K | 35.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 79.0K | 31.6% |
| kadantiscam | blocklist | domain | 43.0K | 12.5K | 29.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 11.5K | 26.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 22.5K | 25.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 87 | 25.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 49.3K | 24.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 5.9K | 21.4% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 23 | 21.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.9K | 21.0% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.5% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2.5K | 19.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 54 | 13.9% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 3.9K | 13.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 541 | 10.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 36 | 9.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1.4K | 9.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 13 | 5.6% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 2.2K | 2.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 196 | 1.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 17 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 8 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 75 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 31 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 18 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 108 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 124 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 454 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 387 | 0.6% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1.2K | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 250 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 2.4K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 35 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 9 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1.1K | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 235 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 7 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 22.1K | Targets: 12 | Unique: 21.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 36 | 2.6% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 64 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| EasyList | blocklist | adguard | 67.1K | 70 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 84 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 122 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 67 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 124 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 4 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 23 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 22.1K | Targets: 40 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 6.6K | 10.8% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 7.1K | 9.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 14.8K | 6.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 36 | 2.6% |
| quidsup_notrack-malware | blocklist | domain | 123 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 27 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 10 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 48 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 327 | 0.2% |
| Torrent Trackers | blocklist | domain | 480 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 64 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 10 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 11 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 250 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 28 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 84 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 18 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 91 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 122 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 67 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 13 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 44 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 23 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 29 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 124 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 50 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 8 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 63.3K | Targets: 62 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 94 | 87 | 92.6% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 429 | 75.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.9% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 36.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 58.3K | 32.5% |
| HaGeZi Pro | blocklist | domain | 225.2K | 60.8K | 27.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4.8K | 25.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 62.6K | 24.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42.5K | 20.9% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 20 | 18.5% |
| quidsup_notrack-malware | blocklist | domain | 123 | 22 | 17.9% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 6.3K | 14.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 3.4K | 12.3% |
| Adaway | blocklist | hostname | 6.5K | 745 | 11.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 24 | 6.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 21 | 6.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 921 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.4K | 4.8% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 572 | 4.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 68 | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.6K | 4.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 6 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 79 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 216 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 84 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 938 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 60 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 262 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 11 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 169 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 3 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 207 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 25 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 188 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 77 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 46 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 175 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 76 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 63.3K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 837 | 63.1% |
| EasyList | blocklist | adguard | 67.1K | 40.5K | 60.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 429 | 35.5% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 58.3K | 32.4% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 62.6K | 24.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| Easy Privacy | blocklist | adguard | 55.2K | 2.6K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1.4K | 4.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 68 | 4.1% |
| abpvn_hosts | blocklist | adguard | 993 | 32 | 3.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 84 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdBlockID | blocklist | adguard | 3.7K | 4 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 207 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 93 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 1 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 2 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 233 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 142 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 21 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 3 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 8 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 50 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 3 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 32 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 13 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 3 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 73 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 109 | 73 | 67.0% |
| Torrent Trackers | blocklist | domain | 480 | 72 | 15.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 29.9K | Targets: 24 | Unique: 16.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 441 | 9.2% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.4% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 33 | 2.7% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 1.4K | 2.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 3.9K | 1.5% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 2.3K | 1.3% |
| EasyList | blocklist | adguard | 67.1K | 649 | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 1.6K | 0.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.2K | 336 | 0.6% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 600 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 64 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 46 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 1.8K | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 993 | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.8K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 29.9K | Targets: 76 | Unique: 0 | Conflicts: 163</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 686 | 19.4% |
| tranco | allowlist | domain_top | 500 | 79 | 15.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 60 | 15.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 569 | 13.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 46 | 9.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 21 | 9.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.5K | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1.3K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 123 | 10 | 8.1% |
| Adaway | blocklist | hostname | 6.5K | 520 | 8.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1.9K | 6.9% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 33 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 15 | 5.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 127 | 4.8% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 26 | 3.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1.4K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 9 | 2.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 278 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.3K | 2.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 4.4K | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.6K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 27 | 1.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 3.9K | 1.5% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2.3K | 1.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 58 | 1.3% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 451 | 1.2% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 2 | 1.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 461 | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 1.6K | 0.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 48 | 0.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 3.5K | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 2.3K | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 364 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 185 | 0.4% |
| phishing_army | blocklist | domain | 156.0K | 554 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 14 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 46 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 64 | 0.3% |
| Torrent Trackers | blocklist | domain | 480 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 21 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 75 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 92 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 172 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 57 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 59 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 93 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 75 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 156.0K | Targets: 36 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 19.0K | 49.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 85.7K | 47.4% |
| kadantiscam | blocklist | domain | 43.0K | 18.2K | 42.3% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 105.2K | 42.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 82.9K | 32.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 50 | 21.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 17.1K | 19.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 12.0K | 5.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 26 | 5.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1.8K | 2.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 554 | 1.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 36 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 30 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 1 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 276 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 31 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 9 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 769 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 101 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 8 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 11 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 20 | Unique: 61.7K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 94 | 6.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 92 | 4.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 19 | 0.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 11 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 3 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 14 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 19 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 142 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 76 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 46 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 46 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 296 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 145 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 298 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 300 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 123 | Targets: 25 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 59 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 80 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 48 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 28 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 10 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 20 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 6 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.3K | Targets: 55 | Unique: 0 | Conflicts: 52</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 576 | 16.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 473 | 11.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 34 | 8.8% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 434 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.3K | 4.2% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 905 | 3.3% |
| hufilter | blocklist | hostname | 94 | 3 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 10 | 2.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 319 | 2.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 11 | 1.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 6 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 921 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 189 | 1.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 73 | 1 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.2K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 996 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2.1K | 1.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2.5K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 109 | 1 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 25 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.4K | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 182 | 0.4% |
| Torrent Trackers | blocklist | domain | 480 | 2 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 599 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 9 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 23 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 12 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 67.1K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 892.4K | Targets: 15 | Unique: 617.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 4.8K | 99.9% |
| RPiList_specials-phishing | blocklist | adguard | 160.2K | 105.1K | 65.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 81.6K | 45.2% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 80.9K | 31.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1.8K | 5.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 5 | 0.4% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 93 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 242 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 15 | 0.1% |
| EasyList | blocklist | adguard | 67.1K | 62 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 23 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 6 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 160.2K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 86.8K | 48.1% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 83.7K | 32.6% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 105.1K | 11.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 600 | 2.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.8K | 10 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 8 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.9K | Targets: 22 | Unique: 0 | Conflicts: 13</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 355 | 88 | 24.8% |
| Greensnow | blocklist | ipv4 | 4.3K | 275 | 6.3% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 1.3K | 6.2% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 31 | 5.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 27 | 4.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 334 | 3.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 35 | 2.6% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 1.4K | 2.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 640 | 1.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 204 | 1.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 62 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 46 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 24 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 50 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 14 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 13 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 13 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 959 | Targets: 20 | Unique: 500 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 32 | 12.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 70 | 5.0% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 13 | 1.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 11 | 1.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 1 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 141 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 122 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 4 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 16 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 5 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 13 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 2 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 10 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 3 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 15 | Unique: 1.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 355 | 86 | 24.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 65 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 17 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 180 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 156 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 186 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 7 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 6 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.4K | Targets: 19 | Unique: 0 | Conflicts: 12.6K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 11.4K | 100.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 11.4K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 33 | 2.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 33 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 248 | 1.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 174 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 13 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 377 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 38 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 23 | 0.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 8 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 27 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 31 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 43 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.5K | Targets: 26 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 384 | 29.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 334 | 17.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 585 | 13.5% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 46 | 13.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 2.3K | 10.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 545 | 10.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.4K | 9.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 1.9K | 9.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 4.4K | 8.5% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 5.7K | 7.8% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 39 | 6.8% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 828 | 6.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 34 | 6.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 973 | 5.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 1.2K | 5.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 60 | 5.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 1 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 21 | 1.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 191 | 1.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 97 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 3 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 23 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 23 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.4K | Targets: 18 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.2K | 12 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 36 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 17 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 36 | 0.2% |
| EasyList | blocklist | adguard | 67.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 10 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 712 | Targets: 38 | Unique: 335 | Conflicts: 304</summary>

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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 2 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 40 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 2 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 17 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 26 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 11 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 9 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 12 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 31 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 27.7K | Targets: 53 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 495 | 30.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 966 | 27.3% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 105 | 18.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 672 | 15.7% |
| WaLLy3K | blocklist | domain | 351 | 53 | 15.1% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 94 | 11 | 11.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.7K | 9.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 7 | 6.5% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.9K | 6.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 24 | 6.2% |
| Adaway | blocklist | hostname | 6.5K | 408 | 6.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 905 | 5.9% |
| quidsup_notrack-malware | blocklist | domain | 123 | 7 | 5.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3.4K | 5.4% |
| tranco | allowlist | domain_top | 500 | 23 | 4.6% |
| HaGeZi Pro | blocklist | domain | 225.2K | 8.7K | 3.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 6.1K | 3.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.1K | 3.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 351 | 2.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 5.9K | 2.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.9K | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 11 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 123 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 48 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 161 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 162 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 14 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 204 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 34 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 33 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 13 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 53 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 6 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 220.2K | Targets: 33 | Unique: 162.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 14.8K | 67.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.5K | 28.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 75 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 442 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 473 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 76 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 22 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 235 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 66 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 50 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 189 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 103 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 15 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 106 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 125 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 31 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 3 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 8 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 208 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 21 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 44.0K | Targets: 43 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 123 | 59 | 48.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 6.3K | 9.9% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 38 | 6.7% |
| HaGeZi Pro | blocklist | domain | 225.2K | 12.6K | 5.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 919 | 5.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 11.5K | 4.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 6.0K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.0K | 3.0% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 55 | 1.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 182 | 1.2% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 90 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 185 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 939 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 312 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 253 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 1 | 0.3% |
| kadantiscam | blocklist | domain | 43.0K | 68 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 195 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 16 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 28 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 366 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 38 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 13 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 11 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.3K | Targets: 29 | Unique: 4.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 959 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 27 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 144 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 243 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 244 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 75 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 12 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 6 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 5.8K | Targets: 25 | Unique: 1.2K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 3 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 48 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 23 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 23 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 8 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 55 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 49 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 9 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 2 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 30 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 43 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.1K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 6.6K | 29.9% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 19.4K | 8.8% |
| pexcn Torrent Trackers | blocklist | domain_url | 73 | 2 | 2.7% |
| Torrent Trackers | blocklist | domain | 480 | 9 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 109 | 2 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 615 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 123 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 75 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 216 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 275 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 540 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 23 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 387 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 44 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 9 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 11 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 9 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 34 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 36 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 14 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.6K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 2.6K | 3.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 127 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1.2K | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 31 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.2K | Targets: 33 | Unique: 0 | Conflicts: 85</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3.2K | 85.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 17 | 2.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 27 | 1.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 23 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 32 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 25 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 23 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 28 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 47 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 31 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 11 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 14 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 52 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 30 | Unique: 6.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 123 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 1.6K | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 20 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 53 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 41 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 159 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 23 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 7 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 12 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.7K | 98.8% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.5% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 17.6K | 7.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 4.3K | 1.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 3 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 977 | 0.5% |
| HaGeZi Pro | blocklist | domain | 225.2K | 420 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 80 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 387 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 19 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 88.9K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2.6K | 100.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.5K | 99.7% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 12.9K | 99.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 367 | 95.3% |
| kadantiscam | blocklist | domain | 43.0K | 39.5K | 91.7% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 50.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.6K | 41.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 402 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 66 | 17.0% |
| quidsup_notrack-malware | blocklist | domain | 123 | 20 | 16.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 49 | 13.3% |
| hufilter | blocklist | hostname | 94 | 12 | 12.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 899 | 12.3% |
| phishing_army | blocklist | domain | 156.0K | 17.1K | 11.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 22.5K | 8.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1.2K | 8.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 13.8K | 7.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1.9K | 7.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 14.9K | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3.6K | 5.7% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 32 | 5.6% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.6K | 5.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 31 | 4.4% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 7.6K | 3.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 4.9K | 2.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 345 | 2.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 3 | 1.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 29 | 0.9% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 3.9K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Spam404 | blocklist | domain | 8.1K | 53 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 312 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 134 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 43 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 813 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 251 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 44 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 25 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 428 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 8 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 20 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 106 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 188 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 45 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.1K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7.1K | 32.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 21.5K | 9.8% |
| pexcn Torrent Trackers | blocklist | domain_url | 73 | 2 | 2.7% |
| hufilter | blocklist | hostname | 94 | 2 | 2.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 73 | 2.1% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 480 | 9 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 109 | 2 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 133 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 157 | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 703 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 161 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 262 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 92 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 225.2K | 630 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 344 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 13 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 454 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 26 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 50 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 9 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 42 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 31 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 13 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 43 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 34 | Unique: 0 | Conflicts: 91</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 3.2K | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 42 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 15 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 17 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 35 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 65 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 29 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 45.7K | Targets: 29 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 12.3K | 6.8% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 4.3K | 5.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 26.7K | 5.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 7 | 1.8% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 461 | 1.5% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 2 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 1.1K | 0.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 12 | 0.2% |
| HaGeZi Pro | blocklist | domain | 225.2K | 470 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 46 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 250 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 167 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 20 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 23 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 3 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 31 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 516.1K | Targets: 48 | Unique: 427.7K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 26.7K | 58.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 385 | 218 | 56.6% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 24.8K | 30.8% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 3.5K | 11.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 441 | 8.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 13.2K | 7.3% |
| quidsup_notrack-malware | blocklist | domain | 123 | 6 | 4.9% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 8.2K | 3.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 387 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 6 | 1.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 2.1K | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 2.4K | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 366 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| phishing_army | blocklist | domain | 156.0K | 769 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 428 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 77 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 23 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 736 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 233 | 1 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 175 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 36 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 129 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 50 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 6 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 282 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 630 | 0.1% |
| kadantiscam | blocklist | domain | 43.0K | 62 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 36 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 125 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 34 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20.8K | Targets: 27 | Unique: 0 | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 569 | 50.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.2K | 28.2% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 5.9K | 27.5% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 18.4K | 25.1% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 3.1K | 24.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 1.9K | 17.8% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 99 | 17.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 85 | 15.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 650 | 15.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 172 | 13.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 204 | 10.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 5.1K | 10.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 2.1K | 9.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 360 | 7.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 896 | 5.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 38 | 2.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 191 | 1.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 16 | 1.2% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 10 | 1.0% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 3 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 58 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 1 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 480 | Targets: 9 | Unique: 276 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 109 | 108 | 99.1% |
| pexcn Torrent Trackers | blocklist | domain_url | 73 | 72 | 98.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 46 | Unique: 0 | Conflicts: 561</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 8 | 18.6% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 129 | 7.7% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 43 | 2 | 4.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 94 | 2 | 2.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 2 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 31 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 27 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 27 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 29 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 79 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 38 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 35 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.8K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 23 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 11 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 36 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 30 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 16 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 28 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 67.1K | 51 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 180.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 30 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 892.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 31 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 11 | Unique: 25 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.2K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 839 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 43 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 53 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 7 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 65 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 385 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 367 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 320 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 60 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 243 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 323 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 218 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 7 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 47 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 60.3K | Targets: 1 | Unique: 60.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 4 | 4.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 13.4K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 13.0K | 17.8% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 191 | 1.8% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 7 | 1.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 7 | 1.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 191 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 3 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 88 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 7 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 217 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 44 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 20 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 10 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 2 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 1 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 12 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 26 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 3 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.4K | Targets: 34 | Unique: 13.2K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 279 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 51 | 3.7% |
| BlockListDE_Strong | blocklist | ipv4 | 355 | 9 | 2.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 564 | 10 | 1.8% |
| BruteforceBlocker | blocklist | ipv4_find | 572 | 10 | 1.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 51 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.4K | 16 | 1.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.3K | 13 | 1.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.5K | 97 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 187 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 959 | 9 | 0.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 41 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 416 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 517 | 0.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 14 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 88 | 0.7% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 48 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 54 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 1 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 91 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 54 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 58 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 6 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 2 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 2 | 0.1% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 4 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 4 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 2.1K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 541 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 441 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 62 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 202 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 135 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 12 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 15 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 279 | 1.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 256 | 1 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 10 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 108 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.7K | 4 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.5K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.8K | 3 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.8K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.6K | 2 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 97 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 60.3K | 4 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 6 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.5K | 45 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 73.2K | 45 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 5 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 351 | Targets: 32 | Unique: 0 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| YousList | blocklist | hostname | 625 | 9 | 1.4% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 123 | 1 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 83 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 19 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 3 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 53 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 35 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 85 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 19 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 139 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 84 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 35 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 20 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 65 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.6K | 74.2% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3.0K | 69.9% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6.3K | 34.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 3.6K | 28.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 20.7K | 23.3% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 94 | 14 | 14.9% |
| quidsup_notrack-malware | blocklist | domain | 123 | 14 | 11.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 187 | 11.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 41 | 10.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 996 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 1.8K | 6.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 15.9K | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 3.1K | 4.9% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 1.4K | 4.6% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 25 | 4.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 27 | 3.8% |
| HaGeZi Pro | blocklist | domain | 225.2K | 6.8K | 3.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 3.2K | 1.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 38 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 127 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 161 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 335 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 29 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 15 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 43 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 21 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 30 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 50 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 18 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |

</details>

---

### YousList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 625 | Targets: 32 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 9 | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 199 | 1.1% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 3 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 108 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 24 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 2 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 86 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 267 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 94 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 152 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 42 | 0.1% |
| HaGeZi Pro | blocklist | domain | 225.2K | 201 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 21 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 17 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 1 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 7 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| AdGuard DNS filter | blocklist | adguard | 180.1K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 10 | 0.0% |
| EasyList | blocklist | adguard | 67.1K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 256.8K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.3K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 29.9K | 2 | 0.0% |

</details>

---

### youtube_GoodbyeAds

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 23 | Unique: 97.2K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| WaLLy3K | blocklist | domain | 351 | 7 | 2.0% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 50 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 74 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| HaGeZi Pro | blocklist | domain | 225.2K | 42 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 8 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 45 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.5K | Targets: 60 | Unique: 0 | Conflicts: 44</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.8K | 42.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 396 | 24.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2.0K | 11.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 123 | 9 | 7.3% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.5K | 4.0% |
| Adaway | blocklist | hostname | 6.5K | 259 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.3K | 576 | 3.8% |
| YousList | blocklist | hostname | 625 | 24 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 4 | 3.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.7K | 966 | 3.5% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 421 | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.3K | 1.6K | 2.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 10 | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 568 | 13 | 2.3% |
| ph00lt0_blocklist | blocklist | domain | 29.9K | 686 | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 346 | 5 | 1.4% |
| HaGeZi Pro | blocklist | domain | 225.2K | 3.0K | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 256.8K | 2.4K | 0.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.4K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 44.0K | 55 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 16 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 45.7K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.0K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 30 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 80.7K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 220.2K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 38.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 250.1K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 516.1K | 6 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 180.6K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 466.1K | 15 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 9 | 0.0% |
| phishing_army | blocklist | domain | 156.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 1 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| USOM-Blocklists-ips | blocklist | ipv4 | 15.4K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

