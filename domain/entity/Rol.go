package entity

import (
	"context"

	"github.com/google/uuid"
)

type Rol struct {
	Code        uuid.UUID
	Name        string
	Description string
}

func (r Rol) Table() string {
	return "roles"
}

type RolRepository interface {
	GetAll(c context.Context) ([]*Rol, error)
	//GetByCode
	//Store
}

type RolService interface {
	GetAll(c context.Context) ([]*Rol, error)
	//GetByCode
	//Store
}
