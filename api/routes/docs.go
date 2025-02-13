package routes

import (
	"github.com/EduartePaiva/payment-gateways/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func DocsRouter(api fiber.Router) {
	api.Get("/docs/openapi.yaml", handlers.GetOpenApiSpec())
	api.Get("/docs/reference", handlers.GetScalarHtml())
}
