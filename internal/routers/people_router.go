package routers

import (
	"github.com/gorilla/mux"
	find_people_info_by_iin_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_iin"
	find_people_info_by_phone_handler "github.com/wildegor/kaspi-rest/internal/handlers/find_people_info_by_phone"
	get_people_info_handler "github.com/wildegor/kaspi-rest/internal/handlers/get_people_info"
)

type PeopleRouter struct {
	pih  *get_people_info_handler.GetPeopleHandler
	fph  *find_people_info_by_iin_handler.FindPeopleInfoByIINHandler
	fpph *find_people_info_by_phone_handler.FindPeopleInfoByPhoneHandler
}

func NewPeopleRouter(pih *get_people_info_handler.GetPeopleHandler, fph *find_people_info_by_iin_handler.FindPeopleInfoByIINHandler, fpph *find_people_info_by_phone_handler.FindPeopleInfoByPhoneHandler) *PeopleRouter {
	return &PeopleRouter{pih: pih, fph: fph, fpph: fpph}
}

func (pr *PeopleRouter) Setup(api *mux.Route) {

	pc := api.PathPrefix("/people")
	pci := pc.PathPrefix("/info")

	pci.Handler(pr.pih)
	pci.PathPrefix("/iin/{iin:[0-9]+}").Handler(pr.fph)
	pci.PathPrefix("/iin/{iin:[0-9]+}").Handler(pr.fpph)
	pci.PathPrefix("/phone/{key}").Handler(pr.fpph)
}
