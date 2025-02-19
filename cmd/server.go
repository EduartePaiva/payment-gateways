package cmd

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/api"
	"github.com/EduartePaiva/payment-gateways/pkg/env"
	"github.com/EduartePaiva/payment-gateways/types"
)

func RunServer(ctx context.Context, db types.Database, redis types.RedisDB) error {
	app := api.CreateApp(ctx, db, redis)
	return app.Listen(":" + env.Config.Port)
}
