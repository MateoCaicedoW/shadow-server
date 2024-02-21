package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
		r.Get("/ws", messages.ServeWs)
	})

	return nil
}
