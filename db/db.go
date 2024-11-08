package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/norrico31/it210-auth-service-backend/config"
)

const addr = "127.0.0.1"

// note in dev mode: change the 3rd arg to localhost or 127.0.0.1 in docker image use config.Envs.DBUser
func NewPostgresStorage() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable",
		config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBUser, config.Envs.DBName, 5432)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
