package processors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/processors"
	"github.com/phani-kb/multilog"
	"github.com/stretchr/testify/assert"
)

func TestNewIpv4FindProcessor(t *testing.T) {
	t.Parallel()

	sourceType := "ipv4_find"
	listType := "blocklist"

	processor := processors.NewIpv4FindProcessor(sourceType, listType)

	assert.NotNil(t, processor)
	assert.Equal(t, sourceType, processor.GetSourceType())
	assert.Equal(t, listType, processor.GetListType())
}

func TestIpv4FindProcessor_Process(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewIpv4FindProcessor("ipv4_find", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "empty content",
			content:         "",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "single valid IPv4",
			content:         "192.168.1.1",
			expectedValid:   []string{"192.168.1.1"},
			expectedInvalid: nil,
		},
		{
			name:            "multiple valid IPv4s",
			content:         "192.168.1.1\n10.0.0.1\n172.16.0.1",
			expectedValid:   []string{"192.168.1.1", "10.0.0.1", "172.16.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "IPv4 with text around it",
			content:         "Server IP: 192.168.1.1 (production)\nBackup: 10.0.0.1 server",
			expectedValid:   []string{"192.168.1.1", "10.0.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "entries with comments",
			content:         "# This is a comment\n192.168.1.1\n! Another comment\n10.0.0.1",
			expectedValid:   []string{"192.168.1.1", "10.0.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "invalid formats",
			content:         "256.256.256.256\n192.168.1\n999.999.999.999",
			expectedValid:   []string{"256.256.256.256", "999.999.999.999"},
			expectedInvalid: []string{"192.168.1"},
		},
		{
			name:            "IPv6 addresses (should be invalid)",
			content:         "2001:db8::1\n::1\nfe80::1",
			expectedValid:   nil,
			expectedInvalid: []string{"2001:db8::1", "::1", "fe80::1"},
		},
		{
			name:            "malware IPs with port and metadata",
			content:         "198.51.100.22:53,cobaltstrike-2 (malware),(static)\n203.0.113.45:443,cobaltstrike-2 (malware),(static)\n198.51.100.78:8080,trojan-downloader,(dynamic)",
			expectedValid:   []string{"198.51.100.22", "203.0.113.45", "198.51.100.78"},
			expectedInvalid: nil,
		},
		{
			name:            "mixed threat indicators",
			content:         "198.51.100.164:8443,banking-trojan,(zeus)\n203.0.113.60:21,ftp-attack,(brute)\n198.51.100.200:3389,rdp-scanner,(automated)\n203.0.113.42:1433,sql-injection,(attempted)",
			expectedValid:   []string{"198.51.100.164", "203.0.113.60", "198.51.100.200", "203.0.113.42"},
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			assert.Equal(t, tt.expectedValid, valid, "Valid entries should match expected")
			assert.Equal(t, tt.expectedInvalid, invalid, "Invalid entries should match expected")
		})
	}
}

func TestIpv4FindProcessor_Integration(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewIpv4FindProcessor("ipv4_find", "blocklist")

	content := `# IP Address Blacklist

# Malicious IPs
192.168.1.100   # compromised host
10.0.0.50       # botnet C&C
172.16.0.25     # malware server

# Threat Intelligence Feed
198.51.100.22:53,cobaltstrike-2 (malware),(static)
203.0.113.45:443,cobaltstrike-2 (malware),(static)
203.0.113.89:22,ssh-bruteforce,(honeypot)
198.51.100.123:80,botnet-cc,(verified)

# Invalid entries
256.256.256.256  # out of range
192.168.1        # incomplete
not.an.ip.address

# More valid IPs
8.8.8.8          # Google DNS (example)
1.1.1.1          # Cloudflare DNS
208.67.222.222   # OpenDNS

# Mixed content
Server 192.168.1.1 is down
Connect to 10.0.0.1:8080 for admin
Check logs at http://172.16.0.1/logs

# Additional threat indicators
198.51.100.164:8443,banking-trojan,(zeus)
203.0.113.60:21,ftp-attack,(brute)

# IPv6 (should be invalid for this processor)
2001:db8::1
::1

# Edge cases
0.0.0.0
127.0.0.1
255.255.255.255`

	expectedValid := []string{
		"192.168.1.100",
		"10.0.0.50",
		"172.16.0.25",
		"198.51.100.22",
		"203.0.113.45",
		"203.0.113.89",
		"198.51.100.123",
		"256.256.256.256",
		"8.8.8.8",
		"1.1.1.1",
		"208.67.222.222",
		"192.168.1.1",
		"10.0.0.1",
		"172.16.0.1",
		"198.51.100.164",
		"203.0.113.60",
		"0.0.0.0",
		"127.0.0.1",
		"255.255.255.255",
	}

	expectedInvalid := []string{
		"192.168.1        # incomplete",
		"not.an.ip.address",
		"2001:db8::1",
		"::1",
	}

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Valid entries should match expected")
	assert.Equal(t, expectedInvalid, invalid, "Invalid entries should match expected")
}

func TestIpv4FindProcessor_EdgeCases(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewIpv4FindProcessor("ipv4_find", "blocklist")

	tests := []struct {
		name            string
		content         string
		expectedValid   []string
		expectedInvalid []string
	}{
		{
			name:            "entries with whitespace",
			content:         "  192.168.1.1  \n\t10.0.0.1\t",
			expectedValid:   []string{"192.168.1.1", "10.0.0.1"},
			expectedInvalid: nil,
		},
		{
			name:            "only comments",
			content:         "# Comment 1\n! Comment 2\n// Comment 3",
			expectedValid:   nil,
			expectedInvalid: nil,
		},
		{
			name:            "partial IP addresses",
			content:         "192\n192.168\n192.168.1\n192.168.1.1",
			expectedValid:   []string{"192.168.1.1"},
			expectedInvalid: []string{"192", "192.168", "192.168.1"},
		},
		{
			name:            "boundary values",
			content:         "0.0.0.0\n255.255.255.255\n256.0.0.0",
			expectedValid:   []string{"0.0.0.0", "255.255.255.255", "256.0.0.0"},
			expectedInvalid: nil,
		},
		{
			name:            "leading zeros in octets",
			content:         "192.168.001.001\n010.000.000.001",
			expectedValid:   []string{"192.168.001.001", "010.000.000.001"},
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, invalid := processor.Process(logger, tt.content)
			assert.Equal(t, tt.expectedValid, valid, "Valid entries should match expected")
			assert.Equal(t, tt.expectedInvalid, invalid, "Invalid entries should match expected")
		})
	}
}

