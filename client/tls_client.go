package client

import (
	"log"
	"fmt"
	"os"
	"crypto/tls"
)

func TlsClient(server string){

	//without config fails to connect
	config := tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp",server, &config)
	if err != nil {
		log.Fatal(err)
	}


	for n:=1;n<10;n++{
		fmt.Println("Writing...")
		conn.Write([]byte("Hiii "+string(n+48)))

        var buf [512]byte

        n,err:=conn.Read(buf[0:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[0:n]))
	}

	os.Exit(0)
}
