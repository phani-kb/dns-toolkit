# DNS Toolkit - Detailed Overlap Analysis

This document provides comprehensive overlap analysis between different DNS sources, showing how entries are shared across blocklists and allowlists.

**Last Updated:** 2025-08-16 17:10:36 UTC

## Overview

| Metric | Value |
|--------|-------|
| Total Sources Analyzed | 112 |
| Total Entries Analyzed | 5.0M |

**Sources by List Type:**

| List Type | Count |
|-----------|-------|
| blocklist | 94 |
| allowlist | 18 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 19 |
| cidr_ipv4 | 2 |
| domain | 59 |
| ipv4 | 32 |

## Detailed Source Analysis

### abpvn_hosts

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 990 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CJX Annoyance | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 33 | 0.1% |

---

### AdGuard Base filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 94.4K | **Unique Entries:** 27.9K | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.5K | 28.9K | 64.9% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 37.3K | 30.0% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 54 | 3.7% |
| abpvn_hosts | blocklist | adguard | 1.1K | 10 | 0.9% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 2 | 0.6% |
| AdBlockID | blocklist | adguard | 3.8K | 24 | 0.6% |
| RedDragonWebDesign_block-everything | blocklist | adguard | 652 | 3 | 0.5% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.5K | 83 | 0.4% |
| AntiAdBlockFilters | blocklist | adguard | 1.7K | 5 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| Easy Privacy | blocklist | adguard | 53.2K | 91 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 16 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 2 | 0.1% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |

---

### AdGuard DNS filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 124.0K | **Unique Entries:** 20.2K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 92.9% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 38.3K | 86.0% |
| Easy Privacy | blocklist | adguard | 53.2K | 26.1K | 49.1% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 37.3K | 39.4% |
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 2 | 28.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 49 | 14.2% |
| abpvn_hosts | blocklist | adguard | 1.1K | 24 | 2.3% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 13 | 1.2% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.5K | 162 | 0.8% |
| CJX Annoyance | blocklist | adguard | 1.8K | 9 | 0.5% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 40 | 0.5% |
| AdBlockID | blocklist | adguard | 3.8K | 2 | 0.1% |

---

### AdGuard DNS filter

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 180 | **Unique Entries:** 51 | **Target Sources:** 10

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | allowlist | domain_adguard | 654 | 9 | 1.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 3 | 0.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 9 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 26 | 0.1% |

---

### AdguardTeam_HttpsExclusions_android

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 97 | **Unique Entries:** 78 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |

---

### AdguardTeam_HttpsExclusions_banks

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 4.0K | **Unique Entries:** 3.9K | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 3 | 1.8% |
| Easy Privacy | allowlist | domain_adguard | 654 | 7 | 1.1% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| BlahDNS_whitelist | allowlist | domain | 773 | 3 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |

---

### AdguardTeam_HttpsExclusions_firefox

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 18 | **Unique Entries:** 13 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |

---

### AdguardTeam_HttpsExclusions_issues

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 68 | **Unique Entries:** 60 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |

---

### AdguardTeam_HttpsExclusions_mac

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 11 | **Unique Entries:** 5 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |

---

### AdguardTeam_HttpsExclusions_sensitive

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 164 | **Unique Entries:** 135 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 2 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |

---

### AdguardTeam_HttpsExclusions_windows

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 7 | **Unique Entries:** 6 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |

---

### bigdargon_hostsVN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 19.0K | **Unique Entries:** 0 | **Target Sources:** 24

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList | blocklist | hostname | 624 | 204 | 32.7% |
| WaLLy3K | blocklist | domain | 350 | 85 | 24.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| Easy Privacy | allowlist | domain_adguard | 654 | 108 | 16.5% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 26 | 14.4% |
| quidsup_notrack-malware | blocklist | domain | 150 | 17 | 11.3% |
| tranco | allowlist | domain_top | 1.0K | 106 | 10.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1.7K | 7.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 50 | 6.5% |
| hufilter | blocklist | hostname | 100 | 6 | 6.0% |
| HaGeZi Pro | blocklist | domain | 409.8K | 12.7K | 3.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 7.8K | 2.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 1.2K | 2.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 14 | 2.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 33 | 1.5% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 33 | 0.2% |
| Torrent Trackers | blocklist | domain | 479 | 1 | 0.2% |
| Frogeye trackers | blocklist | hostname | 33.3K | 67 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 24 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 52 | 0.1% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 9 | 0.1% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 1 | 0.1% |

---

### BinaryDefense_Banlist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 3.0K | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.1K | 7.0% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 626 | 6.8% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 15 | 5.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 856 | 4.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 19 | 4.2% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 619 | 3.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 2.2K | 2.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 30 | 1.5% |
| Greensnow | blocklist | ipv4 | 5.1K | 72 | 1.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 485 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 14 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 13 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 3 | 0.1% |

---

### BlahDNS_whitelist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 773 | **Unique Entries:** 304 | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 49 | 7.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 107 | 4.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 6 | 3.3% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 6 | 2.5% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| tranco | allowlist | domain_top | 1.0K | 24 | 2.4% |
| Dogino_Discord_Official | allowlist | domain | 43 | 1 | 2.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 13 | 2.0% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 50 | 0.3% |
| YousList | blocklist | hostname | 624 | 2 | 0.3% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 19 | 0.1% |

---

### BlockListDE_Brute

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 613 | **Unique Entries:** 0 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level2 | blocklist | ipv4 | 14.2K | 572 | 4.0% |
| Greensnow | blocklist | ipv4 | 5.1K | 127 | 2.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 47 | 2.1% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 15 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 379 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 28 | 0.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 2 | 0.1% |

---

### BlockListDE_Strong

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 276 | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 87 | 4.4% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 93 | 3.6% |
| Greensnow | blocklist | ipv4 | 5.1K | 148 | 2.9% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 212 | 1.5% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 4 | 0.9% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 15 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 254 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 15 | 0.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 17 | 0.1% |

---

### Blocklists UT1 Cryptojacking

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 16.3K | **Unique Entries:** 15.4K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 4 | 2.7% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 8 | 1.0% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 17 | 0.1% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 4 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 232 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 46 | 0.1% |

---

