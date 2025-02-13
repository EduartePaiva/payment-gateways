package handlers

import (
	"path"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/gofiber/fiber/v2"
)

func GetOpenApiSpec() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendFile(path.Join(env.Config.BasePath, "/docs/openapi.yaml"), true)
	}
}

func GetScalarHtml() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile(path.Join(env.Config.BasePath, "/docs/scalar.html"), true)
	}
}
