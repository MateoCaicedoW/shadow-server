package chats

import (
	"net/http"

	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	chatService := r.Context().Value("chatService").(models.ChatService)
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

	json.Response(w, http.StatusCreated, chat)
}
