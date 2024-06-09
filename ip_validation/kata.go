package kata

import (
	"net"
)

// Is_valid_ip / IP Validation
// https://www.codewars.com/kata/515decfd9dcfc23bb6000006/train/go
func Is_valid_ip(ip string) bool {
	parsedIp := net.ParseIP(ip)
	return parsedIp.String() == ip
}
