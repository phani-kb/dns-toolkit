# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2025-07-12 14:29:29 UTC

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 88 |
| Total Entries Analyzed | 3.2M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 75 |
| allowlist | 13 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 17 |
| cidr_ipv4 | 3 |
| domain | 42 |
| ipv4 | 26 |

## Detailed Source Analysis

### abpvn_hosts

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 51 | **Unique Entries:** 0 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 5 | 7.7% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 1 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 10 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 13 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 8 | 0.1% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 12 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 12 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 12 | 0.0% |

---

### abpvn_hosts

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 1.0K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 24 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 2 | 0.0% |

---

### Adaway

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 6.5K | **Unique Entries:** 0 | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList | blocklist | hostname | 624 | 111 | 17.8% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 8 | 15.7% |
| WaLLy3K | blocklist | domain | 350 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 2.7K | 14.4% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| Easy Privacy | allowlist | domain_adguard | 646 | 52 | 8.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 254 | 7.5% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 4 | 6.2% |
| AdBlockID | allowlist | domain_adguard | 53 | 3 | 5.7% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 133 | 4.9% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 155 | 4.2% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 7 | 4.0% |
| transco | allowlist | domain_top | 100 | 4 | 4.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 14 | 2.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 449 | 2.9% |
| quidsup_notrack-malware | blocklist | domain | 154 | 4 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 49 | 2.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 6.5K | 1.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 407 | 1.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 10 | 1.5% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 4.2K | 1.4% |
| HaGeZi Pro | blocklist | domain | 406.4K | 5.3K | 1.3% |
| Frogeye trackers | blocklist | hostname | 33.4K | 119 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 15 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 26 | 0.0% |

---

### AdBlockID

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 53 | **Unique Entries:** 0 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 2 | 3.9% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 10 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 7 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |

---

### AdBlockID

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 3.8K | **Unique Entries:** 3.8K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 93.5K | 26 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 2 | 0.0% |

---

### AdGuard Base filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 93.5K | **Unique Entries:** 56.5K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 114.5K | 36.8K | 32.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 55 | 3.8% |
| abpvn_hosts | blocklist | adguard | 1.1K | 10 | 1.0% |
| AdBlockID | blocklist | adguard | 3.8K | 26 | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 3 | 0.5% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 5 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 16 | 0.2% |
| Easy Privacy | blocklist | adguard | 53.2K | 89 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.7K | 6 | 0.0% |

---

### AdGuard Base filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 3.7K | **Unique Entries:** 1.2K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 2 | 66.7% |
| transco | allowlist | domain_top | 100 | 38 | 38.0% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 10 | 19.6% |
| AdBlockID | allowlist | domain_adguard | 53 | 10 | 18.9% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 11 | 16.9% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 646 | 74 | 11.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 194 | 8.6% |
| YousList | blocklist | hostname | 624 | 47 | 7.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 40 | 6.2% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 8 | 4.5% |
| WaLLy3K | blocklist | domain | 350 | 11 | 3.1% |
| hufilter | blocklist | hostname | 99 | 3 | 3.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 92 | 2.7% |
| Adaway | blocklist | hostname | 6.5K | 155 | 2.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 268 | 1.4% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 22 | 0.8% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 152 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 108 | 0.7% |
| HaGeZi Pro | blocklist | domain | 406.4K | 398 | 0.1% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 352 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 418 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 16 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 20 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 5 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

---

### AdGuard DNS filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 114.5K | **Unique Entries:** 50.5K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 27.0K | 50.8% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 36.8K | 39.3% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 2 | 28.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 47 | 13.6% |
| abpvn_hosts | blocklist | adguard | 1.1K | 24 | 2.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 30 | 2.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 24 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 13 | 1.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 40 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 1.4K | 1 | 0.1% |

---

### AdGuard DNS filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 177 | **Unique Entries:** 10 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| Easy Privacy | allowlist | domain_adguard | 646 | 9 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 6 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 8 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 7 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 26 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 22 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 5 | 0.0% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 28 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 12 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 25 | 0.0% |

---

### AntiAdBlockFilters

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.7K | **Unique Entries:** 1.7K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 93.5K | 5 | 0.0% |

---

