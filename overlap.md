# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-10-11 00:43:17 UTC

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
| Total Entries Analyzed | 9.3M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 143 |

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
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 71 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 68.2% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 11.2K | 60.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 33.8K | 60.3% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 310 | 47.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 167 | 45.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 142 | 37.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 588 | 35.6% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 32.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 110 | 31.9% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 33 | 31.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 59.0K | 30.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 44.3K | 24.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 6.3K | 22.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 40.6K | 17.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 13.2K | 16.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 6.1K | 13.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 4.4K | 12.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 21.2K | 8.8% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 147 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 3.3K | 1.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 257 | 1.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 26 | 0.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 2 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 451 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 144 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 43 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 433 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 458 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| kadantiscam | blocklist | domain | 40.0K | 87 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.1K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 47 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 738 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| phishing_army | blocklist | domain | 140.3K | 89 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 215 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 630 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 212 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 906 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 29 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.4K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 23 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 35 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 52 | Unique: 0 | Conflicts: 37</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.7K | 14.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 6.5K | 8.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 263 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 6 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 21 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 434 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 9 | 2.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 408 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 10 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 521 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 750 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 2.4K | 1.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.7K | 0.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 90 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1.0K | 0.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 26 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 33 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 50 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 14 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 93 | Targets: 8 | Unique: 32 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 211 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 3 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 35 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 5 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 657 | Targets: 32 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 483 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 13 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 115 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 625 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 54 | 0.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 545 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 487 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 310 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 44 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 34 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 114 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 33 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.0K | 483 | 0.9% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 644 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 487 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 3 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.4K | 44 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| EasyList | blocklist | adguard | 66.5K | 89 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 34 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 114 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 1 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 218.3K | Targets: 14 | Unique: 217.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 433 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 227.5K | Targets: 24 | Unique: 118.0K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 91.6K | 50.6% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.2K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 585 | 3.8% |
| HaGeZi Pro | blocklist | domain | 195.7K | 6.2K | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 909 | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 811 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 30 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1.8K | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 204 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 84 | 0.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 2 | 0.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.2K | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 336 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 59 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 43 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 181.6K | Targets: 26 | Unique: 0 | Conflicts: 200</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 203 | 199 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.5% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 51.0K | 91.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 66.5K | 47.0K | 70.6% |
| Easy Privacy | blocklist | adguard | 55.4K | 28.7K | 51.7% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 644 | 49.2% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 55.7K | 23.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 25.8K | 10.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 2.4K | 6.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 23 | 2.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.7K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 117 | 0.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 179 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 7 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 180.9K | Targets: 69 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 657 | 625 | 95.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.5K | 93.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 51.0K | 91.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 91.6K | 40.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 76.0K | 38.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5.0K | 27.1% |
| YousList | blocklist | hostname | 625 | 151 | 24.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 55.7K | 23.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 966 | 22.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 44.3K | 21.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 6.0K | 21.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 21 | 19.8% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5.8K | 12.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.7K | 11.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 25.8K | 10.8% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 25 | 7.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2.4K | 6.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 4.9K | 6.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 20 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 554 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 117 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 346 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 277 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 187 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 273 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 194 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 29 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 636 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 274 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 364 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 107 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 118 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 7 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 5 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.0K | 841 | 1.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 825 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 71 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.4K | 77 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 7 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 203 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 211 | 199 | 94.3% |

</details>

---

