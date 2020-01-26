package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MahletH/intern-seek-RestAPI/entity"
	"github.com/MahletH/intern-seek-RestAPI/internship"

	"github.com/julienschmidt/httprouter"
)

// InternshipHandler handles comment related http requests
type InternshipHandler struct {
	internshipService internship.InternshipService
	//fieldService internship.FieldService
}

// NewInternshipHandler returns new AdminCommentHandler object
func NewInternshipHandler(intrnService internship.InternshipService) *InternshipHandler {
	return &InternshipHandler{internshipService: intrnService}
}

// GetInternships handles GET /internship requests
func (ih *InternshipHandler) GetInternships(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	internship, errs := ih.internshipService.Internships()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(internship, "", "\t\t") //marshal indent indents by given text -> two tabs
	//marshal takes object and returns json?
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//GetSingleInternship takes id from url and return it
func (ih *InternshipHandler) GetSingleInternship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	internship, errs := ih.internshipService.Internship(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(internship, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}

/*
// PutInternships handles PUT /internships/:id requests
func (ih *InternshipHandler) PutInternships(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	internship, errs := ih.internshipService.Internship(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &internship)

	internship, errs = ih.internshipService.UpdateInternship(internship)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(internship, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
*/

// PostInternship handles POST /v1/admin/comments request
func (ih *InternshipHandler) PostInternship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	internship := &entity.Internship{}
	err := json.Unmarshal(body, internship)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	internship, errs := ih.internshipService.StoreInternship(internship)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("internship/%d", internship.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// DeleteInternship handles DELETE internship/:id
func (ih *InternshipHandler) DeleteInternship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ih.internshipService.DeleteInternship(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
