package controller

import (
	"net/http"

	"github.com/nissy/bon/render"
	"github.com/nissy/phck"
)

type health struct {
	health phck.Health
}

func NewHealth(h phck.Health) *health {
	return &health{
		health: h,
	}
}

func (h *health) HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.Json(w, healthStatusCode(h.health.IsHealth()), h.health)
}

func healthStatusCode(ok bool) int {
	if ok {
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
