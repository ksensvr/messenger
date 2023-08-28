package models

type Attachment struct {
	Id        int
	Uuid      string
	MessageId int
	FileUrl   string
}

type Attachments []Attachment
