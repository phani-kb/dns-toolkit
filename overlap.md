# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2026-09-12 13:17:44 UTC

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
| Total Entries Analyzed | 6.6M |

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
<summary>List Type: blocklist | Source Type: domain | Total: 203.0K | Targets: 69 | Unique: 0 | Conflicts: 51</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 4.9K | 74.6% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 241 | 68.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 68.3% |
| YousList | blocklist | hostname | 625 | 419 | 67.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 36.4K | 65.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 11.2K | 60.9% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 301 | 51.9% |
| WaLLy3K | blocklist | domain | 351 | 171 | 48.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 168 | 45.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 141 | 36.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 590 | 36.0% |
| hufilter | blocklist | hostname | 94 | 31 | 33.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.4K | 32.9% |
| hkamran80_smarttv | blocklist | domain | 294 | 96 | 32.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 110 | 31.9% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 34 | 31.5% |
| HaGeZi Pro | blocklist | domain | 222.5K | 68.3K | 30.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 49.4K | 27.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3.7K | 24.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 29 | 23.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6.2K | 22.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2.6K | 19.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 43.2K | 17.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2.5K | 16.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 13.2K | 14.9% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 10.8K | 14.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 4.3K | 14.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6.0K | 13.1% |
| tranco | allowlist | domain_top | 500 | 34 | 6.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 15 | 2.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 77 | 2.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 144 | 2.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 52 | 1.6% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 3.3K | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 301 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 19 | 1.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 563 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 669 | 0.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 1.3K | 0.8% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 26 | 0.8% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 18 | 0.7% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 67 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 432 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 2 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 21 | 0.3% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 167 | 0.3% |
| kadantiscam | blocklist | domain | 43.8K | 93 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 723 | 0.2% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 444 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 43 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 16 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 448 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 11 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 46 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 736 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17 | 0.1% |
| phishing_army | blocklist | domain | 152.5K | 101 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 222 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 211 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 39 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 993 | Targets: 8 | Unique: 893 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 32 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 37 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 24 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 262 | 7.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 6.5K | 7.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 6 | 5.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| tranco | allowlist | domain_top | 500 | 21 | 4.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 404 | 3.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 434 | 2.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.9K | 2.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 9 | 2.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 10 | 1.7% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 520 | 1.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 409 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 751 | 1.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 2.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 1.7K | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 25 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 86 | 0.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1.0K | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 14 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 50 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 1 | 0.0% |

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
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 35 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 10 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.2K | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 432 | 0.8% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 581 | 0.3% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 434 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 10 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 33 | 0.1% |
| abpvn_hosts | blocklist | adguard | 993 | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 3 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 43 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 65 | 0.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 6 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 580 | Targets: 30 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 432 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 105 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 562 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 434 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 493 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 40 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 33 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 7 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 301 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 209.7K | Targets: 14 | Unique: 209.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 444 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 224.8K | Targets: 24 | Unique: 116.6K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 89.7K | 50.5% |
| hufilter | blocklist | hostname | 94 | 18 | 19.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2.3K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 587 | 3.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 8.6K | 3.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 933 | 1.7% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.3K | 1.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 813 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 30 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 50 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 204 | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 84 | 0.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.2K | 0.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 2 | 0.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 335 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 43 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 59 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 2 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 177.7K | Targets: 67 | Unique: 0 | Conflicts: 42</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | domain_adguard | 580 | 562 | 96.9% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.5K | 94.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 50.7K | 91.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 296 | 84.1% |
| hufilter | blocklist | hostname | 94 | 72 | 76.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.4% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 89.7K | 39.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 78.3K | 35.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5.0K | 27.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 49.4K | 24.3% |
| YousList | blocklist | hostname | 625 | 151 | 24.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 967 | 22.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 55.2K | 22.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 28 | 22.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6.0K | 21.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 21 | 19.4% |
| Adaway | blocklist | hostname | 6.5K | 1.0K | 15.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.1K | 13.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 5.9K | 12.9% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 1.7K | 11.7% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2.3K | 7.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 24 | 7.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 4.9K | 5.6% |
| tranco | allowlist | domain_top | 500 | 28 | 5.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 19 | 4.9% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 547 | 4.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.2K | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 12 | 1.7% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 35 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 28 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 10 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 53 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 119 | 0.5% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 343 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 274 | 0.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 11 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 181 | 0.2% |
| Torrent Trackers | blocklist | domain | 483 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 188 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 29 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 282 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 216 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 275 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 10 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 10 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 6 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 108 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 5 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 8 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 115 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 178.5K | Targets: 25 | Unique: 0 | Conflicts: 196</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardSDNSFilter_exceptions | allowlist | adguard | 199 | 195 | 98.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.5K | 93.3% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 50.7K | 91.0% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 1.2K | 87.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| EasyList | blocklist | adguard | 65.3K | 45.9K | 70.2% |
| Easy Privacy | blocklist | adguard | 55.2K | 28.8K | 52.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 581 | 47.4% |
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 55.2K | 22.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 53 | 14.4% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 2.3K | 7.6% |
| abpvn_hosts | blocklist | adguard | 993 | 24 | 2.4% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 31 | 2.1% |
| AdBlockID | allowlist | adguard | 93 | 1 | 1.1% |
| AdBlockID | blocklist | adguard | 3.7K | 35 | 0.9% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 10 | 0.7% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 39 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 119 | 0.5% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 216 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 10 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 181 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 8 | 0.0% |

</details>

---

