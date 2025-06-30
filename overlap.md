# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2025-08-09 14:27:33 UTC

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 99 |
| Total Entries Analyzed | 5.0M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 88 |
| allowlist | 11 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 20 |
| cidr_ipv4 | 2 |
| domain | 46 |
| ipv4 | 31 |

## Detailed Source Analysis

### abpvn_hosts

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 983 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 33 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 9 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 23 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 2 | 0.0% |

---

### abpvn_hosts

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 49 | **Unique Entries:** 0 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdBlockID | allowlist | domain_adguard | 59 | 2 | 3.4% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 13 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1 | 0.0% |

---

### AdBlockID

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 3.8K | **Unique Entries:** 3.8K | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 126.3K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 1 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 24 | 0.0% |

---

### AdBlockID

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 59 | **Unique Entries:** 23 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 2 | 4.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 6 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 5 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 4 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |

---

### AdGuard Base filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 98.6K | **Unique Entries:** 28.0K | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 28.8K | 64.9% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 41.4K | 32.8% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 54 | 3.8% |
| abpvn_hosts | blocklist | adguard | 1.1K | 9 | 0.9% |
| AdBlockID | blocklist | adguard | 3.8K | 24 | 0.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 3 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.0K | 83 | 0.4% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 5 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 16 | 0.2% |
| Easy Privacy | blocklist | adguard | 53.3K | 91 | 0.2% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| CJX Annoyance | blocklist | adguard | 1.8K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 74 | 0.0% |
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 6 | 0.0% |

---

### AdGuard DNS filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 180 | **Unique Entries:** 59 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| Easy Privacy | allowlist | domain_adguard | 653 | 9 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 3 | 0.5% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 26 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 22 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 5 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 12 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 28 | 0.0% |

---

### AdGuard DNS filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 126.3K | **Unique Entries:** 21.5K | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 38.2K | 85.9% |
| Easy Privacy | blocklist | adguard | 53.3K | 24.7K | 46.4% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 41.4K | 42.0% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 2 | 28.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 49 | 14.2% |
| abpvn_hosts | blocklist | adguard | 1.1K | 23 | 2.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 24 | 1.4% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 13 | 1.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.0K | 163 | 0.8% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 40 | 0.5% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 173 | 0.0% |

---

### AntiAdBlockFilters

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.7K | **Unique Entries:** 1.7K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 98.6K | 5 | 0.0% |

---

### bigdargon_hostsVN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.9K | **Unique Entries:** 0 | **Target Sources:** 31

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList | blocklist | hostname | 624 | 204 | 32.7% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 13 | 26.5% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 653 | 108 | 16.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 26 | 14.4% |
| quidsup_notrack-malware | blocklist | domain | 151 | 17 | 11.3% |
| tranco | allowlist | domain_top | 1.0K | 106 | 10.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 1.7K | 7.3% |
| AdBlockID | allowlist | domain_adguard | 59 | 4 | 6.8% |
| hufilter | blocklist | hostname | 99 | 6 | 6.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 12.7K | 3.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 7.8K | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1.2K | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 14 | 2.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 32 | 0.2% |
| Frogeye trackers | blocklist | hostname | 33.3K | 66 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 24 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 19 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 31 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 18 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 264 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 136 | 0.0% |

---

### BinaryDefense_Banlist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 3.5K | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 83 | 18.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 836 | 10.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 8.8% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 18 | 6.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 2.8K | 5.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1.2K | 5.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 783 | 4.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 52 | 2.4% |
| Greensnow | blocklist | ipv4 | 5.0K | 102 | 2.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 5 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 503 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 12 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 14 | 0.1% |

---

### BlockListDE_Strong

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 282 | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 116 | 5.3% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 96 | 3.7% |
| Greensnow | blocklist | ipv4 | 5.0K | 159 | 3.2% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 241 | 1.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 3 | 0.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 18 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 247 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 16 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 18 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 3 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |

---

### Blocklists UT1 Cryptojacking

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.3K | **Unique Entries:** 15.5K | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 151 | 4 | 2.6% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 17 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 232 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 46 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 49 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 76 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 49 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 34 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 262 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 8 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 9 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 4 | 0.0% |

---

