package stripe

import (
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	pkgStripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

// initialize the stripe key
func init() {
	pkgStripe.Key = env.Config.StripeKey
}

func CreateCheckoutSessionURL(itemID string, quantity uint, email string) (string, error) {
	params := &pkgStripe.CheckoutSessionParams{
		LineItems: []*pkgStripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    pkgStripe.String("{{" + itemID + "}}"),
				Quantity: pkgStripe.Int64(int64(quantity)),
			},
		},
		Mode:          pkgStripe.String(string(pkgStripe.CheckoutSessionModePayment)),
		SuccessURL:    pkgStripe.String(env.Config.Domain + "?success=true"),
		CancelURL:     pkgStripe.String(env.Config.Domain + "?canceled=true"),
		CustomerEmail: pkgStripe.String(email),
	}

	s, err := session.New(params)
	return s.URL, err
}