### AdGuard Spyware Filter - Mobile

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.3K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 839 | 1.5% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 5 | 1.4% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 1.2K | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 823 | 0.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 71 | 0.2% |
| Easy Privacy | blocklist | adguard | 55.2K | 75 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |

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
| Sinfonietta_Social | blocklist | hostname | 3.2K | 5 | 0.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |

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
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 9 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

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
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 5 | Unique: 61 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |

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
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |

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
| AdGuard DNS filter | blocklist | adguard | 178.5K | 2.0K | 1.1% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 2.1K | 0.8% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 60 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 25 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 1 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.4K | Targets: 55 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.0K | 56.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 43.8% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.7% |
| YousList | blocklist | hostname | 625 | 199 | 31.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 91 | 24.7% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 76 | 21.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 206 | 12.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 12 | 11.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 13 | 10.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 30 | 10.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 52 | 9.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 7.6K | 8.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 4.7K | 8.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.3K | 8.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1.1K | 8.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 7.8% |
| tranco | allowlist | domain_top | 500 | 35 | 7.0% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.7K | 6.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11.2K | 5.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 20 | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.5K | 5.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 7.2K | 3.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 5.0K | 2.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 6.4K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 926 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 14 | 2.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 42 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 32 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 53 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 157 | 0.2% |
| Torrent Trackers | blocklist | domain | 483 | 1 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 141 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 17 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 8 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 43 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 28 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 25 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 9 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 34 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 98 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 49 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 14 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.1K | Targets: 22 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 456 | 8.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 712 | 7.3% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 35 | 5.7% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 20 | 5.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 34 | 5.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 806 | 4.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 1.0K | 4.3% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 2.8K | 4.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 620 | 4.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 1.7K | 3.3% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 411 | 3.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 48 | 3.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 501 | 2.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 102 | 2.3% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 183 | 0.8% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 4 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 23 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 11 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 8 | 0.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 8 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 1 | 0.1% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.1K | Targets: 21 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 22.0K | 986 | 4.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 192 | 4.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 46 | 3.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 569 | 2.7% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 14 | 1.4% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 572 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 2 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 1 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 7 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 91 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 53 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 200 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 16 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 4 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 35 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 18 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 10 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 359 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 81 | 5.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 85 | 3.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 121 | 2.8% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 308 | 1.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 20 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 280 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 43 | 0.4% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 2 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 2 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 218 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 36 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 14 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 19 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 9 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 1 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.5K | Targets: 38 | Unique: 10.5K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 3 | 2.4% |
| WaLLy3K | blocklist | domain | 351 | 4 | 1.1% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 24 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 194 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 22 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 181 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 45 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 79 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 41 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 67 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 50 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 7 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 17 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 49 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 32 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 53 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 17 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 248.5K | Targets: 49 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.2% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 269 | 73.9% |
| phishing_army | blocklist | domain | 152.5K | 110.4K | 72.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 42.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 74.9K | 42.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 81.3K | 33.1% |
| kadantiscam | blocklist | domain | 43.8K | 13.9K | 31.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16.3K | 18.4% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 992 | 7.6% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 2.0K | 5.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 10.2K | 4.6% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 3.1K | 3.7% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 1.2K | 2.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 10.8K | 2.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 6.3K | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 3 | 1.1% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 189 | 0.6% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 194 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 49 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 9 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 15 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 16 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 12 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 235 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 45 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 25 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 108 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 222 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 28 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 13 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 33 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 14 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 55 | Unique: 0 | Conflicts: 71</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.8K | 51.4% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 48 | 13.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.9K | 10.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 166 | 10.1% |
| tranco | allowlist | domain_top | 500 | 29 | 5.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 40 | 5.6% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 514 | 4.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.0K | 3.9% |
| WaLLy3K | blocklist | domain | 351 | 13 | 3.7% |
| YousList | blocklist | hostname | 625 | 22 | 3.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 473 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1.6K | 2.8% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 11 | 2.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 2.2K | 2.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 672 | 2.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 7 | 2.4% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 570 | 1.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 2.0K | 0.9% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 5 | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.9K | 0.8% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.4K | 0.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 967 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 24 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 42 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 14 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 19 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 5 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 12 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 24 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 5 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 4 | 1.5% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 5 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 8 | 0.5% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 59 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 23 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 4 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 55 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 36 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 75 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 11 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 5 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 48 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 51.1K | Targets: 34 | Unique: 0 | Conflicts: 43</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 359 | 280 | 78.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1.7K | 53.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 4 | 2 | 50.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.5K | 48.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.9K | 45.8% |
| Greensnow | blocklist | ipv4 | 4.3K | 2.0K | 44.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 4.0K | 41.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 502 | 36.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 534 | 33.9% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 4.3K | 33.9% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 6.8K | 30.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 7.3K | 30.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 5.4K | 25.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 4.2K | 24.5% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 15.2K | 23.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 200 | 17.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 98 | 16.1% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 95 | 15.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 138 | 14.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 103 | 7.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 156 | 6.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 18 | 5.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 2 | 4.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 8 | 3.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 425 | 2.7% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 9 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 214 | 1.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 30 | 1.1% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 43 | 0.4% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 43 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 11 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 21 | Unique: 220 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 65 | 13.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 61 | 1.3% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 3 | 1.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 6 | 0.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 16 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 4 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 615 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 587 | 96.2% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 600 | 4.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 42 | 2.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 35 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 226 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 443 | 0.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 114 | 0.6% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 2 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 34 | 0.4% |
| Greensnow | blocklist | ipv4 | 4.3K | 13 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 95 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 8 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 13 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 10 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### CF_Torrent_Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 108 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| pexcn Torrent Trackers | blocklist | domain_url | 76 | 74 | 97.4% |
| Torrent Trackers | blocklist | domain | 483 | 107 | 22.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 20 | Unique: 0 | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.7K | 8.1K | 64.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 620 | 19.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 4.1K | 19.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 3.3K | 13.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6.9K | 13.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1.1K | 11.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 547 | 10.7% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 6.9K | 10.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 958 | 5.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 10 | 1.6% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 10 | 1.6% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 275 | 1.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 58 | 1.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 18 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 8 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 72 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 3 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 6 | 0.0% |

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
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 55 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 4 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 84.6K | Targets: 46 | Unique: 36.5K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 4.4K | 9.1% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 24 | 6.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 28.8K | 5.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 5 | 4.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 3.2K | 1.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 4 | 1.5% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 384 | 1.3% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3.1K | 1.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 64 | 1.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 503 | 1.3% |
| phishing_army | blocklist | domain | 152.5K | 1.7K | 1.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 136 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 2.2K | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 256 | 0.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 82 | 0.5% |
| HaGeZi Pro | blocklist | domain | 222.5K | 967 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 74 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 264 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 171 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 32 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 28 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 432 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 5 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 43 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 572 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 13 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 181 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 31 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 32 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 122 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 68 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |

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
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 46 | 11.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 422 | 11.9% |
| Adaway | blocklist | hostname | 6.5K | 404 | 6.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.1K | 5.8% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 20 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.6K | 4.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 9 | 3.1% |
| tranco | allowlist | domain_top | 500 | 11 | 2.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 352 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 9 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.6K | 1.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 7 | 1.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 189 | 1.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 2.5K | 1.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 572 | 1.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 278 | 0.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 1.5K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 992 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 547 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 136 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 24 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 122 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 90 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 133 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 29 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 14 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 34 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 8 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 84 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 15 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 78 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 21 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 36 | 0.0% |

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
| EasyList | blocklist | adguard | 65.3K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 7 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 19 | Unique: 62 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 72 | 7.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 46 | 4.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 502 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 472 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 1 | 0.5% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 87 | 0.4% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 52 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 15 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 4 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 51 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 16 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 1 | 0.0% |

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
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.0K | Targets: 7 | Unique: 371 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 1.4K | 95.6% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 81 | 11.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 92 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 2 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.1K | Targets: 9 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 1.0K | 31.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 890 | 5.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 6 | Unique: 0 | Conflicts: 22</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 79 | 5.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 81 | 4.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 22 | 0.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 22 | 0.2% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 17.5K | Targets: 10 | Unique: 14.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3.0K | 90.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 47 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 36 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 3 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 43 | 0.0% |

</details>

---

