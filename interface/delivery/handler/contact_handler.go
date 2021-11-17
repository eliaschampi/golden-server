package handler

import (
	"golden-server/domain/entity"
	"golden-server/interface/presenter"
	"log"

	"github.com/gofiber/fiber/v2"
)

type contactHandler struct {
	service entity.ContactService
}

func NewContactHandler(service entity.ContactService, router fiber.Router) {

	handler := &contactHandler{service}

	contact := router.Group("/contact")

	contact.Get("/", handler.getAll)

}

func (h *contactHandler) getAll(ctx *fiber.Ctx) error {
	contacts, err := h.service.GetAll(ctx.Context())

	if err != nil {
		log.Println("Error", err.Error())
		return presenter.JsonResponse(ctx, fiber.StatusInternalServerError, nil)
	}

	return presenter.JsonResponse(ctx, fiber.StatusOK, contacts)
}
