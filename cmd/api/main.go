package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/icakces/order-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config: %v", err)
	}

	app := fiber.New()
	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
