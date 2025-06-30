# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2025-07-19 14:27:39 UTC

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 95 |
| Total Entries Analyzed | 3.0M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 82 |
| allowlist | 13 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 17 |
| cidr_ipv4 | 3 |
| domain | 45 |
| ipv4 | 30 |

## Detailed Source Analysis

### abpvn_hosts

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 50 | **Unique Entries:** 0 | **Target Sources:** 17

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
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 8 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 13 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |

---

### abpvn_hosts

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 1.0K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 10 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 24 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 2 | 0.0% |

---

### Adaway

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 6.5K | **Unique Entries:** 0 | **Target Sources:** 29

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList | blocklist | hostname | 624 | 111 | 17.8% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 8 | 16.0% |
| WaLLy3K | blocklist | domain | 350 | 54 | 15.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2.7K | 14.4% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| Easy Privacy | allowlist | domain_adguard | 650 | 52 | 8.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 255 | 7.5% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 4 | 6.2% |
| AdBlockID | allowlist | domain_adguard | 53 | 3 | 5.7% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 155 | 4.2% |
| tranco | allowlist | domain_top | 100 | 4 | 4.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 7 | 4.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 402 | 3.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 449 | 2.9% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 14 | 2.9% |
| quidsup_notrack-malware | blocklist | domain | 152 | 4 | 2.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 49 | 2.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 6.5K | 2.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 407 | 1.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 10 | 1.5% |
| HaGeZi Pro | blocklist | domain | 419.7K | 5.3K | 1.3% |
| Frogeye trackers | blocklist | hostname | 33.4K | 115 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 15 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 28 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 26 | 0.0% |

---

### AdBlockID

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 3.8K | **Unique Entries:** 3.8K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 91.3K | 25 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 2 | 0.0% |

---

### AdBlockID

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 53 | **Unique Entries:** 0 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 2 | 4.0% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 10 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 7 | 0.0% |

---

### AdGuard Base filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 91.3K | **Unique Entries:** 56.6K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 120.8K | 34.5K | 28.5% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 54 | 3.8% |
| abpvn_hosts | blocklist | adguard | 1.1K | 10 | 0.9% |
| AdBlockID | blocklist | adguard | 3.8K | 25 | 0.7% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 3 | 0.5% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 5 | 0.3% |
| Easy Privacy | blocklist | adguard | 53.2K | 89 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 16 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 2 | 0.1% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 6 | 0.0% |

---

### AdGuard Base filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 3.7K | **Unique Entries:** 1.4K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 2 | 66.7% |
| tranco | allowlist | domain_top | 100 | 38 | 38.0% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 10 | 20.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 10 | 18.9% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 11 | 16.9% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 650 | 74 | 11.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 194 | 8.6% |
| YousList | blocklist | hostname | 624 | 47 | 7.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 40 | 6.2% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 8 | 4.5% |
| WaLLy3K | blocklist | domain | 350 | 11 | 3.1% |
| hufilter | blocklist | hostname | 99 | 3 | 3.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 92 | 2.7% |
| Adaway | blocklist | hostname | 6.5K | 155 | 2.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 268 | 1.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 101 | 0.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 108 | 0.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 152 | 0.7% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 12 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 418 | 0.1% |
| HaGeZi Pro | blocklist | domain | 419.7K | 399 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 5 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 7 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 20 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 16 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

---

### AdGuard DNS filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 120.8K | **Unique Entries:** 59.2K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 27.0K | 50.7% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 34.5K | 37.8% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 2 | 28.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 50 | 14.5% |
| abpvn_hosts | blocklist | adguard | 1.1K | 24 | 2.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 24 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 13 | 1.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 40 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 1.4K | 1 | 0.1% |

---

### AdGuard DNS filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 177 | **Unique Entries:** 30 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 1 | 2.0% |
| Easy Privacy | allowlist | domain_adguard | 650 | 9 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 8 | 0.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 6 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 7 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 26 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 7 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 22 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 27 | 0.0% |

---

### AntiAdBlockFilters

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 65 | **Unique Entries:** 12 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 5 | 10.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 2 | 0.3% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 11 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 6 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 8 | 0.0% |

---

### AntiAdBlockFilters

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.7K | **Unique Entries:** 1.7K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 91.3K | 5 | 0.0% |

---

