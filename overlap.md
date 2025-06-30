# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-11-21 17:08:12 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 161 |
| Total Entries Analyzed | 9.7M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 21 |
| blocklist | 140 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| ipv4 | 39 |
| adguard | 32 |
| cidr_ipv4 | 3 |
| domain | 87 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 70 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 68.0% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 11.3K | 60.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 31.5K | 50.8% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 318 | 45.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 165 | 43.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 588 | 35.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 137 | 34.7% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 34 | 32.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 111 | 27.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 7.1K | 23.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 53.8K | 21.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 38.4K | 20.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 38.7K | 17.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 2.5K | 16.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.7K | 15.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 13.2K | 15.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 4.4K | 12.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 16.2K | 7.2% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 142 | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 16 | 2.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 232 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 3.3K | 1.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 26 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 140 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 459 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 1 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 42 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| kadantiscam | blocklist | domain | 51.2K | 85 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 445 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 11 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 33 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 740 | 0.1% |
| phishing_army | blocklist | domain | 146.4K | 79 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 432 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 630 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 255 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 920 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.9K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 22 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 29 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 7 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 36 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2.7K | 14.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.4K | 8.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 6.5K | 7.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 267 | 7.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 6 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 17 | 4.5% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 433 | 2.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 10 | 2.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 409 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 10 | 1.4% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 519 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 745 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 2.4K | 1.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.9K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 3 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 87 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1.0K | 0.5% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 14 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 66 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 50 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 92 | Targets: 8 | Unique: 31 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 216 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 3 | 0.2% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 35 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 5 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 62.0K | 511 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 698 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 516 | 0.2% |
| EasyList | blocklist | adguard | 64.3K | 101 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 11 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 158 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 35 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.9K | 44 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 2 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 702 | Targets: 31 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 511 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 139 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 679 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 57 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 576 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 318 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 516 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 44 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 158 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 35 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 33 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 2 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 308.6K | Targets: 14 | Unique: 308.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| 1Hosts (Lite) | blocklist | domain | 203.0K | 432 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 22 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 306.6K | Targets: 23 | Unique: 190.1K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 100.4K | 52.6% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 2.3K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 586 | 3.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 5.8K | 2.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 995 | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 820 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 207 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 83 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.3K | 0.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 353 | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 2 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 60 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 137 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 191.1K | Targets: 68 | Unique: 0 | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 702 | 679 | 96.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 93.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 56.7K | 91.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.5% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 100.4K | 32.8% |
| HaGeZi Pro | blocklist | domain | 245.0K | 79.8K | 32.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 61.7K | 27.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5.0K | 27.0% |
| YousList | blocklist | hostname | 625 | 150 | 24.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 6.7K | 22.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 21 | 20.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 38.4K | 18.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.5K | 15.4% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 1.8K | 11.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 24.1K | 10.7% |
| WaLLy3K | blocklist | domain | 351 | 37 | 10.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2.4K | 6.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 25 | 6.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 5.0K | 5.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 21 | 5.5% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 21 | 5.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 553 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 119 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 275 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 344 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 186 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| Torrent Trackers | blocklist | domain | 512 | 1 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 263 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 197 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 51.2K | 29 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 142 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 9 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 258 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 361 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 558 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 104 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 191.8K | Targets: 26 | Unique: 0 | Conflicts: 205</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 208 | 204 | 98.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.7% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 56.7K | 91.4% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 64.3K | 44.6K | 69.3% |
| Easy Privacy | blocklist | adguard | 55.9K | 29.0K | 51.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 698 | 51.1% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 61.7K | 27.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 24.1K | 10.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 2.4K | 6.6% |
| abpvn_hosts | blocklist | adguard | 1.0K | 22 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 92 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 119 | 0.7% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 11 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 177 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 62.0K | 847 | 1.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 830 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 71 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.9K | 77 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 15 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 208 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 216 | 204 | 94.4% |

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
| Sinfonietta_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 11 | Unique: 4.0K | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 3 | 1.6% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 9 | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 10 | 0.0% |

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
<summary>List Type: allowlist | Source Type: domain | Total: 67 | Targets: 6 | Unique: 58 | Conflicts: 4</summary>

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
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 9 | Targets: 3 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 91.8% |
| Easy Privacy | blocklist | adguard | 55.9K | 2.0K | 3.7% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 2.0K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 25 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 59 | 0.1% |
| EasyList | blocklist | adguard | 64.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 25 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.6K | Targets: 56 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.0K | 56.3% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 625 | 198 | 31.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 90 | 23.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 205 | 12.5% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 12 | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 7.6K | 9.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 57 | 8.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 4.7K | 7.5% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 2.0K | 6.5% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 21 | 5.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 21 | 5.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.5K | 4.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 7.2K | 2.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 6.4K | 2.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 5.0K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 604 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 14 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1.3K | 0.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 55 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Torrent Trackers | blocklist | domain | 512 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 18 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 43 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 36 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 14 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 12 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 142 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 9 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 465 | Targets: 20 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 285 | 13 | 4.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 140 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 33 | 1.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 71 | 1.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 8 | 1.3% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 7 | 1.2% |
| Greensnow | blocklist | ipv4 | 1.9K | 20 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 314 | 0.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 147 | 0.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 116 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 79 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 59 | 0.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 79 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 253 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 60 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 8 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 5 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 862 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 1.9K | 108 | 5.6% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 748 | 4.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 46 | 3.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 306 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 9 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 2 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 175 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 41 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 196 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 61 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 12 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 7 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 5 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 285 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 118 | 5.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 13 | 2.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 61 | 2.4% |
| Greensnow | blocklist | ipv4 | 1.9K | 47 | 2.4% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 9 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 232 | 1.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 8 | 1.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 235 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 47 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 263 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 2 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 9 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 14 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 17 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 38 | Unique: 9.9K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 338 | 0.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 190 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 190 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 22 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 69 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 35 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 127 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 264 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 40 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 35 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 53 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 4 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 49 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 146.4K | 83.5K | 57.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 240 | 50.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 2.1K | 42.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 61.0K | 27.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 60.1K | 26.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 3.2K | 3.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 370 | 2.2% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 592 | 2.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 4.2K | 1.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1.3K | 1.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 19.5K | 1.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.8K | 1.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 6.3K | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 335 | 0.9% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 119 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| kadantiscam | blocklist | domain | 51.2K | 174 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 12 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 25 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 40 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 647 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 104 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 24 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 12 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 6 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 34 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 661 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 67 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 20 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 30 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 15 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 27 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 112 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 15 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 57.7K | Targets: 33 | Unique: 2.9K | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 285 | 263 | 92.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.2K | 62.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 712 | 60.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 253 | 54.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4.4K | 47.4% |
| Greensnow | blocklist | ipv4 | 1.9K | 877 | 45.4% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 921 | 43.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 205 | 34.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 7.2K | 33.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.9K | 32.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 11.9K | 32.1% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 5.1K | 32.0% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 185 | 30.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 5.5K | 30.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 3.9K | 29.8% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 196 | 22.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 3.7K | 22.6% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 114 | 11.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 10 | 5.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 17 | 5.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 80 | 4 | 5.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 399 | 2.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 42 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 261 | 1.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 368 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 2 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 22 | Unique: 202 | Conflicts: 23</summary>

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
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 10 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 599 | Targets: 21 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 571 | 95.8% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 571 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 9 | 3.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 7 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 31 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 153 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 341 | 0.9% |
| Greensnow | blocklist | ipv4 | 1.9K | 16 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 57 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 84 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 185 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 26 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 27 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 17 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 107 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 71 | 98.6% |
| Torrent Trackers | blocklist | domain | 512 | 106 | 20.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 21 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 13.0K | 10.6K | 81.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 3.6K | 22.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 79 | 17.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 2.8K | 12.9% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 4.6K | 12.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 887 | 9.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4.9K | 8.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 322 | 6.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1.1K | 6.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 26 | 4.4% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 26 | 4.3% |
| Greensnow | blocklist | ipv4 | 1.9K | 54 | 2.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 33 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 207 | 1.3% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 4 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 32 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 9 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 11 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 24 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 8 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 54 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 101.5K | Targets: 47 | Unique: 15.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 1.8K | 6.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 30.9K | 5.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 23 | 4.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 23.6K | 1.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 76 | 1.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16.5K | 1.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 4 | 1.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.2K | 1.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 168 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 463 | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 2.6K | 1.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 185 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.9K | 0.8% |
| phishing_army | blocklist | domain | 146.4K | 1.1K | 0.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 131 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 169 | 0.6% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.0K | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 35 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 249 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 167 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 39 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 28 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 459 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 42 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 603 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 40 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 186 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 622 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 69 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 13.0K | 15.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 48 | 12.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 415 | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.1K | 5.7% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 19 | 5.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 4.6K | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 9 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 364 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 2.5K | 1.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 8 | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 573 | 0.9% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 279 | 0.8% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 673 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 553 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 73 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 168 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| kadantiscam | blocklist | domain | 51.2K | 30 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 691 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 635 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 15 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 7 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 83 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 36 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 21 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 15 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 13.9K | Targets: 7 | Unique: 13.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.4K | 11 | 0.8% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 445 | 2 | 0.4% |
| EasyList | blocklist | adguard | 64.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 20 | Unique: 41 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 46 | 5.3% |
| Greensnow | blocklist | ipv4 | 1.9K | 72 | 3.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 4 | 2.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 1 | 2.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 712 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 63 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 49 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 47 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 107 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 6 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 16 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 2 | 0.0% |
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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.7K | Targets: 10 | Unique: 25 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1.4K | 95.4% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 80 | 10.9% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 82 | 1.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 12 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 3 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2K | Targets: 10 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1.1K | 30.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 878 | 5.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 7 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 557 | 8.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 78 | 5.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 80 | 4.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 10 | Unique: 13.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3.0K | 82.7% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 56 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 165 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 10 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 33 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 14 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 2 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 3 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 40 | Unique: 11.3K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3.0K | 82.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 878 | 74.0% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 56 | 0.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 137 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 15 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 165 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 539 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 5 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 63 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 13 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 3 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 5.1K | 28.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 5.1K | 23.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 71 | 15.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 813 | 8.8% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 2.6K | 7.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.2K | 5.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 717 | 4.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 322 | 2.1% |
| Greensnow | blocklist | ipv4 | 1.9K | 37 | 1.9% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 41 | 1.9% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 3 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 1 | 0.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 2 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 15 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 22 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.9K | Targets: 19 | Unique: 13.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 92.4% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 29.0K | 15.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.7% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 2.7K | 4.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 44 | 3.2% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 6.4K | 2.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 344 | 1.0% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 206 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 3 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 11 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 64.3K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 62.0K | 35.3K | 56.9% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 44.6K | 23.3% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 35.8K | 15.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 19.9K | 8.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 101 | 7.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 659 | 1.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 66 | 0.4% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 682 | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 42 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 11 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 1 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 596 | Targets: 21 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 599 | 571 | 95.3% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 549 | 4.2% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 8 | 2.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 8 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 32 | 1.5% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 339 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 139 | 0.9% |
| Greensnow | blocklist | ipv4 | 1.9K | 16 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 56 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 80 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 205 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 26 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 17 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 27 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 7 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.7K | 99.7% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.9% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 42 | Unique: 1.2K | Conflicts: 211</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 14 | 32.6% |
| local_source_domain_allowlist | allowlist | domain | 42 | 13 | 31.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 500 | 123 | 24.6% |
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 9 | 1 | 11.1% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 56 | 7.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 3 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 3 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 6 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 42 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |

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
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 25 | 1.5% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 32 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 19 | Unique: 4.8K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 8 | 8.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 25 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 33 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 22 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 35 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 23 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 179 | Targets: 10 | Unique: 141 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 5 | 0.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 8 | 0.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 10 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 9 | Unique: 471 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 1 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 5 | Unique: 8 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 5 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.0K | Targets: 20 | Unique: 647 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 5 | 2.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 50 | 2.0% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 9 | 1.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| Greensnow | blocklist | ipv4 | 1.9K | 7 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 62 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 114 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 31 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 24 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 18 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 11 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 9 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.6K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.6% |
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.1% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.1K | Targets: 28 | Unique: 0 | Conflicts: 578</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 748 | 86.8% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 232 | 81.4% |
| Greensnow | blocklist | ipv4 | 1.9K | 1.5K | 77.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.1K | 52.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 9.2K | 42.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 5.8K | 35.6% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 153 | 25.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 116 | 24.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 139 | 23.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.8K | 19.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 6.0K | 16.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.5K | 9.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 63 | 5.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 578 | 5.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 578 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 8 | 4.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 90 | 3.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 31 | 3.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 350 | 2.7% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 28 | 1.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 73 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 35 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 59 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 13.0K | Targets: 27 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 571 | 95.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 549 | 92.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10.6K | 70.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 9.2K | 50.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 147 | 31.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.6K | 17.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 2.7K | 16.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 5.8K | 15.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7.2K | 12.6% |
| Greensnow | blocklist | ipv4 | 1.9K | 158 | 8.2% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 15 | 5.3% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 73 | 3.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 24 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 20 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 350 | 2.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 107 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 102 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 48 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 48 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 171 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.5K | Targets: 14 | Unique: 2.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 72 | 21.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 50 | 4.9% |
| Greensnow | blocklist | ipv4 | 1.9K | 21 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 28 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 3 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 15 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 335 | Targets: 6 | Unique: 235 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.0K | Targets: 19 | Unique: 6.1K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 321 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 87 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1.8K | 0.9% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2.3K | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.1K | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 347 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 55 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 131 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 22 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 116 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 62 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 12 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 60 | 0.0% |

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
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 487 | 1.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 205 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.6K | 0.7% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 402 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 588 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 187 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 26 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 66 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 32 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 4 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 1.5K | 2.7% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 26 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 66 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 32 | 0.0% |

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
| HaGeZi Pro | blocklist | domain | 245.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.9K | Targets: 25 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 285 | 47 | 16.5% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 108 | 12.5% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 1.5K | 9.3% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 72 | 6.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 110 | 5.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 20 | 4.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 16 | 2.7% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 16 | 2.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 228 | 2.5% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 758 | 2.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 877 | 1.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 228 | 1.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 178 | 1.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 21 | 0.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 19 | 0.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 7 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 158 | 0.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 37 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 1 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 54 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 31 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 23 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 380 | Targets: 18 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 90 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 17 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 46 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 31 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 60 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 19 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 333 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 165 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 36 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 6 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 105 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 23 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 20 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 65 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 224.1K | Targets: 61 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 353 | 74.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 137.0K | 60.4% |
| phishing_army | blocklist | domain | 146.4K | 65.3K | 44.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 13.1K | 44.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 22.6K | 36.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 165 | 33.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 76.9K | 31.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 37 | 29.6% |
| kadantiscam | blocklist | domain | 51.2K | 14.1K | 27.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 61.0K | 25.4% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 158 | 22.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 3.6K | 19.9% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 3.0K | 17.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 945 | 15.8% |
| Spam404 | blocklist | domain | 8.1K | 1.2K | 15.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 661 | 14.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 11.4K | 13.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 24.1K | 12.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 510 | 10.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16.2K | 8.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.4K | 7.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.3K | 7.0% |
| YousList | blocklist | hostname | 625 | 37 | 5.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1.6K | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.9K | 5.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 673 | 5.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 310 | 5.0% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 338 | 2.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 98 | 2.8% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 2.6K | 2.6% |
| WaLLy3K | blocklist | domain | 351 | 9 | 2.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 32 | 2.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 5 | 1.7% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 242 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 5 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 17 | 1.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 186 | 1.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 165 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 66 | 1.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 282 | 1.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 3 | 0.9% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 4.5K | 0.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8.8K | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.9K | 0.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.6K | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 3.4K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 35 | 0.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 60 | 0.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 255 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 213 | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 441 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 5 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 6 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 22 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 137 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 224.1K | Targets: 23 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 226.7K | 137.0K | 60.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1.0K | 49.0% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 65.3K | 44.8% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 22.6K | 36.4% |
| EasyList | blocklist | adguard | 64.3K | 19.9K | 30.9% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 24.1K | 12.5% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 79.8K | 11.6% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 158 | 11.6% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1.9K | 5.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 32 | 1.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 242 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 17 | 1.2% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 15 | 1.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 165 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 7 | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| Easy Privacy | blocklist | adguard | 55.9K | 206 | 0.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 4 | 0.3% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 5 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.6K | Targets: 5 | Unique: 600 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3.0K | 18.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 2 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 12 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1.1K | 90.6% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3.0K | 18.1% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 46 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 698.5K | Targets: 48 | Unique: 671.9K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 3.0K | 8.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 3.9K | 4.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 3.4K | 1.5% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 248 | 0.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.9K | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.4K | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 603 | 0.6% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.2K | 0.5% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 3 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 2.3K | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 12 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 647 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 36 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 61 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 101 | 0.2% |
| kadantiscam | blocklist | domain | 51.2K | 108 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.1K | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 21 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 255 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 142 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.6K | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 39 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 17 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 6 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 21 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| phishing_army | blocklist | domain | 146.4K | 141 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 46 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 7 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 34 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 395 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 48 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 33 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 44 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 137 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 26 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 70 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 343 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 57 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 24 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 10 | 0.0% |

