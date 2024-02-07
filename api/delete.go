package api

import (
	"github.com/gofiber/fiber/v2"
	"setu-engine/database"
	"setu-engine/models"
)

func DeleteBridge(service database.Service) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var bridge models.Bridge
		if err := ctx.BodyParser(&bridge); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := service.DeleteBridge(bridge.FromChatID, bridge.SecondChatID); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Bridge deleted successfully"})
	}
}