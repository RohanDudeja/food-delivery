package main

import (
	"context"

	"food-delivery/cmd/common"
	"food-delivery/internal/infra/config"
)

func main() {
	// initalise context
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// initialise config
	cnf := config.BuildConfig()

	// initialise App
	app, _ := common.NewApp(ctx, cnf)

	app.Server.Run(cnf.ServerURL())
}
