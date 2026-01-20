package main

import (
	"html/template"
	"io"

	"github.com/facktoreal/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/facktoreal/ip/lib/controllers"
	"github.com/facktoreal/ip/lib/providers/mock"
	"github.com/facktoreal/ip/lib/services"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &Template{
		templates: template.Must(template.ParseGlob("lib/views/*.html")),
	}
	e.Renderer = renderer

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
	controllers.NewDefaultController(statsSrv).Routes(e.Group(""))

	e.Static("/static", "static")

	e.Logger.Infof("Server started, v%s | port: %s", echo.Version, port)
	e.Logger.Fatal(e.Start(":" + port))
}
