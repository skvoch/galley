package model

type User struct {
	Hash       string `json:"hash"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