### Blocklists UT1 Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 222.5K | **Unique Entries:** 65.2K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 17.6K | 97.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.1K | 43.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 90.8K | 29.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.7K | 21.0% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 1.6K | 9.0% |
| quidsup_notrack-malware | blocklist | domain | 150 | 9 | 6.0% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 14 | 5.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 46 | 4.7% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 1.1K | 4.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 19.0K | 2.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 11.5K | 2.5% |
| HaGeZi Pro | blocklist | domain | 409.8K | 6.2K | 1.5% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 613 | 1.2% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 3.9K | 1.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 22 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 49 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 27 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 31 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 20 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 12 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 16 | 0.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 5 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| Spam404 | blocklist | domain | 8.1K | 11 | 0.1% |

---

### Blocklists UT1 Shortener

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 4.5K | **Unique Entries:** 59 | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 4.1K | 87.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 93 | 39.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 60 | 14.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 5 | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 7 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 1 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 26 | 0.1% |

---

### Borestad_AbuseIPDB_S100_3d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 81.7K | **Unique Entries:** 19.3K | **Target Sources:** 24

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 276 | 254 | 92.0% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 1.7K | 85.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 7.2K | 79.1% |
| Greensnow | blocklist | ipv4 | 5.1K | 4.0K | 78.6% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 348 | 77.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 11.5K | 77.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 2.2K | 74.2% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 10.0K | 70.0% |
| Firehol_level3 | blocklist | ipv4 | 12.9K | 8.7K | 67.2% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 379 | 61.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 484 | 21.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 4 | 20.0% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 431 | 16.8% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 319 | 16.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 36 | 7.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 23 | 4.2% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 7 | 1.9% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 275 | 1.6% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 3.1K | 1.5% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 36 | 1.5% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 133 | 1.1% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 1 | 1.0% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 9 | 0.2% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 6 | 0.1% |

---

### Boutetnico_URL_Shorteners

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 418 | **Unique Entries:** 222 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Korlabs_UrlShortener | blocklist | domain | 237 | 56 | 23.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 60 | 1.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 6 | 0.9% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 6 | 0.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 12 | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 15 | 0.3% |

---

### CINSScore_BadGuys_Army

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 15.0K | **Unique Entries:** 0 | **Target Sources:** 14

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.9K | 10.3K | 79.7% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 1.1K | 34.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.4K | 15.3% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 11.5K | 14.1% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 1.3K | 8.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 73 | 3.7% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 6 | 2.2% |
| Greensnow | blocklist | ipv4 | 5.1K | 95 | 1.9% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 8 | 1.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 24 | 1.0% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 1.4K | 0.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 6 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 15 | 0.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 24 | 0.1% |

---

### CJX Annoyance

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 1.8K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| abpvn_hosts | blocklist | adguard | 1.1K | 1 | 0.1% |

---

### CybercrimeTracker_All

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.9K | **Unique Entries:** 1.8K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 55 | 53.4% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 488 | 10.7% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 50 | 0.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 473 | 0.2% |

---

### CybercrimeTracker_All

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 13.0K | **Unique Entries:** 543 | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 585 | 59.6% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 1.8K | 36.6% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.6K | 14.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 2.7K | 1.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 3.9K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 746 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 17 | 0.1% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 103 | **Unique Entries:** 37 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 55 | 1.9% |

---

### CybercrimeTracker_CCPMGate

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 982 | **Unique Entries:** 177 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 585 | 4.5% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 47 | 0.3% |

---

### cyberhost_malware-blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 17.2K | **Unique Entries:** 1.2K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 9.7K | 1.4% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 3.7K | 0.9% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 1.6K | 0.7% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 17 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 2 | 0.2% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 83 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 24 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 3 | 0.1% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 16 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 565 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 302 | 0.1% |

---

### DandelionSprout-Anti-Malware-List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 32.6K | **Unique Entries:** 32.6K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Most Abused TLDs | blocklist | adguard | 425 | 2 | 0.5% |

---

### Dogino_Discord_Official

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 43 | **Unique Entries:** 5 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 26 | 1.2% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.3K | **Unique Entries:** 80 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1.2K | 81.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 9 | 0.4% |
| Torrent Trackers | blocklist | domain | 479 | 1 | 0.2% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |

---

### DoH_IP_blocklists

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.3K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 25 | 37.9% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 11 | 17.7% |
| DoH_IP_list | blocklist | ipv4 | 731 | 94 | 12.9% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 100 | 0.2% |

---

### DoH_IP_list

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 731 | **Unique Entries:** 40 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 26 | 39.4% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 94 | 3.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 569 | 0.9% |

---

### Easy Privacy

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 654 | **Unique Entries:** 4 | **Target Sources:** 16

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| tranco | allowlist | domain_top | 1.0K | 60 | 6.0% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 2 | 5.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 9 | 5.0% |
| Dogino_Discord_Official | allowlist | domain | 43 | 1 | 2.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 46 | 2.0% |
| BlahDNS_whitelist | allowlist | domain | 773 | 13 | 1.7% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 10 | 1.5% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 108 | 0.6% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 78 | 0.5% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 7 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 27 | 0.1% |

---

### Easy Privacy

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 53.2K | **Unique Entries:** 22.8K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1.6K | 92.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 156 | 45.2% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 26.1K | 21.1% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 2.4K | 5.4% |
| abpvn_hosts | blocklist | adguard | 1.1K | 2 | 0.2% |
| CJX Annoyance | blocklist | adguard | 1.8K | 4 | 0.2% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 10 | 0.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 2 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 91 | 0.1% |

---

### EmergingThreats_CompromisedIPs

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 450 | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level3 | blocklist | ipv4 | 12.9K | 406 | 3.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 37 | 1.9% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 4 | 1.4% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 117 | 0.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 19 | 0.6% |
| Greensnow | blocklist | ipv4 | 5.1K | 30 | 0.6% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 348 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 27 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 8 | 0.1% |

---

### ET_fwip

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 1.6K | **Unique Entries:** 138 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_level1 | blocklist | cidr_ipv4 | 4.6K | 1.5K | 32.9% |

---

