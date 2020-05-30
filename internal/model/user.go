package model

type User struct {
	Hash       string `json:"hash" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
}
