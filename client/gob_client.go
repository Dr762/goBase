package client

import (
	"net"
	"log"
	"github.com/abondar24/GoBase/network"
	"fmt"
	"encoding/gob"
)

func GobClient(server string){
	conn, err := net.Dial("tcp",server)
	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n:=1;n<10;n++{
		encoder.Encode(network.PrepareData())

		var newPerson network.Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.PersonString())
	}


}
