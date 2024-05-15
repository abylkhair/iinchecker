package routers

import (
	"github.com/gorilla/mux"
	iin_check_handler "github.com/wildegor/kaspi-rest/internal/handlers/iin_check"
)

type CheckRouter struct {
	ich *iin_check_handler.CheckIINHandler
}

func NewCheckRouter(ich *iin_check_handler.CheckIINHandler) *CheckRouter {
	return &CheckRouter{ich}
}

func (cr *CheckRouter) Setup(r *mux.Router) {
	r.Handle("/api/v1/iin_check/{iin:[0-9]+}", cr.ich).Methods("GET")
}
