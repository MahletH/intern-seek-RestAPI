package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/abdimussa87/Intern-Seek-Version-1/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type SignInHandler struct {
	userServ user.UserService
}

func NewSignInHandler(US user.UserService) *SignInHandler {
	return &SignInHandler{userServ: US}
}

type Claims struct {
	UserID uint
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
	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &Claims{
		UserID: usr.ID,
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
