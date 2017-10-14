package server

import (
	"crypto/tls"
	"log"
	"time"
	"crypto/rand"
	"net"
	"fmt"
)

func TlsServer(){
	cert, err := tls.LoadX509KeyPair("certificate.pem", "private.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}


	now := time.Now()
	config.Time = func() time.Time {
		 return  now
	}

	config.Rand = rand.Reader


	listener, err := tls.Listen("tcp", "localhost:1900",&config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		fmt.Println("Accepted")
		go handleTlsConnect(conn)
	}
}


func handleTlsConnect(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

