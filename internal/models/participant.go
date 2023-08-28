package models

type Participant struct {
	Id        int
	Uuid      string
	ChannelId int
	UserId    int
	CanWrite  bool
}

type Participants []Participant