func TestIpv4FindProcessor_Registration(t *testing.T) {
	t.Parallel()

	processor, exists := processors.Processors.GetProcessor("ipv4_find", "blocklist")
	assert.True(t, exists, "Processor should be registered")
	assert.NotNil(t, processor)

	ipv4Processor, ok := processor.(*processors.Ipv4FindProcessor)
	assert.True(t, ok, "Processor should be of type *Ipv4FindProcessor")
	assert.Equal(t, "ipv4_find", ipv4Processor.GetSourceType())
	assert.Equal(t, "blocklist", ipv4Processor.GetListType())
}

func TestIpv4FindProcessor_Performance(t *testing.T) {
	t.Parallel()

	logger := multilog.NewLogger()
	processor := processors.NewIpv4FindProcessor("ipv4_find", "blocklist")

	var contentBuilder strings.Builder
	expectedValid := make([]string, 0, 1000)

	for i := 0; i < 255; i++ {
		for j := 0; j < 4; j++ {
			ip := fmt.Sprintf("192.168.%d.%d", i, j)
			contentBuilder.WriteString(ip + "\n")
			expectedValid = append(expectedValid, ip)
		}
	}

	content := contentBuilder.String()

	valid, invalid := processor.Process(logger, content)

	assert.Equal(t, expectedValid, valid, "Should process all valid IPs correctly")
	assert.Empty(t, invalid, "Should have no invalid entries for well-formed IPs")
	assert.Equal(t, 1020, len(valid), "Should have processed exactly 1020 IPs")
}