### DoH_VPN_Proxy_Bypass

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.5K | Targets: 38 | Unique: 13.1K | Conflicts: 13</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3.0K | 90.8% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 890 | 77.8% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| tranco | allowlist | domain_top | 500 | 5 | 1.0% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 5 | 0.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 47 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 162 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 6 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 16 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 8 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 10 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 36 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 77 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 43 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 20 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 5.1K | 30.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 5.1K | 21.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 456 | 14.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 778 | 8.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2.5K | 4.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 882 | 4.3% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 2.7K | 4.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 547 | 3.6% |
| Greensnow | blocklist | ipv4 | 4.3K | 113 | 2.6% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 8 | 2.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 16 | 1.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 15 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 15 | 1.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 2 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 2 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 36 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 1 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 6 | 0.0% |

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
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.0K | 74.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 164 | 44.6% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 28.8K | 16.1% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 75 | 5.6% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 2.6K | 4.7% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 43 | 3.5% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 5.6K | 2.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 336 | 1.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 3 | 0.2% |
| abpvn_hosts | blocklist | adguard | 993 | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 10 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 2 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 8 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 3 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 2 | 0.0% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 65.3K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 32.8K | 58.7% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 45.9K | 25.7% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 32.9K | 13.4% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 65 | 5.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 51 | 3.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 648 | 2.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 2 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 69 | 0.3% |
| abpvn_hosts | blocklist | adguard | 993 | 2 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 148 | 0.1% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 677 | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 11 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 47 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 8 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 1 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 3 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 610 | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 615 | 587 | 95.4% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 577 | 4.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 41 | 2.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 34 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 210 | 1.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 443 | 0.7% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 2 | 0.6% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 108 | 0.5% |
| Greensnow | blocklist | ipv4 | 4.3K | 12 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 28 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 98 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 7 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 12 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 11 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.3% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.5% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5 | Targets: 1 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 5 | 0.0% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.7K | Targets: 39 | Unique: 1.2K | Conflicts: 206</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_social_allowlist | allowlist | domain | 1 | 1 | 100.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 14 | 32.6% |
| local_source_domain_allowlist | allowlist | domain | 43 | 14 | 32.6% |
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
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 6 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 42 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 25 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 11 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 6 | Unique: 0 | Conflicts: 16</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 25 | 1.7% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 25 | 1.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
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
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 15 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 25 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 36 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 21 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 51 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 23 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 206 | Targets: 11 | Unique: 155 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 973 | 24 | 2.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 1 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 2 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 8 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 8 | Unique: 476 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 9 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 2 | 0.0% |

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
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 3 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.3K | Targets: 22 | Unique: 884 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 66 | 2.4% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 15 | 1.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 2 | 1.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 7 | 0.6% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 4 | 0.3% |
| Greensnow | blocklist | ipv4 | 4.3K | 11 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 103 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 15 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 16 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 18 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 33 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 16 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 12 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 73 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 5 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.7K | Targets: 2 | Unique: 1.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.7K | 1.6K | 91.3% |
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.6K | 90.7% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 17.1K | Targets: 28 | Unique: 0 | Conflicts: 541</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 3.8K | 87.7% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 986 | 86.9% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 308 | 85.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1.1K | 66.6% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 226 | 36.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 210 | 34.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 8.2K | 34.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 6.4K | 31.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 806 | 25.9% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 13.3K | 20.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1.8K | 18.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6.8K | 13.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 176 | 6.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 958 | 6.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 87 | 6.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 541 | 4.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 541 | 4.7% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 479 | 3.8% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 26 | 2.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 6 | 1.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 4 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 18 | 1.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 197 | 1.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 7 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 28 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 4 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 6 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.7K | Targets: 28 | Unique: 0 | Conflicts: 37</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 45 | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 600 | 97.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 577 | 94.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8.1K | 54.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 8.2K | 48.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 1.0K | 33.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 7.3K | 14.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1.4K | 14.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 2.8K | 13.4% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 6.1K | 9.3% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 19 | 5.3% |
| local_source_ipv4_allowlist | allowlist | ipv4 | 64 | 3 | 4.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 200 | 4.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 61 | 3.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 35 | 3.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 33 | 2.5% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 479 | 2.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 104 | 0.7% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 5 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 34 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 40 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 34 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 3 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 5 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.7K | Targets: 13 | Unique: 2.5K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 4 | 2 | 50.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 55 | 17.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 66 | 5.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 1 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 5 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 30 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 19 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 7 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 21 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 324 | Targets: 8 | Unique: 232 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 68 | 2 | 2.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 55 | 2.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 1 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 7 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 18 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 6 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 2 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 14.8K | Targets: 18 | Unique: 5.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 316 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 86 | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.5K | 1.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1.7K | 1.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2.3K | 1.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 1.9K | 0.8% |
| YousList | blocklist | hostname | 625 | 4 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 342 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 16 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 53 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 127 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 14 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 57 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 21 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 106 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 13 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.6K | Targets: 20 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 352 | 290 | 82.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 396 | 11.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 166 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 495 | 1.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 206 | 1.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1.5K | 0.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 1.6K | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.6K | 0.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 402 | 0.5% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 590 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 188 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 27 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 67 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1.5K | 55.8% |
| Easy Privacy | blocklist | adguard | 55.2K | 1.5K | 2.8% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 1.5K | 0.9% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 1.6K | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 67 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 27 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 1 | 0.0% |

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
| HaGeZi Pro | blocklist | domain | 222.5K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 5 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.3K | Targets: 26 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 359 | 121 | 33.7% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 3.8K | 17.4% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 192 | 16.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 267 | 16.9% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 3.8K | 5.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 536 | 5.5% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 831 | 4.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2.0K | 3.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 102 | 3.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 414 | 2.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 61 | 2.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 113 | 2.2% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 13 | 2.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 12 | 2.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 15 | 1.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 2 | 1.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 200 | 0.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 11 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 58 | 0.5% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 5 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 58 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 54 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 5 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 369 | Targets: 19 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 4 | 1.4% |
| YousList | blocklist | hostname | 625 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 91 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 20 | 0.3% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 338 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 6 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 49 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 168 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 10 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 20 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 6 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 108 | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 4 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 66 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 21 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 23 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 21 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 177.4K | Targets: 48 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 299 | 82.1% |
| phishing_army | blocklist | domain | 152.5K | 80.0K | 52.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 80.9K | 33.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 15.6K | 32.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 74.9K | 30.1% |
| kadantiscam | blocklist | domain | 43.8K | 11.0K | 25.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 45.2K | 20.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 7.2K | 18.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 11.3K | 12.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 32 | 12.2% |
| quidsup_notrack-malware | blocklist | domain | 125 | 11 | 8.8% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.7K | 5.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 975 | 5.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 23 | 4.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 202 | 4.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 3.2K | 3.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 243 | 3.3% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 16.2K | 3.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 860 | 1.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 48 | 1.1% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 6 | 1.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 56 | 0.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1.3K | 0.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 78 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 162 | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 7 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 50 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 203 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 36 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 49 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 48 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 5 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 43 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 7 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 216 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 9 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 258 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 44 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 12 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 36 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 5 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 105 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 16 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 177.4K | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 80.6K | 51.2% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 80.9K | 33.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 986 | 24.2% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 80.2K | 13.5% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1.7K | 5.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 7 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 6 | 0.5% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 203 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 43 | 0.2% |
| EasyList | blocklist | adguard | 65.3K | 148 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 5 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 48 | 0.2% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 216 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 2 | 0.1% |
| Easy Privacy | blocklist | adguard | 55.2K | 2 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.3K | Targets: 11 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1.0K | 90.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 3.0K | 17.1% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 6 | 0.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 9 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 62 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 26 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.3K | Targets: 6 | Unique: 285 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 3.0K | 17.1% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 5 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 2 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 424.0K | Targets: 41 | Unique: 411.4K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1.2K | 44.7% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2.3K | 7.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.9K | 4.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 3 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 15 | 0.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.0K | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| HaGeZi Pro | blocklist | domain | 222.5K | 750 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 12 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.4K | 0.3% |
| kadantiscam | blocklist | domain | 43.8K | 99 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 34 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 21 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 7 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 38 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 115 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 211 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 9 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 78 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 122 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 235 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 258 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 43 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 621 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 28 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 22 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 19 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 10 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 4 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 11 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 387 | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dan Pollock's List | blocklist | hostname | 13.0K | 46 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 11 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 10 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 337 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 34 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 24 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 66 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 20 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 141 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 10 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 24 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 54 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 19 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 222.5K | Targets: 68 | Unique: 0 | Conflicts: 39</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 98.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 341 | 98.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 53.3K | 95.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 338 | 91.6% |
| hufilter | blocklist | hostname | 94 | 82 | 87.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 337 | 87.1% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 493 | 85.0% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 299 | 84.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.0K | 83.8% |
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| quidsup_notrack-malware | blocklist | domain | 125 | 80 | 64.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 66 | 61.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.0K | 47.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 78.3K | 44.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.2K | 39.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 84.0K | 34.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 68.3K | 33.7% |
| YousList | blocklist | hostname | 625 | 201 | 32.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 8.6K | 30.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 12.1K | 26.3% |
| Adaway | blocklist | hostname | 6.5K | 1.7K | 26.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 45.2K | 25.5% |
| WaLLy3K | blocklist | domain | 351 | 83 | 23.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2.5K | 16.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 13.7K | 15.4% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 4.4K | 14.4% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 47 | 12.9% |
| kadantiscam | blocklist | domain | 43.8K | 5.6K | 12.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 1.9K | 12.8% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1.5K | 11.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 6.8K | 9.0% |
| phishing_army | blocklist | domain | 152.5K | 10.8K | 7.1% |
| tranco | allowlist | domain_top | 500 | 30 | 6.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 29 | 5.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 14 | 5.3% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 10.2K | 4.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 8.6K | 3.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 240 | 3.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 135 | 2.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 416 | 2.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 905 | 2.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 29 | 2.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 62 | 1.9% |
| Spam404 | blocklist | domain | 8.1K | 149 | 1.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 65 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 181 | 1.6% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 47 | 1.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 55 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 8 | 1.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 243 | 1.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 967 | 1.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 26 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 536 | 0.9% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 434 | 0.9% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 162 | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 627 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 49 | 0.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 51 | 0.6% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 6 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.6K | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 2.0K | 0.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 467 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 750 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 42 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 345 | Targets: 14 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 341 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 110 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 21 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 21 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 24 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 87 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.4K | Targets: 7 | Unique: 0 | Conflicts: 32</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 1.4K | 68.5% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 79 | 10.8% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 32 | 0.3% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 32 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 95 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 2 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 65.7K | Targets: 33 | Unique: 0 | Conflicts: 480</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | ipv4 | 5 | 5 | 100.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 13.1K | 97.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 2.8K | 88.9% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 18.4K | 88.6% |
| Greensnow | blocklist | ipv4 | 4.3K | 3.8K | 87.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 12.5K | 73.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 443 | 72.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 1.1K | 72.0% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 443 | 72.0% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 218 | 60.7% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 13.3K | 60.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.7K | 52.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 4.9K | 50.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 572 | 50.4% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 5.9K | 46.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.9K | 46.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 472 | 33.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 15.2K | 29.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 6.1K | 25.5% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 7 | 15.6% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 121 | 12.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 73 | 5.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 128 | 5.0% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 480 | 4.2% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 480 | 4.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 6 | 2.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 368 | 2.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 7 | 2.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 21 | 0.8% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 16 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 15 | 0.0% |

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
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 23 | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 3 | 0.4% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 30 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 45 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 53 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 20 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 15 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 108 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 108 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 96 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 21 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 13 | 0.0% |

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
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 87 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 82 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 12 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 31 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 18 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 72 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 85 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 833 | Targets: 3 | Unique: 827 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 468.7K | Targets: 55 | Unique: 431.7K | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 3.6K | 32.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 959 | 13.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 12 | 9.6% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 10.8K | 4.4% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1.7K | 4.4% |
| hufilter | blocklist | hostname | 94 | 4 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 9 | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1.0K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 117 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 75 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 4 | 1.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 2.4K | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 3.2K | 1.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 2.6K | 1.2% |
| YousList | blocklist | hostname | 625 | 7 | 1.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 26 | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 572 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 2.4K | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 98 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 723 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 99 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 34 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1.4K | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 186 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 34 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 189 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 52 | 0.2% |
| phishing_army | blocklist | domain | 152.5K | 278 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 275 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 30 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 20 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 25 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 208 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 45 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 62 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 47 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 17 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 43.8K | Targets: 40 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 37.3K | 42.0% |
| phishing_army | blocklist | domain | 152.5K | 18.4K | 12.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 11.0K | 6.2% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 13.9K | 5.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 11.5K | 4.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 4 | 3.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 5.6K | 2.5% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 706 | 1.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 2 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 22 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| YousList | blocklist | hostname | 625 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 29 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 43 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 8 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 69 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 43 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 19 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 22 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 99 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 24 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 93 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 45 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 54 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 13 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 29 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |

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
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 4 | 1.5% |
| tranco | allowlist | domain_top | 500 | 6 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 7 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 10 | 0.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 46 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 3 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 3 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 5 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 26 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 29 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 5 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 6 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 23 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 4 | 0.0% |

