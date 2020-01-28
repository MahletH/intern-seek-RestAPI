package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abdimussa87/intern-seek-RestAPI/entity"
	"github.com/abdimussa87/intern-seek-RestAPI/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type SignInHandler struct {
	userServ user.UserService
	userRole user.UserRoleService
}

func NewSignInHandler(US user.UserService, UR user.UserRoleService) *SignInHandler {
	return &SignInHandler{userServ: US, userRole: UR}
}

type Claims struct {
	UserID uint
	Role   string
	Name   string
	jwt.StandardClaims
}

func (sih *SignInHandler) SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	user := &entity.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usr, err := sih.userServ.UserByUsernameAndPassword(user.Username, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	role, errs := sih.userRole.UserRole(usr.ID)
	if len(errs) > 0 {
		fmt.Println("got into error from getting user role")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claim := &Claims{
		UserID: usr.ID,
		Role:   role.Role,
		Name:   usr.Name,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
