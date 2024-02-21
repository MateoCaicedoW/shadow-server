package main

import (
	"fmt"

	"github.com/leapkit/core/db"
	"github.com/shadow/backend/internal"
	"github.com/shadow/backend/internal/migrations"
)

func main() {

	err := db.Create(internal.DatabaseURL)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Database created successfully")

	conn, err := internal.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Migrations ran successfully")
}
