package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/shadow/backend/internal/auth"
	"github.com/shadow/backend/internal/messages"
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r *Instance) error {
	// Base middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.AllowAll().Handler)

	go messages.L.Run()

	r.Route("/", func(r chi.Router) {
		// Auth
		r.Post("/auth/login", auth.Login)
		r.Post("/auth/sign-up", auth.SignUp)

		secure := r.With(auth.JWT)
		secure.Get("/current_user", auth.GetCurrentUser)

		// Websocket
		secure.Get("/ws", messages.ServeWs)
	})

	return nil
}
