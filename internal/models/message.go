package models

type Message struct {
	Id        int
	Uuid      string
	ChannelId int
	SenderId  int
	Message   string
}

type Messages []Message
