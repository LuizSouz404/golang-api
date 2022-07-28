package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizsouz404/api-rest-go/config"
	"github.com/luizsouz404/api-rest-go/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	server.Run(":5000")
}
