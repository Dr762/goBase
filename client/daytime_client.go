package client

import (
	"net"
	"log"
	"fmt"
)

func DaytimeClient(service string){
	udpAddr,err := net.ResolveUDPAddr("udp4",service)
	if err != nil {
		log.Fatal(err)
	}

	conn,err := net.DialUDP("udp",nil,udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	_,err = conn.Write([]byte("Client is up"))
	if err != nil {
		log.Fatal(err)
	}

	var buf [512]byte

	n,err := conn.Read(buf[0:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf[0:n]))
}