</details>

---

### local_adg_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 4 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 2 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |

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
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |

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
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 5 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 5 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 3 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 6 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 5 | 0.0% |

</details>

---

### local_miscellaneous_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 9 | Unique: 0 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |

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
| Firehol_level3 | blocklist | ipv4 | 12.7K | 3 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 4.1K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 470 | 1.5% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 4.1K | 0.7% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 986 | 0.6% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 1.3K | 0.5% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 2 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 10 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 39.1K | Targets: 31 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 87 | 33.1% |
| phishing_army | blocklist | domain | 152.5K | 19.9K | 13.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 28.0K | 11.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 7.2K | 4.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 16 | 3.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 544 | 1.8% |
| kadantiscam | blocklist | domain | 43.8K | 706 | 1.6% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 2.0K | 0.8% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 503 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 23 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.7K | 0.4% |
| HaGeZi Pro | blocklist | domain | 222.5K | 905 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 20 | 0.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 121 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 45 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 8 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 3 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 130 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 46 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 12 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 245.4K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 55.1K | 98.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1.6K | 98.6% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 2.1K | 75.3% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 823 | 61.9% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 81.0K | 51.5% |
| EasyList | blocklist | adguard | 65.3K | 32.9K | 50.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 80.9K | 45.6% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 434 | 35.4% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 1.3K | 31.8% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 55.2K | 30.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 65 | 17.7% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 83.7K | 14.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 4.0K | 13.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 5.6K | 10.2% |
| abpvn_hosts | blocklist | adguard | 993 | 37 | 3.7% |
| CJX Annoyance | blocklist | adguard | 1.8K | 55 | 3.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 38 | 2.6% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 12 | 0.9% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 68 | 0.9% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 113 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 9 | 0.3% |
| AdBlockID | blocklist | adguard | 3.7K | 10 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 36 | 0.2% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 7 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 833 | 1 | 0.1% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 245.4K | Targets: 64 | Unique: 0 | Conflicts: 24</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1.6K | 99.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 55.1K | 98.7% |
| hufilter | blocklist | hostname | 94 | 85 | 90.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 310 | 85.2% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 298 | 84.7% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 434 | 74.8% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 28.0K | 71.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.4K | 67.3% |
| phishing_army | blocklist | domain | 152.5K | 80.6K | 52.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 80.9K | 45.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 44.1% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| YousList | blocklist | hostname | 625 | 266 | 42.6% |
| WaLLy3K | blocklist | domain | 351 | 138 | 39.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 84.0K | 37.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 47 | 37.6% |
| hkamran80_smarttv | blocklist | domain | 294 | 108 | 36.7% |
| Adaway | blocklist | hostname | 6.5K | 2.4K | 36.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 6.4K | 34.9% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 81.3K | 32.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 55.2K | 31.0% |
| kadantiscam | blocklist | domain | 43.8K | 11.5K | 26.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 87 | 25.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 11.4K | 24.7% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 21.5K | 24.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.3K | 23.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 23 | 21.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 43.2K | 21.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 15.9K | 21.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 5.8K | 21.0% |
| Spam404 | blocklist | domain | 8.1K | 1.6K | 19.4% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2.5K | 19.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 54 | 14.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 4.0K | 13.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 550 | 11.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 37 | 10.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.4K | 9.0% |
| tranco | allowlist | domain_top | 500 | 16 | 3.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 7 | 2.7% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 2.2K | 2.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 194 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 8 | 1.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 71 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 31 | 1.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 38 | 1.0% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 12 | 0.9% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 106 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3.2K | 0.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 451 | 0.6% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 270 | 0.6% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 385 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 113 | 0.5% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 2.4K | 0.5% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 12 | 0.5% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1.2K | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 9 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 36 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1.0K | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 220 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 45 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 7 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 22.1K | Targets: 12 | Unique: 21.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 32 | 2.3% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 64 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 78 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 119 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 69 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 48 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 113 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 15 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 3 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 1 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 22.1K | Targets: 41 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 6.5K | 10.6% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 7.0K | 9.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 14.9K | 6.7% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 32 | 2.3% |
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 26 | 0.4% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 49 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 64 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| Torrent Trackers | blocklist | domain | 483 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 11 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 119 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 78 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 74 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 17 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 301 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 243 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 28 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 48 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 7 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 26 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 38 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 113 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 25 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 42 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 55.8K | Targets: 62 | Unique: 0 | Conflicts: 20</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 94 | 87 | 92.6% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 432 | 74.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1.6K | 45.7% |
| local_domain_blocklist | blocklist | domain | 7 | 3 | 42.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 36.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 50.7K | 28.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4.7K | 25.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 53.3K | 23.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 55.1K | 22.4% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 21 | 19.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 36.4K | 17.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 22 | 17.6% |
| YousList | blocklist | hostname | 625 | 94 | 15.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 6.2K | 13.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 3.3K | 12.0% |
| Adaway | blocklist | hostname | 6.5K | 751 | 11.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 25 | 7.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 24 | 6.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 924 | 6.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 21 | 6.1% |
| WaLLy3K | blocklist | domain | 351 | 20 | 5.7% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.4K | 4.7% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 572 | 4.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 13 | 4.4% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 67 | 4.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.6K | 4.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3.1K | 4.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| tranco | allowlist | domain_top | 500 | 14 | 2.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 6 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 29 | 0.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 79 | 0.7% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 57 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 933 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 78 | 0.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 260 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 214 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 171 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 8 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 24 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 10 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 203 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 8 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 186 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 174 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 45 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 78 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 70 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 55.8K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 839 | 63.1% |
| EasyList | blocklist | adguard | 65.3K | 32.8K | 50.2% |
| local_adg_blocklist | blocklist | adguard | 7 | 3 | 42.9% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 432 | 35.3% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 50.7K | 28.4% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 55.1K | 22.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 43 | 11.7% |
| Easy Privacy | blocklist | adguard | 55.2K | 2.6K | 4.8% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1.4K | 4.7% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 67 | 4.1% |
| abpvn_hosts | blocklist | adguard | 993 | 32 | 3.2% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 60 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 30 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 78 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 3 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 203 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 5 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 70 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 8 | 0.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 2 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 263 | Targets: 19 | Unique: 14 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 87 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 22 | 0.1% |
| phishing_army | blocklist | domain | 152.5K | 55 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 7 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 32 | 0.0% |

