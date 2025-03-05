package controller

import (
	"net/http"

	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.CreateUserService
}

func NewUserHandler(userService *service.CreateUserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (c *UserHandler) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := c.UserService.Register(user.Username, user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Create user successfully", "user": newUser.Username})
}
