package auth

import (
	"net/http"

	"github.com/shadow/backend/internal/services"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtService := r.Context().Value("jwtService").(services.JWTService)
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

		token = token[len("Bearer "):]
		_, err := jwtService.ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		next.ServeHTTP(w, r)
	})
}
