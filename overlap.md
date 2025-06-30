# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-11-28 03:47:39 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs. blocklist, allowlist vs. allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 163 |
| Total Entries Analyzed | 10.1M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 22 |
| blocklist | 141 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 34 |
| cidr_ipv4 | 3 |
| domain | 87 |
| ipv4 | 39 |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 67.7% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 11.3K | 60.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 31.4K | 49.5% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 318 | 45.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 165 | 43.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 586 | 35.9% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 34 | 34.3% |
| hufilter | blocklist | hostname | 92 | 31 | 33.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 138 | 33.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 111 | 26.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 7.1K | 23.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 52.2K | 21.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.6K | 20.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 37.8K | 19.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 38.6K | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 2.5K | 16.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.7K | 15.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 13.2K | 14.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 4.4K | 11.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 12.8K | 7.6% |
| tranco | allowlist | domain_top | 500 | 33 | 6.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 142 | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 16 | 2.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 232 | 1.3% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 3.3K | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 26 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 143 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 457 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 1 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 42 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 11 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 447 | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 51.3K | 85 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 32 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.2K | 0.1% |
| phishing_army | blocklist | domain | 146.8K | 78 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 630 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 430 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 211 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 742 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 283 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.0K | Targets: 9 | Unique: 919 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 22 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 29 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 36 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 267 | 7.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 6.5K | 7.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 6 | 6.1% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 17 | 4.5% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 433 | 2.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 10 | 2.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 11 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 519 | 1.4% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 10 | 1.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 409 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 744 | 1.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 2.4K | 1.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.9K | 0.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 3 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 91 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1.0K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 50 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 49 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 92 | Targets: 8 | Unique: 31 | Conflicts: 60</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 215 | 1 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 5 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 1 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 35 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 10 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.5K | 515 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 703 | 0.4% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 520 | 0.2% |
| EasyList | blocklist | adguard | 64.8K | 108 | 0.2% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 3 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 35 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 125 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 11 | 0.1% |
| Easy Privacy | blocklist | adguard | 56.0K | 44 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 2 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 8 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 707 | Targets: 31 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 515 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 140 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 684 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 58 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 239.3K | 537 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 45 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 318 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 520 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 35 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 125 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 33 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 332.2K | Targets: 14 | Unique: 331.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 430 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 9 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 367.2K | Targets: 22 | Unique: 248.2K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 103.2K | 53.0% |
| hufilter | blocklist | hostname | 92 | 16 | 17.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 2.3K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 585 | 3.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 5.5K | 2.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1.0K | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 815 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 206 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 83 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1.3K | 0.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 356 | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 2 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 60 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 43 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 24 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 195.5K | Targets: 26 | Unique: 0 | Conflicts: 204</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 207 | 203 | 98.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1.5K | 92.7% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 58.3K | 91.9% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 64.8K | 45.0K | 69.5% |
| Easy Privacy | blocklist | adguard | 56.0K | 28.9K | 51.7% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 703 | 51.3% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 63.5K | 28.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 19.1K | 11.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 2.4K | 6.5% |
| abpvn_hosts | blocklist | adguard | 1.0K | 22 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 92 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.8K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 116 | 0.7% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 10 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 225 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 11 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 194.7K | Targets: 68 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 707 | 684 | 96.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 93.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 58.3K | 91.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 92 | 70 | 76.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| HaGeZi Pro | blocklist | domain | 239.3K | 77.9K | 32.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 63.5K | 28.5% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 103.2K | 28.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5.0K | 27.0% |
| YousList | blocklist | hostname | 625 | 150 | 24.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 6.8K | 22.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 37.8K | 18.6% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 17 | 17.2% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.5K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 1.8K | 11.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 19.1K | 11.3% |
| WaLLy3K | blocklist | domain | 351 | 37 | 10.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2.4K | 6.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 25 | 6.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 5.0K | 5.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 21 | 5.5% |
| tranco | allowlist | domain_top | 500 | 27 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 21 | 5.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 553 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 116 | 0.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 52 | 0.5% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 344 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 275 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 183 | 0.2% |
| Torrent Trackers | blocklist | domain | 519 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 2 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 197 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 10 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| kadantiscam | blocklist | domain | 51.3K | 28 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 263 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 269 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 8 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 103 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 3 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 351 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 165 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 548 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 848 | 1.3% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 831 | 0.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 72 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 56.0K | 77 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 10 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 2 | 0.0% |

</details>

---

### AdGuardSDNSFilter_exceptions

