# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-11-07 01:37:26 UTC

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
| Total Entries Analyzed | 9.6M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 140 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| cidr_ipv4 | 3 |
| domain | 87 |
| ipv4 | 40 |
| adguard | 32 |

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
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 32.8K | 54.0% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 316 | 45.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 166 | 45.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 142 | 37.0% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 590 | 35.8% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 110 | 31.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 34 | 30.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| HaGeZi Pro | blocklist | domain | 222.6K | 53.5K | 24.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 6.5K | 22.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 39.5K | 20.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 39.9K | 17.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 13.1K | 15.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 6.3K | 13.5% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 4.4K | 12.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 15.3K | 7.9% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 147 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 236 | 1.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 3.3K | 1.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 26 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 459 | 0.5% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 143 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 42 | 0.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 1 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 38 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| kadantiscam | blocklist | domain | 44.5K | 85 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 462 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| phishing_army | blocklist | domain | 146.3K | 83 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 739 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 432 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 632 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 229 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 910 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| EasyList | blocklist | adguard | 63.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 30 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 23 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 37 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 6.5K | 7.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 267 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 6 | 5.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 433 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 9 | 2.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 523 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 409 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 10 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 750 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.6K | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 85 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1.0K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 50 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 41 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 93 | Targets: 8 | Unique: 32 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 213 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 3 | 0.2% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 5 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 35 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 5 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.8K | 509 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 686 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 514 | 0.2% |
| EasyList | blocklist | adguard | 63.7K | 97 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.6K | 44 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 137 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 35 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 2 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 691 | Targets: 31 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 509 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 122 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 667 | 0.4% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.6K | 573 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 55 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 316 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 514 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 137 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 35 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 47 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 32 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 6 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 308.6K | Targets: 14 | Unique: 308.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| 1Hosts (Lite) | blocklist | domain | 203.0K | 432 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 4 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 3 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 306.6K | Targets: 23 | Unique: 190.7K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 100.4K | 53.1% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 2.3K | 15.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 586 | 3.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 958 | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| HaGeZi Pro | blocklist | domain | 222.6K | 3.4K | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2.1K | 1.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 820 | 0.9% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 204 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 83 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1.2K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 2 | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 353 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 61 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 189.9K | Targets: 26 | Unique: 0 | Conflicts: 202</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 205 | 201 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 93.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 55.5K | 91.4% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 63.7K | 44.1K | 69.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 28.6K | 51.5% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 686 | 50.8% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 60.5K | 26.0% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 25.7K | 13.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 2.4K | 6.7% |
| abpvn_hosts | blocklist | adguard | 1.0K | 23 | 2.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 119 | 0.6% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 178 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 8 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 189.2K | Targets: 69 | Unique: 0 | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 691 | 667 | 96.5% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 93.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 55.5K | 91.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 222.6K | 74.9K | 33.6% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 100.4K | 32.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5.0K | 27.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 60.5K | 26.0% |
| YousList | blocklist | hostname | 625 | 150 | 24.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 6.1K | 21.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39.5K | 19.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 25.7K | 13.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.7K | 11.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 5.4K | 11.5% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 25 | 7.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2.4K | 6.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 5.0K | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 20 | 5.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 553 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 119 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 275 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 344 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Torrent Trackers | blocklist | domain | 502 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 188 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 29 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 194 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 270 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 10 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 362 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 105 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 127 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 271 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 628 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.8K | 844 | 1.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 828 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 72 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 77 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 7 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 205 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 213 | 201 | 94.4% |

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
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1 | 0.0% |

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
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |

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
| Easy Privacy | blocklist | adguard | 55.6K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 2.1K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 60 | 0.1% |
| EasyList | blocklist | adguard | 63.7K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 18 | 0.0% |

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
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 205 | 12.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 12 | 10.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 7.6K | 8.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 55 | 8.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 4.7K | 7.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.7% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1.8K | 6.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 21 | 6.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 18 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.5K | 4.4% |
| HaGeZi Pro | blocklist | domain | 222.6K | 7.0K | 3.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 6.4K | 2.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 5.0K | 2.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 937 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1.0K | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 54 | 0.4% |
| Torrent Trackers | blocklist | domain | 502 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 18 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 14 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 142 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 10 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 43 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 993 | Targets: 22 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 300 | 23 | 7.7% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 15 | 2.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 13 | 2.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 124 | 2.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 217 | 2.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 340 | 2.3% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 688 | 1.9% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 205 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 32 | 1.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 310 | 1.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 590 | 1.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 175 | 1.0% |
| Greensnow | blocklist | ipv4 | 3.6K | 33 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 130 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 79 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 2 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 6 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 3 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 852 | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 18.1K | 691 | 3.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 43 | 3.6% |
| Greensnow | blocklist | ipv4 | 3.6K | 87 | 2.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 25 | 2.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 235 | 1.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 10 | 1.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 55 | 0.6% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 2 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 76 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 180 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 122 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 12 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 2 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 28 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 25 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 5 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 300 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 117 | 5.7% |
| Greensnow | blocklist | ipv4 | 3.6K | 143 | 4.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 66 | 2.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 23 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 241 | 1.3% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 224 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 277 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 46 | 0.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 10 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 43 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 9 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 17 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 9 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |

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
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 1 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 190 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 22 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 138 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 45 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 191 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 75 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 40 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 264 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 17 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 53 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 126 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 49 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 146.3K | 93.3K | 63.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 252 | 55.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 64.8K | 33.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 67.8K | 29.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 5.1K | 5.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| kadantiscam | blocklist | domain | 44.5K | 2.1K | 4.8% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 734 | 3.5% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 3.2K | 3.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 586 | 2.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 3.8K | 1.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 19.5K | 1.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.8K | 1.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 6.3K | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 340 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 196 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 16 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 43 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 105 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 25 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 465 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 29 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 12 | 0.0% |

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
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 2 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 67 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 23 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 15 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 15 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 55 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 11 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 57.7K | Targets: 34 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 300 | 277 | 92.3% |
| Greensnow | blocklist | ipv4 | 3.6K | 2.2K | 62.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 712 | 60.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.1K | 60.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 590 | 59.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7.9K | 52.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 266 | 52.1% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 267 | 51.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4.6K | 50.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 986 | 47.9% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 5.1K | 45.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 13.1K | 36.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 8.0K | 34.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 335 | 32.2% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 5.7K | 31.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 5.5K | 30.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 5.2K | 29.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 180 | 21.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 104 | 10.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 14 | 10.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 17 | 5.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 80 | 4 | 5.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 395 | 2.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 42 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 249 | 1.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 368 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 12 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 22 | Unique: 207 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 515 | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 486 | 95.1% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 489 | 4.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 15 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 112 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 267 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 11 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 195 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.6K | 13 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 36 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 45 | 0.3% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 22 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 25 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 18 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 4 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 105 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 68 | 100.0% |
| Torrent Trackers | blocklist | domain | 502 | 104 | 20.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 23 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.3K | 8.4K | 74.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 340 | 34.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 4.8K | 20.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 3.6K | 20.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 995 | 19.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 5.4K | 15.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.3K | 14.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7.9K | 13.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 1.6K | 9.1% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 18 | 3.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 15 | 2.9% |
| Greensnow | blocklist | ipv4 | 3.6K | 87 | 2.4% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 360 | 2.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 36 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 2 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 58 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 3 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 3 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 24 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 24 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 13 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 33 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 7 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 54 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 98.8K | Targets: 48 | Unique: 12.8K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 1.8K | 6.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 30.8K | 5.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 25 | 5.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 5 | 1.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 23.5K | 1.8% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.4K | 1.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 73 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 470 | 1.3% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 270 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.2K | 1.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 168 | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2.0K | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.0K | 0.9% |
| phishing_army | blocklist | domain | 146.3K | 1.3K | 0.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 128 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 267 | 0.6% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.1K | 0.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 173 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 267 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 35 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 50 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 28 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 459 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 188 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 617 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 442 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 42 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 18 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 84 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 38 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 21 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 11 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 12.9K | 14.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 46 | 12.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 422 | 11.9% |
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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 363 | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 8 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.5K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 582 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 279 | 0.8% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.3K | 0.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 546 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 553 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 168 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 95 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| kadantiscam | blocklist | domain | 44.5K | 30 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 691 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 635 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 15 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 20 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 9 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 83 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 36 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 14.0K | Targets: 7 | Unique: 14.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.4K | 11 | 0.8% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 445 | 2 | 0.4% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 1 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 7 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 216 | 20.7% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 43 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 4 | 2.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 712 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 280 | 0.8% |
| Greensnow | blocklist | ipv4 | 3.6K | 26 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 85 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 4 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 2 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 2 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 62 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 47 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 6 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 16 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 7 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.1K | Targets: 8 | Unique: 308 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1.5K | 97.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 80 | 10.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 91 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 2 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1.1K | 29.9% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 886 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 7 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 78 | 5.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 80 | 3.9% |
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
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3.0K | 83.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 55 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 33 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 14 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 3 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 3 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 39 | Unique: 11.5K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3.0K | 83.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 886 | 74.6% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 55 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 131 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 539 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 63 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 15 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 33 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 13 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 5.1K | 28.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 5.1K | 22.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 124 | 12.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 2.7K | 7.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 643 | 7.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 995 | 6.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.1K | 5.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 855 | 4.8% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 10 | 3.3% |
| Greensnow | blocklist | ipv4 | 3.6K | 66 | 1.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 30 | 1.5% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 12 | 1.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 7 | 0.7% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 3 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 2 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 35 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 27 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 18 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.6K | Targets: 20 | Unique: 13.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.8% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 28.6K | 15.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.8% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 2.7K | 4.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 44 | 3.3% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 6.3K | 2.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 344 | 1.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 639 | 0.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 3 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 12 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 63.7K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.8K | 35.4K | 58.3% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 44.1K | 23.2% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 35.9K | 15.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 18.2K | 9.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 97 | 7.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 655 | 1.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 65 | 0.4% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 43 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 12 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 511 | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 515 | 486 | 94.4% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 468 | 4.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 13 | 1.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 266 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 98 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 11 | 0.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 187 | 0.5% |
| Greensnow | blocklist | ipv4 | 3.6K | 12 | 0.3% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 1 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 29 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 39 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 15 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 20 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 23 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 5 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.8K | 99.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 34.2% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 5 | 0.0% |

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
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 6 | 0.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 42 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 7 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 25 | 1.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
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
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 23 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 15 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 52 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 46 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 14 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 22 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 138 | Targets: 13 | Unique: 80 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 12 | 1.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 1 | 0.1% |
| Greensnow | blocklist | ipv4 | 3.6K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 14 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 9 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 9 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 8 | Unique: 471 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 3 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 3 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 4 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 985 | Targets: 22 | Unique: 602 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 57 | 2.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 14 | 1.3% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 10 | 1.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| Greensnow | blocklist | ipv4 | 3.6K | 9 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 104 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 57 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 18 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 20 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 26 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 20 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 5 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 7 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 8 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.7K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.3% |
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.1% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 18.1K | Targets: 30 | Unique: 0 | Conflicts: 606</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 691 | 81.1% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 241 | 80.3% |
| Greensnow | blocklist | ipv4 | 3.6K | 2.8K | 77.4% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.1K | 53.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 7.4K | 41.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 8.8K | 37.9% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 112 | 21.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 98 | 19.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.8K | 19.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 175 | 17.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 5.5K | 15.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.6K | 10.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.7K | 9.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 92 | 8.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 85 | 7.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 9 | 6.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 606 | 5.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 606 | 5.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 104 | 4.1% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 445 | 3.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 26 | 2.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 5 | 1.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 31 | 1.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 82 | 0.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 30 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 86 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 11.3K | Targets: 28 | Unique: 0 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 489 | 95.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 468 | 91.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8.4K | 55.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 8.8K | 48.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 310 | 31.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.7K | 18.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 6.2K | 17.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 8.0K | 13.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 2.2K | 12.6% |
| Greensnow | blocklist | ipv4 | 3.6K | 283 | 7.9% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 17 | 5.7% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 28 | 3.3% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 445 | 2.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 44 | 2.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 20 | 2.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 117 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 103 | 0.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 57 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 57 | 0.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 201 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 12 | Unique: 2.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 72 | 21.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 57 | 5.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 11 | 1.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 31 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 24 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 21 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 17 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 335 | Targets: 10 | Unique: 218 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 2 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 4 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.7K | Targets: 19 | Unique: 6.0K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 320 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 85 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1.7K | 0.9% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2.3K | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.0K | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 346 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 54 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 128 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 22 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 55 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 112 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 14 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.4% |
| Easy Privacy | blocklist | adguard | 55.6K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 65 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 26 | 0.1% |
| EasyList | blocklist | adguard | 63.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 24 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 396 | 11.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 490 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 205 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1.5K | 0.8% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.6K | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 401 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 590 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 187 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 65 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 26 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 24 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 4 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |

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
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 13 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.6K | Targets: 23 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 300 | 143 | 47.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 377 | 18.3% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 2.8K | 15.3% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 87 | 10.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 2.2K | 6.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 476 | 5.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 2.2K | 3.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 641 | 3.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 33 | 3.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 79 | 3.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 545 | 3.0% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 13 | 2.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 12 | 2.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 26 | 2.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 18 | 1.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 66 | 1.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 283 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 9 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 87 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 53 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 59 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 27 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 369 | Targets: 18 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 91 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 22 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 49 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 333 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 166 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 20 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 36 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 11 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 110 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 23 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 70 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 21 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 194.8K | Targets: 53 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 375 | 81.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 140.5K | 60.4% |
| phishing_army | blocklist | domain | 146.3K | 69.7K | 47.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 22.6K | 37.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 67.9K | 30.5% |
| kadantiscam | blocklist | domain | 44.5K | 13.5K | 30.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 64.8K | 27.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 10.1K | 21.5% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 4.2K | 19.9% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 137 | 19.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 14.1K | 16.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 25.7K | 13.6% |
| Spam404 | blocklist | domain | 8.1K | 1.1K | 13.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 15.3K | 7.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.7K | 7.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 372 | 7.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.0K | 5.5% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| YousList | blocklist | hostname | 625 | 26 | 4.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 546 | 4.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 292 | 3.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 9 | 3.4% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.2K | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 879 | 3.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 82 | 2.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 2.0K | 2.0% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 24 | 1.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 138 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 136 | 0.9% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 194 | 0.7% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2.1K | 0.7% |
| Adaway | blocklist | hostname | 6.5K | 41 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 3.1K | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.7K | 0.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 2 | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 46 | 0.6% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 9 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.8K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.0K | 0.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 1 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 127 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.6K | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 33 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 105 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 208 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 194.8K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 232.7K | 140.5K | 60.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1.0K | 53.6% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 69.6K | 47.5% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 22.6K | 37.2% |
| EasyList | blocklist | adguard | 63.7K | 18.2K | 28.6% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 25.7K | 13.5% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 76.5K | 11.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 137 | 10.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1.2K | 3.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 24 | 1.4% |
| Easy Privacy | blocklist | adguard | 55.6K | 639 | 1.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 9 | 0.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 18 | 0.7% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 9 | 0.6% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 7 | 0.5% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.6K | Targets: 5 | Unique: 585 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.2% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1.1K | 90.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.2% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 40 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 9 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 630.1K | Targets: 43 | Unique: 610.0K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2.8K | 7.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 3.9K | 4.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1.6K | 0.8% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 201 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1.4K | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.1K | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 442 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 3 | 0.4% |
| HaGeZi Pro | blocklist | domain | 222.6K | 969 | 0.4% |
| kadantiscam | blocklist | domain | 44.5K | 112 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 12 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 465 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1.4K | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 35 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.5K | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 85 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 33 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 229 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 24 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 44 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 11 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 6 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 127 | 0.1% |
| phishing_army | blocklist | domain | 146.3K | 155 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 48 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 30 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 384 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 46 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 10 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 330 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 24 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 66 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 40 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 57 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 25 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 20 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 10 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 222.6K | Targets: 69 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 346 | 98.9% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 98.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 333 | 90.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 53.3K | 87.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 330 | 85.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 573 | 82.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.9K | 82.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 332 | 72.5% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| hufilter | blocklist | hostname | 92 | 65 | 70.7% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 70 | 63.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 77 | 61.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 74.9K | 39.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 111 | 37.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.0K | 37.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 84.1K | 36.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 67.9K | 34.9% |
| YousList | blocklist | hostname | 625 | 197 | 31.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 7.9K | 27.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 53.5K | 26.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 11.4K | 24.3% |
| Adaway | blocklist | hostname | 6.5K | 1.6K | 24.3% |
| WaLLy3K | blocklist | domain | 351 | 68 | 19.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.3K | 15.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 12.8K | 14.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 4.7K | 13.4% |
| kadantiscam | blocklist | domain | 44.5K | 5.9K | 13.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.3K | 10.3% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1.9K | 9.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.7K | 8.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1.0K | 7.1% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| phishing_army | blocklist | domain | 146.3K | 8.3K | 5.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 258 | 5.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 707 | 3.9% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 246 | 3.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 191 | 1.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.8K | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 62 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 62 | 1.6% |
| Spam404 | blocklist | domain | 8.1K | 123 | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 232 | 1.3% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 351 | 1.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 40 | 1.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 3.4K | 1.1% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1.1K | 1.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 24 | 0.9% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 131 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 486 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 2 | 0.8% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 574 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.7K | 0.6% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 52 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 2.3K | 0.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 4 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.6K | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.3K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 469 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 969 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 41 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 350 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.6K | 346 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 87 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 16 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.5K | Targets: 8 | Unique: 0 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 1.5K | 71.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 78 | 10.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 33 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 33 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 95 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 35.9K | Targets: 33 | Unique: 0 | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 224 | 74.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 11.9K | 73.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 688 | 69.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.3K | 64.6% |
| Greensnow | blocklist | ipv4 | 3.6K | 2.2K | 61.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.7K | 53.6% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 195 | 37.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 187 | 36.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.4K | 36.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3.3K | 35.6% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 5.5K | 30.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 5.1K | 28.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 6.2K | 26.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 260 | 25.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 2.8K | 24.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 280 | 23.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 4.2K | 23.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 13.1K | 22.7% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 122 | 14.3% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 9 | 6.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 57 | 5.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 115 | 4.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 302 | 1.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 4 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 21 | 0.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 13 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 537 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 20 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 53 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 13 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 108 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 17 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 111 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 21 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |

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
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 84 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 82 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 6 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 70 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 65 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 12 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 57 | Unique: 420.7K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 995 | 13.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1.3K | 6.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1.0K | 2.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 9 | 2.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2.7K | 1.4% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2.7K | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.8K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| phishing_army | blocklist | domain | 146.3K | 1.2K | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 617 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 98 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 2.4K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.5K | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.2K | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 124 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 2.1K | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 182 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 55 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 179 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 41 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 270 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 23 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 20 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 211 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 9 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 44.5K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 34.1K | 39.2% |
| phishing_army | blocklist | domain | 146.3K | 19.6K | 13.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 13.5K | 6.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 13.1K | 5.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 5.9K | 2.6% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 541 | 2.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.1K | 0.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 12 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 30 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 60 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 14 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 68 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 24 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 74 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 21 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 21 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 29 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 81 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 112 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 85 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 41 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |

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
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 48 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 18 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 4 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 24 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 10 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 189.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 3 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |
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
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |
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
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 3 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 9 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1 | 0.0% |

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
| Firehol_level3 | blocklist | ipv4 | 11.3K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.9K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 400 | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 1.3K | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 1.0K | 0.5% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 1.9K | 0.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3M | Targets: 60 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.1M | 97.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.7% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 272.6K | 49.5% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 9.7K | 33.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 23.5K | 23.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 49 | 10.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 19.5K | 8.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 3.7K | 8.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 11 | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2.0K | 5.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 691 | 5.3% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 539 | 3.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 5.8K | 3.0% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 6.3K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| HaGeZi Pro | blocklist | domain | 222.6K | 4.6K | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1.3K | 1.5% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 8 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 288 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 142 | 0.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 403 | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.2K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 26 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 84 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 628 | 0.3% |
| phishing_army | blocklist | domain | 146.3K | 426 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.5K | 0.2% |
| kadantiscam | blocklist | domain | 44.5K | 81 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 37 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 7 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 15 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 125 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 219.8K | Targets: 37 | Unique: 207.6K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 5.6K | 34.6% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 9 | 3.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 25 | 2.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 258 | 1.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 537 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 16 | 1.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 29 | 1.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 11 | 1.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 5 | 1.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 201 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 167 | 0.9% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 4 | 0.8% |
| Greensnow | blocklist | ipv4 | 3.6K | 27 | 0.8% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 10 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 7 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 7 | 0.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 10 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 59 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 86 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 18 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 35 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 42 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 33 | 0.2% |
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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.1M | 80.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3.9K | 78.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 222.5K | 40.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 16.4K | 16.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 42 | 9.2% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 2.2K | 7.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1.2K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 874 | 2.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 4.0K | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 4.3K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 977 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1.0K | 1.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 126 | 1.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2.3K | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.5K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 87 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 632 | 0.3% |
| kadantiscam | blocklist | domain | 44.5K | 74 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 362 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.2K | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 109 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| phishing_army | blocklist | domain | 146.3K | 363 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 16 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 17 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 5 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 8 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 69 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 21.3K | Targets: 32 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.3K | 17.4K | 11.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 16.0K | 6.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 18 | 3.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 8 | 3.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 4.2K | 2.2% |
| kadantiscam | blocklist | domain | 44.5K | 541 | 1.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 383 | 1.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 1.9K | 0.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 23 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 17 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 734 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 270 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.3K | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 103 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 28 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 100 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 38 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 87 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 24 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 84 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 232.7K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 60.8K | 60.0K | 98.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.3% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 74.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 140.5K | 72.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1.3K | 68.3% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 828 | 61.9% |
| EasyList | blocklist | adguard | 63.7K | 35.9K | 56.3% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 77.1K | 52.7% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 514 | 38.1% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 60.5K | 31.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 66 | 17.9% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 78.9K | 11.7% |
| Easy Privacy | blocklist | adguard | 55.6K | 6.3K | 11.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 3.9K | 10.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 37 | 3.6% |
| CJX Annoyance | blocklist | adguard | 1.8K | 54 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 117 | 0.6% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 33 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 232.7K | Targets: 66 | Unique: 0 | Conflicts: 21</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 60.0K | 98.8% |
| hufilter | blocklist | hostname | 92 | 82 | 89.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 387 | 84.5% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 16.0K | 75.3% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 514 | 74.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 140.5K | 72.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.6% |
| phishing_army | blocklist | domain | 146.3K | 78.1K | 53.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 267 | 42.7% |
| WaLLy3K | blocklist | domain | 351 | 135 | 38.5% |
| HaGeZi Pro | blocklist | domain | 222.6K | 84.1K | 37.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.4K | 34.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 60.5K | 32.0% |
| kadantiscam | blocklist | domain | 44.5K | 13.1K | 29.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 67.8K | 28.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 21.8K | 25.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 87 | 24.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 10.2K | 21.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 23 | 20.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.8K | 20.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 5.9K | 20.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39.9K | 19.7% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 19.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 57 | 14.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 559 | 11.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 3.9K | 10.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 36 | 9.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 10 | 3.8% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 2.0K | 2.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 86 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 112 | 0.8% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 203 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.8K | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 395 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 117 | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 463 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.3K | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 2.3K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1.2K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.3K | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 9 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.4K | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 33 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 307 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.5K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.8K | 9.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.3K | 8.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 12.3K | 5.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 22 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 2 | 0.3% |
| Torrent Trackers | blocklist | domain | 502 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 49 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 58 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 119 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 18 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 50 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 232 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 236 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 76 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 117 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 37 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 22 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 22 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 23 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 18 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 18.5K | Targets: 10 | Unique: 18.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 29 | 2.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 58 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 119 | 0.1% |
| EasyList | blocklist | adguard | 63.7K | 65 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 117 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 76 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 60.8K | Targets: 65 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 84 | 91.3% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 509 | 73.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 55.5K | 29.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 60.0K | 25.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4.7K | 25.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 53.3K | 23.9% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 21 | 19.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 32.8K | 16.2% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 5.5K | 11.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 3.4K | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 22.6K | 11.6% |
| Adaway | blocklist | hostname | 6.5K | 750 | 11.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 25 | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 22 | 6.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 916 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 582 | 4.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 3.6K | 4.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.5K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 65 | 3.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 75 | 0.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 76 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 55 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 215 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 261 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 958 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 173 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 1 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 10 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 24 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 75 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 182 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 85 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 159 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 109 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 403 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 43 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 60.8K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 844 | 63.1% |
| EasyList | blocklist | adguard | 63.7K | 35.4K | 55.6% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 509 | 37.7% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 55.5K | 29.2% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 60.0K | 25.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 22.6K | 11.6% |
| Easy Privacy | blocklist | adguard | 55.6K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1.5K | 4.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 65 | 3.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 30 | 3.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 76 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 3 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 67 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 263 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.3K | 258 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 8 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 68 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 105 | 68 | 64.8% |
| Torrent Trackers | blocklist | domain | 502 | 67 | 13.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 35.4K | Targets: 24 | Unique: 22.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 400 | 20.8% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 72 | 5.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 35 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 1.5K | 2.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 3.9K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 26 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 2.4K | 1.2% |
| EasyList | blocklist | adguard | 63.7K | 655 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.6K | 344 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 1.2K | 0.6% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 55 | 0.3% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 503 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 58 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 1.6K | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 35.4K | Targets: 76 | Unique: 0 | Conflicts: 166</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 697 | 19.6% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 64 | 14.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 48 | 9.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.5K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 523 | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 2.0K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 35 | 5.1% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 10 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1.5K | 2.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 6 | 2.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 279 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 4.7K | 2.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 552 | 1.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1.6K | 1.8% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 383 | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 3.9K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 26 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 67 | 1.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2.4K | 1.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 58 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 4.2K | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1.2K | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 470 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 2.8K | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 191 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 55 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 58 | 0.3% |
| phishing_army | blocklist | domain | 146.3K | 500 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.0K | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 7 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 340 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 874 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 60 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 22 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 124 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 61 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 78 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 146.3K | Targets: 36 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 258 | 98.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 17.4K | 82.0% |
| kadantiscam | blocklist | domain | 44.5K | 19.6K | 44.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 93.3K | 38.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 69.7K | 35.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 78.1K | 33.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 16.9K | 19.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 24 | 4.8% |
| HaGeZi Pro | blocklist | domain | 222.6K | 8.3K | 3.7% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 500 | 1.4% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1.3K | 1.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 33 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 23 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.2K | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 439 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 29 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 7 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 3 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 83 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 155 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 8 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 363 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 426 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

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
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 95 | 6.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 91 | 4.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 31 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 12 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 3 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 48 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 5 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 5 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 76 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 296 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 145 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 46 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 47 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 299 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 298 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 27 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 59 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| kadantiscam | blocklist | domain | 44.5K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 35 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 20 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 47 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 29 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 28 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 77 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 7 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 56 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 576 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 34 | 8.9% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 8 | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 433 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.2K | 6.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.3K | 3.6% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 906 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 10 | 2.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 320 | 2.2% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 6 | 1.7% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 12 | 1.7% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 1 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 916 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1.2K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 991 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2.1K | 1.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2.3K | 1.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 105 | 1 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1.4K | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 180 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Torrent Trackers | blocklist | domain | 502 | 2 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 4 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 586 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 136 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 8 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 6 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 63.7K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 671.4K | Targets: 15 | Unique: 388.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 1.9K | 99.2% |
| RPiList_specials-phishing | blocklist | adguard | 146.4K | 123.7K | 84.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 76.5K | 39.3% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 78.9K | 33.9% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1.6K | 4.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| EasyList | blocklist | adguard | 63.7K | 43 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 67 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 11 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 178 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 14 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 3 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 146.4K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 69.6K | 35.7% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 77.1K | 33.1% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 123.7K | 18.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 503 | 1.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 1.9K | 2 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 8 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.1K | Targets: 25 | Unique: 0 | Conflicts: 25</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 300 | 117 | 39.0% |
| Greensnow | blocklist | ipv4 | 3.6K | 377 | 10.6% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 1.1K | 6.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 400 | 4.3% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 1.3K | 3.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 32 | 3.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 394 | 2.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 11 | 2.2% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 11 | 2.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 326 | 1.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 986 | 1.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 30 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 13 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 3 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 39 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 25 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 44 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 36 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 25 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 13 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.0K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 216 | 18.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 12 | 8.7% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 25 | 2.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 14 | 1.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 4 | 1.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 260 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 335 | 0.6% |
| Greensnow | blocklist | ipv4 | 3.6K | 18 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 92 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 66 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 14 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 3 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 7 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 2 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 20 | Unique: 2.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 300 | 66 | 22.0% |
| Greensnow | blocklist | ipv4 | 3.6K | 79 | 2.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 104 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 2 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 115 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 22 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 1 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 8 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.5K | Targets: 20 | Unique: 0 | Conflicts: 12.6K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 606 | 3.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 33 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 33 | 1.6% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 25 | 1.2% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 57 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 87 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 76 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 3 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 24 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 19 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 31 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 10 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.2K | Targets: 28 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 217 | 21.9% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 400 | 19.4% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 46 | 15.3% |
| Greensnow | blocklist | ipv4 | 3.6K | 476 | 13.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 643 | 12.6% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 1.8K | 9.7% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 3.3K | 9.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 8.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4.6K | 8.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 1.3K | 7.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1.7K | 7.2% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 36 | 7.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 55 | 6.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 1.2K | 6.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 29 | 5.7% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 638 | 5.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 147 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 1 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 7 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 91 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 5 | 0.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 19 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 19 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 35 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 29 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 15 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 18.5K | 29 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 10 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 15 | 0.0% |

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
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 7 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 2 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 2 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 27 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 14 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 11 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 31 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 5 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 28.7K | Targets: 54 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 490 | 29.8% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 975 | 27.4% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 122 | 17.7% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 87 | 13.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1.8K | 9.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 7 | 6.4% |
| Adaway | blocklist | hostname | 6.5K | 409 | 6.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 24 | 6.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 22 | 6.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 906 | 6.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3.4K | 5.5% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2.0K | 5.5% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| HaGeZi Pro | blocklist | domain | 222.6K | 7.9K | 3.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 6.1K | 3.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.5K | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 363 | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 5.9K | 2.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 2.0K | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 879 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 49 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 128 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 168 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 14 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 204 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 55 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 11 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 33 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 288 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 17 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 36 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 236.5K | Targets: 33 | Unique: 180.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 12.3K | 66.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.6K | 28.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 469 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 462 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 78 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 208 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 84 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 10 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 194 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 75 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 307 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 107 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 125 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 6 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 29 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 156 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 69 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 21 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 30 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 211 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.8K | Targets: 44 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 5.5K | 9.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 47 | 6.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 10.1K | 5.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 11.4K | 5.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 937 | 5.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 10.2K | 4.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.3K | 3.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 5.4K | 2.8% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 56 | 1.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 180 | 1.2% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 95 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 191 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 308 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3.7K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 267 | 0.3% |
| kadantiscam | blocklist | domain | 44.5K | 68 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.2K | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 11 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 196 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 371 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 48 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 11 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.5K | Targets: 31 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 995 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 147 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 292 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 246 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 13 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 10 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 10 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 7 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 6 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 12 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 86 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 58 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 17 | 0.1% |
| phishing_army | blocklist | domain | 146.3K | 23 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 24 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 9 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 49 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 43 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 5.8K | 31.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 19.4K | 8.2% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 105 | 2 | 1.9% |
| Torrent Trackers | blocklist | domain | 502 | 9 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 615 | 0.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 215 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 128 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 486 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 12 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 395 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 78 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 105 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 44 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 22 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 275 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 39 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 37 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 18 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 2.7K | 3.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.2K | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 24 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 10 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |

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
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 15 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 29 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 32 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 17 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 62 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 35 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 38 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 31 | Unique: 4.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 1.6K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 1.1K | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 20 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 123 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 41 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 110 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 5 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 18 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.8K | 98.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.8% |

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
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 4.3K | 1.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 2.6K | 1.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.0K | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 2 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 222.6K | 707 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 128 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 389 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 17 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 20 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 7 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 87.0K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 12.9K | 98.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 98.4% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 44.5K | 34.1K | 76.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 7.6K | 41.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 401 | 24.4% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 66 | 17.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 49 | 13.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 899 | 12.1% |
| phishing_army | blocklist | domain | 146.3K | 16.9K | 11.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 21.8K | 9.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 14.1K | 7.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 2.0K | 6.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.1K | 6.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3.6K | 5.9% |
| HaGeZi Pro | blocklist | domain | 222.6K | 12.8K | 5.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 16 | 4.6% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 32 | 4.6% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.6K | 4.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 5.0K | 2.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 346 | 2.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 5.1K | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 308 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 3.9K | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 103 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 267 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 40 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 820 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 37 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.0K | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.3K | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 7 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 107 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 195 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 179 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 12 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 9 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 46 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 6.3K | 34.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 21.6K | 9.1% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 2 | 2.9% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.1% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 105 | 2 | 1.9% |
| Torrent Trackers | blocklist | domain | 502 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 703 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 157 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 168 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 261 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.6K | 574 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 95 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 13 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 344 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 463 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 50 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 127 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 48 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 13 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 20 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 43 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 38 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |

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
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 42 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 15 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 29 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 62 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 17 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 32 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 29.1K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 35 | 7.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 21.2K | 3.8% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1.8K | 1.8% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 552 | 1.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 9.7K | 0.7% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 2 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 586 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.6K | 351 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 10 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.2K | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 194 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 203 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 143 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 28 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 201 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 29 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 550.5K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 21.2K | 72.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 458 | 213 | 46.5% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 30.8K | 31.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 272.6K | 20.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 222.5K | 20.4% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 4.2K | 11.9% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6.3K | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 3.1K | 1.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.3K | 1.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2.3K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 6 | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 371 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 100 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 739 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 24 | 0.4% |
| phishing_army | blocklist | domain | 146.3K | 439 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 159 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 13 | 0.2% |
| kadantiscam | blocklist | domain | 44.5K | 69 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 195 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1.4K | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 18 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 36 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 271 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 12 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 156 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 23 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 17.7K | Targets: 29 | Unique: 0 | Conflicts: 87</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 18.1K | 7.4K | 41.0% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 235 | 27.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3.6K | 23.7% |
| Greensnow | blocklist | ipv4 | 3.6K | 641 | 18.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 855 | 16.7% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 1.9K | 16.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 326 | 15.8% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 43 | 14.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.3K | 14.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 130 | 13.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 4.2K | 11.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 2.2K | 9.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 1.7K | 9.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.2K | 8.9% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 45 | 8.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 39 | 7.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 66 | 6.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 62 | 5.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 20 | 2.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 2 | 1.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 4 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 24 | 1.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 22 | 0.9% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 87 | 0.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 87 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 59 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 167 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 502 | Targets: 8 | Unique: 308 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 105 | 104 | 99.0% |
| pexcn Torrent Trackers | blocklist | domain_url | 68 | 67 | 98.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

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
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 19 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 37 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 82 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 36 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 19 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 31 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 27 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 32 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 63.7K | 51 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 38 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 30 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 671.4K | 1 | 0.0% |

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
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| Easy Privacy | allowlist | adguard | 854 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 43 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 66 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 458 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 375 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 387 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 64 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 252 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 332 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 35 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 213 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 25 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 42 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 49 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.1K | Targets: 24 | Unique: 3.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 11.9K | 33.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 147 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 6 | 1.2% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 5 | 1.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 94 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 5 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 27 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 117 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 249 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 4 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 1 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 9 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 30 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 16 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 20 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 14 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 258 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.2K | Targets: 36 | Unique: 8.7K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 283 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 47 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 300 | 9 | 3.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5.6K | 2.5% |
| BlockListDE_Brute | blocklist | ipv4 | 852 | 15 | 1.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 985 | 18 | 1.8% |
| Greensnow | blocklist | ipv4 | 3.6K | 59 | 1.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.0K | 16 | 1.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 138 | 2 | 1.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 91 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 20 | 1.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 511 | 5 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 515 | 4 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 302 | 0.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 35 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 395 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 94 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 993 | 6 | 0.6% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 60 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 18.1K | 82 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 58 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 103 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 51 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 17.7K | 59 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
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
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 372 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 559 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 442 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 258 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 73 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 2 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 12 | Unique: 362 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 283 | 1.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 9 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 13 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.1K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.3K | 3 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.3K | 45 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.9K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 23 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.1K | 1 | 0.0% |
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
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 18 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 83 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 1 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 54 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 85 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 135 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 19 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 20 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 35 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 68 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 1 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 65 | Unique: 0 | Conflicts: 90</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.4% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 6.2K | 33.7% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 20.7K | 23.8% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 62 | 16.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 187 | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 40 | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 9 | 8.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 15.8K | 6.8% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 991 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 1.8K | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 3.1K | 5.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 1.4K | 3.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 25 | 3.6% |
| HaGeZi Pro | blocklist | domain | 222.6K | 6.7K | 3.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 5.7K | 2.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 128 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 161 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.5K | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 353 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 22 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 977 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 42 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 15 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 44 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 51 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 23 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |

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
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 198 | 1.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 3 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 87 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 241 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 2 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 94 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 150 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.6K | 197 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 42 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 267 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 17 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 60.8K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.4K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 189.9K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.6K | 10 | 0.0% |
| EasyList | blocklist | adguard | 63.7K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 194.8K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 232.7K | 69 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 74 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.6K | 41 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 45 | 0.0% |

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
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 396 | 24.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.5K | 2.0K | 10.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 18 | 5.1% |
| Adaway | blocklist | hostname | 6.5K | 267 | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 87.0K | 3.5K | 4.0% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 576 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 110 | 4 | 3.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 28.7K | 975 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 422 | 3.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 60.8K | 1.6K | 2.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 384 | 10 | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 35.4K | 697 | 2.0% |
| AdGuard Base filter | blocklist | domain_adguard | 691 | 13 | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 350 | 5 | 1.4% |
| HaGeZi Pro | blocklist | domain | 222.6K | 2.9K | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 232.7K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 189.2K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.7K | 19 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.8K | 56 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 236.5K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.5K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 194.8K | 82 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 8 | 0.0% |
| phishing_army | blocklist | domain | 146.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 21.3K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 98.8K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 630.1K | 12 | 0.0% |
| kadantiscam | blocklist | domain | 44.5K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 29.1K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.5K | 13 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 2 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.2K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