</details>

---

### pexcn Torrent Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 108 | 74 | 68.5% |
| Torrent Trackers | blocklist | domain | 483 | 75 | 15.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 30.6K | Targets: 24 | Unique: 17.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_adg_blocklist | blocklist | adguard | 7 | 2 | 28.6% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 470 | 11.6% |
| AdGuard Spyware Filter - Mobile | blocklist | adguard | 1.3K | 71 | 5.3% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 33 | 2.7% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 1.4K | 2.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 7 | 1.9% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 4.0K | 1.6% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 27 | 1.6% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 2.3K | 1.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 1.7K | 1.0% |
| EasyList | blocklist | adguard | 65.3K | 648 | 1.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 25 | 0.9% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 11 | 0.7% |
| Easy Privacy | blocklist | adguard | 55.2K | 336 | 0.6% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 646 | 0.4% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 64 | 0.3% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 47 | 0.3% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 1.5K | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 3.3K | 2 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 1 | 0.1% |
| abpvn_hosts | blocklist | adguard | 993 | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.7K | 1 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 2 | 0.0% |

</details>

---

### ph00lt0_blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 30.6K | Targets: 76 | Unique: 0 | Conflicts: 162</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 688 | 19.3% |
| tranco | allowlist | domain_top | 500 | 79 | 15.8% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 51 | 14.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 570 | 13.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 46 | 9.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 22 | 8.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.3K | 8.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.5K | 8.3% |
| Adaway | blocklist | hostname | 6.5K | 520 | 8.0% |
| quidsup_notrack-malware | blocklist | domain | 125 | 10 | 8.0% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.9K | 6.8% |
| YousList | blocklist | hostname | 625 | 42 | 6.7% |
| hufilter | blocklist | hostname | 94 | 6 | 6.4% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 33 | 5.7% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 15 | 5.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 127 | 4.8% |
| local_ai_blocklist | blocklist | domain | 24 | 1 | 4.2% |
| local_ai_allowlist | allowlist | domain | 24 | 1 | 4.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 26 | 3.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 10 | 2.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1.4K | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 42 | 2.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 9 | 2.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 278 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 4.3K | 2.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 4.4K | 2.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.6K | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.4K | 1.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 4.0K | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 6 | 1.6% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 27 | 1.6% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 544 | 1.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2.3K | 1.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 59 | 1.3% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 525 | 1.1% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 181 | 2 | 1.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 1.7K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 48 | 0.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 3.8K | 0.7% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 384 | 0.5% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 2.3K | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 14 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 185 | 0.4% |
| phishing_army | blocklist | domain | 152.5K | 612 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 17 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 47 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 64 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 9 | 0.2% |
| Torrent Trackers | blocklist | domain | 483 | 1 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 22 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 1 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 43 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 7 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 92 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 21 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 75 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 5 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 189 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 59 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 99 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 75 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 9 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 152.5K | Targets: 36 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 19.9K | 50.8% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 80.0K | 45.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 110.4K | 44.4% |
| kadantiscam | blocklist | domain | 43.8K | 18.4K | 42.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 80.6K | 32.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 55 | 20.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 17.0K | 19.1% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 26 | 5.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 10.8K | 4.8% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1.7K | 2.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 612 | 2.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 36 | 0.8% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 30 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 278 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 32 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 10 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 651 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 8 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 11 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 101 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |

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
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 95 | 6.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 92 | 4.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 19 | 0.7% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 31 | 0.3% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 31 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 1 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 11 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 5 | 0.0% |
| Greensnow | blocklist | ipv4 | 4.3K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 15 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 6 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 1 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 352 | Targets: 19 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 290 | 17.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 142 | 4.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 48 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 76 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 145 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 46 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 46 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 296 | 0.2% |
| tranco | allowlist | domain_top | 500 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 298 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 241 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 67 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 299 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 25 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 1 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 125 | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 9 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 59 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 13 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 80 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 20 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 28 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 12 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 47 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 29 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 8 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 11 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 10 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 14 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 22 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.2K | Targets: 55 | Unique: 0 | Conflicts: 52</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 577 | 16.2% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 473 | 11.1% |
| WaLLy3K | blocklist | domain | 351 | 35 | 10.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 34 | 8.8% |
| tranco | allowlist | domain_top | 500 | 38 | 7.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 8 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 294 | 21 | 7.1% |
| Adaway | blocklist | hostname | 6.5K | 434 | 6.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.2K | 6.4% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.3K | 4.1% |
| YousList | blocklist | hostname | 625 | 21 | 3.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 904 | 3.3% |
| hufilter | blocklist | hostname | 94 | 3 | 3.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 10 | 2.7% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 316 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 11 | 1.9% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 3.7K | 1.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 6 | 1.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 924 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 189 | 1.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.2K | 1.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 996 | 1.3% |
| pexcn Torrent Trackers | blocklist | domain_url | 76 | 1 | 1.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2.1K | 1.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 2.5K | 1.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 108 | 1 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 25 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 26 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.4K | 0.6% |
| Torrent Trackers | blocklist | domain | 483 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 183 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 587 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 5 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 30 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 11 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 12 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 26 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 23 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 677 | Targets: 1 | Unique: 676 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 65.3K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 593.7K | Targets: 15 | Unique: 310.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 4.1K | 99.7% |
| RPiList_specials-phishing | blocklist | adguard | 157.3K | 113.2K | 72.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 80.2K | 45.2% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 83.7K | 34.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1.5K | 5.0% |
| AdGuard Base filter | blocklist | adguard | 1.2K | 4 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard_domain | 1.4K | 3 | 0.2% |
| DoH_VPN_Proxy_Bypass | blocklist | adguard | 17.5K | 11 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.5K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 181 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 47 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 15 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 70 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 14.0K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 3 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 157.3K | Targets: 8 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 80.6K | 45.4% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 81.0K | 33.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 113.2K | 19.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 646 | 2.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 4.1K | 10 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 8 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 1 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.6K | Targets: 20 | Unique: 0 | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 359 | 81 | 22.6% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 42 | 6.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 41 | 6.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 267 | 6.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 1.1K | 4.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 184 | 1.9% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 1.1K | 1.7% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 310 | 1.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 48 | 1.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 534 | 1.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 61 | 0.5% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 53 | 0.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 13 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 53 | 0.5% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 15 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 54 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 19 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 35 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 18 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 973 | Targets: 19 | Unique: 504 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 24 | 11.7% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 72 | 5.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 14 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 15 | 1.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 1 | 0.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 138 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 5 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 121 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 12 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 26 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 13 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 8 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 8 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 1 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 17 | Unique: 1.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 359 | 85 | 23.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 61 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 176 | 0.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 13 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 156 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 1 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 128 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 1 | 0.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 19 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 4 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 6 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 6 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 7 | 0.0% |

