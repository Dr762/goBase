package client

import (
	"net"
	"log"
	"encoding/json"
	"github.com/abondar24/GoBase/network"
	"fmt"
)

func JsonClient(server string) {
	conn, err := net.Dial("tcp",server)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n:=1;n<10;n++{
		encoder.Encode(network.PrepareData())

		var newPerson network.Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.PersonString())
	}


}
