package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ngc224/phck/model"
)

type Controller struct {
	Health model.Health
}

func NewController(h model.Health) *Controller {
	return &Controller{
		Health: h,
	}
}

func (cntr *Controller) HealthCheck(c echo.Context) error {
	cntr.Health.StatusCode = HealthStatusCode(cntr.Health.UpdateHealth())
	return c.JSON(cntr.Health.StatusCode, cntr.Health)
}

func HealthStatusCode(ok bool) int {
	if ok {
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
