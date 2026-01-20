package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/facktoreal/ip/lib/models"
)

// DefaultControllerInterface ...
type DefaultControllerInterface interface {
	Routes(g *echo.Group)
	Public(c echo.Context) error
}

type defaultController struct {
}

// NewDefaultController returns a controller
func NewDefaultController() DefaultControllerInterface {
	return &defaultController{}
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

	return c.JSON(http.StatusOK, models.PublicIpResponse{IP: c.RealIP()})
}
