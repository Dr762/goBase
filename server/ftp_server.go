package server

import (
	"net"
	"log"
	"os"
)

const (
	DIR = "DIR"
	PWD = "PWD"
	CD  = "CD"
)

func FtpServer() {
	listener, err := net.Listen("tcp", "localhost:1700")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleClt(conn)
	}
}

func handleClt(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Print(err)
			conn.Close()
			return
		}

		s := string(buf[0:n])

		if s[0:2] == CD {
			changeDir(conn, s[3:])
		} else if s[0:3] == DIR {
			listDir(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func changeDir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func listDir(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir,err := os.Open(".")
	if err != nil {
		conn.Write([]byte(""))
		return
	}

	names,err := dir.Readdirnames(-1)
	if err != nil {
		conn.Write([]byte(""))
		return
	}

	for _,nm :=range names{
		conn.Write([]byte(nm+"\r\n"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))

}