### bigdargon_hostsVN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.9K | **Unique Entries:** 0 | **Target Sources:** 33

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 2.1K | 62.2% |
| Adaway | blocklist | hostname | 6.5K | 2.7K | 41.5% |
| YousList | blocklist | hostname | 624 | 204 | 32.7% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 13 | 26.0% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 650 | 107 | 16.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 26 | 14.7% |
| quidsup_notrack-malware | blocklist | domain | 152 | 17 | 11.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1.1K | 9.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 42 | 8.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.2% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 5 | 7.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1.7K | 7.6% |
| AdBlockID | allowlist | domain_adguard | 53 | 4 | 7.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 268 | 7.3% |
| hufilter | blocklist | hostname | 99 | 6 | 6.1% |
| tranco | allowlist | domain_top | 100 | 5 | 5.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 12.9K | 3.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1.2K | 2.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 7.8K | 2.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 14 | 2.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| Frogeye trackers | blocklist | hostname | 33.4K | 68 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 24 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 4 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 264 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 30 | 0.0% |

---

### BinaryDefense_Banlist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.8K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 451 | 8.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 743 | 8.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 33 | 8.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 855 | 4.0% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 11 | 3.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 582 | 3.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 2.5K | 2.5% |
| Greensnow | blocklist | ipv4 | 5.2K | 103 | 2.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 39 | 1.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 462 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 3 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 12 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 13 | 0.1% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 12 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

---

### BlockListDE_Strong

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 298 | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 113 | 5.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 96 | 3.7% |
| Greensnow | blocklist | ipv4 | 5.2K | 181 | 3.5% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 254 | 1.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 11 | 0.4% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 288 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 1 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 13 | 0.2% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 4 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |

---

### Blocklists UT1 Cryptojacking

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.3K | **Unique Entries:** 15.6K | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 152 | 4 | 2.6% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 6 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 24 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 46 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 419.7K | 237 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 48 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 7 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 48 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 261 | 0.0% |

---

### Blocklists UT1 Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 211.7K | **Unique Entries:** 84.9K | **Target Sources:** 29

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 78.4K | 25.7% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 405 | 14.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1.6K | 9.9% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 992 | 8.4% |
| quidsup_notrack-malware | blocklist | domain | 152 | 9 | 5.9% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 46 | 4.7% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1.1K | 4.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 8 | 3.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 19.2K | 2.8% |
| HaGeZi Pro | blocklist | domain | 419.7K | 6.7K | 1.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 572 | 1.3% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| Spam404 | blocklist | domain | 8.1K | 13 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 7 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 25 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 30 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 18 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 15 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 4 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 12 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 3 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 5 | 0.0% |

---

### Borestad_AbuseIPDB

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 99.9K | **Unique Entries:** 22.2K | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 379 | 100.0% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 288 | 96.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 8.0K | 93.7% |
| Greensnow | blocklist | ipv4 | 5.2K | 4.7K | 90.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 2.5K | 89.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.8K | 88.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 13.0K | 86.7% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 9.3K | 79.8% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 11.6K | 77.8% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2.6K | 50.4% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 182 | 36.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 645 | 25.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 460 | 21.8% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 411 | 19.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 157 | 5.7% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 1.1K | 4.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 16 | 3.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 4.5K | 2.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 3 | 1.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 3.0K | 1.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 125 | 1.0% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 1 | 1.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 17 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 7 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 20 | 0.0% |

---

### CINSScore_BadGuys_Army

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 15.0K | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.7K | 9.7K | 83.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 1.1K | 39.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 1.9K | 22.3% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 865 | 16.9% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 13.0K | 13.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 1.2K | 7.9% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 75 | 3.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 11 | 2.9% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 6 | 2.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 96 | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 25 | 1.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 1.3K | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 41 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 12 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 66 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |

---

### CJX Annoyance

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 6 | **Unique Entries:** 5 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 419.7K | 1 | 0.0% |

---

### CJX Annoyance

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 1.8K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 4 | 0.0% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 103 | **Unique Entries:** 92 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 7 | 0.0% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 982 | **Unique Entries:** 795 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 47 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 46 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 88 | 0.0% |

---

### cyberhost_malware-blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.4K | **Unique Entries:** 4.7K | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 65 | 2.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 9.1K | 1.3% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1.6K | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 152 | 1 | 0.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 74 | 0.2% |
| HaGeZi Pro | blocklist | domain | 419.7K | 706 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 24 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 16 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 39 | 0.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |

---

