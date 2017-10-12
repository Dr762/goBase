package client

import (
	"log"
	"net"
	"bufio"
	"os"
	"strings"
	"fmt"
	"bytes"
)

const (
	uiDir  = "dir"
	uiPwd  = "pwd"
	uiCd   = "cd"
	uiQuit = "quit"
)

const (
	DIR = "DIR"
	PWD = "PWD"
	CD  = "CD"
)

func FtpClient(server string) {
	conn, err := net.Dial("tcp", server+":1700")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		line = strings.TrimRight(line, "\t\r\n")
		if err != nil {
			log.Println(err)
			break
		}

		//cmd + arg
		strs := strings.SplitN(line, " ", 2)
		switch strs[0] {

		case uiDir:
			dirRequest(conn)

		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}

			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])

		case uiPwd:
			pwdRequest(conn)

		case uiQuit:
			conn.Close()
			os.Exit(0)

		default:
			fmt.Println("unknown command")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
	var buf [512]byte

	res := bytes.NewBuffer(nil)
	for {
		//read till blank line
		n, _ := conn.Read(buf[0:])
		res.Write(buf[0:n])

		length := res.Len()
		contents := res.Bytes()

		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0:length-4]))
			return
		}
	}

}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512] byte

	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512] byte

	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir\""+s+"\"")

}
