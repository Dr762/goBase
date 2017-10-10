package client

import (
	"net"
	"fmt"
	"io/ioutil"
	"log"
)

func TcpClient(service string) {
	conn,err := net.Dial("tcp",service)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	resp,err:=ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp))
}
