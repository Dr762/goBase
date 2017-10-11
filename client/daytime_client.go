package client

import (
	"net"
	"log"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/asn1"
)

func DaytimeUdpClient(server string){
	udpAddr,err := net.ResolveUDPAddr("udp4", server)
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


func DaytimeAsn1Client(server string) {
	conn,err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	resp,err:=ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	var newtime time.Time
	_,err = asn1.Unmarshal(resp,&newtime)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After marshal/unmarshal: ", newtime.String())
}
