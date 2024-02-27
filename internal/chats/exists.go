package chats

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func Exists(w http.ResponseWriter, r *http.Request) {
	chatService := r.Context().Value("chatService").(models.ChatService)
	firstUserID := uuid.FromStringOrNil(r.URL.Query().Get("first_user_id"))
	secondUserID := uuid.FromStringOrNil(r.URL.Query().Get("second_user_id"))

	chatID, err := chatService.Exists(firstUserID, secondUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{"exists": !chatID.IsNil(), "chat_id": chatID}

	json.Response(w, http.StatusOK, response)
}