<details>
<summary>List Type: allowlist | Source Type: adguard | Total: 207 | Targets: 1 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | adguard | 215 | 203 | 94.4% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1 | 0.0% |

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
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 10 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
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
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 9 | Targets: 3 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1.5K | 91.8% |
| Easy Privacy | blocklist | adguard | 56.0K | 2.0K | 3.6% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 2.0K | 1.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 2.0K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 59 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 25 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 24 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 1 | 0.0% |

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
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 90 | 23.6% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 206 | 12.6% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 12 | 12.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 7.6K | 8.4% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 58 | 8.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.2K | 8.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1.1K | 8.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 4.7K | 7.4% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 2.0K | 6.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.3K | 5.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 21 | 5.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 21 | 5.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.5K | 4.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 6.4K | 2.9% |
| HaGeZi Pro | blocklist | domain | 239.3K | 7.0K | 2.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 5.0K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 604 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 14 | 1.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 1.0K | 0.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 55 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| Torrent Trackers | blocklist | domain | 519 | 1 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 43 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 39 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 14 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 142 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 12 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 25 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 27 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 951 | Targets: 21 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 280 | 18 | 6.4% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 15 | 2.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 14 | 2.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 108 | 2.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 181 | 1.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 597 | 1.7% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 185 | 1.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 232 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 27 | 1.5% |
| Greensnow | blocklist | ipv4 | 2.1K | 28 | 1.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 224 | 1.0% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 147 | 0.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 154 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 119 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 468 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 12 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 16.9K | 972 | 5.7% |
| Greensnow | blocklist | ipv4 | 2.1K | 118 | 5.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 302 | 1.9% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 3 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 5 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 49 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 73 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 158 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 1 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 134 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 1 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 18 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 7 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 5 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 6 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 280 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 131 | 7.3% |
| Greensnow | blocklist | ipv4 | 2.1K | 76 | 3.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 60 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 12 | 2.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 12 | 2.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 18 | 1.9% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 240 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 226 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 52 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 257 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 3 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 11 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 19 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 15 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 5 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 38 | Unique: 10.0K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 24 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 35 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 182 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 69 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 188 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 203 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 22 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 8 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 52 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 49 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 6 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 264 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 40 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 127 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 240.0K | Targets: 49 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| phishing_army | blocklist | domain | 146.8K | 80.8K | 55.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 237 | 49.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 2.1K | 42.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 47.4K | 28.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 54.0K | 24.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 3.2K | 3.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 574 | 2.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 308 | 1.9% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 16.8K | 1.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 19.5K | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1.3K | 1.4% |
| HaGeZi Pro | blocklist | domain | 239.3K | 2.9K | 1.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 6.3K | 1.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 290 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 119 | 0.4% |
| kadantiscam | blocklist | domain | 51.3K | 172 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 12 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 40 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1.0K | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 25 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 103 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 12 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 24 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.6K | Targets: 35 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4.1K | 68.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 175 | 35.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 61 | 14.6% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 1 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 67 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 23 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 89 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 15 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 15 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 92 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 76 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 27 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 57.7K | Targets: 33 | Unique: 2.0K | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 280 | 257 | 91.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.4K | 67.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 712 | 60.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 918 | 51.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 468 | 49.2% |
| Greensnow | blocklist | ipv4 | 2.1K | 941 | 43.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 4.3K | 42.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5.6K | 37.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 7.2K | 32.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 11.5K | 32.5% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 4.0K | 31.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 5.3K | 30.7% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 4.8K | 28.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 4.3K | 27.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 139 | 24.1% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 140 | 24.1% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 134 | 12.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 113 | 11.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 14 | 8.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 139 | 5.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 17 | 5.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 80 | 4 | 5.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 412 | 2.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 1 | 2.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 11 | 2.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 42 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 266 | 1.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 42 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 42 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 368 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 1 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 10 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 2 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 581 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 570 | 98.8% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 559 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 12 | 4.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 30 | 1.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 15 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 150 | 0.9% |
| Greensnow | blocklist | ipv4 | 2.1K | 17 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 263 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 76 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 55 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 58 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 140 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 13 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 19 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 4 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 113 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 72 | 100.0% |
| Torrent Trackers | blocklist | domain | 519 | 109 | 21.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 23 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.6K | 9.9K | 78.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 4.0K | 25.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 232 | 24.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 3.5K | 15.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 658 | 12.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 4.3K | 12.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.1K | 10.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 58 | 10.1% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 58 | 10.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.6K | 9.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1.2K | 7.0% |
| Greensnow | blocklist | ipv4 | 2.1K | 50 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 263 | 1.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 27 | 1.5% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 3 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 7 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 44 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 20 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 16 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 16 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 21 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

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
| OISD Blocklist Small | blocklist | adguard | 63.5K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 53 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 4 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 101.7K | Targets: 47 | Unique: 15.1K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 1.7K | 6.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 32.1K | 5.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 24 | 5.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 5 | 1.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 23.6K | 1.8% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 76 | 1.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 16.5K | 1.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 168 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3.2K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 2.1K | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 462 | 1.3% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 165 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1.8K | 0.8% |
| phishing_army | blocklist | domain | 146.8K | 1.1K | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 131 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 169 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 250 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 166 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 35 | 0.3% |
| HaGeZi Pro | blocklist | domain | 239.3K | 752 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 457 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 6 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 32 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 18 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 40 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 13 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 27 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 654 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 624 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 183 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 70 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |

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
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 13.0K | 14.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 422 | 11.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 48 | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.1K | 5.7% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| hufilter | blocklist | hostname | 92 | 5 | 5.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 19 | 5.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 4.4K | 1.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 9 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 364 | 1.2% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 8 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 2.5K | 1.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 573 | 0.9% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 279 | 0.8% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 992 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 553 | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 549 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 168 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 73 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 691 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 635 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 13 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 51.3K | 26 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 36 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 8 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 15 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 83 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 21 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 14.0K | Targets: 7 | Unique: 13.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 1.4K | 11 | 0.8% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 4 | 0.5% |
| HaGeZi Most Abused TLDs | blocklist | adguard | 445 | 2 | 0.4% |
| EasyList | blocklist | adguard | 64.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.2K | Targets: 18 | Unique: 298 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 4 | 2.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 1 | 2.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 712 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 47 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 1 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 32 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 1 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 48 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 4 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 16 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 3 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.7K | Targets: 9 | Unique: 68 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 1.4K | 95.2% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 80 | 10.9% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 47 | 1.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 13 | 0.0% |
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
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 874 | 5.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 6 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 7 | Unique: 454 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 78 | 5.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 80 | 4.7% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 47 | 1.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 16.3K | Targets: 10 | Unique: 13.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 2.9K | 82.2% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 57 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 97 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 33 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 14 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 2 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 10 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 40 | Unique: 11.4K | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 2.9K | 82.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 874 | 73.5% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 57 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 182 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 97 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 15 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 33 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 539 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 18 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 8 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 13 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 10 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 3 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 63 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 42 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 4 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 5.1K | 29.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 5.1K | 23.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 108 | 11.4% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 3.3K | 9.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 892 | 8.8% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 1.1K | 7.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 3.4K | 6.0% |
| Greensnow | blocklist | ipv4 | 2.1K | 121 | 5.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 658 | 4.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 31 | 1.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 2 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 2 | 0.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 16 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 22 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 12 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 56.0K | Targets: 19 | Unique: 13.6K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1.5K | 92.5% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 28.9K | 14.8% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 77 | 5.7% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 2.7K | 4.2% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 44 | 3.2% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 6.4K | 2.9% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 344 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.0K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 166 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| EasyList | blocklist | adguard | 64.8K | 11 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 64.8K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.5K | 36.6K | 57.7% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 45.0K | 23.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 37.2K | 16.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 15.8K | 9.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 108 | 7.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 50 | 3.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 658 | 1.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 67 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 683 | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| Easy Privacy | blocklist | adguard | 56.0K | 11 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 42 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 1 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 577 | Targets: 22 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 581 | 570 | 98.1% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 553 | 4.4% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 12 | 4.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 14 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 26 | 1.5% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 143 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 258 | 0.7% |
| Greensnow | blocklist | ipv4 | 2.1K | 15 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 73 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 52 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 58 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 12 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 17 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 3 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 5 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.8K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 100.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 34.1% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 41 | Unique: 1.2K | Conflicts: 210</summary>

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 56 | 7.8% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 11 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 3 | 1.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 30 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 3 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 8 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 6 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 42 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 11 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |

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
| Public_DNS4 | blocklist | ipv4 | 3.9K | 24 | 0.6% |
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
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 22 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 24 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 23 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 14 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 33 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 15 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 170 | Targets: 9 | Unique: 135 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 14 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 8 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 3 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 8 | Unique: 473 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 1 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 2 | 0.0% |

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
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.0K | Targets: 21 | Unique: 629 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 48 | 1.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 1 | 0.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 5 | 0.5% |
| Greensnow | blocklist | ipv4 | 2.1K | 9 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 64 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 113 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 32 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 18 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 23 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 8 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 13 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 10 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 7 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.6K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.6K | 89.7% |
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.6K | 89.7% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 16.9K | Targets: 27 | Unique: 0 | Conflicts: 699</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 972 | 88.4% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 240 | 85.7% |
| Greensnow | blocklist | ipv4 | 2.1K | 1.8K | 82.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.2K | 64.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 8.2K | 37.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 5.5K | 35.2% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 150 | 25.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 2.6K | 25.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 143 | 24.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 154 | 16.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 5.3K | 15.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 5.3K | 9.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 8.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 699 | 6.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 699 | 6.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 8 | 4.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 97 | 3.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 32 | 3.2% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 387 | 3.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 35 | 1.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 3 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 72 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 3 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 69 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.6K | Targets: 27 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 46 | 100.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 559 | 96.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 553 | 95.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.9K | 66.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 8.2K | 47.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 224 | 23.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 6.1K | 17.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 2.4K | 15.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.6K | 15.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7.2K | 12.5% |
| Greensnow | blocklist | ipv4 | 2.1K | 158 | 7.3% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 15 | 5.4% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 3 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 46 | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 23 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 387 | 2.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 18 | 1.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 109 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 50 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 50 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 58 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 2 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 155 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 80 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 72 | 21.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 48 | 4.8% |
| Greensnow | blocklist | ipv4 | 2.1K | 26 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 35 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 11 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 15 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 29 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 3 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 335 | Targets: 8 | Unique: 230 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 72 | 2.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 3 | 0.0% |
| Greensnow | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 17 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 18 | Unique: 6.2K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 319 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 91 | 1.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1.8K | 0.9% |
| YousList | blocklist | hostname | 625 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 2.3K | 0.6% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.1K | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 353 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 55 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 135 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 59 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 117 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 23 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 12 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 394 | 11.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 485 | 1.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 206 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1.5K | 0.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.6K | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 400 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 586 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 187 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 26 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 31 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 4 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.6K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 54.8% |
| Easy Privacy | blocklist | adguard | 56.0K | 1.5K | 2.7% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 1.5K | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 66 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 26 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 31 | 0.0% |

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
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 13 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.1K | Targets: 23 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 280 | 76 | 27.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 220 | 12.3% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 118 | 10.7% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 1.8K | 10.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 28 | 2.9% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 17 | 2.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 273 | 2.7% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 926 | 2.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 15 | 2.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 121 | 2.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 292 | 1.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 941 | 1.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 214 | 1.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 33 | 1.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 26 | 1.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 9 | 0.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 158 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 1 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 42 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 50 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 12 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 13 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 381 | Targets: 18 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 90 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 17 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 165 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 19 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 333 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 31 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 46 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 60 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 6 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 36 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 21 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 99 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 12 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 17 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 19 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 59 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 169.2K | Targets: 57 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 401 | 83.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 117.8K | 53.0% |
| phishing_army | blocklist | domain | 146.8K | 56.6K | 38.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 19.3K | 30.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 8.7K | 29.6% |
| HaGeZi Pro | blocklist | domain | 239.3K | 62.5K | 26.1% |
| kadantiscam | blocklist | domain | 51.3K | 12.6K | 24.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 27 | 21.6% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 47.4K | 19.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 3.1K | 19.0% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 125 | 17.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 13.9K | 15.3% |
| Spam404 | blocklist | domain | 8.1K | 1.2K | 14.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 19.1K | 9.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 39 | 7.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5.2K | 6.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 12.8K | 6.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2.1K | 5.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.0K | 5.4% |
| YousList | blocklist | hostname | 625 | 32 | 5.1% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 549 | 4.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 192 | 3.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 11 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 821 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 89 | 2.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 133 | 2.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 2.1K | 2.1% |
| WaLLy3K | blocklist | domain | 351 | 7 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 89 | 2.0% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 31 | 1.9% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 203 | 1.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 88 | 1.5% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 285 | 1.1% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 3 | 0.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 132 | 0.9% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 11 | 0.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 4.2K | 0.7% |
| Adaway | blocklist | hostname | 6.5K | 49 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 4.3K | 0.6% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 97 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.1K | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.1K | 0.5% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 4.6K | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 24 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 43 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 136 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 113 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 168 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 24 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 4 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 169.2K | Targets: 23 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 1.1K | 53.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 117.8K | 53.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 56.4K | 38.5% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 19.3K | 30.4% |
| EasyList | blocklist | adguard | 64.8K | 15.8K | 24.3% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 19.1K | 9.8% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 125 | 9.1% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 65.6K | 7.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 2.1K | 5.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 31 | 1.9% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 24 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 11 | 0.8% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 10 | 0.7% |
| abpvn_hosts | blocklist | adguard | 1.0K | 7 | 0.7% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 97 | 0.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| Easy Privacy | blocklist | adguard | 56.0K | 166 | 0.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 12 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 43 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.6K | Targets: 6 | Unique: 616 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 2.9K | 18.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 9 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 3 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 13 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1.1K | 90.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2.9K | 18.0% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 44 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 9 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 718.4K | Targets: 48 | Unique: 686.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1.2K | 45.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3.0K | 8.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 3.9K | 4.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4.3K | 2.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 3.9K | 1.8% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 7 | 1.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 240 | 0.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.5K | 0.7% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.8K | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 654 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 2.3K | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.0K | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 80 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 12 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 17 | 0.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 2 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 21 | 0.2% |
| kadantiscam | blocklist | domain | 51.3K | 98 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.8K | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 39 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2.2K | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 109 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 44 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 50 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 15 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 165 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 283 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| phishing_army | blocklist | domain | 146.8K | 119 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 13 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 6 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 18 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 35 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 11 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 411 | Targets: 15 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.1K | 48 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 33 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 26 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 138 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 339 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 70 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 44 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 24 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 57 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 239.3K | Targets: 69 | Unique: 0 | Conflicts: 38</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 408 | 99.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 333 | 87.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 53.2K | 83.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 339 | 82.5% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 537 | 76.0% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| hufilter | blocklist | hostname | 92 | 64 | 69.6% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 59 | 59.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 74 | 59.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 90.7K | 40.8% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 77.9K | 40.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 113 | 38.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 7.0K | 37.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 62.5K | 36.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 10.3K | 34.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4.4K | 33.9% |
| YousList | blocklist | hostname | 625 | 205 | 32.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 27.4K | 30.1% |
| Adaway | blocklist | hostname | 6.5K | 1.9K | 29.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 8.6K | 28.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 52.2K | 25.7% |
| WaLLy3K | blocklist | domain | 351 | 89 | 25.4% |
| kadantiscam | blocklist | domain | 51.3K | 12.3K | 24.0% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 91 | 18.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.3K | 15.3% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.4K | 13.7% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 4.4K | 11.9% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1.5K | 9.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 45 | 9.0% |
| phishing_army | blocklist | domain | 146.8K | 13.1K | 8.9% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 1.1K | 7.5% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 17 | 6.1% |
| Spam404 | blocklist | domain | 8.1K | 304 | 3.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 30 | 2.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 92 | 2.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 98 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 93 | 1.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 182 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 61 | 1.6% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 5.5K | 1.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 94 | 1.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 281 | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 231 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2.9K | 1.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 44 | 1.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 182 | 1.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 586 | 1.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 25 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 689 | 0.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 6 | 0.8% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 752 | 0.7% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 140 | 0.5% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 6 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 33 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.8K | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1.6K | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3.4K | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1.8K | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 532 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2.2K | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 42 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 412 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 239.3K | 408 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 21 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 111 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 87 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 25 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 22 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 16 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.5K | Targets: 9 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 1.4K | 82.3% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 78 | 10.7% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 48 | 1.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 32 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 12 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 35.4K | Targets: 30 | Unique: 0 | Conflicts: 30</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 226 | 80.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 3.3K | 63.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 597 | 62.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 1.1K | 61.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 9.1K | 53.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 263 | 45.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 258 | 44.7% |
| Greensnow | blocklist | ipv4 | 2.1K | 926 | 43.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 3.6K | 35.8% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 5.3K | 31.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 5.1K | 29.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.3K | 28.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 6.1K | 27.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 4.0K | 25.7% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 2.7K | 21.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 11.5K | 19.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 158 | 14.4% |
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 5 | 10.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 64 | 6.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 115 | 4.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 48 | 4.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 310 | 1.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 3 | 1.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 15 | 0.6% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 14 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 30 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 30 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 517 | 0.2% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 294 | Targets: 24 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 4 | 1.0% |
| tranco | allowlist | domain_top | 500 | 4 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 23 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 3 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 20 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 53 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 107 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 17 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 13 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 113 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 21 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 83 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 70 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 16 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 64 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 81 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 56 | Unique: 422.3K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 692 | 11.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1.1K | 6.9% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 10.8K | 4.5% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 166 | 3.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 118 | 2.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 529 | 1.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 76 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 1 | 1.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 7 | 1.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 2.1K | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 2.5K | 1.1% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| phishing_army | blocklist | domain | 146.8K | 1.1K | 0.7% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.8K | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 624 | 0.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 98 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 3.5K | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 4.5K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.2K | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 2.4K | 0.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 34 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 177 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 127 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 4 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 175 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 69 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 20 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 177 | 0.1% |
| kadantiscam | blocklist | domain | 51.3K | 36 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 20 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 263 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 13 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 51.3K | Targets: 41 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 42.1K | 46.4% |
| phishing_army | blocklist | domain | 146.8K | 22.9K | 15.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 12.6K | 7.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 13.9K | 6.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 12.3K | 5.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 688 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 3 | 1.1% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 103 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 14 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 12 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 60 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 26 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 172 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 14 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 138 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 28 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 7 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 98 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 18 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 24 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 36 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 137 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 85 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 133 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 12 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |

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
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 128 | 2.1% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 1 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 48 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 18 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 10 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 11 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 45 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 39 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 4 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 21 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 5 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 195.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 3 | 0.0% |

