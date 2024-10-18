package main

import (
	"fmt"
	"log"

	"github.com/norrico31/it210-auth-service-backend/cmd/api"
	"github.com/norrico31/it210-auth-service-backend/config"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage(
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName,
		5432,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("PostgreSQL connection established!")
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
