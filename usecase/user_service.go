package usecase

import (
	"context"
	"golden-server/domain/entity"
	"golden-server/domain/interfaces"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(repo *interfaces.UserRepository) interfaces.UserService {
	return &userService{*repo}
}

func (s *userService) GetAll(c context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(c)
}
