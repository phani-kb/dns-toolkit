# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2025-08-26 05:28:55 UTC

## How to read this analysis

- Unique Entries (same list type): number of entries found only in this source when compared with other sources of the same list type (blocklist vs blocklist, allowlist vs allowlist). If this is `0` the source is fully covered by other sources of the same list type.
- Conflicts (cross-list overlaps): entries from this source that also appear in sources of a different list type (for example an entry present in a blocklist and an allowlist). Conflicts may indicate data mismatches.
- Overlap % (in the table): shown relative to the target source (overlap_count / target_total_count). High values mean the target is largely covered by this source.
- High overlap with low Unique: the source is mostly redundant and can be deprioritized or disabled.
- Low overlap with high Unique: the source contributes unique entries and may be valuable to keep.

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 151 |
| Total Entries Analyzed | 7.2M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| allowlist | 23 |
| blocklist | 128 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 26 |
| cidr_ipv4 | 2 |
| domain | 84 |
| ipv4 | 39 |

## Detailed Source Analysis

### 1Hosts (Lite)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 128.6K | Targets: 66 | Unique: 0 | Conflicts: 516</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-annoyance | blocklist | domain | 475 | 416 | 87.6% |
| Adaway | blocklist | hostname | 6.5K | 5.3K | 80.4% |
| YousList | blocklist | hostname | 624 | 450 | 72.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 2.3K | 67.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2.3K | 67.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 11.7K | 61.6% |
| WaLLy3K | blocklist | domain | 350 | 183 | 52.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 23.0K | 51.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 245 | 39.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 168 | 35.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.5K | 34.6% |
| hkamran80_smarttv | blocklist | domain | 293 | 100 | 34.1% |
| hufilter | blocklist | hostname | 100 | 32 | 32.0% |
| quidsup_notrack-malware | blocklist | domain | 150 | 46 | 30.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.8K | 24.4% |
| Easy Privacy | allowlist | domain_adguard | 655 | 156 | 23.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 5.1K | 21.9% |
| HaGeZi Pro | blocklist | domain | 409.7K | 88.8K | 21.7% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2.4K | 20.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 26 | 14.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 28.9K | 14.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 7.4K | 14.2% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 126 | 13.0% |
| tranco | allowlist | domain_top | 1.0K | 104 | 10.4% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 2.8K | 8.4% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 24 | 8.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 53 | 6.9% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 96 | 4.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 14.0K | 4.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 193 | 3.9% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 2.7K | 3.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 65 | 2.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 456 | 2.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 14 | 2.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 46 | 1.4% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 247 | 1.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 185 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 424 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 95 | 0.6% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 16 | 0.6% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 7 | 0.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 30 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1.7K | 0.3% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 329 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 23 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 568 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 16 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 114 | 0.1% |
| kadantiscam | blocklist | domain | 228.8K | 138 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 5 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 502 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 316 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 19 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 267 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 35 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 58 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 6 | 0.0% |

</details>

---

### abpvn_hosts

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.1K | Targets: 7 | Unique: 950 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 36 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 24 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 2 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 43 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 10 | 0.0% |

</details>

---

### Adaway

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 6.5K | Targets: 46 | Unique: 0 | Conflicts: 271</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList | blocklist | hostname | 624 | 111 | 17.8% |
| WaLLy3K | blocklist | domain | 350 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2.7K | 14.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 53 | 8.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 254 | 7.4% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 254 | 7.4% |
| hkamran80_smarttv | blocklist | domain | 293 | 21 | 7.2% |
| tranco | allowlist | domain_top | 1.0K | 52 | 5.2% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 194 | 4.5% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 5.3K | 4.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 7 | 3.9% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 75 | 3.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 22 | 3.6% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 402 | 3.4% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 9 | 3.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 24 | 3.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 14 | 2.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 448 | 2.9% |
| quidsup_notrack-malware | blocklist | domain | 150 | 4 | 2.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 49 | 2.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 6.5K | 2.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 751 | 1.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 407 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 10 | 1.5% |
| HaGeZi Pro | blocklist | domain | 409.7K | 5.3K | 1.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 11 | 1.1% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 4 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.1K | 0.5% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 116 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 54 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 26 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 6 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 15 | 0.0% |

</details>

---

### AdBlockID

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 3.8K | Targets: 4 | Unique: 3.8K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 97.6K | 24 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 2 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 2 | 0.0% |

</details>

---

### AdGuard Base filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 97.6K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 54.6K | 54.6K | 99.9% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 28.7K | 64.6% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 40.2K | 30.4% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 28.7K | 14.3% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 5.6K | 4.7% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 54 | 3.7% |
| abpvn_hosts | blocklist | adguard | 1.1K | 10 | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| AdBlockID | blocklist | adguard | 3.8K | 24 | 0.6% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 3 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 69 | 0.4% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 5 | 0.3% |
| Easy Privacy | blocklist | adguard | 53.3K | 91 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 16 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 66 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 72 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 7 | 0.0% |

</details>

---

### AdGuard CNAME Mail Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 37.6K | Targets: 14 | Unique: 37.6K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 6 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 6 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 6 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 8 | 0.0% |

</details>

---

### AdGuard CNAME Trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 75.3K | Targets: 24 | Unique: 49.6K | Conflicts: 46</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 7.3K | 21.8% |
| hufilter | blocklist | hostname | 100 | 20 | 20.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 668 | 4.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 6 | 3.3% |
| HaGeZi Pro | blocklist | domain | 409.7K | 10.9K | 2.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.1K | 2.4% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2.7K | 2.1% |
| Easy Privacy | allowlist | domain_adguard | 655 | 10 | 1.5% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 19 | 1.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.8K | 0.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 190 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 29 | 0.8% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 29 | 0.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 82 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 2 | 0.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 3 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 821 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 40 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 132.2K | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 92.0% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 38.1K | 85.8% |
| EasyList | blocklist | adguard | 54.6K | 35.1K | 64.4% |
| Easy Privacy | blocklist | adguard | 53.3K | 28.0K | 52.6% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 40.2K | 41.2% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 2 | 28.6% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 42.9K | 21.4% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 49 | 14.2% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 5.9K | 5.0% |
| abpvn_hosts | blocklist | adguard | 1.1K | 24 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 137 | 0.7% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 40 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 2 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 165 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 175 | 0.0% |

</details>

---

### AdGuard DNS filter

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 181 | Targets: 28 | Unique: 0 | Conflicts: 194</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | allowlist | domain_adguard | 655 | 9 | 1.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 3 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 19 | 0.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 7 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 6 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 6 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 7 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 7 | 0.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 26 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 5 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 12 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 22 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 26 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 6 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 28 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_android

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 97 | Targets: 11 | Unique: 68 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 6 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 5 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_banks

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 4.0K | Targets: 11 | Unique: 3.9K | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 3 | 1.8% |
| Easy Privacy | allowlist | domain_adguard | 655 | 7 | 1.1% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 7 | 0.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 3 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 3 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_firefox

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 18 | Targets: 4 | Unique: 10 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |

</details>

---

### AdGuardTeam_HttpsExclusions_issues

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 68 | Targets: 6 | Unique: 60 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 2 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_mac

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 11 | Targets: 3 | Unique: 4 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 1 | 0.1% |

</details>

---

### AdGuardTeam_HttpsExclusions_sensitive

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 164 | Targets: 12 | Unique: 133 | Conflicts: 12</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 2 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 11 | 0.0% |

</details>

---

### AdGuardTeam_HttpsExclusions_windows

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 7 | Targets: 1 | Unique: 6 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |

</details>

---

### AntiAdBlockFilters

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.7K | Targets: 1 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 97.6K | 5 | 0.0% |

</details>

---

### bigdargon_hostsVN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 19.0K | Targets: 54 | Unique: 0 | Conflicts: 435</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 2.1K | 62.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2.1K | 62.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.0K | 47.7% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 624 | 204 | 32.7% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 89 | 18.7% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 102 | 16.6% |
| Easy Privacy | allowlist | domain_adguard | 655 | 108 | 16.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 26 | 14.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 5.0K | 11.4% |
| quidsup_notrack-malware | blocklist | domain | 150 | 17 | 11.3% |
| tranco | allowlist | domain_top | 1.0K | 106 | 10.6% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 11.7K | 9.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1.1K | 9.0% |
| hkamran80_smarttv | blocklist | domain | 293 | 26 | 8.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 42 | 8.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1.7K | 7.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 50 | 6.5% |
| hufilter | blocklist | hostname | 100 | 6 | 6.0% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 17 | 5.9% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 96 | 4.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 38 | 3.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 6.4K | 3.2% |
| HaGeZi Pro | blocklist | domain | 409.7K | 12.6K | 3.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 7.8K | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1.2K | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 14 | 2.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 32 | 1.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 144 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 67 | 0.2% |
| Torrent Trackers | blocklist | domain | 485 | 1 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 24 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 26 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 40 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 58 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 20 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 136 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 264 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 19 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 4 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 30 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 14 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 9 | 0.0% |

</details>

---

