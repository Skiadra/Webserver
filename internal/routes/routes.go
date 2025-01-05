package routes

import (
	"webserver/internal/controllers"
	"webserver/internal/repositories"
	"webserver/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize layers
	userRepo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userController := &controllers.UserController{Service: userService}

	// Define routes
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.ListUsers)
	}
}
