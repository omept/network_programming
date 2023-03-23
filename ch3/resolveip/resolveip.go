package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s domain-name", os.Args[1])
	}

	ipAddr, err := net.ResolveIPAddr("ip", os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP Address %s", ipAddr)
}