### BinaryDefense_Banlist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 3.0K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 445 | 8.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 599 | 6.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 961 | 6.4% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 16 | 6.1% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 653 | 5.3% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 3.0K | 4.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 870 | 4.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 467 | 3.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 2.2K | 2.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 41 | 2.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 8 | 1.9% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 8 | 1.9% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 198 | 1.4% |
| Greensnow | blocklist | ipv4 | 5.6K | 77 | 1.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 486 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 16 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 14 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 3 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### BlahDNS_whitelist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 773 | Targets: 47 | Unique: 0 | Conflicts: 481</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 49 | 7.5% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 134 | 6.8% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 107 | 4.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 6 | 3.3% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 14 | 2.9% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 6 | 2.5% |
| tranco | allowlist | domain_top | 1.0K | 24 | 2.4% |
| Dogino_Discord_Official | allowlist | domain | 43 | 1 | 2.3% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| Easy Privacy | allowlist | domain_adguard | 655 | 13 | 2.0% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 47 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 3 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 293 | 2 | 0.7% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 6 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 24 | 0.4% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 14 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 14 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 17 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 50 | 0.3% |
| YousList | blocklist | hostname | 624 | 2 | 0.3% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 19 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 8 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 4 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 68 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 8 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 5 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 53 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 8 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 63 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 9 | 0.0% |

</details>

---

### BlockListDE_Brute

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 914 | Targets: 12 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 14.5K | 874 | 6.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 132 | 5.5% |
| Greensnow | blocklist | ipv4 | 5.6K | 171 | 3.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 685 | 1.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 748 | 0.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 33 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 13 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 8 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 3 | 0.0% |

</details>

---

### BlockListDE_Strong

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 262 | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 92 | 5.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 88 | 3.4% |
| Greensnow | blocklist | ipv4 | 5.6K | 152 | 2.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 8 | 1.9% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 8 | 1.9% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 214 | 1.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 16 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 248 | 0.3% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 229 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 14 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 6 | 0.0% |

</details>

---

### Blocklists UT1 Cryptojacking

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 16.3K | Targets: 40 | Unique: 15.0K | Conflicts: 25</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 4 | 2.7% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 8 | 1.0% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 9 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 95 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 24 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 46 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 17 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 227 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 216 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 95 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 50 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 7 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 9 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 3 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 34 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 262 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 22 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 76 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 7 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 48 | 0.0% |

</details>

---

### Blocklists UT1 Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 225.4K | Targets: 46 | Unique: 0 | Conflicts: 15</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.0% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 20 | 87.0% |
| phishing_army | blocklist | domain | 121.8K | 96.5K | 79.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 43.4% |
| kadantiscam | blocklist | domain | 228.8K | 96.7K | 42.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 65.4K | 32.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 95.3K | 30.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.7K | 21.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1.5K | 8.7% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 992 | 8.4% |
| quidsup_notrack-malware | blocklist | domain | 150 | 10 | 6.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 6.0K | 5.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 46 | 4.7% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1.0K | 4.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 19.0K | 2.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 11.4K | 2.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 6 | 2.2% |
| HaGeZi Pro | blocklist | domain | 409.7K | 7.4K | 1.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 613 | 1.2% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3.9K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 11 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 568 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 50 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 30 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 9 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 27 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 10 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 21 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 4 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 4 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 64 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 15 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 28 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 44 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 5 | 0.0% |

</details>

---

### Blocklists UT1 Publicite

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.3K | Targets: 53 | Unique: 0 | Conflicts: 254</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1.9K | 55.7% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1.9K | 55.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2.0K | 10.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 19 | 10.5% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 49 | 10.3% |
| tranco | allowlist | domain_top | 1.0K | 73 | 7.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 39 | 6.0% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 514 | 4.4% |
| Easy Privacy | allowlist | domain_adguard | 655 | 28 | 4.3% |
| WaLLy3K | blocklist | domain | 350 | 13 | 3.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.6K | 3.6% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| quidsup_notrack-malware | blocklist | domain | 150 | 5 | 3.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 496 | 3.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 61 | 3.1% |
| Adaway | blocklist | hostname | 6.5K | 194 | 3.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 666 | 2.9% |
| hkamran80_smarttv | blocklist | domain | 293 | 7 | 2.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 17 | 2.2% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 8 | 1.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 12 | 1.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1.5K | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.9K | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 2.2K | 0.7% |
| HaGeZi Pro | blocklist | domain | 409.7K | 2.6K | 0.6% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 33 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 4 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 42 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 14 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 5 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 24 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 18 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 5 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 10 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 5 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 18 | 0.0% |

</details>

---

### Blocklists UT1 Shortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.5K | Targets: 32 | Unique: 0 | Conflicts: 56</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 4.1K | 72.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 93 | 39.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 60 | 14.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 27 | 1.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 5 | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 7 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 120 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 4 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 29 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 25 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 14 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 73 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 19 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 32 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 3 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 23 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |

</details>

---

### Borestad_AbuseIPDB_S100_3d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 80.7K | Targets: 33 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 262 | 248 | 94.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.5K | 87.0% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 748 | 81.8% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 326 | 76.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 324 | 75.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11.2K | 74.5% |
| Greensnow | blocklist | ipv4 | 5.6K | 4.2K | 74.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 2.2K | 73.3% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 8.3K | 68.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 9.8K | 67.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 5.9K | 66.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.4K | 46.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 7.3K | 35.6% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 22.9K | 34.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 599 | 24.7% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 2.7K | 19.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 338 | 18.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 417 | 16.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 37 | 7.5% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 21 | 3.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 9 | 3.2% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 1 | 2.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 289 | 1.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 112 | 2 | 1.8% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 203 | 1.6% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 3.1K | 1.5% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 1 | 1.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 21 | 0.8% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 7 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 7 | 0.1% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 16 | 0.0% |

</details>

---

### Boutetnico_URL_Shorteners

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 418 | Targets: 25 | Unique: 191 | Conflicts: 50</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 237 | 56 | 23.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 60 | 1.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 19 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 6 | 0.9% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 6 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 12 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 16 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 8 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 2 | 0.0% |

</details>

---

### BruteforceBlocker

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 427 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 413 | 96.5% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 397 | 3.2% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 8 | 3.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 41 | 2.4% |
| Greensnow | blocklist | ipv4 | 5.6K | 52 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 118 | 0.8% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 427 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 326 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 38 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 8 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 31 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 7 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 15 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 15 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 2 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 3 | 0.0% |

</details>

---

### CINSScore_BadGuys_Army

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.0K | Targets: 21 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.2K | 9.8K | 79.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 961 | 31.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 4.1K | 19.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 1.7K | 19.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 11.2K | 13.8% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 6.8K | 10.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 470 | 9.2% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 70 | 4.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 500 | 3.7% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 15 | 3.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 14 | 3.3% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 385 | 2.7% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 1 | 2.6% |
| Greensnow | blocklist | ipv4 | 5.6K | 99 | 1.8% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 2 | 0.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 18 | 0.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 1.2K | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 48 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 33 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 6 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 4 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 59 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 2 | 0.0% |

</details>

---

### CJX Annoyance

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 6 | Targets: 2 | Unique: 4 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 409.7K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1 | 0.0% |

</details>

---

### CybercrimeTracker_All

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.9K | Targets: 9 | Unique: 1.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 55 | 53.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 488 | 10.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 50 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 473 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 52 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 1 | 0.0% |

</details>

---

### CybercrimeTracker_All

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 13.0K | Targets: 19 | Unique: 159 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 585 | 59.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1.8K | 36.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.1% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 2.7K | 1.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3.9K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 746 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 362 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 17 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 78 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 19 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 25 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 50 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### CybercrimeTracker_CCPMGate

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 103 | Targets: 5 | Unique: 34 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 55 | 1.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 7 | 0.0% |

</details>

---

### CybercrimeTracker_CCPMGate

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 982 | Targets: 12 | Unique: 165 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 585 | 4.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 47 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 1 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 46 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 14 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 88 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 24 | 0.0% |

</details>

---

### cyberhost_malware-blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 17.4K | Targets: 38 | Unique: 29 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | domain_http_url | 23 | 1 | 4.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 9.9K | 1.4% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3.7K | 0.9% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 1.5K | 0.7% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 21 | 0.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 331 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 542 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 86 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 247 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 529 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 3 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 60 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 16 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 305 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 24 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 35 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 22 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 6 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 24 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |

</details>

---

### Dan Pollock's List

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.8K | Targets: 53 | Unique: 0 | Conflicts: 113</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList | blocklist | hostname | 624 | 108 | 17.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 441 | 12.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 441 | 12.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 514 | 12.0% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| Adaway | blocklist | hostname | 6.5K | 402 | 6.1% |
| WaLLy3K | blocklist | domain | 350 | 20 | 5.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.1K | 5.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 52 | 5.4% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 7 | 3.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 11.7K | 3.7% |
| Easy Privacy | allowlist | domain_adguard | 655 | 23 | 3.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 21 | 3.4% |
| hkamran80_smarttv | blocklist | domain | 293 | 9 | 3.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 4 | 2.7% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 42 | 2.1% |
| tranco | allowlist | domain_top | 1.0K | 20 | 2.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2.4K | 1.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 633 | 1.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 326 | 1.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 2.6K | 1.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 199 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 8 | 1.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 8 | 1.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3.0K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 992 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 92 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 122 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 12 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 82 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 46 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 686 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 24 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 84 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 17 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 5 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 37 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 34 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 13 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 48 | 0.0% |

</details>

---

### DandelionSprout-Anti-Malware-List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 32.6K | Targets: 4 | Unique: 32.6K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Most Abused TLDs | blocklist | adguard | 429 | 2 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 7 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 1 | 0.0% |

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

### Dogino_Discord_Official

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 43 | Targets: 6 | Unique: 0 | Conflicts: 7</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 26 | 1.2% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 7 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.3K | Targets: 10 | Unique: 75 | Conflicts: 13</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1.2K | 81.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| Torrent Trackers | blocklist | domain | 485 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |

</details>

---

