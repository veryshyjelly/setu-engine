package api

import (
	"github.com/gofiber/fiber/v2"
	"setu-engine/database"
)

func GetBridge(service database.Service) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fromChatID := ctx.Params("fromChatID")
		toChatIDs, err := service.GetBridge(fromChatID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"toChatIDs": toChatIDs})
	}
}