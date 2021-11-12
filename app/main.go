package main

import (
	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"
	"golden-server/app/core"
	"golden-server/configuration"
	"log"

	_ "github.com/lib/pq"
)

func main()  {

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

	relRepository := rel.New(adapter)

	app := fiber.New()

	core.StartApp(relRepository, app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("te amo kath")
	})

	log.Fatalln(app.Listen(":3000"))
}
