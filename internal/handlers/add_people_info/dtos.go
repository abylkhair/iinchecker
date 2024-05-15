package add_people_info_handler

import (
	"encoding/json"
	"github.com/wildegor/kaspi-rest/internal/dtos"
	"github.com/wildegor/kaspi-rest/internal/utils"
	"net/http"
)

type AddUserRequestDto struct {
	IIN   string `json:"iin"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (r *AddUserRequestDto) ParseAndValidate(res *dtos.ResponseDto, req *http.Request) error {
	// Checking received data from JSON body. Return status 400 and error message.
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(err.Error())
		return res.JSON()
	}

	// Create a new validator for a AddUserRequestDto.
	validate := utils.NewValidator()

	// Validate fields.
	if err := validate.Struct(r); err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(err.Error()) // ValidatorErrors(err)
	}

	return res.JSON()
}
