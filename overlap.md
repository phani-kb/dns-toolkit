# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-10-24 13:47:13 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 164 |
| Total Entries Analyzed | 9.1M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 142 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 34 |
| cidr_ipv4 | 3 |
| domain | 87 |
| ipv4 | 40 |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 68.0% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 11.3K | 60.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 33.3K | 56.5% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 311 | 46.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 166 | 45.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 142 | 37.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 589 | 35.6% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 110 | 31.9% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 34 | 30.9% |
| HaGeZi Pro | blocklist | domain | 208.8K | 57.7K | 27.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 44.4K | 23.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 6.4K | 22.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 40.5K | 17.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 13.2K | 16.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 6.3K | 13.5% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 4.4K | 12.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 15.3K | 7.2% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 147 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 243 | 1.3% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 3.3K | 1.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 26 | 0.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 458 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 145 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 42 | 0.3% |
| kadantiscam | blocklist | domain | 41.3K | 86 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 42 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 433 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 461 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 630 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| phishing_army | blocklist | domain | 154.9K | 93 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 227 | 0.0% |

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
| OISD Blocklist Small | blocklist | adguard | 59.0K | 30 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 2 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 23 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 36 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 50 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.7K | 14.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 6.5K | 8.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 267 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 6 | 5.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 433 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 9 | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 10 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 523 | 1.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 408 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 749 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1.7K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 84 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1.0K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 50 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 39 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |

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
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 1 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 35 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 5 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 59.0K | 494 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 666 | 0.3% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 499 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.5K | 44 | 0.1% |
| EasyList | blocklist | adguard | 71.0K | 96 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 127 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 35 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 6 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 669 | Targets: 30 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 494 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 116 | 0.4% |
| HaGeZi Pro | blocklist | domain | 208.8K | 564 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 55 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 647 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 311 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 499 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 45 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 35 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 127 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 32 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 243.8K | Targets: 13 | Unique: 243.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 433 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 10 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 4 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 3 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 245.1K | Targets: 23 | Unique: 132.6K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 95.2K | 50.0% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2.3K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 585 | 3.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 7.4K | 3.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 904 | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 824 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 205 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 83 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1.2K | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 351 | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 2 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 61 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 90 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 191.0K | Targets: 26 | Unique: 0 | Conflicts: 201</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 204 | 200 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.8% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 53.8K | 91.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 71.0K | 51.3K | 72.4% |
| Easy Privacy | blocklist | adguard | 55.5K | 28.7K | 51.6% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 666 | 50.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 58.7K | 24.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 21.6K | 10.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 2.4K | 6.8% |
| abpvn_hosts | blocklist | adguard | 1.0K | 23 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 118 | 0.6% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 291 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 12 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 190.3K | Targets: 67 | Unique: 0 | Conflicts: 41</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 669 | 647 | 96.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.5K | 93.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 53.8K | 91.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.1% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 95.2K | 38.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 79.4K | 38.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5.0K | 27.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 58.7K | 24.9% |
| YousList | blocklist | hostname | 625 | 151 | 24.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 44.4K | 21.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 6.1K | 21.6% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 6.1K | 13.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.7K | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 21.6K | 10.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 25 | 7.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2.4K | 6.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 5.0K | 6.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 20 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 554 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 118 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 277 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 346 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 194 | 0.2% |
| Torrent Trackers | blocklist | domain | 496 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 29 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 273 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 195 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 364 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 637 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 107 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 129 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 5 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 3 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 59.0K | 841 | 1.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 825 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 71 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.5K | 77 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 7 | 0.0% |

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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 9 | Unique: 4.0K | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 3 | 1.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 10 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 11 | 0.0% |
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
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 10 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 91.9% |
| Easy Privacy | blocklist | adguard | 55.5K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 2.1K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 25 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 181 | 0.1% |
| EasyList | blocklist | adguard | 71.0K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 3 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 55 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.0K | 56.5% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 625 | 198 | 31.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 205 | 12.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 12 | 10.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 7.6K | 9.4% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 55 | 8.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 4.7K | 7.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.7% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1.7K | 6.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 18 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.5K | 4.4% |
| HaGeZi Pro | blocklist | domain | 208.8K | 7.1K | 3.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 6.4K | 2.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 5.0K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 936 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 1.1K | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 53 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| Torrent Trackers | blocklist | domain | 496 | 1 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 19 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 28 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 33 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 43 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 142 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 23 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 24 | 8.7% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 16 | 3.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 14 | 2.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 315 | 2.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 134 | 2.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 671 | 1.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 273 | 1.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 400 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 34 | 1.7% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 166 | 1.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 795 | 1.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 211 | 1.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 198 | 1.0% |
| Greensnow | blocklist | ipv4 | 5.4K | 48 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 110 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 14 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 20 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 4 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.3K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 21.3K | 1.1K | 5.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 235 | 4.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 838 | 4.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 46 | 3.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 25 | 2.3% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 3 | 1.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 562 | 1.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 90 | 0.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 118 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 7 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 239 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 13 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 21 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 30 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 13 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 277 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 117 | 6.0% |
| Greensnow | blocklist | ipv4 | 5.4K | 165 | 3.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 64 | 2.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 24 | 2.1% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 244 | 1.1% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 5 | 1.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 5 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 236 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 261 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 48 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 3 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 21 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 8 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 12 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 9 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 7 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 38 | Unique: 10.1K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 45 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 41 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 76 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 22 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 190 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 189 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 138 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 126 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 53 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 264 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 34 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 154.9K | 100.2K | 64.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 275 | 62.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 70.9K | 33.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 72.7K | 30.8% |
| kadantiscam | blocklist | domain | 41.3K | 7.0K | 16.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 9.6K | 11.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1.3K | 4.7% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 3.1K | 3.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 3.7K | 1.8% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 580 | 1.8% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.8K | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 19.5K | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 354 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 196 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 16 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 43 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 465 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 107 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 31 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 4 | 1.5% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 65 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 6 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 38 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 16 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 55 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 57.7K | Targets: 34 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 261 | 94.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.5K | 77.7% |
| Greensnow | blocklist | ipv4 | 5.4K | 3.8K | 70.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 795 | 69.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 308 | 63.9% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 310 | 63.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.4K | 62.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 6.7K | 61.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 6.6K | 60.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 712 | 60.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.0K | 57.7% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 17.1K | 45.9% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 9.5K | 44.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 562 | 44.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 8.3K | 41.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 8.5K | 37.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 4.4K | 27.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 278 | 25.0% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 97 | 9.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 17 | 5.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 80 | 4 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 7 | 5.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 368 | 2.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 42 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 229 | 1.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 368 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 12 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 22 | Unique: 206 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 2 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 490 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 478 | 99.2% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 455 | 4.2% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 5 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 16 | 1.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 28 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 279 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 162 | 0.8% |
| Greensnow | blocklist | ipv4 | 5.4K | 26 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 310 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 98 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 46 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 25 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 10 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 32 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 11 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 98 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 67 | 67 | 100.0% |
| Torrent Trackers | blocklist | domain | 496 | 97 | 19.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 22 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 10.9K | 6.6K | 60.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 5.4K | 26.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.3K | 25.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 273 | 23.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 5.1K | 22.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 9.4K | 16.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 1.5K | 13.8% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 5.0K | 13.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 1.2K | 7.6% |
| Greensnow | blocklist | ipv4 | 5.4K | 123 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 9 | 1.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 8 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 32 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 289 | 1.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 9 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 65 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 1 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 3 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 27 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 27 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 34 | 0.0% |

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
| Easy Privacy | blocklist | adguard | 55.5K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 54 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 9 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 96.4K | Targets: 48 | Unique: 40.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 27 | 6.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 1.9K | 5.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 167 | 6 | 3.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 8 | 3.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 23.4K | 1.8% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.3K | 1.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 72 | 1.5% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 402 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 477 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 168 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.1K | 1.3% |
| phishing_army | blocklist | domain | 154.9K | 1.6K | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2.2K | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 2.2K | 0.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 120 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 265 | 0.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 304 | 0.4% |
| HaGeZi Pro | blocklist | domain | 208.8K | 914 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 173 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 28 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 45 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 458 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 38 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 613 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 358 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 194 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 37 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 81 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 11 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 13.1K | Targets: 53 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| YousList | blocklist | hostname | 625 | 108 | 17.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 13.0K | 16.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 46 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 426 | 12.0% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 20 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 361 | 1.3% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 8 | 1.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 2.5K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 582 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 280 | 0.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1.4K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 554 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 538 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 168 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 95 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 635 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 691 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 15 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 30 | 0.1% |
| phishing_army | blocklist | domain | 154.9K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 19 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 83 | 0.0% |

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
| EasyList | blocklist | adguard | 71.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 1 | 0.0% |

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
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 164 | 14.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 46 | 3.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 2 | 1.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 712 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 210 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 74 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 61 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 47 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.4K | 9 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 11 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 16 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 2 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.9K | Targets: 10 | Unique: 306 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 96.3% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 80 | 10.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 91 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1.0K | 31.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 889 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 8 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 78 | 5.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 80 | 4.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 38 | Unique: 11.5K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3.0K | 91.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 889 | 78.7% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 54 | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 135 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 10 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 539 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 4 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 15 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 9 | Unique: 13.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3.0K | 91.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 54 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 8 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 35 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 19 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 2 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 3 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 4.9K | 21.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 2.1K | 13.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 134 | 11.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 8.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 627 | 5.7% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 2.1K | 5.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.0K | 5.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 930 | 4.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 10 | 2.1% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 10 | 2.0% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 5 | 1.8% |
| Greensnow | blocklist | ipv4 | 5.4K | 59 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 13 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 17 | 0.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 8 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 79 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 31 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 13 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 33 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.5K | Targets: 20 | Unique: 13.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.6% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 28.7K | 15.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.8% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 2.7K | 4.5% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 44 | 3.3% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 6.2K | 2.6% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 341 | 1.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 329 | 0.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 13 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 71.0K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 59.0K | 34.2K | 58.1% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 51.3K | 26.9% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 34.7K | 14.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 17.0K | 8.0% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 96 | 7.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 659 | 1.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 64 | 0.3% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 131 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 482 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 490 | 478 | 97.6% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 447 | 4.1% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 5 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 28 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 151 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 277 | 0.7% |
| Greensnow | blocklist | ipv4 | 5.4K | 26 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 308 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 43 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 89 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 10 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 25 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 10 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.8% |

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
| tranco | allowlist | domain_top | 500 | 123 | 24.6% |
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
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 42 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 11 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 8 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.8% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 25 | 1.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 18 | Unique: 4.7K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 8 | 8.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 38 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 15 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 53 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 48 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 15 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 23 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 141 | Targets: 8 | Unique: 92 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 28 | 2.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 4 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 5 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 7 | Unique: 473 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 10 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 1 | 0.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.4K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.0K | Targets: 21 | Unique: 664 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 61 | 2.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 20 | 1.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 7 | 0.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 97 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 13 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 24 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 14 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 26 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 16 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 52 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 12 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 2 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.6K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.5K | 91.8% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.8% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.9K | Targets: 30 | Unique: 0 | Conflicts: 821</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 244 | 88.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 1.1K | 87.1% |
| Greensnow | blocklist | ipv4 | 5.4K | 4.6K | 86.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.3K | 69.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 11.5K | 57.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.1K | 40.6% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 162 | 33.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 151 | 31.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 6.8K | 29.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 2.2K | 20.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 211 | 18.4% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 6.2K | 16.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 9.5K | 16.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 8.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 821 | 7.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 821 | 7.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 74 | 6.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 61 | 5.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 107 | 4.2% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 399 | 3.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 4 | 2.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 24 | 2.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 3 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 93 | 0.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 5 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 21 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 91 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 10 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.9K | Targets: 33 | Unique: 0 | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4.9K | 95.0% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 455 | 92.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 447 | 92.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.6K | 43.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 6.8K | 42.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 400 | 34.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 3.5K | 17.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 6.0K | 16.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 1.8K | 16.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 8.5K | 14.7% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 21 | 7.6% |
| Greensnow | blocklist | ipv4 | 5.4K | 288 | 5.4% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 51 | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 26 | 2.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 30 | 2.4% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 399 | 1.9% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 1 | 1.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 130 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 97 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 423 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| DoH_IP_list | blocklist | ipv4 | 731 | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 15 | Unique: 2.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 72 | 21.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 61 | 5.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 6 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 3 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 3 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 2 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 335 | Targets: 10 | Unique: 221 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 3 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 6 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 4 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.6K | Targets: 19 | Unique: 5.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 311 | 2.0% |
| Adaway | blocklist | hostname | 6.5K | 84 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 2.3K | 0.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1.7K | 0.9% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1.6K | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 18 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 344 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 53 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 128 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 57 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 15 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 114 | 0.0% |

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
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 499 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 205 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1.5K | 0.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1.6K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 401 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 589 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 189 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 167 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 27 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 65 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

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
| AdGuard DNS filter | blocklist | adguard | 191.0K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 27 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 65 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 167 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 1 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 1 | 0.0% |

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
| HaGeZi Pro | blocklist | domain | 208.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 13 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.4K | Targets: 28 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 165 | 59.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 631 | 32.4% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 4.6K | 21.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 235 | 18.5% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 2.7K | 7.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.8K | 6.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 26 | 5.4% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 26 | 5.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 580 | 5.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 48 | 4.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 90 | 3.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 702 | 3.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 318 | 2.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 288 | 1.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 59 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 13 | 1.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 123 | 0.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 9 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 67 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 5 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 64 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 9 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 9 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 368 | Targets: 18 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 91 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 333 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 22 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 166 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 49 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 20 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 36 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 11 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 110 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 7 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 70 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 23 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 213.8K | Targets: 51 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 293 | 66.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 142.7K | 60.4% |
| phishing_army | blocklist | domain | 154.9K | 72.4K | 46.7% |
| HaGeZi Pro | blocklist | domain | 208.8K | 82.6K | 39.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 22.4K | 38.0% |
| kadantiscam | blocklist | domain | 41.3K | 12.8K | 31.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 70.9K | 29.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 11.9K | 25.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 127 | 19.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 5.1K | 18.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 12.6K | 15.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.6% |
| Spam404 | blocklist | domain | 8.1K | 1.1K | 13.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 21.6K | 11.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 167 | 10.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 375 | 7.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 15.3K | 7.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.6K | 7.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.1K | 6.2% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| YousList | blocklist | hostname | 625 | 29 | 4.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 538 | 4.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 294 | 3.9% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.2K | 3.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 985 | 3.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 82 | 2.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 6 | 2.3% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 2.2K | 2.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 138 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 136 | 0.9% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 205 | 0.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 9 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.9K | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 48 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 39 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.5K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.0K | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1.6K | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 104 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 123 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 216 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 90 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 213.8K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 236.1K | 142.7K | 60.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 883 | 49.9% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 72.3K | 47.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 22.4K | 38.0% |
| EasyList | blocklist | adguard | 71.0K | 17.0K | 24.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 87.2K | 11.5% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 21.6K | 11.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 167 | 10.0% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 127 | 9.6% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 181 | 6.6% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1.2K | 3.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 9 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.5K | 329 | 0.6% |
| abpvn_hosts | blocklist | adguard | 1.0K | 6 | 0.6% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 7 | 0.5% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.3K | Targets: 5 | Unique: 256 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.4% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.3K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.4% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| HaGeZi Pro | blocklist | domain | 208.8K | 41 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 9 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 608.5K | Targets: 42 | Unique: 590.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2.7K | 7.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 3.9K | 4.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 1.6K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1.4K | 0.6% |
| HaGeZi Pro | blocklist | domain | 208.8K | 891 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 358 | 0.4% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 126 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.1K | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 3 | 0.4% |
| kadantiscam | blocklist | domain | 41.3K | 120 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 33 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 465 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 34 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 45 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 227 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 83 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 766 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 129 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 19 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 42 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.0K | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| phishing_army | blocklist | domain | 154.9K | 161 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 8 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 30 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 7 | 0.0% |

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
| HaGeZi Pro | blocklist | domain | 208.8K | 331 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 40 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 24 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 66 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 20 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 57 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 208.8K | Targets: 68 | Unique: 0 | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 342 | 99.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 99.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 56.9K | 96.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 333 | 90.5% |
| hufilter | blocklist | hostname | 92 | 81 | 88.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 331 | 86.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 564 | 84.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 70 | 63.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 79 | 63.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 79.4K | 41.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 82.6K | 38.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.1K | 38.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 111 | 37.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 83.7K | 35.5% |
| YousList | blocklist | hostname | 625 | 202 | 32.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 8.4K | 29.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 57.7K | 28.4% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 25.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 11.7K | 25.1% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.4K | 16.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 11.8K | 14.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 4.4K | 12.7% |
| kadantiscam | blocklist | domain | 41.3K | 4.8K | 11.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1.6K | 11.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 49 | 11.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.4K | 10.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| phishing_army | blocklist | domain | 154.9K | 5.9K | 3.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 246 | 3.3% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 7.4K | 3.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 133 | 2.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 450 | 2.5% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 620 | 2.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| Spam404 | blocklist | domain | 8.1K | 149 | 1.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.7K | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 237 | 1.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 41 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 520 | 0.9% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 914 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 24 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 135 | 0.8% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 611 | 0.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 53 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.5K | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 1 | 0.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 4 | 0.4% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 101 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.4K | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 476 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.0K | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 891 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 41 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 3 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 345 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 342 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 16 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 87 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 10 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1.4K | 70.1% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 78 | 10.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 94 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 37.2K | Targets: 35 | Unique: 0 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 14.1K | 87.6% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 236 | 85.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.3K | 68.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 671 | 58.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 277 | 57.5% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 279 | 56.9% |
| Greensnow | blocklist | ipv4 | 5.4K | 2.7K | 50.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.1K | 40.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 4.1K | 37.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.0K | 33.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17.1K | 29.6% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 6.2K | 29.1% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 2.9K | 26.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 6.0K | 26.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 3.8K | 23.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 4.6K | 22.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 222 | 20.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 239 | 18.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 210 | 17.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 52 | 5.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 120 | 4.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 5 | 3.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 314 | 2.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 15 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 33 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 559 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 33 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 7 | 0.0% |

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
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 111 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 20 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 53 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 13 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 17 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 108 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 21 | 0.0% |

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
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 84 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 12 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 70 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 82 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 81 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 11 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 56 | Unique: 422.8K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 995 | 13.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1.4K | 5.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 8 | 1.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2.9K | 1.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 2.5K | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 2.9K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| phishing_army | blocklist | domain | 154.9K | 1.4K | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 613 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 98 | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.4K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.2K | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 1 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 2.1K | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 181 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 120 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 187 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 54 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| kadantiscam | blocklist | domain | 41.3K | 41 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 210 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 273 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 22 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 20 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 8 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 41.3K | Targets: 40 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 28.6K | 35.3% |
| phishing_army | blocklist | domain | 154.9K | 17.9K | 11.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 12.8K | 6.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 12.1K | 5.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7.0K | 2.9% |
| HaGeZi Pro | blocklist | domain | 208.8K | 4.8K | 2.3% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 496 | 1.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 2 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 30 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 58 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 68 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 41 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 120 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 86 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 24 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 21 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 29 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 29 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 175 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 126 | 2.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 3 | 1.2% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 47 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 17 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 24 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 10 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 29 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 5 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 236.1K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 2 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 24 | Targets: 5 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### local_domain_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7 | Targets: 21 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 9 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

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
| Firehol_level3 | blocklist | ipv4 | 10.9K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 477 | 1.4% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 1.2K | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 883 | 0.4% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 1.8K | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.5K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 219.8K | Targets: 37 | Unique: 207.7K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.7% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5.4K | 34.1% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 9 | 3.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 423 | 1.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 20 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 281 | 1.7% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 559 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 16 | 1.4% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 6 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 29 | 1.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 6 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 13 | 1.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 10 | 0.9% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 10 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 7 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 33 | 0.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 11 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 59 | 0.5% |
| Greensnow | blocklist | ipv4 | 5.4K | 29 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 56 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 10 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 91 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 62 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 41 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 34 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 8 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 48 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 8 | 0.1% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3M | Targets: 60 | Unique: 152.8K | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.1M | 100.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 13.7K | 43.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 23.4K | 24.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 167 | 29 | 17.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 51 | 11.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 19.5K | 8.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 3.7K | 8.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 11 | 6.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2.0K | 5.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 691 | 5.3% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 539 | 3.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 6.5K | 3.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 6.4K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 4.4K | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1.3K | 1.6% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 8 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 288 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 404 | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.2K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 26 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 104 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 637 | 0.3% |
| phishing_army | blocklist | domain | 154.9K | 485 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1.0K | 0.2% |
| kadantiscam | blocklist | domain | 41.3K | 25 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 16 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 37 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 124 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1M | Targets: 49 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.1M | 80.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 16.3K | 16.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 167 | 20 | 12.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 41 | 9.3% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2.3K | 7.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 840 | 2.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 4.0K | 1.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 4.2K | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1.0K | 1.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 2.0K | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.4K | 0.9% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 98 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 1 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 630 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 106 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| phishing_army | blocklist | domain | 154.9K | 363 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 364 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 766 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 65 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 18 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 19 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 9 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 9 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 27.7K | Targets: 31 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 154.9K | 23.2K | 15.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 20.2K | 8.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 17 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 5.1K | 2.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 6 | 2.3% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 481 | 1.4% |
| kadantiscam | blocklist | domain | 41.3K | 496 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 21 | 0.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.3K | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 402 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 15 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 620 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 34 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 70 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 11 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 98 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 104 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 2 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 236.1K | Targets: 63 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1.6K | 99.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 58.3K | 98.8% |
| hufilter | blocklist | hostname | 92 | 82 | 89.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 367 | 83.0% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 499 | 74.6% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 20.2K | 73.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 142.7K | 66.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.6% |
| phishing_army | blocklist | domain | 154.9K | 82.7K | 53.4% |
| YousList | blocklist | hostname | 625 | 269 | 43.0% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 208.8K | 83.7K | 40.1% |
| WaLLy3K | blocklist | domain | 351 | 136 | 38.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.4K | 34.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 58.7K | 30.8% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 72.7K | 30.3% |
| kadantiscam | blocklist | domain | 41.3K | 12.1K | 29.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 87 | 25.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 20.3K | 25.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 10.6K | 22.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.8K | 20.9% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 23 | 20.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 5.9K | 20.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 40.5K | 19.9% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 57 | 14.9% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 4.0K | 11.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 556 | 11.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 36 | 9.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 8 | 3.1% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 2.2K | 2.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 189 | 1.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 8 | 1.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 85 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 114 | 0.8% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 212 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.9K | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 122 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 463 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 395 | 0.6% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 1.2K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.4K | 0.5% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.2K | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 9 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1.4K | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 305 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 4 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 236.1K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 59.0K | 58.3K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.3% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1.2K | 67.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 142.7K | 66.7% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 825 | 61.8% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 81.6K | 53.0% |
| EasyList | blocklist | adguard | 71.0K | 34.7K | 48.9% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 499 | 37.7% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 58.7K | 30.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 66 | 17.9% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 92.6K | 12.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 4.0K | 11.5% |
| Easy Privacy | blocklist | adguard | 55.5K | 6.2K | 11.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 36 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 54 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 122 | 0.6% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 35 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 18.8K | Targets: 10 | Unique: 18.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 28 | 2.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 60 | 0.2% |
| EasyList | blocklist | adguard | 71.0K | 64 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 76 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 118 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 122 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 14 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 2 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.8K | Targets: 38 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.9K | 9.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.4K | 8.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 12.7K | 5.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 28 | 2.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 25 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 47 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 60 | 0.2% |
| Torrent Trackers | blocklist | domain | 496 | 1 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 122 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 76 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 41 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 19 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 118 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 237 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 243 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 6 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 45 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 24 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 16 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 22 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 22 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 59.0K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 841 | 62.9% |
| EasyList | blocklist | adguard | 71.0K | 34.2K | 48.3% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 494 | 37.4% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 53.8K | 28.1% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 58.3K | 24.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 22.4K | 10.5% |
| Easy Privacy | blocklist | adguard | 55.5K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1.5K | 4.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 65 | 3.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 30 | 2.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 76 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 151 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 2 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 59.0K | Targets: 62 | Unique: 0 | Conflicts: 17</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 84 | 91.3% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 494 | 73.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.6% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 53.8K | 28.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 56.9K | 27.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4.7K | 25.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 58.3K | 24.7% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33.3K | 16.4% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5.7K | 12.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 3.3K | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 749 | 11.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 22.4K | 10.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 25 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 22 | 6.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 918 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 3.6K | 4.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 582 | 4.4% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.5K | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 65 | 3.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 76 | 0.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 216 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 57 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 76 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 904 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 262 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 173 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 24 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 10 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 83 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 43 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 181 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 75 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 404 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 106 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 259 | Targets: 16 | Unique: 134 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 6 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 68 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 1 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 67 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 98 | 67 | 68.4% |
| Torrent Trackers | blocklist | domain | 496 | 66 | 13.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 34.6K | Targets: 24 | Unique: 19.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 477 | 27.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 35 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 1.5K | 2.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 4.0K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 2.4K | 1.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| EasyList | blocklist | adguard | 71.0K | 659 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.5K | 341 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 1.2K | 0.6% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 3.8K | 0.5% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 613 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 60 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 54 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 34.6K | Targets: 75 | Unique: 0 | Conflicts: 166</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 87 | 19.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 698 | 19.6% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 47 | 9.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.5K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 523 | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1.9K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 35 | 5.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 11 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1.5K | 2.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 280 | 2.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 4.4K | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1.7K | 2.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 560 | 1.8% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 481 | 1.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 4.0K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 27 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 6 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 65 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2.4K | 1.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 54 | 0.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 1.2K | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 477 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 192 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 2.7K | 0.4% |
| phishing_army | blocklist | domain | 154.9K | 606 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 54 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 60 | 0.3% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 22 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.0K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 354 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 58 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 840 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 7 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 120 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 61 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 77 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 154.9K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 23.2K | 84.0% |
| kadantiscam | blocklist | domain | 41.3K | 17.9K | 43.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 100.2K | 41.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 82.7K | 35.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 72.4K | 33.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 68 | 26.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 13.6K | 16.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 24 | 4.8% |
| HaGeZi Pro | blocklist | domain | 208.8K | 5.9K | 2.8% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 606 | 1.7% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1.6K | 1.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 167 | 2 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 38 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 28 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 47 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 93 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 12 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 485 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 12 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 161 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 363 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 7 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 20 | Unique: 61.6K | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 94 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 91 | 4.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 4 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 10 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 3 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 48 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 4 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 12 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 7 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 290 | 17.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 76 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 296 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 145 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 46 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 299 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 298 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 47 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 25 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 59 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 79 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 35 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 47 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 29 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 22 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 20 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 54 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 576 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 34 | 8.9% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 8 | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 433 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.3K | 3.7% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 904 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 311 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 6 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 918 | 1.6% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 11 | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1.2K | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 991 | 1.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 2.4K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2.1K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 98 | 1 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1.4K | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 180 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Torrent Trackers | blocklist | domain | 496 | 2 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 5 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 585 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 136 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 6 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 7 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 22 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 25 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 71.0K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 761.6K | Targets: 17 | Unique: 442.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 1.8K | 99.5% |
| RPiList_specials-phishing | blocklist | adguard | 153.9K | 133.6K | 86.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 87.2K | 40.8% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 92.6K | 39.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 3.8K | 10.9% |
| AdGuard Base filter | blocklist | adguard | 1.3K | 6 | 0.5% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 151 | 0.3% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 291 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| EasyList | blocklist | adguard | 71.0K | 131 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 14 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 19 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 8 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 153.9K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 236.1K | 81.6K | 34.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 72.3K | 33.8% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 133.6K | 17.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 613 | 1.8% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.8K | 5 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 12 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.9K | Targets: 26 | Unique: 0 | Conflicts: 26</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 117 | 42.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 631 | 11.7% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 1.3K | 6.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 28 | 5.8% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 28 | 5.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 468 | 4.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 1.3K | 3.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 34 | 3.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1.5K | 2.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 25 | 1.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 210 | 1.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 135 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 1 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 41 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 17 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 51 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 32 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 26 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 26 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 16 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 18 | Unique: 204 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 28 | 19.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 164 | 13.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 25 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 20 | 1.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 3 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 222 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 278 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 74 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 39 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 15 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.4K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 4 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 20 | Unique: 2.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 277 | 64 | 23.1% |
| Greensnow | blocklist | ipv4 | 5.4K | 90 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 25 | 1.3% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 107 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 2 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 2 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 120 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 8 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 7 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 4 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.5K | Targets: 21 | Unique: 0 | Conflicts: 13.0K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 821 | 3.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 264 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 33 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 26 | 1.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 82 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 5 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 32 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 27 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 9 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 24 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 33 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.9K | Targets: 28 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 315 | 27.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 468 | 24.0% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 48 | 17.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 627 | 12.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 6.7K | 11.5% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 4.1K | 11.0% |
| Greensnow | blocklist | ipv4 | 5.4K | 580 | 10.8% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 2.2K | 10.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.5K | 10.0% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 46 | 9.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 43 | 8.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 1.7K | 8.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 1.8K | 7.7% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 840 | 7.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 90 | 7.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 973 | 6.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 249 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 14 | 1.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 91 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 24 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 24 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 59 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 28 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 10 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 18.8K | 28 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 10 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 718 | Targets: 39 | Unique: 373 | Conflicts: 272</summary>

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
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 2 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 1 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 27 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 12 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 8 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 31 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.4K | Targets: 54 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 499 | 30.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 973 | 27.3% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 116 | 17.3% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.7K | 9.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 24 | 6.3% |
| Adaway | blocklist | hostname | 6.5K | 408 | 6.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 22 | 6.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 904 | 5.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3.3K | 5.6% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.9K | 5.6% |
| HaGeZi Pro | blocklist | domain | 208.8K | 8.4K | 4.0% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.4K | 3.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 6.1K | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 361 | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 5.9K | 2.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 2.0K | 2.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 985 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 47 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 125 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 165 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 15 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 205 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 17 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 288 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 54 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 34 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 234.4K | Targets: 33 | Unique: 178.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 12.7K | 67.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.6K | 28.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 461 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 77 | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 476 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 106 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 216 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 75 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 305 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 21 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 10 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 81 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 195 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 210 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 2 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 30 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 124 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 65 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 29 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.7K | Targets: 43 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 5.7K | 9.7% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 45 | 6.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 11.9K | 5.6% |
| HaGeZi Pro | blocklist | domain | 208.8K | 11.7K | 5.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 936 | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 10.6K | 4.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 6.1K | 3.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.3K | 3.1% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 55 | 1.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 180 | 1.2% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 95 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 192 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 313 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 265 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3.7K | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 41.3K | 68 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 196 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 22 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 11 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 12 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 45 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.4K | Targets: 29 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 899 | 1.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 995 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 147 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 246 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 294 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 25 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 10 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 85 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 24 | Unique: 1.4K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 2 | 0.8% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 54 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 15 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 49 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 28 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 5.9K | 31.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 19.4K | 8.3% |
| pexcn Torrent Trackers | blocklist | domain_url | 67 | 2 | 3.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 98 | 2 | 2.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Torrent Trackers | blocklist | domain | 496 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 615 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 125 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 216 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 12 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 78 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 395 | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 520 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 277 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 44 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 22 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 17 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 104 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 37 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 8 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 2.7K | 3.3% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1.2K | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 24 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 10 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |

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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 15 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 38 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 32 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 35 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 61 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 29 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 17 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 30 | Unique: 4.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 1.6K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 1.1K | 0.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 20 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 41.3K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 52 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 149 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 110 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 18 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 5 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 17.6K | 7.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 4.3K | 1.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 2.6K | 1.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.0K | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 2 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 208.8K | 450 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 120 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 22 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 2 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 2 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 80.8K | Targets: 67 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 98.6% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 363 | 82.1% |
| kadantiscam | blocklist | domain | 41.3K | 28.6K | 69.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.6K | 41.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 401 | 24.3% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 66 | 17.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 49 | 13.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 899 | 12.1% |
| phishing_army | blocklist | domain | 154.9K | 13.6K | 8.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 20.3K | 8.6% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 2.0K | 6.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3.6K | 6.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 12.6K | 5.9% |
| HaGeZi Pro | blocklist | domain | 208.8K | 11.8K | 5.7% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 32 | 4.8% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.7K | 4.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9.6K | 4.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 5.0K | 2.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 344 | 2.3% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 313 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 3.9K | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 41 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 304 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 70 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 824 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 41 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.3K | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.0K | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 37 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 106 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 187 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 46 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 6.4K | 34.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 21.6K | 9.2% |
| pexcn Torrent Trackers | blocklist | domain_url | 67 | 2 | 3.0% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| CF_Torrent_Trackers | blocklist | domain_url | 98 | 2 | 2.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 496 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 703 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 157 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 165 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 262 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 95 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 208.8K | 611 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 346 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 463 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 13 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 123 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 50 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 43 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 19 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 37 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 13 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |

</details>

---

### StevenBlack_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 34 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 15 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 17 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 38 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 29 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 35 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 61 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 32 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 31.5K | Targets: 29 | Unique: 11.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 442 | 34 | 7.7% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1.9K | 1.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 560 | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13.7K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 2 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 580 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 8 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.3K | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 205 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 212 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 145 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 34 | 0.1% |
| phishing_army | blocklist | domain | 154.9K | 47 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 126 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 101 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 7 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 167 | Targets: 4 | Unique: 110 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 29 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 20 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 2 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20.0K | Targets: 28 | Unique: 0 | Conflicts: 82</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 838 | 66.1% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 11.5K | 54.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.4K | 35.8% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 2.9K | 26.3% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 98 | 20.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 89 | 18.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 930 | 18.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 198 | 17.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 1.7K | 16.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 3.5K | 15.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 8.3K | 14.3% |
| Greensnow | blocklist | ipv4 | 5.4K | 702 | 13.1% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 4.6K | 12.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 210 | 10.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 1.4K | 9.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 74 | 6.3% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 12 | 4.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 39 | 3.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 4 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 12 | 1.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 82 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 82 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 10 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 52 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 21 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 62 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 496 | Targets: 8 | Unique: 310 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 98 | 97 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 67 | 66 | 98.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 45 | Unique: 0 | Conflicts: 528</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| local_ai_allowlist | allowlist | domain | 24 | 3 | 12.5% |
| local_ai_blocklist | blocklist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 123 | 7.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 2 | 4.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 3 | 3.1% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 259 | 1 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 19 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 35 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 37 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 82 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 19 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 1 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 11 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 32 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 27 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 14 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 31 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 71.0K | 51 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 30 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.5K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 38 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 761.6K | 1 | 0.0% |

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
| Easy Privacy | allowlist | adguard | 852 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 43 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 66 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 53 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 7 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 442 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 363 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 87 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 367 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 275 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 34 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 293 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 27 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 51 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 49 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 41 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.1K | Targets: 25 | Unique: 861 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 14.1K | 37.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 2 | 4.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 249 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 11 | 2.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 10 | 2.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 130 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 94 | 0.6% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 1 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 229 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 4 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 13 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 3 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 9 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.4K | 3 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 20 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 21 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 281 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 16 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 21 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 88.3K | Targets: 1 | Unique: 88.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.8K | Targets: 35 | Unique: 8.6K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 281 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 47 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 277 | 8 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5.4K | 2.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 16 | 1.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.1K | 15 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| Greensnow | blocklist | ipv4 | 5.4K | 64 | 1.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.3K | 14 | 1.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 482 | 5 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 490 | 5 | 1.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.9K | 91 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 314 | 0.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 16 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 73 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 141 | 1 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 31 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 94 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 93 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 65 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 97 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 52 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 36 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.1K | 0.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 3.9K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 556 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 375 | 0.2% |
| HaGeZi Pro | blocklist | domain | 208.8K | 133 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 72 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 8 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 14 | Unique: 347 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 281 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 9 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.0K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 21.3K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.9K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 10.9K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.8K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 15 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 10.9K | 45 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 37.2K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 23 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 88.3K | 1 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 351 | Targets: 33 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| YousList | blocklist | hostname | 625 | 9 | 1.4% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 83 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 54 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 85 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 136 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 19 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 83 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 20 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 35 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 64 | Unique: 0 | Conflicts: 90</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.2K | 33.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 20.7K | 25.6% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 189 | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 40 | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 15.8K | 6.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 991 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 1.8K | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 3.1K | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 1.4K | 4.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 25 | 3.7% |
| HaGeZi Pro | blocklist | domain | 208.8K | 6.8K | 3.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 5.6K | 2.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 128 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 161 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.5K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 351 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 24 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 977 | 0.1% |
| kadantiscam | blocklist | domain | 41.3K | 15 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 51 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 41 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 4 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 42 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |

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
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 3 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 2 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 241 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 86 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 94 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 42 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 269 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 151 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 208.8K | 202 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 17 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 29 | 0.0% |
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
| Easy Privacy | blocklist | adguard | 55.5K | 10 | 0.0% |
| EasyList | blocklist | adguard | 71.0K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 213.8K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 236.1K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 59.0K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 34.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.0K | 39 | 0.0% |

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
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 50 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 52 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 74 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 45 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 208.8K | 41 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 8 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 60 | Unique: 0 | Conflicts: 43</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.7K | 397 | 24.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.0K | 10.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 80.8K | 3.5K | 4.3% |
| Adaway | blocklist | hostname | 6.5K | 267 | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 576 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 4 | 3.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.4K | 973 | 3.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 426 | 3.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 59.0K | 1.6K | 2.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 383 | 10 | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 34.6K | 698 | 2.0% |
| AdGuard Base filter | blocklist | domain_adguard | 669 | 13 | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Pro | blocklist | domain | 208.8K | 3.0K | 1.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 5 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 368 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 236.1K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 190.3K | 1.6K | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.7K | 55 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.6K | 18 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 608.5K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 234.4K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 31.5K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| phishing_army | blocklist | domain | 154.9K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.8K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| kadantiscam | blocklist | domain | 41.3K | 9 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 245.1K | 30 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 27.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 96.4K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.4K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 243.8K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 213.8K | 82 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.8K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

