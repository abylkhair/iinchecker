package get_people_info_handler

import "net/http"

type GetPeopleHandler struct {
}

func NewGetPeopleHandler() *GetPeopleHandler {
	return &GetPeopleHandler{}
}

func (h *GetPeopleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(w, r)
}

func (h *GetPeopleHandler) Handle(writer http.ResponseWriter, request *http.Request) error {
	// TODO
	return nil
}
