/*
IP Mask
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("usage: %s dotted-ip-address ones bits \n", os.Args[0])
	}
	dotIp := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])
	addrs := net.ParseIP(dotIp)

	if bits != 64 && bits != 32 {
		log.Fatal("invalid bits value")
	}

	if addrs == nil {
		log.Fatal("invalid address")
	}

	//The simplest function to create a netmask uses the CIDR notation of ones followed by zeroes up to the number of bits:
	mask := net.CIDRMask(ones, bits) // Classless Inter-Domain Routing (a.k.a. CIDR)
	computedOnes, computedBits := mask.Size()
	network := addrs.Mask(mask)

	fmt.Println("Address is ", addrs.String(),
		"\nMask Length is : ", computedBits,
		"\nLeading ones count is : ", computedOnes,
		"\nMask is (hex): ", mask.String(),
		"\nNetwork is : ", network.String())
}
