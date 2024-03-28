package users

import (
	"net/http"

	"food-delivery/internal/application/users"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Users users.IUsers
}

func NewController(
	users users.IUsers,
) *controller {
	return &controller{
		Users: users,
	}
}

func (c *controller) GetUsersById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	user, err := c.Users.GetUserById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func BindUserController(r *gin.Engine, c *controller) {
	users := r.Group("/users")
	users.GET("/:id", c.GetUsersById)
}
