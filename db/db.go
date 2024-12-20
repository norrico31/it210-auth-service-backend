package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/norrico31/it210-auth-service-backend/config"
)

// note in dev mode: change the 3rd arg to localhost or 127.0.0.1 in docker image use config.Envs.DBUser
func NewPostgresStorage() (*sql.DB, error) {
	prod := config.Envs.DATABASE_URL
	// staging := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable",
	// 	config.Envs.DBUser, config.Envs.DBPassword, "localhost", config.Envs.DBName, 5432)
	db, err := sql.Open("postgres", prod)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