### AntiAdBlockFilters

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 65 | **Unique Entries:** 6 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 5 | 9.8% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 11 | 0.3% |
| YousList | blocklist | hostname | 624 | 2 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 7 | 0.0% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 8 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 5 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |

---

### bigdargon_hostsVN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.8K | **Unique Entries:** 0 | **Target Sources:** 33

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers | blocklist | hostname | 3.4K | 2.1K | 62.4% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 624 | 204 | 32.7% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 13 | 25.5% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 646 | 106 | 16.4% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 401 | 14.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 26 | 14.7% |
| quidsup_notrack-malware | blocklist | domain | 154 | 18 | 11.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 42 | 8.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.2% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 5 | 7.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1.7K | 7.6% |
| AdBlockID | allowlist | domain_adguard | 53 | 4 | 7.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 268 | 7.3% |
| hufilter | blocklist | hostname | 99 | 6 | 6.1% |
| transco | allowlist | domain_top | 100 | 5 | 5.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 12.5K | 4.2% |
| HaGeZi Pro | blocklist | domain | 406.4K | 12.8K | 3.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1.2K | 2.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 7.8K | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 14 | 2.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| Frogeye trackers | blocklist | hostname | 33.4K | 66 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 24 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 264 | 0.0% |

---

### BinaryDefense_Banlist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 1.2K | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 14 | 3.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 302 | 3.4% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 10 | 3.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 468 | 3.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 31 | 1.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 323 | 1.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 181 | 1.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 1.1K | 0.9% |
| Greensnow | blocklist | ipv4 | 5.7K | 43 | 0.8% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 85 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 6 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 8 | 0.0% |

---

### BlockListDE_Strong

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 293 | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 115 | 6.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 97 | 3.8% |
| Greensnow | blocklist | ipv4 | 5.7K | 186 | 3.3% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 251 | 1.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 10 | 0.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 26 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 1 | 0.2% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 285 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 5 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |

---

### Blocklists UT1 Cryptojacking

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.3K | **Unique Entries:** 15.4K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 154 | 4 | 2.6% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 6 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 46 | 0.1% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 201 | 0.1% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 3 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 406.4K | 233 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 261 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 48 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 7 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 48 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 4 | 0.0% |

---

### Blocklists UT1 Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 211.7K | **Unique Entries:** 72.2K | **Target Sources:** 26

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 84.1K | 25.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1.6K | 10.1% |
| quidsup_notrack-malware | blocklist | domain | 154 | 9 | 5.8% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1.2K | 4.9% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 99 | 3.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 19.0K | 2.8% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 7.4K | 2.5% |
| HaGeZi Pro | blocklist | domain | 406.4K | 7.8K | 1.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 572 | 1.3% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 25 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 7 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 13 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 18 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 30 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 15 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 4 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |

---

### Borestad_AbuseIPDB

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 117.3K | **Unique Entries:** 43.1K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 404 | 100.0% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 285 | 97.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 8.5K | 94.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.8K | 92.9% |
| Greensnow | blocklist | ipv4 | 5.7K | 5.2K | 90.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 1.1K | 90.1% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 13.1K | 87.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 13.1K | 87.6% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 9.4K | 80.8% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 221 | 63.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 673 | 26.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 463 | 22.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 216 | 8.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 17 | 3.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 7 | 2.9% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 4.9K | 2.5% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 128 | 1.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 2.1K | 1.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 17 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 25 | 0.0% |

---

### CINSScore_BadGuys_Army

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 15.0K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.6K | 9.0K | 77.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 468 | 39.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 2.0K | 22.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 13.1K | 11.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 1.6K | 8.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 11 | 2.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 49 | 2.6% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 6 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 25 | 1.2% |
| Greensnow | blocklist | ipv4 | 5.7K | 64 | 1.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 1 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 758 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 4 | 0.2% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 13 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 74 | 0.0% |

---

### CJX Annoyance

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 6 | **Unique Entries:** 4 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 406.4K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 1 | 0.0% |

---

### CJX Annoyance

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 1.8K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 4 | 0.0% |

---

### cyberhost_malware-blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.0K | **Unique Entries:** 4.5K | **Target Sources:** 19

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 8.8K | 1.3% |
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1.6K | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 154 | 1 | 0.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 74 | 0.2% |
| HaGeZi Pro | blocklist | domain | 406.4K | 617 | 0.2% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 303 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 24 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 16 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 2 | 0.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 36 | 0.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 1 | 0.0% |