### DoH_IP_blocklists

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 10 | Unique: 664 | Conflicts: 9</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 1.7K | 97.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 94 | 12.9% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 76 | 8 | 10.5% |
| Local AI Blocklist (Domain) | blocklist | ipv4_from_domain | 55 | 1 | 1.8% |
| Local AI Allowlist (Domain) | allowlist | ipv4_from_domain | 55 | 1 | 1.8% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 100 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 14 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 2 | 0.0% |

</details>

---

### DoH_IP_list

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 731 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 92 | 5.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 94 | 3.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 2 | 0.0% |

</details>

---

### DShield

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.1K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 5.1K | 38.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 5.1K | 25.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 445 | 14.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 910 | 10.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 470 | 3.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 2.4K | 3.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 1.7K | 2.6% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 7 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 6 | 1.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 1.7K | 0.8% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 12 | 0.7% |
| Greensnow | blocklist | ipv4 | 5.6K | 40 | 0.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 33 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 19 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 20 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 19 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 53.3K | Targets: 15 | Unique: 16.1K | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | allowlist | adguard | 1 | 1 | 100.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 91.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 156 | 45.2% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 28.0K | 21.2% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 2.4K | 5.5% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 4.8K | 2.4% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.1K | 2 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 91 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 17 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 7 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 18 | 0.0% |

</details>

---

### Easy Privacy

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 655 | Targets: 42 | Unique: 0 | Conflicts: 949</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 47 | 3 | 6.4% |
| tranco | allowlist | domain_top | 1.0K | 60 | 6.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 9 | 5.0% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 1 | 4.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 60 | 3.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 1 | 2.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 46 | 2.0% |
| BlahDNS_whitelist | allowlist | domain | 773 | 13 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 10 | 1.5% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 45 | 1.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 45 | 1.3% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 53 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 28 | 0.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 108 | 0.6% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 78 | 0.5% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 7 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 23 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 36 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 27 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 156 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 2 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 101 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 10 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 158 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 49 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |

</details>

---

### EasyList

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 54.6K | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 97.6K | 54.6K | 55.9% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 23.8K | 53.6% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 35.1K | 26.6% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 23.8K | 11.9% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 5.6K | 4.7% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 54 | 3.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 69 | 0.4% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 1 | 0.2% |
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 5 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 65 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 70 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 5 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 1 | 0.0% |

</details>

---

### EmergingThreats_CompromisedIPs

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 428 | Targets: 18 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 427 | 413 | 96.7% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 391 | 3.2% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 8 | 3.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 36 | 2.1% |
| Greensnow | blocklist | ipv4 | 5.6K | 46 | 0.8% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 106 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 425 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 33 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 324 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 8 | 0.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 6 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 13 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 27 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 3 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 2 | 0.0% |

</details>

---

### ET_fwip

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 1.6K | Targets: 1 | Unique: 126 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.5K | 32.7% |

</details>

---

### fabriziosalmi_allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 2.3K | Targets: 63 | Unique: 560 | Conflicts: 646</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dogino_Discord_Official | allowlist | domain | 43 | 26 | 60.5% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 20 | 42.6% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 1.0K | 268 | 26.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 480 | 24.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 107 | 13.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 74 | 11.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 2 | 8.7% |
| Easy Privacy | allowlist | domain_adguard | 655 | 46 | 7.0% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 6 | 6.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 32 | 5.2% |
| Local AI Blocklist (Domain) | blocklist | domain | 24 | 1 | 4.2% |
| Local AI Allowlist (Domain) | allowlist | domain | 24 | 1 | 4.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 7 | 3.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 12 | 2.9% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 7 | 2.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 3 | 1.8% |
| hkamran80_smarttv | blocklist | domain | 293 | 4 | 1.4% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 43 | 1.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 12 | 1.2% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 2 | 1.1% |
| Adaway | blocklist | hostname | 6.5K | 49 | 0.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 9 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 15 | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 6 | 0.4% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 8 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 32 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 9 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 33 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 8 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 65 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 15 | 0.1% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 8 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 9 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 10 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 8 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 7 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 10 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 4 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 3 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 77 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 95 | 0.0% |

</details>

---

### FabrizioSalmi_DNS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66 | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 25 | 1.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 25 | 1.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 4 | 0.0% |

</details>

---

### FakeWebshopListHUN

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.2K | Targets: 17 | Unique: 4.7K | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 100 | 9 | 9.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 42 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| phishing_army | blocklist | domain | 121.8K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 29 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 28 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 35 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 7 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 24 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 29 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 22 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |

</details>

---

### Firehol_BitcoinNodes_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 7.3K | Targets: 12 | Unique: 7.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4_cidr_expand | 106 | 47 | 44.3% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 47 | 0.6% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 7 | 0.0% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 18 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 4 | 0.0% |

</details>

---

### Firehol_Botscout_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 617 | Targets: 12 | Unique: 503 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 64 | 3.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 4 | 1.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 6 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 6 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 21 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 1 | 0.0% |

</details>

---

### Firehol_CleanTalk

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 494 | Targets: 12 | Unique: 411 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 5 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.6K | 6 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 2 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 37 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 15 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 2 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 7 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 3 | 0.0% |

</details>

---

### Firehol_CleanTalk_Top20

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 20 | Targets: 12 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 5 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |

</details>

---

### Firehol_GPF_Comics

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.4K | Targets: 23 | Unique: 972 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Brute | blocklist | ipv4 | 914 | 132 | 14.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 51 | 2.8% |
| Greensnow | blocklist | ipv4 | 5.6K | 84 | 1.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 4 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 178 | 1.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 6 | 1.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 112 | 1 | 0.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 599 | 0.7% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 278 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 28 | 0.3% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 1 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 3 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 26 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 13 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 18 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 11 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 3 | 0.0% |

</details>

---

### Firehol_level1

<details>
<summary>List Type: blocklist | Source Type: cidr_ipv4 | Total: 4.5K | Targets: 1 | Unique: 3.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.5K | 92.2% |

</details>

---

### Firehol_level2

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 13.4K | Targets: 30 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 874 | 95.6% |
| Greensnow | blocklist | ipv4 | 5.6K | 5.3K | 93.8% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 214 | 81.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.4K | 81.0% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 118 | 27.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 5.2K | 25.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 106 | 24.8% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 12.8K | 19.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 467 | 15.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 333 | 13.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 9.8K | 12.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 987 | 11.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 178 | 7.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 425 | 3.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 500 | 3.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 186 | 1.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 7 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 186 | 1.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 24 | 1.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 1.7K | 0.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 1 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 24 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 33 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |

</details>

---

### Firehol_level3

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.2K | Targets: 29 | Unique: 0 | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 38 | 100.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 397 | 93.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 391 | 91.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.8K | 65.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 5.2K | 38.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 870 | 28.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 1.8K | 20.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 8.3K | 10.3% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 5.1K | 7.6% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 16 | 6.1% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 651 | 5.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 87 | 5.0% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 76 | 3 | 3.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 651 | 3.2% |
| Greensnow | blocklist | ipv4 | 5.6K | 173 | 3.1% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 425 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 3.2K | 1.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 26 | 1.1% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 3 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 99 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 52 | 0.4% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 3 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 7 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 2 | 0.1% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 6 | 0.0% |

</details>

---

### Firehol_SocksProxy_7d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 16 | Unique: 2.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 112 | 54 | 48.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 31 | 10.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 54 | 2.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 7 | 1.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 20 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 4 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 8 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 34 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 7 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 21 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 2 | 0.0% |

</details>

---

### Firehol_SSLProxies_1d

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 285 | Targets: 12 | Unique: 207 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 18 | 9 | 50.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 9 | 3.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 31 | 1.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 112 | 1 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 4 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 4 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 3 | 0.2% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 9 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 1 | 0.0% |

</details>

---

### Frogeye-firstparty-trackers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 33.3K | Targets: 28 | Unique: 10.0K | Conflicts: 53</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 7.3K | 9.6% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 12 | 6.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 539 | 3.4% |
| HaGeZi Pro | blocklist | domain | 409.7K | 11.0K | 2.7% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2.8K | 2.2% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| Adaway | blocklist | hostname | 6.5K | 116 | 1.8% |
| Easy Privacy | allowlist | domain_adguard | 655 | 8 | 1.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 19 | 1.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 33 | 1.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 33 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 33 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 46 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 67 | 0.4% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 152 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 621 | 0.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 2 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 565 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 26 | 0.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 1 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 1 | 0.0% |

</details>

---

### GetAdmiral Domains Filter List

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.8K | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.3K | 1.6K | 3.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 1.6K | 1.2% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 1.7K | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 80 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 17 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 4 | 0.0% |

</details>

---

### GlobalAntiScamOrg-blocklist-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.1K | Targets: 17 | Unique: 7.3K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 5 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 33 | 0.0% |

</details>

---

### Greensnow

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 5.6K | Targets: 28 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 262 | 152 | 58.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 669 | 38.5% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 5.3K | 36.3% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 171 | 18.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 318 | 12.4% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 52 | 12.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 46 | 10.7% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 5.4K | 8.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 4.2K | 5.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 318 | 3.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 84 | 3.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 77 | 2.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 201 | 1.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 6 | 1.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 173 | 0.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 40 | 0.8% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 86 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 13 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 99 | 0.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 1 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 138 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 3 | 0.0% |

</details>

---

### HaGeZi Amazon Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 615 | Targets: 19 | Unique: 0 | Conflicts: 38</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| hkamran80_smarttv | blocklist | domain | 293 | 4 | 1.4% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 102 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 2 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 22 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 8 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 245 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 21 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 10 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 601 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 31 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 53 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 10 | 0.0% |

</details>

---

