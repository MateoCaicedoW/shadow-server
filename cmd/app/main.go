package main

import (
	"fmt"
	"os"

	"github.com/shadow/backend/internal"
	"github.com/shadow/backend/server"
)

func main() {
	s := server.New("Shadow Chat")

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
