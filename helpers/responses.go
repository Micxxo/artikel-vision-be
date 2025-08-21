package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(c *fiber.Ctx, status int, err interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   true,
		"message": err,
	})
}

func SendSuccessResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
