# DNS Toolkit - Daily Processing Results

This branch contains the daily processed and consolidated DNS blocklists and allowlists.

**Last Updated:** 2025-08-09 14:27:33 UTC

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
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_ipv4_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_adguard_blocklist.txt
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
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/cryptocurrency_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/dns_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/dns_ipv6_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fake_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fake_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fakenews_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/gambling_domain_blocklist.txt
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
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_domain_allowlist.txt
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
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_blocklist.txt

# High-confidence lists (top entries by number of sources)
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min8.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min7.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min6.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min3.txt
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
| Total Sources | 92 |
| Successful Downloads | 92 |
| Failed Downloads | 0 |
| Success Rate | 100.0% |
| Last Update | 20250809_142534 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 23 |
| cidr_ipv4 | 3 |
| domain | 21 |
| domain_adguard | 15 |
| domain_comment | 1 |
| domain_csv_http_url_find | 1 |
| domain_custom_csv_blackbook | 1 |
| domain_custom_csv_maltrail | 1 |
| domain_custom_html_puppyscams | 1 |
| domain_http_url | 2 |
| domain_top | 1 |
| domain_url | 1 |
| domain_with_comment_suffix | 1 |
| hostname | 7 |
| ipv4 | 25 |
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
| adguard | 33 | 22 | 55 |
| cidr_ipv4 | 2 | 1 | 3 |
| domain | 47 | 30 | 77 |
| ipv4 | 37 | 12 | 49 |
| ipv6 | 3 | 3 | 6 |
| **Last Update** | | | 20250809_142543 |

### Consolidation Statistics

| Type | Blocklist Entries | Allowlist Entries | Total Files |
|------|-------------------|-------------------|-------------|
| adguard | 291.9K (-2 ignored) | 1.2K | 31 |
| cidr_ipv4 | 4.7K | - | 2 |
| domain | 989.4K (-678 ignored) | 4.4K | 37 |
| ipv4 | 126.5K (-4 ignored) | 62 | 29 |
| ipv6 | 68 | - | 2 |
| **Last Update** | | | 20250809_142544 |

### Size Groups Summary

| Group | Total Entries |
|-------|---------------|
| big | 2.4M |
| lite | 229.3K |
| mini | 11.6K |
| normal | 345.4K |
| **Last Update** | 20250809_142549 |

### Categories Summary

| Category | Total Entries |
|----------|---------------|
| ads | 1.0M |
| adult | 575.6K |
| annoyance | 5.9K |
| anonymizer | 2.8K |
| botnet | 16.3K |
| cryptocurrency | 67.9K |
| dns | 62.7K |
| doh | 6.2K |
| fake | 321.7K |
| fakenews | 312.6K |
| gambling | 492.6K |
| malicious | 547.5K |
| malware | 1.7M |
| others | 3.7K |
| phishing | 1.3M |
| privacy | 57.6K |
| proxy | 2.8K |
| ransomware | 44.4K |
| scam | 479.4K |
| spam | 10.5K |
| threat | 1.4M |
| trackers | 660.6K |
| **Last Update** | 20250809_142600 |

### Overlap Analysis Summary

| Metric | Count |
|--------|-------|
| Total Sources Analyzed | 99 |
| **Last Update** | 2025-08-09 14:27:33 |

**[View Detailed Overlap Analysis →](overlap.md)**

### Top Entries Summary

| Type | List Type | Min Sources | Entries Count | Files Generated |
|------|-----------|-------------|---------------|----------------|
| adguard | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min5.txt) | 5 | 1 |
| adguard | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min4.txt) | 234 | 1 |
| adguard | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min3.txt) | 30.8K | 1 |
| domain | allowlist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min4.txt) | 3 | 1 |
| domain | allowlist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min3.txt) | 38 | 1 |
| domain | blocklist | [8](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min8.txt) | 6 | 1 |
| domain | blocklist | [7](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min7.txt) | 41 | 1 |
| domain | blocklist | [6](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min6.txt) | 404 | 1 |
| domain | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min5.txt) | 2.2K | 1 |
| domain | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min4.txt) | 7.4K | 1 |
| domain | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min3.txt) | 41.1K | 1 |
| ipv4 | blocklist | [9](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min9.txt) | 5 | 1 |
| ipv4 | blocklist | [8](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min8.txt) | 20 | 1 |
| ipv4 | blocklist | [7](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min7.txt) | 126 | 1 |
| ipv4 | blocklist | [6](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min6.txt) | 757 | 1 |
| ipv4 | blocklist | [5](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min5.txt) | 2.5K | 1 |
| ipv4 | blocklist | [4](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min4.txt) | 7.2K | 1 |
| ipv4 | blocklist | [3](https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min3.txt) | 21.5K | 1 |
| **Last Update** | | | | 2025-08-09 14:27:33 |

## About

These lists are automatically generated daily by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) from multiple reputable sources.

