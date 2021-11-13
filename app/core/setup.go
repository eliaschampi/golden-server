package core

import (
	"golden-server/infrastructure/repository"
	"golden-server/interface/deliveryhttp"
	"golden-server/interface/presenter"
	"golden-server/usecase"

	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartApp(kathRelInstance *rel.Repository, app *fiber.App) {

	rolRepo := repository.NewRolRepository(kathRelInstance)
	userRepo := repository.NewUserRepository(kathRelInstance)

	rolService := usecase.NewRolService(&rolRepo)
	userService := usecase.NewUserService(&userRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("te amo kath")
	})

	router := setupApiV1(app)

	//send to delivery here
	presenter.NewMessages()

	deliveryhttp.NewRolHandler(rolService, router)
	deliveryhttp.NewUserHandler(userService, router)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Elias, i can't find that!")
	})
}

func setupApiV1(app *fiber.App) fiber.Router {

	app.Use(logger.New())

	return app.Group("/api/v1")

}
