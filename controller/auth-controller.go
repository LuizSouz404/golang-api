package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

//NewAuthController creates a new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello ma men",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello ma girl",
	})
}
