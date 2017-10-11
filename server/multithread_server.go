package server

import (
	"net"
	"log"
	"fmt"
)

func MultithreadServer(){
	listener, err := net.Listen("tcp", "localhost:1201")
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleClient(conn)

	}
}

func handleClient(conn net.Conn){
	defer conn.Close()

	var buf [512]byte

	for {
		n,err := conn.Read(buf[0:])
		if err != nil {
			log.Print(err)
		}

		fmt.Println(string(buf[0:]))

		_,err = conn.Write(buf[0:n])
		if err != nil {
			log.Print(err)
		}
	}
}