### Blocklists UT1 Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 222.0K | **Unique Entries:** 64.8K | **Target Sources:** 32

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 43.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 90.1K | 28.8% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.7K | 21.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1.6K | 9.2% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 22 | 7.8% |
| quidsup_notrack-malware | blocklist | domain | 151 | 9 | 6.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 46 | 4.7% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 1.1K | 4.4% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 19.0K | 2.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 11.5K | 2.5% |
| HaGeZi Pro | blocklist | domain | 405.7K | 6.7K | 1.6% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 589 | 1.2% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 3.9K | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 49 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 27 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 31 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 12 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 16 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 18 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 27 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 40 | 0.0% |

---

### Borestad_AbuseIPDB_S100_3d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 52.7K | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 282 | 247 | 87.6% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1.8K | 81.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 2.8K | 79.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 6.3K | 76.8% |
| Greensnow | blocklist | ipv4 | 5.0K | 3.7K | 75.3% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 10.4K | 70.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10.6K | 70.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 297 | 67.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 7.6K | 32.7% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 5 | 25.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 367 | 16.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 370 | 14.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 233 | 11.9% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 33 | 6.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 7 | 3.1% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 15 | 2.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 2.9K | 1.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 215 | 1.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 96 | 0.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 16 | 0.6% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 6 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 13 | 0.0% |

---

### CINSScore_BadGuys_Army

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 15.0K | **Unique Entries:** 0 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 14.5K | 10.0K | 68.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 1.3K | 37.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 1.7K | 20.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 10.6K | 20.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 1.2K | 7.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 60 | 2.7% |
| Greensnow | blocklist | ipv4 | 5.0K | 102 | 2.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 6 | 1.4% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 3 | 1.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 22 | 1.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 1.2K | 0.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 31 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 3 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 16 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 2 | 0.0% |

---

### CJX Annoyance

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 1.8K | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 3 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 9 | 0.0% |
| Easy Privacy | blocklist | adguard | 53.3K | 4 | 0.0% |

---

### CJX Annoyance

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 6 | **Unique Entries:** 5 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Pro | blocklist | domain | 405.7K | 1 | 0.0% |

---

### CybercrimeTracker_All

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.9K | **Unique Entries:** 1.8K | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 55 | 53.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 488 | 10.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 50 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 473 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 1 | 0.0% |

---

### CybercrimeTracker_All

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 13.0K | **Unique Entries:** 551 | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 585 | 59.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1.8K | 36.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 2.7K | 1.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 3.9K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 746 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 17 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 50 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 95 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 6 | 0.0% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 103 | **Unique Entries:** 38 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 55 | 1.9% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 3 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 7 | 0.0% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 982 | **Unique Entries:** 177 | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 585 | 4.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 47 | 0.3% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 46 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 5 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 88 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 24 | 0.0% |

---

### cyberhost_malware-blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.9K | **Unique Entries:** 1.2K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 9.5K | 1.4% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 3.6K | 0.9% |
| quidsup_notrack-malware | blocklist | domain | 151 | 1 | 0.7% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 1.6K | 0.7% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 17 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 79 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 16 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 293 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 24 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 528 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 32 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |

---

### DandelionSprout-Anti-Malware-List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 32.6K | **Unique Entries:** 32.6K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Most Abused TLDs | blocklist | adguard | 426 | 2 | 0.5% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 6 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.3K | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| DoH_IP_list | blocklist | ipv4 | 731 | 94 | 12.9% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 8 | 12.9% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 100 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 12 | 0.0% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.3K | **Unique Entries:** 81 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1.2K | 81.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 3 | 0.0% |

---

### DoH_IP_list

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 731 | **Unique Entries:** 40 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 94 | 3.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 2 | 0.0% |

---

### Easy Privacy

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 53.3K | **Unique Entries:** 24.3K | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 91.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 156 | 45.2% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 24.7K | 19.6% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 2.4K | 5.4% |
| abpvn_hosts | blocklist | adguard | 1.1K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 91 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 7 | 0.0% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 653 | **Unique Entries:** 28 | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 60 | 6.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 9 | 5.0% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 2 | 5.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 46 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 10 | 1.5% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 108 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 78 | 0.5% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 27 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 101 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 3 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 8 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 159 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1 | 0.0% |

---

### EmergingThreats_CompromisedIPs

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 440 | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 14.5K | 410 | 2.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 83 | 2.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 34 | 1.5% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 3 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 150 | 1.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 297 | 0.6% |
| Greensnow | blocklist | ipv4 | 5.0K | 31 | 0.6% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 32 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 2 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 6 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 10 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |

---

### ET_fwip

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.6K | **Unique Entries:** 135 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level1 | blocklist | cidr_ipv4 | 4.5K | 1.5K | 32.6% |

