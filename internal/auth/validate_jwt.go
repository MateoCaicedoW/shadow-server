package auth

import (
	"fmt"
	"net/http"

	"github.com/shadow/backend/internal/response"
	"github.com/shadow/backend/internal/services"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	auth2 := r.URL.Query().Get("Authorization")

	if auth == "" && auth2 == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := auth
	if auth2 != "" {
		token = auth2
	}

	jwtService := r.Context().Value("jwtService").(services.JWTService)
	token = token[len("Bearer "):]
	user, err := jwtService.ValidateToken(token)

	if err != nil {
		fmt.Println("error", err)
		response.JSON(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token."})
		return
	}

	response.JSON(w, http.StatusOK, user)
}
