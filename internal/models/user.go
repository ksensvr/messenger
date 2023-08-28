package models

type User struct {
	Uuid string
	Name string
}

type DbUser struct {
	Id       int
	Uuid     string
	Name     string
	Password string
}

type Users []User

type DbUsers []DbUser
