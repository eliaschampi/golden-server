package interfaces

import (
	"context"
	"golden-server/domain/entity"
)

type RolRepository interface {
	GetAll(c context.Context) ([]entity.Role, error)
	//GetByCode
	//Store
}

type RolService interface {
	GetAll(c context.Context) ([]entity.Role, error)
	//GetByCode
	//Store
}
