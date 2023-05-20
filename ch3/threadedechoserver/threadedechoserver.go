package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//run as a goroutine
		go handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}

func handleClient(conn net.Conn){
	// close connection on exit
	defer conn.Close()
	var buf [512]byte

	for{
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		//write the n bytes read
		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}