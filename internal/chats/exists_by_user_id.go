package chats

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
	"github.com/shadow/backend/internal/services"
)

func ExistsByUserID(w http.ResponseWriter, r *http.Request) {
	chatID := uuid.FromStringOrNil(chi.URLParam(r, "id"))
	token := r.Header.Get("Authorization")
	token = token[len("Bearer "):]

	jwtService := r.Context().Value("jwtService").(services.JWTService)
	user, err := jwtService.ValidateToken(token)
	if err != nil {
		json.Response(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token."})
		return
	}

	userID := user["id"].(string)
	chatService := r.Context().Value("chatService").(models.ChatService)

	exists, err := chatService.ExistsByUserID(uuid.FromStringOrNil(userID), chatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Response(w, http.StatusOK, exists)
}
