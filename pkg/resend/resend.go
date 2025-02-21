package resend

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/pkg/env"
	rsd "github.com/resend/resend-go/v2"
)

var client *rsd.Client

func init() {
	client = rsd.NewClient(env.Config.ResendKey)
}

func SendEmail(ctx context.Context, email string) (*rsd.SendEmailResponse, error) {

	return client.Emails.SendWithContext(ctx, &rsd.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"delivered@resend.dev"},
		Html:    "<strong>hello world</strong>",
		Subject: "Hello from Golang",
		Cc:      []string{"cc@example.com"},
		Bcc:     []string{"bcc@example.com"},
		ReplyTo: "replyto@example.com",
	})
}
