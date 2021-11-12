package Rol

import (
	"context"
	"golden-server/domain/entity"
)

type service struct {
	rolRepository domainentity.RolRepository
}

func NewService(repo *domainentity.RolRepository) domainentity.RolService {
	return &service{*repo}
}

func (s *service) GetAll(c context.Context) ([]*domainentity.Rol, error) {
	roles, err := s.rolRepository.GetAll(c)

	if err != nil {
		return nil, err
	}

	return roles, nil
}