### fabriziosalmi_allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 2.3K | **Unique Entries:** 1.3K | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Dogino_Discord_Official | allowlist | domain | 43 | 26 | 60.5% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 16 | 40.0% |
| AdguardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| tranco | allowlist | domain_top | 1.0K | 268 | 26.8% |
| BlahDNS_whitelist | allowlist | domain | 773 | 107 | 13.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 74 | 11.4% |
| AdguardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| Easy Privacy | allowlist | domain_adguard | 654 | 46 | 7.0% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 6 | 6.2% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 7 | 3.0% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 12 | 2.9% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 3 | 1.8% |
| WaLLy3K | blocklist | domain | 350 | 5 | 1.4% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 2 | 1.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 9 | 0.7% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 6 | 0.4% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 10 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 9 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 32 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 33 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 15 | 0.1% |

---

### FabrizioSalmi_DNS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 66 | **Unique Entries:** 0 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 26 | 3.6% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 25 | 1.0% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 32 | 0.1% |

---

### FakeWebshopListHUN

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.2K | **Unique Entries:** 4.8K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| hufilter | blocklist | hostname | 100 | 9 | 9.0% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.2K | 0.7% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 42 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |

---

### Firehol_BitcoinNodes_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 7.1K | **Unique Entries:** 7.0K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_BitcoinNodes_1d | blocklist | ipv4_cidr_expand | 94 | 42 | 44.7% |

---

### Firehol_Botscout_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 549 | **Unique Entries:** 423 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 72 | 3.7% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 10 | 2.7% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 9 | 0.4% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 4 | 0.2% |

---

### Firehol_CleanTalk

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 494 | **Unique Entries:** 432 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.1K | 6 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 2 | 0.1% |

---

### Firehol_CleanTalk_Top20

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 20 | **Unique Entries:** 7 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 3 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 2 | 0.1% |

---

### Firehol_GPF_Comics

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.3K | **Unique Entries:** 1.4K | **Target Sources:** 15

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 2 | 10.0% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 47 | 7.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 55 | 2.8% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 9 | 1.6% |
| Greensnow | blocklist | ipv4 | 5.1K | 65 | 1.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 4 | 1.1% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 112 | 0.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 484 | 0.6% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 27 | 0.3% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 24 | 0.2% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 3 | 0.1% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 30 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 3 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 1 | 0.1% |

---

### Firehol_level1

**List Type:** blocklist | **Source Type:** cidr_ipv4 | **Total Entries:** 4.6K | **Unique Entries:** 3.1K | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ET_fwip | blocklist | cidr_ipv4 | 1.6K | 1.5K | 91.6% |

---

### Firehol_level2

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 14.2K | **Unique Entries:** 0 | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Greensnow | blocklist | ipv4 | 5.1K | 4.8K | 94.7% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 572 | 93.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 1.6K | 78.1% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 212 | 76.8% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 8.0K | 39.7% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 117 | 26.0% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 619 | 20.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.7K | 18.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 355 | 13.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 10.0K | 12.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.3K | 9.0% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 112 | 4.9% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 37 | 1.9% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 204 | 1.3% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 2.4K | 1.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 6 | 1.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 5 | 0.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 19 | 0.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 67 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 1 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 17 | 0.1% |

---

### Firehol_level3

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 12.9K | **Unique Entries:** 0 | **Target Sources:** 23

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 406 | 90.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 10.3K | 68.6% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 8.0K | 49.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 856 | 28.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 1.5K | 16.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 8.7K | 10.6% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 17 | 6.2% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 3 | 4.8% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 84 | 4.2% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 720 | 3.6% |
| Greensnow | blocklist | ipv4 | 5.1K | 121 | 2.4% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 3.2K | 1.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 30 | 1.3% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 5 | 0.8% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 101 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 51 | 0.4% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 6 | 0.3% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 1 | 0.3% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 6 | 0.2% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 5 | 0.1% |

---

### Firehol_SocksProxy_7d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.5K | **Unique Entries:** 2.2K | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SocksProxy_7d | blocklist | ipv4_cidr_expand | 108 | 52 | 48.1% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 21 | 5.7% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 4 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 14 | 0.7% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 34 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 2 | 0.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 3 | 0.1% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 2 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 19 | 0.1% |

---

### Firehol_SSLProxies_1d

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 369 | **Unique Entries:** 296 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_SSLProxies_1d | blocklist | ipv4_cidr_expand | 22 | 11 | 50.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 10 | 1.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 21 | 0.8% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 4 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 2 | 0.1% |

---

### Frogeye trackers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 33.3K | **Unique Entries:** 21.1K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 12 | 6.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 548 | 3.5% |
| HaGeZi Pro | blocklist | domain | 409.8K | 11.0K | 2.7% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| Easy Privacy | allowlist | domain_adguard | 654 | 8 | 1.2% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| WaLLy3K | blocklist | domain | 350 | 2 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 67 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| BlahDNS_whitelist | allowlist | domain | 773 | 2 | 0.3% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 568 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 26 | 0.1% |

---

### GetAdmiral Domains Filter List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.8K | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 1.6K | 3.0% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 1.6K | 1.3% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 80 | 0.2% |

---

### GlobalAntiScamOrg-blocklist-domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 11.0K | **Unique Entries:** 7.4K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.6K | 0.8% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |

---

### Greensnow

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 5.1K | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 276 | 148 | 53.6% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 696 | 34.9% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 4.8K | 33.9% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 127 | 20.7% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 319 | 12.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 30 | 6.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 4.0K | 4.9% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 296 | 3.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 65 | 2.8% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 72 | 2.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 6 | 1.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 13 | 0.7% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 121 | 0.6% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 95 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 9 | 0.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 109 | 0.1% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |

---

### HaGeZi Encrypted DNS Servers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.4K | **Unique Entries:** 254 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1.2K | 92.5% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 6 | 0.3% |
| Torrent Trackers | blocklist | domain | 479 | 1 | 0.2% |

---

### HaGeZi Gambling Only Domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 185.3K | **Unique Entries:** 180.7K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Security Filter | blocklist | domain | 1.7K | 93 | 5.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 3.3K | 1.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 409.8K | 568 | 0.1% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 13 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 18 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 7 | 0.1% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 3 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 384 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 17 | 0.1% |

---

