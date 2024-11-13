package entities

import "time"

type UserStore interface {
	Login(UserLoginPayload) (string, User, error)
}

type User struct {
	ID           int        `json:"id"`
	FirstName    string     `json:"firstName"`
	Age          int        `json:"age"`
	LastName     string     `json:"lastName"`
	Email        string     `json:"email"`
	Password     string     `json:"-"`
	LastActiveAt *time.Time `json:"lastActiveAt"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	Token        string     `json:"token"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty"` // Optional, for password update
}
