package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/pkg/stripe"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
)

// This handler handle a form post request
func CreateCheckoutStripe(db types.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		quantityInt, err := strconv.Atoi(c.FormValue("quantity"))
		email := c.FormValue("email")
		if err != nil || quantityInt < 1 || email == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		session, err := stripe.CreateCheckoutSessionURL(env.Config.StripePriceID, uint(quantityInt), email)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}

		// insert a payment into database
		payment := types.Payment{
			SessionID: session.ID,
			UserEmail: email,
			Status:    "unpaid",
			Item:      "Mystery Box",
			Price:     100,
			Quantity:  uint(quantityInt),
		}
		err = db.CreatePayment(c.Context(), payment)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.Redirect(session.URL)
	}
}
