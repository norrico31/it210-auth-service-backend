package main

import (
	"fmt"
	"log"

	"github.com/norrico31/it210-auth-service-backend/config"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage(
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName,
		5432, // Default PostgreSQL port
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// You can now use the db connection...
	fmt.Println("PostgreSQL connection established!")
}
