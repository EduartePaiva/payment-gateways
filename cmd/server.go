package cmd

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/api"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/types"
)

func RunServer(ctx context.Context, db types.Database) error {
	app := api.CreateApp(ctx, db)
	return app.Listen(":" + env.Config.Port)
}
