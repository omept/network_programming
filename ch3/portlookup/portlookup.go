package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s network-type service\n", os.Args[0])
	}
	netType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(netType, service)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Printf("Port: %d", port)
}