</details>

---

### local_ai_allowlist

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 52 | Targets: 1 | Unique: 0 | Conflicts: 52</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_blocklist | blocklist | ipv4_from_domain | 52 | 52 | 100.0% |

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
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |

</details>

---

### local_ai_blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 52 | Targets: 1 | Unique: 0 | Conflicts: 52</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_ai_allowlist | allowlist | ipv4_from_domain | 52 | 52 | 100.0% |

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
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
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
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 3 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 1 | 0.0% |

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
| Firehol_level3 | blocklist | ipv4 | 12.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.2K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 301 | 0.8% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 1.5K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 1.1K | 0.7% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 2.1K | 0.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 2 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 219.8K | Targets: 37 | Unique: 207.9K | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4.4K | 86.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 23 | 50.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 5.7K | 33.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 8 | 2.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 6 | 1.8% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 62 | 1 | 1.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 517 | 1.5% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 16 | 1.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 232 | 1.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 2 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 29 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 10 | 1.0% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 5 | 0.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 137 | 0.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 5 | 0.9% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 13 | 0.8% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 12 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 8 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 155 | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 17 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 6 | 0.6% |
| Greensnow | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 368 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 46 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 69 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 47 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 6 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 33 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 12 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 21 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 4 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 8 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 15 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 8 | 0.1% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3M | Targets: 60 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1.1M | 92.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 4.0K | 80.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 272.7K | 48.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 35 | 28.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 6.4K | 23.9% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 23.6K | 23.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 42 | 8.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 19.5K | 8.1% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 11 | 6.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2.2K | 6.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1.6K | 5.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 691 | 5.3% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 6.1K | 3.6% |
| WaLLy3K | blocklist | domain | 351 | 12 | 3.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 539 | 3.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.5K | 3.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 264 | 2.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 4.9K | 2.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 10 | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1.4K | 1.5% |
| tranco | allowlist | domain_top | 500 | 7 | 1.4% |
| HaGeZi Pro | blocklist | domain | 239.3K | 3.4K | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 110 | 1.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5.2K | 1.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 8 | 1.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 292 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 142 | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.2K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 26 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 303 | 0.5% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 65 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2.8K | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 3 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 26 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| phishing_army | blocklist | domain | 146.8K | 438 | 0.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 548 | 0.3% |
| kadantiscam | blocklist | domain | 51.3K | 137 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 15 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 43 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 37 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 98 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |

