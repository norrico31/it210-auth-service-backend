package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/norrico31/it210-auth-service-backend/cmd/api"
	"github.com/norrico31/it210-auth-service-backend/config"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	godotenv.Load()
	PORT := config.Envs.Port
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL connection established!")
	server := api.NewApiServer(":"+PORT, db, config.Envs)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
