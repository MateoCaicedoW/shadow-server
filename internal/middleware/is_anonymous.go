package middleware

//echo middleware to check if the user is anonymous

import (
	"encoding/json"
	"net/http"

	"github.com/shadow/backend/internal"
)

func Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		auth2 := r.URL.Query().Get("Authorization")

		if auth != internal.Auth && auth2 != internal.Auth {
			response := map[string]string{"error": "Unauthorized"}

			jsonResponse, _ := json.Marshal(response)

			w.Write(jsonResponse)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
