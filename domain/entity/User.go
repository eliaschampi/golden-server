package entity

import (
	"context"
	"time"
)

type GetUsersRow struct {
	Dni       string    `json:"dni"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Rol       *Role     `json:"rol"`
	Gender    string    `json:"gender"`
	Image     *string   `json:"image,omitempty"`
	Address   *string   `json:"address,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	GetAll(context context.Context) ([]*GetUsersRow, error)
}

type UserService interface {
	GetAll(context context.Context) ([]*GetUsersRow, error)
}