---

### fabriziosalmi_allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 2.3K | **Unique Entries:** 1.5K | **Target Sources:** 31

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 40 | 16 | 40.0% |
| tranco | allowlist | domain_top | 1.0K | 268 | 26.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 74 | 11.4% |
| Easy Privacy | allowlist | domain_adguard | 653 | 46 | 7.0% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 2 | 4.1% |
| AdBlockID | allowlist | domain_adguard | 59 | 2 | 3.4% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 2 | 1.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 2 | 0.7% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 9 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 6 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 32 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 33 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 15 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 4 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 8 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 3 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 95 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 10 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 5 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 77 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 8 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 2 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 10 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |

---

### FabrizioSalmi_DNS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 66 | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 25 | 1.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 4 | 0.0% |

---

### FakeWebshopListHUN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.2K | **Unique Entries:** 4.8K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 41 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 22 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 27 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 29 | 0.0% |

---

### Firehol_BitcoinNodes_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.1K | **Unique Entries:** 7.0K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4_cidr_expand | 108 | 48 | 44.4% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 1 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 2 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 22 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 6 | 0.0% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 2 | 0.0% |

---

### Firehol_Botscout_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 620 | **Unique Entries:** 476 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 88 | 4.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 5 | 2.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 21 | 1.0% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 5 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 15 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 1 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 4 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |

---

### Firehol_CleanTalk

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 494 | **Unique Entries:** 430 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 2 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 9 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 33 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 2 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 6 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 2 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 2 | 0.0% |

---

### Firehol_CleanTalk_Top20

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 20 | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 6 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 5 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 3 | 0.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 2 | 0.0% |

---

### Firehol_GPF_Comics

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.2K | **Unique Entries:** 1.4K | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 21 | 3.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 63 | 3.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 3 | 1.3% |
| Greensnow | blocklist | ipv4 | 5.0K | 61 | 1.2% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 117 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 367 | 0.7% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 24 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 22 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 28 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 5 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 11 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 4 | 0.0% |

---

### Firehol_level1

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 4.5K | **Unique Entries:** 3.1K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.5K | 91.6% |

---

### Firehol_level2

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 14.8K | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 5.0K | 4.6K | 93.3% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 241 | 85.5% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1.7K | 78.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 8.0K | 34.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 150 | 34.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 783 | 22.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 1.8K | 22.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 10.4K | 19.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 406 | 15.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 8.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 117 | 5.3% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 73 | 3.7% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 9 | 1.8% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 254 | 1.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 2.7K | 1.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 4 | 0.6% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 9 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 57 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 21 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 2 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 8 | 0.0% |

---

### Firehol_level3

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 14.5K | **Unique Entries:** 0 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 410 | 93.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10.0K | 66.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 8.0K | 47.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 1.2K | 32.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 1.8K | 22.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 7.6K | 14.4% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 18 | 6.4% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 3 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 82 | 3.7% |
| Greensnow | blocklist | ipv4 | 5.0K | 158 | 3.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 514 | 2.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 3.8K | 1.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 28 | 1.3% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 103 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 53 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 5 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 5 | 0.1% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 1 | 0.0% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

---

### Firehol_SocksProxy_7d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.3K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 88 | 40 | 45.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 22 | 9.6% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 5 | 0.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 11 | 0.6% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 3 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 9 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 6 | 0.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 16 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 3 | 0.0% |

---

### Firehol_SSLProxies_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 228 | **Unique Entries:** 163 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 18 | 9 | 50.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 22 | 0.9% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 5 | 0.8% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 3 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 3 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 7 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 3 | 0.0% |

---

### Frogeye trackers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 33.3K | **Unique Entries:** 21.1K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 12 | 6.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 552 | 3.5% |
| HaGeZi Pro | blocklist | domain | 405.7K | 11.0K | 2.7% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| Easy Privacy | allowlist | domain_adguard | 653 | 8 | 1.2% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 4 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 66 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 562 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 26 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 1 | 0.0% |

---

### GetAdmiral Domains Filter List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 35 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.3K | 1.6K | 3.0% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 78 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 24 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 4 | 0.0% |

---

### GlobalAntiScamOrg-blocklist-domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 11.0K | **Unique Entries:** 7.4K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 27 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 13 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 1 | 0.0% |

---

