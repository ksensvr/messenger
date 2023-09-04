package models

import "github.com/oklog/ulid"

type User struct {
	Id   ulid.ULID
	Name string
}

type Users []User
