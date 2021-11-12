package deliveryhttp

import (
	domain "golden-server/domain/entity"
	"golden-server/interface/presenter"

	"github.com/gofiber/fiber/v2"
)

type rolHandler struct {
	service domain.RolService
}

func NewRolHandler(service domain.RolService, router fiber.Router) {
	handler := &rolHandler{service: service}

	rol := router.Group("/roles")
	rol.Get("/", handler.getAll)

}

func (rh *rolHandler) getAll(c *fiber.Ctx) error {

	roles, err := rh.service.GetAll(c.Context())

	if err != nil {
		return presenter.JsonResponse(c, fiber.StatusInternalServerError, nil)
		// log error to file
	}

	data := make([]interface{}, len(roles))

	for i, rol := range roles {
		data[i] = presenter.MapRol(*rol)
	}

	return presenter.JsonResponse(c, fiber.StatusOK, data)

}
