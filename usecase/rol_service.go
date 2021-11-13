package usecase

import (
	"context"
	"golden-server/domain/entity"
)

type rolService struct {
	rolRepo entity.RolRepository
}

func NewRolService(repo *entity.RolRepository) entity.RolService {
	return &rolService{*repo}
}

func (s *rolService) GetAll(c context.Context) ([]*entity.Rol, error) {
	return s.rolRepo.GetAll(c)
}