### Greensnow

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 5.0K | **Unique Entries:** 0 | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 282 | 159 | 56.4% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1.0K | 45.6% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 4.6K | 31.5% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 316 | 12.3% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 3.7K | 7.1% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 31 | 7.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 328 | 4.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 102 | 2.9% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 61 | 2.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 28 | 1.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 6 | 1.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 102 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 158 | 0.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 114 | 0.1% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 9 | 0.1% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 4 | 0.0% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 2 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 2 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.4K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 1 | 0.0% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.4K | **Unique Entries:** 258 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1.2K | 92.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| HaGeZi Pro | blocklist | domain | 405.7K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 1 | 0.0% |

---

### HaGeZi Gambling Only Domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 183.3K | **Unique Entries:** 178.7K | **Target Sources:** 21

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Security Filter | blocklist | domain | 1.7K | 93 | 5.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 3.3K | 1.1% |
| quidsup_notrack-malware | blocklist | domain | 151 | 1 | 0.7% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 384 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 13 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 18 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 567 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 16 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 40 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 2 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 7 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 23 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 62 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 8 | 0.0% |

---

### HaGeZi Most Abused TLDs

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 426 | **Unique Entries:** 424 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DandelionSprout-Anti-Malware-List | blocklist | adguard | 32.6K | 2 | 0.0% |

---

### HaGeZi Pro

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 405.7K | **Unique Entries:** 305.5K | **Target Sources:** 42

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 99 | 91 | 91.9% |
| quidsup_notrack-malware | blocklist | domain | 151 | 123 | 81.5% |
| YousList | blocklist | hostname | 624 | 446 | 71.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 12.7K | 67.1% |
| WaLLy3K | blocklist | domain | 350 | 163 | 46.6% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 8.1K | 35.4% |
| Frogeye trackers | blocklist | hostname | 33.3K | 11.0K | 33.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 14.1K | 28.1% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 12 | 24.5% |
| Easy Privacy | allowlist | domain_adguard | 653 | 159 | 24.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 378 | 21.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.2K | 20.1% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 28 | 15.6% |
| tranco | allowlist | domain_top | 1.0K | 106 | 10.6% |
| AdBlockID | allowlist | domain_adguard | 59 | 6 | 10.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 26.6K | 8.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 6.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 293 | 6.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 1.2K | 4.7% |
| Spam404 | blocklist | domain | 8.1K | 349 | 4.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 77 | 3.4% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 36 | 3.4% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 528 | 3.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 6.7K | 3.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 217 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 407 | 1.9% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 232 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 7 | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 5.9K | 0.9% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 3.0K | 0.8% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 95 | 0.7% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 5 | 0.5% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 4 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.4K | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 724 | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 567 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 27 | 0.3% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 27 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 3 | 0.2% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 47 | 0.0% |

---

### hufilter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 99 | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 9 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 6 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 4 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 14 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 91 | 0.0% |

---

### jarelllama_Scam-Blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 457.7K | **Unique Entries:** 419.5K | **Target Sources:** 35

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 3.6K | 32.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 44 | 15.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 1.7K | 15.5% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 3.7K | 14.7% |
| quidsup_notrack-malware | blocklist | domain | 151 | 16 | 10.6% |
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 7 | 6.9% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 11.5K | 5.2% |
| hufilter | blocklist | hostname | 99 | 5 | 5.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1.1K | 2.3% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 293 | 1.7% |
| YousList | blocklist | hostname | 624 | 7 | 1.1% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 4.6K | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 4.9K | 0.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 136 | 0.7% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 7 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| HaGeZi Pro | blocklist | domain | 405.7K | 1.4K | 0.4% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 50 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 34 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 507 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 49 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 397 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 384 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 38 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |

---

### Local Allowlist (Domain)

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 40 | **Unique Entries:** 18 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 16 | 0.7% |
| Easy Privacy | allowlist | domain_adguard | 653 | 2 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 1 | 0.0% |

---

### Local Allowlist (ipv4)

**List Type:** allowlist | **Source Type:** ipv4 | **Total Entries:** 62 | **Unique Entries:** 45 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 8 | 0.3% |
| Firehol_level3 | blocklist | ipv4 | 14.5K | 3 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 5 | 0.0% |

---