### HaGeZi Pro

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 409.8K | **Unique Entries:** 310.1K | **Target Sources:** 42

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| hufilter | blocklist | hostname | 100 | 91 | 91.0% |
| quidsup_notrack-malware | blocklist | domain | 150 | 123 | 82.0% |
| YousList | blocklist | hostname | 624 | 444 | 71.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 12.7K | 67.0% |
| WaLLy3K | blocklist | domain | 350 | 163 | 46.6% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 5 | 41.7% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 8.3K | 35.7% |
| Frogeye trackers | blocklist | hostname | 33.3K | 11.0K | 32.9% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 14.7K | 28.4% |
| Easy Privacy | allowlist | domain_adguard | 654 | 159 | 24.3% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 376 | 21.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 3.1K | 19.8% |
| CJX Annoyance | allowlist | domain_adguard | 6 | 1 | 16.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 28 | 15.6% |
| tranco | allowlist | domain_top | 1.0K | 107 | 10.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 68 | 8.8% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 25.7K | 8.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 1.1K | 6.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 296 | 6.0% |
| Spam404 | blocklist | domain | 8.1K | 355 | 4.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 1.1K | 4.2% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 36 | 3.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 77 | 3.4% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 565 | 3.3% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 9 | 3.2% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 6.2K | 2.8% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 225 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 408 | 1.9% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 232 | 1.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 7 | 1.1% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 6.3K | 0.9% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 103 | 0.8% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 2.8K | 0.7% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 5 | 0.5% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 36 | 0.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 274.1K | 795 | 0.3% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.5K | 0.3% |
| HaGeZi Gambling Only Domains | blocklist | domain | 185.3K | 568 | 0.3% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 26 | 0.3% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 3 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 3 | 0.2% |

---

### hufilter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 100 | **Unique Entries:** 0 | **Target Sources:** 5

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 9 | 0.1% |

---

### jarelllama_Scam-Blocklist

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 457.7K | **Unique Entries:** 419.5K | **Target Sources:** 41

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| FakeWebshopListHUN | blocklist | domain | 8.2K | 3.2K | 39.4% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 3.6K | 32.8% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 54 | 19.4% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 1.7K | 15.3% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 3.6K | 14.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 16 | 10.7% |
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 7 | 6.9% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 11.5K | 5.1% |
| hufilter | blocklist | hostname | 100 | 5 | 5.0% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 166 | 3.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 107 | 2.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 1.1K | 2.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 302 | 1.8% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 73 | 1.6% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 4.6K | 1.1% |
| YousList | blocklist | hostname | 624 | 7 | 1.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 136 | 0.7% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 4.9K | 0.7% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 7 | 0.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 115 | 0.6% |
| HaGeZi Pro | blocklist | domain | 409.8K | 1.5K | 0.4% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 50 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 3 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 34 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 17 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 472 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 4 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 38 | 0.2% |
| HaGeZi Gambling Only Domains | blocklist | domain | 185.3K | 384 | 0.2% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 49 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 3 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 274.1K | 399 | 0.1% |

---

### Korlabs_UrlShortener

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 237 | **Unique Entries:** 0 | **Target Sources:** 9

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 56 | 13.4% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 93 | 2.1% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 53 | 1.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 6 | 0.8% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 5 | 0.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 1 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 7 | 0.3% |

---

### Local Allowlist (Domain)

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 40 | **Unique Entries:** 16 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 1 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 16 | 0.7% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 1 | 0.6% |
| Easy Privacy | allowlist | domain_adguard | 654 | 2 | 0.3% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |

---

### Local Allowlist (ipv4)

**List Type:** allowlist | **Source Type:** ipv4 | **Total Entries:** 62 | **Unique Entries:** 43 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 11 | 0.4% |

---

### Local Blocklist (Domain)

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1 | **Unique Entries:** 0 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| YousList | blocklist | hostname | 624 | 1 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 690.4K | **Unique Entries:** 601.9K | **Target Sources:** 36

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 4.0K | 80.9% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 9.7K | 56.4% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 3.9K | 29.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 5.0K | 27.5% |
| quidsup_notrack-malware | blocklist | domain | 150 | 39 | 26.0% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 88 | 9.0% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 19.0K | 8.5% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 29.5K | 7.3% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 11 | 6.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 3.1K | 5.9% |
| WaLLy3K | blocklist | domain | 350 | 12 | 3.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 13 | 3.1% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 2 | 2.9% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 46 | 2.6% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 4 | 1.7% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 262 | 1.6% |
| HaGeZi Pro | blocklist | domain | 409.8K | 6.3K | 1.5% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 264 | 1.4% |
| Spam404 | blocklist | domain | 8.1K | 108 | 1.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 280 | 1.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 9 | 1.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4.9K | 1.1% |
| tranco | allowlist | domain_top | 1.0K | 9 | 0.9% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 25 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 3 | 0.5% |
| Easy Privacy | allowlist | domain_adguard | 654 | 3 | 0.5% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 1.6K | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 10 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 21 | 0.4% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 1 | 0.4% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 50 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 35 | 0.2% |
| GlobalAntiScamOrg-blocklist-domains | blocklist | domain | 11.0K | 13 | 0.1% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 8 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 19 | 0.1% |

---

### Maltrail_StaticTrails

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 203.4K | **Unique Entries:** 181.5K | **Target Sources:** 29

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 4.0K | 87.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 4.4K | 35.2% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 473 | 16.5% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 485 | 16.0% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 3.2K | 16.0% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 2.4K | 14.9% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.4K | 9.1% |
| Local Allowlist (ipv4) | allowlist | ipv4 | 62 | 5 | 8.1% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 707 | 7.7% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 7 | 6.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 4 | 6.1% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 973 | 5.7% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 3.1K | 3.8% |
| Greensnow | blocklist | ipv4 | 5.1K | 109 | 2.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 28 | 1.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 6 | 1.3% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 3 | 1.1% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 6 | 1.0% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 3 | 0.8% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 17 | 0.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 11 | 0.6% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 16 | 0.6% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 11 | 0.5% |
| Yoyo AdServers-IPList | blocklist | ipv4 | 8.9K | 49 | 0.5% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| DoH_IP_list | blocklist | ipv4 | 731 | 2 | 0.3% |
| Firehol_BitcoinNodes_1d | blocklist | ipv4 | 7.1K | 20 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 7 | 0.3% |
| Public_DNS4 | blocklist | ipv4 | 62.6K | 40 | 0.1% |

