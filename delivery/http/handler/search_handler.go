package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abdimussa87/intern-seek-RestAPI/field"
	"github.com/julienschmidt/httprouter"
)

type SearchHandler struct {
	fieldService field.FieldService
}

func NewSearchHandler(fldSrv field.FieldService) *SearchHandler {

	return &SearchHandler{fieldService: fldSrv}
}

func (sh *SearchHandler) GetInternshipsbyFieldName(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	name := ps.ByName("id")

	field, errs := sh.fieldService.GetFieldbyName(name)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	intern, errss := sh.fieldService.FieldInternships(field)

	fmt.Println(intern)

	if len(errss) > 0 {
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
