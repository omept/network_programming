package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"slices"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("required format: %s <domain> <type>", os.Args[0])
	}

	allowed := []string{"ip", "ip4", "ip6"}
	domain := os.Args[1]
	ntype := os.Args[2]

	if !slices.Contains(allowed, ntype) {
		log.Fatalln("invalid type. allowed:", allowed)
	}

	ip, err := net.ResolveIPAddr(ntype, domain)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Domain : %s,\nIP : %s,\nZone : %s\n", domain, ip.IP, ip.Zone)
}
