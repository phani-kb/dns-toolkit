# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-09-19 13:37:52 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 165 |
| Total Entries Analyzed | 9.1M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 143 |
| allowlist | 22 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 34 |
| cidr_ipv4 | 3 |
| domain | 88 |
| ipv4 | 40 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 70 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 68.3% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 36.1K | 64.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 11.2K | 60.9% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 303 | 51.4% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 168 | 45.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 141 | 36.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 590 | 35.9% |
| hufilter | blocklist | hostname | 94 | 31 | 33.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 32.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 110 | 31.9% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 34 | 31.5% |
| HaGeZi Pro | blocklist | domain | 229.9K | 68.2K | 29.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 49.4K | 27.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6.2K | 22.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 42.9K | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 13.2K | 14.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 4.3K | 13.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6.0K | 13.1% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 16 | 2.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 144 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 3.3K | 1.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 301 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 26 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1.3K | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 437 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 161 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 43 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 444 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 448 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 90 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 736 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 219 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 47 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 624 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.1K | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| phishing_army | blocklist | domain | 150.6K | 100 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 214 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 984 | Targets: 8 | Unique: 882 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 33 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 38 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 24 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 2 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 51 | Unique: 0 | Conflicts: 37</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2.7K | 14.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 262 | 7.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 6.5K | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 6 | 5.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 21 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 434 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 9 | 2.3% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 10 | 1.7% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 520 | 1.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 408 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 11 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 752 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.7K | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 87 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1.0K | 0.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 14 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 26 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 50 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 93 | Targets: 8 | Unique: 32 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 207 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 35 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 589 | Targets: 32 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 439 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 106 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 53 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 571 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 504 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 440 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 303 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 34 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 40 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 32 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 8 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 5 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.2K | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.4K | 439 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 590 | 0.3% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 440 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 43 | 0.1% |
| EasyList | blocklist | adguard | 67.8K | 69 | 0.1% |
| abpvn_hosts | blocklist | adguard | 984 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 34 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 10 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 5 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 4 | 0.0% |

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
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 224.8K | Targets: 24 | Unique: 116.4K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 89.7K | 49.9% |
| hufilter | blocklist | hostname | 94 | 18 | 19.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.3K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 587 | 3.9% |
| HaGeZi Pro | blocklist | domain | 229.9K | 8.8K | 3.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 929 | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 813 | 0.9% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 30 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 204 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 84 | 0.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 2 | 0.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.2K | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 335 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 59 | 0.2% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 180.6K | Targets: 25 | Unique: 0 | Conflicts: 196</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 199 | 195 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 93.1% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 51.3K | 90.9% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 67.8K | 48.3K | 71.3% |
| Easy Privacy | blocklist | adguard | 55.3K | 28.8K | 52.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 590 | 47.7% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 55.9K | 22.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 2.3K | 7.5% |
| abpvn_hosts | blocklist | adguard | 984 | 24 | 2.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.7K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 122 | 0.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 215 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 181 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 8 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 179.9K | Targets: 69 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 589 | 571 | 96.9% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 93.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 51.3K | 90.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 94 | 72 | 76.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.3% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 89.7K | 39.9% |
| HaGeZi Pro | blocklist | domain | 229.9K | 80.5K | 35.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5.0K | 27.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 49.4K | 24.3% |
| YousList | blocklist | hostname | 625 | 151 | 24.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 55.9K | 22.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 966 | 22.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6.0K | 21.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 21 | 19.4% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6.0K | 13.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.7K | 11.7% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2.3K | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 24 | 7.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 4.9K | 5.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 19 | 4.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 553 | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 122 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 274 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 343 | 0.4% |
| Torrent Trackers | blocklist | domain | 488 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 183 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 275 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 276 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 29 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 215 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 641 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 188 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 108 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 6 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 120 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 5 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 365 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 8 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.4K | 836 | 1.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 820 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 70 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 76 | 0.1% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 199 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 207 | 195 | 94.2% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 10 | Unique: 4.0K | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_issues

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 6 | Unique: 59 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2 | 0.0% |

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
<summary>List Type: allowlist | Source Type: domain | Total: 181 | Targets: 12 | Unique: 154 | Conflicts: 15</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 1 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.7% |
| Easy Privacy | blocklist | adguard | 55.3K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 2.1K | 0.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 25 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 2 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.4K | Targets: 57 | Unique: 0 | Conflicts: 48</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.0K | 56.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.8% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.7% |
| YousList | blocklist | hostname | 625 | 199 | 31.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 91 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 206 | 12.5% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 12 | 11.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 53 | 9.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 7.6K | 8.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.3K | 8.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 4.7K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.8% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.7K | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.2K | 5.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 20 | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.5K | 4.9% |
| HaGeZi Pro | blocklist | domain | 229.9K | 7.3K | 3.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 5.0K | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 6.4K | 2.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 14 | 2.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 926 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 52 | 0.4% |
| Torrent Trackers | blocklist | domain | 488 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 19 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 43 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 25 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 28 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 12 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 9 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 50 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 142 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 34 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.8K | Targets: 22 | Unique: 0 | Conflicts: 13</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 704 | 13.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 949 | 9.2% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 54 | 7.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 51 | 7.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.1% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 26 | 7.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 1.5K | 6.6% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 4.3K | 5.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 1.1K | 5.7% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 624 | 5.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 3.1K | 5.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 762 | 3.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 56 | 3.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 125 | 2.5% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 278 | 1.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 36 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 13 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 15 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 13 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 51 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.3K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 5.0K | 224 | 4.4% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 851 | 3.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 34 | 2.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 575 | 2.8% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 698 | 1.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 174 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 8 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 1 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 58 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 285 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 1 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 11 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 20 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 28 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 369 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 96 | 5.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 87 | 3.4% |
| Greensnow | blocklist | ipv4 | 5.0K | 157 | 3.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 312 | 1.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 26 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 37 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 257 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 3 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 227 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 2 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 39 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 20 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 28 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 12 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 3 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 40 | Unique: 10.1K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 45 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 79 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 193 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 22 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 184 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 32 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 53 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 51 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 49 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 7 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 264 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 126 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 17 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 41 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 246.3K | Targets: 51 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 150.6K | 107.5K | 71.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 266 | 69.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 74.2K | 40.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 79.4K | 32.2% |
| kadantiscam | blocklist | domain | 42.9K | 12.0K | 28.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 15.4K | 17.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 2.1K | 5.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 10.9K | 4.7% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 3.1K | 3.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 16.8K | 1.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18.9K | 1.6% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 634 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 3 | 1.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 6.2K | 1.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 201 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 195 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 16 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 262 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 45 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 108 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 219 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 28 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 57 | Unique: 0 | Conflicts: 71</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.8K | 51.5% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 48 | 13.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.9K | 10.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 167 | 10.2% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 40 | 5.6% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 514 | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 13 | 3.7% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 473 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1.6K | 2.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 11 | 2.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 672 | 2.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 2.2K | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 570 | 1.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 6 | 1.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2.0K | 0.9% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 5 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.9K | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 966 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 42 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 19 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 24 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 10 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 34 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 5 | 2.1% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 59 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 23 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 49 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 75 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 15 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 55 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 36 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 11 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 60.9K | Targets: 32 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 369 | 257 | 69.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 3.1K | 65.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.4K | 62.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 673 | 58.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 2.9K | 56.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.0K | 55.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 373 | 55.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 5.6K | 54.8% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 6.2K | 54.3% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 376 | 54.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.8K | 54.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 7.4K | 36.6% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 8.8K | 35.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 8.1K | 35.0% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 21.3K | 29.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 4.3K | 22.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 285 | 21.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 212 | 19.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 20 | 7.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 85 | 6.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 6 | 3.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 411 | 2.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 232 | 1.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 38 | 1.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 42 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 420 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 8 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 22 | Unique: 208 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 2 | 0.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 695 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 657 | 97.3% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 675 | 5.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 101 | 5.6% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 306 | 1.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 59 | 1.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 54 | 1.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 152 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 520 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 376 | 0.6% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 2 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 38 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 19 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 8 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 11 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 6 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 101 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 67 | 94.4% |
| Torrent Trackers | blocklist | domain | 488 | 100 | 20.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 20 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.4K | 7.7K | 67.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 4.7K | 23.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 1.1K | 22.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 4.5K | 19.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 888 | 17.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 1.6K | 15.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 9.4K | 15.4% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 9.5K | 13.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 1.5K | 7.8% |
| Greensnow | blocklist | ipv4 | 5.0K | 78 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 335 | 1.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 18 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 5 | 0.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 9 | 0.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 4 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 57 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 35 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 35 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 26 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 7 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 984 | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 56 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 9 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 87.8K | Targets: 48 | Unique: 1.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 4.7K | 10.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 25 | 6.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 29.3K | 5.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 7 | 3.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 3.5K | 1.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 21.8K | 1.8% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 15.2K | 1.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 68 | 1.4% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3.1K | 1.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 516 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 359 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 149 | 1.1% |
| phishing_army | blocklist | domain | 150.6K | 1.7K | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 2.2K | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 256 | 0.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 85 | 0.5% |
| HaGeZi Pro | blocklist | domain | 229.9K | 970 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 170 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 32 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 276 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 14 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 28 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 7 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 437 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 52 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 183 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 581 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 26 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 15 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 31 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 69 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 34 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 190 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 13.1K | Targets: 55 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| YousList | blocklist | hostname | 625 | 108 | 17.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 13.0K | 14.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 46 | 11.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 422 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 8 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 9 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 355 | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 578 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 2.5K | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 279 | 0.9% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.5K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 992 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 553 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 149 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 94 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 690 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 29 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 14 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 635 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 9 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 8 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 84 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 36 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 22 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 79 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 15 | 0.0% |

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
| EasyList | blocklist | adguard | 67.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 115 | 10.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 34 | 2.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 673 | 1.1% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 480 | 0.7% |
| Greensnow | blocklist | ipv4 | 5.0K | 24 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 61 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 80 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 51 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 10 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 22 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 6 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

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
| Sinfonietta_Social | blocklist | hostname | 3.8K | 7 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 7 | 0.2% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 8 | Unique: 348 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 96.3% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 81 | 11.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 93 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1.0K | 30.9% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 891 | 5.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 7 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 79 | 5.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 81 | 4.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 40 | Unique: 11.3K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3.0K | 91.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 891 | 78.2% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 47 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 162 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 526 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 13 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 45 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 15 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 35 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 63 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 10 | Unique: 13.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3.0K | 91.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 47 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 4 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 12 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 45 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 35 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 5.1K | 27.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 5.1K | 22.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 704 | 14.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 782 | 7.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 888 | 5.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 2.8K | 4.5% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 2.5K | 3.4% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 8 | 2.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 455 | 2.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 81 | 1.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 15 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 18 | 1.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 11 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 41 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 13 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 59 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.3K | Targets: 21 | Unique: 13.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 28.8K | 16.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 76 | 5.7% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 2.6K | 4.7% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 43 | 3.5% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 5.7K | 2.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 337 | 1.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| abpvn_hosts | blocklist | adguard | 984 | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 2 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 13 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 67.8K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.4K | 33.7K | 59.7% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 48.3K | 26.8% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 33.9K | 13.8% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 69 | 5.6% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 646 | 2.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 71 | 0.3% |
| abpvn_hosts | blocklist | adguard | 984 | 2 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 146 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 47 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 675 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 695 | 657 | 94.5% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 642 | 5.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 98 | 5.4% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 280 | 1.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 56 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 51 | 1.1% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 3 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 507 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 138 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 373 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 34 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 7 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 17 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 10 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 9 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.7% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 5 | 0.0% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 40 | Unique: 1.2K | Conflicts: 212</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 14 | 32.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 13 | 31.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 500 | 122 | 24.4% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 56 | 7.9% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 6 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 42 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 7 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 19 | Unique: 4.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 94 | 8 | 8.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 38 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 16 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 36 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 16 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 51 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 15 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 159 | Targets: 11 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 30 | 2.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 54 | 1 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 3 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 6 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 3 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 2 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 10 | Unique: 469 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 11 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 6 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 22 | Unique: 820 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 3 | 1.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 54 | 1 | 1.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 19 | 1.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 5 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.0K | 13 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 18 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 85 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 17 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 65 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 18 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 19 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 34 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 13 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 6 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 4 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 1 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.7K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.5% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.0% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 24.6K | Targets: 29 | Unique: 0 | Conflicts: 2.1K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 4.6K | 90.3% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 312 | 84.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.3K | 69.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 851 | 64.0% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 306 | 44.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 280 | 41.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 7.4K | 36.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 8.3K | 35.6% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 16.1K | 22.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 1.1K | 22.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 2.2K | 21.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 2.1K | 18.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 2.1K | 18.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8.8K | 14.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.5K | 9.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 188 | 7.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 80 | 6.9% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 596 | 5.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 29 | 2.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 3 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 19 | 1.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 208 | 1.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 117 | 0.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 2 | 0.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 9 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 322 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 7 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 23.2K | Targets: 29 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 675 | 97.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 642 | 95.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.7K | 51.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 8.3K | 44.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 1.5K | 31.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 1.6K | 15.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8.1K | 13.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 2.7K | 13.2% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 7.4K | 10.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 112 | 6.2% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 20 | 5.4% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Greensnow | blocklist | ipv4 | 5.0K | 237 | 4.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 34 | 2.8% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 596 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 28 | 2.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 134 | 1.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 116 | 0.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 38 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 38 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 474 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 5 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 13 | Unique: 2.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 56 | 20.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 72 | 5.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 3 | 1.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 7 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 38 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 9 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 7 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 23 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 279 | Targets: 8 | Unique: 187 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 56 | 2.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 20 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 2 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.7K | Targets: 19 | Unique: 5.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 315 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 87 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1.7K | 1.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2.3K | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.9K | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 16 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 342 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 129 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 59 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 14 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 14 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 19 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 109 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 392 | 11.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 167 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 497 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 206 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1.5K | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.6K | 0.7% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 402 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 590 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 188 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 27 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 66 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 10 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.7% |
| Easy Privacy | blocklist | adguard | 55.3K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 1.5K | 0.9% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 27 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |

