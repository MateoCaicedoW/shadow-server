package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Group struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Groups []Group

type GroupService interface {
	Create(group *Group) error
	AddMember(groupID, userID uuid.UUID) error
}
