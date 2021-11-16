package interfaces

import (
	"context"
	"golden-server/domain/entity"
)

type UserRepository interface {
	GetAll(context context.Context) ([]entity.User, error)
}

type UserService interface {
	GetAll(context context.Context) ([]entity.User, error)
}
