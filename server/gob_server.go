package server

import (
	"net"
	"log"
	"github.com/abondar24/GoBase/network"
	"fmt"
	"encoding/gob"
)

func GobServer(){
	listener, err := net.Listen("tcp", "localhost:1600")
	if err != nil {
		log.Fatal(err)
	}

	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Print(err)
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n:=1;n<10;n++{
			var person network.Person
			decoder.Decode(&person)
			fmt.Println(person.PersonString())
			encoder.Encode(network.PrepareData())
		}

		conn.Close()
	}
}
