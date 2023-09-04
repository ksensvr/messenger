package models

import "github.com/oklog/ulid"

type Participant struct {
	Id        ulid.ULID
	ChannelId int
	UserId    int
	CanWrite  bool
}

type Participants []Participant
