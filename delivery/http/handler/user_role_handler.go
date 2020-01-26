package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MahletH/intern-seek-RestAPI/entity"
	"github.com/MahletH/intern-seek-RestAPI/user"
	"github.com/julienschmidt/httprouter"
)

type UserRoleHandler struct {
	userRoleService user.UserRoleService
}

func NewUserRoleHandler(userRoleSrv user.UserRoleService) *UserRoleHandler {

	return &UserRoleHandler{userRoleService: userRoleSrv}
}

//GetSingleRoles handles GET/v1/admin/roles/:id  requests
func (urh *UserRoleHandler) GetSingleUserRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	userole, errs := urh.userRoleService.UserRole(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(userole, "", "\t\t")

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
func (urh *UserRoleHandler) PostUserRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	userole := &entity.UserRole{}

	err := json.Unmarshal(body, userole)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	userole, errs := urh.userRoleService.StoreUserRole(userole)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(userole, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	p := fmt.Sprintf("/v1/userrole/%d", userole.UserId)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return

}

func (urh *UserRoleHandler) DeleteUserRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := urh.userRoleService.DeleteUserRole(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return

}
