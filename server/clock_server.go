package server

import (
	"io"
	"log"
	"net"
	"time"
)

func ClockServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(con)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("23:59:35\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
