package common

import (
	"context"
	"log"

	"food-delivery/internal/infra/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config   *config.Config
	Services *Services
	Storage  *Storage
	Server   *gin.Engine
}

func NewApp(ctx context.Context, config *config.Config) (*App, context.Context) {
	app := &App{
		Config:   config,
		Services: &Services{},
		Storage:  &Storage{},
	}
	if err := bindStorage(app); err != nil {
		log.Fatalf("bindStorage: error=%v\n", err)
	}
	if err := bindServices(ctx, app); err != nil {
		log.Fatalf("bindServices: error=%v\n", err)
	}
	if err := bindServer(ctx, app); err != nil {
		log.Fatalf("bindServer: error=%v\n", err)
	}

	return app, ctx
}