### Dan Pollock's List

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 11.8K | **Unique Entries:** 0 | **Target Sources:** 32

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList | blocklist | hostname | 624 | 108 | 17.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 441 | 12.9% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 5 | 10.0% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| Adaway | blocklist | hostname | 6.5K | 402 | 6.1% |
| WaLLy3K | blocklist | domain | 350 | 20 | 5.7% |
| AdBlockID | allowlist | domain_adguard | 53 | 3 | 5.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.1K | 5.6% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 7 | 4.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 11.7K | 3.8% |
| Easy Privacy | allowlist | domain_adguard | 650 | 23 | 3.5% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 101 | 2.8% |
| quidsup_notrack-malware | blocklist | domain | 152 | 4 | 2.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 325 | 1.5% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 202 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 8 | 1.2% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 3.1K | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 992 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 20 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 92 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 50 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 24 | 0.1% |
| Frogeye trackers | blocklist | hostname | 33.4K | 45 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 686 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 5 | 0.0% |

---

### DandelionSprout-Anti-Malware-List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 32.6K | **Unique Entries:** 32.6K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Most Abused TLDs | blocklist | adguard | 422 | 2 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 6 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.3K | **Unique Entries:** 80 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1.2K | 81.8% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| HaGeZi Pro | blocklist | domain | 419.7K | 4 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.3K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 94 | 12.9% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 100 | 0.2% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 5 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 8 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 11 | 0.0% |

---

### DoH_IP_list

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 731 | **Unique Entries:** 40 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 94 | 3.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 2 | 0.0% |

---

### DShield

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 5.1K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 5.1K | 32.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 5.1K | 24.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 451 | 15.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 906 | 10.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 865 | 5.8% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 2.6K | 2.6% |
| Greensnow | blocklist | ipv4 | 5.2K | 47 | 0.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 1.8K | 0.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 2 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 8 | 0.4% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 5 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 24 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 7 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 25 | 0.0% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 650 | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| tranco | allowlist | domain_top | 100 | 10 | 10.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 9 | 5.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 46 | 2.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 74 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 11 | 1.7% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 44 | 1.3% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Adaway | blocklist | hostname | 6.5K | 52 | 0.8% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 107 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 76 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 23 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 27 | 0.1% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 3 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 156 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 101 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 8 | 0.0% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** adguard | **Total Entries:** 750 | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 91.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 156 | 45.2% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 27.0K | 22.3% |
| abpvn_hosts | blocklist | adguard | 1.1K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 89 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |

---

### EmergingThreats_CompromisedIPs

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 379 | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.7K | 362 | 3.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 30 | 1.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 33 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 93 | 0.6% |
| Greensnow | blocklist | ipv4 | 5.2K | 27 | 0.5% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 379 | 0.4% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 1 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 18 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 9 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 13 | 0.0% |

---

### ET_fwip

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.6K | **Unique Entries:** 0 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.5K | 1.5K | 99.9% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.4K | 32.1% |

---

### fabriziosalmi_allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 2.3K | **Unique Entries:** 1.5K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 2 | 66.7% |
| tranco | allowlist | domain_top | 100 | 50 | 50.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 74 | 11.4% |
| Easy Privacy | allowlist | domain_adguard | 650 | 46 | 7.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 194 | 5.3% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 2 | 4.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 2 | 1.1% |
| Adaway | blocklist | hostname | 6.5K | 49 | 0.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 9 | 0.7% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 20 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 6 | 0.4% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 8 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 32 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 33 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 15 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 95 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 10 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 4 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 78 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 10 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |

---

### FabrizioSalmi_DNS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 66 | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 25 | 1.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 4 | 0.0% |

---

### FakeWebshopListHUN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.2K | **Unique Entries:** 8.1K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 41 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 650 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 419.7K | 27 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 25 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 18 | 0.0% |

---

### Firehol_abusers_30d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 194.6K | **Unique Entries:** 159.5K | **Target Sources:** 24

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 1.6K | 78.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 329 | 61.8% |
| Firehol_abusers_30d | blocklist | ipv4_cidr_expand | 28.9K | 11.4K | 39.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 56 | 27.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 524 | 24.8% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 61 | 12.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 306 | 11.1% |
| Greensnow | blocklist | ipv4 | 5.2K | 290 | 5.6% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 4.5K | 4.5% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 637 | 4.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 315 | 3.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 13 | 3.4% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 100 | 0.9% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 13 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 74 | 0.6% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 25 | 0.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 66 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 12 | 0.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 8 | 0.3% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 21 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 35 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 59 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 256 | 0.1% |

---

### Firehol_BitcoinNodes_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.2K | **Unique Entries:** 7.1K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4_cidr_expand | 98 | 43 | 43.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 2 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 21 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 7 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 6 | 0.0% |

---

### Firehol_Botscout_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 532 | **Unique Entries:** 32 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 67 | 3.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 1 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 8 | 0.4% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 329 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 4 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 1 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 16 | 0.0% |

