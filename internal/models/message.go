package models

import "github.com/oklog/ulid"

type Message struct {
	Id        ulid.ULID
	ChannelId int
	SenderId  int
	Message   string
}

type Messages []Message
