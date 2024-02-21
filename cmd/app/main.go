package main

import (
	"fmt"
	"os"

	"github.com/shadow/backend/internal"
)

func main() {
	s := internal.NewServer("Shadow Chat")

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.StartServer(); err != nil {
		fmt.Println(err)
	}
}
