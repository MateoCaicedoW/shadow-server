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
		CONCAT(fuser.first_name, ' ', fuser.last_name) AS first_user_name,
		fuser.picture AS first_user_picture,
		CONCAT(suser.first_name, ' ', suser.last_name) AS second_user_name,
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

func (c *chat) Exists(firstUserID, secondUserID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT (SELECT id FROM chats WHERE (first_user_id = $1 AND second_user_id = $2) OR (first_user_id = $2 AND second_user_id = $1))`
	var id uuid.NullUUID
	err := c.db.Get(&id, query, firstUserID, secondUserID)
	if err != nil {
		return id.UUID, err
	}

	return id.UUID, nil
}

func (c *chat) Messages(chatID uuid.UUID) (models.MessagesSummary, error) {
	messages := models.MessagesSummary{}
	query := `
		SELECT 
		messages.id,
		messages.sender_id as user_id,
		messages.content,
		users.first_name AS user_first_name,
		users.last_name AS user_last_name,
		users.email AS user_email,
		users.picture AS user_picture,
		messages.element_id,
		messages.kind
	FROM 
		messages
	JOIN 
		users ON (users.id = messages.sender_id)
	JOIN 
		chats ON (chats.id = messages.element_id)
	WHERE 
		messages.element_id = $1
	ORDER BY messages.created_at ASC;
	`

	err := c.db.Select(&messages, query, chatID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *chat) ExistsByUserID(userID, chatID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM chats WHERE id = $1 AND (first_user_id = $2 OR second_user_id = $2))`
	var exists bool
	err := c.db.Get(&exists, query, chatID, userID)
	if err != nil {
		return false, err
	}

	return exists, nil
}
