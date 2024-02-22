package services

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shadow/backend/internal/models"
)

var _ models.UserService = (*user)(nil)

type user struct {
	db *sqlx.DB
}

func Users(db *sqlx.DB) *user {
	return &user{db: db}
}

func (u *user) CreateUser(user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email, picture) VALUES ($1, $2, $3, $4) RETURNING *`

	err := u.db.Get(user, query, user.FirstName, user.LastName, user.Email, user.Picture)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (u *user) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE email = $1`

	err := u.db.Get(&user, query, email)
	if err != nil {
		return user, fmt.Errorf("could not find user with email %s", email)
	}

	return user, nil
}
