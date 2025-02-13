package stripe

import (
	pkgStripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

func CreateCheckoutSessionURL(itemID string, quantity uint, domain string) (string, error) {
	params := &pkgStripe.CheckoutSessionParams{
		LineItems: []*pkgStripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    pkgStripe.String("{{" + itemID + "}}"),
				Quantity: pkgStripe.Int64(int64(quantity)),
			},
		},
		Mode:       pkgStripe.String(string(pkgStripe.CheckoutSessionModePayment)),
		SuccessURL: pkgStripe.String(domain + "?success=true"),
		CancelURL:  pkgStripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)
	return s.URL, err
}
