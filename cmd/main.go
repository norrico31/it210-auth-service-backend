package main

import (
	"fmt"
	"log"

	"github.com/norrico31/it210-auth-service-backend/cmd/api"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("PostgreSQL connection established!")
	server := api.NewApiServer(":8081", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
