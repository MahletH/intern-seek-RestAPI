package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/abdimussa87/Intern-Seek-Version-1/user"
	"github.com/julienschmidt/httprouter"
)

type Companyhandler struct {
	companyService user.CompanyService
	userService    user.UserService
}

func NewCompanyHandler(compSrv user.CompanyService, userSrv user.UserService) *Companyhandler {

	return &Companyhandler{companyService: compSrv, userService: userSrv}
}

//GetCompanies handles GET?v1/admin/roles requests
func (ch *Companyhandler) GetCompanies(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	companies, errs := ch.companyService.Companies()

	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}

	output, err := json.MarshalIndent(companies, "", "\t\t")

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//GetSingleRoles handles GET/v1/admin/roles/:id  requests
func (ch *Companyhandler) GetSingleCompany(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	company, errs := ch.companyService.Company(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(company, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (ch *Companyhandler) PutCompany(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	company, errs := ch.companyService.Company(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return

	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &company)

	company, errs = ch.companyService.UpdateCompany(company)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(company, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//TO DO change ps to not get userid
func (ch *Companyhandler) PostCompany(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, errr := strconv.Atoi(ps.ByName("id"))

	if errr != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := ch.userService.User(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	company := &entity.CompanyDetail{}

	err := json.Unmarshal(body, company)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	company.UserID = user.ID
	company, errs = ch.companyService.StoreCompany(company)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/company/%d", company.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return

}

func (ch *Companyhandler) DeleteCompany(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ch.companyService.DeleteCompany(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return

}
