package configuration

import (
	"errors"
	"github.com/joho/godotenv"
	"time"
)

func LoadEnvFile() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = errors.New("Cannot load your .env files.")
	}
	return
}

func TokenExpirationTime(seconds int16) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}
