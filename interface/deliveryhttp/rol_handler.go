package deliveryhttp

import (
	"github.com/gofiber/fiber/v2"
	"golden-server/domain/entity"
	"golden-server/interface/presenter"
)

type rolHandler struct {
	service domainentity.RolService
}

func NewRolHandler(service domainentity.RolService, router fiber.Router) {
	handler := &rolHandler{service: service}

	rol := router.Group("/roles")
	rol.Get("/", handler.getAll)

}

func (rh *rolHandler) getAll(c *fiber.Ctx) error  {

	roles, err := rh.service.GetAll(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	data := make([]interface{}, len(roles))

	for i, rol := range roles{
		data[i] = presenter.MapRol(*rol)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": data,
	})

}









