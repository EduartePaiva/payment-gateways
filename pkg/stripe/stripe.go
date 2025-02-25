package stripe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/gofiber/fiber/v2"
	pkgStripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/webhook"
)

// initialize the stripe key
func init() {
	pkgStripe.Key = env.Config.StripeKey
}

func CreateCheckoutSessionURL(priceID string, quantity uint, email string) (*pkgStripe.CheckoutSession, error) {
	params := &pkgStripe.CheckoutSessionParams{
		LineItems: []*pkgStripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    pkgStripe.String(priceID),
				Quantity: pkgStripe.Int64(int64(quantity)),
			},
		},
		Mode:          pkgStripe.String(string(pkgStripe.CheckoutSessionModePayment)),
		SuccessURL:    pkgStripe.String(env.Config.Domain + "?success=true"),
		CancelURL:     pkgStripe.String(env.Config.Domain + "?canceled=true"),
		CustomerEmail: pkgStripe.String(email),
	}
	return session.New(params)
}

func CheckPaymentIsPaid(sessionID string) bool {
	params := &pkgStripe.CheckoutSessionParams{}
	params.AddExpand("line_items")
	cs, _ := session.Get(sessionID, params)
	return cs.PaymentStatus == pkgStripe.CheckoutSessionPaymentStatusPaid
}

func ValidateWebhookAndFullfilCheckout(c *fiber.Ctx, fullfilCheckout func(SessionID string) int) error {
	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// Use the secret provided by Stripe CLI for local testing
	// or your webhook endpoint's secret.
	event, err := webhook.ConstructEvent(c.Body(), c.Get("Stripe-Signature"), env.Config.StripeEndpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		return c.SendStatus(http.StatusBadRequest) // Return a 400 error on a bad signature
	}

	if event.Type == pkgStripe.EventTypeCheckoutSessionCompleted ||
		event.Type == pkgStripe.EventTypeCheckoutSessionAsyncPaymentSucceeded {
		var cs pkgStripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &cs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			return c.SendStatus(http.StatusBadRequest) // Return a 400 error on a bad signature
		}

		return c.SendStatus(fullfilCheckout(cs.ID))

	}
	return c.SendStatus(http.StatusAccepted)
}
