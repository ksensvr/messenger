package models

import "github.com/oklog/ulid"

type User struct {
	Uuid ulid.ULID
	Name string
}

type Users []User
