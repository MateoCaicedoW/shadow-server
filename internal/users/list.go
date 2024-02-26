package users

import (
	"net/http"

	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)
	users, err := userService.List()
	if err != nil {
		json.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Response(w, http.StatusOK, users)
}
