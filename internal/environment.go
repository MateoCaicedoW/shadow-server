package internal

import "github.com/leapkit/core/envor"

var (
	// Environment in which the application is running, this is useful
	// to determine the way we'll run the application, for example, if
	// we're running in production we might want to disable debug mode.
	Environment = envor.Get("GO_ENV", "development")

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envor.Get("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/shadow?sslmode=disable")

	// Port in which the web application listens on.
	Port = envor.Get("PORT", "3000")

	Auth = envor.Get("AUTH", "")
)
