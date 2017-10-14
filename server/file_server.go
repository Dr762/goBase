package server

import (
	"net/http"
	"log"
)

func FileServer(){
	fileServer := http.FileServer(http.Dir("/home/abondar/"))

	err:= http.ListenAndServe("localhost:2000",fileServer)
	if err != nil {
		log.Fatal(err)
	}


}