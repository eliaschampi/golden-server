package repository

import (
	"context"
	"golden-server/domain/entity"

	"github.com/go-rel/rel"
)

type userRepository struct {
	kath rel.Repository
}

func NewUserRepository(kath *rel.Repository) entity.UserRepository {
	return &userRepository{kath: *kath}
}

func (u *userRepository) GetAll(c context.Context) ([]*entity.User, error) {
	var users []*entity.User

	if err := u.kath.FindAll(c, &users); err != nil {
		return nil, err
	}

	return users, nil
}
