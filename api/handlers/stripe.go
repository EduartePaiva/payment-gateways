package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EduartePaiva/payment-gateways/pkg/stripe"
	"github.com/gofiber/fiber/v2"
)

// This handler handle a form post request
func CreateCheckout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		itemID := c.FormValue("item_id")
		quantityInt, err := strconv.Atoi(c.FormValue("quantity"))
		email := c.FormValue("email")
		if err != nil || quantityInt < 1 || itemID == "" || email == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		url, err := stripe.CreateCheckoutSessionURL(itemID, uint(quantityInt), email)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.Redirect(url)
	}
}
