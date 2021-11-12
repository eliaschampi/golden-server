package presenter

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var messages map[int]string

func NewMessages() {
	messages = map[int]string{
		fiber.StatusOK:                  "Su operación finalizó con exito",
		fiber.StatusInternalServerError: "Ocurrió un error inesperado",
		fiber.StatusNotFound:            "La información solicidada no esta disponible",
		fiber.StatusUnauthorized:        "Esta información no es accesible",
		fiber.StatusForbidden:           "Se requiere un token de seguridad",
		fiber.StatusBadRequest:          "Información proporcionada ha sido rechazada",
	}
}

func JsonResponse(c *fiber.Ctx, statusCode int, data interface{}) error {

	resp := Response{
		Code:    statusCode,
		Message: messages[statusCode],
		Data:    data,
	}

	return c.Status(statusCode).JSON(resp)
}
