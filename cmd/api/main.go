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

	app := fiber.New()
	app.Get("/healthz", func(c fiber.Ctx) error {
		reqCtx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()
		err := db.WithContext(reqCtx).Exec("SELECT 1").Error
		if err != nil {
			log.Printf("Database ping failed: %v", err)
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "503"})
		}
		return c.JSON(fiber.Map{"status": "ok"})
	})
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
