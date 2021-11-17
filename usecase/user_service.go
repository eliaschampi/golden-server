package usecase

import (
	"context"
	"golden-server/domain/entity"
)

type userService struct {
	userRepository entity.UserRepository
}

func NewUserService(repo *entity.UserRepository) entity.UserService {
	return &userService{*repo}
}

func (s *userService) GetAll(c context.Context) ([]*entity.GetUsersRow, error) {
	return s.userRepository.GetAll(c)
}

func (s *userService) Create(c context.Context) (code string, err error) {
	return
}