### AdGuardTeam_HttpsExclusions_android

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 97 | Targets: 11 | Unique: 70 | Conflicts: 17</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 10 | Unique: 4.0K | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1 | 0.0% |
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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
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
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 2.1K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 25 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 180 | 0.1% |
| EasyList | blocklist | adguard | 66.5K | 2 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 57 | Unique: 0 | Conflicts: 48</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.0K | 56.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.8% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.4% |
| YousList | blocklist | hostname | 625 | 198 | 31.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 205 | 12.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 12 | 11.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 7.6K | 9.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 4.7K | 8.3% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 54 | 8.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.7% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1.7K | 6.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.2K | 5.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 18 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.5K | 4.5% |
| HaGeZi Pro | blocklist | domain | 195.7K | 7.1K | 3.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 5.0K | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 6.4K | 2.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 934 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1.2K | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 55 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 43 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 12 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 142 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 29 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 30 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 14 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.1K | Targets: 23 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 445 | 8.7% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 40 | 7.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 645 | 7.0% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 26 | 6.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 37 | 6.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1.9K | 5.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 812 | 5.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 1.1K | 4.6% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 463 | 4.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 61 | 3.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 707 | 3.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 453 | 2.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 1.5K | 2.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 64 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 148 | 1.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 26 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 14 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 9 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 9 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 25 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 15.4K | 945 | 6.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 183 | 4.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 544 | 3.1% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 18 | 1.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 79 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 3 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 1 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 7 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 114 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 165 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 14 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 205 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 14 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 21 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 14 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 46 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 382 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 107 | 6.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 87 | 3.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 126 | 2.9% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 205 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 6 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 26 | 0.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 5 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 205 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 268 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 31 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 17 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 26 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 22 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 12 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 12 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 3 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 3 | 0.0% |

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
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 77 | 0.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 187 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 45 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 189 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 22 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 41 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 126 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 264 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 49 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 33 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 107 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 53 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 3 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 241.0K | Targets: 50 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 140.3K | 101.8K | 72.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 278 | 69.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 75.5K | 31.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 74.6K | 31.1% |
| kadantiscam | blocklist | domain | 40.0K | 10.4K | 26.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 12.7K | 15.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1.9K | 5.3% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 3.1K | 3.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3.6K | 1.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18.7K | 1.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 16.8K | 1.6% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 568 | 1.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 6.2K | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 284 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 195 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 16 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 215 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 44 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 107 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 491 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 57 | Unique: 0 | Conflicts: 71</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.8K | 51.4% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 48 | 13.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 169 | 10.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.9K | 10.1% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 40 | 5.6% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 514 | 3.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 13 | 3.7% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 473 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 11 | 2.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1.6K | 2.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 2.2K | 2.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 672 | 2.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 573 | 1.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 6 | 1.6% |
| HaGeZi Pro | blocklist | domain | 195.7K | 2.1K | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 5 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.9K | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 966 | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 26 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 42 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 24 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 19 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 112 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 5 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 33 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 2 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 64 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 21 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 55 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 34 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 6 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 11 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 15 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 60.9K | Targets: 32 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 382 | 268 | 70.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 673 | 58.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 881 | 56.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.7K | 53.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 313 | 52.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.5K | 50.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1.5K | 47.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4.2K | 46.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 2.0K | 45.5% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 4.9K | 44.8% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 237 | 43.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 5.4K | 34.7% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 11.5K | 32.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 7.4K | 31.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 5.3K | 30.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 4.4K | 24.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 263 | 21.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 205 | 17.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 10 | 8.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 80 | 7.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 20 | 7.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 412 | 2.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 249 | 1.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 38 | 1.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 420 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 8 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 23 | Unique: 206 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 10 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 545 | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 513 | 85.5% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 509 | 4.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 33 | 2.1% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 6 | 1.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 40 | 1.3% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 131 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 263 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 86 | 0.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 19 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 37 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 237 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 11 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 17 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 26 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 5 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 6 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 101 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 68 | 100.0% |
| Torrent Trackers | blocklist | domain | 495 | 100 | 20.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 21 | Unique: 0 | Conflicts: 36</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.0K | 7.8K | 70.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 812 | 26.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 4.5K | 25.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 5.1K | 21.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.3K | 13.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 5.0K | 13.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 7.5K | 12.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 990 | 5.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 190 | 3.7% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 261 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 25 | 1.6% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 8 | 1.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 65 | 1.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 6 | 1.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 6 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 67 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 36 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 36 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 12 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 32 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 8 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.4K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 55 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 9 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 94.7K | Targets: 49 | Unique: 8.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 3.2K | 8.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 26 | 6.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 30.3K | 5.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 5 | 2.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 22.4K | 1.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 15.9K | 1.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 70 | 1.4% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 470 | 1.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 3.0K | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 3.1K | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 424 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 163 | 1.2% |
| phishing_army | blocklist | domain | 140.3K | 1.5K | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 2.2K | 0.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 111 | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 264 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 291 | 0.4% |
| HaGeZi Pro | blocklist | domain | 195.7K | 776 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 33 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 170 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 7 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 29 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 451 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 15 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 44 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 26 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 187 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 597 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 15 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 36 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 35 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 72 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 275 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 13.0K | 15.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 46 | 12.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 421 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 359 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 8 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 2.5K | 1.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 581 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 280 | 0.8% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.4K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 992 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 617 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 554 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 163 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 95 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 635 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 690 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 29 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 13 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 36 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 15 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 7 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 84 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 18 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |

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
| EasyList | blocklist | adguard | 66.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 176 | 14.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 4 | 3.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 18 | 1.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 673 | 1.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 20 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 146 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 59 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 51 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 46 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 22 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 3 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1 | 0.0% |
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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 9 | Unique: 307 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 96.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 81 | 11.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 92 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1.0K | 31.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 882 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |

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
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 81 | 4.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 39 | Unique: 11.4K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3.0K | 90.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 882 | 78.0% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 47 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 131 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 13 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 6 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 526 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 36 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 16 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 16 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 63 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 5 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 10 | Unique: 13.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3.0K | 90.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 47 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 13 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 36 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 5.1K | 28.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 5.1K | 21.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 445 | 14.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 2.4K | 6.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 567 | 6.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 2.7K | 4.5% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 17 | 4.5% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 11 | 2.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 315 | 1.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 10 | 1.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 69 | 1.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 190 | 1.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 14 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 14 | 0.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 6 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 50 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 18 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 54 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 2 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.4K | Targets: 20 | Unique: 13.3K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.3% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 28.7K | 15.8% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.8% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 2.7K | 4.7% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 44 | 3.4% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 5.7K | 2.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 342 | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 825 | 0.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 3 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 13 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 66.5K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.0K | 32.8K | 58.6% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 47.0K | 25.9% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 33.2K | 13.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 19.2K | 8.0% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 89 | 6.8% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 653 | 1.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 67 | 0.3% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 45 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 600 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 545 | 513 | 94.1% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 512 | 4.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 27 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 5 | 1.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 37 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 261 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 107 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 79 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 313 | 0.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 17 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 27 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 10 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 26 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 16 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 5 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.3% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.4% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 5 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 56 | 7.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 42 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 11 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

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
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4 | 0.0% |

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
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 38 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 25 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 48 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 16 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 51 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 15 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 3 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 123 | Targets: 12 | Unique: 60 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 23 | 1.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 8 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 10 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 8 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 8 | Unique: 471 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 1 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 11 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 3 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 3 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 20 | Unique: 687 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 68 | 2.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 30 | 2.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 2 | 1.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 15 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 15 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 17 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 17 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 80 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 14 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 27 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 51 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 9 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 6 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.6K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.5K | 91.9% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.7% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 18.2K | Targets: 32 | Unique: 0 | Conflicts: 656</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 945 | 82.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 3.4K | 78.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1.0K | 66.1% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 205 | 53.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 9.5K | 40.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 6.4K | 36.5% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 131 | 24.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 707 | 22.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 107 | 17.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.6K | 17.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 5.4K | 15.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 5.4K | 8.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 990 | 6.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 8 | 6.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 75 | 6.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 656 | 5.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 656 | 5.7% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 46 | 4.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 99 | 3.9% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 374 | 3.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 17 | 1.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 79 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 25 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 85 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 11.0K | Targets: 29 | Unique: 0 | Conflicts: 63</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 509 | 93.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 512 | 85.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 9.5K | 52.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.8K | 51.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1.1K | 34.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 6.4K | 17.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.5K | 16.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 2.8K | 15.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 7.4K | 12.2% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 26 | 6.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Greensnow | blocklist | ipv4 | 4.3K | 151 | 3.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 55 | 3.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 27 | 2.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 374 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 21 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 128 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 103 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 60 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 60 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 436 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 15 | Unique: 2.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 56 | 20.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 68 | 6.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 1 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 1 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 38 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 6 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 6 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 279 | Targets: 9 | Unique: 196 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 56 | 2.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 20 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 1 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.6K | Targets: 20 | Unique: 5.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 314 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 90 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1.7K | 1.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 2.2K | 1.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.5K | 0.8% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 16 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 340 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 55 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 130 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 22 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 57 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 36 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 108 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 13 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.7K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 392 | 11.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 169 | 4.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 500 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 205 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1.5K | 0.9% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.6K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 400 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 588 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 191 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 167 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 27 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.5% |
| Easy Privacy | blocklist | adguard | 55.4K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 27 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 167 | 0.1% |
| EasyList | blocklist | adguard | 66.5K | 1 | 0.0% |

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
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.3K | Targets: 27 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 382 | 126 | 33.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 406 | 26.2% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 3.4K | 21.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 183 | 15.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1.9K | 5.3% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 19 | 3.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 2.0K | 3.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 77 | 3.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 269 | 2.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 17 | 2.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 478 | 2.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 64 | 2.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 20 | 1.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 304 | 1.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 15 | 1.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 69 | 1.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 14 | 1.1% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 74 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 151 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 60 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 65 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 7 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 7 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 15 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 368 | Targets: 19 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 91 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 334 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 167 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 21 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 49 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 20 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 37 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 106 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 66 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 23 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 21 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 21 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 239.9K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 237.9K | 140.5K | 59.0% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 73.3K | 51.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 21.7K | 38.7% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 935 | 33.7% |
| EasyList | blocklist | adguard | 66.5K | 19.2K | 28.9% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 25.8K | 14.2% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 75.6K | 12.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 167 | 10.0% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 114 | 8.7% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 180 | 6.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 2.1K | 6.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 825 | 1.5% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 9 | 0.7% |
| abpvn_hosts | blocklist | adguard | 1.0K | 5 | 0.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 7 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 9 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 2 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 239.9K | Targets: 53 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 271 | 68.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 140.5K | 59.0% |
| phishing_army | blocklist | domain | 140.3K | 72.2K | 51.4% |
| HaGeZi Pro | blocklist | domain | 195.7K | 80.5K | 41.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 21.7K | 38.7% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 13.5K | 35.7% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 74.6K | 31.0% |
| kadantiscam | blocklist | domain | 40.0K | 10.8K | 27.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 12.2K | 26.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 25 | 20.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 6.8K | 19.5% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 114 | 17.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 12.0K | 14.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 25.8K | 14.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.5K | 13.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 31 | 12.8% |
| Spam404 | blocklist | domain | 8.1K | 1.0K | 12.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21.2K | 10.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 167 | 10.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.6K | 7.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 362 | 7.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.5% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2.1K | 6.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 617 | 4.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 307 | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1.1K | 3.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 18.7K | 3.4% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 3.0K | 3.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 112 | 2.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 64 | 1.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 107 | 0.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 9.4K | 0.8% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 1.8K | 0.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 9 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 109 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 48 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.0K | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 33 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 5.4K | 0.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1.8K | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 129 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 158 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 36 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 320 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 6 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.3K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.3% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 9 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 41 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.3K | Targets: 5 | Unique: 286 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.3% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 9 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 580.9K | Targets: 43 | Unique: 563.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 44.9% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2.6K | 7.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 3.9K | 4.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1.8K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.5K | 0.6% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 3 | 0.5% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.0K | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 275 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.5K | 0.3% |
| kadantiscam | blocklist | domain | 40.0K | 114 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 88 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 491 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 856 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 212 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| phishing_army | blocklist | domain | 140.3K | 150 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 18 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 118 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 27 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 42 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 78 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 484 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 8 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 11 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 327 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 31 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 383 | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 46 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 10 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 331 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 40 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 24 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 66 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 20 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 55 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 25 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 195.7K | Targets: 70 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 342 | 99.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 98.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 54.0K | 96.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 334 | 90.8% |
| hufilter | blocklist | hostname | 92 | 81 | 88.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 331 | 86.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.0K | 83.6% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 545 | 83.0% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 66 | 62.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 75 | 60.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.1K | 50.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 76.0K | 42.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.1K | 38.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 111 | 37.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 80.5K | 33.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 79.0K | 33.2% |
| YousList | blocklist | hostname | 625 | 201 | 32.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 59.0K | 29.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 8.2K | 29.1% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 25.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 11.5K | 24.9% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.4K | 15.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 10.6K | 13.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 4.3K | 12.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.4K | 10.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.5K | 10.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 8.9% |
| kadantiscam | blocklist | domain | 40.0K | 3.2K | 8.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 28 | 7.0% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 245 | 3.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 7 | 2.9% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 6.2K | 2.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 127 | 2.6% |
| phishing_army | blocklist | domain | 140.3K | 3.4K | 2.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 426 | 2.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 586 | 1.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 187 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| Spam404 | blocklist | domain | 8.1K | 120 | 1.5% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 3.6K | 1.5% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 41 | 1.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 237 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 24 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 521 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 611 | 0.8% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 776 | 0.8% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 131 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 51 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.5K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.2K | 0.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 3 | 0.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1.6K | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1.9K | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1.0K | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 80 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 480 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 41 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 345 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 342 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 88 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 16 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 22 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 9 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 70.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 79 | 10.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 95 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 35.8K | Targets: 35 | Unique: 0 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 14.8K | 95.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1.0K | 65.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1.9K | 63.0% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 205 | 53.7% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 263 | 48.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4.4K | 47.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.4K | 46.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 1.9K | 44.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 261 | 43.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 5.4K | 35.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.0K | 33.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 6.4K | 27.2% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 2.7K | 24.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 4.2K | 23.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 4.1K | 23.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 11.5K | 18.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 177 | 14.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 165 | 14.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 146 | 12.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 8 | 6.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 51 | 4.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 98 | 3.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 305 | 1.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 33 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 15 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 481 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 4 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 23 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 53 | 0.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 111 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 20 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 108 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 13 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 21 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |

