package server

import (
	"net"
	"log"
	"time"
)

func RunDaytimeTcpServer(){
	listener, err := net.Listen("tcp", "localhost:1200")
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Print(err)
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}


func RunDaytimeUdpServer()  {
	udp4Addr,err := net.ResolveUDPAddr("udp4","localhost:1300")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udp4Addr)
	if err != nil {
		log.Fatal(err)
	}

	for{
		handleUdpClient(conn)

	}
}

func handleUdpClient(conn *net.UDPConn)  {
     var buf [512]byte

     _,addr,err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		log.Print(err)
	}

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime),addr)
}
