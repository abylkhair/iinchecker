package health_check_handler

import (
	"github.com/wildegor/kaspi-rest/internal/dtos"
	"net/http"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *HealthCheckHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.NewResponse(w)
	return resp.JSON()
}
