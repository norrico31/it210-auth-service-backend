package main

import (
	"log"

	"github.com/norrico31/it210-auth-service-backend/cmd/seed/seeders"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}

	seeders.SeedUsers(db)
	log.Println("Seeding successfully complete.")
}
