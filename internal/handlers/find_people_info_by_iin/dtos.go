package find_people_info_by_iin_handler

import (
	"github.com/gorilla/mux"
	"github.com/wildegor/kaspi-rest/internal/dtos"
	"github.com/wildegor/kaspi-rest/internal/utils"
	"net/http"
)

type QueryIINDto struct {
	IIN string `json:"iin"`
}

func (r QueryIINDto) ParseAndValidate(res *dtos.ResponseDto, req *http.Request) error {
	params := mux.Vars(req)
	r.IIN = params["iin"]

	validate := utils.NewValidator()

	// Validate fields.
	if err := validate.Struct(&r); err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(err.Error()) // ValidatorErrors(err)
	}

	return res.JSON()
}
