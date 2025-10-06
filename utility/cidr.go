// Package utility provides network utility functions for CIDR operations.
package utility

import "net"

// WhetherCidrContainsIp checks if the given IP address is within the specified CIDR range.
//
// Parameters:
//   - cidr: A string representing the CIDR notation (e.g., "192.168.1.0/24")
//   - ip: A string representing the IP address to check (e.g., "192.168.1.10")
//
// Returns:
//   - bool: true if the IP is within the CIDR range, false otherwise
//   - error: any error that occurred during parsing
//
// Example:
//
//	contains, err := WhetherCidrContainsIp("192.168.1.0/24", "192.168.1.10")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(contains) // true
func WhetherCidrContainsIp(cidr, ip string) (bool, error) {
	if _, subnet, err := net.ParseCIDR(cidr); err != nil {
		return false, err
	} else {
		return subnet.Contains(net.ParseIP(ip)), nil
	}
}

// GetAllIpsOfCidr returns all IP addresses within the specified CIDR range.
// This function excludes the network and broadcast addresses.
//
// Parameters:
//   - cidr: A string representing the CIDR notation (e.g., "192.168.1.0/30")
//
// Returns:
//   - []string: a slice of IP addresses as strings within the CIDR range
//   - error: any error that occurred during parsing
//
// Example:
//
//	ips, err := GetAllIpsOfCidr("192.168.1.0/30")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(ips) // ["192.168.1.1", "192.168.1.2"]
//
// Note:
//   - For CIDR ranges with less than 2 usable IPs, all IPs are returned
//   - Network and broadcast addresses are excluded for ranges with 2 or more IPs
func GetAllIpsOfCidr(cidr string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); {
		ips = append(ips, ip.String())

		for i := len(ip) - 1; i >= 0; i-- {
			ip[i]++
			if ip[i] > 0 {
				break
			}
		}
	}

	switch {
	case len(ips) < 2:
		return ips, nil
	default:
		return ips[1 : len(ips)-1], nil
	}
}
