package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/abdimussa87/Intern-Seek-Version-1/internship"
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

//GetSingleInternship handles GET/v1/admin/roles/:id  requests
func (ih *InternshipHandler) GetSingleInternship(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	intern, errs := ih.internshipService.Internship(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(intern, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// PutInternships handles PUT /internships/:id requests
func (ih *InternshipHandler) PutInternship(w http.ResponseWriter,
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
		fmt.Println("Unable to parse json")
		return
	}

	internship, errs := ih.internshipService.StoreInternship(internship)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		fmt.Println("Unable to store json")
		return
	}

	p := fmt.Sprintf("internship/%d", internship.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

func (ih *InternshipHandler) DeleteInternship(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

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
