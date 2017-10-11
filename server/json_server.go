package server

import (
	"log"
	"net"
	"encoding/json"
	"github.com/abondar24/GoBase/network"
	"fmt"
)

func JsonServer(){

	listener, err := net.Listen("tcp", "localhost:1500")
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Print(err)
		}

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		for n:=1;n<10;n++{
			var person network.Person
			decoder.Decode(&person)
			fmt.Println(person.PersonString())
			encoder.Encode(network.PrepareData())
		}

		conn.Close()
	}
}