</details>

---

### Sefinek_Known_Bots_IP

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 11.4K | Targets: 21 | Unique: 0 | Conflicts: 12.9K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 11.4K | 100.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 11.4K | 100.0% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 16 | 24.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 53 | 3.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 22 | 3.0% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 541 | 2.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 32 | 2.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 32 | 1.6% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 480 | 0.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 64 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 8 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 34 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 27 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 49 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 2 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 31 | 0.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 5 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 43 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 31 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 9.7K | Targets: 26 | Unique: 0 | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 712 | 22.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 778 | 15.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 536 | 12.3% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 43 | 12.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 184 | 11.7% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 1.8K | 8.2% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 1.6K | 7.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 4.0K | 7.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.5% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 4.9K | 7.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 1.2K | 7.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 802 | 6.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 1.4K | 5.8% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 34 | 5.5% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 53 | 4.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 28 | 4.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 220 | 1.6% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 12 | 1.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 12 | 0.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 76 | 0.5% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 27 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 27 | 0.2% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |

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
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 32 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 16 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 14 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 10 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 19 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 29 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 12 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |

</details>

---

### ShadowWhisperer's Dating List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 22.1K | 32 | 0.1% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 1 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 10 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 12 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 712 | Targets: 39 | Unique: 334 | Conflicts: 305</summary>

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
| Sinfonietta_Social | blocklist | hostname | 3.2K | 17 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 2 | 0.5% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 20 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 9 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 11 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 26 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 14 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 27 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 15 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 6 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 31 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 8 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 27.8K | Targets: 53 | Unique: 0 | Conflicts: 23</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 495 | 30.2% |
| local_domain_blocklist | blocklist | domain | 7 | 2 | 28.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 978 | 27.5% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 105 | 18.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 672 | 15.7% |
| WaLLy3K | blocklist | domain | 351 | 54 | 15.4% |
| YousList | blocklist | hostname | 625 | 86 | 13.8% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 46 | 13.1% |
| hufilter | blocklist | hostname | 94 | 11 | 11.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1.7K | 9.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 20 | 6.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 7 | 6.5% |
| Adaway | blocklist | hostname | 6.5K | 409 | 6.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 24 | 6.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.9K | 6.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3.3K | 6.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 904 | 5.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 7 | 5.6% |
| tranco | allowlist | domain_top | 500 | 23 | 4.6% |
| HaGeZi Pro | blocklist | domain | 222.5K | 8.6K | 3.9% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 6.0K | 3.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 11 | 3.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.2K | 3.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 352 | 2.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 5.8K | 2.4% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 1.8K | 2.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1.9K | 2.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 3 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 3 | 0.6% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 15 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 11 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 123 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 161 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 49 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 162 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 13 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 204 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 11 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 35 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 12 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 52 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 6 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 221.5K | Targets: 33 | Unique: 163.4K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 14.9K | 67.3% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 19.4K | 31.7% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 21.5K | 28.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 467 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 75 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 448 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 50 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 105 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 22 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 106 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 15 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 68 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 70 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 220 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 188 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 28 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 22 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 9 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 4 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 8 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 208 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 139 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 46.0K | Targets: 43 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 59 | 47.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 6.2K | 11.0% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 40 | 6.9% |
| HaGeZi Pro | blocklist | domain | 222.5K | 12.1K | 5.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 926 | 5.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 11.4K | 4.6% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 5.9K | 3.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 6.0K | 3.0% |
| YousList | blocklist | hostname | 625 | 17 | 2.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 55 | 1.5% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 5 | 1.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 183 | 1.2% |
| hufilter | blocklist | hostname | 94 | 1 | 1.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 5 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 90 | 0.7% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 185 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 860 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 41 | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 313 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 45 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 256 | 0.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 1 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1.0K | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 161 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| kadantiscam | blocklist | domain | 43.8K | 69 | 0.2% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 50 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 16 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 28 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 368 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 194 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 44 | 0.1% |
| phishing_army | blocklist | domain | 152.5K | 12 | 0.0% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 12 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 38 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 7.3K | Targets: 29 | Unique: 4.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 899 | 1.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 38 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 959 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 26 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 144 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 243 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 240 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 10 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 12 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 12 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 13 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 6 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 4 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 7 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 71 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 2 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 12 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.0K | Targets: 25 | Unique: 1.3K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4.1K | 90.0% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 126 | 25.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 2 | 0.8% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 48 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 20 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 117 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 8 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 23 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 56 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 5 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 30 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 9 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 49 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 61.2K | Targets: 43 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Porn | blocklist | hostname | 76.8K | 61.1K | 79.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 6.5K | 29.4% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 19.4K | 8.8% |
| pexcn Torrent Trackers | blocklist | domain_url | 76 | 2 | 2.6% |
| Torrent Trackers | blocklist | domain | 483 | 9 | 1.9% |
| CF_Torrent_Trackers | blocklist | domain_url | 108 | 2 | 1.9% |
| YousList | blocklist | hostname | 625 | 11 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 65 | 1.8% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 876 | 1.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 14 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 122 | 0.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 141 | 0.8% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 615 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 123 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 214 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 563 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 222.5K | 536 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 23 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 274 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 75 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 385 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 44 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 1 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 9 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 13 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 39 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 47 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 36 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 14 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 11 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.6K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 2.6K | 3.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 127 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1.2K | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 12 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 26 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 18 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 13 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 23 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 25 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 32 | 0.2% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 38 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 47 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 52 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 11 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 31 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 28 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 23 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 29 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 14 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 30 | Unique: 6.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 1.6K | 0.6% |
| WaLLy3K | blocklist | domain | 351 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 20 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| kadantiscam | blocklist | domain | 43.8K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 41 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 52 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 149 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 11 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 4 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 11 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 6 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 7 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 12 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 5 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |

</details>

---

