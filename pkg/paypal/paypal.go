package paypal

import (
	"log"

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
	client = newClient
}

func PaySomething() {
}
