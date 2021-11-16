package repository

import (
	"context"
	"golden-server/domain/entity"
	"golden-server/domain/interfaces"
)

type rolRepository struct {
	kath *entity.Queries
}

func NewRolRepository(kath *entity.Queries) interfaces.RolRepository {
	return &rolRepository{kath}
}

func (r *rolRepository) GetAll(c context.Context) ([]entity.Role, error) {
	return r.kath.GetRoles(c)
}
