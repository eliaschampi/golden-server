package configuration

import "github.com/gofiber/fiber/v2"

const (
	success         = fiber.StatusOK
	notfound        = fiber.StatusNotFound
	internal        = fiber.StatusInternalServerError
	unauthorized    = fiber.StatusUnauthorized
	unauthenticated = fiber.StatusForbidden
	validated       = fiber.StatusBadRequest
)

func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(success).JSON(fiber.Map{
		"status":  success,
		"message": message,
		"data":    data,
	})
}

func ResponseNotFound(c *fiber.Ctx) error {
	return c.Status(notfound).JSON(fiber.Map{
		"status":  notfound,
		"message": "La información solicitada no ha sido encontrada",
		"data":    nil,
	})
}

func ResponseInternalError(c *fiber.Ctx, message string) error {
	return c.Status(internal).JSON(fiber.Map{
		"status":  internal,
		"message": message,
		"data":    nil,
	})
}

func ResponseUnauthorized(c *fiber.Ctx) error {
	return c.Status(unauthorized).JSON(fiber.Map{
		"status":  unauthorized,
		"message": "Respuesta no autorizada",
		"data":    nil,
	})
}

func ResponseUnauthenticated(c *fiber.Ctx) error {
	return c.Status(unauthenticated).JSON(fiber.Map{
		"status":  unauthenticated,
		"message": "Se requiere un token de seguridad",
		"data":    nil,
	})
}

func ResponseValidated(c *fiber.Ctx) error {
	return c.Status(validated).JSON(fiber.Map{
		"status":  validated,
		"message": "Algunos campos son inválidos",
		"data":    nil,
	})
}
