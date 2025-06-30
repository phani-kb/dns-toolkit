# DNS Toolkit - Daily Processing Results

This branch contains the daily processed and consolidated DNS blocklists and allowlists.

**Last Updated:** 2025-07-12 14:29:29 UTC

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
| Total Sources | 91 |
| Successful Downloads | 89 |
| Failed Downloads | 2 |
| Success Rate | 97.8% |
| Last Update | 20250712_142353 |

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
| hostname | 12 |
| ipv4 | 23 |
| ipv4_cidr_expand | 6 |
| ipv4_custom_html_ccam | 1 |
| ipv4_find | 5 |
| ipv4_http_url | 2 |
| ipv4_url | 1 |
| ipv6 | 1 |
| ipv6_find | 1 |
| ipv6_htaccess | 1 |

**Failed Sources:**
- Dan Pollock's List
- Sblam_Blocklist

### Processing Statistics

| Source Type | Valid Files | Invalid Files | Total |
|-------------|-------------|---------------|-------|
| adguard | 31 | 23 | 54 |
| cidr_ipv4 | 3 | 1 | 4 |
| domain | 43 | 25 | 68 |
| ipv4 | 33 | 11 | 44 |
| ipv6 | 1 | 1 | 2 |
| **Last Update** | | | 20250712_142822 |

### Consolidation Statistics

| Type | Blocklist Entries | Allowlist Entries | Total Files |
|------|-------------------|-------------------|-------------|
| adguard | 252.2K (-6 ignored) | 7.0K | 31 |
| cidr_ipv4 | 4.6K | - | 3 |
| domain | 1.7M (-1.2K ignored) | 7.2K | 43 |
| ipv4 | 644.6K | - | 33 |
| ipv6 | 11 | - | 1 |
| **Last Update** | | | 20250712_142822 |

### Size Groups Summary

| Group | Total Entries |
|-------|---------------|
| big | 2.7M |
| lite | 221.9K |
| mini | 8.7K |
| normal | 246.9K |
| **Last Update** | 20250712_142831 |

### Categories Summary

| Category | Total Entries |
|----------|---------------|
| ads | 1.0M |
| adult | 336.5K |
| annoyance | 2.4K |
| anonymizer | 2.9K |
| dns | 6.2K |
| doh | 2.9K |
| fakenews | 336.5K |
| gambling | 336.5K |
| malware | 437.8K |
| mobile | 6.5K |
| others | 1.7M |
| phishing | 243.4K |
| privacy | 346 |
| proxy | 2.9K |
| scam | 19.7K |
| security | 1.7K |
| spam | 423 |
| trackers | 133.7K |
| **Last Update** | 20250712_142836 |

### Overlap Analysis Summary

| Metric | Count |
|--------|-------|
| Total Sources Analyzed | 88 |
| **Last Update** | 2025-07-12 14:29:29 |

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
| domain | blocklist | 9 | 7 | 1 |
| domain | blocklist | 8 | 85 | 1 |
| domain | blocklist | 7 | 467 | 1 |
| domain | blocklist | 6 | 1.6K | 1 |
| domain | blocklist | 5 | 6.5K | 1 |
| domain | blocklist | 4 | 21.6K | 1 |
| domain | blocklist | 3 | 70.3K | 1 |
| ipv4 | blocklist | 8 | 11 | 1 |
| ipv4 | blocklist | 7 | 53 | 1 |
| ipv4 | blocklist | 6 | 325 | 1 |
| ipv4 | blocklist | 5 | 1.7K | 1 |
| ipv4 | blocklist | 4 | 7.1K | 1 |
| ipv4 | blocklist | 3 | 23.2K | 1 |
| **Last Update** | | | | 2025-07-12 14:29:29 |

## About

These lists are automatically generated daily by the [DNS Toolkit](https://github.com/phani-kb/dns-toolkit) from multiple reputable sources.

