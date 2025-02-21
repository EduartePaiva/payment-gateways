package resend

import (
	"bytes"
	"context"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	rsd "github.com/resend/resend-go/v2"
)

var client *rsd.Client

func init() {
	client = rsd.NewClient(env.Config.ResendKey)
}

func SendEmail(ctx context.Context, email string) (*rsd.SendEmailResponse, error) {
	buff := new(bytes.Buffer)
	GenerateEmail().Render(ctx, buff)

	return client.Emails.SendWithContext(ctx, &rsd.SendEmailRequest{
		From:    "Payment Gateway <payment@eduarte.pro>",
		To:      []string{email},
		Html:    buff.String(),
		Subject: "Mystery Box Delivery",
	})
}
