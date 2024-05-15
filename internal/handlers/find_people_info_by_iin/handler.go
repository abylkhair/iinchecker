package find_people_info_by_iin_handler

import "net/http"

type FindPeopleInfoByIINHandler struct {
}

func NewFindPeopleInfoByIINHandler() *FindPeopleInfoByIINHandler {
	return &FindPeopleInfoByIINHandler{}
}

func (h *FindPeopleInfoByIINHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *FindPeopleInfoByIINHandler) Handle(writer http.ResponseWriter, request *http.Request) error {
	// TODO
	return nil
}