### Local Blocklist (AdGuard)

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7 | **Unique Entries:** 1 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard DNS filter | blocklist | adguard | 126.3K | 2 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 4 | 0.0% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 203.2K | **Unique Entries:** 180.7K | **Target Sources:** 29

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 4.0K | 87.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 4.4K | 35.2% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 473 | 16.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 2.7K | 16.3% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 3.8K | 16.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 503 | 14.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 812 | 9.9% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 5 | 8.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.2K | 7.8% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 7 | 6.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 1.0K | 5.9% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 2.9K | 5.5% |
| Greensnow | blocklist | ipv4 | 5.0K | 114 | 2.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 10 | 2.3% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 3 | 1.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 22 | 1.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 12 | 0.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 12 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 9.0K | 49 | 0.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 11 | 0.5% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 1 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 22 | 0.3% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 6 | 0.2% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 40 | 0.1% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 687.3K | **Unique Entries:** 599.6K | **Target Sources:** 31

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.9% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 9.5K | 56.5% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 3.9K | 29.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 151 | 39 | 25.8% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 88 | 9.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 19.0K | 8.6% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 29.5K | 7.4% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 3.0K | 6.1% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 262 | 1.6% |
| HaGeZi Pro | blocklist | domain | 405.7K | 5.9K | 1.5% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 264 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 108 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 280 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4.9K | 1.1% |
| tranco | allowlist | domain_top | 1.0K | 9 | 0.9% |
| Easy Privacy | allowlist | domain_adguard | 653 | 3 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1.6K | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 3 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 48 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 35 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 20 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 8 | 0.1% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 62 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 86 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |

---

### malware-filter_phishing-filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 25.5K | **Unique Entries:** 19.2K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 21 | 7.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.7K | 0.8% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 1.1K | 0.5% |
| HaGeZi Pro | blocklist | domain | 405.7K | 1.2K | 0.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 48 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 90 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 2 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 12 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 2 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 45 | 0.0% |

---

### OISD Blocklist NSFW Small

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 21.0K | **Unique Entries:** 0 | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 88 | 8.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 15.1K | 5.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 8.3K | 2.7% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| quidsup_notrack-malware | blocklist | domain | 151 | 2 | 1.3% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 47 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 77 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 32 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 48 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 20 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 407 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 18 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 3 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 38 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 16 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 20 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 5 | 0.0% |

---

### OISD Blocklist NSFW Small

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 21.0K | **Unique Entries:** 20.6K | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 88 | 8.3% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 94 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 83 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 163 | 0.1% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 24 | 0.0% |

---

### OISD Blocklist Small

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 44.4K | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 4 | 57.1% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 38.2K | 30.2% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 28.8K | 29.2% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 39 | 11.3% |
| Easy Privacy | blocklist | adguard | 53.3K | 2.4K | 4.5% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 78 | 4.4% |
| abpvn_hosts | blocklist | adguard | 1.1K | 33 | 3.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.0K | 94 | 0.4% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 25 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 3 | 0.2% |
| AdBlockID | blocklist | adguard | 3.8K | 1 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 83 | 0.0% |

---

### OpenPhish_Feed

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 283 | **Unique Entries:** 188 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 21 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 44 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 22 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 2 | 0.0% |

---

### Public_DNS4

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 62.6K | **Unique Entries:** 61.8K | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 100 | 3.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 3 | 1.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 32 | 1.3% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 1 | 0.4% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 4 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 8 | 0.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 13 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 40 | 0.0% |
| Firehol_level3 | blocklist | ipv4 | 14.5K | 3 | 0.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 1 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 7 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 2 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 5 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

---

### PuppyScams

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 102 | **Unique Entries:** 84 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 10 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |

---

### quidsup_notrack-malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 151 | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 17 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 74 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 9 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 26 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 123 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 16 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 39 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 9 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |

---

### quidsup_notrack-tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 15.7K | **Unique Entries:** 7.7K | **Target Sources:** 26

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | allowlist | domain_adguard | 653 | 78 | 11.9% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| tranco | allowlist | domain_top | 1.0K | 117 | 11.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.3K | 6.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 9 | 5.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 969 | 4.2% |
| hufilter | blocklist | hostname | 99 | 4 | 4.0% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| AdBlockID | allowlist | domain_adguard | 59 | 2 | 3.4% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 12 | 1.9% |
| Frogeye trackers | blocklist | hostname | 33.3K | 552 | 1.7% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| HaGeZi Pro | blocklist | domain | 405.7K | 3.2K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 325 | 0.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1.3K | 0.4% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 20 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 7 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 35 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 41 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 16 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |

---

### RedDragonWebDesign_block-everything

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 652 | **Unique Entries:** 649 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdGuard Base filter | blocklist | adguard | 98.6K | 3 | 0.0% |

---