---

### DandelionSprout-Anti-Malware-List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 32.7K | **Unique Entries:** 32.7K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Most Abused TLDs | blocklist | adguard | 423 | 2 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 6 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.3K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 94 | 12.9% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 100 | 0.2% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 6 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 11 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 8 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.3K | **Unique Entries:** 80 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1.2K | 82.7% |
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 4 | 0.0% |

---

### DoH_IP_list

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 731 | **Unique Entries:** 40 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 94 | 3.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 2 | 0.0% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** adguard | **Total Entries:** 746 | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 91.1% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 155 | 44.9% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 27.0K | 23.6% |
| abpvn_hosts | blocklist | adguard | 1.1K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 89 | 0.1% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 646 | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| transco | allowlist | domain_top | 100 | 10 | 10.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 9 | 5.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 74 | 2.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 46 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 11 | 1.7% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 44 | 1.3% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 52 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 106 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 76 | 0.5% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 6 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 27 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 101 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 4 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 125 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 8 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 155 | 0.0% |

---

### EmergingThreats_CompromisedIPs

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 404 | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.6K | 378 | 3.3% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 24 | 1.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 14 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 88 | 0.6% |
| Greensnow | blocklist | ipv4 | 5.7K | 25 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 1 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 24 | 0.3% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 404 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 7 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 12 | 0.0% |

---

### ET_fwip

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.5K | **Unique Entries:** 0 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.5K | 1.5K | 99.3% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.4K | 31.6% |

---

### fabriziosalmi_allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 2.3K | **Unique Entries:** 1.5K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 2 | 66.7% |
| transco | allowlist | domain_top | 100 | 52 | 52.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 74 | 11.4% |
| Easy Privacy | allowlist | domain_adguard | 646 | 46 | 7.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 194 | 5.3% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 2 | 3.9% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 2 | 1.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 9 | 0.7% |
| Adaway | blocklist | hostname | 6.5K | 49 | 0.7% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 11 | 0.4% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 6 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 32 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 33 | 0.2% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 8 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 15 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 10 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 78 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 29 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 95 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |

---

### FabrizioSalmi_DNS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 66 | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 25 | 1.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 4 | 0.0% |

---

### FakeWebshopListHUN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.2K | **Unique Entries:** 8.1K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 41 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 646 | 1 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 18 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 27 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 27 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 24 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |

---

### Firehol_abusers_30d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 196.7K | **Unique Entries:** 162.5K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 275 | 62.9% |
| Firehol_abusers_30d | blocklist | ipv4_cidr_expand | 30.2K | 11.8K | 39.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 505 | 24.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 42 | 17.4% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 52 | 14.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 385 | 14.4% |
| Greensnow | blocklist | ipv4 | 5.7K | 440 | 7.7% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 769 | 5.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 4.9K | 4.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 315 | 3.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 12 | 3.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 15 | 0.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 8 | 0.7% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 80 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 127 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 74 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 8 | 0.3% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 22 | 0.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 262 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 59 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 6 | 0.1% |

---

### Firehol_BitcoinNodes_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.2K | **Unique Entries:** 7.0K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4_cidr_expand | 104 | 47 | 45.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 2 | 0.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 7 | 0.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 23 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 6 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |

---

### Firehol_Botscout_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 437 | **Unique Entries:** 70 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 1 | 0.4% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 5 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 275 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.7K | 3 | 0.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 17 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

---

### Firehol_GPF_Comics

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.1K | **Unique Entries:** 766 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 9 | 2.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 5 | 1.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 463 | 0.4% |
| Greensnow | blocklist | ipv4 | 5.7K | 24 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 47 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 26 | 0.3% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 505 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 25 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 31 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 2 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 13 | 0.0% |

---

### Firehol_level1

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 4.5K | **Unique Entries:** 1.7K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.5K | 1.4K | 93.0% |
| ET_fwip | blocklist | cidr_ipv4 | 1.5K | 1.4K | 92.2% |

---