### HaGeZi Apple Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 290 | Targets: 14 | Unique: 0 | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlahDNS_whitelist | allowlist | domain | 773 | 3 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 17 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 9 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 243 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 11 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 12 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 15 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 7 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 24 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 120.0K | Targets: 14 | Unique: 62.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1.4K | 67.7% |
| EasyList | blocklist | adguard | 54.6K | 5.6K | 10.3% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 18.4K | 9.2% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 3.0K | 6.8% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 5.6K | 5.8% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 5.9K | 4.5% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 12.6K | 1.6% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 17 | 1.0% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 4.4K | 0.7% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 824 | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 18 | 0.1% |
| Easy Privacy | blocklist | adguard | 53.3K | 18 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 1 | 0.0% |

</details>

---

### HaGeZi DNS TIF Mini

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 120.0K | Targets: 47 | Unique: 2.8K | Conflicts: 3</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | domain_http_url | 23 | 7 | 30.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 7.7K | 30.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 50.7K | 12.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 18.4K | 9.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 20 | 8.4% |
| phishing_army | blocklist | domain | 121.8K | 9.7K | 7.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 3.0K | 6.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 13 | 4.7% |
| quidsup_notrack-malware | blocklist | domain | 150 | 5 | 3.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 120 | 2.7% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 6.0K | 2.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 141 | 2.5% |
| kadantiscam | blocklist | domain | 228.8K | 5.2K | 2.3% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 331 | 1.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 5.5K | 1.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 176 | 1.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 725 | 1.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 37 | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 3 | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 162 | 0.7% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 2.5K | 0.6% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 4.0K | 0.6% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 2.1K | 0.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 67 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 48 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 58 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 23 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 28 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 329 | 0.3% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 33 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 25 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 5 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 22 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 18 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 11 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 120 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 17 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |

</details>

---

### HaGeZi Encrypted DNS Servers

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.4K | Targets: 7 | Unique: 248 | Conflicts: 11</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1.2K | 92.4% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Torrent Trackers | blocklist | domain | 485 | 1 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |

</details>

---

### HaGeZi Gambling Only Domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 184.0K | Targets: 39 | Unique: 177.7K | Conflicts: 11</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1.2K | 44.0% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 105 | 6.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 3.3K | 1.0% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 5 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 7 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 7 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 183 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 12 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 17 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 401 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 590 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 114 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 17 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 20 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 7 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 63 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 24 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 9 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 2 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 60 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 11 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 70 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 44 | 0.0% |

</details>

---

### HaGeZi Microsoft Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 971 | Targets: 19 | Unique: 0 | Conflicts: 36</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 17 | 0.9% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 12 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 52 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 11 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 12 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 11 | 0.3% |
| HaGeZi Pro | blocklist | domain | 409.7K | 774 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 11 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 38 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 38 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 20 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 25 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 126 | 0.1% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 67 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 83 | 0.0% |

</details>

---

### HaGeZi Most Abused TLDs

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 429 | Targets: 1 | Unique: 427 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 2 | 0.0% |

</details>

---

### HaGeZi Pro

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 409.7K | Targets: 66 | Unique: 8.8K | Conflicts: 500</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 473 | 99.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 43.7K | 98.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 601 | 97.7% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3.1K | 91.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3.1K | 91.2% |
| hufilter | blocklist | hostname | 100 | 91 | 91.0% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 425 | 89.5% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 243 | 83.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 122 | 81.3% |
| Adaway | blocklist | hostname | 6.5K | 5.3K | 80.9% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 774 | 79.7% |
| YousList | blocklist | hostname | 624 | 443 | 71.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 88.8K | 69.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 12.6K | 66.6% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.6K | 59.8% |
| hkamran80_smarttv | blocklist | domain | 293 | 140 | 47.8% |
| WaLLy3K | blocklist | domain | 350 | 161 | 46.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 50.7K | 42.2% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 8.2K | 35.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 70.0K | 34.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 11.0K | 32.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 14.0K | 27.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 3.0K | 25.6% |
| Easy Privacy | allowlist | domain_adguard | 655 | 158 | 24.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 368 | 21.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.1K | 19.7% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 4 | 17.4% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 28 | 15.5% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 10.9K | 14.5% |
| tranco | allowlist | domain_top | 1.0K | 107 | 10.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 68 | 8.8% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 26.4K | 8.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 290 | 5.9% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 5.9% |
| phishing_army | blocklist | domain | 121.8K | 6.6K | 5.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1.3K | 5.2% |
| Spam404 | blocklist | domain | 8.1K | 340 | 4.2% |
| kadantiscam | blocklist | domain | 228.8K | 9.3K | 4.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 77 | 3.4% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 7.4K | 3.3% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 529 | 3.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 71 | 2.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 7 | 2.5% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 49 | 2.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 202 | 1.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 346 | 1.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 54 | 1.7% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1.0K | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 227 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 7 | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 5.5K | 0.8% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3.0K | 0.7% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 78 | 0.6% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 29 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.4K | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 590 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 24 | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 750 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 3 | 0.2% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 2 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 3 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 47 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 7 | 0.0% |

</details>

---

### HaGeZi Xiaomi Tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 475 | Targets: 14 | Unique: 0 | Conflicts: 15</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlahDNS_whitelist | allowlist | domain | 773 | 14 | 1.8% |
| hkamran80_smarttv | blocklist | domain | 293 | 4 | 1.4% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 42 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 5 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 168 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 473 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 111 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 29 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 36 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 3 | 0.0% |

</details>

---

### HaGeZi_DoH

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.7K | Targets: 7 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 1.7K | 64.7% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 92 | 12.6% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 99 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 14 | 0.0% |

</details>

---

### HaGeZi_TIF

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 66.4K | Targets: 33 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BruteforceBlocker | blocklist | ipv4_find | 427 | 427 | 100.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 3.0K | 99.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 12.5K | 99.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 425 | 99.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 15.6K | 98.4% |
| Greensnow | blocklist | ipv4 | 5.6K | 5.4K | 95.4% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 12.8K | 87.8% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 229 | 87.4% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1.4K | 83.1% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 685 | 74.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6.8K | 45.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 3.8K | 43.5% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 4.4K | 35.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.7K | 33.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 22.9K | 28.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 5.1K | 24.8% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 9 | 23.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 395 | 15.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 1.9K | 14.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 220 | 12.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 278 | 11.5% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 430 | 9.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 8.4K | 4.1% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 15 | 3.0% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 3 | 2.9% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 52 | 1.8% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 5 | 1.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 6 | 1.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 4 | 0.1% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 10 | 0.0% |

</details>

---

### hkamran80_smarttv

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 293 | Targets: 24 | Unique: 0 | Conflicts: 31</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| tranco | allowlist | domain_top | 1.0K | 8 | 0.8% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 4 | 0.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 13 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 23 | 0.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 4 | 0.7% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 23 | 0.7% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 3 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 21 | 0.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 2 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 100 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 109 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 9 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 21 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 26 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 18 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 53 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 13 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 140 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |

</details>

---

### hufilter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 100 | Targets: 23 | Unique: 0 | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 90 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 5 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 1 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 9 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 32 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 89 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 91 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 14 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 10 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 20 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 5 | 0.0% |

</details>

---

### iam-py-test_my-filters-001-antitypo

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 824 | Targets: 2 | Unique: 822 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 1 | 0.0% |

</details>

---

### jarelllama_Scam-Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 457.7K | Targets: 57 | Unique: 410.5K | Conflicts: 21</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 3.6K | 32.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1.7K | 15.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 3.4K | 13.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 16 | 10.7% |
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 7 | 6.9% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 11.4K | 5.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 12 | 4.4% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 1 | 4.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 5.2K | 2.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1.1K | 2.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 117 | 2.1% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 305 | 1.8% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 2.1K | 1.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 73 | 1.6% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| YousList | blocklist | hostname | 624 | 7 | 1.1% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 4.6K | 1.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 32 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 32 | 0.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 356 | 0.8% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 7 | 0.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 5.0K | 0.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 136 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| phishing_army | blocklist | domain | 121.8K | 783 | 0.6% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 50 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 502 | 0.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 8 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 34 | 0.3% |
| HaGeZi Pro | blocklist | domain | 409.7K | 1.4K | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 14 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 34 | 0.2% |
| kadantiscam | blocklist | domain | 228.8K | 411 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 33 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 1 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 401 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 49 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 4 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 407 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 41 | 0.1% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 434 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |

</details>

---

### kadantiscam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 228.8K | Targets: 40 | Unique: 0 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 121.8K | 108.7K | 89.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 194.4K | 61.2% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 96.7K | 42.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 57.4K | 28.6% |
| quidsup_notrack-malware | blocklist | domain | 150 | 7 | 4.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 5.2K | 4.3% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1.1K | 4.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 9.3K | 2.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 16 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 16 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 31 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 1 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 192 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 29 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 37 | 0.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 14 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 22 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 15 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 418 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 14 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 411 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 27 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 8 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 21 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 14 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 138 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 12 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 5 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 60 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 155 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 21 | 0.0% |

</details>

---

### Korlabs_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 237 | Targets: 29 | Unique: 0 | Conflicts: 44</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 56 | 13.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 93 | 2.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 20 | 1.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 59 | 1.0% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 5 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 2 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 3 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 20 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 16 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 3 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 14 | 0.0% |

</details>

---

### Local AI Allowlist (Domain)

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 55 | Targets: 2 | Unique: 0 | Conflicts: 56</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local AI Blocklist (Domain) | blocklist | ipv4_from_domain | 55 | 55 | 100.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### Local AI Allowlist (Domain)

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 24 | Targets: 4 | Unique: 0 | Conflicts: 25</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local AI Blocklist (Domain) | blocklist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |

</details>

---

