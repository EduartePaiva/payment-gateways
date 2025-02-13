package cmd

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/api"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
)

func RunServer(ctx context.Context) error {
	app := api.CreateApp(ctx)
	return app.Listen(":" + env.Config.Port)
}
