package client

import (
	"net"
	"log"
	"unicode/utf16"
	"fmt"
)

const BOM = '\ufffe'

func Utf16Client(server string) {

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	shorts := readShorts(conn)
	ints := utf16.Decode(shorts)
	str := string(ints)

	fmt.Println(str)

}

func readShorts(conn net.Conn) []uint16 {
	var buf [512] byte

	n, err := conn.Read(buf[0:2])

	for true {
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
	}

	if err != nil {
		log.Fatal(err)
	}

	var shorts []uint16

	shorts = make([]uint16, n/2)

	if buf[0] == 0xff && buf[1] == 0xfe {
		//big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
	} else if buf[1] == 0xfe && buf[0] == 0xff {
		//little endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		fmt.Println("Unknown order")
	}

	return shorts

}
