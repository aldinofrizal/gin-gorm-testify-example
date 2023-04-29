package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ControllerRegister(ctx *gin.Context) {
	var user UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userService := UserService{Repository: UserRepository{}}
	createdUser, err := userService.Register(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success register",
		"user":    createdUser.GetBasicResponse(),
	})
}

func ControllerLogin(ctx *gin.Context) {
	var loginReq LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userService := UserService{Repository: UserRepository{}}
	user, err := userService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user":    user.GetBasicResponse(),
	})
}