### Local AI Blocklist (Domain)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 24 | Targets: 4 | Unique: 0 | Conflicts: 28</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local AI Allowlist (Domain) | allowlist | domain | 24 | 24 | 100.0% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |

</details>

---

### Local AI Blocklist (Domain)

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 55 | Targets: 2 | Unique: 0 | Conflicts: 55</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local AI Allowlist (Domain) | allowlist | ipv4_from_domain | 55 | 55 | 100.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 1 | 0.0% |

</details>

---

### Local Allowlist (Domain)

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 47 | Targets: 13 | Unique: 6 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| VXVault_URLList | blocklist | domain_http_url | 23 | 2 | 8.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 20 | 0.9% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| Easy Privacy | allowlist | domain_adguard | 655 | 3 | 0.5% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 7 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |

</details>

---

### Local Allowlist (ipv4)

<details>
<summary>List Type: allowlist | Source Type: ipv4 | Total: 76 | Targets: 4 | Unique: 58 | Conflicts: 18</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 8 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 3 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 5 | 0.0% |

</details>

---

### Local Blocklist (AdGuard)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7 | Targets: 3 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 4 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 4 | 0.0% |

</details>

---

### Local Blocklist (Domain)

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1 | Targets: 9 | Unique: 0 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 1 | 0.0% |

</details>

---

### Malicious URL Blocklist (URLHaus)

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 2.1K | Targets: 5 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 1.4K | 1.2% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 1.5K | 0.8% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 2.1K | 0.3% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 319 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 2 | 0.0% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 203.9K | Targets: 37 | Unique: 172.7K | Conflicts: 5</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 4.0K | 87.1% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 29 | 76.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 4.4K | 35.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.7K | 32.6% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 473 | 16.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 486 | 16.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 3.2K | 15.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 1.7K | 12.7% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 8.4K | 12.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 834 | 9.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 8.2% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 7 | 6.8% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 76 | 5 | 6.6% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 18 | 1 | 5.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 896 | 5.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 3.1K | 3.9% |
| Greensnow | blocklist | ipv4 | 5.6K | 138 | 2.5% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 247 | 2.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 24 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 175 | 1.2% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 3 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 8 | 0.9% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 14 | 0.8% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 3 | 0.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 3 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 10 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 49 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 14 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 11 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 1 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 7 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 18 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 40 | 0.1% |

</details>

---

### Maltrail_StaticTrails

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 694.7K | Targets: 57 | Unique: 586.3K | Conflicts: 93</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.9% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 9.9K | 56.8% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 3.9K | 29.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 150 | 39 | 26.0% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 3 | 13.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 88 | 9.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 19.0K | 8.4% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 29.6K | 7.3% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 11 | 6.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 12.0K | 6.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 3.1K | 5.9% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 686 | 5.8% |
| Local AI Allowlist (Domain) | allowlist | domain | 24 | 1 | 4.2% |
| Local AI Blocklist (Domain) | blocklist | domain | 24 | 1 | 4.2% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 4.0K | 3.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.1K | 2.5% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 45 | 2.3% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 262 | 1.6% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1.7K | 1.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 264 | 1.4% |
| HaGeZi Pro | blocklist | domain | 409.7K | 5.5K | 1.3% |
| Spam404 | blocklist | domain | 8.1K | 108 | 1.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 9 | 1.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 280 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 5.0K | 1.1% |
| tranco | allowlist | domain_top | 1.0K | 9 | 0.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 25 | 0.6% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 16 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 16 | 0.5% |
| Easy Privacy | allowlist | domain_adguard | 655 | 3 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 3 | 0.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1.7K | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 24 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 18 | 0.4% |
| phishing_army | blocklist | domain | 121.8K | 458 | 0.4% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 1 | 0.2% |
| kadantiscam | blocklist | domain | 228.8K | 418 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 35 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 53 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 14 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 8 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 13 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 21 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 86 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 70 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |

</details>

---

### malware-filter_phishing-filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 25.7K | Targets: 24 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 21.7K | 10.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 23 | 8.4% |
| phishing_army | blocklist | domain | 121.8K | 9.2K | 7.5% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 7.7K | 6.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 9 | 3.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.4K | 0.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 29 | 0.6% |
| kadantiscam | blocklist | domain | 228.8K | 1.1K | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 28 | 0.5% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 1.0K | 0.5% |
| HaGeZi Pro | blocklist | domain | 409.7K | 1.3K | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 53 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 6 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 16 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 76 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 41 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 12 | 0.0% |

</details>

---

### Notracking_Hosts_whitelist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 2.0K | Targets: 68 | Unique: 0 | Conflicts: 1.3K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 330 | 33.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 480 | 21.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 134 | 17.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 3 | 16.7% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 7 | 14.9% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 62 | 9.5% |
| Easy Privacy | allowlist | domain_adguard | 655 | 60 | 9.2% |
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 1 | 9.1% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 2 | 8.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 20 | 8.4% |
| Dogino_Discord_Official | allowlist | domain | 43 | 3 | 7.0% |
| WaLLy3K | blocklist | domain | 350 | 22 | 6.3% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 6 | 6.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 19 | 4.5% |
| hkamran80_smarttv | blocklist | domain | 293 | 13 | 4.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 7 | 3.9% |
| quidsup_notrack-malware | blocklist | domain | 150 | 3 | 2.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 17 | 1.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 51 | 1.5% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 51 | 1.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 61 | 1.4% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 4 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 2 | 1.2% |
| Adaway | blocklist | hostname | 6.5K | 75 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 4 | 0.8% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 23 | 0.7% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 4 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 97 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 27 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 96 | 0.5% |
| Torrent Trackers | blocklist | domain | 485 | 2 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 42 | 0.4% |
| YousList | blocklist | hostname | 624 | 2 | 0.3% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 5 | 0.2% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 1 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 7 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 3 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 45 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 5 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 217 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 17 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 2 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 96 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 19 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 49 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 16 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 4 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 45 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 5 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 11 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 19 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 11 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 17 | 0.0% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 200.8K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 44.0K | 99.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.7K | 98.2% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 1.5K | 72.1% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 4 | 57.1% |
| EasyList | blocklist | adguard | 54.6K | 23.8K | 43.7% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 42.9K | 32.5% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 28.7K | 29.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 18.4K | 15.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 43 | 12.5% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 59.9K | 9.8% |
| Easy Privacy | blocklist | adguard | 53.3K | 4.8K | 8.9% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 70.3K | 8.8% |
| abpvn_hosts | blocklist | adguard | 1.1K | 43 | 4.0% |
| CJX Annoyance | blocklist | adguard | 1.8K | 59 | 3.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 30 | 2.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 73 | 1.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 176 | 0.9% |
| iam-py-test_my-filters-001-antitypo | blocklist | adguard | 824 | 1 | 0.1% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |

</details>

---

### OISD Blocklist Big

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 200.8K | Targets: 67 | Unique: 0 | Conflicts: 145</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 44.0K | 99.2% |
| hufilter | blocklist | hostname | 100 | 89 | 89.0% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 421 | 88.6% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 21.7K | 84.7% |
| quidsup_notrack-malware | blocklist | domain | 150 | 117 | 78.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 2.3K | 65.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2.3K | 65.9% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 12 | 52.2% |
| phishing_army | blocklist | domain | 121.8K | 59.0K | 48.4% |
| YousList | blocklist | hostname | 624 | 288 | 46.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 44.9% |
| WaLLy3K | blocklist | domain | 350 | 151 | 43.1% |
| hkamran80_smarttv | blocklist | domain | 293 | 109 | 37.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 6.4K | 33.7% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 65.4K | 29.0% |
| kadantiscam | blocklist | domain | 228.8K | 57.4K | 25.1% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 3 | 25.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 111 | 23.4% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 28.9K | 22.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4.0K | 22.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2.6K | 21.8% |
| Spam404 | blocklist | domain | 8.1K | 1.7K | 21.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 4.9K | 21.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 66.2K | 20.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 9.7K | 18.6% |
| HaGeZi Pro | blocklist | domain | 409.7K | 70.0K | 17.1% |
| Adaway | blocklist | hostname | 6.5K | 1.1K | 16.9% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 18.4K | 15.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.7K | 10.5% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 448 | 9.1% |
| Easy Privacy | allowlist | domain_adguard | 655 | 49 | 7.5% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 67 | 6.9% |
| tranco | allowlist | domain_top | 1.0K | 67 | 6.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 16 | 5.8% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 15 | 5.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 31 | 5.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 6 | 3.3% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 542 | 3.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 362 | 2.8% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 1.8K | 2.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 621 | 1.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 12.0K | 1.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 153 | 1.4% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 14 | 1.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 216 | 1.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 5.2K | 1.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 176 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 29 | 0.9% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 425 | 0.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 3 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 19 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 35 | 0.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 3 | 0.4% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1.6K | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 6 | 0.3% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 9 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 18 | 0.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 334 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 183 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 44 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 18.7K | Targets: 8 | Unique: 18.1K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 78 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 69 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 137 | 0.1% |
| EasyList | blocklist | adguard | 54.6K | 69 | 0.1% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 176 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 18 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 27 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 26 | 0.0% |

</details>

---

