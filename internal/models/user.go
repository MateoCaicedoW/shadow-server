package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Picture   string    `db:"picture" json:"picture"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Users []User

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// UserService is the interface for the user service
type UserService interface {
	// Create a new user
	CreateUser(user *User) error
	// Get a user by their email
	GetUserByEmail(email string) (User, error)
	// List all users
	List() (Users, error)
}
