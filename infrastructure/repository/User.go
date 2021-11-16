package repository

import (
	"context"
	"golden-server/domain/entity"
	"golden-server/domain/interfaces"
)

type userRepository struct {
	kath *entity.Queries
}

func NewUserRepository(kath *entity.Queries) interfaces.UserRepository {
	return &userRepository{kath}
}

func (u *userRepository) GetAll(c context.Context) ([]entity.User, error) {
	return u.kath.GetUsers(c)
}
