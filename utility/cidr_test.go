package utility

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWhetherCidrContainsIp(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		ip       string
		expected bool
		hasError bool
	}{
		{
			name:     "IP within CIDR range",
			cidr:     "192.168.1.0/24",
			ip:       "192.168.1.10",
			expected: true,
			hasError: false,
		},
		{
			name:     "IP outside CIDR range",
			cidr:     "192.168.1.0/24",
			ip:       "192.168.2.10",
			expected: false,
			hasError: false,
		},
		{
			name:     "IP at network address",
			cidr:     "192.168.1.0/24",
			ip:       "192.168.1.0",
			expected: true,
			hasError: false,
		},
		{
			name:     "IP at broadcast address",
			cidr:     "192.168.1.0/24",
			ip:       "192.168.1.255",
			expected: true,
			hasError: false,
		},
		{
			name:     "Single host CIDR",
			cidr:     "192.168.1.10/32",
			ip:       "192.168.1.10",
			expected: true,
			hasError: false,
		},
		{
			name:     "Single host CIDR - different IP",
			cidr:     "192.168.1.10/32",
			ip:       "192.168.1.11",
			expected: false,
			hasError: false,
		},
		{
			name:     "IPv6 CIDR",
			cidr:     "2001:db8::/32",
			ip:       "2001:db8::1",
			expected: true,
			hasError: false,
		},
		{
			name:     "IPv6 CIDR - outside range",
			cidr:     "2001:db8::/32",
			ip:       "2001:db9::1",
			expected: false,
			hasError: false,
		},
		{
			name:     "Invalid CIDR format",
			cidr:     "invalid.cidr",
			ip:       "192.168.1.10",
			expected: false,
			hasError: true,
		},
		{
			name:     "Invalid IP format",
			cidr:     "192.168.1.0/24",
			ip:       "invalid.ip",
			expected: false,
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := WhetherCidrContainsIp(tt.cidr, tt.ip)

			if tt.hasError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestGetAllIpsOfCidr(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		expected []string
		hasError bool
	}{
		{
			name:     "Small CIDR range /30",
			cidr:     "192.168.1.0/30",
			expected: []string{"192.168.1.1", "192.168.1.2"},
			hasError: false,
		},
		{
			name:     "Very small CIDR range /31",
			cidr:     "192.168.1.0/31",
			expected: []string{}, // /31 has exactly 2 IPs, so after excluding first and last, result is empty
			hasError: false,
		},
		{
			name:     "Single host /32",
			cidr:     "192.168.1.10/32",
			expected: []string{"192.168.1.10"},
			hasError: false,
		},
		{
			name: "Medium CIDR range /29",
			cidr: "192.168.1.0/29",
			expected: []string{
				"192.168.1.1", "192.168.1.2", "192.168.1.3",
				"192.168.1.4", "192.168.1.5", "192.168.1.6",
			},
			hasError: false,
		},
		{
			name:     "IPv6 small range",
			cidr:     "2001:db8::0/126",
			expected: []string{"2001:db8::1", "2001:db8::2"},
			hasError: false,
		},
		{
			name:     "Invalid CIDR format",
			cidr:     "invalid.cidr",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Empty CIDR",
			cidr:     "",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAllIpsOfCidr(tt.cidr)

			if tt.hasError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !tt.hasError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkWhetherCidrContainsIp(b *testing.B) {
	cidr := "192.168.1.0/24"
	ip := "192.168.1.10"

	for i := 0; i < b.N; i++ {
		_, _ = WhetherCidrContainsIp(cidr, ip)
	}
}

func BenchmarkGetAllIpsOfCidr(b *testing.B) {
	cidr := "192.168.1.0/29"

	for i := 0; i < b.N; i++ {
		_, _ = GetAllIpsOfCidr(cidr)
	}
}

// Example tests for documentation
func ExampleWhetherCidrContainsIp() {
	contains, err := WhetherCidrContainsIp("192.168.1.0/24", "192.168.1.10")
	if err != nil {
		panic(err)
	}
	fmt.Println(contains)
	// Output: true
}

func ExampleGetAllIpsOfCidr() {
	ips, err := GetAllIpsOfCidr("192.168.1.0/30")
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	// Output:
	// 192.168.1.1
	// 192.168.1.2
}
