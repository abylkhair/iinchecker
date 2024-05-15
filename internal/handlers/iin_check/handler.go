package iin_check_handler

import (
	"github.com/wildegor/kaspi-rest/internal/dtos"
	"net/http"
)

type CheckIINHandler struct {
}

func NewCheckIINHandler() *CheckIINHandler {
	return &CheckIINHandler{}
}

func (h *CheckIINHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *CheckIINHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.NewResponse(w)

	dto := QueryByIINDto{}
	if err := dto.ParseAndValidate(resp, r); err != nil {
		return err
	}

	return resp.JSON()
}