### Firehol_level2

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 14.9K | **Unique Entries:** 0 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 5.7K | 5.3K | 93.3% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 251 | 85.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 1.5K | 78.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 9.0K | 40.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 88 | 21.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 466 | 18.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 1.6K | 17.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 181 | 15.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 13.1K | 11.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.6K | 10.3% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 32 | 9.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 47 | 2.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 296 | 1.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 5 | 1.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 1.2K | 0.6% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 769 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 5 | 0.2% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 11 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 6 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

---

### Firehol_level3

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 11.6K | **Unique Entries:** 0 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 378 | 93.6% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 268 | 76.4% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.0K | 59.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 9.0K | 50.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 323 | 27.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 1.9K | 20.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 9.4K | 8.0% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 64 | 3.4% |
| Greensnow | blocklist | ipv4 | 5.7K | 176 | 3.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 558 | 2.5% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 6 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 31 | 1.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 1.8K | 0.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 1 | 0.4% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 32 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 5 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 5 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 127 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 4 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

---

### Firehol_SocksProxy_7d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.7K | **Unique Entries:** 1.8K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 138 | 64 | 46.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 55 | 22.8% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 216 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 385 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 34 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 5 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 2 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 10 | 0.0% |

---

### Firehol_SSLProxies_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 241 | **Unique Entries:** 112 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 14 | 7 | 50.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 55 | 2.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 1 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 42 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 6 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |

---

### Frogeye trackers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 33.4K | **Unique Entries:** 16.5K | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 12 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 555 | 3.5% |
| HaGeZi Pro | blocklist | domain | 406.4K | 10.7K | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 119 | 1.8% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 4.7K | 1.6% |
| Easy Privacy | allowlist | domain_adguard | 646 | 8 | 1.2% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 32 | 0.9% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 66 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 7 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 567 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 26 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1 | 0.0% |

---

### GetAdmiral Domains Filter List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 130 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 1.6K | 3.0% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 24 | 0.0% |

---

### Greensnow

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 5.7K | **Unique Entries:** 0 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 293 | 186 | 63.5% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 850 | 44.9% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 5.3K | 35.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 397 | 15.5% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 29 | 8.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 25 | 6.2% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 5.2K | 4.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 43 | 3.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 286 | 3.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 24 | 1.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 176 | 0.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 3 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 64 | 0.4% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 440 | 0.2% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 4 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 63 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 5 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.4K | **Unique Entries:** 233 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1.2K | 92.6% |
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| HaGeZi Pro | blocklist | domain | 406.4K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 2 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.4K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 114.5K | 1 | 0.0% |

---

### HaGeZi Most Abused TLDs

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 423 | **Unique Entries:** 421 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.7K | 2 | 0.0% |

---

### HaGeZi Normal

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 299.9K | **Unique Entries:** 0 | **Target Sources:** 40

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 463 | 96.7% |
| hufilter | blocklist | hostname | 99 | 90 | 90.9% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 3.0K | 89.3% |
| HaGeZi Pro | blocklist | domain | 406.4K | 299.9K | 73.8% |
| quidsup_notrack-malware | blocklist | domain | 154 | 112 | 72.7% |
| YousList | blocklist | hostname | 624 | 437 | 70.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 12.5K | 66.6% |
| Adaway | blocklist | hostname | 6.5K | 4.2K | 64.1% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| WaLLy3K | blocklist | domain | 350 | 145 | 41.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 7.2K | 32.5% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 685 | 25.3% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 12 | 23.5% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 376 | 21.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 9.6K | 21.2% |
| Easy Privacy | allowlist | domain_adguard | 646 | 125 | 19.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2.9K | 18.4% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| Frogeye trackers | blocklist | hostname | 33.4K | 4.7K | 14.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 25 | 14.1% |
| AdBlockID | allowlist | domain_adguard | 53 | 7 | 13.2% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 7 | 10.8% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 352 | 9.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 29.7K | 8.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 958 | 5.3% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1.3K | 5.1% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 7.4K | 3.5% |
| Spam404 | blocklist | domain | 8.1K | 269 | 3.3% |
| transco | allowlist | domain_top | 100 | 3 | 3.0% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 27 | 2.5% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 303 | 1.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 29 | 1.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 201 | 1.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 5 | 0.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 89 | 0.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 3.7K | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 24 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 2 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 44 | 0.0% |

---

