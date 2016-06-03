package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ngc224/hck/config"
)

type Controller struct {
	Config *config.Config
}

func NewController(c *config.Config) *Controller {
	return &Controller{
		Config: c,
	}
}

func (cntr *Controller) HealthCheck(c echo.Context) error {
	cntr.Config.Health.StatusCode = http.StatusOK

	for _, v := range cntr.Config.Health.Process {
		if !v.IsProcess() {
			cntr.Config.Health.StatusCode = http.StatusInternalServerError
		}
	}

	return c.JSON(cntr.Config.Health.StatusCode, cntr.Config.Health)
}
