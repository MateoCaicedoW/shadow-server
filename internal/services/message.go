package services

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"github.com/shadow/backend/internal/models"
)

var _ models.MessageService = (*message)(nil)

type message struct {
	db *sqlx.DB
}

func Messages(db *sqlx.DB) *user {
	return &user{db: db}
}

func (m *message) Create(message *models.Message) error {
	query := `INSERT INTO messages (element_id, content, sender_id, kind) VALUES ($1, $2, $3, $4) RETURNING *`

	err := m.db.Get(message, query, message.ElementID, message.Content, message.SenderID, message.Kind)
	if err != nil {
		return fmt.Errorf("could not create message: %w", err)
	}

	return nil
}

func (m *message) GetByElementID(elementID uuid.UUID) (models.MessagesSummary, error) {
	messages := models.MessagesSummary{}
	query := `
	SELECT 
		m.id, 
		m.content, 
		u.first_name as user_first_name, 
		u.last_name as user_last_name, 
		u.email as user_email, 
		u.picture as user_picture, 
		m.sender_id as user_id, 
		m.kind, 
		m.created_at, 
		m.updated_at 
	FROM messages m 
	JOIN users u ON m.sender_id = u.id 
	WHERE m.element_id = $1 
	ORDER BY m.created_at ASC`

	err := m.db.Select(&messages, query, elementID)
	if err != nil {
		return messages, fmt.Errorf("could not find messages with element_id %s", elementID)
	}

	return messages, nil
}
