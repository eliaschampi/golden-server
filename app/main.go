package main

import (
	"fmt"
	"golden-server/app/core"
	"golden-server/configuration"
	"log"
	"os"

	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {

	var err error
	var adapter rel.Adapter

	err = configuration.LoadEnvFile()
	if err != nil {
		panic(err)
	}

	adapter, err = configuration.ConnectToDB()

	if err != nil {
		panic(err)
	}

	kathRelInstance := rel.New(adapter)

	app := fiber.New()

	core.StartApp(&kathRelInstance, app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
