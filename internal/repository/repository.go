package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"messenger/internal/infr/database"
	"messenger/internal/models"
)

type Repository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(cfg database.PostgresConfig) (Repository, error) {
	db, err := sqlx.Open("postgres", GetPostgresDsn(cfg))
	repo := Repository{
		DB: db,
	}
	if err != nil {
		return repo, err
	}

	return repo, nil
}

func GetPostgresDsn(cfg database.PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Server,
		cfg.User,
		cfg.Password,
		cfg.DatabaseName,
		cfg.Port,
	)
}

type usersRepository interface {
	CreateUser(name, password string) (int, error)
	GetUser(id int) (models.User, error)
	DeleteUser(id int) error
}

type channelsRepository interface {
	CreateChannel(name string, creatorId int) error
	DeleteChannel(id int) error
	GetUserChannels(userId int) (models.Channels, error)
	GetChannelsByIds(channelIds []int) (models.Channels, error)
}

type participantsRepository interface {
	AddParticipant(userId int, channelId int) error
	DeleteParticipant(userId int, channelId int) error
	GetParticipants(channelId string) (models.Participants, error) //или сразу вернуть models.Users
}

type messagesRepository interface {
	CreateMessage(channelId int, senderId int, message string) error
	GetMessages(channelId int) (models.Messages, error)
}

type attachmentsRepository interface {
	CreateAttachment(messageId string, fileUrl string) error
	GetAttachments(messageId string) (models.Attachment, error)
}
