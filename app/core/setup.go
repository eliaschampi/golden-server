package core

import (
	"database/sql"
	"golden-server/infrastructure/repository"
	"golden-server/interface/delivery/handler"
	"golden-server/interface/presenter"
	"golden-server/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartApp(sqlInstance *sql.DB, app *fiber.App) {

	rolRepo := repository.NewRolRepository(sqlInstance)
	userRepo := repository.NewUserRepository(sqlInstance)
	contactRepo := repository.NewContactRepository(sqlInstance)

	rolService := usecase.NewRolService(&rolRepo)
	userService := usecase.NewUserService(&userRepo)
	contactService := usecase.NewContactService(&contactRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("te amo kath")
	})

	router := setupApiV1(app)

	//send to delivery here
	presenter.NewMessages()

	handler.NewRolHandler(rolService, router)
	handler.NewUserHandler(userService, router)
	handler.NewContactHandler(contactService, router)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Elias, i can't find that!")
	})
}

func setupApiV1(app *fiber.App) fiber.Router {

	app.Use(logger.New())

	return app.Group("/api/v1")

}
