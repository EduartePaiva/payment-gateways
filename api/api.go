package api

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/api/routes"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/resend/resend-go/v2"
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
	resendClient := resend.NewClient(env.Config.ResendKey)

	routes.DocsRouter(api)
	routes.StripeRouter(api, db, resendClient)
	return app
}
