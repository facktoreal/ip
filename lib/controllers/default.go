package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/facktoreal/ip/lib/models"
	"github.com/facktoreal/ip/lib/services"
)

// DefaultControllerInterface ...
type DefaultControllerInterface interface {
	Routes(g *echo.Group)
	Public(c echo.Context) error
}

type defaultController struct {
	stats services.StatsService
}

// NewDefaultController returns a controller
func NewDefaultController(stats services.StatsService) DefaultControllerInterface {
	return &defaultController{
		stats: stats,
	}
}

// Routes registers routes
// @tag.name Public
// @tag.description Public Endpoints
func (ctl *defaultController) Routes(g *echo.Group) {
	g.GET("/", ctl.Public)
}

// Public godoc
// @Summary Public IP service
// @Description return public IP
// @tags Public
// @Produce  json
// @Success 200 {object} models.PublicIpResponse
// @Failure 500 {object} echo.HTTPError
// @Router / [get]
func (ctl *defaultController) Public(c echo.Context) error {
	ip := c.RealIP()

	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusOK, models.PublicIpResponse{IP: ip})
	}

	stats := ctl.stats.Get(c.Request().Context())

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"IP":     ip,
		"Uptime": stats.Uptime.Format("2006-01-02 15:04:05"),
	})
}
