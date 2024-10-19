package seeders

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/norrico31/it210-auth-service-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check user table: %v", err)
	}

	if count > 0 {
		// log.Println("User table already has data. Skipping seeing.")
		return
	}
	hashPassword := func(password string) string {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		return string(hashedPassword)
	}

	users := []entities.User{
		{
			FirstName: "Mary Grace",
			LastName:  "Bitmal",
			Age:       20,
			Email:     "mvbitmal@up.edu.ph",
			Password:  hashPassword("secret.123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			FirstName: "Chester",
			LastName:  "Francisco",
			Age:       19,
			Email:     "cgfrancisco@up.edu.ph",
			Password:  hashPassword("secret.123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			FirstName: "Norrico Gerald",
			LastName:  "Biason",
			Age:       18,
			Email:     "nmbiason@up.edu.ph",
			Password:  hashPassword("secret.123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, user := range users {
		go func(u entities.User) {
			defer wg.Done()

			_, err := db.Exec(`
			INSERT INTO users (firstName, lastName, age, email, password, createdAt, updatedAt) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
				u.FirstName,
				u.LastName,
				u.Age,
				u.Email,
				u.Password,
				u.CreatedAt,
				u.UpdatedAt)
			if err != nil {
				log.Printf("Failed to insert user %s: %v", u.Email, err)
			} else {
				log.Printf("Inserted user: %s", u.Email)
			}
		}(user)
	}
	wg.Wait()
	log.Println("Users table seeded successfully.")
}
