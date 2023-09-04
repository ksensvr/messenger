package models

import "github.com/oklog/ulid"

type Channel struct {
	Id        ulid.ULID
	Name      string
	CreatorId int
}

type Channels []Channel
