package routes

import (
	"github.com/EduartePaiva/payment-gateways/api/handlers"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
)

func StripeRouter(api fiber.Router, db types.Database) {
	apiG := api.Group("/stripe")
	apiG.Post("/create-checkout-session", handlers.CreateCheckoutStripe(db))
}
