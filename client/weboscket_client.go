package client

import (
	"golang.org/x/net/websocket"
	"log"
	"io"
	"fmt"
	"os"
	"github.com/abondar24/GoBase/xmlcodec"
)

type Person struct {
	Name   string
	Emails []string
}

func WebSocketClient(server string, msg string) {

	conn, err := websocket.Dial("ws://"+server, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	var incomingMsg string

	for {
		err = websocket.Message.Receive(conn, &incomingMsg)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}
		fmt.Println("Received from server: " + incomingMsg)

		err = websocket.Message.Send(conn, msg)
		if err != nil {
			log.Fatal(err)
		}

	}

	os.Exit(0)
}

func WebSocketJsonClient(server string) {
	conn, err := websocket.Dial("ws://"+server, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	person := Person{Name: "Alex",
		Emails: []string{"alex@mail.ru", "desertalex@icloud.com"},
	}

	fmt.Println("Sending ",person)

	err = websocket.JSON.Send(conn,person)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}

func WebSocketXmlClient(server string) {
	conn, err := websocket.Dial("ws://"+server, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	person := Person{Name: "Alex",
		Emails: []string{"alex@mail.ru", "desertalex@icloud.com"},
	}

	fmt.Println("Sending ",person)

	err = xmlcodec.XmlCodec.Send(conn,person)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}