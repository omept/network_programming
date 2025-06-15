package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/omept/go_network_programming/ipconv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ip-addr\n", os.Args[0])
	}

	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	fmt.Printf("IPV4 version : %s \n", ipconv.IpTo4(name))
	fmt.Printf("IPV6 version : %s \n", ipconv.IpTo6(name))
}