</details>

---

### Maltrail_StaticTrails_Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.2M | Targets: 47 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.1M | 80.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 3.9K | 78.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 222.9K | 39.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.4% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 16.5K | 16.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 18 | 14.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 41 | 8.5% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 2.1K | 7.8% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 16.8K | 7.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 635 | 4.9% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.2K | 3.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4.6K | 2.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 629 | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 3.2K | 1.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1.1K | 1.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 980 | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 127 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4.5K | 1.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 2.2K | 0.9% |
| Spam404 | blocklist | domain | 8.1K | 56 | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 78 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2.2K | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 630 | 0.3% |
| phishing_army | blocklist | domain | 146.8K | 420 | 0.3% |
| kadantiscam | blocklist | domain | 51.3K | 133 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 351 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 134 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 17 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 95 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 18 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 9 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 8 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 16 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.2K | Targets: 30 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.8K | 12.7K | 8.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 12.6K | 5.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 18 | 3.6% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 9 | 3.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 3.1K | 1.8% |
| kadantiscam | blocklist | domain | 51.3K | 688 | 1.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 292 | 0.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.5K | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 23 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.1K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 165 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 308 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 22 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 120 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 78 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 18 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 32 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 65 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 86 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 222.5K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 63.5K | 62.8K | 98.9% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1.6K | 98.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 73.7% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 1.5K | 70.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 117.8K | 69.6% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 831 | 61.9% |
| EasyList | blocklist | adguard | 64.8K | 37.2K | 57.4% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 72.4K | 49.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 520 | 37.9% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 63.5K | 32.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 65 | 17.7% |
| Easy Privacy | blocklist | adguard | 56.0K | 6.4K | 11.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 3.8K | 10.4% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 72.9K | 8.6% |
| abpvn_hosts | blocklist | adguard | 1.0K | 36 | 3.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 53 | 2.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 37 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 15 | 1.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 69 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 119 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 9 | 0.3% |
| AdBlockID | blocklist | adguard | 3.8K | 10 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 33 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.1% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 222.5K | Targets: 67 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 62.8K | 98.9% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 98.9% |
| hufilter | blocklist | hostname | 92 | 81 | 88.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 405 | 83.9% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 12.6K | 77.5% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 520 | 73.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 117.8K | 69.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 66.8% |
| phishing_army | blocklist | domain | 146.8K | 73.5K | 50.1% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 264 | 42.2% |
| WaLLy3K | blocklist | domain | 351 | 134 | 38.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 90.7K | 37.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 107 | 36.4% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6.4K | 34.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 63.5K | 32.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 8.5K | 28.8% |
| kadantiscam | blocklist | domain | 51.3K | 13.9K | 27.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 23.2K | 25.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 54.0K | 22.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 6.4K | 21.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 87 | 21.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.7K | 20.7% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 19 | 19.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 38.6K | 19.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2.5K | 18.9% |
| Spam404 | blocklist | domain | 8.1K | 1.5K | 18.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 57 | 13.9% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3.8K | 10.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 36 | 9.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| tranco | allowlist | domain_top | 500 | 15 | 3.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 7 | 2.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 11 | 2.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1.8K | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 188 | 1.6% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 83 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 15 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 260 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 7 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 117 | 0.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 119 | 0.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 454 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 388 | 0.6% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 3.9K | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.5K | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.9K | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 2.2K | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 17 | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 10 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 1.3K | 0.4% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 51 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 9 | 0.3% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 3.2K | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 14 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 33 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 12 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 281 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.5K | Targets: 39 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 5.5K | 9.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 6.0K | 7.8% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 11.7K | 5.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 28 | 1.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 23 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 2 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 52 | 0.2% |
| Torrent Trackers | blocklist | domain | 519 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 60 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 79 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 8 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 231 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 232 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 17 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 119 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 116 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 22 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 43 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 32 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 36 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 5 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 17.5K | Targets: 10 | Unique: 17.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 28 | 1.9% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 60 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| EasyList | blocklist | adguard | 64.8K | 67 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 119 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 79 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 116 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 43 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 12 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 3 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 63.5K | Targets: 63 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 92 | 83 | 90.2% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 515 | 72.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.5% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 58.3K | 30.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 62.8K | 28.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 4.7K | 25.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 53.2K | 22.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 16 | 16.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4.7K | 16.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31.4K | 15.5% |
| YousList | blocklist | hostname | 625 | 93 | 14.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 4.0K | 13.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 19.3K | 11.4% |
| Adaway | blocklist | hostname | 6.5K | 744 | 11.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 910 | 6.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 24 | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 22 | 5.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 573 | 4.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 66 | 4.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 3.6K | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.4K | 3.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 11 | 2.9% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 5 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 69 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 79 | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 59 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 214 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 1.0K | 0.3% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 260 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 166 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 177 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 95 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 77 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 40 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 24 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 303 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 109 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 150 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 63.5K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 848 | 63.1% |
| EasyList | blocklist | adguard | 64.8K | 36.6K | 56.5% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 515 | 37.6% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 58.3K | 29.8% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 62.8K | 28.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 19.3K | 11.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 42 | 11.4% |
| Easy Privacy | blocklist | adguard | 56.0K | 2.7K | 4.8% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 66 | 4.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1.4K | 3.9% |
| abpvn_hosts | blocklist | adguard | 1.0K | 29 | 2.8% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 59 | 2.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 79 | 0.5% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 2 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 2 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 72 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 8 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 280 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 146.8K | 275 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 9 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 11 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 17 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 11 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 7 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 72 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 113 | 72 | 63.7% |
| Torrent Trackers | blocklist | domain | 519 | 71 | 13.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 36.7K | Targets: 24 | Unique: 23.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 301 | 13.9% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 72 | 5.4% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 35 | 2.6% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 1.4K | 2.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 3.8K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 26 | 1.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 2.1K | 1.2% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 2.4K | 1.2% |
| EasyList | blocklist | adguard | 64.8K | 658 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 56.0K | 344 | 0.6% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 57 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 60 | 0.3% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 471 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 1.6K | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.6K | 3 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 1.0K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 36.7K | Targets: 76 | Unique: 0 | Conflicts: 166</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 697 | 19.6% |
| tranco | allowlist | domain_top | 500 | 82 | 16.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 47 | 13.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 52 | 10.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 48 | 9.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.5K | 8.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| Adaway | blocklist | hostname | 6.5K | 519 | 7.9% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 7 | 7.1% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 2.0K | 6.5% |
| hufilter | blocklist | hostname | 92 | 6 | 6.5% |
| hkamran80_smarttv | blocklist | domain | 294 | 17 | 5.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 144 | 5.4% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 35 | 5.0% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 11 | 3.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 27 | 3.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 10 | 2.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 10 | 2.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1.4K | 2.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.4K | 2.2% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 558 | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 279 | 2.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 4.4K | 1.8% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 292 | 1.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1.6K | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 3.8K | 1.7% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 26 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 6 | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 67 | 1.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 2.4K | 1.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 2.1K | 1.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 5.5K | 1.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 58 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 3 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 169 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 462 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 3.0K | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| phishing_army | blocklist | domain | 146.8K | 472 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 57 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 60 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.2K | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 23 | 0.2% |
| kadantiscam | blocklist | domain | 51.3K | 103 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 7 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1.2K | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 290 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 95 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 78 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 127 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 76 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 60 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 146.8K | Targets: 36 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 275 | 98.2% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 12.7K | 77.9% |
| kadantiscam | blocklist | domain | 51.3K | 22.9K | 44.7% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 80.8K | 33.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 56.6K | 33.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 73.5K | 33.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 18.7K | 20.5% |
| HaGeZi Pro | blocklist | domain | 239.3K | 13.1K | 5.5% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 21 | 4.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 472 | 1.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1.1K | 1.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 27 | 0.6% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 20 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.1K | 0.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 7 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 456 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 34 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 78 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 119 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 3 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 420 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 438 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 11 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 8 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 2 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.9K | Targets: 11 | Unique: 3.7K | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 24 | 36.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 47 | 6.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 48 | 3.3% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 47 | 2.8% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 20 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 20 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 76 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 145 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 296 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 46 | 0.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 298 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 298 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 47 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |

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
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 49 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 22 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 28 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 27 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 74 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 18 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 47 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 35 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 20 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |

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
| HaGeZi Apple Tracker | blocklist | domain | 99 | 8 | 8.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 33 | 8.0% |
| tranco | allowlist | domain_top | 500 | 37 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 433 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1.2K | 6.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.3K | 3.5% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| hufilter | blocklist | hostname | 92 | 3 | 3.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 912 | 3.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 10 | 2.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 319 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 12 | 1.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 6 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 189 | 1.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 910 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 10 | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 991 | 1.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1.2K | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 2.1K | 1.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 2.3K | 1.0% |
| CF_Torrent_Trackers | blocklist | domain_url | 113 | 1 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1.4K | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 159 | 0.5% |
| Torrent Trackers | blocklist | domain | 519 | 2 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 4 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 585 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 132 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 9 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 22 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 25 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 7 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 25 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 683 | Targets: 1 | Unique: 682 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 64.8K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 846.0K | Targets: 17 | Unique: 577.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 2.1K | 99.3% |
| RPiList_specials-phishing | blocklist | adguard | 146.7K | 125.5K | 85.6% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 65.6K | 38.8% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 72.9K | 32.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1.6K | 4.3% |
| AdGuard Base filter | blocklist | adguard | 1.4K | 8 | 0.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 72 | 0.1% |
| EasyList | blocklist | adguard | 64.8K | 42 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 225 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 16.3K | 14 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 12 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 3 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 146.7K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 56.4K | 33.3% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 72.4K | 32.6% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 125.5K | 14.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 471 | 1.3% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.2K | 2 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 11 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.8K | Targets: 22 | Unique: 0 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 280 | 131 | 46.8% |
| Greensnow | blocklist | ipv4 | 2.1K | 220 | 10.2% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 1.2K | 6.9% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 30 | 5.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 26 | 4.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 458 | 4.5% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 1.1K | 3.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 27 | 2.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 918 | 1.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 215 | 1.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 25 | 1.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 31 | 0.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 84 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 46 | 0.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 27 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 2 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 46 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 9 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 14 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 6 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 14 | Unique: 2.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 280 | 60 | 21.4% |
| Greensnow | blocklist | ipv4 | 2.1K | 33 | 1.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 25 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 97 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 115 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 139 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 17 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 2 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 5 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 8 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.5K | Targets: 22 | Unique: 0 | Conflicts: 12.6K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 11.5K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 11.5K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 699 | 4.1% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 32 | 1.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 14 | 0.8% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 20 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 50 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 4 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 78 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 46 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 19 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 1 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 30 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 16 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 42 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 8 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 8 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 10.1K | Targets: 26 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 458 | 25.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 181 | 19.0% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 52 | 18.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 892 | 17.4% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 2.6K | 15.1% |
| Greensnow | blocklist | ipv4 | 2.1K | 273 | 12.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 1.7K | 11.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 3.6K | 10.3% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 55 | 9.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 52 | 9.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1.3K | 7.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4.3K | 7.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 1.6K | 7.1% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 803 | 6.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 49 | 4.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 164 | 1.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 8 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 94 | 0.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 4 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 19 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 19 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 33 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 17.5K | 28 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 10 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 15 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.4K | Targets: 18 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 12 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 28 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 11 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 15 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 30 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 719 | Targets: 38 | Unique: 381 | Conflicts: 265</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 9 | 1 | 11.1% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 56 | 3.3% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 67 | 2 | 3.0% |
| tranco | allowlist | domain_top | 500 | 10 | 2.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 7 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 2 | 1.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 3 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 351 | 3 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 2 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 27 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 14 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 6 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 5 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 7 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 31 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 30.1K | Targets: 53 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 485 | 29.7% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 974 | 27.3% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 140 | 19.8% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 87 | 13.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 92 | 11 | 12.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2.0K | 10.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 31 | 8.1% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 7 | 7.1% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| Adaway | blocklist | hostname | 6.5K | 409 | 6.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 26 | 6.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 4.0K | 6.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 912 | 6.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2.0K | 5.4% |
| tranco | allowlist | domain_top | 500 | 19 | 3.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 8.6K | 3.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 7.1K | 3.5% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 6.8K | 3.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 6.4K | 2.9% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 364 | 2.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 2.0K | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 821 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 52 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 131 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 170 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 12 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 206 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 44 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 39 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 11 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 292 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 17 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 12 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 69 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 215.2K | Targets: 33 | Unique: 162.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 11.7K | 66.7% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 18.3K | 29.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 20.3K | 26.5% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 2 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 447 | 0.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 532 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 76 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 168 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 15 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 70 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 103 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 51 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 281 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 77 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 197 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 134 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 140 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 18 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 24 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 35 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 98 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 177 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 3 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 29.5K | Targets: 43 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 49 | 39.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 4.7K | 7.5% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 45 | 6.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 8.7K | 5.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 10.3K | 4.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 8.5K | 3.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 604 | 3.2% |
| YousList | blocklist | hostname | 625 | 16 | 2.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 4.5K | 2.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.7K | 2.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 53 | 1.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| hufilter | blocklist | hostname | 92 | 1 | 1.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 159 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 73 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 169 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 31 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 35 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 258 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 169 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 132 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 15 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 529 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 629 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 41 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 22 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 11 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 36 | 0.1% |
| kadantiscam | blocklist | domain | 51.3K | 60 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.6K | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 119 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 8 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 6 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 80 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 201 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.2K | Targets: 29 | Unique: 4.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 897 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 14 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 692 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 133 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 142 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 23 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 10 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 12 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 12 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 83 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 17 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 94 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 28 | Unique: 1.2K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 128 | 25.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 2 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 58 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 88 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 16 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 26 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 9 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 5 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 20 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 118 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 93 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 7 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 25 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 44 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.2K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 5.5K | 31.5% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 18.3K | 8.5% |
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 2 | 2.8% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| CF_Torrent_Trackers | blocklist | domain_url | 113 | 2 | 1.8% |
| Torrent Trackers | blocklist | domain | 519 | 9 | 1.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 613 | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 131 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 214 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 13 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 78 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 586 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 388 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 36 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 275 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 113 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 22 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 11 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 16 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 37 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 40 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 19 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.7K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 2.7K | 3.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 144 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1.2K | 0.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 25 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 2 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 42 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 29 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 61 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 15 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 32 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 38 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 35 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 31 | Unique: 4.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 1.2K | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 1.5K | 0.7% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 14 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 52 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 304 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 31 | 0.1% |
| kadantiscam | blocklist | domain | 51.3K | 22 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 18 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 11 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 56 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 110 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 6 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 5 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 1 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.8K | 1.7K | 98.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.6K | 33.8% |

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
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 1.1K | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 5.0K | 0.4% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 5.0K | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 2 | 0.4% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 390 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 281 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 131 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 6 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 20 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 51 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 13 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 4 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 90.9K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.6K | 99.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 13.0K | 99.3% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 51.3K | 42.1K | 82.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 7.6K | 41.0% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 400 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 84 | 23.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 70 | 17.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 897 | 14.5% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| hufilter | blocklist | hostname | 92 | 12 | 13.0% |
| phishing_army | blocklist | domain | 146.8K | 18.7K | 12.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 46 | 12.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 27.4K | 11.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 23.2K | 10.4% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 9 | 9.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 13.9K | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 2.0K | 6.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 3.6K | 5.7% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 33 | 4.7% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.6K | 4.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 31 | 4.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 16 | 3.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 5.0K | 2.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 353 | 2.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 613 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 701 | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 258 | 0.9% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 280 | 2 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 120 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 3.9K | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1.3K | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 40 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 815 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 36 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 250 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 14 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 20 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1.4K | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 294 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 7 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 1.1K | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 103 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 175 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 6 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 47 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.2K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 6.0K | 34.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 20.3K | 9.4% |
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 2 | 2.8% |
| hufilter | blocklist | hostname | 92 | 2 | 2.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.0% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 113 | 2 | 1.8% |
| Torrent Trackers | blocklist | domain | 519 | 9 | 1.7% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 133 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 157 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 701 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 170 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 260 | 0.4% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 95 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 239.3K | 689 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 454 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 14 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 344 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 136 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 9 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 41 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 20 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 40 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 14 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 43 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 11 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 49 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 20 | 2.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 30 | 1.8% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 42 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 61 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 29 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 35 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 32 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 38 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 15 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 26.6K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 58 | 12.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 18.3K | 3.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1.7K | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 558 | 1.5% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 4 | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 6.4K | 0.5% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 574 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 285 | 0.2% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2.1K | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 8 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 260 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 140 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 143 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 22 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 8 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 240 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 34 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 6 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 562.1K | Targets: 48 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 18.3K | 68.7% |
| URLHaus (Abuse.ch) | blocklist | hostname | 483 | 215 | 44.5% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 32.1K | 31.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 272.7K | 20.7% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 222.9K | 19.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 5.5K | 15.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 5.0K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 6.3K | 2.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4.2K | 2.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 390 | 2.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 2.2K | 1.0% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 6 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 1.6K | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 201 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 86 | 0.5% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 63 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 742 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 25 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 36 | 0.3% |
| phishing_army | blocklist | domain | 146.8K | 456 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 294 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 2.3K | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| kadantiscam | blocklist | domain | 51.3K | 138 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 10 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 150 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 140 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 15 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 40 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 49 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 39 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 12 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 269 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.7K | Targets: 27 | Unique: 0 | Conflicts: 46</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 16.9K | 5.5K | 32.6% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 302 | 27.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.0K | 26.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.1K | 21.8% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 2.4K | 19.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 1.7K | 17.0% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 76 | 13.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 73 | 12.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 119 | 12.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 215 | 12.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 4.0K | 11.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 2.4K | 10.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1.8K | 10.3% |
| Greensnow | blocklist | ipv4 | 2.1K | 214 | 10.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 4.3K | 7.5% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 19 | 6.8% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 32 | 2.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 13 | 1.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 1 | 0.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 46 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 46 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 47 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 5 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 23 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 137 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 6 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 519 | Targets: 8 | Unique: 316 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 72 | 71 | 98.6% |
| CF_Torrent_Trackers | blocklist | domain_url | 113 | 109 | 96.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 1 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 45 | Unique: 0 | Conflicts: 533</summary>

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
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 10 | 1.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 28 | 0.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 182 | 1 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 19 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 37 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 36 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 82 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 10 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.2K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.6K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 19 | 0.1% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 11 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 31 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 33 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 15 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 32 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 27 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 34 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 64.8K | 50 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 846.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 56.0K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 37 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 29 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 368 | Targets: 12 | Unique: 24 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 5 | 0.4% |
| Easy Privacy | blocklist | adguard | 56.0K | 164 | 0.3% |
| Easy Privacy | allowlist | adguard | 859 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.6K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 42 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 65 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 483 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 405 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 401 | 0.2% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 58 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 52 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 237 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 24 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 42 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 215 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 41 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 91 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 176.1K | Targets: 1 | Unique: 176.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 17.0K | Targets: 26 | Unique: 7.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 9.1K | 25.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 1 | 2.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 164 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 4 | 0.7% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 4 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 100 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 266 | 0.5% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 1 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 4 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 16 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 58 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 20 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 3 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 9 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 20 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 15 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 20 | 0.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 23 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 232 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 17.0K | Targets: 35 | Unique: 9.6K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 5 | 10.9% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 283 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 47 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 280 | 9 | 3.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 5.7K | 2.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.0K | 18 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 951 | 12 | 1.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 10.1K | 94 | 0.9% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 310 | 0.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 412 | 0.7% |
| Greensnow | blocklist | ipv4 | 2.1K | 12 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 100 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 170 | 1 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 335 | 2 | 0.6% |
| BruteforceBlocker | blocklist | ipv4_find | 581 | 3 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 6 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 109 | 0.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.8K | 9 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 577 | 3 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 51 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 22 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 16.9K | 72 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 44 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 47 | 0.3% |
| HaGeZi_DoH | blocklist | ipv4 | 1.5K | 3 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 1.7K | 3 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 34 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 3.9K | 2 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.5K | 4 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.5K | 4 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |

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
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 3.9K | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 4.0K | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 76 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 192 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 442 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 98 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 14 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 13 | Unique: 360 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 46 | 3 | 6.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 4.4K | 2.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 283 | 1.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 9 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.6K | 3 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 15.7K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 57.7K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.0K | 4 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 14 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.4K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 176.1K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 46 | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.6K | 46 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 35.4K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 23 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.0K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 5 | 0.0% |
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
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 83 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 20 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 54 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 84 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 19 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 134 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 89 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 12 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 20 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 37 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 7 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 65 | Unique: 0 | Conflicts: 90</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.3% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| YousList | blocklist | hostname | 625 | 231 | 37.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 6.2K | 33.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 2 | 28.6% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 3.6K | 27.7% |
| WaLLy3K | blocklist | domain | 351 | 81 | 23.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 20.7K | 22.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 67 | 19.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 60 | 15.7% |
| hkamran80_smarttv | blocklist | domain | 294 | 45 | 15.3% |
| hufilter | blocklist | hostname | 92 | 14 | 15.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 187 | 11.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 44 | 10.7% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 9 | 9.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 15.7K | 7.1% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 991 | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 1.8K | 5.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 3.1K | 4.8% |
| HaGeZi Pro | blocklist | domain | 239.3K | 10.4K | 4.3% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 1.4K | 3.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 27 | 3.8% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 25 | 3.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 5.2K | 3.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 8 | 1.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 3.2K | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 135 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 132 | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 2.5K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 980 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.7K | 3 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 21 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 356 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 15 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 41 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 51 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 25 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 50 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 33 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 3 | 0.0% |

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
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 3 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 108 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 25 | 0.7% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 2 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 87 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 241 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 150 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 42 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 93 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 264 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 16 | 0.1% |
| HaGeZi Pro | blocklist | domain | 239.3K | 205 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 32 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 1 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| Easy Privacy | blocklist | adguard | 56.0K | 10 | 0.0% |
| EasyList | blocklist | adguard | 64.8K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 169.2K | 12 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 222.5K | 69 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 63.5K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 36.7K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 195.5K | 39 | 0.0% |

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
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 74 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 7 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| HaGeZi Pro | blocklist | domain | 239.3K | 42 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 9 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 45 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |

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
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 394 | 24.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.6K | 2.0K | 10.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| WaLLy3K | blocklist | domain | 351 | 18 | 5.1% |
| hufilter | blocklist | hostname | 92 | 4 | 4.3% |
| Adaway | blocklist | hostname | 6.5K | 267 | 4.1% |
| YousList | blocklist | hostname | 625 | 25 | 4.0% |
| HaGeZi Apple Tracker | blocklist | domain | 99 | 4 | 4.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 90.9K | 3.6K | 3.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 576 | 3.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.4% |
| Dan Pollock's List | blocklist | hostname | 13.1K | 422 | 3.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 30.1K | 974 | 3.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 411 | 11 | 2.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 63.5K | 1.6K | 2.6% |
| ph00lt0_blocklist | blocklist | domain | 36.7K | 697 | 1.9% |
| AdGuard Base filter | blocklist | domain_adguard | 707 | 13 | 1.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 719 | 11 | 1.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 412 | 5 | 1.2% |
| HaGeZi Pro | blocklist | domain | 239.3K | 3.0K | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 222.5K | 2.4K | 1.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 381 | 4 | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 194.7K | 1.6K | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 29.5K | 53 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 169.2K | 89 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 15.2K | 19 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 101.7K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 215.2K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 562.1K | 13 | 0.0% |
| phishing_army | blocklist | domain | 146.8K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 1.3M | 15 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 367.2K | 30 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 16.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 332.2K | 8 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 16.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 17.5K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 240.0K | 4 | 0.0% |
| Maltrail_StaticTrails_Domains | blocklist | domain | 1.2M | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 6.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 51.3K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 26.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 718.4K | 12 | 0.0% |

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
| Maltrail_StaticTrails | blocklist | ipv4_find | 219.8K | 46 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 17.0K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

