package handler

import (
	"github.com/nebyubeyene/Intern-Seek-Version-1/user"
)

type UserHandler struct {
	userServ user.UserService
}

func NewUserHandler(US user.UserService) *UserHandler {
	return &UserHandler{userServ: US}
}

// //SignUp handles requests coming at /signup
// func (uh UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		user := entity.User{}
// 		user.Name = r.FormValue("fullname")
// 		// user.UUID = r.FormValue("username")
// 		user.Email = r.FormValue("email")
// 		user.Phone = r.FormValue("phone")
// 		user.Password = r.FormValue("password")

// 		err := uh.userServ.StoreUser(&user)
// 		if err != nil {
// 			panic(err)
// 		}

// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	} else if r.Method == http.MethodGet {
// 		uh.t.ExecuteTemplate(w, "signup.html", nil)
// 	}
// }
