package common

import (
	"context"
	"log"

	"food-delivery/internal/infra/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config     *config.Config
	Services   *Services
	Storage    *Storage
	Server     *gin.Engine
	Repository *Repository
}

func NewApp(ctx context.Context, config *config.Config) (*App, context.Context) {
	app := &App{
		Config:     config,
		Services:   &Services{},
		Storage:    &Storage{},
		Repository: &Repository{},
	}
	if err := bindStorage(app); err != nil {
		log.Fatalf("bindStorage: error=%v\n", err)
	}
	if err := bindRepo(app); err != nil {
		log.Fatalf("binRepo: error=%v\n", err)
	}
	if err := bindServices(app); err != nil {
		log.Fatalf("bindServices: error=%v\n", err)
	}
	if err := bindServer(app); err != nil {
		log.Fatalf("bindServer: error=%v\n", err)
	}

	return app, ctx
}
