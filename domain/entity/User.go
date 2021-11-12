package entity

import "time"

type User struct {
	Dni       string
	Name      string
	Lastname  string
	RolCode   string
	Rol       Rol
	Gender    string
	Image     string
	Address   string
	Phone     string
	Email     string
	Password  string
	CreatedAt time.Time
}
