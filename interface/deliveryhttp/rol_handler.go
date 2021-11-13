package deliveryhttp

import (
	"golden-server/domain/entity"
	"golden-server/interface/presenter"
	"log"

	"github.com/gofiber/fiber/v2"
)

type rolHandler struct {
	service entity.RolService
}

func NewRolHandler(service entity.RolService, router fiber.Router) {
	handler := &rolHandler{service}

	rol := router.Group("/roles")
	rol.Get("/", handler.getAll)

}

func (rh *rolHandler) getAll(c *fiber.Ctx) error {

	roles, err := rh.service.GetAll(c.Context())

	if err != nil {

		log.Println("Error", err.Error()) //save logs to file

		return presenter.JsonResponse(c, fiber.StatusInternalServerError, nil)
	}

	data := make([]interface{}, len(roles))

	for i, rol := range roles {
		data[i] = presenter.MapRol(*rol)
	}

	return presenter.JsonResponse(c, fiber.StatusOK, data)

}
