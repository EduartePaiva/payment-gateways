package api

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/api/routes"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp(ctx context.Context, db types.Database, redis types.RedisDB) *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")
	if env.Config.GoEnv != "production" {
		api.Use(logger.New())
	}
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})
	routes.DocsRouter(api)
	routes.StripeRouter(api, db)
	return app
}