### HaGeZi Pro

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 406.4K | **Unique Entries:** 0 | **Target Sources:** 40

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Normal | blocklist | hostname | 299.9K | 299.9K | 100.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 477 | 99.6% |
| hufilter | blocklist | hostname | 99 | 91 | 91.9% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 3.1K | 91.0% |
| Adaway | blocklist | hostname | 6.5K | 5.3K | 81.1% |
| quidsup_notrack-malware | blocklist | domain | 154 | 122 | 79.2% |
| YousList | blocklist | hostname | 624 | 450 | 72.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 12.8K | 68.2% |
| WaLLy3K | blocklist | domain | 350 | 162 | 46.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 7.7K | 34.7% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 917 | 33.8% |
| Frogeye trackers | blocklist | hostname | 33.4K | 10.7K | 32.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 12.4K | 27.4% |
| Easy Privacy | allowlist | domain_adguard | 646 | 155 | 24.0% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 12 | 23.5% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 389 | 22.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.2K | 20.2% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 28 | 15.8% |
| AdBlockID | allowlist | domain_adguard | 53 | 7 | 13.2% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 8 | 12.3% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 398 | 10.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 32.6K | 9.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 5.9% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1.3K | 5.2% |
| transco | allowlist | domain_top | 100 | 4 | 4.0% |
| Spam404 | blocklist | domain | 8.1K | 316 | 3.9% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 617 | 3.9% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 7.8K | 3.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 78 | 3.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 36 | 3.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 158 | 1.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 233 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 7 | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 6.5K | 1.0% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 5 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 27 | 0.3% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 4 | 0.3% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 47 | 0.0% |

---

### HaGeZi Xiaomi Tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 479 | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 26 | 1.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 463 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 42 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 5 | 0.1% |
| HaGeZi Pro | blocklist | domain | 406.4K | 477 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 36 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 3 | 0.0% |

---

### hufilter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 99 | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 646 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 5 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 9 | 0.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 3 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 6 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 91 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 90 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |

---

### Local Allowlist (Domain)

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 3 | **Unique Entries:** 0 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| transco | allowlist | domain_top | 100 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 646 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 2 | 0.1% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1 | 0.0% |

---

### Local Blocklist (AdGuard)

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7 | **Unique Entries:** 5 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 114.5K | 2 | 0.0% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 195.6K | **Unique Entries:** 180.1K | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 4.0K | 53.9% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 4.3K | 35.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 1.8K | 8.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 85 | 7.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 17.7K | 1.2K | 6.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 758 | 5.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 421 | 4.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 241 | 6 | 2.5% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 2.1K | 1.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 7 | 1.7% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 5 | 1.7% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 22 | 1.2% |
| Greensnow | blocklist | ipv4 | 5.7K | 63 | 1.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 15 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| Yoyo AdServers | blocklist | ipv4 | 9.0K | 49 | 0.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 10 | 0.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 23 | 0.3% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 262 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 40 | 0.1% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 676.3K | **Unique Entries:** 627.0K | **Target Sources:** 26

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 8.8K | 55.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 154 | 39 | 25.3% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 374 | 13.8% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 19.0K | 9.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 3.0K | 6.7% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 261 | 1.6% |
| HaGeZi Pro | blocklist | domain | 406.4K | 6.5K | 1.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 264 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 108 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 279 | 1.3% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 3.7K | 1.2% |
| Easy Privacy | allowlist | domain_adguard | 646 | 4 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1.6K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 15 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 35 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 7 | 0.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 5 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 37 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |

---

### malware-filter_phishing-filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 24.8K | **Unique Entries:** 20.3K | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1.2K | 0.6% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 1.3K | 0.4% |
| HaGeZi Pro | blocklist | domain | 406.4K | 1.3K | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 717 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 37 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 13 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 2 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |

---

### MyIP_MS_Blocklist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 351 | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.6K | 268 | 2.3% |
| Greensnow | blocklist | ipv4 | 5.7K | 29 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 9 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 20 | 0.2% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 221 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4_cidr_expand | 30.2K | 52 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 32 | 0.2% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |

---

### Public_DNS4

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 62.6K | **Unique Entries:** 61.7K | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 100 | 4.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 34 | 1.3% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 1 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 437 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.7K | 3 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 2 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 59 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 5 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 40 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 25 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 2 | 0.0% |

