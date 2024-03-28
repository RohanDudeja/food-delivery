package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

func bindServer(_ context.Context, app *App) error {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{})
	})

	app.Server = r
	return nil
}
