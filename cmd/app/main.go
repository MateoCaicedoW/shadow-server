package main

import (
	"fmt"
	"os"

	"github.com/leapkit/core/server"
	"github.com/shadow/backend/internal"
	"github.com/shadow/backend/internal/services"
)

func main() {
	s := internal.NewServer("Shadow")
	dbConn, err := internal.Connection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.Use(server.InCtxMiddleware("userService", services.Users(dbConn)))
	s.Use(server.InCtxMiddleware("jwtService", services.JWT()))
	s.Use(server.InCtxMiddleware("messagesService", services.Messages(dbConn)))
	s.Use(server.InCtxMiddleware("chatService", services.Chats(dbConn)))

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.StartServer(); err != nil {
		fmt.Println(err)
	}
}
