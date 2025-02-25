package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/pkg/resend"
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

// This handler finalizes the payment
func StripeWebhook(db types.Database, redis types.RedisDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return stripe.ValidateWebhookAndFullfilCheckout(c, func(sessionID string) int {
			return fulfillCheckout(c.Context(), redis, db, sessionID)
		})
	}
}

func fulfillCheckout(ctx context.Context, redis types.RedisDB, db types.Database, sessionID string) int {
	// redis lock logic
	log.Println(sessionID)
	ok, err := redis.LockSessionID(ctx, sessionID)
	if err != nil {
		return http.StatusInternalServerError
	}
	log.Println("ok: ", ok)
	log.Println("err: ", err)
	if !ok {
		return http.StatusServiceUnavailable
	}
	defer redis.DelSessionID(ctx, sessionID)

	// check in mongo if it's already paid
	payment, err := db.GetPayment(ctx, sessionID)
	if err != nil {
		// if there's not a payment on db
		log.Println(err)
		return http.StatusNotFound
	}
	if payment.Status == "paid" {
		// if it's already paid just return ok
		return http.StatusOK
	}

	// send email to the user
	_, err = resend.SendEmail(ctx, payment.UserEmail)
	if err != nil {
		log.Println(err)
		log.Println("Failed to send email to the user")
		return http.StatusServiceUnavailable
	}

	// update payment status to paid
	err = db.MarkStatusAsPaid(ctx, sessionID)
	if err != nil {
		log.Println("Failed to update the database status to paid")
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
