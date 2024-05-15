package routers

import "github.com/gorilla/mux"

type IRouter interface {
	Setup(api *mux.Route)
}
