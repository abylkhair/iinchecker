package find_people_info_by_phone_handler

import "net/http"

type FindPeopleInfoByPhoneHandler struct {
}

func NewFindPeopleInfoByPhoneHandler() *FindPeopleInfoByPhoneHandler {
	return &FindPeopleInfoByPhoneHandler{}
}

func (h *FindPeopleInfoByPhoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *FindPeopleInfoByPhoneHandler) Handle(writer http.ResponseWriter, request *http.Request) error {
	// TODO
	return nil
}
