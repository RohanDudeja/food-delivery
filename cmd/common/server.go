package common

import (
	"food-delivery/internal/application/controller/users"

	"github.com/gin-gonic/gin"
)

func bindServer(app *App) error {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{})
	})

	userController := users.NewController(app.Services.Users)
	users.BindUserController(r, userController)

	app.Server = r
	return nil
}
