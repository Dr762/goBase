package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:alex21@tcp(172.17.0.2:3306)/test1?charset=utf8")

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func GetPerson(db *sql.DB, personId int) (*Person, error) {

	rows, err := db.Query("SELECT * FROM person where person_id=?", personId)
	if err != nil {
		return &Person{}, err
	}

	person := Person{}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&person.Name, &person.PersonId, &person.Admin, &person.Id)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &person, err
}

func GetPersons(db *sql.DB) (*[]Person, error) {

	rows, err := db.Query("SELECT * FROM person ")
	if err != nil {
		return &[]Person{}, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return &[]Person{}, err
	}

	persons := make([]Person, len(cols))

	defer rows.Close()

	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.Name, &person.PersonId, &person.Admin, &person.Id)
		if err != nil {
			log.Fatalln(err)
		}
		persons = append(persons, person)
	}

	return &persons, err
}

func InsertPerson(db *sql.DB, person *Person) (int, error) {

	res, err := db.Exec("INSERT INTO person(name,person_id,admin) values(?,?,?)",
		&person.Name, &person.PersonId, &person.Admin)

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return int(id), err
}

func GetJob(db *sql.DB, jobId int) (*Job, error) {

	rows, err := db.Query("SELECT * FROM job where id=?", jobId)
	if err != nil {
		return &Job{}, err
	}

	job := Job{}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&job.Id, &job.JobName, &job.PersonId)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &job, err
}

func GetJobForPerson(db *sql.DB, personId int) (*[]Job, error) {
	rows, err := db.Query("SELECT * FROM job j join person p on p.id=j.person_id where p.id=?", personId)
	if err != nil {
		return &[]Job{}, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return &[]Job{}, err
	}

	res := make([]Job, len(cols))

	defer rows.Close()

	for rows.Next() {
		job := Job{}
		person := Person{}
		err := rows.Scan(&job.Id, &job.JobName, &job.PersonId,
			&person.Name, &person.PersonId, &person.Admin, &person.Id)
		if err != nil {
			log.Fatalln(err)
		}

		res = append(res, job)
	}

	return &res, err
}

func InsertJob(db *sql.DB, job *Job) (int, error) {

	res, err := db.Exec("INSERT INTO job(job_name,person_id) values(?,?)", &job.JobName, &job.PersonId)

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return int(id), err
}
