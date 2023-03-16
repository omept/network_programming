package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s dotted-ip-address", os.Args[0])
	}

	ip := os.Args[1]
	addrs := net.ParseIP(ip)
	if addrs == nil {
		fmt.Println("invalid address")
	}

	mask := addrs.DefaultMask()
	network := addrs.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address is: ", addrs.String(),
		"\nDefault Mask Length is : ", bits,
		"\nLeading ones count is : ", ones,
		"\nMask is (hex): ", mask.String(),
		"\nNetwork is: ", network.String())
}
