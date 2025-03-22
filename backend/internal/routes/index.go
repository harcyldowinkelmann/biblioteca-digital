package routes

import (
	controller "biblioteca-digital/internal/controllers"
	"biblioteca-digital/internal/middlewares"
	"biblioteca-digital/internal/repositories"
	"biblioteca-digital/internal/services"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	auth := e.Group("/auth")
	auth.POST("/login", userController.Login)

	protected := e.Group("/api")
	protected.Use(middlewares.VerifyTokenHandler())

	protected.POST("/usuarios", userController.Created)
	protected.GET("/usuarios", userController.Listagem)
	protected.GET("/usuarios/:id", userController.Get)
	protected.PUT("/usuarios/:id", userController.Updated)
	protected.DELETE("/usuarios/:id", userController.Deleted)
}
