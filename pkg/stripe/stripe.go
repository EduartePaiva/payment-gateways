package stripe

import (
	"fmt"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	pkgStripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
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

func FulfillCheckout(sessionId string) {
	fmt.Println("Fulfilling Checkout Session " + sessionId)
	// TODO: Make this function safe to run multiple times,
	// even concurrently, with the same session ID

	// TODO: Make sure fulfillment hasn't already been
	// peformed for this Checkout Session

	// Retrieve the Checkout Session from the API with line_items expanded
	params := &pkgStripe.CheckoutSessionParams{}
	params.AddExpand("line_items")

	cs, _ := session.Get(sessionId, params)

	// Check the Checkout Session's payment_status property
	// to determine if fulfillment should be peformed
	if cs.PaymentStatus != pkgStripe.CheckoutSessionPaymentStatusUnpaid {
		// TODO: Perform fulfillment of the line items

		// TODO: Record/save fulfillment status for this
		// Checkout Session
	}
}
