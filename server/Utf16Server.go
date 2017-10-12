package server

import (
	"net"
	"log"
	"unicode/utf16"
)

const BOM = '\ufffe'

func Utf16Server() {
	listener, err := net.Listen("tcp", "localhost:1800")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		str := "Всем привет!"
		shorts := utf16.Encode([]rune(str))

		writeShorts(conn, shorts)
		conn.Close()
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte

	// send bom as first 2 bytes
	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255

	_, err := conn.Write(bytes[0:])
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)

		_, err := conn.Write(bytes[0:])
		if err != nil {
			log.Println(err)
			return
		}
	}

}
