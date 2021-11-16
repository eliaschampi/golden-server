package entity

import (
	"context"

	"github.com/google/uuid"
)

type Role struct {
	Code        uuid.UUID `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type RolRepository interface {
	GetAll(c context.Context) ([]*Role, error)
	//GetByCode
	//Store
}

type RolService interface {
	GetAll(c context.Context) ([]*Role, error)
	//GetByCode
	//Store
}