</details>

---

### GlobalAntiScamOrg-blocklist-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.2K | Targets: 21 | Unique: 7.5K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 13 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.0K | Targets: 29 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 369 | 157 | 42.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 591 | 32.7% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 4.6K | 18.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 224 | 16.9% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 59 | 8.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 56 | 8.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 1.3K | 6.3% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 4.5K | 6.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 2.9K | 4.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 451 | 4.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 106 | 4.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 125 | 2.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 24 | 2.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 314 | 1.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 81 | 1.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 2 | 1.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 13 | 1.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 237 | 1.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 94 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 7 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 78 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 71 | 0.5% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 5 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 371 | Targets: 19 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 91 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 49 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 168 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 340 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 37 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 20 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 6 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 108 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 21 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 66 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 23 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 21 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 183.8K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 81.3K | 52.1% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 80.3K | 32.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 1.1K | 31.8% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 77.4K | 12.9% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1.9K | 6.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 8 | 0.6% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 5 | 0.4% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 197 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 45 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 38 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 5 | 0.2% |
| EasyList | blocklist | adguard | 67.8K | 146 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 215 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 2 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 183.8K | Targets: 52 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 320 | 83.6% |
| phishing_army | blocklist | domain | 150.6K | 80.5K | 53.4% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 14.1K | 32.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 80.3K | 32.6% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 74.2K | 30.1% |
| kadantiscam | blocklist | domain | 42.9K | 12.3K | 28.7% |
| HaGeZi Pro | blocklist | domain | 229.9K | 46.7K | 20.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 7.6K | 19.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 43 | 18.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 12.9K | 14.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 11 | 8.8% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.9K | 6.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 980 | 5.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 23 | 4.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 204 | 4.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 3.5K | 4.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 246 | 3.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 16.6K | 3.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 866 | 1.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 49 | 1.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 58 | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 5 | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.3K | 0.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7.5K | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 171 | 0.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 79 | 0.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 8 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.5K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 36 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 4.1K | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 51 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 197 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 45 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 50 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 7 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 5 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 38 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 35 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 594 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 12 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 215 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 107 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.3K | Targets: 13 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 63 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 9 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.3K | Targets: 6 | Unique: 274 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.6% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 9 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 522.4K | Targets: 43 | Unique: 508.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2.3K | 7.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.9K | 4.4% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 3 | 0.5% |
| HaGeZi Pro | blocklist | domain | 229.9K | 839 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 15 | 0.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.1K | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 1 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 12 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 594 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.5K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 190 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 34 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 104 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 22 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 43 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 7 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 664 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 214 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 39 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 37 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 78 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 262 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 120 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 29 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 27 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 21 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 330 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 248 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 388 | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 46 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 10 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 20 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 24 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 141 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 338 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 66 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 54 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 24 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 19 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 10 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 229.9K | Targets: 70 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 341 | 98.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 54.0K | 95.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 340 | 91.6% |
| hufilter | blocklist | hostname | 94 | 82 | 87.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 338 | 87.1% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 504 | 85.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.0K | 83.7% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 80 | 64.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 66 | 61.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.0K | 47.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 80.5K | 44.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.3K | 39.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 86.1K | 35.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 68.2K | 33.6% |
| YousList | blocklist | hostname | 625 | 202 | 32.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 8.7K | 31.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 12.3K | 26.7% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 26.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 46.7K | 25.4% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.5K | 16.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 14.4K | 16.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 57 | 14.9% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 4.5K | 14.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| kadantiscam | blocklist | domain | 42.9K | 6.1K | 14.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.9K | 12.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.5K | 11.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| phishing_army | blocklist | domain | 150.6K | 11.4K | 7.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 17 | 7.3% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 10.9K | 4.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 8.8K | 3.9% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 240 | 3.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 137 | 2.8% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1.0K | 2.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 428 | 2.4% |
| Spam404 | blocklist | domain | 8.1K | 191 | 2.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 63 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 65 | 1.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 65 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 184 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 254 | 1.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 970 | 1.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 8 | 1.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 162 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 539 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 25 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 630 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.6K | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 51 | 0.6% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.6K | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1.8K | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 114 | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 2.1K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 470 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 839 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 42 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 345 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 341 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 21 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 24 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 87 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 21 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 8 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 69.1% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 79 | 10.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 95 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 72.3K | Targets: 34 | Unique: 0 | Conflicts: 4.7K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 13.2K | 97.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 4.5K | 89.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 18.1K | 89.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 4.3K | 89.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 507 | 75.1% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 520 | 74.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.3K | 74.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 13.8K | 73.5% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 16.1K | 65.5% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 7.4K | 65.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.5K | 63.5% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 227 | 61.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 6.0K | 58.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 698 | 52.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.5K | 47.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 480 | 41.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 4.7K | 41.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 4.7K | 41.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 21.3K | 34.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 6.9K | 29.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 174 | 16.3% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 7 | 15.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 144 | 5.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 65 | 5.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 6 | 3.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 10 | 3.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 405 | 2.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 23 | 0.9% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 18 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 616 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 16 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 4 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 23 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 30 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 20 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 53 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 21 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 108 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 13 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 15 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 108 | 0.0% |

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
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 87 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 18 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 85 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 82 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 72 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 12 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 58 | Unique: 421.9K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 960 | 13.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 10.8K | 4.4% |
| hufilter | blocklist | hostname | 94 | 4 | 4.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1.7K | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 6 | 2.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 9 | 2.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 75 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 2.5K | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 3.2K | 1.3% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2.6K | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 581 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 98 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 2.4K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 4.4K | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.2K | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 186 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 105 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1.5K | 0.3% |
| phishing_army | blocklist | domain | 150.6K | 277 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 194 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 52 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 275 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 208 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 49 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 24 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 21 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 17 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 42.9K | Targets: 42 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 35.5K | 39.8% |
| phishing_army | blocklist | domain | 150.6K | 17.9K | 11.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 12.3K | 6.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 12.6K | 5.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 12.0K | 4.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 6.1K | 2.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 5 | 2.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 801 | 2.0% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 29 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 58 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 19 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 69 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 22 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 44 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 49 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 90 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 30 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 24 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 104 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 29 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 26 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 23 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 33 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 175 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 126 | 2.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 4 | 1.7% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 46 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 23 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 4 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 26 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 29 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 180.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 3 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |

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
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 11 | Unique: 0 | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 3 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.6K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 465 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 1.1K | 0.6% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 3.6K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 1.3K | 0.5% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 2 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2M | Targets: 62 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 975.4K | 98.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 270.9K | 51.2% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 20.6K | 47.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 21.8K | 24.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 41 | 10.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 3.7K | 8.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 18.9K | 7.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 11 | 6.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.8K | 5.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 690 | 5.3% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 7.5K | 4.1% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 526 | 3.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 6.4K | 2.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 4.6K | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 8 | 1.4% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.3K | 1.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 286 | 1.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 2 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 426 | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.1K | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 16 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 25 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 3 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 641 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 25 | 0.4% |
| phishing_army | blocklist | domain | 150.6K | 521 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 96 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 119 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 330 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 7 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 30 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 29 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 34 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 215.2K | Targets: 38 | Unique: 202.7K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 5.3K | 33.9% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 474 | 2.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 253 | 1.9% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 7 | 1.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 22 | 1.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 322 | 1.7% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 59 | 1.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 51 | 1.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 19 | 1.1% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 616 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 9 | 0.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 5 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 145 | 0.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 5 | 0.7% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 10 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 420 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 1 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 59 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 10 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 54 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 7 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Greensnow | blocklist | ipv4 | 5.0K | 17 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 81 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 26 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 48 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 990.5K | Targets: 50 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 975.4K | 81.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 221.9K | 41.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 15.2K | 17.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 35 | 9.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2.9K | 6.8% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 16.8K | 6.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 690 | 2.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 4.1K | 2.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 4.2K | 1.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 4 | 1.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.0K | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2.1K | 0.9% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.4K | 0.9% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 109 | 0.3% |
| phishing_army | blocklist | domain | 150.6K | 404 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 624 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 114 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 365 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 16 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 5 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 23 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 18 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 61 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 15 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 9 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 248 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 39.7K | Targets: 33 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 81 | 34.6% |
| phishing_army | blocklist | domain | 150.6K | 20.4K | 13.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 28.2K | 11.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 7.6K | 4.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 17 | 3.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 625 | 2.0% |
| kadantiscam | blocklist | domain | 42.9K | 801 | 1.9% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 2.1K | 0.9% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 516 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 23 | 0.5% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.0K | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.7K | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 19 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 125 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 41 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 130 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 96 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 109 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 47 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 4 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 246.3K | Targets: 66 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 98.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 55.7K | 98.7% |
| hufilter | blocklist | hostname | 94 | 85 | 90.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 324 | 84.6% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 440 | 74.7% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 28.2K | 71.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 66.7% |
| phishing_army | blocklist | domain | 150.6K | 79.9K | 53.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 80.3K | 43.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.6% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 267 | 42.7% |
| WaLLy3K | blocklist | domain | 351 | 138 | 39.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| HaGeZi Pro | blocklist | domain | 229.9K | 86.1K | 37.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6.4K | 34.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 81 | 34.6% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 79.4K | 32.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 55.9K | 31.1% |
| kadantiscam | blocklist | domain | 42.9K | 12.6K | 29.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 22.8K | 25.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 87 | 25.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 11.3K | 24.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 23 | 21.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42.9K | 21.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.9K | 21.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 5.8K | 20.9% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 54 | 13.9% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 4.1K | 13.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 554 | 11.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 37 | 10.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 2.2K | 2.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 193 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 8 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 71 | 1.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 109 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 382 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 448 | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 278 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 121 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 2.4K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 6.4K | 0.5% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1.2K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 4.2K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 11 | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 9 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1.1K | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 35 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 219 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 246.3K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.4K | 55.7K | 98.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 820 | 61.7% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 80.3K | 51.5% |
| EasyList | blocklist | adguard | 67.8K | 33.9K | 50.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 80.3K | 43.7% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 1.3K | 36.8% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 440 | 35.6% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 55.9K | 30.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 65 | 17.7% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 79.8K | 13.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 4.1K | 13.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 5.7K | 10.2% |
| abpvn_hosts | blocklist | adguard | 984 | 38 | 3.9% |
| CJX Annoyance | blocklist | adguard | 1.8K | 56 | 3.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 68 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 12 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 121 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 9 | 0.3% |
| AdBlockID | blocklist | adguard | 3.7K | 10 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 35 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.1% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 22.1K | Targets: 42 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 6.5K | 10.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 7.1K | 9.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 15.0K | 6.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 33 | 2.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 23 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 64 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 48 | 0.2% |
| Torrent Trackers | blocklist | domain | 488 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 122 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 80 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 19 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 301 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 46 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 254 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 52 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 28 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 27 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 24 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 121 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 12 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 38 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 6 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 22.1K | Targets: 12 | Unique: 21.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 33 | 2.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 64 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 80 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 122 | 0.1% |
| EasyList | blocklist | adguard | 67.8K | 71 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 121 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 4 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 15 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 38 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 56.4K | Targets: 64 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 94 | 87 | 92.6% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 439 | 74.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.8% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 36.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 51.3K | 28.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4.7K | 25.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 54.0K | 23.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 55.7K | 22.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 21 | 19.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 36.1K | 17.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6.1K | 13.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3.3K | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 752 | 11.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 24 | 6.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 927 | 6.1% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.4K | 4.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 578 | 4.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.6K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 66 | 4.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 6 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 79 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 80 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 929 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 59 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 213 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 259 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 170 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 24 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 10 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 197 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 426 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 78 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 186 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 70 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 114 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 45 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 167 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.4K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 836 | 62.9% |
| EasyList | blocklist | adguard | 67.8K | 33.7K | 49.7% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 439 | 35.5% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 51.3K | 28.4% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 55.7K | 22.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| Easy Privacy | blocklist | adguard | 55.3K | 2.6K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1.4K | 4.6% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 66 | 4.0% |
| abpvn_hosts | blocklist | adguard | 984 | 33 | 3.4% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 80 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 197 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 69 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 234 | Targets: 21 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 81 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 29 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 3 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 81 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 17 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 43 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 56 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 71 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 67 | 66.3% |
| Torrent Trackers | blocklist | domain | 488 | 70 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 31.1K | Targets: 78 | Unique: 0 | Conflicts: 162</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 76 | 19.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 687 | 19.4% |
| tranco | allowlist | domain_top | 500 | 80 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 570 | 13.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 29 | 12.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 46 | 9.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.5K | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.3% |
| Adaway | blocklist | hostname | 6.5K | 520 | 8.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.9K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 34 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 141 | 5.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 15 | 5.1% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 26 | 3.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1.4K | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 9 | 2.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.3K | 2.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 279 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 4.5K | 1.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.6K | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 4.1K | 1.7% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 625 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 6 | 1.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2.3K | 1.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 59 | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1.9K | 1.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 462 | 1.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 48 | 0.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 3.9K | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| phishing_army | blocklist | domain | 150.6K | 717 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 359 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 2.3K | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 185 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 64 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 47 | 0.3% |
| Torrent Trackers | blocklist | domain | 488 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 201 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 58 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 92 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 19 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 7 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 690 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 75 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.8K | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 59 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 105 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 76 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 31.1K | Targets: 24 | Unique: 17.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 465 | 12.9% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 70 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 34 | 2.7% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 1.4K | 2.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 4.1K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 2.3K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 1.9K | 1.1% |
| EasyList | blocklist | adguard | 67.8K | 646 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.3K | 337 | 0.6% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 753 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 47 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 1.5K | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 64 | 0.3% |
| abpvn_hosts | blocklist | adguard | 984 | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 150.6K | Targets: 37 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 20.4K | 51.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 80.5K | 43.8% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 107.5K | 43.7% |
| kadantiscam | blocklist | domain | 42.9K | 17.9K | 41.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 79.9K | 32.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 56 | 23.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 17.0K | 19.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 26 | 5.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 11.4K | 4.9% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 717 | 2.3% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1.7K | 1.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 36 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 30 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 277 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 598 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 33 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 521 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 12 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 100 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 404 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 21 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 23 | Unique: 61.6K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 95 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 93 | 4.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 2 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 2 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 4 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 7 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 3 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 5 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 48 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 142 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 76 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 296 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 47 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 46 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 145 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 299 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 298 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 59 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 22 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 80 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 20 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 28 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 18 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 47 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 11 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 57 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 575 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 473 | 11.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 34 | 8.8% |
| tranco | allowlist | domain_top | 500 | 39 | 7.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 434 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.3K | 4.1% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 94 | 3 | 3.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 904 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 315 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 11 | 1.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 6 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 927 | 1.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 11 | 1.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 1 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.2K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 996 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2.1K | 1.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 2.5K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 1 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.4K | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 183 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Torrent Trackers | blocklist | domain | 488 | 2 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 587 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 23 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 12 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 6 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 8 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 67.8K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 597.9K | Targets: 15 | Unique: 327.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 3.6K | 99.4% |
| RPiList_specials-phishing | blocklist | adguard | 156.1K | 107.9K | 69.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 77.4K | 42.1% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 79.8K | 32.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1.5K | 4.8% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 69 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 181 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 12 | 0.1% |
| EasyList | blocklist | adguard | 67.8K | 47 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 15 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.3K | 3 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 156.1K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 81.3K | 44.2% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 80.3K | 32.6% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 107.9K | 18.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 753 | 2.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 3.6K | 9 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 8 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.8K | Targets: 25 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 369 | 96 | 26.0% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 101 | 14.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 98 | 14.5% |
| Greensnow | blocklist | ipv4 | 5.0K | 591 | 11.7% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 1.3K | 5.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 405 | 2.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 1.3K | 1.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 195 | 1.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 1.0K | 1.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 56 | 1.2% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 112 | 1.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 22 | 0.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 88 | 0.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 18 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 39 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 39 | 0.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 39 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 18 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 13 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 19 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 21 | Unique: 390 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 30 | 18.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 115 | 9.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 19 | 1.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 8 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 212 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 7 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 174 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 29 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 9 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 7 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 7 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 22 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 13 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 9 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 18 | Unique: 1.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 369 | 87 | 23.6% |
| Greensnow | blocklist | ipv4 | 5.0K | 106 | 2.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 22 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 188 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 139 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 31 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 144 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 8 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 2 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.4K | Targets: 21 | Unique: 0 | Conflicts: 20.3K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 11.4K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 11.4K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 2.1K | 8.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 1.6K | 8.4% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 4.7K | 6.5% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 39 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 32 | 1.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 219 | 1.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 38 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 13 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 35 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 21 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 5 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 42 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 2 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.3K | Targets: 27 | Unique: 0 | Conflicts: 21</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 949 | 19.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 782 | 15.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 195 | 10.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 2.2K | 10.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.6K | 10.6% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 37 | 10.0% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 1.1K | 9.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 5.6K | 9.3% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 2.2K | 9.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 451 | 8.9% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 6.0K | 8.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 1.6K | 7.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 1.2K | 6.3% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 38 | 5.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 34 | 5.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 58 | 4.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 18 | 1.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 203 | 1.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 7 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 69 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 21 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 21 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 59 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 33 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 12 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.4K | Targets: 18 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 33 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 10 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 29 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 12 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 713 | Targets: 40 | Unique: 327 | Conflicts: 313</summary>

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
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 2 | 0.9% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 40 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 2 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 26 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 12 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 31 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 27.8K | Targets: 55 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 497 | 30.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 975 | 27.5% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 106 | 18.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 672 | 15.7% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 94 | 11 | 11.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.7K | 9.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 7 | 6.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 24 | 6.2% |
| Adaway | blocklist | hostname | 6.5K | 408 | 6.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.9K | 6.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3.3K | 5.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 904 | 5.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| tranco | allowlist | domain_top | 500 | 23 | 4.6% |
| HaGeZi Pro | blocklist | domain | 229.9K | 8.7K | 3.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 6.0K | 3.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 11 | 3.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.2K | 3.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 355 | 2.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 5.8K | 2.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1.9K | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 48 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 161 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 123 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 14 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 204 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 171 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 286 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 35 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 52 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 29 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 12 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 16 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 222.1K | Targets: 35 | Unique: 163.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 15.0K | 67.8% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.5K | 28.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 448 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 76 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 470 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 22 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 106 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 50 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 219 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 107 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 70 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 69 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 188 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 119 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 27 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 61 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 139 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 208 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 28 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.0K | Targets: 45 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 6.1K | 10.8% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 40 | 6.8% |
| HaGeZi Pro | blocklist | domain | 229.9K | 12.3K | 5.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 926 | 5.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 11.3K | 4.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 6.0K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.0K | 3.0% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 55 | 1.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 183 | 1.2% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 94 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 185 | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 866 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 313 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 256 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 1 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3.7K | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| kadantiscam | blocklist | domain | 42.9K | 69 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 195 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 368 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1.2K | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 28 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 13 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 39 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 12 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.3K | Targets: 32 | Unique: 4.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 960 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 240 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 23 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 144 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 246 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 71 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 10 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 16 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 13 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 27 | Unique: 1.3K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 3 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 48 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 19 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 49 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 58 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 23 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 2 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 6.5K | 29.5% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 19.4K | 8.7% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 2 | 2.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 65 | 1.8% |
| Torrent Trackers | blocklist | domain | 488 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 1 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 615 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 213 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 123 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 75 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 274 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 23 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 539 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 382 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 44 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 39 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 29 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 13 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 35 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 11 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 2.7K | 3.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 141 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1.2K | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 36 | Unique: 0 | Conflicts: 93</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 42 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 17 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 29 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 15 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 65 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 32 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 38 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 32 | Unique: 5.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 1.6K | 0.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 20 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 191 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 12 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 5 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 14 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 56 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 11 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 110 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 6 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.5% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 17.6K | 7.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 4.3K | 1.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 2 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 5.0K | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 980 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 229.9K | 428 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 389 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 85 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 24 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 19 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 4 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 89.3K | Targets: 69 | Unique: 0 | Conflicts: 77</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.5K | 99.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 355 | 92.7% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 42.9K | 35.5K | 82.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 50.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.6K | 41.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 402 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 66 | 17.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 49 | 13.2% |
| hufilter | blocklist | hostname | 94 | 12 | 12.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 899 | 12.3% |
| phishing_army | blocklist | domain | 150.6K | 17.0K | 11.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 22.8K | 9.2% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 12.9K | 7.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.9K | 6.9% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3.6K | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 15.4K | 6.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 14.4K | 6.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 32 | 5.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.6K | 5.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 31 | 4.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 4.9K | 2.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 342 | 2.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 313 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 3.9K | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 41 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 813 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 125 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 276 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 46 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.3K | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 450 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 24 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 1.0K | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 106 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 194 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 20 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 46 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7.1K | 32.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 21.5K | 9.7% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 2 | 2.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 73 | 2.1% |
| hufilter | blocklist | hostname | 94 | 2 | 2.1% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 488 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 1 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 157 | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 703 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 161 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 259 | 0.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 229.9K | 630 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 92 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 343 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 448 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 13 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 50 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 34 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 43 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 34 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 48 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 18 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 36 | Unique: 0 | Conflicts: 93</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 42 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 29 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 17 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 65 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 32 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 38 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 35 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 43.1K | Targets: 30 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 14.1K | 7.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 28.9K | 5.5% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 4.7K | 5.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 12 | 3.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 20.6K | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 462 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 1 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 2.9K | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 634 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 2 | 0.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 278 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 41 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 161 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| phishing_army | blocklist | domain | 150.6K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 37 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 114 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 6 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 529.5K | Targets: 50 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 28.9K | 67.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 383 | 234 | 61.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 29.3K | 33.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 270.9K | 22.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 221.9K | 22.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 3.9K | 12.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 16.6K | 9.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 6.2K | 2.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 6 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 2.4K | 1.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 1.8K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 368 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 450 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 23 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 1 | 0.4% |
| phishing_army | blocklist | domain | 150.6K | 598 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 736 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 130 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 167 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 276 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 13 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 6 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 664 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 35 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 27 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 139 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 44 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20.3K | Targets: 29 | Unique: 0 | Conflicts: 219</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 575 | 43.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.7K | 31.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 7.4K | 30.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 1.3K | 25.3% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 18.1K | 25.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 2.7K | 23.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 405 | 22.4% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 152 | 21.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 2.2K | 21.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 138 | 20.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 762 | 15.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 7.4K | 12.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 2.7K | 11.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 455 | 8.9% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 28 | 7.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 1.2K | 6.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 61 | 5.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 22 | 2.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 219 | 1.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 159 | 3 | 1.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 219 | 1.9% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 230 | 1.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 18 | 1.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 31 | 1.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 88 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 7 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 145 | 0.1% |
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
<summary>List Type: blocklist | Source Type: domain | Total: 488 | Targets: 9 | Unique: 294 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 100 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 71 | 70 | 98.6% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 47 | Unique: 0 | Conflicts: 580</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 122 | 7.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 2 | 4.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 94 | 2 | 2.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 32 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 29 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 234 | 1 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 39 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 80 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 34 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 11 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 23 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 33 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 33 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 27 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 16 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.4K | 30 | 0.1% |
| EasyList | blocklist | adguard | 67.8K | 51 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 597.9K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 183.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 38 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 11 | Unique: 25 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.3K | 164 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 43 | 0.1% |
| Easy Privacy | allowlist | adguard | 842 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 53 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 65 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 2 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 383 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 355 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 76 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 320 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 266 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 324 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 25 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 234 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 41 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 35 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 57 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 13.5K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 13.2K | 18.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 4 | 8.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 203 | 2.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 230 | 1.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 7 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 6 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 3 | 0.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 117 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 134 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 84 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 232 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 13 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 15 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 9 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 253 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 16 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 4 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 11 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.2K | Targets: 1 | Unique: 56.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.6K | Targets: 34 | Unique: 8.1K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 280 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 51 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 369 | 12 | 3.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5.3K | 2.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 675 | 10 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.2K | 17 | 1.4% |
| Greensnow | blocklist | ipv4 | 5.0K | 71 | 1.4% |
| BruteforceBlocker | blocklist | ipv4_find | 695 | 8 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 15 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 4.8K | 36 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 208 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 9 | 0.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 41 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 411 | 0.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.3K | 69 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 79 | 0.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 13 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 405 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 84 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 116 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 57 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 88 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 66 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 4 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 2.1K | 0.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 3.9K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 554 | 0.2% |
| HaGeZi Pro | blocklist | domain | 229.9K | 137 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 204 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 442 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 68 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 15 | Unique: 361 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 280 | 1.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.4K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 24.6K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 18 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.3K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.7K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.2K | 5 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 56.2K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 6 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.4K | 45 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 72.3K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 23 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.5K | 4 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 351 | Targets: 34 | Unique: 0 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| YousList | blocklist | hostname | 625 | 9 | 1.4% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 19 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 83 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 3 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 54 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 138 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 19 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 85 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 2 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 35 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 20 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 84 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 1 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 66 | Unique: 0 | Conflicts: 94</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.6K | 74.0% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3.0K | 69.9% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6.3K | 34.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 20.7K | 23.2% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 62 | 16.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 94 | 14 | 14.9% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 188 | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 41 | 10.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 15.9K | 6.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 996 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.8K | 6.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 3.1K | 5.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 1.4K | 4.4% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 25 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 27 | 3.8% |
| HaGeZi Pro | blocklist | domain | 229.9K | 6.8K | 3.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 3.2K | 1.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 129 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 161 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2.5K | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 335 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 26 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 977 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 31 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 50 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 33 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 21 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 43 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 17 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |

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
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 199 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 3 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 24 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 86 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 2 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 94 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 42 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 151 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 267 | 0.1% |
| HaGeZi Pro | blocklist | domain | 229.9K | 202 | 0.1% |
| kadantiscam | blocklist | domain | 42.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 3 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 7 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| AdGuard DNS filter | blocklist | adguard | 180.6K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.3K | 10 | 0.0% |
| EasyList | blocklist | adguard | 67.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 246.3K | 68 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.4K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 31.1K | 2 | 0.0% |

