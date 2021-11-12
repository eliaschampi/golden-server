package repository

import (
	"context"
	"golden-server/domain/entity"

	"github.com/go-rel/rel"
)

type rolRepository struct {
	kath rel.Repository
}

func NewRolRepository(kath *rel.Repository) entity.RolRepository {
	return &rolRepository{kath: *kath}
}

func (r *rolRepository) GetAll(c context.Context) ([]*entity.Rol, error) {
	var roles []*entity.Rol
	if err := r.kath.FindAll(c, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}
