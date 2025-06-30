# DNS Toolkit - Daily Processing Results

This branch contains the daily processed and consolidated DNS blocklists and allowlists.

**Last Updated:** 2025-07-19 14:27:39 UTC

## Quick Start

Add any of these URLs to your DNS filtering solution:

```
# Blocklists or Allowlists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/adguard_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/adguard_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/cidr_ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/ipv4_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/ipv6_blocklist.txt

# Size-based lists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/big_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/lite_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/mini_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/groups/normal_domain_blocklist.txt

# Category-based lists
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/ads_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/adult_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/annoyance_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/annoyance_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/anonymizer_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/dns_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/doh_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/fakenews_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/gambling_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/malware_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/mobile_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/others_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/phishing_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/privacy_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/proxy_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/scam_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/security_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/spam_domain_blocklist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_allowlist.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/categories/trackers_domain_blocklist.txt

# High-confidence lists (top entries)
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_adguard_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_cidr_ipv4_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min10.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min9.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min8.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min7.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min6.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min5.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min4.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_blocklist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_domain_allowlist_min3.txt
https://raw.githubusercontent.com/phani-kb/dns-toolkit/output/top/top_ipv4_blocklist_min10.txt
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
| Total Sources | 90 |
| Successful Downloads | 90 |
| Failed Downloads | 0 |
| Success Rate | 100.0% |
| Last Update | 20250719_142626 |

**Sources by Type:**

| Source Type | Count |
|-------------|-------|
| adguard | 21 |
| cidr_ipv4 | 3 |
| domain | 18 |
| domain_adguard | 14 |
| domain_comment | 1 |
| domain_csv_http_url_find | 1 |
| domain_custom_csv_blackbook | 1 |
| domain_custom_csv_maltrail | 1 |
| domain_custom_html_ccam | 1 |
| domain_custom_html_puppyscams | 1 |
| domain_http_url | 3 |
| domain_top | 1 |
| domain_url | 1 |
| domain_with_comment_suffix | 1 |
| hostname | 10 |
| ipv4 | 23 |
| ipv4_cidr_expand | 6 |
| ipv4_custom_html_ccam | 1 |
| ipv4_find | 5 |
| ipv4_http_url | 2 |
| ipv4_range_expand | 1 |
| ipv4_url | 1 |
| ipv6 | 1 |
| ipv6_find | 1 |
| ipv6_htaccess | 1 |

### Processing Statistics

| Source Type | Valid Files | Invalid Files | Total |
|-------------|-------------|---------------|-------|
| adguard | 31 | 23 | 54 |
| cidr_ipv4 | 3 | 1 | 4 |
| domain | 46 | 27 | 73 |
| ipv4 | 37 | 13 | 50 |
| ipv6 | 1 | 1 | 2 |
| **Last Update** | | | 20250719_142635 |

### Consolidation Statistics

| Type | Blocklist Entries | Allowlist Entries | Total Files |
|------|-------------------|-------------------|-------------|
| adguard | 258.7K (-6 ignored) | 7.0K | 31 |
| cidr_ipv4 | 4.6K | - | 3 |
| domain | 1.7M (-1.2K ignored) | 7.2K | 46 |
| ipv4 | 647.9K | - | 37 |
| ipv6 | 11 | - | 1 |
| **Last Update** | | | 20250719_142635 |

### Size Groups Summary

| Group | Total Entries |
|-------|---------------|
| big | 2.7M |
| lite | 227.5K |
| mini | 10.9K |
| normal | 264.0K |
| **Last Update** | 20250719_142643 |

### Categories Summary

| Category | Total Entries |
|----------|---------------|
| ads | 999.6K |
| adult | 304.9K |
| annoyance | 2.4K |
| anonymizer | 3.0K |
| dns | 6.2K |
| doh | 2.9K |
| fakenews | 304.9K |
| gambling | 304.9K |
| malware | 441.9K |
| mobile | 6.5K |
| others | 1.5M |
| phishing | 243.7K |
| privacy | 346 |
| proxy | 3.0K |
| scam | 30.8K |
| security | 1.7K |
| spam | 422 |
| trackers | 149.0K |
| **Last Update** | 20250719_142648 |

### Overlap Analysis Summary

| Metric | Count |
|--------|-------|
| Total Sources Analyzed | 95 |
| **Last Update** | 2025-07-19 14:27:39 |

**[View Detailed Overlap Analysis →](overlap.md)**

### Top Entries Summary

| Type | List Type | Min Sources | Entries Count | Files Generated |
|------|-----------|-------------|---------------|----------------|
| adguard | blocklist | 4 | 2 | 1 |
| adguard | blocklist | 3 | 130 | 1 |
| cidr_ipv4 | blocklist | 3 | 1.4K | 1 |
| domain | allowlist | 5 | 3 | 1 |
| domain | allowlist | 4 | 11 | 1 |
| domain | allowlist | 3 | 69 | 1 |
| domain | blocklist | 10 | 1 | 1 |
| domain | blocklist | 9 | 4 | 1 |
| domain | blocklist | 8 | 29 | 1 |
| domain | blocklist | 7 | 154 | 1 |
| domain | blocklist | 6 | 635 | 1 |
| domain | blocklist | 5 | 2.1K | 1 |
| domain | blocklist | 4 | 7.4K | 1 |
| domain | blocklist | 3 | 32.4K | 1 |
| ipv4 | blocklist | 10 | 1 | 1 |
| ipv4 | blocklist | 9 | 5 | 1 |
| ipv4 | blocklist | 8 | 83 | 1 |
| ipv4 | blocklist | 7 | 457 | 1 |
| ipv4 | blocklist | 6 | 1.5K | 1 |
| ipv4 | blocklist | 5 | 3.3K | 1 |
| ipv4 | blocklist | 4 | 8.8K | 1 |
| ipv4 | blocklist | 3 | 26.5K | 1 |
| **Last Update** | | | | 2025-07-19 14:27:39 |

## About

These lists are automatically generated daily by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) from multiple reputable sources.