### OISD Blocklist NSFW Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.7K | Targets: 39 | Unique: 0 | Conflicts: 33</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 7.1K | 12.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 13.9K | 5.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 7.8K | 2.5% |
| quidsup_notrack-malware | blocklist | domain | 150 | 2 | 1.3% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 17 | 0.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 44 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 12 | 0.4% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 12 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 67 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| Torrent Trackers | blocklist | domain | 485 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 78 | 0.2% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 6 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 176 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 36 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 185 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 346 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 17 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 26 | 0.1% |
| kadantiscam | blocklist | domain | 228.8K | 14 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 12 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 14 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 21 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 5 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 33 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 44.4K | Targets: 16 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 4 | 57.1% |
| EasyList | blocklist | adguard | 54.6K | 23.8K | 43.5% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 28.7K | 29.4% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 38.1K | 28.8% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 44.0K | 21.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 39 | 11.3% |
| Easy Privacy | blocklist | adguard | 53.3K | 2.4K | 4.6% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 80 | 4.5% |
| abpvn_hosts | blocklist | adguard | 1.1K | 36 | 3.4% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 3.0K | 2.5% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 78 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 3 | 0.2% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 74 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 81 | 0.0% |

</details>

---

### OISD Blocklist Small

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 44.4K | Targets: 61 | Unique: 0 | Conflicts: 95</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 100 | 90 | 90.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1.7K | 49.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1.7K | 49.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.6K | 37.8% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 5.0K | 26.6% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 44.0K | 21.9% |
| quidsup_notrack-malware | blocklist | domain | 150 | 30 | 20.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 23.0K | 17.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 8.8K | 17.0% |
| YousList | blocklist | hostname | 624 | 98 | 15.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 3.4K | 14.6% |
| Adaway | blocklist | hostname | 6.5K | 751 | 11.5% |
| HaGeZi Pro | blocklist | domain | 409.7K | 43.7K | 10.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.1K | 6.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 29 | 6.1% |
| WaLLy3K | blocklist | domain | 350 | 21 | 6.0% |
| Easy Privacy | allowlist | domain_adguard | 655 | 36 | 5.5% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 26 | 5.5% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 633 | 5.4% |
| tranco | allowlist | domain_top | 1.0K | 46 | 4.6% |
| hkamran80_smarttv | blocklist | domain | 293 | 13 | 4.4% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 12 | 4.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 25 | 2.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 3.0K | 2.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 4.5K | 1.4% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 1.1K | 1.4% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 7 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 22 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 95 | 0.6% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 152 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 78 | 0.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 3 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 227 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 60 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 2 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1.1K | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 12 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 6 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 356 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 7 | 0.1% |
| phishing_army | blocklist | domain | 121.8K | 6 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 55 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 64 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 114 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 63 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 27 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2 | 0.0% |

</details>

---

### OpenPhish_Feed

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 275 | Targets: 9 | Unique: 33 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| phishing_army | blocklist | domain | 121.8K | 163 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 23 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 7 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 16 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 13 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 12 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |

</details>

---

### Peter Lowe's Blocklist

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.4K | Targets: 56 | Unique: 0 | Conflicts: 214</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3.4K | 100.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 44.6% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 166 | 34.9% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2.1K | 11.2% |
| hkamran80_smarttv | blocklist | domain | 293 | 23 | 7.8% |
| tranco | allowlist | domain_top | 1.0K | 78 | 7.8% |
| Easy Privacy | allowlist | domain_adguard | 655 | 45 | 6.9% |
| quidsup_notrack-malware | blocklist | domain | 150 | 10 | 6.7% |
| WaLLy3K | blocklist | domain | 350 | 19 | 5.4% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Adaway | blocklist | hostname | 6.5K | 254 | 3.9% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.7K | 3.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 589 | 3.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 441 | 3.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 847 | 3.6% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 6 | 3.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 51 | 2.6% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2.3K | 1.8% |
| BlahDNS_whitelist | allowlist | domain | 773 | 14 | 1.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 11 | 1.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 3.4K | 1.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 2.3K | 1.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 11 | 1.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 5 | 1.1% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 3 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3.1K | 0.8% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 3 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 66 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 33 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 12 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 66 | 0.1% |
| kadantiscam | blocklist | domain | 228.8K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 32 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 16 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 4 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 29 | 0.0% |

</details>

---

### phishing_army

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 121.8K | Targets: 33 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 163 | 59.3% |
| kadantiscam | blocklist | domain | 228.8K | 108.7K | 47.5% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 96.5K | 42.8% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 9.2K | 35.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 59.0K | 29.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 92.7K | 29.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 9.7K | 8.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 14 | 5.9% |
| HaGeZi Pro | blocklist | domain | 409.7K | 6.6K | 1.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 32 | 0.7% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 37 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 382 | 0.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 783 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 458 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 24 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 2 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 58 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 2 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 3 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 157 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 6 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |

</details>

---

### Public_DNS4

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 62.6K | Targets: 23 | Unique: 61.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 99 | 5.8% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 100 | 3.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 34 | 1.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 112 | 1 | 0.9% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 1 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 1 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 10 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 40 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 4 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 4 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 16 | 0.0% |

</details>

---

### PuppyScams

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 102 | Targets: 2 | Unique: 85 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |

</details>

---

### quidsup_notrack-annoyance

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 475 | Targets: 18 | Unique: 0 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 166 | 4.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 166 | 4.9% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 49 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 89 | 0.5% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 416 | 0.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 421 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 54 | 0.2% |
| HaGeZi Pro | blocklist | domain | 409.7K | 425 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 26 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 171 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 7 | 0.0% |

</details>

---

### quidsup_notrack-malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 150 | Targets: 29 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 10 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 10 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 30 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 74 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 117 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 17 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 10 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 16 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 39 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 5 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 46 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 122 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 26 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 7 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 4 | 0.0% |

</details>

---

### quidsup_notrack-tracker

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 15.7K | Targets: 49 | Unique: 0 | Conflicts: 366</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 589 | 17.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 589 | 17.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 78 | 11.9% |
| tranco | allowlist | domain_top | 1.0K | 117 | 11.7% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 496 | 11.6% |
| hkamran80_smarttv | blocklist | domain | 293 | 21 | 7.2% |
| Adaway | blocklist | hostname | 6.5K | 448 | 6.9% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.3K | 6.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 9 | 5.0% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 97 | 4.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 963 | 4.1% |
| hufilter | blocklist | hostname | 100 | 4 | 4.0% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 38 | 3.9% |
| YousList | blocklist | hostname | 624 | 21 | 3.4% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 9 | 3.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 3.8K | 3.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 12 | 2.5% |
| BlahDNS_whitelist | allowlist | domain | 773 | 19 | 2.5% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.1K | 2.4% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 12 | 1.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 199 | 1.7% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 539 | 1.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 10 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 668 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 25 | 0.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.7K | 0.8% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3.1K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 304 | 0.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1.3K | 0.4% |
| Torrent Trackers | blocklist | domain | 485 | 2 | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 17 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 15 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 23 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 8 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 41 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 35 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 9 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |

</details>

---

### RedDragonWebDesign_block-everything

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 652 | Targets: 2 | Unique: 648 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 97.6K | 3 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 1 | 0.0% |

</details>

---

### RPiList_specials-malware

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 611.4K | Targets: 12 | Unique: 315.3K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 2.1K | 99.8% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 59.9K | 29.8% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 229.1K | 28.8% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 4.4K | 3.7% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 74 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 165 | 0.1% |
| EasyList | blocklist | adguard | 54.6K | 65 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 27 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 66 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 1 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 17 | 0.0% |

</details>

---

### RPiList_specials-phishing

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 794.5K | Targets: 12 | Unique: 481.7K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-malware | blocklist | adguard | 611.4K | 229.1K | 37.5% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 70.3K | 35.0% |
| Malicious URL Blocklist (URLHaus) | blocklist | adguard | 2.1K | 319 | 15.1% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 12.6K | 10.5% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 4 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 81 | 0.2% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 175 | 0.1% |
| EasyList | blocklist | adguard | 54.6K | 70 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 18.7K | 26 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 72 | 0.1% |
| Easy Privacy | blocklist | adguard | 53.3K | 7 | 0.0% |

</details>

---

### Rutgers_DROP

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.7K | Targets: 19 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 262 | 92 | 35.1% |
| Greensnow | blocklist | ipv4 | 5.6K | 669 | 11.9% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 1.4K | 9.7% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 41 | 9.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 36 | 8.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 75 | 2.9% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 1.4K | 2.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 1.5K | 1.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 155 | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 41 | 1.4% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 87 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 70 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 47 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 39 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 12 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 24 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 1 | 0.0% |

</details>

---

### Sblam_Blocklist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 1.8K | Targets: 17 | Unique: 1.0K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 64 | 10.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 51 | 2.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 285 | 3 | 1.1% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 5 | 1.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 20 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 338 | 0.4% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 220 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.6K | 13 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 24 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 14 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 10 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

</details>

---

### ScriptzTeam_BadIPS

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 2.6K | Targets: 16 | Unique: 897 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 262 | 88 | 33.6% |
| Greensnow | blocklist | ipv4 | 5.6K | 318 | 5.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 75 | 4.3% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 333 | 2.3% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 395 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 417 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 12 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 3 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 16 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 5 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |

</details>

---

