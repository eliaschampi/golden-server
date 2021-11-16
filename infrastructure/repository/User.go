package repository

import (
	"context"
	"database/sql"
	"golden-server/domain/entity"
)

type userRepository struct {
	kath *sql.DB
}

func NewUserRepository(kath *sql.DB) entity.UserRepository {
	return &userRepository{kath}
}

func (u *userRepository) GetAll(c context.Context) ([]*entity.GetUsersRow, error) {

	query := `select dni, users.name, lastname, gender, image, address, phone, email, created_at,
				roles.code, roles.name from users
				inner join roles on users.rol_code = roles.code
	         `

	var rows *sql.Rows
	var err error

	rows, err = u.kath.QueryContext(c, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*entity.GetUsersRow, 0)

	for rows.Next() {
		var user entity.GetUsersRow
		var role entity.Role
		err = rows.Scan(
			&user.Dni,
			&user.Name,
			&user.Lastname,
			&user.Gender,
			&user.Image,
			&user.Address,
			&user.Phone,
			&user.Email,
			&user.CreatedAt,
			&role.Code,
			&role.Name,
		)

		if err != nil {
			return nil, err
		}

		if role != (entity.Role{}) {
			user.Rol = &role
		}

		users = append(users, &user)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, nil

}
