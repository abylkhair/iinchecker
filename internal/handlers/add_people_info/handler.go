package add_people_info_handler

import (
	"github.com/wildegor/kaspi-rest/internal/dtos"
	users_repository "github.com/wildegor/kaspi-rest/internal/repositories/users"
	"net/http"
)

type AddPeopleHandler struct {
	ur users_repository.IUserRepository
}

func NewAddPeopleHandler(ur users_repository.IUserRepository) *AddPeopleHandler {
	return &AddPeopleHandler{ur}
}

func (h *AddPeopleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *AddPeopleHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.NewResponse(w)
	dto := &AddUserRequestDto{}

	if err := dto.ParseAndValidate(resp, r); err != nil {
		return err
	}

	create := &users_repository.CreateUserModel{
		IIN:   dto.IIN,
		Name:  dto.Name,
		Phone: dto.Phone,
	}

	_, err := h.ur.CreateIfNotExists(create)
	if err != nil {
		resp.SetError(err.Error())
	}

	return resp.JSON()
}
