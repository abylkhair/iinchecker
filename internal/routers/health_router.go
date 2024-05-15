package routers

import (
	"github.com/gorilla/mux"
	health_check_handler "github.com/wildegor/kaspi-rest/internal/handlers/health_check"
)

type HealthRouter struct {
	hh *health_check_handler.HealthCheckHandler
}

func NewHealthRouter(hh *health_check_handler.HealthCheckHandler) *HealthRouter {
	return &HealthRouter{hh}
}

func (hr *HealthRouter) Setup(r *mux.Router) {
	r.Handle("/api/v1/health", hr.hh).Methods("GET")
}
