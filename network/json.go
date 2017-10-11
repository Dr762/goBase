package network

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p* Person) PersonString() string{
	s:= p.Name.Personal + " " + p.Name.Family

	for _,v:=range p.Email{
		s+= "\n" + v.Kind + ": "+v.Address
	}

	return s
}

func JsonMarshall() {
	b,err := json.Marshal(PrepareData())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Marshalled:" ,string(b))

	var person Person
	err = json.Unmarshal(b,&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unmarshalled:" ,person)
}

func PrepareData() *Person {
	person := Person{
		Name: Name{Family: "Stark", Personal: "Rob"},
		Email: []Email{{Kind: "work", Address: "rob.stark@north.com"},
			{Kind: "work", Address: "rob.stark@north.com"},
		},
	}

	return &person
}