---

### Firehol_GPF_Comics

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.1K | **Unique Entries:** 638 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 57 | 2.7% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 9 | 1.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 8 | 1.5% |
| Greensnow | blocklist | ipv4 | 5.2K | 44 | 0.9% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 100 | 0.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 460 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 1 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 24 | 0.3% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 524 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 25 | 0.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 8 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 27 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 3 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 12 | 0.0% |

---

### Firehol_level1

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 4.5K | **Unique Entries:** 1.6K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| spamhaus_drop | blocklist | cidr_ipv4 | 1.5K | 1.4K | 93.1% |
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.4K | 92.6% |

---

### Firehol_level2

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 14.9K | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 4.8K | 92.3% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 254 | 85.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 1.5K | 74.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 7.2K | 33.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 93 | 24.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 582 | 20.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 416 | 16.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 1.3K | 15.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 11.6K | 11.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 8.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 100 | 4.7% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 13 | 2.6% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 39 | 1.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 257 | 1.6% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 2.4K | 1.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 4 | 0.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 8 | 0.3% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 637 | 0.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 53 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 15 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

---

### Firehol_level3

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 11.7K | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DShield | blocklist | ipv4_range_expand | 5.1K | 5.1K | 100.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 362 | 95.5% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 411 | 82.5% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 9.7K | 64.7% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 7.2K | 46.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 855 | 30.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 1.9K | 22.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 9.3K | 9.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 96 | 4.6% |
| Greensnow | blocklist | ipv4 | 5.2K | 148 | 2.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 460 | 2.2% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 6 | 2.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 3.6K | 1.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 27 | 1.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 85 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 7 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 22 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 4 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 5 | 0.1% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 100 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

---

### Firehol_SocksProxy_7d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.8K | **Unique Entries:** 2.1K | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 106 | 48 | 45.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 37 | 18.3% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 9 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 2 | 0.4% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 306 | 0.2% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 157 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 34 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 8 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 5 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 7 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 2 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 11 | 0.0% |

---

### Firehol_SSLProxies_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 20 | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 20 | 10 | 50.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 37 | 1.3% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 4 | 0.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 2 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 3 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 56 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

---

### Frogeye trackers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 33.4K | **Unique Entries:** 21.1K | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 12 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 558 | 3.5% |
| HaGeZi Pro | blocklist | domain | 419.7K | 10.8K | 2.6% |
| Adaway | blocklist | hostname | 6.5K | 115 | 1.8% |
| Easy Privacy | allowlist | domain_adguard | 650 | 8 | 1.2% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 32 | 0.9% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 68 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 45 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 566 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 4 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 26 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1 | 0.0% |

---

### GetAdmiral Domains Filter List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 127 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 1.6K | 3.0% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 24 | 0.0% |

---

### GlobalAntiScamOrg-blocklist-domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 11.0K | **Unique Entries:** 11.0K | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 30 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 13 | 0.0% |

---

### Greensnow

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 5.2K | **Unique Entries:** 0 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 298 | 181 | 60.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 784 | 37.9% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 4.8K | 32.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 354 | 13.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 27 | 7.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 4.7K | 4.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 103 | 3.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 279 | 3.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 44 | 2.1% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 6 | 1.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 47 | 0.9% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 14 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 148 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 96 | 0.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 1 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 290 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 8 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 109 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.4K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 120.8K | 1 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.4K | **Unique Entries:** 249 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1.2K | 92.6% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 5 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1 | 0.0% |

---

### HaGeZi Most Abused TLDs

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 422 | **Unique Entries:** 420 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 2 | 0.0% |

---

### HaGeZi Pro

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 419.7K | **Unique Entries:** 308.2K | **Target Sources:** 43

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 477 | 99.6% |
| hufilter | blocklist | hostname | 99 | 91 | 91.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3.1K | 91.2% |
| quidsup_notrack-malware | blocklist | domain | 152 | 124 | 81.6% |
| Adaway | blocklist | hostname | 6.5K | 5.3K | 81.3% |
| YousList | blocklist | hostname | 624 | 451 | 72.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 12.9K | 68.5% |
| WaLLy3K | blocklist | domain | 350 | 163 | 46.6% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 7.7K | 34.9% |
| Frogeye trackers | blocklist | hostname | 33.4K | 10.8K | 32.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 12.7K | 27.9% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 3.1K | 26.0% |
| Easy Privacy | allowlist | domain_adguard | 650 | 156 | 24.0% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 12 | 24.0% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 393 | 22.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.2K | 20.2% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 27 | 15.3% |
| AdBlockID | allowlist | domain_adguard | 53 | 7 | 13.2% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 8 | 12.3% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 399 | 10.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 32.5K | 10.6% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 278 | 9.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 6.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1.4K | 5.8% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 706 | 4.3% |
| Spam404 | blocklist | domain | 8.1K | 338 | 4.2% |
| tranco | allowlist | domain_top | 100 | 4 | 4.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 78 | 3.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 36 | 3.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 8 | 3.1% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 6.7K | 3.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 199 | 1.8% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 237 | 1.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 7 | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 6.7K | 1.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 4 | 0.4% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 4 | 0.3% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 30 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 27 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 5 | 0.3% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 47 | 0.0% |

