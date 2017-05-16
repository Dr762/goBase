package db

type Job struct {
	Id       int    `json:"id"`
	JobName  string `json:"job_name"`
	PersonId int    `json:"personId"`
}
