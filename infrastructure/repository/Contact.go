package repository

import (
	"context"
	"database/sql"
	"golden-server/domain/entity"
	"golden-server/domain/rules"
)

type contactRepository struct {
	kath *sql.DB
}

func NewContactRepository(kath *sql.DB) entity.ContactRepository {
	return &contactRepository{kath}
}

func (ct *contactRepository) GetAll(c context.Context) ([]*entity.Contact, error) {

	var rows *sql.Rows
	var err error

	query := "select * from contacts"

	rows, err = ct.kath.QueryContext(c, query)

	if err != nil {
		return nil, err
	}

	contacts := make([]*entity.Contact, 0)

	for rows.Next() {
		var contact entity.Contact

		err = rows.Scan(
			&contact.Code,
			&contact.Name,
			&contact.Type,
			&contact.Email,
			&contact.Phone,
			&contact.Address,
			&contact.Description,
			&contact.CreatedAt,
			&contact.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (ct *contactRepository) Create(c context.Context, contact *rules.ContactStruct) (code string, err error) {

	query := `insert into contacts (name, type, email, phone, address, description)
			  values ($1, $2, $3, $4, $5, $6) returning code
	         `

	row := ct.kath.QueryRowContext(c, query,
		contact.Name,
		contact.Type,
		contact.Email,
		contact.Phone,
		contact.Address,
		contact.Description,
	)

	err = row.Scan(&code)

	return

}