### RPiList_specials-phishing

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 777.7K | **Unique Entries:** 777.4K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 4 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 83 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 173 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 74 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | adguard | 1.4K | 1 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.0K | 24 | 0.1% |
| Easy Privacy | blocklist | adguard | 53.3K | 7 | 0.0% |

---

### Rutgers_DROP

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.2K | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 282 | 116 | 41.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 1.0K | 20.1% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 1.7K | 11.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 34 | 7.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 117 | 4.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 1.8K | 3.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 241 | 2.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 52 | 1.5% |
| Firehol_level3 | blocklist | ipv4 | 14.5K | 82 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 60 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 22 | 0.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 1 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 1 | 0.0% |

---

### Sblam_Blocklist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 1.9K | **Unique Entries:** 1.4K | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 6 | 30.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 88 | 14.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 63 | 2.9% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 228 | 3 | 1.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 4 | 0.8% |
| Greensnow | blocklist | ipv4 | 5.0K | 28 | 0.6% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 73 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 233 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 11 | 0.4% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 12 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 12 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 4 | 0.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 3 | 0.0% |

---

### ScriptzTeam_BadIPS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.6K | **Unique Entries:** 1.2K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 282 | 96 | 34.0% |
| Greensnow | blocklist | ipv4 | 5.0K | 316 | 6.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 117 | 5.3% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 406 | 2.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 370 | 0.7% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 11 | 0.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 2 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 5 | 0.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 16 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 3 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |

---

### Sentinel_Greylist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 8.2K | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 836 | 23.5% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 1.8K | 12.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 6.3K | 12.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.7K | 11.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 241 | 11.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 1.8K | 7.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 32 | 7.3% |
| Greensnow | blocklist | ipv4 | 5.0K | 328 | 6.6% |
| BlockListDE_Strong | blocklist | ipv4 | 282 | 16 | 5.7% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 1 | 1.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 229 | 1.3% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 24 | 1.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 812 | 0.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 11 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 3 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 15 | 0.1% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.1K | **Unique Entries:** 908 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 88 | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 3 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 36 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 23 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3 | 0.0% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 952 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 21.0K | 88 | 0.4% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 3 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 13 | 0.0% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 3 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 3 | 0.0% |

---

### ShadowWhisperer_Allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 648 | **Unique Entries:** 466 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 74 | 3.3% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 3 | 1.7% |
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| Easy Privacy | allowlist | domain_adguard | 653 | 10 | 1.5% |
| tranco | allowlist | domain_top | 1.0K | 11 | 1.1% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 1 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 14 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 3 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 32 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 5 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 7 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |

---

### ShadowWhisperer_BlockLists Ads

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 23.0K | **Unique Entries:** 9.5K | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 624 | 85 | 13.6% |
| hufilter | blocklist | hostname | 99 | 9 | 9.1% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.7K | 8.9% |
| AdBlockID | allowlist | domain_adguard | 59 | 4 | 6.8% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 969 | 6.2% |
| tranco | allowlist | domain_top | 1.0K | 61 | 6.1% |
| quidsup_notrack-malware | blocklist | domain | 151 | 7 | 4.6% |
| Easy Privacy | allowlist | domain_adguard | 653 | 27 | 4.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 5 | 2.8% |
| HaGeZi Pro | blocklist | domain | 405.7K | 8.1K | 2.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1.9K | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 77 | 0.4% |
| Frogeye trackers | blocklist | hostname | 33.3K | 26 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 49 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 280 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 9 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 13 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 12 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 6 | 0.0% |

---

### ShadowWhisperer_BlockLists Adult

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 264.3K | **Unique Entries:** 220.6K | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 15.1K | 71.9% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 27.2K | 8.7% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 724 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 397 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 9 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 86 | 0.0% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 7 | 0.0% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 17 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 2 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 27 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 23 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 6 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 3 | 0.0% |

---

### ShadowWhisperer_BlockLists Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 50.1K | **Unique Entries:** 28.8K | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 151 | 74 | 49.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 1.2K | 6.2% |
| HaGeZi Pro | blocklist | domain | 405.7K | 14.1K | 3.5% |
| YousList | blocklist | hostname | 624 | 19 | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 325 | 2.1% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 45 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 79 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 3.0K | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 589 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 46 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.1K | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 48 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 499 | 0.2% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 7 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 8 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 147 | 0.0% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 12 | 0.0% |

---