### spamhaus_drop

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.7K | Targets: 2 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.7K | 1.7K | 99.0% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.7K | 1.6K | 33.6% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 17.6K | 7.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 4.3K | 1.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 975 | 0.5% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 2 | 0.5% |
| HaGeZi Pro | blocklist | domain | 222.5K | 416 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 389 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 82 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 25 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 115 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 19 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 2 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 17 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 7 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 88.9K | Targets: 68 | Unique: 0 | Conflicts: 76</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2.6K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3.5K | 99.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 12.9K | 99.3% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 334 | 91.8% |
| local_domain_blocklist | blocklist | domain | 7 | 6 | 85.7% |
| kadantiscam | blocklist | domain | 43.8K | 37.3K | 85.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 50.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 7.6K | 41.6% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 145 | 41.2% |
| YousList | blocklist | hostname | 625 | 241 | 38.6% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 20.7K | 27.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 402 | 24.5% |
| WaLLy3K | blocklist | domain | 351 | 85 | 24.2% |
| hkamran80_smarttv | blocklist | domain | 294 | 53 | 18.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 66 | 17.1% |
| quidsup_notrack-malware | blocklist | domain | 125 | 20 | 16.0% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 49 | 13.3% |
| hufilter | blocklist | hostname | 94 | 12 | 12.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 899 | 12.3% |
| phishing_army | blocklist | domain | 152.5K | 17.0K | 11.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 21.5K | 8.7% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1.2K | 8.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.9K | 7.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 16.3K | 6.6% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 13.2K | 6.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3.6K | 6.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 11.3K | 6.4% |
| tranco | allowlist | domain_top | 500 | 32 | 6.4% |
| HaGeZi Pro | blocklist | domain | 222.5K | 13.7K | 6.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 32 | 5.5% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.6K | 5.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 16 | 4.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 31 | 4.4% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 4.9K | 2.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 342 | 2.3% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 6 | 1.2% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 615 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 3.9K | 0.9% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 703 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 29 | 0.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 2 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 32 | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 313 | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 11 | 0.7% |
| Spam404 | blocklist | domain | 8.1K | 52 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 16 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 41 | 0.4% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 813 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 121 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 264 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 42 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 3 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 16 | 0.2% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 424 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 25 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 8 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 106 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 189 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 10 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |

</details>

---

### StevenBlack_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 76.8K | Targets: 45 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 61.1K | 100.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 7.0K | 31.9% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 21.5K | 9.7% |
| pexcn Torrent Trackers | blocklist | domain_url | 76 | 2 | 2.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 73 | 2.1% |
| hufilter | blocklist | hostname | 94 | 2 | 2.1% |
| CF_Torrent_Trackers | blocklist | domain_url | 108 | 2 | 1.9% |
| YousList | blocklist | hostname | 625 | 12 | 1.9% |
| Torrent Trackers | blocklist | domain | 483 | 9 | 1.9% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 958 | 1.3% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 16 | 1.2% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 133 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 157 | 0.9% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 703 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 26 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 161 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 260 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 669 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 627 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 92 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 13 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 16 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 343 | 0.2% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 451 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 50 | 0.1% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 2 | 0.1% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 2 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 13 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 50 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 14 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 32 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 10 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 44 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 48 | 0.0% |

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
| Adaway | blocklist | hostname | 6.5K | 25 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 14 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 26 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 42 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 29 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 15 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 17 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 41 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 77 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 65 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 38 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 35 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |

</details>

---

### ThreatFox_Hostfile

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 48.7K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 15.6K | 8.8% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 30.6K | 5.8% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 4.4K | 5.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 525 | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 125 | 2 | 1.6% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 6 | 1.6% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 1.2K | 0.5% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 2 | 0.3% |
| HaGeZi Pro | blocklist | domain | 222.5K | 434 | 0.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 12 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 167 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 270 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 45 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 6 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 1 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 32 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 19 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 17 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 16 | 0.0% |

</details>

---

### ThreatView_Domain_High-Confidence

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 526.2K | Targets: 47 | Unique: 428.5K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 30.6K | 62.9% |
| URLHaus (Abuse.ch) | blocklist | hostname | 364 | 213 | 58.5% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 28.8K | 34.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 3.8K | 12.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 16.2K | 9.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 442 | 8.9% |
| quidsup_notrack-malware | blocklist | domain | 125 | 8 | 6.4% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 6.3K | 2.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 389 | 2.1% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 6 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 2.4K | 1.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 2.0K | 0.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 368 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 4 | 0.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 2.4K | 0.5% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 424 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 736 | 0.4% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 77 | 0.4% |
| phishing_army | blocklist | domain | 152.5K | 651 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 23 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 174 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 23 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 15 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 130 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 36 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 6 | 0.2% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 282 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 12 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 38 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 35 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 12 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 39 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 48 | 0.1% |
| kadantiscam | blocklist | domain | 43.8K | 54 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 139 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 621 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 21 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 1 | 0.0% |

</details>

---

### ThreatView_IP_HighConfidence

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20.7K | Targets: 26 | Unique: 0 | Conflicts: 49</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 569 | 50.2% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 6.4K | 29.2% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 18.4K | 27.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4.1K | 27.1% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 2.8K | 21.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 310 | 19.7% |
| Greensnow | blocklist | ipv4 | 4.3K | 831 | 19.1% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 114 | 18.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 108 | 17.7% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 882 | 17.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 1.6K | 16.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 501 | 16.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 5.4K | 10.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 2.5K | 10.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 1.4K | 8.1% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 14 | 3.9% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 52 | 3.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 212 | 1.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 15 | 1.2% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 8 | 0.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 19 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 324 | 2 | 0.6% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 49 | 0.4% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 49 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 66 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 4 | 0.1% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 483 | Targets: 9 | Unique: 277 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CF_Torrent_Trackers | blocklist | domain_url | 108 | 107 | 99.1% |
| pexcn Torrent Trackers | blocklist | domain_url | 76 | 75 | 98.7% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 9 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 2 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 9 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 500 | Targets: 47 | Unique: 0 | Conflicts: 563</summary>

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 31 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 27 | 0.8% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 27 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 29 | 0.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 263 | 1 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 79 | 0.3% |
| YousList | blocklist | hostname | 625 | 2 | 0.3% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 1 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 35 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 38 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 23 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 3.3K | 3 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.1K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 11 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 6 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 30 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 28 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 16 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.2K | 1 | 0.0% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 36 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 14 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 32 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 34 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.5K | Targets: 8 | Unique: 1.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 55.8K | 30 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 51 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 177.4K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 38 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 11 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 593.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 31 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 3 | 0.0% |

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
| OISD Blocklist Small | blocklist | adguard | 55.8K | 43 | 0.1% |
| Easy Privacy | allowlist | adguard | 839 | 1 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.7K | 1 | 0.1% |
| EasyList | blocklist | adguard | 65.3K | 2 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 7 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.4K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 53 | 0.0% |
| AntiAdBlockFilters | blocklist | adguard | 2.8K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 65 | 0.0% |

</details>

---