---

### HaGeZi Xiaomi Tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 479 | **Unique Entries:** 0 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 42 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 14 | 0.2% |
| HaGeZi Pro | blocklist | domain | 419.7K | 477 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 36 | 0.0% |

---

### hufilter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 99 | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 650 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 9 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 5 | 0.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 3 | 0.1% |
| HaGeZi Pro | blocklist | domain | 419.7K | 91 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 14 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 6 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |

---

### Local Allowlist (Domain)

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 3 | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 650 | 1 | 0.2% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 2 | 0.1% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 2 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |

---

### Local Blocklist (AdGuard)

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7 | **Unique Entries:** 5 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 120.8K | 2 | 0.0% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 679.8K | **Unique Entries:** 632.5K | **Target Sources:** 28

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 9.1K | 55.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 152 | 39 | 25.7% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 629 | 21.7% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 19.2K | 9.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 88 | 9.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 3.0K | 6.7% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 686 | 5.8% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 261 | 1.6% |
| HaGeZi Pro | blocklist | domain | 419.7K | 6.7K | 1.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 264 | 1.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 279 | 1.3% |
| Spam404 | blocklist | domain | 8.1K | 108 | 1.3% |
| Easy Privacy | allowlist | domain_adguard | 650 | 4 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1.6K | 0.5% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 15 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 26 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 35 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 41 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 13 | 0.1% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 5 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 7 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 198.6K | **Unique Entries:** 175.0K | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 4.0K | 53.9% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1.8K | 36.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 4.3K | 35.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 3.6K | 16.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 462 | 16.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 2.4K | 15.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 842 | 9.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 8.4% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 7 | 6.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 860 | 3.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 3.0K | 3.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 9 | 2.4% |
| Greensnow | blocklist | ipv4 | 5.2K | 109 | 2.1% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 4 | 1.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 2 | 1.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 18 | 0.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 12 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 15 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 11 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 9.0K | 49 | 0.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 11 | 0.4% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 21 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 40 | 0.1% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 256 | 0.1% |

---

### malware-filter_phishing-filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 25.0K | **Unique Entries:** 22.2K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 75 | 28.8% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1.1K | 0.5% |
| HaGeZi Pro | blocklist | domain | 419.7K | 1.4K | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 2 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 41 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 12 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 108 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |

---

### MyIP_MS_Blocklist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 498 | **Unique Entries:** 0 | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 11.7K | 411 | 3.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 9 | 0.4% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 182 | 0.2% |
| Firehol_abusers_30d | blocklist | ipv4_cidr_expand | 28.9K | 61 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 13 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.2K | 6 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 10 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 1 | 0.0% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 1 | 0.0% |

---

### OpenPhish_Feed

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 260 | **Unique Entries:** 165 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 75 | 0.3% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 8 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 8 | 0.0% |

---

### Public_DNS4

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 62.6K | **Unique Entries:** 61.7K | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 100 | 4.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 34 | 1.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 2 | 1.0% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 1 | 0.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 3 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 9 | 0.0% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 40 | 0.0% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 20 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 1 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 59 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 2 | 0.0% |

---

### quidsup_notrack-malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 152 | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 10 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 74 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 17 | 0.1% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 26 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 7 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 124 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 39 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 4 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 9 | 0.0% |

---

### quidsup_notrack-tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 15.7K | **Unique Entries:** 6.5K | **Target Sources:** 26

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 590 | 17.3% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| Easy Privacy | allowlist | domain_adguard | 650 | 76 | 11.7% |
| Adaway | blocklist | hostname | 6.5K | 449 | 6.9% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.3K | 6.8% |
| tranco | allowlist | domain_top | 100 | 6 | 6.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 9 | 5.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 969 | 4.4% |
| hufilter | blocklist | hostname | 99 | 4 | 4.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 108 | 3.0% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 12 | 2.5% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 1 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 12 | 1.9% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 202 | 1.7% |
| Frogeye trackers | blocklist | hostname | 33.4K | 558 | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| HaGeZi Pro | blocklist | domain | 419.7K | 3.2K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 332 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1.3K | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 35 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 15 | 0.0% |

