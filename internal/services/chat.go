package services

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"github.com/shadow/backend/internal/models"
)

var _ models.ChatService = (*chat)(nil)

type chat struct {
	db *sqlx.DB
}

func Chats(db *sqlx.DB) *chat {
	return &chat{db: db}
}

func (c *chat) Create(chat *models.Chat) error {
	query := `INSERT INTO chats (first_user_id, second_user_id) VALUES ($1, $2) RETURNING *`

	err := c.db.Get(chat, query, chat.FirstUserID, chat.SecondUserID)
	if err != nil {
		return err
	}

	return nil
}

func (c *chat) Chats(userID uuid.UUID) (models.ChatSummaries, error) {
	chats := models.ChatSummaries{}
	query := `
		SELECT 
		chats.id,
		chats.first_user_id,
		chats.second_user_id,
		fuser.first_name AS first_user_name,
		fuser.picture AS first_user_picture,
		suser.first_name AS second_user_name,
		suser.picture AS second_user_picture

	FROM 
		chats
	JOIN 
		users fuser ON (fuser.id = chats.first_user_id)
	JOIN
		users suser ON (suser.id = chats.second_user_id) 
	WHERE 
		first_user_id = $1 OR second_user_id = $1
	ORDER BY chats.created_at DESC;
	`

	err := c.db.Select(&chats, query, userID)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (c *chat) Exists(firstUserID, secondUserID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM chats WHERE (first_user_id = $1 AND second_user_id = $2) OR (first_user_id = $2 AND second_user_id = $1))`
	var exists bool
	err := c.db.Get(&exists, query, firstUserID, secondUserID)
	if err != nil {
		return false, err
	}

	return exists, nil
}
