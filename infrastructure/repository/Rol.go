package repository

import (
	"context"
	"database/sql"
	"golden-server/domain/entity"
)

type rolRepository struct {
	kath *sql.DB
}

func NewRolRepository(kath *sql.DB) entity.RolRepository {
	return &rolRepository{kath}
}

func (r *rolRepository) GetAll(c context.Context) ([]*entity.Role, error) {

	query := "select code, name, COALESCE(description, '') from roles"

	var err error
	var rows *sql.Rows

	rows, err = r.kath.QueryContext(c, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []*entity.Role

	for rows.Next() {
		var role entity.Role
		err = rows.Scan(&role.Code, &role.Name, &role.Description)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return roles, nil
}
