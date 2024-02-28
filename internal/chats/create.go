package chats

import (
	"net/http"
	"time"

	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
	"github.com/shadow/backend/internal/websocket"
)

func Create(w http.ResponseWriter, r *http.Request) {
	chatService := r.Context().Value("chatService").(models.ChatService)
	users := r.Context().Value("userService").(models.UserService)
	var chat models.Chat

	if err := json.Decode(r, &chat); err != nil {
		json.Response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := chatService.Create(&chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	firstUser, err := users.GetByID(chat.FirstUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	secondUser, err := users.GetByID(chat.SecondUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatSummary := models.ChatSummary{
		ID:                chat.ID,
		FirstUserID:       chat.FirstUserID,
		SecondUserID:      chat.SecondUserID,
		FirstUserName:     firstUser.FullName(),
		SecondUserName:    secondUser.FullName(),
		FirstUserPicture:  firstUser.Picture,
		SecondUserPicture: secondUser.Picture,
		LastMessage:       "",
		LastMessageAt:     time.Time{},
	}

	chatSummaryMap := map[string]interface{}{
		"element_id": "chats",
		"action":     "create",
		"chat":       chatSummary,
	}

	json.Response(w, http.StatusCreated, chatSummaryMap)

	messageByte, err := json.Marshal(chatSummaryMap)
	if err != nil {
		json.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	websocket.Broadcast(messageByte)
}
