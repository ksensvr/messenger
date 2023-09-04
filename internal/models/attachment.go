package models

import "github.com/oklog/ulid"

type Attachment struct {
	Id        ulid.ULID
	MessageId int
	FileUrl   string
}

type Attachments []Attachment
