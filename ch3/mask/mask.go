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
		log.Fatalf("supported format: %s ip ones bits", os.Args[0])
	}

	ipString := os.Args[1]
	onesString := os.Args[2]
	bitsString := os.Args[3]

	ip := net.ParseIP(ipString)

	if ip == nil {
		log.Fatalln("invalid IP")
	}
	ones, err := strconv.Atoi(onesString)
	if err != nil {
		log.Fatalln("invalid ones passed")
	}

	bits, err := strconv.Atoi(bitsString)
	if err != nil {
		log.Fatalln("invalid bits passed")
	}

	mask := net.CIDRMask(ones, bits)

	network := ip.Mask(mask)
	mo, mb := mask.Size()
	fmt.Printf("IP Address: %s, \nNetwork : %s, \nmask (hex): %s,\nones : %d,\nbits: %d\n", ip.String(), network, mask.String(), mo, mb)
}
