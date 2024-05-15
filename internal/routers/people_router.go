package routers

import (
	"github.com/gorilla/mux"
	get_people_info_handler "github.com/wildegor/kaspi-rest/internal/handlers/add_people_info"
	find_people_info_by_iin_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_iin"
	find_people_info_by_phone_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_key"
)

type PeopleRouter struct {
	pih  *get_people_info_handler.AddPeopleHandler
	fph  *find_people_info_by_iin_handler.FindPeopleInfoByIINHandler
	fpph *find_people_info_by_phone_handler.FindPeopleInfoByKeyHandler
}

func NewPeopleRouter(pih *get_people_info_handler.AddPeopleHandler, fph *find_people_info_by_iin_handler.FindPeopleInfoByIINHandler, fpph *find_people_info_by_phone_handler.FindPeopleInfoByKeyHandler) *PeopleRouter {
	return &PeopleRouter{pih: pih, fph: fph, fpph: fpph}
}

func (pr *PeopleRouter) Setup(r *mux.Router) {
	r.Handle("/api/v1/people/info", pr.pih).Methods("POST")
	r.Handle("/iin/{iin:[0-9]+}", pr.fph).Methods("GET")
	r.Handle("/iin/{iin:[0-9]+}", pr.fpph).Methods("GET")
	r.Handle("/phone/{key}", pr.fpph).Methods("GET")
}
