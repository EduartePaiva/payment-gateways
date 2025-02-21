package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/pkg/stripe"
	"github.com/EduartePaiva/payment-gateways/types"
	"github.com/gofiber/fiber/v2"
	"github.com/resend/resend-go/v2"
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
func StripeWebhook(db types.Database, redis types.RedisDB, email *resend.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// guessing it's this way
		sessionID := c.Params("session_id")
		if sessionID == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.SendStatus(fulfillCheckout(c.Context(), redis, db, sessionID, email))
	}
}

func fulfillCheckout(ctx context.Context, redis types.RedisDB, db types.Database, sessionID string, email *resend.Client) int {
	// redis lock logic
	ok, err := redis.LockSessionID(ctx, sessionID)
	if err != nil {
		return http.StatusInternalServerError
	}
	if !ok {
		return http.StatusServiceUnavailable
	}
	defer redis.DelSessionID(ctx, sessionID)

	// check in mongo if it's already paid
	payment, err := db.GetPayment(sessionID)
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
	// // TODO: make a nice email template
	// resp, err := email.Emails.SendWithContext(ctx, &resend.SendEmailRequest{
	// 	From:    "Acme <onboarding@resend.dev>",
	// 	To:      []string{"delivered@resend.dev"},
	// 	Html:    "<strong>hello world</strong>",
	// 	Subject: "Hello from Golang",
	// 	Cc:      []string{"cc@example.com"},
	// 	Bcc:     []string{"bcc@example.com"},
	// 	ReplyTo: "replyto@example.com",
	// })

	return http.StatusOK
}