---

### malware-filter_phishing-filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 25.8K | **Unique Entries:** 19.5K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 227 | 81.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 10 | 4.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 3.6K | 0.8% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 26 | 0.6% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 1.1K | 0.5% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 19 | 0.4% |
| HaGeZi Pro | blocklist | domain | 409.8K | 1.1K | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |

---

### OISD Blocklist NSFW Small

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 21.5K | **Unique Entries:** 0 | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 89 | 8.3% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 274.1K | 15.6K | 5.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 8.4K | 2.7% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 2 | 1.3% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 48 | 0.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 78 | 0.3% |
| Torrent Trackers | blocklist | domain | 479 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 33 | 0.2% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 21 | 0.1% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 49 | 0.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 408 | 0.1% |

---

### OISD Blocklist NSFW Small

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 21.5K | **Unique Entries:** 21.0K | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 89 | 8.3% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 94 | 0.2% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 83 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 162 | 0.1% |

---

### OISD Blocklist Small

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 44.5K | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (AdGuard) | blocklist | adguard | 7 | 4 | 57.1% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 38.3K | 30.9% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 28.9K | 30.6% |
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 39 | 11.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 80 | 4.6% |
| Easy Privacy | blocklist | adguard | 53.2K | 2.4K | 4.5% |
| abpvn_hosts | blocklist | adguard | 1.1K | 33 | 3.1% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 29 | 2.0% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.5K | 94 | 0.4% |
| YousList-AdGuard | blocklist | adguard | 7.3K | 25 | 0.3% |
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| CJX Annoyance | blocklist | adguard | 1.8K | 3 | 0.2% |

---

### OpenPhish_Feed

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 279 | **Unique Entries:** 0 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 227 | 0.9% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 1 | 0.4% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 1 | 0.2% |

---

### Public_DNS4

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 62.6K | **Unique Entries:** 61.8K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_list | blocklist | ipv4 | 731 | 569 | 77.8% |
| FabrizioSalmi_DNS | blocklist | ipv4 | 66 | 32 | 48.5% |
| DoH_IP_blocklists | blocklist | ipv4 | 2.5K | 100 | 3.9% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 34 | 1.4% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 1 | 0.4% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 1 | 0.3% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 1 | 0.2% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 1 | 0.1% |
| Greensnow | blocklist | ipv4 | 5.1K | 3 | 0.1% |

---

### PuppyScams

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 102 | **Unique Entries:** 84 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 10 | 0.1% |

---

### quidsup_notrack-malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 150 | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 74 | 0.1% |
| tranco | allowlist | domain_top | 1.0K | 1 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 17 | 0.1% |

---

### quidsup_notrack-tracker

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 15.7K | **Unique Entries:** 7.7K | **Target Sources:** 20

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | allowlist | domain_adguard | 654 | 78 | 11.9% |
| WaLLy3K | blocklist | domain | 350 | 41 | 11.7% |
| tranco | allowlist | domain_top | 1.0K | 117 | 11.7% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.3K | 6.8% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 9 | 5.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 968 | 4.2% |
| hufilter | blocklist | hostname | 100 | 4 | 4.0% |
| YousList | blocklist | hostname | 624 | 22 | 3.5% |
| BlahDNS_whitelist | allowlist | domain | 773 | 19 | 2.5% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 2 | 2.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 12 | 1.8% |
| Frogeye trackers | blocklist | hostname | 33.3K | 548 | 1.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 32 | 1.4% |
| HaGeZi Pro | blocklist | domain | 409.8K | 3.1K | 0.8% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 318 | 0.6% |
| Torrent Trackers | blocklist | domain | 479 | 2 | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 1.3K | 0.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 21 | 0.1% |

---

### RPiList_specials-phishing

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 785.5K | **Unique Entries:** 785.1K | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| ShadowWhisperer's Dating List | blocklist | adguard | 1.1K | 3 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 4 | 0.2% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 81 | 0.2% |
| Ukrainian Ad Filter | blocklist | adguard | 1.4K | 1 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 71 | 0.1% |
| AdGuard DNS filter | blocklist | adguard | 124.0K | 174 | 0.1% |
| OISD Blocklist NSFW Small | blocklist | adguard | 21.5K | 26 | 0.1% |

---

### Rutgers_DROP

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.0K | **Unique Entries:** 0 | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 276 | 87 | 31.5% |
| Greensnow | blocklist | ipv4 | 5.1K | 696 | 13.7% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 1.6K | 11.0% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 37 | 8.2% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 344 | 3.8% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 72 | 2.8% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 1.7K | 2.1% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 30 | 1.0% |
| Firehol_level3 | blocklist | ipv4 | 12.9K | 84 | 0.7% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 73 | 0.5% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 2 | 0.3% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |

---

### Sblam_Blocklist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 1.9K | **Unique Entries:** 1.4K | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 3 | 15.0% |
| Firehol_Botscout_1d | blocklist | ipv4 | 549 | 72 | 13.1% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 55 | 2.4% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 15 | 2.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 4 | 0.8% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 14 | 0.6% |
| Firehol_SSLProxies_1d | blocklist | ipv4 | 369 | 2 | 0.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 319 | 0.4% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 37 | 0.3% |
| Greensnow | blocklist | ipv4 | 5.1K | 13 | 0.3% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 11 | 0.1% |

---

### ScriptzTeam_BadIPS

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 2.6K | **Unique Entries:** 1.2K | **Target Sources:** 7

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BlockListDE_Strong | blocklist | ipv4 | 276 | 93 | 33.7% |
| Greensnow | blocklist | ipv4 | 5.1K | 319 | 6.3% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 72 | 3.6% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 355 | 2.5% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 431 | 0.5% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 14 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 2 | 0.1% |

---

### Sentinel_Greylist

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 9.2K | **Unique Entries:** 0 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 626 | 20.7% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 344 | 17.2% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 1.7K | 12.0% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 1.4K | 9.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 7.2K | 8.9% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 1.5K | 7.4% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 27 | 6.0% |
| Greensnow | blocklist | ipv4 | 5.1K | 296 | 5.8% |
| BlockListDE_Strong | blocklist | ipv4 | 276 | 15 | 5.4% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 28 | 4.6% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 254 | 1.5% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 27 | 1.2% |
| ScriptzTeam_BadIPS | blocklist | ipv4 | 2.6K | 14 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 707 | 0.3% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 1 | 0.2% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 4 | 0.2% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 16 | 0.1% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.1K | **Unique Entries:** 955 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | adguard | 21.5K | 89 | 0.4% |

