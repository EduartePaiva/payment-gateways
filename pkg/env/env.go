package env

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type envVariables struct {
	BasePath             string `env:"BASE_PATH,default=."`
	GoEnv                string `env:"GO_ENV,default=production"`
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
