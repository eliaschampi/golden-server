package main

import (
	"database/sql"
	"fmt"
	"golden-server/app/core"
	"golden-server/configuration"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {

	var err error
	var sqlInstance *sql.DB

	err = configuration.LoadEnvFile()
	if err != nil {
		panic(err)
	}

	sqlInstance, err = configuration.ConnectToDB()

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	core.StartApp(sqlInstance, app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