### Sentinel_Greylist

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.8K | Targets: 26 | Unique: 0 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 599 | 19.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 910 | 17.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.7K | 11.3% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 38 | 8.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 155 | 8.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 1.8K | 8.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 33 | 7.7% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 943 | 7.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 987 | 7.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 5.9K | 7.3% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 957 | 6.6% |
| Greensnow | blocklist | ipv4 | 5.6K | 318 | 5.7% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 3.8K | 5.7% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 14 | 5.3% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 33 | 3.6% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 1 | 2.6% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 76 | 2 | 2.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 236 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 28 | 1.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 12 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 834 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 28 | 0.2% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_Allowlist

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 654 | Targets: 38 | Unique: 220 | Conflicts: 219</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_windows | allowlist | domain | 7 | 1 | 14.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 49 | 6.3% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 74 | 3.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 62 | 3.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 3 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| Easy Privacy | allowlist | domain_adguard | 655 | 10 | 1.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 2 | 1.2% |
| tranco | allowlist | domain_top | 1.0K | 11 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 293 | 3 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 39 | 0.9% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 16 | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 11 | 0.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 11 | 0.3% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 2 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 14 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 8 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 2 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 32 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Ads

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 23.3K | Targets: 51 | Unique: 0 | Conflicts: 160</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 847 | 24.8% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 847 | 24.8% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 666 | 15.6% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 624 | 85 | 13.6% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 54 | 11.4% |
| hufilter | blocklist | hostname | 100 | 10 | 10.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.7K | 8.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 3.4K | 7.7% |
| Adaway | blocklist | hostname | 6.5K | 407 | 6.2% |
| hkamran80_smarttv | blocklist | domain | 293 | 18 | 6.1% |
| tranco | allowlist | domain_top | 1.0K | 61 | 6.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 963 | 6.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 7 | 4.7% |
| Easy Privacy | allowlist | domain_adguard | 655 | 27 | 4.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 5.1K | 4.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 326 | 2.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 5 | 2.8% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 7 | 2.4% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 4.9K | 2.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 45 | 2.3% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 20 | 2.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 8.2K | 2.0% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 10 | 1.6% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1.9K | 0.6% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 3 | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 67 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 11 | 0.3% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 190 | 0.3% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 116 | 0.2% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 162 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 26 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 9 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 17 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 49 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 280 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 12 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 278.1K | Targets: 34 | Unique: 210.7K | Conflicts: 27</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 13.9K | 74.2% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 24.1K | 40.8% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 27.3K | 8.6% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 12 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 114 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 267 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 334 | 0.2% |
| HaGeZi Pro | blocklist | domain | 409.7K | 750 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 4 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 120 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 13 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 4 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 407 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 28 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 17 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 24 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 6 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 21 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 86 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Malware

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 52.0K | Targets: 42 | Unique: 3.1K | Conflicts: 14</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 74 | 49.3% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 8.8K | 19.9% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.2K | 6.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 7.4K | 5.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 9.7K | 4.8% |
| HaGeZi Pro | blocklist | domain | 409.7K | 14.0K | 3.4% |
| YousList | blocklist | hostname | 624 | 19 | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 304 | 1.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 66 | 1.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 66 | 1.9% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 7 | 1.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 42 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 92 | 0.8% |
| Spam404 | blocklist | domain | 8.1K | 45 | 0.6% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 725 | 0.6% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 9 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 86 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3.1K | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 46 | 0.3% |
| phishing_army | blocklist | domain | 121.8K | 382 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 613 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 15 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 510 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.1K | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 36 | 0.2% |
| kadantiscam | blocklist | domain | 228.8K | 192 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 43 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 9 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 12 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 147 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 8 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 5 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### ShadowWhisperer_BlockLists Scam

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 11.1K | Targets: 28 | Unique: 7.6K | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 10 | 9.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 42 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 34 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.7K | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 991 | 0.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 3 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 44 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 153 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 176 | 0.1% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 15 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 8 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 4 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 30 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 14 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 12 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 27 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 202 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

</details>

---

### ShadowWhisperer_UrlShortener

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 5.7K | Targets: 26 | Unique: 1.1K | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4.1K | 90.8% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 59 | 24.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 16 | 3.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 5 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 141 | 0.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 4 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 28 | 0.1% |
| phishing_army | blocklist | domain | 121.8K | 37 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 5 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 24 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 15 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 9 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 3 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 18 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 117 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 6 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |

</details>

---

### Sinfonietta_Adult

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 58.9K | Targets: 42 | Unique: 0 | Conflicts: 45</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 7.1K | 38.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 58.9K | 18.5% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 24.1K | 8.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 66 | 1.9% |
| Torrent Trackers | blocklist | domain | 485 | 9 | 1.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 66 | 1.9% |
| YousList | blocklist | hostname | 624 | 11 | 1.8% |
| BlahDNS_whitelist | allowlist | domain | 773 | 8 | 1.0% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 122 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 144 | 0.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 16 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 24 | 0.6% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 227 | 0.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 116 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 424 | 0.3% |
| HaGeZi Pro | blocklist | domain | 409.7K | 1.0K | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 425 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 15 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 43 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 23 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 21 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 41 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 12 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 17 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 6 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 5 | 0.0% |

</details>

---

### Sinfonietta_Gambling

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 2.6K | Targets: 20 | Unique: 0 | Conflicts: 6</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 2.6K | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 1.2K | 0.6% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 11 | 0.6% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 5 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 9 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 16 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 6 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 71 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### Sinfonietta_Social

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.2K | Targets: 32 | Unique: 2.8K | Conflicts: 178</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dogino_Discord_Official | allowlist | domain | 43 | 7 | 16.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 47 | 6.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 5 | 5.2% |
| tranco | allowlist | domain_top | 1.0K | 33 | 3.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 16 | 2.4% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 43 | 1.9% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 23 | 1.2% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 23 | 0.4% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 13 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 13 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 655 | 2 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 25 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 32 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 22 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 30 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 46 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 11 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 29 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 54 | 0.0% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 1 | 0.0% |

</details>

---

### Spam404

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 8.1K | Targets: 31 | Unique: 5.7K | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 2 | 1.3% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.7K | 0.9% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 7 | 0.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 45 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 340 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 23 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 31 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 108 | 0.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 23 | 0.0% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 63 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 11 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 17 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1 | 0.0% |

</details>

---

### Stamparm_Blackbook

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 18.1K | Targets: 23 | Unique: 0 | Conflicts: 4</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.4% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.6K | 19.7% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 17.6K | 7.8% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 47 | 4.8% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 4.0K | 2.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 5.0K | 0.7% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 456 | 0.4% |
| HaGeZi Pro | blocklist | domain | 409.7K | 1.1K | 0.3% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 955 | 0.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 4 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 16 | 0.1% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 67 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 2 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 6 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 21 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 115 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 8 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 32 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |

</details>

---

### StevenBlack_Fake_Gambling_Porn

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 317.8K | Targets: 67 | Unique: 0 | Conflicts: 640</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 58.9K | 100.0% |
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 2.6K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 11.7K | 99.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3.4K | 98.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3.4K | 98.9% |
| kadantiscam | blocklist | domain | 228.8K | 194.4K | 85.0% |
| phishing_army | blocklist | domain | 121.8K | 92.7K | 76.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 2.2K | 52.2% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 95.3K | 42.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 7.8K | 41.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 7.8K | 41.2% |
| YousList | blocklist | hostname | 624 | 243 | 38.9% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 171 | 36.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 66.2K | 33.0% |
| WaLLy3K | blocklist | domain | 350 | 86 | 24.6% |
| hkamran80_smarttv | blocklist | domain | 293 | 53 | 18.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 26 | 17.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 101 | 15.4% |
| hufilter | blocklist | hostname | 100 | 14 | 14.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 22 | 12.2% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 217 | 11.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 14.0K | 10.9% |
| tranco | allowlist | domain_top | 1.0K | 107 | 10.7% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 4.5K | 10.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 27.3K | 9.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 991 | 8.9% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 53 | 8.6% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 83 | 8.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 63 | 8.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1.9K | 8.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 36 | 7.6% |
| HaGeZi Pro | blocklist | domain | 409.7K | 26.4K | 6.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 32 | 4.9% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 5.5K | 4.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 95 | 4.2% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 11 | 3.8% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| Torrent Trackers | blocklist | domain | 485 | 9 | 1.9% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 3.3K | 1.8% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 565 | 1.7% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 821 | 1.1% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 510 | 1.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 30 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| Spam404 | blocklist | domain | 8.1K | 63 | 0.8% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 29 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 275 | 1 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 76 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 14 | 0.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 32 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 1.7K | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 35 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 434 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 382 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 6 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 6 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 8 | 0.0% |

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
<summary>List Type: blocklist | Source Type: domain | Total: 485 | Targets: 9 | Unique: 457 | Conflicts: 2</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 9 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 2 | 0.0% |

</details>

---

### tranco

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 1.0K | Targets: 53 | Unique: 0 | Conflicts: 1.1K</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 10 | 23.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 330 | 16.7% |
| Local AI Allowlist (Domain) | allowlist | domain | 24 | 3 | 12.5% |
| Local AI Blocklist (Domain) | blocklist | domain | 24 | 3 | 12.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 268 | 11.9% |
| AdGuardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| Easy Privacy | allowlist | domain_adguard | 655 | 60 | 9.2% |
| VXVault_URLList | blocklist | domain_http_url | 23 | 1 | 4.3% |
| AdGuardTeam_HttpsExclusions_android | allowlist | domain | 97 | 4 | 4.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 24 | 3.1% |
| hufilter | blocklist | hostname | 100 | 3 | 3.0% |
| hkamran80_smarttv | blocklist | domain | 293 | 8 | 2.7% |
| AdGuardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 4 | 2.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 78 | 2.3% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 78 | 2.3% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 73 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 11 | 1.7% |
| AdGuardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 33 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 52 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 117 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 106 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 61 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 7 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 20 | 0.2% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 1 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 3 | 0.2% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 46 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 104 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 10 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 67 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 2 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 107 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 107 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 10 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 4 | 0.0% |

</details>

---

### Ukrainian Ad Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 1.4K | Targets: 9 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EasyList | blocklist | adguard | 54.6K | 54 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 29 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 54 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 794.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 29 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 2 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | adguard | 120.0K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 30 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 2 | 0.0% |

</details>

---

