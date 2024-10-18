package entities

import (
	"time"

	_ "github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	UpdateUser(User) error
	DeleteUser(int) error
	SetUserActive(int) error
	UpdateLastActiveTime(int, time.Time) error
}

type User struct {
	ID           int        `json:"id"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	Email        string     `json:"email"`
	Password     string     `json:"-"`
	LastActiveAt *time.Time `json:"lastActiveAt"`
	// IsActive     bool       `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserRegisterPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdatePayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password,omitempty"` // Optional, for password update
}