### ShadowWhisperer_BlockLists Scam

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 10.9K | **Unique Entries:** 7.8K | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 10 | 9.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 41 | 0.5% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.7K | 0.4% |
| Spam404 | blocklist | domain | 8.1K | 34 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 991 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 47 | 0.2% |
| HaGeZi Pro | blocklist | domain | 405.7K | 217 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 2 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 8 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 4 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 6 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 27 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2 | 0.0% |

---

### Spam404

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.1K | **Unique Entries:** 7.5K | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 151 | 2 | 1.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 34 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| HaGeZi Pro | blocklist | domain | 405.7K | 349 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 45 | 0.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 11 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 63 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 2 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 2 | 0.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 2 | 0.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 17 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 108 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 7 | 0.0% |

---

### Stamparm_Blackbook

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.1K | **Unique Entries:** 0 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.5% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.6K | 19.7% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 17.6K | 7.9% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 47 | 4.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 5.0K | 0.7% |
| HaGeZi Pro | blocklist | domain | 405.7K | 1.1K | 0.3% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 955 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 16 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 8 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 115 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 29 | 0.0% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 2 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 3 | 0.0% |

---

### StevenBlack_Fake_Gambling_Porn

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 312.6K | **Unique Entries:** 140.3K | **Target Sources:** 40

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 7.8K | 41.2% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 90.1K | 40.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 8.3K | 39.4% |
| YousList | blocklist | hostname | 624 | 243 | 38.9% |
| WaLLy3K | blocklist | domain | 350 | 86 | 24.6% |
| abpvn_hosts | allowlist | domain_adguard | 49 | 12 | 24.5% |
| quidsup_notrack-malware | blocklist | domain | 151 | 26 | 17.2% |
| Easy Privacy | allowlist | domain_adguard | 653 | 101 | 15.5% |
| hufilter | blocklist | hostname | 99 | 14 | 14.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 22 | 12.2% |
| tranco | allowlist | domain_top | 1.0K | 107 | 10.7% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 27.2K | 10.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 10.9K | 991 | 9.1% |
| AdBlockID | allowlist | domain_adguard | 59 | 5 | 8.5% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 1.9K | 8.2% |
| HaGeZi Pro | blocklist | domain | 405.7K | 26.6K | 6.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 32 | 4.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 95 | 4.2% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 23 | 2.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 3.3K | 1.8% |
| Frogeye trackers | blocklist | hostname | 33.3K | 562 | 1.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 499 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 63 | 0.8% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 29 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 1 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 90 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 49 | 0.3% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 29 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 1.6K | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 32 | 0.2% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 376 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 507 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 6 | 0.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2 | 0.0% |

---

### tranco

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 1.0K | **Unique Entries:** 101 | **Target Sources:** 27

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 268 | 11.9% |
| Easy Privacy | allowlist | domain_adguard | 653 | 60 | 9.2% |
| hufilter | blocklist | hostname | 99 | 3 | 3.0% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 11 | 1.7% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 2 | 0.7% |
| quidsup_notrack-malware | blocklist | domain | 151 | 1 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 117 | 0.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 106 | 0.6% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 61 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 3 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 2 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 107 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 10 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 4 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 10 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 106 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 9 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 1 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |

---

### Ukrainian Ad Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.3K | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.4K | 29 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 54 | 0.1% |
| Easy Privacy | blocklist | adguard | 53.3K | 2 | 0.0% |
| RPiList_specials-phishing | blocklist | adguard | 777.7K | 1 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 29 | 0.0% |

---

### Ukrainian Privacy Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 345 | **Unique Entries:** 96 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.3K | 156 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 39 | 0.1% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 1 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 2 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 49 | 0.0% |

---

### Ukrainian Security Filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.7K | **Unique Entries:** 1.2K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 93 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 378 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 46 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 1 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 13 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 3 | 0.0% |

---

### URLHaus_Text

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 17.6K | **Unique Entries:** 15.8K | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 229 | 2.8% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 28 | 0.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 2 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 1.0K | 0.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 57 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 103 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 215 | 0.4% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 57 | 0.3% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 12 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 31 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.0K | 2 | 0.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 7 | 0.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 1 | 0.0% |

---

