package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Message struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ElementID uuid.UUID `json:"element_id" db:"element_id"`
	Content   string    `json:"content" db:"content"`
	SenderID  uuid.UUID `json:"sender_id" db:"sender_id"`
	Kind      string    `json:"kind" db:"kind"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type MessageInfo struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Content       string    `json:"content" db:"content"`
	UserFirstName string    `json:"user_first_name" db:"user_first_name"`
	UserLastName  string    `json:"user_last_name" db:"user_last_name"`
	UserEmail     string    `json:"user_email" db:"user_email"`
	UserPicture   string    `json:"user_picture" db:"user_picture"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	ElementID     uuid.UUID `json:"element_id" db:"element_id"`
	Kind          string    `json:"kind" db:"kind"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessagesSummary []MessageInfo

type MessageService interface {
	Create(message *Message) error
	GetByElementID(elementID uuid.UUID) (MessagesSummary, error)
}