### URLHaus (Abuse.ch)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 364 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 334 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 299 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 51 | 0.2% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 1 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 269 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 310 | 0.1% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 6 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 24 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 47 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 213 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 57.0K | Targets: 1 | Unique: 57.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | adguard_http_url | 101 | 1 | 1.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 13.4K | Targets: 26 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 13.1K | 19.9% |
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 4 | 8.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 220 | 2.3% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 8 | 1.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 7 | 1.1% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 212 | 1.0% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 3 | 0.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 85 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 11 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 214 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 8 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 40 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 28 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 17 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 1 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 16 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 2 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 1 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 1 | 0.1% |
| Greensnow | blocklist | ipv4 | 4.3K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.5K | Targets: 34 | Unique: 13.4K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 5 | 11.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 280 | 5.5% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 51 | 3.7% |
| BlockListDE_Strong | blocklist | ipv4 | 359 | 9 | 2.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 610 | 11 | 1.8% |
| BruteforceBlocker | blocklist | ipv4_find | 615 | 10 | 1.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 1.3K | 16 | 1.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.6K | 19 | 1.2% |
| Greensnow | blocklist | ipv4 | 4.3K | 54 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 197 | 0.9% |
| BlockListDE_Brute | blocklist | ipv4 | 1.1K | 10 | 0.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.7K | 76 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 973 | 8 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 425 | 0.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 36 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.1K | 23 | 0.7% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 72 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 85 | 0.6% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 368 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 206 | 1 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 72 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 104 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 53 | 0.3% |
| ThreatView_IP_HighConfidence | blocklist | ipv4 | 20.7K | 66 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 6 | 0.2% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.0K | 2 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.4K | 2 | 0.1% |
| Sefinek_Known_Bots_IP | allowlist | ipv4 | 11.4K | 4 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.7K | 1 | 0.0% |
| Sefinek_Known_Bots_IP | blocklist | ipv4 | 11.4K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
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
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 2.1K | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 550 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 64 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 135 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 202 | 0.1% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 442 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 3 | 0.0% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 166 | 0.0% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 12 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 11 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 2 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 12 | Unique: 4.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 45 | 3 | 6.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 280 | 1.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 8 | 0.1% |
| DanMeUK_TorExitNodes | blocklist | ipv4 | 1.4K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 22.0K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.1K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.7K | 3 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.9K | 5 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 16 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 101 | Targets: 1 | Unique: 100 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | adguard_http_url | 57.0K | 1 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 45 | Targets: 6 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.7K | 45 | 0.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 5.1K | 3 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 51.1K | 2 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 65.7K | 7 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 13.4K | 4 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 5 | 0.0% |

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
| quidsup_notrack-malware | blocklist | domain | 125 | 1 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 294 | 2 | 0.7% |
| tranco | allowlist | domain_top | 500 | 3 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 83 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 19 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 3 | 0.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 1 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 20 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 54 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 138 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 171 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 1 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 19 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 81 | 0.1% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 85 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 2 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 35 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 83 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 20 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 1 | 0.0% |

</details>

---

### Warui_Adhosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.8K | Targets: 64 | Unique: 0 | Conflicts: 92</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.4K | 97.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 2.6K | 73.9% |
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
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 188 | 11.5% |
| quidsup_notrack-malware | blocklist | domain | 125 | 14 | 11.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 41 | 10.6% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 9 | 8.3% |
| tranco | allowlist | domain_top | 500 | 36 | 7.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 996 | 6.5% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 15.9K | 6.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 1.8K | 6.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 3.1K | 5.5% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 10.8K | 5.3% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 1.4K | 4.5% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 25 | 4.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 27 | 3.8% |
| HaGeZi Pro | blocklist | domain | 222.5K | 6.8K | 3.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 8 | 2.3% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 3.2K | 1.8% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 25 | 1.5% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 876 | 1.4% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 958 | 1.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 38 | 1.2% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 41 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 127 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 161 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 21 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 7 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 75 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 26 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 17 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 335 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain | 1.4K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 4 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 33 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 62 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 9 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 50 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 43 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 21 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 31 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 15 | 0.0% |

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
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 24 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| tranco | allowlist | domain_top | 500 | 2 | 0.4% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 241 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 86 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 294 | 1 | 0.3% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 231 | 0.3% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 2 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 419 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 94 | 0.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 42 | 0.1% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 151 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 21 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 266 | 0.1% |
| HaGeZi Pro | blocklist | domain | 222.5K | 201 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 17 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 4 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 11 | 0.0% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 1 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.4K | Targets: 7 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 368 | 1 | 0.3% |
| OISD Blocklist Big | blocklist | adguard | 245.4K | 68 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 55.8K | 25 | 0.0% |
| ph00lt0_blocklist | blocklist | adguard_domain | 30.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 178.5K | 39 | 0.0% |
| Easy Privacy | blocklist | adguard | 55.2K | 10 | 0.0% |
| EasyList | blocklist | adguard | 65.3K | 11 | 0.0% |

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
| Dan Pollock's List | blocklist | hostname | 13.0K | 50 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 52 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 4 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.6K | 8 | 0.2% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 74 | 0.1% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 75 | 0.1% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 45 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 39 | 0.0% |
| HaGeZi Pro | blocklist | domain | 222.5K | 42 | 0.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 8 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 9 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.6K | Targets: 60 | Unique: 0 | Conflicts: 44</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| local_domain_blocklist | blocklist | domain | 7 | 5 | 71.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.8K | 42.9% |
| quidsup_notrack-annoyance | blocklist | domain | 352 | 142 | 40.3% |
| GetAdmiral Domains Filter List | blocklist | domain_adguard | 1.6K | 396 | 24.1% |
| local_miscellaneous_allowlist | allowlist | domain | 7 | 1 | 14.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.4K | 2.0K | 11.0% |
| hkamran80_smarttv | blocklist | domain | 294 | 23 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 125 | 9 | 7.2% |
| tranco | allowlist | domain_top | 500 | 31 | 6.2% |
| WaLLy3K | blocklist | domain | 351 | 19 | 5.4% |
| hufilter | blocklist | hostname | 94 | 5 | 5.3% |
| StevenBlack_Fake_Gambling | blocklist | hostname | 88.9K | 3.5K | 4.0% |
| Adaway | blocklist | hostname | 6.5K | 262 | 4.0% |
| YousList | blocklist | hostname | 625 | 24 | 3.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.2K | 577 | 3.8% |
| HaGeZi Apple Tracker | blocklist | domain | 108 | 4 | 3.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 27.8K | 978 | 3.5% |
| Warui_Adhosts | blocklist | hostname | 75.8K | 2.6K | 3.5% |
| Dan Pollock's List | blocklist | hostname | 13.0K | 422 | 3.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 55.8K | 1.6K | 2.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 387 | 10 | 2.6% |
| AdGuard Base filter | blocklist | domain_adguard | 580 | 13 | 2.2% |
| ph00lt0_blocklist | blocklist | domain | 30.6K | 688 | 2.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 712 | 11 | 1.5% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 345 | 5 | 1.4% |
| HaGeZi Pro | blocklist | domain | 222.5K | 3.0K | 1.3% |
| 1Hosts (Lite) | blocklist | domain | 203.0K | 2.4K | 1.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 369 | 4 | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 245.4K | 2.4K | 1.0% |
| AdGuard DNS filter | blocklist | domain_adguard | 177.7K | 1.6K | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| StevenBlack_Social | blocklist | hostname | 3.8K | 14 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 499 | 2 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 61.2K | 65 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 1.7K | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.6K | 4 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 14.8K | 16 | 0.1% |
| StevenBlack_Porn | blocklist | hostname | 76.8K | 73 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 46.0K | 55 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 7.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 22.1K | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 424.0K | 15 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 84.6K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 221.5K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 43.8K | 9 | 0.0% |
| ThreatFox_Hostfile | blocklist | hostname | 48.7K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 6.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 11.5K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 39.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 177.4K | 7 | 0.0% |
| DoH_VPN_Proxy_Bypass | blocklist | domain_adguard | 17.5K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 209.7K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 248.5K | 4 | 0.0% |
| ThreatView_Domain_High-Confidence | blocklist | domain | 526.2K | 6 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 224.8K | 30 | 0.0% |
| phishing_army | blocklist | domain | 152.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 468.7K | 26 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.7K | Targets: 1 | Unique: 8.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| USOM-Blocklists-ips | blocklist | ipv4 | 15.5K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

