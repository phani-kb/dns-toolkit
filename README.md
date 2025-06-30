# DNS Toolkit - Daily Processing Results

This branch contains the daily processed and consolidated DNS blocklists and allowlists.

**Last Updated:** 2025-08-16 17:10:36 UTC

## Quick Start

Add any of these URLs to your DNS filtering solution:

```
# Consolidated Blocklists or Allowlists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/ipv6_blocklist.txt

# Size-based lists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_ipv6_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_ipv4_blocklist.txt

# Category-based lists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/adult_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/adult_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/annoyance_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/annoyance_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/annoyance_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/anonymizer_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/botnet_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/botnet_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/browser_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/discord_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/dns_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/dns_ipv6_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fake_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fake_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fakenews_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/finance_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/gambling_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/issues_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/mac_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malicious_ipv6_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/mobile_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/phishing_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/phishing_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/proxy_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ransomware_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/scam_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/scam_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/spam_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/spam_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/spam_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/threat_cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/threat_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/threat_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/threat_ipv6_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/torrent_trackers_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/url_shorteners_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/windows_domain_allowlist.txt

# High-confidence lists (top entries by number of sources)
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min8.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min7.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min6.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min9.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min8.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min7.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min6.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min3.txt

```

## Daily Workflow Summary

### Download Statistics

| Metric | Count |
|--------|-------|
| Total Sources | 113 |
| Successful Downloads | 113 |
| Failed Downloads | 0 |
| Success Rate | 100.0% |
| Last Update | 20250816_170828 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 25 |
| cidr_ipv4 | 3 |
| domain | 37 |
| domain_adguard | 15 |
| domain_comment | 1 |
| domain_csv_http_url_find | 1 |
| domain_custom_csv_blackbook | 1 |
| domain_custom_csv_maltrail | 1 |
| domain_custom_html_puppyscams | 1 |
| domain_http_url | 2 |
| domain_top | 1 |
| domain_url | 3 |
| domain_with_comment_suffix | 1 |
| hostname | 7 |
| ipv4 | 26 |
| ipv4_cidr_expand | 5 |
| ipv4_csv_http_url_find | 1 |
| ipv4_find | 4 |
| ipv4_http_url | 2 |
| ipv4_url | 1 |
| ipv6 | 1 |
| ipv6_find | 1 |
| ipv6_htaccess | 1 |

### Processing Statistics

| Source Type | Valid Files | Invalid Files | Total |
|-------------|-------------|---------------|-------|
| adguard | 34 | 23 | 57 |
| cidr_ipv4 | 2 | 1 | 3 |
| domain | 60 | 39 | 99 |
| ipv4 | 38 | 13 | 51 |
| ipv6 | 3 | 3 | 6 |
| **Last Update** | | | 20250816_170838 |

### Consolidation Statistics

| Type | Blocklist Entries | Allowlist Entries | Total Files |
|------|-------------------|-------------------|-------------|
| adguard | 288.6K (-14 ignored) | 1.5K | 32 |
| cidr_ipv4 | 4.7K | - | 2 |
| domain | 997.6K (-587 ignored) | 7.5K (-1 ignored) | 43 |
| ipv4 | 121.9K (-3 ignored) | 62 | 30 |
| ipv6 | 68 | - | 2 |
| **Last Update** | | | 20250816_170838 |

### Size Groups Summary

| Group | Total Entries |
|-------|---------------|
| big | 2.3M |
| lite | 225.8K |
| mini | 10.3K |
| normal | 677.5K |
| **Last Update** | 20250816_170845 |

### Categories Summary

| Category | Total Entries |
|----------|---------------|
| ads | 1.0M |
| adult | 586.7K |
| annoyance | 5.9K |
| anonymizer | 2.9K |
| botnet | 16.3K |
| browser | 18 |
| cryptocurrency | 68.0K |
| discord | 43 |
| dns | 62.7K |
| doh | 6.1K |
| fake | 322.7K |
| fakenews | 313.6K |
| finance | 4.0K |
| gambling | 495.6K |
| issues | 68 |
| mac | 11 |
| malicious | 571.5K |
| malware | 1.7M |
| mobile | 97 |
| others | 4.8K |
| phishing | 1.3M |
| privacy | 57.7K |
| proxy | 2.9K |
| ransomware | 44.5K |
| scam | 479.6K |
| spam | 10.5K |
| threat | 1.4M |
| torrent_trackers | 479 |
| trackers | 661.0K |
| url_shorteners | 5.6K |
| windows | 7 |
| **Last Update** | 20250816_170857 |

### Overlap Analysis Summary

| Metric | Count |
|--------|-------|
| Total Sources Analyzed | 112 |
| **Last Update** | 2025-08-16 17:10:36 |

**[View Detailed Overlap Analysis →](overlap.md)**

### Top Entries Summary

| Type | List Type | Min Sources | Entries Count | Files Generated |
|------|-----------|-------------|---------------|----------------|
| adguard | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min5.txt) | 6 | 1 |
| adguard | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min4.txt) | 291 | 1 |
| adguard | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min3.txt) | 32.4K | 1 |
| domain | allowlist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min5.txt) | 1 | 1 |
| domain | allowlist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min4.txt) | 16 | 1 |
| domain | allowlist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min3.txt) | 92 | 1 |
| domain | blocklist | [8](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min8.txt) | 6 | 1 |
| domain | blocklist | [7](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min7.txt) | 43 | 1 |
| domain | blocklist | [6](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min6.txt) | 406 | 1 |
| domain | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min5.txt) | 2.2K | 1 |
| domain | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min4.txt) | 7.4K | 1 |
| domain | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min3.txt) | 41.0K | 1 |
| ipv4 | blocklist | [9](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min9.txt) | 2 | 1 |
| ipv4 | blocklist | [8](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min8.txt) | 17 | 1 |
| ipv4 | blocklist | [7](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min7.txt) | 137 | 1 |
| ipv4 | blocklist | [6](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min6.txt) | 660 | 1 |
| ipv4 | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min5.txt) | 2.2K | 1 |
| ipv4 | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min4.txt) | 6.7K | 1 |
| ipv4 | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min3.txt) | 22.0K | 1 |
| **Last Update** | | | | 2025-08-16 17:10:36 |

## About

These lists are automatically generated daily by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) from multiple reputable sources.

