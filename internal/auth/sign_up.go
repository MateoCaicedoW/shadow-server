package auth

import (
	"net/http"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.Decode(r, &user); err != nil {
		json.Response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userService := r.Context().Value("userService").(models.UserService)
	verrs := validateSignUp(user, userService)
	if verrs.HasAny() {
		json.Response(w, http.StatusUnprocessableEntity, verrs.Errors)
		return
	}

	// Continue with the sign-up
	err := userService.CreateUser(&models.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	if err != nil {
		json.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Response(w, http.StatusCreated, map[string]string{"message": "User created successfully."})
}

func validateSignUp(user models.User, userService models.UserService) *validate.Errors {
	return validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: user.Email, Message: "Email is required."},
		&validators.FuncValidator{
			Fn: func() bool {
				user, _ := userService.GetUserByEmail(user.Email)

				return user.ID.IsNil()
			},
			Name:    "Email",
			Message: "%sUser already exists.",
		},
		&validators.StringIsPresent{Name: "FirstName", Field: user.FirstName, Message: "First name is required."},
		&validators.StringIsPresent{Name: "LastName", Field: user.LastName, Message: "Last name is required."},
	)
}
