# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-10-17 13:25:43 UTC

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
| Total Entries Analyzed | 9.5M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 143 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| domain | 88 |
| ipv4 | 40 |
| adguard | 34 |
| cidr_ipv4 | 3 |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 67.9% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 11.3K | 60.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 33.5K | 59.0% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 310 | 47.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 167 | 45.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 142 | 37.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 588 | 35.6% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 32.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 110 | 31.9% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 33 | 31.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 57.5K | 28.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 44.3K | 24.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 6.3K | 22.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 40.3K | 17.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.4K | 16.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 13.2K | 15.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 6.1K | 13.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 4.4K | 12.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 20.5K | 8.9% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 147 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 3.3K | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 243 | 1.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 26 | 0.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 2 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 455 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 145 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 43 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| kadantiscam | blocklist | domain | 39.2K | 86 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 459 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 433 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 44 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 630 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 738 | 0.1% |
| phishing_army | blocklist | domain | 157.2K | 99 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.1K | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 224 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 913 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 29 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 22 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 35 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 5 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 2 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 6.5K | 7.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 266 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 6 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 21 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 434 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 9 | 2.3% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 10 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 523 | 1.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 408 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 749 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.7K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 87 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1.0K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 50 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 26 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 34 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 14 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 1 | 0.0% |

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
| Easy Privacy | blocklist | adguard | 55.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 35 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 655 | Targets: 32 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 486 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 114 | 0.4% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 633 | 0.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 551 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 55 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 490 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 310 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 35 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 44 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 6 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 112 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 32 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 8 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 486 | 0.9% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 652 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 490 | 0.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 35 | 0.1% |
| EasyList | blocklist | adguard | 68.8K | 91 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 44 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 112 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 6 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 228.4K | Targets: 14 | Unique: 227.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 433 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.6K | Targets: 24 | Unique: 129.0K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 93.1K | 50.4% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.2K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 585 | 3.8% |
| HaGeZi Pro | blocklist | domain | 201.5K | 6.8K | 3.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 908 | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 825 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.8K | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 205 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 84 | 0.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 349 | 0.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.2K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 2 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 60 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 185.6K | Targets: 26 | Unique: 0 | Conflicts: 200</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 203 | 199 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.9% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 51.6K | 91.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 68.8K | 49.2K | 71.5% |
| Easy Privacy | blocklist | adguard | 55.5K | 28.7K | 51.7% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 652 | 49.9% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 56.5K | 24.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 25.5K | 11.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 2.4K | 6.8% |
| abpvn_hosts | blocklist | adguard | 1.0K | 22 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 113 | 0.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 292 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 13 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 184.8K | Targets: 69 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 655 | 633 | 96.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 93.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 51.6K | 91.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.3% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 93.1K | 38.7% |
| HaGeZi Pro | blocklist | domain | 201.5K | 75.7K | 37.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5.0K | 27.1% |
| YousList | blocklist | hostname | 625 | 151 | 24.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 56.5K | 24.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 966 | 22.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 44.3K | 21.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 6.1K | 21.5% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 21 | 19.8% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5.9K | 12.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.7K | 11.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 25.5K | 11.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 25 | 7.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2.4K | 6.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 5.0K | 6.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 20 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 554 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 113 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 346 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 277 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 192 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 636 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 273 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 195 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 29 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 364 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 125 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 107 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 274 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 5 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 3 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 12 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 842 | 1.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 826 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 72 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 77 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 7 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |

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
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |

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
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 10 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 2.1K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 25 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 181 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 3 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 2 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 57 | Unique: 0 | Conflicts: 48</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.0K | 56.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.8% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 625 | 198 | 31.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 204 | 12.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 12 | 11.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 7.6K | 9.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 55 | 8.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 4.7K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.7% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1.7K | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 18 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.5K | 4.4% |
| HaGeZi Pro | blocklist | domain | 201.5K | 7.1K | 3.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 6.4K | 2.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 5.0K | 2.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 935 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.2K | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 52 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 18 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 14 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 29 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 43 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 142 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 10 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 327 | Targets: 21 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 249 | 11 | 4.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 14 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 68 | 0.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 24 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 73 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 187 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 61 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 87 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 2 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 31 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 156 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 45 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.2K | 8 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 22 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 6 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 16.7K | 869 | 5.2% |
| Greensnow | blocklist | ipv4 | 4.2K | 187 | 4.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 405 | 2.4% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 5 | 2.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 65 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 6 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 79 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 174 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 1 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 16 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 1 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 152 | 0.2% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 1 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 31 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 14 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 16 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 8 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 249 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 116 | 6.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 11 | 3.4% |
| Greensnow | blocklist | ipv4 | 4.2K | 124 | 3.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 56 | 2.2% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 205 | 1.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 4 | 0.8% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 4 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 211 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 5 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 226 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 41 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 22 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 10 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 8 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 3 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 1 | 0.0% |

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
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 188 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 45 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 76 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 189 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 22 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 34 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 264 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 126 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 53 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 105 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 41 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 6 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 51 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 285 | 68.2% |
| phishing_army | blocklist | domain | 157.2K | 103.7K | 66.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 74.2K | 32.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 74.9K | 32.0% |
| kadantiscam | blocklist | domain | 39.2K | 9.3K | 23.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 12.4K | 14.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1.5K | 5.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 3.1K | 3.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 4.4K | 2.2% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 575 | 1.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.8K | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18.7K | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 380 | 1.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 6.3K | 1.1% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 1 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 196 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 16 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 44 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 25 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 500 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 107 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 12 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 7 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 29 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 57 | Unique: 0 | Conflicts: 71</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.8K | 51.1% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 48 | 13.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 166 | 10.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.9K | 10.1% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 40 | 5.6% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 514 | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 13 | 3.7% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 473 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 11 | 2.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.6K | 2.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 2.2K | 2.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 672 | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 574 | 1.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 6 | 1.6% |
| HaGeZi Pro | blocklist | domain | 201.5K | 2.1K | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 5 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.9K | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 966 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 25 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 42 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 113 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 19 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 18 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 24 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 5 | 2.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 65 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 23 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 6 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 15 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 55 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 37 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 60.9K | Targets: 32 | Unique: 2.5K | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 249 | 226 | 90.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 673 | 58.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8.7K | 57.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 946 | 52.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 156 | 47.7% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 4.6K | 47.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.3K | 44.8% |
| Greensnow | blocklist | ipv4 | 4.2K | 1.8K | 43.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 4.2K | 41.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 7.7K | 32.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 5.4K | 32.3% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 5.2K | 31.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 10.9K | 29.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 136 | 27.4% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 128 | 26.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 249 | 21.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 3.5K | 20.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 152 | 14.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 76 | 7.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 20 | 7.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 9 | 4.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 416 | 2.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 252 | 1.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 38 | 1.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
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
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 1 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 4 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 493 | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 450 | 90.5% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 463 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 30 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 4 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 165 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 268 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 2 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 99 | 0.6% |
| Greensnow | blocklist | ipv4 | 4.2K | 26 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 55 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 128 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 26 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 8 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 16 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 102 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 67 | 98.5% |
| Torrent Trackers | blocklist | domain | 495 | 101 | 20.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 23 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 9.6K | 5.7K | 59.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 4.7K | 28.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 5.9K | 24.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.2K | 22.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 73 | 22.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 5.7K | 15.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.6K | 15.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8.7K | 14.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 1.9K | 11.0% |
| Greensnow | blocklist | ipv4 | 4.2K | 87 | 2.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 10 | 2.0% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 10 | 2.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 32 | 1.8% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 291 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 3 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 5 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 63 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 11 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 30 | 0.0% |

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
| OISD Blocklist Small | blocklist | adguard | 56.8K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 56 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 1 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.9K | Targets: 49 | Unique: 12.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 1.9K | 6.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 25 | 6.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 30.6K | 5.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 12 | 5.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 22.7K | 1.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.1K | 1.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 71 | 1.4% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 435 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 166 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.1K | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 463 | 1.3% |
| phishing_army | blocklist | domain | 157.2K | 1.6K | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 2.2K | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 2.2K | 0.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 118 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 264 | 0.6% |
| HaGeZi Pro | blocklist | domain | 201.5K | 801 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 298 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 174 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 29 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 47 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 455 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| kadantiscam | blocklist | domain | 39.2K | 31 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 608 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 39 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 192 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 286 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 78 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 36 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 13.0K | 15.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 46 | 12.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 425 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 359 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 8 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 2.5K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 582 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 280 | 0.8% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.4K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 554 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 627 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 166 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 95 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 690 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 29 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 14 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 635 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 19 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 15 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 84 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 6 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 36 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |

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
| RPiList_specials-malware | blocklist | adguard | 751.4K | 1 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 7 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 19 | Unique: 33 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 174 | 15.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 673 | 1.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 51 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 123 | 0.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 30 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 21 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.2K | 3 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 2 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 22 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 10 | Unique: 303 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 97.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 81 | 11.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 93 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1.0K | 31.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 884 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 8 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 79 | 5.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 81 | 4.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 39 | Unique: 11.4K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3.0K | 90.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 884 | 77.9% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 201.5K | 131 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 47 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 16 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 6 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 36 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 526 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 10 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 63 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 9 | Unique: 13.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3.0K | 90.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 47 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 18 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 2 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 4.9K | 27.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 5.1K | 21.7% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 2.9K | 7.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 7.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 24 | 7.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 1.2K | 7.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 680 | 6.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 2.3K | 3.8% |
| Greensnow | blocklist | ipv4 | 4.2K | 118 | 2.8% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 6 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 16 | 1.5% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 6 | 1.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 6 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 9 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 30 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 12 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 28 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.5K | Targets: 20 | Unique: 13.3K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.7% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 28.7K | 15.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.8% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 2.7K | 4.7% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 44 | 3.4% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 5.8K | 2.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 342 | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 825 | 0.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| EasyList | blocklist | adguard | 68.8K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 8 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 68.8K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 33.3K | 58.8% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 49.2K | 26.5% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 33.7K | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 18.9K | 8.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 91 | 7.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 659 | 1.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 64 | 0.3% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 132 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 497 | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 493 | 450 | 91.3% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 429 | 4.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 28 | 1.6% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 4 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 132 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 249 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 2 | 0.6% |
| Greensnow | blocklist | ipv4 | 4.2K | 26 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 51 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 92 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 136 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 26 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 16 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 8 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.3% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.5% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 5 | 0.0% |

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
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 42 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 8 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |

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
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 49 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 23 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 23 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 51 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 2 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 205 | Targets: 11 | Unique: 138 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 31 | 2.7% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 11 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 8 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 9 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 8 | Unique: 471 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 11 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 8 | Unique: 1 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 6 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.0K | Targets: 19 | Unique: 677 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 68 | 2.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 30 | 2.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 6 | 0.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.2K | 12 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 24 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 17 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 22 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 45 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 76 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 6 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.6K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.7% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.6% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.7K | Targets: 31 | Unique: 0 | Conflicts: 1.7K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 4.9K | 95.0% |
| Greensnow | blocklist | ipv4 | 4.2K | 3.5K | 84.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 869 | 82.7% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 205 | 82.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.3K | 72.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 5.9K | 35.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 8.0K | 33.8% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 165 | 33.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 132 | 26.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.8K | 17.6% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 5.9K | 16.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1.7K | 15.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1.7K | 15.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 45 | 13.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.9K | 12.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 5.2K | 8.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 11 | 5.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 439 | 4.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 44 | 3.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 95 | 3.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 22 | 2.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 54 | 1 | 1.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 21 | 1.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 69 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 5 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 22 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 63 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.6K | Targets: 34 | Unique: 0 | Conflicts: 38</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 463 | 93.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 429 | 86.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 8.0K | 45.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.9K | 39.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 87 | 26.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 3.8K | 22.9% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 7.2K | 19.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.8K | 17.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 7.7K | 12.6% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 22 | 8.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Greensnow | blocklist | ipv4 | 4.2K | 180 | 4.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 31 | 2.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 50 | 2.8% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 439 | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 24 | 2.3% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 1 | 1.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 132 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 99 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 35 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 35 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 358 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| DoH_IP_list | blocklist | ipv4 | 731 | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 14 | Unique: 2.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 56 | 20.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 68 | 6.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 38 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 7 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 9 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.2K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 5 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 279 | Targets: 8 | Unique: 194 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 56 | 2.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 20 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 3 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.6K | Targets: 20 | Unique: 5.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 311 | 2.0% |
| Adaway | blocklist | hostname | 6.5K | 87 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 2.2K | 0.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1.7K | 0.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.6K | 0.8% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 25 | 0.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 339 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 16 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 52 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 130 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 19 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 14 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 56 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 106 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 32 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 14 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 395 | 11.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 166 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 502 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 204 | 1.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.6K | 0.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 400 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 588 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 189 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 66 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 168 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 27 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 10 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.6% |
| Easy Privacy | blocklist | adguard | 55.5K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 168 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 27 | 0.1% |
| EasyList | blocklist | adguard | 68.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 1 | 0.0% |

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
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 13 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.2K | Targets: 27 | Unique: 0 | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 249 | 124 | 49.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 414 | 23.1% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 3.5K | 21.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 187 | 17.8% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 2.0K | 5.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 539 | 5.4% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 26 | 5.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 26 | 5.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 603 | 3.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 1.8K | 3.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 72 | 2.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 479 | 2.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 8 | 2.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 118 | 2.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 12 | 1.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 180 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 73 | 0.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 87 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 47 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 12 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 12 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 5 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 18 | 0.0% |

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
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 334 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 49 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 167 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 6 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 20 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 106 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 23 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 66 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 229.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 234.2K | 139.7K | 59.7% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 71.3K | 42.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 853 | 40.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 21.7K | 38.2% |
| EasyList | blocklist | adguard | 68.8K | 18.9K | 27.4% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 25.5K | 13.7% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 89.6K | 11.9% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 168 | 10.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 112 | 8.6% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 181 | 6.6% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1.6K | 4.6% |
| Easy Privacy | blocklist | adguard | 55.5K | 825 | 1.5% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 9 | 0.7% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 7 | 0.5% |
| abpvn_hosts | blocklist | adguard | 1.0K | 5 | 0.5% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 9 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 2 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 229.1K | Targets: 53 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 273 | 65.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 139.7K | 59.7% |
| phishing_army | blocklist | domain | 157.2K | 71.1K | 45.2% |
| HaGeZi Pro | blocklist | domain | 201.5K | 82.6K | 41.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 21.7K | 38.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 74.2K | 30.9% |
| kadantiscam | blocklist | domain | 39.2K | 11.0K | 28.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 12.1K | 26.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 24 | 19.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 112 | 17.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 5.2K | 17.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 13.1K | 15.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 25.5K | 13.8% |
| Spam404 | blocklist | domain | 8.1K | 1.1K | 13.0% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 168 | 10.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 20.5K | 10.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.6K | 7.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 362 | 7.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.5% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 627 | 4.8% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.6K | 4.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 10 | 4.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 302 | 4.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1.1K | 4.0% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 113 | 2.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 2.2K | 2.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 62 | 1.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 5.9K | 1.1% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 105 | 0.9% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 1.8K | 0.8% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 237 | 0.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 109 | 0.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 9 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 49 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.0K | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 6.6K | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 34 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.0K | 0.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1.8K | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 130 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 32 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 158 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 319 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 6 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.3K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.2% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 41 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.3K | Targets: 5 | Unique: 300 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.2% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 9 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 601.2K | Targets: 43 | Unique: 583.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2.6K | 7.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 3.9K | 4.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.8K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.5K | 0.7% |
| HaGeZi Pro | blocklist | domain | 201.5K | 915 | 0.5% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 3 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 88 | 0.3% |
| kadantiscam | blocklist | domain | 39.2K | 107 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.6K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 286 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 500 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 33 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 879 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| phishing_army | blocklist | domain | 157.2K | 141 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 30 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 42 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 125 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 19 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 224 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 6 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 81 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 42 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 494 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 30 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 551 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 383 | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 46 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 10 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| HaGeZi Pro | blocklist | domain | 201.5K | 331 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 40 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 66 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 24 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 20 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 55 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 10 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 201.5K | Targets: 70 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 342 | 99.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 54.7K | 96.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 334 | 90.8% |
| hufilter | blocklist | hostname | 92 | 81 | 88.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 331 | 86.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 551 | 84.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.7% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 66 | 62.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 77 | 61.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.1K | 50.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 75.7K | 41.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.1K | 38.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 110 | 37.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 82.6K | 36.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 80.4K | 34.3% |
| YousList | blocklist | hostname | 625 | 201 | 32.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 8.3K | 29.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 57.5K | 28.3% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 25.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 11.5K | 24.8% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.4K | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 11.8K | 14.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 4.4K | 12.6% |
| kadantiscam | blocklist | domain | 39.2K | 4.2K | 10.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.4K | 10.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.6K | 10.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 37 | 8.9% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 248 | 3.3% |
| phishing_army | blocklist | domain | 157.2K | 4.9K | 3.1% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 6.8K | 2.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 134 | 2.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 451 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 613 | 2.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4.4K | 1.9% |
| Spam404 | blocklist | domain | 8.1K | 154 | 1.9% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 189 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 3 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 227 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 41 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 523 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 24 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 614 | 0.8% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 801 | 0.8% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 131 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 51 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.6K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.3K | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1.6K | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 85 | 0.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 3 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 915 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 474 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.9K | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 41 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 345 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 342 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 21 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 88 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 16 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 22 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 3 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 10 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 70.6% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 79 | 10.8% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 94 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 10 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 36.6K | Targets: 35 | Unique: 0 | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 14.8K | 93.6% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 211 | 84.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.3K | 69.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 187 | 57.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.9K | 56.2% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 268 | 54.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 249 | 50.1% |
| Greensnow | blocklist | ipv4 | 4.2K | 2.0K | 49.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 4.9K | 49.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.7K | 38.3% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 5.9K | 35.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 5.3K | 31.9% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 2.9K | 30.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 7.2K | 30.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 4.5K | 25.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 10.9K | 17.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 174 | 16.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 164 | 14.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 6 | 13.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 123 | 10.6% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 45 | 4.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 8 | 3.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 101 | 3.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 285 | 1.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 3 | 1.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 9 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 14 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 473 | 0.2% |
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
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 4 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 23 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| HaGeZi Pro | blocklist | domain | 201.5K | 110 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 20 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 53 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 108 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 13 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 17 | 0.0% |

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
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 85 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 81 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 70 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 83 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 12 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 58 | Unique: 420.6K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 995 | 13.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1.5K | 4.9% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 9 | 2.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 3 | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 3.0K | 1.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 2.6K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 3.0K | 1.3% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| phishing_army | blocklist | domain | 157.2K | 1.4K | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 608 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 98 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 2.4K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.4K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.2K | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 117 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1.6K | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 181 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 53 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 189 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 21 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 210 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 21 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 43 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 273 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 8 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 39.2K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 29.7K | 35.7% |
| phishing_army | blocklist | domain | 157.2K | 16.9K | 10.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 11.0K | 4.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 11.0K | 4.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9.3K | 3.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 201.5K | 4.2K | 2.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 479 | 1.6% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 75 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 29 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 68 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 31 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 86 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 14 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 24 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 43 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 22 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 107 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 29 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 2 | 0.9% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 17 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 47 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 24 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 29 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 10 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 185.6K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 3 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
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
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |

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
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
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
| HaGeZi Pro | blocklist | domain | 201.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |

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
| Firehol_level3 | blocklist | ipv4 | 9.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.1K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 534 | 1.5% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 1.2K | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 853 | 0.4% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 2.1K | 0.3% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 5 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2M | Targets: 61 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 975.4K | 92.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 271.0K | 49.4% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 13.6K | 44.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 22.7K | 23.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 41 | 9.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 3.7K | 7.9% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 18.7K | 7.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 11 | 6.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.8K | 5.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 690 | 5.3% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 526 | 3.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 6.6K | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 6.3K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 4.3K | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1.3K | 1.5% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 8 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 286 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 412 | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.1K | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 25 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 636 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| phishing_army | blocklist | domain | 157.2K | 452 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 76 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 121 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 13 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 551 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 15 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 29 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 34 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 215.2K | Targets: 37 | Unique: 203.3K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 85.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5.3K | 33.6% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 8 | 3.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 22 | 1.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 6 | 1.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 358 | 1.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 224 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 473 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 5 | 1.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 5 | 1.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 145 | 0.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 8 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 420 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 8 | 0.7% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 10 | 0.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 6 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 28 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 10 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 51 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 47 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 63 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.2K | 18 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 49 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 6 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 30 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 48 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1M | Targets: 49 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 975.4K | 81.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 222.4K | 40.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 16.1K | 16.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 36 | 8.6% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2.4K | 7.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 839 | 2.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 4.2K | 1.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 4.0K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1.0K | 1.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.4K | 0.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.9K | 0.9% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 91 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 630 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 364 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 111 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| phishing_army | blocklist | domain | 157.2K | 373 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 494 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 17 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 19 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 8 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 14 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 65 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 30.5K | Targets: 33 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 187 | 86.2% |
| phishing_army | blocklist | domain | 157.2K | 26.2K | 16.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 21.8K | 9.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 17 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 5.2K | 2.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 509 | 1.5% |
| kadantiscam | blocklist | domain | 39.2K | 479 | 1.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.5K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 23 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 435 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.5K | 0.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 613 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 17 | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 34 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 86 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 91 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 9 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 111 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 44 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 76 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 12 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 2 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 234.2K | Targets: 65 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 56.1K | 98.8% |
| hufilter | blocklist | hostname | 92 | 83 | 90.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 343 | 82.1% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 490 | 74.8% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 21.8K | 71.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 139.7K | 61.0% |
| phishing_army | blocklist | domain | 157.2K | 84.2K | 53.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.7% |
| YousList | blocklist | hostname | 625 | 270 | 43.2% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 80.4K | 39.9% |
| WaLLy3K | blocklist | domain | 351 | 137 | 39.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.4K | 34.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 74.9K | 31.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 56.5K | 30.5% |
| kadantiscam | blocklist | domain | 39.2K | 11.0K | 28.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 21.5K | 25.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 88 | 25.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 10.5K | 22.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 23 | 21.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.9K | 20.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 5.8K | 20.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 40.3K | 19.9% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 55 | 14.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 4.0K | 11.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 561 | 11.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 36 | 9.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 11 | 5.1% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 2.2K | 2.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 188 | 1.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 86 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 106 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.0K | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 453 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 118 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 387 | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 184 | 0.6% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 1.2K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 6.3K | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 2.3K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.2K | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 9 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1.5K | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 36 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 306 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 234.2K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 56.1K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.3% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.2% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 826 | 61.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 139.7K | 61.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1.2K | 54.5% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 83.5K | 49.4% |
| EasyList | blocklist | adguard | 68.8K | 33.7K | 49.0% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 490 | 37.5% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 56.5K | 30.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 66 | 17.9% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 92.8K | 12.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 4.0K | 11.4% |
| Easy Privacy | blocklist | adguard | 55.5K | 5.8K | 10.4% |
| abpvn_hosts | blocklist | adguard | 1.0K | 35 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 56 | 3.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 118 | 0.6% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 36 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 19.2K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.9K | 9.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.4K | 8.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 12.9K | 5.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 25 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 59 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 47 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 113 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 75 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 243 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 227 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 118 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 47 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 39 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 22 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 8 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 21 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 19.2K | Targets: 11 | Unique: 18.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 29 | 2.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 59 | 0.2% |
| EasyList | blocklist | adguard | 68.8K | 64 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 113 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 118 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 75 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 2 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 14 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 56.8K | Targets: 64 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 85 | 92.4% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 486 | 74.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.5% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 37.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 51.6K | 27.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 54.7K | 27.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4.7K | 25.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 56.1K | 23.9% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 21 | 19.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33.5K | 16.5% |
| YousList | blocklist | hostname | 625 | 95 | 15.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5.6K | 12.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 3.3K | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 749 | 11.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 21.7K | 9.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 25 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 22 | 6.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 925 | 6.1% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 582 | 4.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 3.6K | 4.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.5K | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 66 | 4.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 76 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 216 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 908 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 56 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 75 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 262 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 174 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 24 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 181 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 81 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 160 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 44 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 76 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 412 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 111 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.8K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 842 | 63.0% |
| EasyList | blocklist | adguard | 68.8K | 33.3K | 48.5% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 486 | 37.2% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 51.6K | 27.8% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 56.1K | 23.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 21.7K | 9.5% |
| Easy Privacy | blocklist | adguard | 55.5K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1.5K | 4.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 66 | 4.0% |
| abpvn_hosts | blocklist | adguard | 1.0K | 29 | 2.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 75 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 153 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 217 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 187 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| phishing_army | blocklist | domain | 157.2K | 209 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 18 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 3 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 68 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 102 | 67 | 65.7% |
| Torrent Trackers | blocklist | domain | 495 | 67 | 13.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 34.7K | Targets: 78 | Unique: 0 | Conflicts: 165</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 114 | 27.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 699 | 19.6% |
| tranco | allowlist | domain_top | 500 | 81 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 574 | 13.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 47 | 9.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.5K | 8.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 18 | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 523 | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1.9K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 7 | 6.6% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 35 | 5.3% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.5K | 2.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| HaGeZi Pro | blocklist | domain | 201.5K | 4.4K | 2.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 280 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1.7K | 2.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 532 | 1.7% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 509 | 1.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 4.0K | 1.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 6 | 1.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 27 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 65 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2.4K | 1.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 54 | 0.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 4.4K | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.6K | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 463 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| phishing_army | blocklist | domain | 157.2K | 687 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 194 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 2.6K | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 47 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 59 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.8K | 0.2% |
| kadantiscam | blocklist | domain | 39.2K | 75 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 380 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Torrent Trackers | blocklist | domain | 495 | 1 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 839 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 19 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 60 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 78 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 34.7K | Targets: 24 | Unique: 19.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 534 | 25.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 72 | 5.4% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 35 | 2.7% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 1.5K | 2.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 4.0K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 2.4K | 1.3% |
| EasyList | blocklist | adguard | 68.8K | 659 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 1.6K | 0.7% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.5K | 342 | 0.6% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 3.8K | 0.5% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 681 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 47 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 59 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 157.2K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 209 | 96.3% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 26.2K | 86.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 103.7K | 43.2% |
| kadantiscam | blocklist | domain | 39.2K | 16.9K | 43.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 84.2K | 36.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 71.1K | 31.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 14.7K | 17.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 24 | 4.8% |
| HaGeZi Pro | blocklist | domain | 201.5K | 4.9K | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 687 | 2.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1.6K | 1.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 37 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 27 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 49 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 512 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 373 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 12 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 141 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 7 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 13 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 99 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 452 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 22 | Unique: 61.7K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 94 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 93 | 4.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 48 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 3 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 6 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 2 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 6 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.2K | 1 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 76 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 145 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 296 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 46 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 47 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 298 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 299 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 5 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 59 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 77 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 28 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 24 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 47 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 10 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 22 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 20 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 56 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 576 | 16.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 473 | 11.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 34 | 8.9% |
| tranco | allowlist | domain_top | 500 | 39 | 7.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 8 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 434 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.3K | 3.7% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 905 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 311 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 11 | 1.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 6 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 925 | 1.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1.2K | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 996 | 1.3% |
| HaGeZi Pro | blocklist | domain | 201.5K | 2.4K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2.1K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 102 | 1 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.4K | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Torrent Trackers | blocklist | domain | 495 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 183 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 585 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 109 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 68.8K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 751.4K | Targets: 17 | Unique: 414.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 2.1K | 98.6% |
| RPiList_specials-phishing | blocklist | adguard | 169.1K | 147.8K | 87.4% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 92.8K | 39.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 89.6K | 39.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 3.8K | 11.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 6 | 0.5% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 153 | 0.3% |
| EasyList | blocklist | adguard | 68.8K | 132 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 292 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 14 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 18 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 8 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 169.1K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 234.2K | 83.5K | 35.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 71.3K | 31.1% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 147.8K | 19.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 681 | 2.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 5 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 13 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 2 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.8K | Targets: 24 | Unique: 0 | Conflicts: 56</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 249 | 116 | 46.6% |
| Greensnow | blocklist | ipv4 | 4.2K | 414 | 10.0% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 1.3K | 7.8% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 30 | 6.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 28 | 5.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 14 | 4.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 351 | 3.5% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 1.3K | 3.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 323 | 1.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 946 | 1.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 23 | 0.9% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 50 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 90 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 56 | 0.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 56 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 38 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 9 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 32 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 13 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 6 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 19 | Unique: 399 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 31 | 15.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 174 | 15.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 30 | 2.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 2 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 249 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 164 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 44 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 3 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 17 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 19 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 8 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.2K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 2 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 19 | Unique: 2.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 249 | 56 | 22.5% |
| Greensnow | blocklist | ipv4 | 4.2K | 72 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 23 | 1.3% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 95 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 2 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 2 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 3 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 101 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 139 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 12 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 8 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 17 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.5K | Targets: 21 | Unique: 0 | Conflicts: 14.5K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 1.7K | 10.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 866 | 4.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 56 | 3.1% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 32 | 1.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 112 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 35 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.2K | 12 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 1 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 32 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 14 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 42 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 31 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 4 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.1K | Targets: 29 | Unique: 0 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 68 | 20.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 351 | 19.6% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 41 | 16.5% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 4.9K | 13.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 680 | 13.3% |
| Greensnow | blocklist | ipv4 | 4.2K | 539 | 13.0% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 55 | 11.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 1.8K | 10.9% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 1.8K | 10.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.6K | 10.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 51 | 10.3% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 992 | 10.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 1.8K | 7.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 1.3K | 7.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 4.2K | 6.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 65 | 6.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 231 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 13 | 1.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 4 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 51 | 0.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 14 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 14 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 47 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 19.2K | 29 | 0.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 2 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 29 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 29 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 15 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 718 | Targets: 39 | Unique: 333 | Conflicts: 312</summary>

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
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 40 | 0.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 2 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 27 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 6 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 31 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 12 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.2K | Targets: 55 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 502 | 30.4% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 982 | 27.5% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 114 | 17.4% |
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
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 905 | 5.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3.3K | 5.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 21 | 5.7% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.9K | 5.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| tranco | allowlist | domain_top | 500 | 23 | 4.6% |
| HaGeZi Pro | blocklist | domain | 201.5K | 8.3K | 4.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 6.1K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.3K | 3.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 359 | 2.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 5.8K | 2.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1.9K | 2.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.1K | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 165 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 47 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 125 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 205 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 14 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 11 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 286 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 13 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 53 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 232.5K | Targets: 35 | Unique: 175.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 12.9K | 66.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.6K | 28.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 474 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 459 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 78 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 78 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 76 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 107 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 195 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 319 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 22 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 306 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 153 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 29 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 121 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 30 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 7 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 210 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 65 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.3K | Targets: 45 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 5.6K | 9.9% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 44 | 6.7% |
| HaGeZi Pro | blocklist | domain | 201.5K | 11.5K | 5.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 12.1K | 5.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 935 | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 10.5K | 4.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 5.9K | 3.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.1K | 3.0% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 55 | 1.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 183 | 1.2% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 95 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 194 | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 313 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3.7K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 264 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| kadantiscam | blocklist | domain | 39.2K | 68 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 196 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 371 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 22 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 12 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 13 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 42 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.4K | Targets: 31 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 899 | 1.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 995 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 302 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 147 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 25 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 248 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 86 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 10 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 12 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 13 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 4 | 1.8% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 54 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 17 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 24 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 27 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 8 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 49 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 45 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 5.9K | 30.7% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 19.4K | 8.3% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 102 | 2 | 2.0% |
| Torrent Trackers | blocklist | domain | 495 | 9 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 615 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 216 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 125 | 0.4% |
| HaGeZi Pro | blocklist | domain | 201.5K | 523 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 23 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 78 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 387 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 277 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 44 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 130 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 29 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 39 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 16 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 11 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 2.7K | 3.2% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1.2K | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 24 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |

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
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 29 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 38 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 35 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 61 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 32 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
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
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 1.6K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 1.1K | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 20 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 52 | 0.1% |
| kadantiscam | blocklist | domain | 39.2K | 22 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 154 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 41 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 110 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 18 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 6 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 7 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 1 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.5% |

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
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 4.3K | 1.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 2.6K | 1.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.0K | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 2 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 201.5K | 451 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 118 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 389 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 22 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 17 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 83.0K | Targets: 69 | Unique: 0 | Conflicts: 77</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 98.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 363 | 86.8% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 39.2K | 29.7K | 75.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 50.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.6K | 41.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 400 | 24.2% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 66 | 17.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 49 | 13.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 899 | 12.1% |
| phishing_army | blocklist | domain | 157.2K | 14.7K | 9.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 21.5K | 9.2% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 9 | 8.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1.9K | 6.9% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3.6K | 6.4% |
| HaGeZi Pro | blocklist | domain | 201.5K | 11.8K | 5.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 13.1K | 5.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12.4K | 5.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 32 | 4.9% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.7K | 4.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 5.0K | 2.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 339 | 2.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 3.9K | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 313 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 41 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 298 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 86 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 825 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 39 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 428 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 30 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1.3K | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.0K | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 107 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 189 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 47 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 6.4K | 33.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 21.6K | 9.3% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| CF_Torrent_Trackers | blocklist | domain_url | 102 | 2 | 2.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 495 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 157 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 703 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 165 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 262 | 0.5% |
| HaGeZi Pro | blocklist | domain | 201.5K | 614 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 95 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 346 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 13 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 453 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 50 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 158 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 34 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 19 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 36 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 48 | 0.0% |

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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 29 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 15 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 32 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 61 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 38 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 17 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 30.9K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 25.6K | 4.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 17 | 4.1% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1.9K | 1.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 532 | 1.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 13.6K | 1.1% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 2 | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.4K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 575 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 9 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 184 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 145 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 237 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 34 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 85 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 88 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 30 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 5 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 49 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 548.6K | Targets: 50 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 25.6K | 82.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 418 | 251 | 60.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 30.6K | 31.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 271.0K | 22.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 222.4K | 21.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 4.4K | 12.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6.3K | 2.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 5.9K | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 217 | 3 | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 2.3K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 6 | 0.9% |
| HaGeZi Pro | blocklist | domain | 201.5K | 1.6K | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 371 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 428 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 111 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 24 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 738 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| phishing_army | blocklist | domain | 157.2K | 512 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 160 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 13 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 153 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 274 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 21 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 879 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 35 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 7 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.8K | Targets: 30 | Unique: 0 | Conflicts: 112</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 405 | 38.5% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 5.9K | 35.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.7K | 31.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.2K | 23.2% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 99 | 20.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 61 | 18.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 92 | 18.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.8K | 18.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 323 | 18.0% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 1.6K | 17.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 3.8K | 16.3% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 5.3K | 14.6% |
| Greensnow | blocklist | ipv4 | 4.2K | 603 | 14.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 1.9K | 11.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 5.4K | 8.9% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 10 | 4.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 30 | 2.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 54 | 1 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 17 | 1.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 19 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 159 | 1.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 112 | 1.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 2 | 1.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 112 | 1.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 12 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 279 | 1 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 59 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 7 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 145 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 495 | Targets: 9 | Unique: 303 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 102 | 101 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 67 | 98.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 46 | Unique: 0 | Conflicts: 580</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 122 | 7.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 2 | 4.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 32 | 0.9% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 29 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 39 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 81 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 34 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 23 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 11 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 33 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 33 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 27 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 56.8K | 30 | 0.1% |
| EasyList | blocklist | adguard | 68.8K | 51 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 38 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 751.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 3 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 11 | Unique: 24 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.5K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 853 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 43 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 68.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 66 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 363 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 114 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 343 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 273 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 17 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 285 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 251 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 37 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 25 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 36 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 41 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.8K | Targets: 25 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 14.8K | 40.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 2 | 4.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 231 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 8 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 8 | 1.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 159 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 90 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 2 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 132 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 252 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 1 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 12 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 22 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.2K | 5 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 12 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 12 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 224 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 73.2K | Targets: 1 | Unique: 73.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.8K | Targets: 35 | Unique: 8.7K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 281 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 51 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 249 | 8 | 3.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 5.3K | 2.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 16 | 1.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 16 | 1.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.2K | 17 | 1.5% |
| Greensnow | blocklist | ipv4 | 4.2K | 47 | 1.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 327 | 3 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 285 | 0.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 13 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 416 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 69 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 90 | 0.6% |
| BruteforceBlocker | blocklist | ipv4_find | 493 | 3 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 497 | 3 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 30 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 51 | 0.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 59 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.7K | 69 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 99 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 63 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 59 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
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
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.1K | 0.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 3.9K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 561 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 362 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 71 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 442 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 134 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 12 | Unique: 389 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 281 | 1.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.6K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.8K | 4 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 60.9K | 8 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.5K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 14 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 73.2K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 205 | 1 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 9.6K | 45 | 0.5% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 36.6K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 23 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.8K | 2 | 0.0% |

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
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 83 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 1 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 54 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 85 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 19 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 137 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 84 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 35 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 20 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 12 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 66 | Unique: 0 | Conflicts: 94</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.5% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3.0K | 69.9% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.2K | 33.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 20.7K | 24.9% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 189 | 11.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 40 | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 9 | 8.5% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 15.9K | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 996 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 1.8K | 6.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 3.1K | 5.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 1.4K | 4.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 25 | 3.8% |
| HaGeZi Pro | blocklist | domain | 201.5K | 6.8K | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 5.6K | 2.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 130 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 161 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 2.5K | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 349 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 977 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 22 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 39 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 22 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 42 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 51 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 15 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |

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
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 198 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 3 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 86 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 2 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 95 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 270 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 151 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 42 | 0.1% |
| HaGeZi Pro | blocklist | domain | 201.5K | 201 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 25 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 17 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 3 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| EasyList | blocklist | adguard | 68.8K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 229.1K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 234.2K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 56.8K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 185.6K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 10 | 0.0% |

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
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 50 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 52 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 74 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 45 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 201.5K | 41 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 62 | Unique: 0 | Conflicts: 45</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.8K | 42.7% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 395 | 23.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.0K | 10.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 83.0K | 3.5K | 4.2% |
| Adaway | blocklist | hostname | 6.5K | 266 | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 576 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 106 | 4 | 3.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.2K | 982 | 3.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 425 | 3.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 56.8K | 1.6K | 2.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 655 | 13 | 2.0% |
| ph00lt0_blocklist | blocklist | domain | 34.7K | 699 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Pro | blocklist | domain | 201.5K | 3.0K | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 5 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 234.2K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 184.8K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 16 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.3K | 55 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 30.5K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 240.6K | 30 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 601.2K | 11 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 97.9K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.2M | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 548.6K | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 229.1K | 62 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 232.5K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 39.2K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 157.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 228.4K | 9 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 19.2K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 30.9K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 215.2K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

