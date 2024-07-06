package kata

import "fmt"

// IpsBetween / https://www.codewars.com/kata/526989a41034285187000de4/train/go
func IpsBetween(start, end string) int {
	ipStart, _ := parseIP(start)
	ipEnd, _ := parseIP(end)

	return ipStart.between(*ipEnd)
}

type ip struct {
	a, b, c, d int
}

func parseIP(s string) (*ip, error) {
	var a, b, c, d int

	num, err := fmt.Sscanf(s, "%d.%d.%d.%d", &a, &b, &c, &d)

	if num != 4 || err != nil {
		return nil, fmt.Errorf("invalid ip address: %s", s)
	}

	return &ip{a, b, c, d}, nil
}

func (ip ip) between(other ip) int {
	return (other.a-ip.a)*256*256*256 +
		(other.b-ip.b)*256*256 +
		(other.c-ip.c)*256 +
		(other.d - ip.d)
}
