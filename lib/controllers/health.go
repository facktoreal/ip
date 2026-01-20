package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stiks/helpers"

	"github.com/facktoreal/ip/lib/models"
	services2 "github.com/facktoreal/ip/lib/services"
)

// HealthControllerInterface ...
type HealthControllerInterface interface {
	Routes(g *echo.Group)
	HealthCheck(c echo.Context) error
}

type healthController struct {
	health services2.HealthService
	stats  services2.StatsService
}

// NewHealthController returns a controller
func NewHealthController(healthSrv services2.HealthService, statsSrv services2.StatsService) HealthControllerInterface {
	return &healthController{
		health: healthSrv,
		stats:  statsSrv,
	}
}

// Routes registers routes
// @tag.name Public
// @tag.description Public Endpoints
func (ctl *healthController) Routes(g *echo.Group) {
	g.GET("/healthz", ctl.HealthCheck)
}

// HealthCheck godoc
// @Summary Health Status
// @Description Report system health status
// @tags Public
// @Produce  json
// @Success 200 {object} models.Health
// @Failure 500 {object} echo.HTTPError
// @Router /healthz [get]
func (ctl *healthController) HealthCheck(c echo.Context) error {
	ctx := c.Request().Context()

	if err := ctl.health.Check(ctx); err != nil {
		return c.JSON(http.StatusOK, echo.Map{"healthy": false, "error": err.Error()})
	}

	stats := ctl.stats.Get(ctx)

	return c.JSON(http.StatusOK, models.Health{
		Healthy: true,
		Uptime:  helpers.DurationToString(time.Now().Sub(stats.Uptime)),
	})
}
