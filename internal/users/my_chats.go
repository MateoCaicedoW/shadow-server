package users

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func MyChats(w http.ResponseWriter, r *http.Request) {
	userID := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	chatService := r.Context().Value("chatService").(models.ChatService)

	chats, err := chatService.Chats(userID)
	if err != nil {
		fmt.Println("error getting chats: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Response(w, http.StatusOK, chats)
}