</details>

---

### HaGeZi Most Abused TLDs

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 445 | Targets: 1 | Unique: 443 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 2 | 0.0% |

</details>

---

### HaGeZi Pro

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 245.0K | Targets: 69 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 408 | 99.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 54.7K | 88.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 333 | 87.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 343 | 86.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.7% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 576 | 82.1% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| hufilter | blocklist | hostname | 92 | 65 | 70.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 320 | 67.7% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 65 | 61.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 77 | 61.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 95.9K | 42.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 79.8K | 41.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 7.2K | 38.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 113 | 38.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4.6K | 35.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 10.4K | 35.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 76.9K | 34.3% |
| YousList | blocklist | hostname | 625 | 199 | 31.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 25.8K | 30.3% |
| Adaway | blocklist | hostname | 6.5K | 1.9K | 29.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 8.7K | 28.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 53.8K | 26.5% |
| kadantiscam | blocklist | domain | 51.2K | 10.6K | 20.7% |
| WaLLy3K | blocklist | domain | 351 | 72 | 20.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.3K | 15.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.4K | 13.7% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 4.8K | 13.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 53 | 10.6% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1.7K | 10.3% |
| phishing_army | blocklist | domain | 146.4K | 14.0K | 9.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 1.1K | 7.3% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 277 | 5.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 822 | 4.5% |
| Spam404 | blocklist | domain | 8.1K | 306 | 3.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 112 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 30 | 2.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 5.8K | 1.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 110 | 1.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 480 | 1.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4.2K | 1.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 96 | 1.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 4 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 231 | 1.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 46 | 1.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1.0K | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 604 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 707 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 25 | 0.9% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 137 | 0.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 6 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.9K | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 33 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 2.3K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.4K | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1.2K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 524 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.6K | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 42 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 411 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 408 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 111 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 16 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 87 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.5K | Targets: 10 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 1.4K | 82.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 78 | 10.7% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 83 | 1.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 37.1K | Targets: 31 | Unique: 0 | Conflicts: 34</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 235 | 82.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 314 | 67.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.4K | 64.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 9.9K | 58.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 339 | 56.9% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 341 | 56.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 4.7K | 51.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.6K | 51.3% |
| Greensnow | blocklist | ipv4 | 1.9K | 758 | 39.3% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 6.0K | 37.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.6K | 30.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 5.3K | 29.3% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 3.4K | 26.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 5.8K | 26.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 3.6K | 22.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11.9K | 20.6% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 175 | 20.3% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 107 | 9.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 62 | 6.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 119 | 4.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 5 | 2.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 308 | 1.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 15 | 0.6% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 34 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 34 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 13 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 530 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 1 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 4 | 1.1% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 23 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 30 | 0.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 20 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 53 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 17 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 113 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 13 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 107 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 83 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 12 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 81 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 65 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 70 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 57 | Unique: 421.3K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 692 | 11.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1.1K | 6.9% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 166 | 3.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 118 | 2.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 529 | 1.8% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 8 | 1.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 2.9K | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 2.6K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.9K | 0.8% |
| phishing_army | blocklist | domain | 146.4K | 1.1K | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 622 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 98 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 3.4K | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 2.4K | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 126 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.5K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.2K | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 4 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 178 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 176 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 69 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 21 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 20 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 51.2K | 37 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 263 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 174 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 11 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 51.2K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 30.9K | 36.2% |
| phishing_army | blocklist | domain | 146.4K | 22.9K | 15.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 14.1K | 6.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 13.5K | 5.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 10.6K | 4.3% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 588 | 3.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 119 | 0.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 60 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 30 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 12 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 14 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 174 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 14 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 24 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 18 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 136 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 136 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 10 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 13 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 29 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 126 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 37 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 108 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 85 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 3 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 499 | Targets: 32 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 65 | 15.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 175 | 3.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 128 | 2.1% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 15 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 165 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 48 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 18 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 53 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 10 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 23 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 4 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 191.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 3 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
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
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
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
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 5 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 10 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 9 | 1 | 11.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 1 | 0.0% |

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
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 3 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.1K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 331 | 0.9% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 1.5K | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 1.0K | 0.5% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 2.1K | 0.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3M | Targets: 60 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.1M | 92.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 4.0K | 80.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 272.6K | 49.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 7.0K | 25.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 23.6K | 23.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 44 | 9.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 19.5K | 8.1% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 11 | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2.0K | 5.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 691 | 5.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1.6K | 5.3% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 8.8K | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 539 | 3.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 6.2K | 2.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 4.4K | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1.3K | 1.5% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 8 | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 291 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 142 | 0.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 26 | 0.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.2K | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 306 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 67 | 0.4% |
| phishing_army | blocklist | domain | 146.4K | 446 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 558 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2.1K | 0.3% |
| kadantiscam | blocklist | domain | 51.2K | 136 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 13 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 37 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 98 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 219.8K | Targets: 37 | Unique: 207.8K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 23 | 51.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 5.7K | 34.2% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 8 | 2.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 8 | 1.7% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 245 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 16 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 530 | 1.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 7 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 29 | 1.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 2 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 10 | 1.0% |
| Greensnow | blocklist | ipv4 | 1.9K | 17 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 131 | 0.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 171 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 7 | 0.8% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 5 | 0.8% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 11 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 12 | 0.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 11 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 36 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 48 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 59 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 11 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 24 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 42 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 8 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 8 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 8 | 0.1% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1M | Targets: 47 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.1M | 80.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 3.9K | 78.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 222.5K | 40.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 16.5K | 16.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 44 | 9.3% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2.1K | 7.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 918 | 2.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 5.6K | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 625 | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 4.4K | 1.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 980 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1.0K | 1.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 2.6K | 1.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 127 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.5K | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 94 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| kadantiscam | blocklist | domain | 51.2K | 136 | 0.3% |
| phishing_army | blocklist | domain | 146.4K | 424 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 630 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 361 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 97 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1.6K | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 17 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 18 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 124 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.7K | Targets: 29 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.4K | 13.2K | 9.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 12.8K | 5.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 15 | 3.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 8 | 2.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 3.0K | 1.3% |
| kadantiscam | blocklist | domain | 51.2K | 588 | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 287 | 0.8% |
| HaGeZi Pro | blocklist | domain | 245.0K | 1.7K | 0.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 20 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 15 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 185 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.1K | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 370 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 94 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 94 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 88 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 67 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 226.7K | Targets: 67 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 61.3K | 98.9% |
| hufilter | blocklist | hostname | 92 | 81 | 88.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 413 | 87.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 12.8K | 76.9% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 516 | 73.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 137.0K | 61.1% |
| phishing_army | blocklist | domain | 146.4K | 72.9K | 49.8% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 264 | 42.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 95.9K | 39.2% |
| WaLLy3K | blocklist | domain | 351 | 134 | 38.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 107 | 36.4% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6.4K | 34.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 61.7K | 32.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 9.0K | 30.5% |
| kadantiscam | blocklist | domain | 51.2K | 13.5K | 26.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 60.1K | 25.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.4K | 24.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 19.4K | 22.7% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 23 | 21.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 6.5K | 21.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 87 | 21.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.8K | 20.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 38.7K | 19.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 18.9% |
| Spam404 | blocklist | domain | 8.1K | 1.5K | 18.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 57 | 14.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 579 | 11.7% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 3.8K | 10.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 36 | 9.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 18 | 3.6% |
| tranco | allowlist | domain_top | 500 | 15 | 3.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 6 | 2.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1.9K | 1.9% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 190 | 1.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 84 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 7 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 244 | 0.9% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 116 | 0.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 118 | 0.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 27 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.6K | 0.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 460 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 391 | 0.6% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 27 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.2K | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 2.3K | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1.3K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 4.4K | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1.9K | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 33 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 286 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 226.7K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 62.0K | 61.3K | 98.9% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 73.7% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1.5K | 68.5% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 830 | 61.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 137.0K | 61.1% |
| EasyList | blocklist | adguard | 64.3K | 35.8K | 55.6% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 71.9K | 49.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 516 | 37.8% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 61.7K | 32.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 65 | 17.7% |
| Easy Privacy | blocklist | adguard | 55.9K | 6.4K | 11.4% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 76.6K | 11.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 3.8K | 10.6% |
| abpvn_hosts | blocklist | adguard | 1.0K | 36 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 54 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 37 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 16 | 1.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 118 | 0.7% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 9 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 33 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 7 | 0.1% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.8K | Targets: 39 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.5K | 9.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.0K | 7.9% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 11.8K | 5.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 22 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 2 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| Torrent Trackers | blocklist | domain | 512 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 59 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 52 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 9 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 118 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 231 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 18 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 119 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 78 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 242 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 22 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 232 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 21 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 37 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 17 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 39 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 17.8K | Targets: 10 | Unique: 17.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 29 | 2.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 59 | 0.2% |
| EasyList | blocklist | adguard | 64.3K | 66 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 242 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 118 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 78 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 119 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 62.0K | Targets: 63 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 83 | 90.2% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 511 | 72.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.3% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 56.7K | 29.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 61.3K | 27.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 4.7K | 25.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 54.7K | 22.3% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 20 | 19.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.8K | 16.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31.5K | 15.5% |
| YousList | blocklist | hostname | 625 | 93 | 14.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 4.0K | 13.3% |
| Adaway | blocklist | hostname | 6.5K | 745 | 11.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 22.6K | 10.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 24 | 6.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 911 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 22 | 5.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 573 | 4.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 3.6K | 4.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 66 | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.5K | 4.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 11 | 2.9% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 5 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 69 | 0.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 62 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 78 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 214 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 995 | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 260 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 167 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 40 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 24 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 77 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 148 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 101 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 178 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 306 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 97 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 62.0K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 847 | 63.1% |
| EasyList | blocklist | adguard | 64.3K | 35.3K | 54.8% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 511 | 37.4% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 56.7K | 29.5% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 61.3K | 27.0% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 42 | 11.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 22.6K | 10.1% |
| Easy Privacy | blocklist | adguard | 55.9K | 2.7K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1.5K | 4.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 66 | 4.0% |
| abpvn_hosts | blocklist | adguard | 1.0K | 29 | 2.8% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 59 | 2.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 78 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 61 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 291 | Targets: 11 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.4K | 284 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 6 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 72 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 107 | 71 | 66.4% |
| Torrent Trackers | blocklist | domain | 512 | 71 | 13.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 35.9K | Targets: 76 | Unique: 0 | Conflicts: 166</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 697 | 19.6% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 56 | 11.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 48 | 9.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.5K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 519 | 7.9% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 7 | 6.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 2.0K | 6.6% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 35 | 5.0% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 10 | 2.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1.5K | 2.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 6 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 279 | 2.1% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 571 | 2.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 4.8K | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1.6K | 1.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 287 | 1.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 3.8K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 26 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 6 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 67 | 1.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2.4K | 1.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 58 | 1.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 4.2K | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1.9K | 0.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 3 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 168 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 463 | 0.5% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 3.0K | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 59 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 56 | 0.3% |
| phishing_army | blocklist | domain | 146.4K | 475 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| kadantiscam | blocklist | domain | 51.2K | 119 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 7 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 22 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.0K | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 918 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 335 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 76 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 60 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 126 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 35.9K | Targets: 24 | Unique: 22.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 331 | 15.6% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 35 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 1.5K | 2.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 3.8K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 26 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 2.4K | 1.2% |
| EasyList | blocklist | adguard | 64.3K | 659 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 1.9K | 0.8% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.9K | 344 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 59 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 56 | 0.3% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 477 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 1.5K | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 146.4K | Targets: 37 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 284 | 97.6% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 13.2K | 79.1% |
| kadantiscam | blocklist | domain | 51.2K | 22.9K | 44.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 83.5K | 34.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 72.9K | 32.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 65.3K | 29.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 16.0K | 18.8% |
| HaGeZi Pro | blocklist | domain | 245.0K | 14.0K | 5.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 23 | 4.6% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 475 | 1.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1.1K | 1.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 30 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 21 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.1K | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 1 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 445 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 7 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 35 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 424 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 446 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 141 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 79 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 4 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 6.7K | Targets: 14 | Unique: 5.9K | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 557 | 76.2% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 83 | 5.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 82 | 4.8% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 27 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 27 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 3 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 76 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 296 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 145 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 46 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 47 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 299 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 298 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 49 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 10 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 20 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 47 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 28 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 35 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 8 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 18 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 77 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 4 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 55 | Unique: 0 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 576 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 33 | 8.4% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 8 | 7.6% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 433 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.2K | 6.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.3K | 3.5% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 913 | 3.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 10 | 2.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 321 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 12 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 911 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 6 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1.2K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 991 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2.1K | 1.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 2.3K | 1.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 107 | 1 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.4K | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 159 | 0.5% |
| Torrent Trackers | blocklist | domain | 512 | 2 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 586 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 4 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 186 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 22 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 6 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 1 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 682 | Targets: 1 | Unique: 681 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 64.3K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 685.7K | Targets: 15 | Unique: 400.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 2.1K | 99.3% |
| RPiList_specials-phishing | blocklist | adguard | 145.7K | 124.7K | 85.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 79.8K | 35.6% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 76.6K | 33.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1.5K | 4.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 14 | 0.1% |
| EasyList | blocklist | adguard | 64.3K | 42 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 61 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 10 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 177 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.9K | 3 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 13.9K | 1 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 145.7K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | adguard | 226.7K | 71.9K | 31.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 65.3K | 29.2% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 124.7K | 18.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 477 | 1.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 3 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 11 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 2 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.1K | Targets: 23 | Unique: 0 | Conflicts: 36</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 285 | 118 | 41.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 33 | 7.1% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 1.1K | 7.0% |
| Greensnow | blocklist | ipv4 | 1.9K | 110 | 5.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 32 | 5.4% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 31 | 5.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 396 | 4.3% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 1.4K | 3.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 564 | 3.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 433 | 2.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 921 | 1.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 41 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 1 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 13 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 56 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 73 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 36 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 36 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 33 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 19 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 11 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 14 | Unique: 2.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 285 | 61 | 21.4% |
| Greensnow | blocklist | ipv4 | 1.9K | 19 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 90 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 119 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 17 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 8 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 7 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 11.5K | Targets: 23 | Unique: 0 | Conflicts: 11.5K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 578 | 3.6% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 32 | 1.9% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 36 | 1.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 5 | 1.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 77 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 48 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 74 | 0.4% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 27 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 28 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 1 | 0.2% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 1 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 34 | 0.1% |
| Greensnow | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 9 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.2K | Targets: 26 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 140 | 30.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 396 | 18.5% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 47 | 16.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 813 | 15.9% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 4.7K | 12.7% |
| Greensnow | blocklist | ipv4 | 1.9K | 228 | 11.8% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 1.8K | 11.2% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 57 | 9.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 56 | 9.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4.4K | 7.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 1.6K | 7.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1.3K | 7.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 887 | 5.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 971 | 5.9% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 621 | 4.8% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 41 | 4.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 173 | 1.0% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 5 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 69 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 28 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 28 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 36 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 29 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 10 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 17 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 16 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 30 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 17.8K | 29 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 10 | 0.0% |
| EasyList | blocklist | adguard | 64.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 17 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 16 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 718 | Targets: 39 | Unique: 377 | Conflicts: 268</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 9 | 1 | 11.1% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 56 | 3.3% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 7 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 2 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 2 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 14 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 27 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 31 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 30.0K | Targets: 53 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 487 | 29.7% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 973 | 27.4% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 139 | 19.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 87 | 13.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2.0K | 10.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 31 | 8.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 7 | 6.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 26 | 6.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 4.0K | 6.4% |
| Adaway | blocklist | hostname | 6.5K | 409 | 6.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 913 | 6.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2.0K | 5.5% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| HaGeZi Pro | blocklist | domain | 245.0K | 8.7K | 3.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 7.1K | 3.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 6.7K | 3.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 6.5K | 2.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 364 | 2.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 2.0K | 2.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 3 | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1.6K | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 52 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 131 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 170 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 12 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 207 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 39 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 291 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 11 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 17 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 69 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 39 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 214.0K | Targets: 33 | Unique: 160.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 11.8K | 66.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 18.3K | 29.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20.3K | 26.5% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 524 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 76 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 445 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 441 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 69 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 286 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 197 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 77 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 103 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 18 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 34 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 98 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 124 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 4 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 140 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 174 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 24 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 29.5K | Targets: 43 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 49 | 39.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 4.8K | 7.7% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 44 | 6.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 13.1K | 5.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 10.4K | 4.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 9.0K | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 604 | 3.2% |
| YousList | blocklist | hostname | 625 | 16 | 2.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 4.5K | 2.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.7K | 2.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 53 | 1.5% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 159 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 73 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 168 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 31 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 35 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 258 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 169 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 132 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 41 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 22 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 625 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 11 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 36 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.6K | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 529 | 0.1% |
| kadantiscam | blocklist | domain | 51.2K | 60 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 6 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 61 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 194 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 119 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.2K | Targets: 29 | Unique: 3.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 897 | 1.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 14 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 310 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 692 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 22 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| kadantiscam | blocklist | domain | 51.2K | 12 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 84 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 96 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 6 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 10 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 10 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 7 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 27 | Unique: 326 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 128 | 25.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 945 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 58 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 15 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 110 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 118 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 21 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 27 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 5.5K | 31.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 18.3K | 8.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 2 | 2.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 107 | 2 | 1.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 63 | 1.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Torrent Trackers | blocklist | domain | 512 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 613 | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 131 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 214 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 78 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 391 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 13 | 0.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 604 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 213 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 22 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 36 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 275 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 37 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 19 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 39 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 16 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 11 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 2.7K | 3.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1.2K | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 10 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 25 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.8K | Targets: 34 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3.8K | 100.0% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 42 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 32 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 38 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 61 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 15 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 35 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 31 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 1.5K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1.2K | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 14 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 31 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 52 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 306 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 110 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 22 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 18 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 56 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.7K | 99.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.9% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 2.4K | 47.9% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 17.6K | 7.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 4.4K | 1.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 3.6K | 1.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 2 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.0K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 5.0K | 0.4% |
| HaGeZi Pro | blocklist | domain | 245.0K | 822 | 0.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 389 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 131 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 17 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 4 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 20 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 85.2K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 99.8% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 51.2K | 30.9K | 60.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 7.6K | 41.0% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 402 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 70 | 17.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 897 | 14.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 46 | 12.1% |
| phishing_army | blocklist | domain | 146.4K | 16.0K | 10.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 25.8K | 10.5% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 9 | 8.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 19.4K | 8.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 2.0K | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 3.6K | 5.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 11.4K | 5.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 33 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.6K | 4.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 31 | 4.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 16 | 3.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 5.0K | 2.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 347 | 2.3% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 613 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 258 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 701 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 94 | 0.6% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 3.9K | 0.6% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.3K | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 291 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 40 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 820 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 14 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 249 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 37 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.3K | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 7 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 20 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 1.0K | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 176 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 103 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 220 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 47 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 6.0K | 34.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 20.3K | 9.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 2 | 2.8% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 71 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 107 | 2 | 1.9% |
| Torrent Trackers | blocklist | domain | 512 | 9 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 701 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 157 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 170 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 260 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 95 | 0.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 707 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 14 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 344 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 460 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 255 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 41 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 51.2K | 13 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 20 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 11 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 48 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 43 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 40 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |

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
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 42 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 17 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 38 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 61 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 32 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 27.9K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 50 | 10.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 18.6K | 3.4% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1.8K | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 571 | 1.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7.0K | 0.5% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 245.0K | 480 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2.1K | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 8 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 592 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 140 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 22 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 244 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 282 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 248 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 7 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 9 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 550.7K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 18.6K | 66.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 473 | 205 | 43.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 30.9K | 30.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 272.6K | 20.6% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 222.5K | 19.4% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 4.2K | 11.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6.3K | 2.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 4.5K | 2.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 2.3K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 6 | 0.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 2.3K | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 194 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 88 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 25 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 740 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 220 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| phishing_army | blocklist | domain | 146.4K | 445 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 2.3K | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 148 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| kadantiscam | blocklist | domain | 51.2K | 126 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 16 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 12 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 258 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 39 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 140 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 23 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.4K | Targets: 28 | Unique: 0 | Conflicts: 77</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 16.1K | 5.8K | 36.3% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 306 | 35.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 564 | 26.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3.6K | 24.2% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 2.7K | 20.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 79 | 17.0% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 84 | 14.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 717 | 14.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 80 | 13.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 971 | 10.6% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 3.6K | 9.8% |
| Greensnow | blocklist | ipv4 | 1.9K | 178 | 9.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1.4K | 7.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 1.5K | 7.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.7K | 6.4% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 17 | 6.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 49 | 4.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 11 | 1.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 77 | 0.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 77 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 179 | 1 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 54 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 32 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 5 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 131 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 1 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 512 | Targets: 8 | Unique: 312 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 107 | 106 | 99.1% |
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 71 | 98.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 1 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 45 | Unique: 0 | Conflicts: 530</summary>

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
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 19 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 82 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 36 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 37 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 19 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 27 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 31 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 32 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 15 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 64.3K | 51 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 685.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 37 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 29 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 12 | Unique: 24 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 55.9K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 856 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 42 | 0.1% |
| EasyList | blocklist | adguard | 64.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 65 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 473 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 353 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 56 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 413 | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 50 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 240 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 320 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 23 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 44 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 44 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 205 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 8 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.8K | Targets: 24 | Unique: 5.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 9.9K | 26.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 173 | 1.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 5 | 0.8% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 4 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 96 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 5 | 0.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 102 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 261 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 1 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 32 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 35 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 10 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 4 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 20 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 245 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 13 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.6K | Targets: 34 | Unique: 9.2K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 283 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 47 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 285 | 9 | 3.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5.7K | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 18 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 465 | 8 | 1.7% |
| BlockListDE_Brute | blocklist | ipv4 | 862 | 12 | 1.4% |
| Greensnow | blocklist | ipv4 | 1.9K | 23 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 19 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 308 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 69 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 399 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 96 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 16.1K | 73 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 596 | 3 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 107 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 22 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 55 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 16.4K | 54 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 599 | 2 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 44 | 0.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 3 | 0.2% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 3 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 32 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 6.7K | 3 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 5.0K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.1K | 0.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.0K | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 579 | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 3.9K | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 510 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 442 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 76 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 277 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 11 | Unique: 367 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 283 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.8K | 10 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.9K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 18.2K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 13.0K | 3 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 13 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 13.0K | 45 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 37.1K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 23 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 5 | 0.0% |

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
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 18 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 83 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 54 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 134 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 19 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 84 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 72 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 37 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 20 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 65 | Unique: 0 | Conflicts: 90</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.2% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6.2K | 33.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 20.7K | 24.3% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 60 | 15.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 187 | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 44 | 11.1% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 9 | 8.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 15.8K | 7.0% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 991 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 1.8K | 5.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 3.1K | 4.9% |
| HaGeZi Pro | blocklist | domain | 245.0K | 10.4K | 4.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 1.4K | 3.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 25 | 3.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 5.4K | 2.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 8 | 1.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 131 | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 132 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.5K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 980 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 353 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 21 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 42 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 23 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 51 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 46 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 15 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 198 | 1.1% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 3 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 87 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 42 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 93 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 264 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 245.0K | 199 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 16 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 150 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 37 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| EasyList | blocklist | adguard | 64.3K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 224.1K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 226.7K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 62.0K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 35.9K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 191.8K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.9K | 10 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 74 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 45 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 245.0K | 42 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |

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
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2.0K | 10.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 18 | 5.1% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 85.2K | 3.5K | 4.2% |
| Adaway | blocklist | hostname | 6.5K | 267 | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| HaGeZi Apple Tracker | blocklist | domain | 105 | 4 | 3.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 576 | 3.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 415 | 3.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.0K | 973 | 3.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 395 | 11 | 2.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 62.0K | 1.6K | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 702 | 13 | 1.9% |
| ph00lt0_blocklist | blocklist | domain | 35.9K | 697 | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 718 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 411 | 5 | 1.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Pro | blocklist | domain | 245.0K | 3.0K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 380 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 226.7K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 191.1K | 1.6K | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 53 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 71 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.0K | 19 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 63 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.1M | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 14 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.8K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 308.6K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 698.5K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| phishing_army | blocklist | domain | 146.4K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 51.2K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 27.9K | 2 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 224.1K | 98 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 214.0K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 550.7K | 13 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.7K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 306.6K | 30 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.5K | 7 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 4 | Unique: 8.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 16.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 46 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

