package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"github.com/abondar24/GoBase/xmlcodec"
)

type Person struct {
	Name string
	Emails []string
}

func WebSocketServer() {
	http.Handle("/", websocket.Handler(echo))
	log.Fatal(http.ListenAndServe("localhost:2400",nil))
}

func WebSocketJsonServer(){
	http.Handle("/", websocket.Handler(ReceivePerson))
	log.Fatal(http.ListenAndServe("localhost:2500",nil))
}

func WebSocketXmlServer(){
	http.Handle("/", websocket.Handler(ReceivePersonXml))
	log.Fatal(http.ListenAndServe("localhost:2600",nil))
}

func echo(ws *websocket.Conn) {
	for n := 0; n < 10; n++ {
		msg := "Hello" + string(n+48)
		fmt.Println("Sending to client: " + msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			log.Fatal(err)
		}

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Recieved from client: ",reply)
	}

}

func ReceivePerson(ws *websocket.Conn){
	var person Person

	err:= websocket.JSON.Receive(ws,&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name: " + person.Name)
	for _, e := range person.Emails {
		fmt.Println("An email: " + e)
	}
}


func ReceivePersonXml(ws *websocket.Conn){
	var person Person

	err:= xmlcodec.XmlCodec.Receive(ws,&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name: " + person.Name)
	for _, e := range person.Emails {
		fmt.Println("An email: " + e)
	}
}

