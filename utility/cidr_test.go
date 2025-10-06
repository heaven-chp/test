package utility

import (
	"fmt"
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
