package api

import (
	"github.com/gofiber/fiber/v2"
	"setu-engine/database"
	"setu-engine/models"
)

func CreateBridge(service database.Service) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var bridge models.Bridge
		if err := ctx.BodyParser(&bridge); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

		}
		if err := service.CreateBridge(bridge); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Bridge created successfully"})
	}
}