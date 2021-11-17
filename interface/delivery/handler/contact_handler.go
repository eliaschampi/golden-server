package handler

import (
	"golden-server/domain/entity"
	"golden-server/domain/rules"
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
	contact.Post("/", handler.create)

}

func (h *contactHandler) getAll(ctx *fiber.Ctx) error {
	contacts, err := h.service.GetAll(ctx.Context())

	if err != nil {
		log.Println("Error", err.Error())
		return presenter.JsonResponse(ctx, fiber.StatusInternalServerError, nil)
	}

	return presenter.JsonResponse(ctx, fiber.StatusOK, contacts)
}

func (h *contactHandler) create(ctx *fiber.Ctx) error {
	contact := &rules.ContactStruct{}

	err := ctx.BodyParser(contact)

	if err != nil {
		log.Println("Error", err.Error())
		return presenter.JsonResponse(ctx, fiber.StatusInternalServerError, nil)
	}

	code, valErrors, err := h.service.Create(ctx.Context(), contact)

	if valErrors != nil {
		return presenter.JsonResponse(ctx, fiber.StatusBadRequest, valErrors)
	}

	if err != nil {
		log.Println("Error", err.Error())
		return presenter.JsonResponse(ctx, fiber.StatusInternalServerError, nil)
	}

	return presenter.JsonResponse(ctx, fiber.StatusOK, fiber.Map{
		"code": code,
	})
}
