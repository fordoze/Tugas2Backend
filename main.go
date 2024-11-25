package main

import (
	"log"
	"perpus/app/controllers"
	"perpus/app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.Connet()

	perpustakaan := app.Group("/perpustakaan")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	perpustakaan.Post("/createPerpustakaan", controllers.CreatePerpustakaan)
	perpustakaan.Get("/getPerpustakaan", controllers.GetAllDataPerpustakaan)
	perpustakaan.Put("/update", controllers.UpdatedAtPerpustakaan)
	perpustakaan.Delete("/delete", controllers.DeleteDataPerpustakaan)

	log.Fatal(app.Listen(":3000"))
}