---

### ShadowWhisperer's Dating List

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.1K | **Unique Entries:** 911 | **Target Sources:** 1

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 89 | 0.4% |

---

### ShadowWhisperer_Allowlist

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 650 | **Unique Entries:** 399 | **Target Sources:** 17

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdguardTeam_HttpsExclusions_windows | allowlist | domain | 7 | 1 | 14.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 49 | 6.3% |
| AdguardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 1 | 5.6% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 74 | 3.3% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 3 | 1.7% |
| Easy Privacy | allowlist | domain_adguard | 654 | 10 | 1.5% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 2 | 1.2% |
| tranco | allowlist | domain_top | 1.0K | 11 | 1.1% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| WaLLy3K | blocklist | domain | 350 | 3 | 0.9% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 1 | 0.4% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 14 | 0.1% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 5 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 12 | 0.1% |

---

### ShadowWhisperer_BlockLists Ads

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 23.3K | **Unique Entries:** 9.7K | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| YousList-AdGuard | allowlist | domain_adguard | 12 | 2 | 16.7% |
| WaLLy3K | blocklist | domain | 350 | 53 | 15.1% |
| YousList | blocklist | hostname | 624 | 85 | 13.6% |
| hufilter | blocklist | hostname | 100 | 10 | 10.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.7K | 8.9% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 968 | 6.2% |
| tranco | allowlist | domain_top | 1.0K | 61 | 6.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 7 | 4.7% |
| Easy Privacy | allowlist | domain_adguard | 654 | 27 | 4.1% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 5 | 2.8% |
| HaGeZi Pro | blocklist | domain | 409.8K | 8.3K | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 4 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 15 | 0.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 1.9K | 0.6% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 78 | 0.4% |
| Frogeye trackers | blocklist | hostname | 33.3K | 26 | 0.1% |

---

### ShadowWhisperer_BlockLists Adult

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 274.1K | **Unique Entries:** 229.8K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 15.6K | 72.7% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 27.3K | 8.7% |
| tranco | allowlist | domain_top | 1.0K | 10 | 1.0% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 409.8K | 795 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 399 | 0.1% |

---

### ShadowWhisperer_BlockLists Malware

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 51.6K | **Unique Entries:** 29.6K | **Target Sources:** 22

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 74 | 49.3% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 1.2K | 6.2% |
| HaGeZi Pro | blocklist | domain | 409.8K | 14.7K | 3.6% |
| YousList | blocklist | hostname | 624 | 19 | 3.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 318 | 2.0% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 3 | 1.3% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 45 | 0.6% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 83 | 0.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 2 | 0.5% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 3.1K | 0.4% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 613 | 0.3% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 46 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 49 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 518 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 4 | 0.2% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.1K | 0.2% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 6 | 0.1% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 1 | 0.1% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 3 | 0.1% |

---

### ShadowWhisperer_BlockLists Scam

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 11.0K | **Unique Entries:** 7.9K | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| PuppyScams | blocklist | domain_custom_html_puppyscams | 102 | 10 | 9.8% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 42 | 0.5% |
| Spam404 | blocklist | domain | 8.1K | 34 | 0.4% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 1.7K | 0.4% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 991 | 0.3% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 48 | 0.2% |
| BlahDNS_whitelist | allowlist | domain | 773 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 225 | 0.1% |

---

### ShadowWhisperer_UrlShortener

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 4.7K | **Unique Entries:** 356 | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 4.1K | 90.7% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 53 | 22.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 15 | 3.6% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 2 | 0.1% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 19 | 0.1% |

---

### Spam404

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 8.1K | **Unique Entries:** 7.5K | **Target Sources:** 6

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| quidsup_notrack-malware | blocklist | domain | 150 | 2 | 1.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 34 | 0.3% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| HaGeZi Pro | blocklist | domain | 409.8K | 355 | 0.1% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 45 | 0.1% |

---

### Stamparm_Blackbook

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 18.1K | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 2.4K | 48.5% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 2.6K | 19.7% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 17.6K | 7.9% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 47 | 4.8% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 5.0K | 0.7% |
| HaGeZi Pro | blocklist | domain | 409.8K | 1.1K | 0.3% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 955 | 0.2% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 16 | 0.1% |

---

### StevenBlack_Fake_Gambling_Porn

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 313.6K | **Unique Entries:** 141.3K | **Target Sources:** 43

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 7.8K | 41.2% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 90.8K | 40.8% |
| OISD Blocklist NSFW Small | blocklist | domain_adguard | 21.5K | 8.4K | 39.1% |
| YousList | blocklist | hostname | 624 | 243 | 38.9% |
| WaLLy3K | blocklist | domain | 350 | 86 | 24.6% |
| quidsup_notrack-malware | blocklist | domain | 150 | 26 | 17.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 101 | 15.4% |
| hufilter | blocklist | hostname | 100 | 14 | 14.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 22 | 12.2% |
| tranco | allowlist | domain_top | 1.0K | 107 | 10.7% |
| ShadowWhisperer_BlockLists Adult | blocklist | domain | 274.1K | 27.3K | 9.9% |
| ShadowWhisperer_BlockLists Scam | blocklist | domain | 11.0K | 991 | 9.0% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 1.3K | 8.3% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 63 | 8.2% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 1.9K | 8.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 25.7K | 6.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 32 | 4.9% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 95 | 4.2% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| ShadowWhisperer's Dating List | blocklist | domain_adguard | 1.1K | 23 | 2.2% |
| Torrent Trackers | blocklist | domain | 479 | 9 | 1.9% |
| HaGeZi Gambling Only Domains | blocklist | domain | 185.3K | 3.3K | 1.8% |
| Frogeye trackers | blocklist | hostname | 33.3K | 568 | 1.7% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 518 | 1.0% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 1 | 1.0% |
| Spam404 | blocklist | domain | 8.1K | 63 | 0.8% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 2 | 0.8% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 13 | 0.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 2 | 0.7% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 29 | 0.4% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 48 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 14 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 84 | 0.3% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 30 | 0.2% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 31 | 0.2% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 1.6K | 0.2% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 1 | 0.2% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 6 | 0.1% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 472 | 0.1% |
| youtube_GoodbyeAds | blocklist | hostname | 97.6K | 74 | 0.1% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 382 | 0.1% |

