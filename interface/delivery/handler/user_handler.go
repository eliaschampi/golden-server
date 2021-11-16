package handler

import (
	"golden-server/domain/entity"
	"golden-server/interface/presenter"
	"log"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service entity.UserService
}

func NewUserHandler(service entity.UserService, router fiber.Router) {
	handler := &userHandler{service}

	user := router.Group("/user")

	user.Get("/", handler.GetAll)
}

func (h *userHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll(c.Context())

	if err != nil {
		log.Println(err.Error())
		return presenter.JsonResponse(c, fiber.StatusInternalServerError, nil)
	}

	return presenter.JsonResponse(c, fiber.StatusOK, users)
}
