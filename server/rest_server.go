package server

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/GoBase/model"
	"encoding/json"
	"strconv"
)

func RunRestServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/getPerson/{person_id}", getPerson)
	router.HandleFunc("/insertPerson",insertPerson)
	router.HandleFunc("/getJob/{job_id}",getJob)
	router.HandleFunc("/getJobForPerson/{job_id,person_id}",getJobForPerson)
	router.HandleFunc("/insertJob",insertJob)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hiiii")
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id := vars["person_id"]

	personId,_ := strconv.Atoi(person_id)


	person := model.Person{
		Id: 1,
		Name: "Alex",
		PersonId: personId,
		Admin:true,
	}

	json.NewEncoder(w).Encode(person)
}

func insertPerson(w http.ResponseWriter, r *http.Request)  {
//	json.NewEncoder(w).Encode()
}

func getJob(w http.ResponseWriter, r *http.Request)  {
	
}

func getJobForPerson(w http.ResponseWriter, r *http.Request)  {
	
}

func insertJob(w http.ResponseWriter, r *http.Request) {

}

