package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Supported format: %s <domain>", os.Args[0])
	}
	sAddr := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp", sAddr)
	evalErr(err)

	listener, err := net.ListenTCP("tcp", addr)
	evalErr(err)

	for {

		acceptdCon, err := listener.Accept()
		if err != nil {
			continue
		}
		go processConnection(acceptdCon)

	}

}
func processConnection(con net.Conn) {
	con.Write([]byte(time.Now().String()))
	log.Default().Printf("Sent date time to %s", con.RemoteAddr().String())
	con.Close()
}

func evalErr(e error) {
	if e != nil {
		log.Fatalln("error: ", e.Error())
	}
}
