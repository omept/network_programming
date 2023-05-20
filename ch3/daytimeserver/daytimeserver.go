/* Daytime server -- on terminal, run 
$ telnet localhost 1200 
*/

package main

import (
	"net"
	"time"
	"log"
)

func main(){
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil{
		log.Fatalln("Fatal error: %s", err.Error())
	}
}