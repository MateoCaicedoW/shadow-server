package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Chat struct {
	ID           uuid.UUID `json:"id" db:"id"`
	FirstUserID  uuid.UUID `json:"first_user_id" db:"first_user_id"`
	SecondUserID uuid.UUID `json:"second_user_id" db:"second_user_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Chats []Chat

type ChatSummary struct {
	ID                uuid.UUID `json:"id" db:"id"`
	FirstUserID       uuid.UUID `json:"first_user_id" db:"first_user_id"`
	FirstUserName     string    `json:"first_user_name" db:"first_user_name"`
	FirstUserPicture  string    `json:"first_user_picture" db:"first_user_picture"`
	SecondUserID      uuid.UUID `json:"second_user_id" db:"second_user_id"`
	SecondUserName    string    `json:"second_user_name" db:"second_user_name"`
	SecondUserPicture string    `json:"second_user_picture" db:"second_user_picture"`
	LastMessage       string    `json:"last_message" db:"last_message"`
	LastMessageAt     time.Time `json:"last_message_at" db:"last_message_at"`
}

type ChatSummaries []ChatSummary

type ChatService interface {
	Create(chat *Chat) error
	Chats(userID uuid.UUID) (ChatSummaries, error)
	Exists(firstUserID, secondUserID uuid.UUID) (bool, error)
}
