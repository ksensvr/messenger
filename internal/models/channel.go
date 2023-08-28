package models

type Channel struct {
	Id        int
	Uuid      string
	Name      string
	CreatorId int
}

type Channels []Channel
