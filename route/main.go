package route

import (
	"golang-web-testing/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world ok!",
		})
	})

	r.POST("/users", user.ControllerRegister)
	r.POST("/users/login", user.ControllerLogin)
	return r
}