### Ukrainian Privacy Filter

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 345 | Targets: 9 | Unique: 51 | Conflicts: 1</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.3K | 156 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 39 | 0.1% |
| Easy Privacy | allowlist | adguard | 754 | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 2 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 2 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 43 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 49 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 1 | 0.0% |

</details>

---

### Ukrainian Security Filter

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 1.7K | Targets: 14 | Unique: 1.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sinfonietta_Gambling | blocklist | hostname | 2.6K | 11 | 0.4% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 105 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 368 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 46 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 7 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 4 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 13 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3 | 0.0% |

</details>

---

### URLHaus_Text

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 15.9K | Targets: 22 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 15.6K | 23.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 236 | 2.7% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 1 | 2.6% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 3 | 0.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 3 | 0.7% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 34 | 0.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 72 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 33 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 14 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 99 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 896 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 289 | 0.4% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 36 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 33 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 33 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.6K | 3 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 17 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |

</details>

---

### USOM-Blocklists-domains

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 404.1K | Targets: 51 | Unique: 351.0K | Conflicts: 40</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 3.7K | 21.2% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 810 | 16.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 16 | 6.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 9 | 6.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 746 | 5.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 955 | 5.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 29.6K | 4.3% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 24 | 2.4% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 2.5K | 2.1% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 8 | 1.9% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 3.9K | 1.7% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4.6K | 1.0% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 17 | 0.9% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 1.6K | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3.0K | 0.7% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 84 | 0.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| YousList | blocklist | hostname | 624 | 4 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 76 | 0.5% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 23 | 0.5% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| Torrent Trackers | blocklist | domain | 485 | 2 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 147 | 0.3% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 15 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 316 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 41 | 0.2% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 3 | 0.1% |
| phishing_army | blocklist | domain | 121.8K | 157 | 0.1% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 382 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 19 | 0.1% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 55 | 0.1% |
| kadantiscam | blocklist | domain | 228.8K | 155 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 6 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 6 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 17 | 0.0% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 9 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 5 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 2 | 0.0% |
| AdGuardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 7 | 0.0% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 1 | 0.0% |

</details>

---

### USOM-Blocklists-ips

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 12.6K | Targets: 33 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 12.5K | 18.8% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 4 | 10.5% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 269 | 5.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 3 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 4.4K | 2.2% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 50 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 262 | 4 | 1.5% |
| Sblam_Blocklist | blocklist | ipv4 | 1.8K | 14 | 0.8% |
| BruteforceBlocker | blocklist | ipv4_find | 427 | 2 | 0.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 428 | 2 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 16 | 0.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 72 | 0.5% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 49 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 20 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 203 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 48 | 0.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 52 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 28 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.7K | 3 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.5K | 24 | 0.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 617 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 21 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.4K | 4 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.6K | 9 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 914 | 1 | 0.1% |
| HaGeZi_DoH | blocklist | ipv4 | 1.7K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 4.6K | Targets: 13 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 488 | 17.0% |
| VXVault_URLList | blocklist | ipv4_http_url | 38 | 4 | 10.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 269 | 2.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 4.0K | 2.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 430 | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 34 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.4K | 7 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 2 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 7 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 12.2K | 4 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 13.4K | 1 | 0.0% |

</details>

---

### Viriback_Dump

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 4.9K | Targets: 17 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1.8K | 13.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 2.1K | 0.9% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 4.0K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 810 | 0.2% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 193 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 448 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 21 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 290 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 37 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 166 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 3 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 23 | Targets: 12 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 47 | 2 | 4.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 20 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 4 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 12 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3 | 0.0% |

</details>

---

### VXVault_URLList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 38 | Targets: 9 | Unique: 0 | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.2K | 38 | 0.3% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 4 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.8K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 4 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 80.7K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 29 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 15.9K | 1 | 0.0% |

</details>

---

### WaLLy3K

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 350 | Targets: 34 | Unique: 0 | Conflicts: 45</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList | blocklist | hostname | 624 | 9 | 1.4% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 22 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 655 | 5 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| hkamran80_smarttv | blocklist | domain | 293 | 2 | 0.7% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 1 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 19 | 0.6% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 19 | 0.6% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 85 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 13 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 53 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 20 | 0.2% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 1 | 0.2% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 151 | 0.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 183 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 86 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 12 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 21 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 161 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

</details>

---

### YousList

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 624 | Targets: 33 | Unique: 0 | Conflicts: 8</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 204 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 108 | 0.9% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 22 | 0.6% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 22 | 0.6% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 3 | 0.5% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 22 | 0.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 85 | 0.4% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 450 | 0.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 2 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| hkamran80_smarttv | blocklist | domain | 293 | 1 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 1 | 0.2% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 98 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 21 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.7K | 443 | 0.1% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 288 | 0.1% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 2 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 243 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 19 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 4 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 3 | 0.0% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 11 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 3 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: allowlist | Source Type: domain | Total: 12 | Targets: 11 | Unique: 0 | Conflicts: 19</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hkamran80_smarttv | blocklist | domain | 293 | 1 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 1 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 2 | 0.0% |

</details>

---

### YousList-AdGuard

<details>
<summary>List Type: blocklist | Source Type: adguard | Total: 7.3K | Targets: 8 | Unique: 7.2K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 25 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 97.6K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 132.2K | 40 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 10 | 0.0% |
| EasyList | blocklist | adguard | 54.6K | 10 | 0.0% |
| OISD Blocklist Big | blocklist | adguard | 200.8K | 73 | 0.0% |
| RPiList_specials-malware | blocklist | adguard | 611.4K | 1 | 0.0% |

</details>

---

### youtube_GoodbyeAds

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 97.6K | Targets: 26 | Unique: 97.2K | Conflicts: 35</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 2 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| hkamran80_smarttv | blocklist | domain | 293 | 3 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 11 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| BlahDNS_whitelist | allowlist | domain | 773 | 4 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 50 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 52 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 655 | 2 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 2 | 0.3% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 7 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 8 | 0.2% |
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 8 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 35 | 0.0% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 409.7K | 47 | 0.0% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 44 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 6 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 74 | 0.0% |

</details>

---

### Yoyo Adservers-Hosts

<details>
<summary>List Type: blocklist | Source Type: domain | Total: 3.4K | Targets: 56 | Unique: 0 | Conflicts: 214</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Peter Lowe's Blocklist | blocklist | domain | 3.4K | 3.4K | 100.0% |
| Blocklists UT1 Publicite | blocklist | domain | 4.3K | 1.9K | 44.6% |
| quidsup_notrack-annoyance | blocklist | domain | 475 | 166 | 34.9% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 2.1K | 11.2% |
| hkamran80_smarttv | blocklist | domain | 293 | 23 | 7.8% |
| tranco | allowlist | domain_top | 1.0K | 78 | 7.8% |
| Easy Privacy | allowlist | domain_adguard | 655 | 45 | 6.9% |
| quidsup_notrack-malware | blocklist | domain | 150 | 10 | 6.7% |
| WaLLy3K | blocklist | domain | 350 | 19 | 5.4% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Adaway | blocklist | hostname | 6.5K | 254 | 3.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 589 | 3.8% |
| OISD Blocklist Small | blocklist | domain_adguard | 44.4K | 1.7K | 3.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 441 | 3.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 847 | 3.6% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 181 | 6 | 3.3% |
| Notracking_Hosts_whitelist | allowlist | domain | 2.0K | 51 | 2.6% |
| Local Allowlist (Domain) | allowlist | domain | 47 | 1 | 2.1% |
| 1Hosts (Lite) | blocklist | domain | 128.6K | 2.3K | 1.8% |
| BlahDNS_whitelist | allowlist | domain | 773 | 14 | 1.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 654 | 11 | 1.7% |
| OISD Blocklist Big | blocklist | domain_adguard | 200.8K | 2.3K | 1.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 317.8K | 3.4K | 1.1% |
| HaGeZi Microsoft Tracker | blocklist | domain | 971 | 11 | 1.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 475 | 5 | 1.1% |
| HaGeZi Apple Tracker | blocklist | domain | 290 | 3 | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| HaGeZi Pro | blocklist | domain | 409.7K | 3.1K | 0.8% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| HaGeZi Amazon Tracker | blocklist | domain | 615 | 3 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Sinfonietta_Social | blocklist | hostname | 3.2K | 13 | 0.4% |
| Sinfonietta_Adult | blocklist | hostname | 58.9K | 66 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 52.0K | 66 | 0.1% |
| Frogeye-firstparty-trackers | blocklist | hostname | 33.3K | 33 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 18.7K | 12 | 0.1% |
| AdGuard CNAME Mail Trackers | blocklist | domain | 37.6K | 6 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 278.1K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 184.0K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 694.7K | 16 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 225.4K | 4 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| AdGuard CNAME Trackers | blocklist | domain | 75.3K | 29 | 0.0% |
| phishing_army | blocklist | domain | 121.8K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 404.1K | 3 | 0.0% |
| kadantiscam | blocklist | domain | 228.8K | 16 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 32 | 0.0% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 5.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.1K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| HaGeZi DNS TIF Mini | blocklist | domain_adguard | 120.0K | 5 | 0.0% |

</details>

---

### Yoyo AdServers-IPList

<details>
<summary>List Type: blocklist | Source Type: ipv4 | Total: 8.9K | Targets: 4 | Unique: 8.9K | Conflicts: 0</summary>

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.3K | 1 | 0.0% |
| HaGeZi_TIF | blocklist | ipv4 | 66.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.9K | 49 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.6K | 1 | 0.0% |

</details>

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources.

**Note:** Per-source percentages are computed as (overlap_count / source_total_count) × 100. In `Overlap with Other Sources` table the displayed Overlap % is computed relative to the target (overlap_count / target_total_count) × 100.