---

### RedDragonWebDesign_block-everything

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 652 | **Unique Entries:** 649 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 91.3K | 3 | 0.0% |

---

### Rutgers_DROP

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.1K | **Unique Entries:** 0 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 298 | 113 | 37.9% |
| Greensnow | blocklist | ipv4 | 5.2K | 784 | 15.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 1.5K | 10.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 30 | 7.9% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 97 | 3.8% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 150 | 1.8% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 1.8K | 1.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 39 | 1.4% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 96 | 0.8% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 75 | 0.5% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 2 | 0.4% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 5 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 18 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 13 | 0.0% |

---

### Sblam_Blocklist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.1K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_Botscout_1d | blocklist | ipv4 | 532 | 67 | 12.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 57 | 2.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 202 | 4 | 2.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 1.6K | 0.8% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 411 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 9 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.2K | 14 | 0.3% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 39 | 0.3% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 8 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 14 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 4 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 11 | 0.0% |

---

### ScriptzTeam_BadIPS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.6K | **Unique Entries:** 923 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 298 | 96 | 32.2% |
| Greensnow | blocklist | ipv4 | 5.2K | 354 | 6.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 97 | 4.7% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 416 | 2.8% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 645 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 5 | 0.1% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 15 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 1 | 0.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 1 | 0.0% |

---

### Sentinel_Greylist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 8.5K | **Unique Entries:** 0 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 743 | 26.2% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 906 | 17.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.9K | 12.6% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 1.9K | 8.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 1.3K | 8.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 8.0K | 8.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 150 | 7.3% |
| Greensnow | blocklist | ipv4 | 5.2K | 279 | 5.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 18 | 4.7% |
| BlockListDE_Strong | blocklist | ipv4 | 298 | 13 | 4.4% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 10 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 24 | 1.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 147 | 0.6% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 842 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 8 | 0.4% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 315 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 5 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 20 | 0.2% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 1.0K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 120.8K | 13 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 3 | 0.0% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.1K | **Unique Entries:** 997 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 36 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 23 | 0.0% |

---

### ShadowWhisperer_Allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 647 | **Unique Entries:** 405 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 100 | 4 | 4.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 74 | 3.3% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 1 | 2.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Easy Privacy | allowlist | domain_adguard | 650 | 11 | 1.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 3 | 1.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 40 | 1.1% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 11 | 0.3% |
| Adaway | blocklist | hostname | 6.5K | 10 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 14 | 0.1% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 8 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 3 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 32 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 7 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |

---

### ShadowWhisperer_BlockLists Ads

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 22.1K | **Unique Entries:** 7.6K | **Target Sources:** 28

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 849 | 24.8% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 624 | 85 | 13.6% |
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.7K | 8.9% |
| AdBlockID | allowlist | domain_adguard | 53 | 4 | 7.5% |
| Adaway | blocklist | hostname | 6.5K | 407 | 6.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 969 | 6.2% |
| quidsup_notrack-malware | blocklist | domain | 152 | 7 | 4.6% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 152 | 4.2% |
| Easy Privacy | allowlist | domain_adguard | 650 | 27 | 4.2% |
| tranco | allowlist | domain_top | 100 | 3 | 3.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 5 | 2.8% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 325 | 2.8% |
| HaGeZi Pro | blocklist | domain | 419.7K | 7.7K | 1.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 3 | 0.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1.9K | 0.6% |
| Frogeye trackers | blocklist | hostname | 33.4K | 26 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 279 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 12 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |

---

### ShadowWhisperer_BlockLists Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 45.4K | **Unique Entries:** 26.7K | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 152 | 74 | 48.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.2K | 6.2% |
| YousList | blocklist | hostname | 624 | 19 | 3.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 12.7K | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 332 | 2.1% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 1 | 2.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 66 | 1.9% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 92 | 0.8% |
| Spam404 | blocklist | domain | 8.1K | 45 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 74 | 0.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 16 | 0.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 3.0K | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 572 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 46 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 650 | 1 | 0.2% |
| Adaway | blocklist | hostname | 6.5K | 15 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 522 | 0.2% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 2 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 8 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 12 | 0.0% |

---

### ShadowWhisperer_BlockLists Scam

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 10.8K | **Unique Entries:** 9.5K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 41 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 34 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 991 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 7 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 25 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 199 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 4 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1 | 0.0% |

