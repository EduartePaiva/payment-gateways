package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/pkg/stripe"
	"github.com/gofiber/fiber/v2"
)

// This handler handle a form post request
func CreateCheckoutStripe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		quantityInt, err := strconv.Atoi(c.FormValue("quantity"))
		email := c.FormValue("email")
		if err != nil || quantityInt < 1 || email == "" {
			fmt.Println("aqui?")
			fmt.Println(quantityInt)
			fmt.Println("email:" + email)
			return c.SendStatus(http.StatusBadRequest)
		}

		url, err := stripe.CreateCheckoutSessionURL(env.Config.StripePriceID, uint(quantityInt), email)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.Redirect(url)
	}
}
