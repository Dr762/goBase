package db

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	PersonId int    `json:"personId"`
	Admin    bool   `json:"admin"`
}
