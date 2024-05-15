package find_people_info_by_iin_handler

import (
	"github.com/wildegor/kaspi-rest/internal/dtos"
	users_repository "github.com/wildegor/kaspi-rest/internal/repositories/users"
	"net/http"
)

type FindPeopleInfoByIINHandler struct {
	ur users_repository.IUserRepository
}

func NewFindPeopleInfoByIINHandler(ur users_repository.IUserRepository) *FindPeopleInfoByIINHandler {
	return &FindPeopleInfoByIINHandler{}
}

func (h *FindPeopleInfoByIINHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *FindPeopleInfoByIINHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.NewResponse(w)

	dto := QueryIINDto{}

	if err := dto.ParseAndValidate(resp, r); err != nil {
		return err
	}

	um, err := h.ur.FindByIIN(dto.IIN)
	if err != nil {
		resp.SetError(err.Error())
		return resp.JSON()
	}

	resp.SetData([]dtos.UserInfoResponseDto{
		{
			Name:  um.Name(),
			IIN:   um.IIN,
			Phone: um.Phone,
		},
	})

	return resp.JSON()
}
