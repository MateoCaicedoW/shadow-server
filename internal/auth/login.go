package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/shadow/backend/internal/models"
	"github.com/shadow/backend/internal/response"
	"github.com/shadow/backend/internal/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	userService := r.Context().Value("userService").(models.UserService)
	jwtService := r.Context().Value("jwtService").(services.JWTService)

	verrs := validateLogin(user.Email, userService)
	if verrs.HasAny() {
		response.JSON(w, http.StatusUnprocessableEntity, verrs.Errors)
		return
	}

	user, err := userService.GetUserByEmail(user.Email)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := jwtService.GenerateToken(user)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}

func validateLogin(email string, userService models.UserService) *validate.Errors {
	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: email, Message: "Email is required."},
		&validators.FuncValidator{
			Fn: func() bool {
				if email == "" {
					return true
				}

				_, err := userService.GetUserByEmail(email)
				return err == nil
			},
			Name:    "Email",
			Message: "%sUser not found with this email address.",
		},
	)
}
