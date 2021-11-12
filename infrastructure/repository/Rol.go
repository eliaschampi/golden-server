package repository

import (
	"context"
	"github.com/go-rel/rel"
	"golden-server/domain/entity"
)

type rolRepository struct {
	kath rel.Repository
}

func NewRolRepository(kath rel.Repository) domainentity.RolRepository {
	return &rolRepository{kath}
}

func (r *rolRepository) GetAll(c context.Context)([] *domainentity.Rol, error)  {
	var roles []*domainentity.Rol
	if err := r.kath.FindAll(c, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}