package main

import (
	"github.com/jgreat/fog-wallet/api"
	"github.com/jgreat/fog-wallet/db"
	"github.com/jgreat/fog-wallet/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Setup and migrate DB
	db := db.Connect()
	models.DoMigrations(db)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(api.ValidationMiddleware())

	// Regster generated API handlers with echo
	fogWalletApi := api.New()
	fogWalletApi.DB = db
	api.RegisterHandlers(e, fogWalletApi)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
