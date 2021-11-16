package configuration

import (
	"database/sql"
	"os"
)

func ConnectToDB() (sqlInstance *sql.DB, err error) {

	url := os.Getenv("DATABASE_URL")

	sqlInstance, err = sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return sqlInstance, nil

}
