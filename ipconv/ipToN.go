package ipconv

import "net"

func IpTo4(s string) net.IP {
	ip := net.ParseIP(s)
	return ip.To4()
}

func IpTo6(s string) net.IP {
	ip := net.ParseIP(s)
	return ip.To16()
}
