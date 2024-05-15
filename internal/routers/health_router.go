package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HealthRouter struct {
}

func NewHealthRouter() *HealthRouter {
	return &HealthRouter{}
}

func (hr *HealthRouter) Setup(api *mux.Route) {
	v1 := api.Subrouter().PathPrefix("/v1")
	
	v1.Path("/health/check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
}
