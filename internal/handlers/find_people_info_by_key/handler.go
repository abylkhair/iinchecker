package find_people_info_by_key_handler

import (
	"github.com/wildegor/kaspi-rest/internal/dtos"
	users_repository "github.com/wildegor/kaspi-rest/internal/repositories/users"
	"net/http"
)

type FindPeopleInfoByKeyHandler struct {
	ur users_repository.IUserRepository
}

func NewFindPeopleInfoByKeyHandler(ur users_repository.IUserRepository) *FindPeopleInfoByKeyHandler {
	return &FindPeopleInfoByKeyHandler{}
}

func (h *FindPeopleInfoByKeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *FindPeopleInfoByKeyHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	resp := dtos.NewResponse(w)

	dto := QueryByKeyDto{}

	if err := dto.ParseAndValidate(resp, r); err != nil {
		return err
	}

	um, err := h.ur.FindByKey(dto.Key)
	if err != nil {
		resp.SetError(err.Error())
		return resp.JSON()
	}

	respData := make([]dtos.UserInfoResponseDto, 0)

	for _, model := range um {
		respData = append(respData, dtos.UserInfoResponseDto{
			Name:  model.Name(),
			IIN:   model.IIN,
			Phone: model.Phone,
		})
	}

	resp.SetData(respData)

	return resp.JSON()
}
