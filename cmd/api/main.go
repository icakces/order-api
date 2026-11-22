package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/icakces/order-api/internal/config"
	"github.com/icakces/order-api/internal/platform"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	db, err := platform.NewDatabase(ctx, cfg)
	cancel()
	if err != nil {
		log.Fatalf("Database: %v", err)
	}
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	app := fiber.New()
	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