---

### Spam404

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.1K | **Unique Entries:** 7.5K | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 152 | 2 | 1.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 20 | 0.2% |
| HaGeZi Pro | blocklist | domain | 419.7K | 338 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 45 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 63 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 108 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1 | 0.0% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 13 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |

---

### spamhaus_drop

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.5K | **Unique Entries:** 0 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.5K | 98.8% |
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.4K | 31.9% |

---

### Stamparm_Blackbook

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.1K | **Unique Entries:** 0 | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 17.6K | 8.3% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 47 | 4.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 5.0K | 0.7% |
| HaGeZi Pro | blocklist | domain | 419.7K | 1.1K | 0.3% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 4 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 16 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |

---

### StevenBlack_Fake_Gambling_Porn

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 304.9K | **Unique Entries:** 155.7K | **Target Sources:** 40

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Adaway | blocklist | hostname | 6.5K | 6.5K | 99.8% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 3.4K | 99.4% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 11.7K | 99.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 7.8K | 41.4% |
| YousList | blocklist | hostname | 624 | 243 | 38.9% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 78.4K | 37.0% |
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| WaLLy3K | blocklist | domain | 350 | 86 | 24.6% |
| abpvn_hosts | allowlist | domain_adguard | 50 | 12 | 24.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 523 | 18.1% |
| quidsup_notrack-malware | blocklist | domain | 152 | 26 | 17.1% |
| Easy Privacy | allowlist | domain_adguard | 650 | 101 | 15.5% |
| hufilter | blocklist | hostname | 99 | 14 | 14.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 22 | 12.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 418 | 11.4% |
| AdBlockID | allowlist | domain_adguard | 53 | 5 | 9.4% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 6 | 9.2% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 991 | 9.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 1.9K | 8.6% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.3% |
| HaGeZi Pro | blocklist | domain | 419.7K | 32.5K | 7.7% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 36 | 7.5% |
| tranco | allowlist | domain_top | 100 | 5 | 5.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 32 | 4.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 95 | 4.2% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 23 | 2.2% |
| Frogeye trackers | blocklist | hostname | 33.4K | 566 | 1.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 522 | 1.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 2 | 0.8% |
| Spam404 | blocklist | domain | 8.1K | 63 | 0.8% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 108 | 0.4% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 25 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 29 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 39 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 1.6K | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |

---

### tranco

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 100 | **Unique Entries:** 0 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 50 | 2.2% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Easy Privacy | allowlist | domain_adguard | 650 | 10 | 1.5% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 38 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 4 | 0.6% |
| Adaway | blocklist | hostname | 6.5K | 4 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 5 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 3 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 5 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |

---

### Ukrainian Ad Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.4K | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 91.3K | 54 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 29 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 2 | 0.0% |

---

### Ukrainian Privacy Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 345 | **Unique Entries:** 134 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 156 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 1 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 50 | 0.0% |

---

### Ukrainian Security Filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.7K | **Unique Entries:** 1.3K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 419.7K | 393 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 46 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 1 | 0.0% |

---

### URLHaus_Text

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 2.9K | **Unique Entries:** 940 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 2 | 66.7% |
| tranco | allowlist | domain_top | 100 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 20 | 0.9% |
| Easy Privacy | allowlist | domain_adguard | 650 | 3 | 0.5% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 65 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 1 | 0.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 12 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 2 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 523 | 0.2% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 405 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 629 | 0.1% |
| HaGeZi Pro | blocklist | domain | 419.7K | 278 | 0.1% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 2 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |

---

### URLHaus_Text

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 23.0K | **Unique Entries:** 20.5K | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 147 | 1.7% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 1.1K | 1.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 24 | 0.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 13 | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 65 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 860 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 85 | 0.4% |
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 25 | 0.3% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 15.6K | 53 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 379 | 1 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 41 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 5 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 35 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 9 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 1 | 0.0% |

---

### USOM-Blocklists-ips

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 12.2K | **Unique Entries:** 7.1K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_find | 7.5K | 280 | 3.7% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 3 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 4.3K | 2.2% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 14 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 2.8K | 12 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 65 | 0.3% |
| MyIP_MS_Blocklist | blocklist | ipv4_find | 498 | 1 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.5K | 20 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 4 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.2K | 8 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 15 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 21.3K | 22 | 0.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 125 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 12 | 0.1% |
| DShield | blocklist | ipv4_range_expand | 5.1K | 7 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.8K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 74 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 9.0K | 1 | 0.0% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.5K | **Unique Entries:** 3.1K | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 280 | 2.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 4.0K | 2.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 5 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.1K | 2 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 23.0K | 25 | 0.1% |
| Borestad_AbuseIPDB | blocklist | ipv4_find | 99.9K | 17 | 0.0% |
| Firehol_abusers_30d | blocklist | ipv4 | 194.6K | 21 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.9K | 2 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.2K | 1 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.2K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 11.7K | 5 | 0.0% |
| Sblam_Blocklist | blocklist | ipv4 | 2.1K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

