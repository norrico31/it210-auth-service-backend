package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file" // File source driver
	_ "github.com/lib/pq"                                // PostgreSQL driver

	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/norrico31/it210-auth-service-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	driver, err := postgresMigrate.WithInstance(db, &postgresMigrate.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Get current migration version
	v, d, err := m.Version()
	if err != nil {
		log.Printf("Failed to get migration version: %v", err)

		// Check if the error indicates no migrations have been applied
		if err.Error() == "no migration" {
			log.Println("No migrations have been applied yet.")
		}
		log.Fatalf("Exiting due to the above error")
	}
	log.Printf("Version: %d, dirty: %v", v, d)

	// Handle command line arguments for migration up/down
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations up: %v", err)
		} else {
			log.Println("Migrations applied successfully.")
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations down: %v", err)
		}
	}
}