---

### quidsup_notrack-malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 154 | **Unique Entries:** 0 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 10 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 75 | 0.2% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 18 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 39 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 112 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 122 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 26 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |

---

### quidsup_notrack-tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 15.7K | **Unique Entries:** 3.7K | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers | blocklist | hostname | 3.4K | 589 | 17.3% |
| Easy Privacy | allowlist | domain_adguard | 646 | 76 | 11.8% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 449 | 6.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 1.3K | 6.8% |
| transco | allowlist | domain_top | 100 | 6 | 6.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 9 | 5.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 970 | 4.4% |
| hufilter | blocklist | hostname | 99 | 4 | 4.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 108 | 3.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 12 | 2.5% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 61 | 2.3% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 12 | 1.9% |
| Frogeye trackers | blocklist | hostname | 33.4K | 555 | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 2.9K | 1.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 3.2K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 335 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1.3K | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 15 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 35 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |

---

### RedDragonWebDesign_block-everything

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 652 | **Unique Entries:** 649 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 93.5K | 3 | 0.0% |

---

### Rutgers_DROP

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 1.9K | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 293 | 115 | 39.2% |
| Greensnow | blocklist | ipv4 | 5.7K | 850 | 14.9% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 1.5K | 9.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 24 | 5.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 101 | 3.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 231 | 2.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 31 | 2.6% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 1.8K | 1.5% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 64 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 49 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 22 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 2 | 0.0% |

---

### ScriptzTeam_BadIPS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.6K | **Unique Entries:** 793 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 293 | 97 | 33.1% |
| Greensnow | blocklist | ipv4 | 5.7K | 397 | 6.9% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 101 | 5.3% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 466 | 3.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 673 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 6 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 1 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 15 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 4 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |

---

### Sentinel_Greylist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 9.0K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 302 | 25.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2.0K | 13.6% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 231 | 12.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 1.6K | 10.5% |
| BlockListDE_Strong | blocklist | ipv4 | 293 | 26 | 8.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 1.9K | 8.3% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 8.5K | 7.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 24 | 5.9% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 20 | 5.7% |
| Greensnow | blocklist | ipv4 | 5.7K | 286 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 26 | 1.3% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 315 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 421 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 6 | 0.2% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 16 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 1.0K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 93.5K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 13 | 0.0% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.1K | **Unique Entries:** 970 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 23 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 27 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 36 | 0.0% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 1 | 0.0% |

---

### ShadowWhisperer_Allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 647 | **Unique Entries:** 407 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| transco | allowlist | domain_top | 100 | 4 | 4.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 74 | 3.3% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Easy Privacy | allowlist | domain_adguard | 646 | 11 | 1.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 3 | 1.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 40 | 1.1% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 14 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 32 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 3 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 7 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1 | 0.0% |

---

### ShadowWhisperer_BlockLists Ads

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 22.1K | **Unique Entries:** 711 | **Target Sources:** 28

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers | blocklist | hostname | 3.4K | 847 | 24.9% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 624 | 85 | 13.6% |
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 1.7K | 8.9% |
| AdBlockID | allowlist | domain_adguard | 53 | 4 | 7.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 970 | 6.2% |
| Adaway | blocklist | hostname | 6.5K | 407 | 6.2% |
| quidsup_notrack-malware | blocklist | domain | 154 | 7 | 4.5% |
| Easy Privacy | allowlist | domain_adguard | 646 | 27 | 4.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 152 | 4.2% |
| transco | allowlist | domain_top | 100 | 3 | 3.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 5 | 2.8% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 7.2K | 2.4% |
| HaGeZi Pro | blocklist | domain | 406.4K | 7.7K | 1.9% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 46 | 1.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1.9K | 0.6% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 3 | 0.6% |
| Frogeye trackers | blocklist | hostname | 33.4K | 26 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 279 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 12 | 0.0% |

---

### ShadowWhisperer_BlockLists Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 45.4K | **Unique Entries:** 17.3K | **Target Sources:** 24

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 154 | 75 | 48.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 1.2K | 6.2% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 9.6K | 3.2% |
| HaGeZi Pro | blocklist | domain | 406.4K | 12.4K | 3.1% |
| YousList | blocklist | hostname | 624 | 19 | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 335 | 2.1% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 66 | 1.9% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 30 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 45 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 74 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 3.0K | 0.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 16 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 572 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 46 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 575 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 15 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 646 | 1 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 13 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 8 | 0.0% |

