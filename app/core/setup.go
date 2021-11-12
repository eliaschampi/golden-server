package core

import (
	"github.com/go-rel/rel"
	"github.com/gofiber/fiber/v2"
	"golden-server/infrastructure/repository"
	"golden-server/interface/deliveryhttp"
	"golden-server/usecase/Rol"
)

func StartApp(relRepository rel.Repository, app *fiber.App)  {

	rolRepo := repository.NewRolRepository(relRepository)

	rolService := Rol.NewService(&rolRepo)

	router := setupApiV1(app)
	//send to delivery here
	deliveryhttp.NewRolHandler(rolService, router)
}

func setupApiV1(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
	//middlewares here
}