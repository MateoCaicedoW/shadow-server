package chats

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func Messages(w http.ResponseWriter, r *http.Request) {
	chatID := chi.URLParam(r, "id")
	chatService := r.Context().Value("chatService").(models.ChatService)
	messages, err := chatService.Messages(uuid.FromStringOrNil(chatID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Response(w, http.StatusOK, messages)
}
