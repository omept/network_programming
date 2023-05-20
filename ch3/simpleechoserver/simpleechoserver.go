/* Simple echo server.  You need two terminals
to run. On the first terminal run :

$ telnet localhost 1201

on the second, run:

telnet localhost 1201
*/

package main 

import (
	"net"
	"log"
	"fmt"
)

func main(){
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close() // we're done 
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}

func checkError(err error){
	if err != nil {
		log.Fatalln("Fatal error: %s", err.Error())
	}
}