### USOM-Blocklists-domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 401.1K | **Unique Entries:** 353.2K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 3.6K | 21.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 809 | 16.5% |
| quidsup_notrack-malware | blocklist | domain | 151 | 9 | 6.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 746 | 5.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 955 | 5.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 29.5K | 4.3% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 24 | 2.4% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 3.9K | 1.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4.6K | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 5 | 0.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 283 | 2 | 0.7% |
| HaGeZi Pro | blocklist | domain | 405.7K | 3.0K | 0.7% |
| YousList | blocklist | hostname | 624 | 4 | 0.6% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 76 | 0.5% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 147 | 0.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.5K | 45 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 19 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 376 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 6 | 0.1% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 9 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 17 | 0.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 7 | 0.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.0K | 5 | 0.0% |
| HaGeZi Gambling Only Domains | blocklist | domain | 183.3K | 2 | 0.0% |

---

### USOM-Blocklists-ips

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 12.4K | **Unique Entries:** 7.4K | **Target Sources:** 24

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 267 | 5.9% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 3 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 4.4K | 2.1% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 50 | 1.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 12 | 0.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.5K | 14 | 0.4% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 57 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 440 | 1 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.0K | 9 | 0.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 8.2K | 15 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 96 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 53 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.2K | 4 | 0.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.6K | 21 | 0.1% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 16 | 0.1% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 9.0K | 1 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 5 | 0.0% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 1 | 0.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.2K | 1 | 0.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 1 | 0.0% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 4.6K | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 488 | 17.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 267 | 2.2% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 4.0K | 2.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 620 | 1 | 0.2% |
| URLHaus_Text | blocklist | ipv4_http_url | 17.6K | 28 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 52.7K | 6 | 0.0% |
| Firehol_level2 | blocklist | ipv4 | 14.8K | 1 | 0.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 23.1K | 5 | 0.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 1 | 0.0% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 4.9K | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1.8K | 13.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 2.1K | 1.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 4.0K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 809 | 0.2% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 293 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 16.9K | 17 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 1 | 0.0% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 264.3K | 1 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 166 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 3 | 0.0% |

---

### WaLLy3K

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 350 | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 2 | 4.1% |
| AdBlockID | allowlist | domain_adguard | 59 | 2 | 3.4% |
| YousList | blocklist | hostname | 624 | 9 | 1.4% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 653 | 5 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 151 | 1 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 3 | 0.5% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 85 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 53 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 86 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 2 | 0.0% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 4 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 7 | 0.0% |
| Spam404 | blocklist | domain | 8.1K | 1 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 1 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 12 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 2 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 163 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1 | 0.0% |

---

### YousList

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 624 | **Unique Entries:** 0 | **Target Sources:** 19

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 5 | 10.2% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| AdBlockID | allowlist | domain_adguard | 59 | 1 | 1.7% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 204 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 85 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 653 | 1 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 243 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 22 | 0.1% |
| HaGeZi Pro | blocklist | domain | 405.7K | 446 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 3 | 0.0% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 50.1K | 19 | 0.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 687.3K | 3 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 5 | 0.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 4 | 0.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.0K | 3 | 0.0% |
| USOM-Blocklists-domains | blocklist | domain | 401.1K | 4 | 0.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 7 | 0.0% |

---

### YousList-AdGuard

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7.3K | **Unique Entries:** 7.3K | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.4K | 25 | 0.1% |
| Easy Privacy | blocklist | adguard | 53.3K | 10 | 0.0% |
| AdGuard Base filter | blocklist | adguard | 98.6K | 16 | 0.0% |
| AdGuard DNS filter | blocklist | adguard | 126.3K | 40 | 0.0% |

---

### YousList-AdGuard

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 12 | **Unique Entries:** 0 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 1 | 2.0% |
| Frogeye trackers | blocklist | hostname | 33.3K | 1 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 5 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 2 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 1 | 0.0% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 2 | 0.0% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 1 | 0.0% |

---

### youtube_GoodbyeAds

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 97.6K | **Unique Entries:** 97.4K | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | allowlist | domain_adguard | 49 | 5 | 10.2% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdBlockID | allowlist | domain_adguard | 59 | 3 | 5.1% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 2 | 1.1% |
| hufilter | blocklist | hostname | 99 | 1 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 653 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 18.9K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 648 | 2 | 0.3% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 6 | 0.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.0K | 6 | 0.0% |
| HaGeZi Pro | blocklist | domain | 405.7K | 47 | 0.0% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 312.6K | 74 | 0.0% |

---

### Yoyo AdServers-IPList

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 9.0K | **Unique Entries:** 8.9K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.2K | 49 | 0.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.4K | 1 | 0.0% |

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources. High overlap percentages may indicate redundant sources, while low overlap percentages suggest unique content.

**Note:** Overlap percentages are calculated as: (overlap_count / source_total_count) × 100

