package iin_check_handler

import "net/http"

type CheckIINHandler struct {
}

func NewCheckIINHandler() *CheckIINHandler {
	return &CheckIINHandler{}
}

func (h *CheckIINHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *CheckIINHandler) Handle(writer http.ResponseWriter, request *http.Request) error {
	// TODO
	return nil
}
