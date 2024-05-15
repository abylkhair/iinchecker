package models

import "net/http"

type IHandler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}