---

### ShadowWhisperer_BlockLists Scam

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 10.8K | **Unique Entries:** 9.4K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 41 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 34 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 991 | 0.3% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 89 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 7 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 25 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 158 | 0.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 4 | 0.0% |

---

### Spam404

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.1K | **Unique Entries:** 7.3K | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 154 | 2 | 1.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 34 | 0.3% |
| HaGeZi Pro | blocklist | domain | 406.4K | 316 | 0.1% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 269 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 45 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 64 | 0.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 108 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 13 | 0.0% |

---

### spamhaus_drop

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.5K | **Unique Entries:** 0 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.5K | 1.5K | 98.6% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.4K | 31.6% |

---

### Stamparm_Blackbook

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.1K | **Unique Entries:** 0 | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 17.6K | 8.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 5.0K | 0.7% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 958 | 0.3% |
| HaGeZi Pro | blocklist | domain | 406.4K | 1.1K | 0.3% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 3 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 16 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 31 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |

---

### StevenBlack_Adhoc_list

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 2.7K | **Unique Entries:** 0 | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 26 | 5.4% |
| WaLLy3K | blocklist | domain | 350 | 15 | 4.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 401 | 2.1% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| Adaway | blocklist | hostname | 6.5K | 133 | 2.0% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| YousList | blocklist | hostname | 624 | 8 | 1.3% |
| quidsup_notrack-malware | blocklist | domain | 154 | 2 | 1.3% |
| Easy Privacy | allowlist | domain_adguard | 646 | 6 | 0.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 2.7K | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 1 | 0.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 22 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 11 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 61 | 0.4% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 10 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 46 | 0.2% |
| HaGeZi Pro | blocklist | domain | 406.4K | 917 | 0.2% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 685 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 30 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 374 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 7 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 99 | 0.0% |

---

### StevenBlack_Fake_Gambling_Porn

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 336.5K | **Unique Entries:** 160.7K | **Target Sources:** 39

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 2.7K | 100.0% |
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.8% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 3.4K | 99.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 7.8K | 41.5% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 84.1K | 39.7% |
| YousList | blocklist | hostname | 624 | 243 | 38.9% |
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| WaLLy3K | blocklist | domain | 350 | 86 | 24.6% |
| abpvn_hosts | allowlist | domain_adguard | 51 | 12 | 23.5% |
| quidsup_notrack-malware | blocklist | domain | 154 | 26 | 16.9% |
| Easy Privacy | allowlist | domain_adguard | 646 | 101 | 15.6% |
| hufilter | blocklist | hostname | 99 | 14 | 14.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 22 | 12.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 418 | 11.4% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 29.7K | 9.9% |
| AdBlockID | allowlist | domain_adguard | 53 | 5 | 9.4% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 6 | 9.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 991 | 9.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1.9K | 8.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| HaGeZi Pro | blocklist | domain | 406.4K | 32.6K | 8.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 36 | 7.5% |
| transco | allowlist | domain_top | 100 | 5 | 5.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 32 | 4.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 95 | 4.2% |
| malware-filter_phishing-filter | blocklist | hostname | 24.8K | 717 | 2.9% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 23 | 2.2% |
| Frogeye trackers | blocklist | hostname | 33.4K | 567 | 1.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 575 | 1.3% |
| Spam404 | blocklist | domain | 8.1K | 64 | 0.8% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 27 | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 1.6K | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 36 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 31 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |

---

### transco

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 100 | **Unique Entries:** 0 | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 52 | 2.3% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Easy Privacy | allowlist | domain_adguard | 646 | 10 | 1.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 38 | 1.0% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 4 | 0.6% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 5 | 0.1% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 3 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 3 | 0.0% |

---

### Ukrainian Ad Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.4K | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 93.5K | 55 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 30 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 2 | 0.0% |

---

### Ukrainian Privacy Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 345 | **Unique Entries:** 138 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 155 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 47 | 0.0% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 1 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 2 | 0.0% |

---

