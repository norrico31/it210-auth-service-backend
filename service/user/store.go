package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/norrico31/it210-auth-service-backend/entities"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserById(id int) (*entities.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := new(entities.User)
	for rows.Next() {
		err := scanRowIntoUser(rows, user)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*entities.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := new(entities.User)
	for rows.Next() {
		err = scanRowIntoUser(rows, user)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil

}

func (s *Store) CreateUser(user entities.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password, lastActiveAt) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, nil)
	return err
}

func (s *Store) UpdateUser(user entities.User) error {
	_, err := s.db.Exec("UPDATE users SET firstName = ?, lastName = ?, email = ?, password = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	return err
}

func (s *Store) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (s *Store) SetUserActive(userId int) error {
	_, err := s.db.Exec("UPDATE users SET lastActiveAt = NULL WHERE id = ?", userId)
	return err
}

func (s *Store) UpdateLastActiveTime(userId int, time time.Time) error {
	_, err := s.db.Exec("UPDATE users SET lastActiveAt = ? WHERE id = ?", time, userId)
	return err
}

func scanRowIntoUser(rows *sql.Rows, user *entities.User) error {
	return rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
}
