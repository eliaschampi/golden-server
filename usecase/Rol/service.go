package Rol

import (
	"context"
	"golden-server/domain/entity"
)

type service struct {
	rolRepository entity.RolRepository
}

func NewService(repo *entity.RolRepository) entity.RolService {
	return &service{*repo}
}

func (s *service) GetAll(c context.Context) ([]*entity.Rol, error) {
	roles, err := s.rolRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return roles, nil
}