### Ukrainian Security Filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.7K | **Unique Entries:** 911 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Normal | blocklist | hostname | 299.9K | 376 | 0.1% |
| HaGeZi Pro | blocklist | domain | 406.4K | 389 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 46 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 13 | 0.0% |

---

### USOM-Blocklists

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 12.1K | **Unique Entries:** 7.2K | **Target Sources:** 19

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 279 | 3.7% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 4.3K | 2.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 1.2K | 6 | 0.5% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 351 | 1 | 0.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 4 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 404 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 16 | 0.2% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 128 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 13 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.7K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 22.3K | 32 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 11 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 1.9K | 2 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.7K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 80 | 0.0% |
| Yoyo AdServers | blocklist | ipv4 | 9.0K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.5K | **Unique Entries:** 3.1K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 279 | 2.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 4.0K | 2.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.7K | 4 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 196.7K | 22 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 6 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 117.3K | 17 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.6K | 4 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.0K | 1 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |

---

### WaLLy3K

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 350 | **Unique Entries:** 0 | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 2 | 3.9% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| transco | allowlist | domain_top | 100 | 2 | 2.0% |
| YousList | blocklist | hostname | 624 | 9 | 1.4% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| Easy Privacy | allowlist | domain_adguard | 646 | 5 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 1 | 0.6% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 15 | 0.6% |
| quidsup_notrack-malware | blocklist | domain | 154 | 1 | 0.6% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 19 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 85 | 0.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 11 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 53 | 0.2% |
| HaGeZi Pro | blocklist | domain | 406.4K | 162 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 2 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 145 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 12 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 86 | 0.0% |

---

### YousList

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 624 | **Unique Entries:** 0 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 5 | 9.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 47 | 1.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 204 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 22 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 85 | 0.4% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 8 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 646 | 1 | 0.2% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 437 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 243 | 0.1% |
| HaGeZi Pro | blocklist | domain | 406.4K | 450 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 22 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 19 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 5 | 0.0% |

---

### YousList-AdGuard

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7.3K | **Unique Entries:** 7.3K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| Easy Privacy | blocklist | adguard | 53.2K | 10 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 93.5K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 114.5K | 40 | 0.0% |

---

### YousList-AdGuard

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 12 | **Unique Entries:** 0 | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 1 | 2.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 2 | 0.0% |

---

### youtube_GoodbyeAds

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 97.6K | **Unique Entries:** 97.3K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 51 | 5 | 9.8% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdBlockID | allowlist | domain_adguard | 53 | 3 | 5.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| transco | allowlist | domain_top | 100 | 3 | 3.0% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 2 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 20 | 0.5% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 646 | 2 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 52 | 0.3% |
| Yoyo Adservers | blocklist | hostname | 3.4K | 8 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 47 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 74 | 0.0% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 44 | 0.0% |

---

### Yoyo Adservers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 3.4K | **Unique Entries:** 0 | **Target Sources:** 31

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.8K | 2.1K | 11.3% |
| Easy Privacy | allowlist | domain_adguard | 646 | 44 | 6.8% |
| quidsup_notrack-malware | blocklist | domain | 154 | 10 | 6.5% |
| WaLLy3K | blocklist | domain | 350 | 19 | 5.4% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| transco | allowlist | domain_top | 100 | 5 | 5.0% |
| Adaway | blocklist | hostname | 6.5K | 254 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 847 | 3.8% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 589 | 3.7% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 6 | 3.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 92 | 2.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 11 | 1.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| HaGeZi Normal | blocklist | hostname | 299.9K | 3.0K | 1.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 336.5K | 3.4K | 1.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 5 | 1.0% |
| HaGeZi Pro | blocklist | domain | 406.4K | 3.1K | 0.8% |
| StevenBlack_Adhoc_list | blocklist | hostname | 2.7K | 10 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 66 | 0.1% |
| Frogeye trackers | blocklist | hostname | 33.4K | 32 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 676.3K | 15 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |

---

### Yoyo AdServers

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 9.0K | **Unique Entries:** 8.9K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 195.6K | 49 | 0.0% |
| USOM-Blocklists | blocklist | ipv4 | 12.1K | 1 | 0.0% |

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources. High overlap percentages may indicate redundant sources, while low overlap percentages suggest unique content.

**Note:** Overlap percentages are calculated as: (overlap_count / source_total_count) × 100

