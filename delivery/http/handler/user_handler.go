package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abdimussa87/Intern-Seek-Version-1/user"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	userServ user.UserService
}

func NewUserHandler(US user.UserService) *UserHandler {
	return &UserHandler{userServ: US}
}

//GetUsers handles GET/v1/users requests
func (uh *UserHandler) GetUsers(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	users, errs := uh.userServ.Users()

	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//GetSingleUser handles GET/v1/users/:id  requests
func (uh *UserHandler) GetSingleUser(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userServ.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