</details>

---

### youtube_GoodbyeAds

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 23 | Unique: 97.2K | Conflicts: 10</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| WaLLy3K | blocklist | domain | 351 | 7 | 2.0% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 50 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 2 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 74 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 229.9K | 42 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 45 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.5K | Targets: 62 | Unique: 0 | Conflicts: 45</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.8K | 42.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 392 | 23.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2.0K | 11.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Adaway | blocklist | hostname | 6.5K | 262 | 4.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 89.3K | 3.5K | 3.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 575 | 3.8% |
| YousList | blocklist | hostname | 625 | 24 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 4 | 3.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 975 | 3.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 422 | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.4K | 1.6K | 2.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 388 | 10 | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 589 | 13 | 2.2% |
| ph00lt0_blocklist | blocklist | domain | 31.1K | 687 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 713 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 5 | 1.4% |
| HaGeZi Pro | blocklist | domain | 229.9K | 3.0K | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 371 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 246.3K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 179.9K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 16 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 55 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 990.5K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 246.3K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 183.8K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 87.8K | 7 | 0.0% |
| phishing_army | blocklist | domain | 150.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 16 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 43.1K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 222.1K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 529.5K | 6 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 522.4K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| kadantiscam | blocklist | domain | 42.9K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 30 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.6K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

