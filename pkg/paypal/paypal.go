package paypal

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	sdk "github.com/plutov/paypal/v4"
)

var client *sdk.Client

func init() {
	apiBase := sdk.APIBaseLive
	if env.Config.GoEnv == "development" {
		apiBase = sdk.APIBaseSandBox
	}
	newClient, err := sdk.NewClient(env.Config.PaypalClientID, env.Config.PaypalSecret, apiBase)
	if err != nil {
		log.Fatal("Error while creating paypal client: ", err)
	}
	newClient.SetLog(os.Stdout)
	client = newClient
}

func CreateOrder(ctx context.Context, quantity uint64) (*sdk.Order, error) {
	units := []sdk.PurchaseUnitRequest{
		sdk.PurchaseUnitRequest{
			Amount: &sdk.PurchaseUnitAmount{
				Currency: "USD",
				Value:    strconv.FormatUint(quantity*1, 10),
				Breakdown: &sdk.PurchaseUnitAmountBreakdown{
					ItemTotal: &sdk.Money{
						Currency: "USD",
						Value:    strconv.FormatUint(quantity*1, 10),
					},
				},
			},
			Items: []sdk.Item{
				sdk.Item{
					Name:     "Mystery Box",
					Quantity: strconv.FormatUint(quantity, 10),
					UnitAmount: &sdk.Money{
						Currency: "USD",
						Value:    "1.00",
					},
				},
			},
		},
	}
	source := &sdk.PaymentSource{}
	appCtx := &sdk.ApplicationContext{}
	return client.CreateOrder(ctx, sdk.OrderIntentCapture, units, source, appCtx)
}
