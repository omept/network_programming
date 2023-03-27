package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Lookup the host IPs from a canonical name
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s hostname", os.Args[0])
	}
	host := os.Args[1]
	cname, err := net.LookupCNAME(host)
	checkErr(err)

	hosts, err := net.LookupHost(cname)
	checkErr(err)

	for _, v := range hosts {
		fmt.Println(v)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Error :", err.Error())
	}
}
