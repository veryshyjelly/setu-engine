package routes

import (
	"github.com/gofiber/fiber/v2"
	"setu-engine/api"
	"setu-engine/database"
)

func RegisterBridgeRoutes(service database.Service, app *fiber.App) {
	app.Post("/bridge", api.CreateBridge(service))
	app.Get("/bridges", api.GetAllBridges(service))
	app.Get("/bridge/:fromChatID", api.GetBridge(service))
	app.Delete("/bridge", api.DeleteBridge(service))
}