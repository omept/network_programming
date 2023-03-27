package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s address", os.Args[1])
	}

	addr := os.Args[1]
	mxs, err := net.LookupMX(addr)

	if err != nil {
		log.Fatal("Lookup error ", err.Error())
	}
	for v, mx := range mxs {
		fmt.Printf("%d) %s \n", v+1, mx.Host)
	}
}
