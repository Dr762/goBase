package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func EchoServer() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(con)
	}
}

func echoReq(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConnection(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echoReq(c, input.Text(), 1*time.Second)
	}
	input.Err()
	c.Close()
}