---

### Torrent Trackers

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 479 | **Unique Entries:** 462 | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 1 | 0.1% |

---

### tranco

**List Type:** allowlist | **Source Type:** domain | **Total Entries:** 1.0K | **Unique Entries:** 32 | **Target Sources:** 25

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| AdguardTeam_HttpsExclusions_mac | allowlist | domain | 11 | 3 | 27.3% |
| Dogino_Discord_Official | allowlist | domain | 43 | 10 | 23.3% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 268 | 11.9% |
| AdguardTeam_HttpsExclusions_firefox | allowlist | domain | 18 | 2 | 11.1% |
| Easy Privacy | allowlist | domain_adguard | 654 | 60 | 9.2% |
| AdguardTeam_HttpsExclusions_android | allowlist | domain | 97 | 4 | 4.1% |
| BlahDNS_whitelist | allowlist | domain | 773 | 24 | 3.1% |
| hufilter | blocklist | hostname | 100 | 3 | 3.0% |
| Local Allowlist (Domain) | allowlist | domain | 40 | 1 | 2.5% |
| AdguardTeam_HttpsExclusions_sensitive | allowlist | domain | 164 | 4 | 2.4% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 5 | 2.1% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 11 | 1.7% |
| AdguardTeam_HttpsExclusions_issues | allowlist | domain | 68 | 1 | 1.5% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 6 | 1.4% |
| WaLLy3K | blocklist | domain | 350 | 4 | 1.1% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 117 | 0.7% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 106 | 0.6% |
| YousList | blocklist | hostname | 624 | 3 | 0.5% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 61 | 0.3% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 7 | 0.2% |
| HaGeZi Encrypted DNS Servers | blocklist | domain_adguard | 1.4K | 3 | 0.2% |
| DoH_IP_blocklists | blocklist | domain_comment | 1.3K | 1 | 0.1% |
| AdguardTeam_HttpsExclusions_banks | allowlist | domain | 4.0K | 5 | 0.1% |

---

### Ukrainian Ad Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 1.4K | **Unique Entries:** 1.3K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| OISD Blocklist Small | blocklist | adguard | 44.5K | 29 | 0.1% |
| AdGuard Base filter | blocklist | adguard | 94.4K | 54 | 0.1% |

---

### Ukrainian Privacy Filter

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 345 | **Unique Entries:** 96 | **Target Sources:** 3

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Easy Privacy | blocklist | adguard | 53.2K | 156 | 0.3% |
| GetAdmiral Domains Filter List | blocklist | adguard | 1.8K | 1 | 0.1% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 39 | 0.1% |

---

### Ukrainian Security Filter

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 1.7K | **Unique Entries:** 1.2K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| HaGeZi Gambling Only Domains | blocklist | domain | 185.3K | 93 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 376 | 0.1% |

---

### URLHaus_Text

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 16.9K | **Unique Entries:** 15.1K | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 254 | 2.8% |
| EmergingThreats_CompromisedIPs | blocklist | ipv4 | 450 | 3 | 0.7% |
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 29 | 0.6% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 67 | 0.5% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 101 | 0.5% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 973 | 0.5% |
| Firehol_level2 | blocklist | ipv4_cidr_expand | 16.2K | 67 | 0.4% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 13 | 0.4% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 275 | 0.3% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 1 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 24 | 0.2% |
| Firehol_SocksProxy_7d | blocklist | ipv4 | 2.5K | 4 | 0.2% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 2 | 0.1% |

---

### USOM-Blocklists-domains

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 402.6K | **Unique Entries:** 354.7K | **Target Sources:** 30

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 3.7K | 21.3% |
| Viriback_Dump | blocklist | domain_csv_http_url_find | 4.9K | 809 | 16.5% |
| Korlabs_UrlShortener | blocklist | domain | 237 | 16 | 6.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 9 | 6.0% |
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 746 | 5.7% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 955 | 5.3% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 29.5K | 4.3% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 24 | 2.4% |
| Boutetnico_URL_Shorteners | blocklist | domain | 418 | 8 | 1.9% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 3.9K | 1.8% |
| jarelllama_Scam-Blocklist | blocklist | domain | 457.7K | 4.6K | 1.0% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 5 | 0.8% |
| HaGeZi Pro | blocklist | domain | 409.8K | 2.8K | 0.7% |
| OpenPhish_Feed | blocklist | domain_http_url | 279 | 2 | 0.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| YousList | blocklist | hostname | 624 | 4 | 0.6% |
| Blocklists UT1 Shortener | blocklist | domain | 4.5K | 23 | 0.5% |
| Blocklists UT1 Cryptojacking | blocklist | domain | 16.3K | 76 | 0.5% |
| Torrent Trackers | blocklist | domain | 479 | 2 | 0.4% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| ShadowWhisperer_UrlShortener | blocklist | domain | 4.7K | 14 | 0.3% |
| WaLLy3K | blocklist | domain | 350 | 1 | 0.3% |
| ShadowWhisperer_BlockLists Malware | blocklist | domain | 51.6K | 147 | 0.3% |
| malware-filter_phishing-filter | blocklist | hostname | 25.8K | 43 | 0.2% |
| Ukrainian Security Filter | blocklist | domain | 1.7K | 3 | 0.2% |
| Spam404 | blocklist | domain | 8.1K | 7 | 0.1% |
| FakeWebshopListHUN | blocklist | domain | 8.2K | 6 | 0.1% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 19 | 0.1% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 382 | 0.1% |

---

