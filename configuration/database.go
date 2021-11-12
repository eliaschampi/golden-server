package configuration

import (
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"os"
)

func ConnectToDB() (adapter rel.Adapter, err error) {

	url := os.Getenv("DATABASE_URL")

	adapter, err = postgres.Open(url)

	if err != nil {
		return nil, err
	}

	return adapter, nil

}

