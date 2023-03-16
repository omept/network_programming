/*
IP
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage %s ip-address\n", os.Args[0])
	}
	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}

}