### USOM-Blocklists-ips

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 12.5K | **Unique Entries:** 7.4K | **Target Sources:** 18

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Viriback_Dump | blocklist | ipv4_csv_http_url_find | 4.6K | 267 | 5.8% |
| Firehol_CleanTalk_Top20 | blocklist | ipv4 | 20 | 1 | 5.0% |
| CybercrimeTracker_CCPMGate | blocklist | ipv4_http_url | 103 | 3 | 2.9% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 4.4K | 2.2% |
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 50 | 1.7% |
| Sblam_Blocklist | blocklist | ipv4 | 1.9K | 11 | 0.6% |
| BinaryDefense_Banlist | blocklist | ipv4 | 3.0K | 14 | 0.5% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 67 | 0.4% |
| Firehol_CleanTalk | blocklist | ipv4 | 494 | 2 | 0.4% |
| Firehol_level3 | blocklist | ipv4_cidr_expand | 20.1K | 51 | 0.3% |
| Sentinel_Greylist | blocklist | ipv4_find | 9.2K | 16 | 0.2% |
| Firehol_GPF_Comics | blocklist | ipv4 | 2.3K | 4 | 0.2% |
| Greensnow | blocklist | ipv4 | 5.1K | 9 | 0.2% |
| Borestad_AbuseIPDB_S100_3d | blocklist | ipv4_find | 81.7K | 133 | 0.2% |
| BlockListDE_Brute | blocklist | ipv4 | 613 | 1 | 0.2% |
| CINSScore_BadGuys_Army | blocklist | ipv4 | 15.0K | 15 | 0.1% |
| Rutgers_DROP | blocklist | ipv4 | 2.0K | 1 | 0.1% |
| Firehol_level2 | blocklist | ipv4 | 14.2K | 17 | 0.1% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** ipv4 | **Total Entries:** 4.6K | **Unique Entries:** 0 | **Target Sources:** 4

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | ipv4_url | 2.9K | 488 | 17.0% |
| USOM-Blocklists-ips | blocklist | ipv4 | 12.5K | 267 | 2.1% |
| Maltrail_StaticTrails | blocklist | ipv4_find | 203.4K | 4.0K | 2.0% |
| URLHaus_Text | blocklist | ipv4_http_url | 16.9K | 29 | 0.2% |

---

### Viriback_Dump

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 4.9K | **Unique Entries:** 0 | **Target Sources:** 8

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| CybercrimeTracker_All | blocklist | domain_url | 13.0K | 1.8K | 13.8% |
| Stamparm_Blackbook | blocklist | domain_custom_csv_blackbook | 18.1K | 2.4K | 13.1% |
| Blocklists UT1 Malware | blocklist | domain | 222.5K | 2.1K | 1.0% |
| Maltrail_StaticTrails | blocklist | domain_custom_csv_maltrail | 690.4K | 4.0K | 0.6% |
| USOM-Blocklists-domains | blocklist | domain | 402.6K | 809 | 0.2% |
| CybercrimeTracker_CCPMGate | blocklist | domain_http_url | 982 | 1 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 296 | 0.1% |
| cyberhost_malware-blocklist | blocklist | domain | 17.2K | 17 | 0.1% |

---

### WaLLy3K

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 350 | **Unique Entries:** 0 | **Target Sources:** 13

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList | blocklist | hostname | 624 | 9 | 1.4% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| Easy Privacy | allowlist | domain_adguard | 654 | 5 | 0.8% |
| quidsup_notrack-malware | blocklist | domain | 150 | 1 | 0.7% |
| BlahDNS_whitelist | allowlist | domain | 773 | 5 | 0.6% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 1 | 0.6% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 3 | 0.5% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 85 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 4 | 0.4% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 41 | 0.3% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 53 | 0.2% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 5 | 0.2% |

---

### YousList

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 624 | **Unique Entries:** 0 | **Target Sources:** 11

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| WaLLy3K | blocklist | domain | 350 | 9 | 2.6% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 204 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| ShadowWhisperer_BlockLists Ads | blocklist | domain | 23.3K | 85 | 0.4% |
| tranco | allowlist | domain_top | 1.0K | 3 | 0.3% |
| BlahDNS_whitelist | allowlist | domain | 773 | 2 | 0.3% |
| Easy Privacy | allowlist | domain_adguard | 654 | 1 | 0.2% |
| StevenBlack_Fake_Gambling_Porn | blocklist | hostname | 313.6K | 243 | 0.1% |
| HaGeZi Pro | blocklist | domain | 409.8K | 444 | 0.1% |
| quidsup_notrack-tracker | blocklist | domain | 15.7K | 22 | 0.1% |

---

### YousList-AdGuard

**List Type:** blocklist | **Source Type:** adguard | **Total Entries:** 7.3K | **Unique Entries:** 7.3K | **Target Sources:** 2

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Ukrainian Privacy Filter | blocklist | adguard | 345 | 1 | 0.3% |
| OISD Blocklist Small | blocklist | adguard | 44.5K | 25 | 0.1% |

---

### youtube_GoodbyeAds

**List Type:** blocklist | **Source Type:** domain | **Total Entries:** 97.6K | **Unique Entries:** 97.4K | **Target Sources:** 12

**Overlap with Other Sources:**

| Target Source | List Type | Source Type | Target Count | Overlap Count | Overlap % |
|---------------|-----------|-------------|--------------|---------------|----------|
| Local Blocklist (Domain) | blocklist | domain | 1 | 1 | 100.0% |
| YousList-AdGuard | allowlist | domain_adguard | 12 | 1 | 8.3% |
| WaLLy3K | blocklist | domain | 350 | 7 | 2.0% |
| AdGuard DNS filter | allowlist | domain_adguard | 180 | 2 | 1.1% |
| hufilter | blocklist | hostname | 100 | 1 | 1.0% |
| YousList | blocklist | hostname | 624 | 5 | 0.8% |
| BlahDNS_whitelist | allowlist | domain | 773 | 4 | 0.5% |
| tranco | allowlist | domain_top | 1.0K | 5 | 0.5% |
| fabriziosalmi_allowlist | allowlist | domain | 2.3K | 8 | 0.4% |
| Easy Privacy | allowlist | domain_adguard | 654 | 2 | 0.3% |
| bigdargon_hostsVN | blocklist | hostname | 19.0K | 52 | 0.3% |
| ShadowWhisperer_Allowlist | allowlist | domain_with_comment_suffix | 650 | 2 | 0.3% |

---

## About

This overlap analysis is automatically generated by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) to help understand relationships between different DNS sources. High overlap percentages may indicate redundant sources, while low overlap percentages suggest unique content.

**Note:** Overlap percentages are calculated as: (overlap_count / source_total_count) × 100

