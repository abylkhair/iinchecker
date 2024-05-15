package find_people_info_by_key_handler

import (
	"github.com/gorilla/mux"
	"github.com/wildegor/kaspi-rest/internal/dtos"
	"github.com/wildegor/kaspi-rest/internal/utils"
	"net/http"
)

type QueryByKeyDto struct {
	Key string `json:"key"`
}

func (r QueryByKeyDto) ParseAndValidate(res *dtos.ResponseDto, req *http.Request) error {
	params := mux.Vars(req)
	r.Key = params["key"]

	validate := utils.NewValidator()

	// Validate fields.
	if err := validate.Struct(&r); err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(err.Error()) // ValidatorErrors(err)
	}

	return res.JSON()
}