---

### WaLLy3K

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 350 | **Unique Entries:** 0 | **Target Sources:** 26

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 2 | 4.0% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| tranco | allowlist | domain_top | 100 | 2 | 2.0% |
| YousList | blocklist | hostname | 624 | 9 | 1.4% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 650 | 5 | 0.8% |
| Adaway | blocklist | hostname | 6.5K | 54 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 152 | 1 | 0.7% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 19 | 0.6% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 1 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 85 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 3 | 0.5% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 11 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 20 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 53 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 86 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 12 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 163 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

---

### YousList

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 624 | **Unique Entries:** 0 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 5 | 10.0% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| AdBlockID | allowlist | domain_adguard | 53 | 1 | 1.9% |
| Adaway | blocklist | hostname | 6.5K | 111 | 1.7% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 47 | 1.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 204 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 108 | 0.9% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 22 | 0.6% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 85 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 650 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 22 | 0.1% |
| HaGeZi Pro | blocklist | domain | 419.7K | 451 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 243 | 0.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 3 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 19 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |

---

### YousList-AdGuard

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 12 | **Unique Entries:** 0 | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 1 | 2.0% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 2 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.4K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 2 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2 | 0.0% |
| Adaway | blocklist | hostname | 6.5K | 1 | 0.0% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 1 | 0.0% |

---

### YousList-AdGuard

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7.3K | **Unique Entries:** 7.3K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| AdGuard Base filter | blocklist | adguard | 91.3K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 120.8K | 40 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 10 | 0.0% |

---

### youtube_GoodbyeAds

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 97.6K | **Unique Entries:** 97.3K | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 50 | 5 | 10.0% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdBlockID | allowlist | domain_adguard | 53 | 3 | 5.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 2 | 3.1% |
| tranco | allowlist | domain_top | 100 | 3 | 3.0% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 2 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 20 | 0.5% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 50 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Adaway | blocklist | hostname | 6.5K | 28 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 52 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 650 | 2 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 2 | 0.3% |
| Yoyo Adservers-Hosts | blocklist | hostname | 3.4K | 8 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 6 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 74 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 47 | 0.0% |

---

### Yoyo Adservers-Hosts

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 3.4K | **Unique Entries:** 0 | **Target Sources:** 33

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 3 | 1 | 33.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2.1K | 11.3% |
| Easy Privacy | allowlist | domain_adguard | 650 | 44 | 6.8% |
| quidsup_notrack-malware | blocklist | domain | 152 | 10 | 6.6% |
| WaLLy3K | blocklist | domain | 350 | 19 | 5.4% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| tranco | allowlist | domain_top | 100 | 5 | 5.0% |
| Adaway | blocklist | hostname | 6.5K | 255 | 3.9% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 22.1K | 849 | 3.8% |
| AdBlockID | allowlist | domain_adguard | 53 | 2 | 3.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 590 | 3.7% |
| Dan Pollock's List | blocklist | hostname | 11.8K | 441 | 3.7% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 177 | 6 | 3.4% |
| AdGuard Base filter | allowlist | domain_adguard | 3.7K | 92 | 2.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 647 | 11 | 1.7% |
| AntiAdBlockFilters | allowlist | domain_adguard | 65 | 1 | 1.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 304.9K | 3.4K | 1.1% |
| HaGeZi Xiaomi Tracker | blocklist | domain | 479 | 5 | 1.0% |
| HaGeZi Pro | blocklist | domain | 419.7K | 3.1K | 0.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 260 | 1 | 0.4% |
| Frogeye trackers | blocklist | hostname | 33.4K | 32 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 45.4K | 66 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.8K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 679.8K | 15 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| URLHaus_Text | blocklist | domain_http_url | 2.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 211.7K | 4 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |

---

### Yoyo AdServers-IPList

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 9.0K | **Unique Entries:** 8.9K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 198.6K | 49 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.2K | 1 | 0.0% |

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources. High overlap percentages may indicate redundant sources, while low overlap percentages suggest unique content.

**Note:** Overlap percentages are calculated as: (overlap_count / source_total_count) × 100

