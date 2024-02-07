package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"setu-engine/bridge"
	"setu-engine/database"
	"setu-engine/routes"
)

func main() {
	app := fiber.New()

	bsr := bridge.NewService()
	bsr.Run()

	dbs := database.Connect(bsr)
	routes.RegisterBridgeRoutes(dbs, app)

	log.Fatalln(app.Listen(":8070"))
}