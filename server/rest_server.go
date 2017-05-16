package server

import (
	"encoding/json"
	"github.com/GoBase/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var base = db.Connect()

func RunRestServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/getPerson/{person_id}", getPerson).Methods("GET")
	router.HandleFunc("/getPersons", getPersons).Methods("GET")
	router.HandleFunc("/insertPerson", insertPerson).Methods("POST")
	router.HandleFunc("/getJob/{job_id}", getJob).Methods("GET")
	router.HandleFunc("/getJobForPerson/{person_id}", getJobForPerson).Methods("GET")
	router.HandleFunc("/insertJob", insertJob).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hiiii")
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id := vars["person_id"]

	personId, err := strconv.Atoi(person_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	person, err := db.GetPerson(base, personId)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(person)
	}

}

func getPersons(w http.ResponseWriter, r *http.Request) {

	persons, err := db.GetPersons(base)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(persons)
	}

}

func insertPerson(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	person_id := r.FormValue("person_id")
	personId, err := strconv.Atoi(person_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	adm := r.FormValue("admin")
	admin, err := strconv.ParseBool(adm)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	person := &db.Person{Name: name, PersonId: personId, Admin: admin}

	id, err := db.InsertPerson(base, person)
	person.Id = id
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(person)
	}
}

func getJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	job_id := vars["job_id"]

	jobId, err := strconv.Atoi(job_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	job, err := db.GetJob(base, jobId)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(job)
	}
}

func getJobForPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id := vars["person_id"]

	personId, err := strconv.Atoi(person_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	job, err := db.GetJobForPerson(base, personId)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(job)
	}
}

func insertJob(w http.ResponseWriter, r *http.Request) {
	jobName := r.FormValue("job_name")

	person_id := r.FormValue("person_id")
	personId, err := strconv.Atoi(person_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	job := &db.Job{JobName: jobName, PersonId: personId}

	id, err := db.InsertJob(base, job)
	job.Id = id
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(job)
	}
}
