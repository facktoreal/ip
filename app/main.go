package main

import (
	"github.com/facktoreal/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/facktoreal/ip/lib/controllers"
	"github.com/facktoreal/ip/lib/providers/mock"
	"github.com/facktoreal/ip/lib/services"
)

func main() {
	e := echo.New()

	// Hide banner
	e.HideBanner = true

	// Load environment vars
	if err := env.Init(true); err != nil {
		e.Logger.Fatalf("Unable to load environment variables, err: %s", err.Error())
	}

	port := env.MayGetString("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Infof("Defaulting to port %s", port)
	}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	var (
		statsSrv  = services.NewStatsService()
		healthSrv = services.NewHealthService(mock.NewHealthRepository())
	)

	// Core
	controllers.NewHealthController(healthSrv, statsSrv).Routes(e.Group("api"))
	controllers.NewDefaultController().Routes(e.Group(""))

	e.Logger.Infof("Server started, v%s | port: %s", echo.Version, port)
	e.Logger.Fatal(e.Start(":" + port))
}
