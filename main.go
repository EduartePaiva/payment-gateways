package main

import (
	"context"
	"log"

	"github.com/EduartePaiva/payment-gateways/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Fatal(cmd.RunServer(ctx))
}
