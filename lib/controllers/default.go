package controllers

import (
	"net/http"

	"github.com/facktoreal/env"
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
	stats    services.StatsService
	hostname string
}

// NewDefaultController returns a controller
func NewDefaultController(stats services.StatsService) DefaultControllerInterface {
	return &defaultController{
		stats:    stats,
		hostname: env.MayGetString("HOSTNAME"),
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

	if c.Request().Header.Get("Accept") == "application/json" || c.QueryParam("format") == "json" || c.QueryParam("json") != "" {
		return c.JSON(http.StatusOK, models.PublicIpResponse{IP: ip})
	}

	if c.Request().Header.Get("Accept") == "application/xml" || c.QueryParam("format") == "xml" || c.QueryParam("xml") != "" {
		return c.XML(http.StatusOK, models.PublicIpResponse{IP: ip})
	}

	if c.Request().Header.Get("Accept") == "text/plain" || c.QueryParam("format") == "text" || c.QueryParam("plain") != "" {
		return c.String(http.StatusOK, ip)
	}

	stats := ctl.stats.Get(c.Request().Context())

	// add backup for empty hostname
	if ctl.hostname == "" {
		ctl.hostname = c.Request().Host
	}

	if ctl.hostname != "" && ctl.hostname[len(ctl.hostname)-1] != '/' {
		ctl.hostname += "/"
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"IP":       ip,
		"HOSTNAME": ctl.hostname,
		"Uptime":   stats.Uptime.Format("2006-01-02 15:04:05"),
	})
}
