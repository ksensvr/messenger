package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"messenger/internal/models"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *Repository {
	return &Repository{DB: db}
}

type DbUser struct {
	Id       int
	Uuid     string
	Name     string
	Password string
}

type DbUsers []DbUser

type UsersRepository interface {
	CreateUser(name, password string) (int, error)
	GetUser(id int) (models.User, error)
	DeleteUser(id int) error
}

type ChannelsRepository interface {
	CreateChannel(name string, creatorId int) error
	DeleteChannel(id int) error
	GetUserChannels(userId int) (models.Channels, error)
	GetChannelsByIds(channelIds []int) (models.Channels, error)
}

type ParticipantsRepository interface {
	AddParticipant(userId int, channelId int) error
	DeleteParticipant(userId int, channelId int) error
	GetParticipants(channelId string) (models.Participants, error) //или сразу вернуть models.Users
}

type MessagesRepository interface {
	CreateMessage(channelId int, senderId int, message string) error
	GetMessages(channelId int) (models.Messages, error)
}

type AttachmentsRepository interface {
	CreateAttachment(messageId string, fileUrl string) error
	GetAttachments(messageId string) (models.Attachment, error)
}
