package routes

import (
	"github.com/EduartePaiva/payment-gateways/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func StripeRouter(api fiber.Router) {
	apiG := api.Group("/stripe")
	apiG.Post("/create-checkout-session", handlers.CreateCheckoutStripe())
}
