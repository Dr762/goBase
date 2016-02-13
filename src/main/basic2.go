package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int `json:"Released"`
	Color  bool `json:"omitempty"`
	Actors []string
}

func main() {
	jsonmarshaller()
}


func jsonmarshaller() {


	var movies = []Movie{
		{Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey Bogart","Ingrid Bergman"}},
		{Title:"Star Wars  Episdoe 4",Year:1977,Color:true,Actors:[]string{"Mark Hummil","Harrison Ford"}},
		{Title:"Die hard",Year:1984,Color:true,Actors:[]string{"Bruce Willis","Alan Rickman"}},

	}

    data,err := json.MarshalIndent(movies,"","   ")
	if err!=nil{
		log.Fatalf("JSON marshalling failed: %s",err)
	}

	fmt.Printf("%s\n",data)
}