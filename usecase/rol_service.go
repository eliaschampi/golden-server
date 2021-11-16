package usecase

import (
	"context"
	"golden-server/domain/entity"
	"golden-server/domain/interfaces"
)

type rolService struct {
	rolRepo interfaces.RolRepository
}

func NewRolService(repo *interfaces.RolRepository) interfaces.RolService {
	return &rolService{*repo}
}

func (s *rolService) GetAll(c context.Context) ([]entity.Role, error) {
	return s.rolRepo.GetAll(c)
}
