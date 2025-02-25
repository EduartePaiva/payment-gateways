package env

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
)

type goEnv string

// this enforces goEnv type when running env decode
func (i *goEnv) Decode(repl string) error {
	switch repl {
	case "development":
	case "Production":
	default:
		return fmt.Errorf("error decoding the value of: %s for GO_ENV", repl)
	}
	*i = goEnv(repl)
	return nil
}

type envVariables struct {
	BasePath string `env:"BASE_PATH,default=."`
	// GoEnv is default of "production" or it can be set to "development"
	GoEnv                goEnv  `env:"GO_ENV,default=production"`
	FrontendURL          string `env:"FRONTEND_URL,default="`
	Port                 string `env:"PORT,default=3000"`
	StripeKey            string `env:"STRIPE_KEY"`
	StripeMysteryBoxID   string `env:"STRIPE_MYSTERY_BOX_ID"`
	StripePriceID        string `env:"STRIPE_PRICE_ID"`
	StripeEndpointSecret string `env:"STRIPE_ENDPOINT_SECRET"`
	Domain               string `env:"DOMAIN"`
	MongoURI             string `env:"MONGODB_URI,required"`
	RedisURI             string `env:"REDIS_URI,required"`
	ResendKey            string `env:"RESEND_KEY,required"`
	PaypalClientID       string `env:"PAYPAL_CLIENT_ID,required"`
	PaypalSecret         string `env:"PAYPAL_SECRET,required"`
}

var (
	Config envVariables
)

func init() {
	loadAndValidateEnv()
}

// loadAndValidateEnv loads environment variables and validates their presence.
func loadAndValidateEnv() {
	Config = envVariables{}
	err := envdecode.Decode(&Config)
	if err != nil {
		log.Fatal("❌ ERROR DECODING ENVIRONMENT VARIABLES: ", err)
	}
}
