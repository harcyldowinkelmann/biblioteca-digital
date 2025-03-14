package server

import (
	"ApiSup/internal/config"
	"ApiSup/internal/middlewares"
	"ApiSup/internal/routes"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Validator = &config.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.ConfigureCORS())

	config.Initialize()

	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))
}
