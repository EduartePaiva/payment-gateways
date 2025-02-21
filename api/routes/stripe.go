package routes

import (
	"github.com/EduartePaiva/payment-gateways/api/handlers"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
	"github.com/resend/resend-go/v2"
)

func StripeRouter(api fiber.Router, db types.Database, email *resend.Client) {
	apiG := api.Group("/stripe")
	apiG.Post("/create-checkout-session", handlers.CreateCheckoutStripe(db))
}
