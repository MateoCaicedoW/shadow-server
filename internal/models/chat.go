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

type ChatService interface {
	Create(chat *Chat) error
}