</details>

---

### hufilter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 92 | Targets: 25 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 85 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 70 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 83 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 6 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 81 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 11 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 58 | Unique: 421.6K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 995 | 13.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1.6K | 4.7% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 9 | 2.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 4 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi Pro | blocklist | domain | 195.7K | 2.5K | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 3.1K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 3.0K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 26 | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 597 | 0.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 98 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 2.4K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 4.4K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.2K | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 182 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 112 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1.5K | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 273 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 52 | 0.2% |
| phishing_army | blocklist | domain | 140.3K | 264 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 194 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 49 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 21 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 21 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 210 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 2 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 40.0K | Targets: 42 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 30.1K | 36.8% |
| phishing_army | blocklist | domain | 140.3K | 15.9K | 11.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 11.2K | 4.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 10.8K | 4.5% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 10.4K | 4.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3.2K | 1.6% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 526 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 2 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 9 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 29 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 70 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 68 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 19 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 17 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 26 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 87 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 49 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 23 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 114 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 24 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 24 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 13 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 31 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 175 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 126 | 2.1% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 47 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 5 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 17 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 24 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 4 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 3 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |

</details>

---

### local_domain_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |

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
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1 | 0.0% |
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
| Firehol_level3 | blocklist | ipv4 | 11.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 513 | 1.5% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 1.3K | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 935 | 0.4% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 2.5K | 0.4% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.4K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 3 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2M | Targets: 61 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 973.5K | 93.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 271.0K | 49.7% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 15.4K | 40.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 22.4K | 23.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 43 | 10.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 3.7K | 7.9% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 18.7K | 7.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 11 | 6.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.8K | 5.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 690 | 5.3% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 9.4K | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 526 | 3.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 6.3K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 4.2K | 2.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1.3K | 1.6% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 8 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 286 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 416 | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.1K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 25 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 25 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 636 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 15 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| phishing_army | blocklist | domain | 140.3K | 437 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 80 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 121 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 484 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 13 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 34 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 29 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 19 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 215.2K | Targets: 37 | Unique: 203.1K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 85.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 5.3K | 33.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 46 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 12 | 3.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 22 | 1.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 436 | 1.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 235 | 1.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 481 | 1.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 54 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 25 | 0.8% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 10 | 0.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 4 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 8 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 4 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 420 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 130 | 0.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 85 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 6 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 53 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 50 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 8 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 10 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 69 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 15 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 32 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 48 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.0M | Targets: 49 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 973.5K | 81.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 220.7K | 40.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 15.9K | 16.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 38 | 9.5% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 16.8K | 7.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 2.6K | 6.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 846 | 2.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 5.4K | 2.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 4.2K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1.0K | 1.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.9K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.4K | 0.9% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| phishing_army | blocklist | domain | 140.3K | 355 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 99 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 630 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 364 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 111 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 5 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 327 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 8 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 18 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 65 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 17 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 15 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 9 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 34.8K | Targets: 34 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 172 | 70.8% |
| phishing_army | blocklist | domain | 140.3K | 17.5K | 12.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 24.5K | 10.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 17 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 6.8K | 2.8% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 558 | 1.6% |
| kadantiscam | blocklist | domain | 40.0K | 526 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 1.9K | 0.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 21 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 470 | 0.5% |
| HaGeZi Pro | blocklist | domain | 195.7K | 586 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 18 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.6K | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 7 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 36 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 84 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 47 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 11 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 115 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 99 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 80 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 237.9K | Targets: 65 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 55.3K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 98.7% |
| hufilter | blocklist | hostname | 92 | 83 | 90.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 333 | 83.7% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 487 | 74.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 24.5K | 70.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.4K | 66.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 140.5K | 58.6% |
| phishing_army | blocklist | domain | 140.3K | 72.7K | 51.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.7% |
| YousList | blocklist | hostname | 625 | 270 | 43.2% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 195.7K | 79.0K | 40.4% |
| WaLLy3K | blocklist | domain | 351 | 137 | 39.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.4K | 34.7% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 75.5K | 31.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 55.7K | 30.8% |
| kadantiscam | blocklist | domain | 40.0K | 11.2K | 28.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 88 | 25.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 20.7K | 25.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 10.7K | 23.0% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 23 | 21.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.9K | 21.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 5.8K | 20.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 40.6K | 20.0% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 55 | 14.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 4.0K | 11.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 559 | 11.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 37 | 10.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 16 | 6.6% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 2.2K | 2.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 189 | 1.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 88 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.1K | 0.7% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 262 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 108 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 390 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 126 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 455 | 0.6% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 1.2K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 6.3K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 4.2K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 2.3K | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 9 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1.5K | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 36 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 308 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 7 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 237.9K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.0K | 55.3K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.3% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 825 | 61.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 140.5K | 58.6% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 73.6K | 51.3% |
| EasyList | blocklist | adguard | 66.5K | 33.2K | 49.9% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 1.3K | 45.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 487 | 37.2% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 55.7K | 30.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 66 | 17.9% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 75.0K | 12.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 4.0K | 11.7% |
| Easy Privacy | blocklist | adguard | 55.4K | 5.7K | 10.3% |
| abpvn_hosts | blocklist | adguard | 1.0K | 35 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 55 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 126 | 0.6% |
| AdBlockID | blocklist | adguard | 3.7K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 36 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 19.4K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.9K | 9.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.5K | 8.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 13.2K | 5.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 30 | 2.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 24 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 44 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 58 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 237 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 17 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 126 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 257 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 79 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 117 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 22 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 20 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 40 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 44 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 19.5K | Targets: 10 | Unique: 19.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 30 | 2.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 58 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 117 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 79 | 0.1% |
| EasyList | blocklist | adguard | 66.5K | 67 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 126 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 13 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 56.0K | Targets: 64 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 85 | 92.4% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 483 | 73.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1.6K | 45.8% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 37.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 51.0K | 28.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 54.0K | 27.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4.7K | 25.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 55.3K | 23.3% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 21 | 19.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33.8K | 16.6% |
| YousList | blocklist | hostname | 625 | 95 | 15.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5.7K | 12.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 3.3K | 11.6% |
| Adaway | blocklist | hostname | 6.5K | 750 | 11.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 21.7K | 9.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 25 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 22 | 6.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 925 | 6.1% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 581 | 4.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 3.6K | 4.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.4K | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 66 | 4.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 77 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 57 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 79 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 216 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 909 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 1 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 262 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 170 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 24 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 111 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 163 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 78 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 76 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 44 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 416 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 182 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.0K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 841 | 63.0% |
| EasyList | blocklist | adguard | 66.5K | 32.8K | 49.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 483 | 36.9% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 51.0K | 28.1% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 55.3K | 23.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 21.7K | 9.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1.4K | 4.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 66 | 4.0% |
| abpvn_hosts | blocklist | adguard | 1.0K | 29 | 2.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 79 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 68 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 243 | Targets: 18 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 172 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 3 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 31 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 13 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 5 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 60 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 2 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 68 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 68 | 67.3% |
| Torrent Trackers | blocklist | domain | 495 | 67 | 13.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 34.1K | Targets: 78 | Unique: 0 | Conflicts: 165</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 87 | 21.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 690 | 19.5% |
| tranco | allowlist | domain_top | 500 | 81 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 573 | 13.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 47 | 9.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.5K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 521 | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1.9K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 7 | 6.6% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 143 | 5.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 13 | 5.3% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 34 | 5.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1.4K | 2.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| HaGeZi Pro | blocklist | domain | 195.7K | 4.3K | 2.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 280 | 2.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1.7K | 2.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 4.0K | 1.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 6 | 1.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 27 | 1.6% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 558 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 64 | 1.4% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 528 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2.4K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 2.1K | 0.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 53 | 0.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 4.3K | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| phishing_army | blocklist | domain | 140.3K | 663 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 2.6K | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 193 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 424 | 0.4% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 58 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 47 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| kadantiscam | blocklist | domain | 40.0K | 70 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.8K | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 22 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 284 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 846 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 78 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 112 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 59 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 34.1K | Targets: 24 | Unique: 20.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 513 | 18.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 34 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 1.4K | 2.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 4.0K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 2.4K | 1.3% |
| EasyList | blocklist | adguard | 66.5K | 653 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 2.1K | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.4K | 342 | 0.6% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 773 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 58 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 47 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 1.6K | 0.3% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 140.3K | Targets: 37 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 17.5K | 50.4% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 101.8K | 42.2% |
| kadantiscam | blocklist | domain | 40.0K | 15.9K | 39.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 72.7K | 30.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 72.2K | 30.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 60 | 24.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 12.7K | 15.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 24 | 4.8% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 663 | 1.9% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3.4K | 1.7% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1.5K | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 34 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 28 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 1 | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 35 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 495 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 264 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 437 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 89 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 355 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 150 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 23 | Unique: 61.7K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 95 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 92 | 4.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 3 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 48 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 2 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 5 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 3 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 290 | 17.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 142 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 76 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 145 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 296 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 299 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 46 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 47 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 298 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 25 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 59 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 47 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 10 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 28 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 20 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 18 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 22 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 75 | 0.0% |

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
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 34 | 8.9% |
| tranco | allowlist | domain_top | 500 | 39 | 7.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 8 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 434 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.3K | 3.7% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 904 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 314 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 925 | 1.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 6 | 1.7% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 11 | 1.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1.2K | 1.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 1 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 996 | 1.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 2.4K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2.1K | 1.2% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 1 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.4K | 0.6% |
| Torrent Trackers | blocklist | domain | 495 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 183 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 585 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 109 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 23 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 12 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 2 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 6 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 66.5K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 624.1K | Targets: 15 | Unique: 368.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 2.5K | 91.7% |
| RPiList_specials-phishing | blocklist | adguard | 143.4K | 100.1K | 69.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 75.6K | 31.5% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 75.0K | 31.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1.6K | 4.6% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 13 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 68 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 179 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 13 | 0.1% |
| EasyList | blocklist | adguard | 66.5K | 45 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 3 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 143.4K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 237.9K | 73.6K | 30.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 73.3K | 30.5% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 100.1K | 16.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 773 | 2.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.8K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 7 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.6K | Targets: 27 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 382 | 107 | 28.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 406 | 9.4% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1.0K | 6.6% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 33 | 6.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 27 | 4.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 1.0K | 2.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 210 | 2.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 61 | 2.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 881 | 1.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 242 | 1.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 24 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 1 | 0.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 102 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 55 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 14 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 38 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 25 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 3 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 14 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 8 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 3 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 1 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 19 | Unique: 368 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 23 | 18.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 176 | 15.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 30 | 2.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 7 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 177 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 75 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 263 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 14 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 40 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 3 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 15 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 8 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 4 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 19 | Unique: 2.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 382 | 87 | 22.8% |
| Greensnow | blocklist | ipv4 | 4.3K | 77 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 24 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 99 | 0.6% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 2 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 98 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 1 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 139 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 7 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 8 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 9 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 7 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 2 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.4K | Targets: 21 | Unique: 0 | Conflicts: 12.7K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 11.4K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 11.4K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 656 | 4.2% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 32 | 1.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 210 | 1.2% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 60 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 9 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 36 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 19 | 0.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 7 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 38 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 33 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 42 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 10 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.2K | Targets: 26 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 645 | 20.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 210 | 13.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 4.4K | 12.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 567 | 11.1% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1.6K | 10.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 8.5% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 31 | 8.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 1.4K | 7.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 79 | 6.9% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 762 | 6.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 4.2K | 6.9% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 37 | 6.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 1.5K | 6.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 269 | 6.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1.0K | 5.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 27 | 4.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 274 | 1.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 15 | 1.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 91 | 0.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 19 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 19 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 3 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 50 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 19.5K | 30 | 0.2% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 9 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 30 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 29 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 9 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 10 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 718 | Targets: 40 | Unique: 331 | Conflicts: 314</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 7 | 1 | 14.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 56 | 3.3% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 7 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 2 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 40 | 0.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 2 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 2 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 11 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 27 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 31 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.1K | Targets: 55 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 500 | 30.3% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 979 | 27.6% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 115 | 17.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 672 | 15.7% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.7K | 9.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 7 | 6.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 24 | 6.3% |
| Adaway | blocklist | hostname | 6.5K | 408 | 6.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 904 | 5.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3.3K | 5.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 21 | 5.7% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.9K | 5.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| tranco | allowlist | domain_top | 500 | 23 | 4.6% |
| HaGeZi Pro | blocklist | domain | 195.7K | 8.2K | 4.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 6.0K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.3K | 3.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 359 | 2.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1.9K | 2.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 5.8K | 2.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1.1K | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 125 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 165 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 44 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 13 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 204 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 13 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 27 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 17 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 52 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 286 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 35 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 232.0K | Targets: 35 | Unique: 174.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 13.2K | 68.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.6K | 28.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 480 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 78 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 458 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 194 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 72 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 308 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 320 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 108 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 23 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 4 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 76 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 10 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 31 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 210 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 121 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 152 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 65 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 29 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.3K | Targets: 45 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 5.7K | 10.1% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 44 | 6.7% |
| HaGeZi Pro | blocklist | domain | 195.7K | 11.5K | 5.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 934 | 5.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 12.2K | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 10.7K | 4.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 5.8K | 3.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.1K | 3.0% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 55 | 1.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 183 | 1.2% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 95 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 193 | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 313 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 1 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 264 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3.7K | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| kadantiscam | blocklist | domain | 40.0K | 68 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1.2K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 195 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 371 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 2 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 42 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 13 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 12 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.4K | Targets: 32 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 899 | 1.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 995 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 245 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 24 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 147 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 307 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 88 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 13 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 16 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 26 | Unique: 1.3K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 3 | 1.2% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 53 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 18 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 28 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 49 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 24 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 8 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 5.9K | 30.5% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 19.4K | 8.4% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 2 | 2.0% |
| Torrent Trackers | blocklist | domain | 495 | 9 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 615 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 141 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 125 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 216 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 521 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 78 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 390 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 277 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 23 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 129 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 44 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 29 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 8 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 15 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 39 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 15 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 2.7K | 3.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 143 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1.2K | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 24 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 35 | Unique: 0 | Conflicts: 93</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 29 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 17 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 38 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 32 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 61 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 32 | Unique: 4.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 1.6K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 1.0K | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 20 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 120 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 52 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 22 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 110 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 15 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 7 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 56 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.7K | 98.8% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.5K | 33.4% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 27 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 17.6K | 7.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 4.3K | 1.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 2.5K | 1.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 2 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 5.0K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 195.7K | 426 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 389 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 111 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 7 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 22 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 81.8K | Targets: 69 | Unique: 0 | Conflicts: 77</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3.5K | 98.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 352 | 88.4% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 40.0K | 30.1K | 75.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 50.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.6K | 41.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 400 | 24.2% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 66 | 17.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 49 | 13.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 899 | 12.1% |
| phishing_army | blocklist | domain | 140.3K | 12.7K | 9.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 20.7K | 8.7% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 9 | 8.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1.9K | 6.9% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3.6K | 6.4% |
| HaGeZi Pro | blocklist | domain | 195.7K | 10.6K | 5.4% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 12.7K | 5.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 12.0K | 5.0% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 33 | 5.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.7K | 4.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 4.9K | 2.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 340 | 2.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 313 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 3.9K | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 811 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 41 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 291 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 40 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 84 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.3K | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 1.0K | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 414 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 22 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 194 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 108 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 6.5K | 33.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 21.6K | 9.3% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 73 | 2.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 2 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 495 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 703 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 157 | 0.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 165 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 262 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 95 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 195.7K | 611 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 13 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 455 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 346 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 158 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 50 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 35 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 48 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 18 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 34 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 14 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 35 | Unique: 0 | Conflicts: 93</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 29 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 15 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 32 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 61 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 38 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 35 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 37.9K | Targets: 30 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 13.5K | 5.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 29.5K | 5.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 18 | 4.5% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 3.2K | 3.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 528 | 1.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 15.4K | 1.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 2 | 0.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 9 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 568 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 2.6K | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 36 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 3 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 262 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 144 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 22 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 88 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 35 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 80 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 544.7K | Targets: 50 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 29.5K | 78.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 398 | 239 | 60.1% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 30.3K | 31.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 271.0K | 22.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 220.7K | 21.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 4.3K | 12.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 18.7K | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 6.2K | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 2.3K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 6 | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 371 | 0.8% |
| HaGeZi Pro | blocklist | domain | 195.7K | 1.6K | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 414 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 738 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 24 | 0.4% |
| phishing_army | blocklist | domain | 140.3K | 495 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 115 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 163 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 13 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 6 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 274 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 20 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| kadantiscam | blocklist | domain | 40.0K | 24 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 856 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 35 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 152 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 17.5K | Targets: 30 | Unique: 0 | Conflicts: 38</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 544 | 47.2% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 6.4K | 41.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.5K | 29.9% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 2.4K | 21.6% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 86 | 15.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 242 | 15.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.4K | 14.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 453 | 14.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 79 | 13.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 2.8K | 11.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 4.1K | 11.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 478 | 11.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 5.3K | 8.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 315 | 6.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 978 | 5.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 59 | 5.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 40 | 3.3% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 12 | 3.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 3 | 2.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 212 | 1.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 14 | 1.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 68 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 9 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 38 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 38 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 130 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 495 | Targets: 9 | Unique: 304 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 101 | 100 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 67 | 98.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 47 | Unique: 0 | Conflicts: 581</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 122 | 7.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 2 | 4.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 32 | 0.9% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 29 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 243 | 1 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 39 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 81 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 34 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 1 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 23 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 11 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 27 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 33 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 16 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 35 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 66.5K | 51 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 30 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 38 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 624.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 3 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 11 | Unique: 24 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.4K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 850 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 43 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 66.5K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 66 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 398 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 352 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 87 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 278 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 333 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 271 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 43 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 28 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 38 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 239 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 26 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.5K | Targets: 25 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 14.8K | 41.3% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 2 | 4.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 274 | 3.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 212 | 1.2% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 5 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 3 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 8 | 0.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 4 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 92 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 128 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 14 | 0.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 18 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 249 | 0.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 3 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 235 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 25 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 23 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 13 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 12 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 1 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 72.6K | Targets: 1 | Unique: 72.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.7K | Targets: 35 | Unique: 8.4K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 281 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 51 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 382 | 13 | 3.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5.3K | 2.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.1K | 17 | 1.6% |
| Greensnow | blocklist | ipv4 | 4.3K | 60 | 1.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 14 | 1.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 15 | 1.2% |
| BruteforceBlocker | blocklist | ipv4_find | 545 | 6 | 1.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 50 | 1.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 91 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 14 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 305 | 0.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 600 | 5 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 123 | 1 | 0.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 26 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 412 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 92 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 70 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 79 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 64 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 67 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 68 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 103 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 4 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 2.1K | 0.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 3.9K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 559 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 362 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 442 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 70 | 0.1% |
| HaGeZi Pro | blocklist | domain | 195.7K | 127 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 2 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 9 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 13 | Unique: 385 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 281 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 8 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.0K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.7K | 4 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 15 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 72.6K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.0K | 45 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 15.4K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 23 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.5K | 2 | 0.0% |

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
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 19 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 83 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 1 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 54 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 137 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 85 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 19 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 35 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 20 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 84 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 66 | Unique: 0 | Conflicts: 94</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 2.6K | 73.8% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3.0K | 69.9% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.2K | 33.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 20.7K | 25.3% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 191 | 11.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 40 | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 9 | 8.5% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 15.9K | 6.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 996 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 1.8K | 6.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 3.1K | 5.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 1.4K | 4.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 25 | 3.8% |
| HaGeZi Pro | blocklist | domain | 195.7K | 6.8K | 3.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 5.6K | 2.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 3.2K | 1.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 130 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 161 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2.5K | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 336 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 977 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 22 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 51 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 22 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 36 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 4 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 33 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 41 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 15 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### YousList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 625 | Targets: 34 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 9 | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 198 | 1.1% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 3 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 24 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 241 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 86 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 2 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 95 | 0.2% |
| HaGeZi Pro | blocklist | domain | 195.7K | 201 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 270 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 42 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 151 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 17 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 25 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 5 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| EasyList | blocklist | adguard | 66.5K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 239.9K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 237.9K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.0K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 181.6K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.4K | 10 | 0.0% |

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
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 50 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.5K | 8 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 74 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 195.7K | 41 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 45 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.5K | Targets: 62 | Unique: 0 | Conflicts: 45</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.8K | 42.7% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 392 | 23.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.0K | 10.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 81.8K | 3.5K | 4.3% |
| Adaway | blocklist | hostname | 6.5K | 263 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 575 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 4 | 3.8% |
| YousList | blocklist | hostname | 625 | 24 | 3.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.1K | 979 | 3.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 421 | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.0K | 1.6K | 2.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 657 | 13 | 2.0% |
| ph00lt0_blocklist | blocklist | domain | 34.1K | 690 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Pro | blocklist | domain | 195.7K | 3.0K | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 5 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 237.9K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 180.9K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 55 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 16 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.0M | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.0K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 140.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 241.0K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 227.5K | 30 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.4K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 37.9K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 544.7K | 6 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 580.9K | 11 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 34.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 239.9K | 64 | 0.0% |
| kadantiscam | blocklist | domain | 40.0K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 94.7K | 7 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 218.3K | 9 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.7K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

