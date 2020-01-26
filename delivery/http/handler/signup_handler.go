package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/abdimussa87/Intern-Seek-Version-1/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type SignUpHandler struct {
	userServ user.UserService
}
type claims struct {
	UserID uint
	Name   string
	jwt.StandardClaims
}

//TO DO
// var jwtKey = os.Getenv("MY_JWT_TOKEN")

func NewSignUpHandler(US user.UserService) *SignUpHandler {
	return &SignUpHandler{userServ: US}
}

func (suh *SignUpHandler) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	user := &entity.User{}

	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {

		fmt.Println(err)
		json.NewEncoder(w).Encode(err)

	}

	user.Password = string(pass)

	createdUser, errs := suh.userServ.StoreUser(user)

	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &claims{
		UserID: createdUser.ID,
		Name:   createdUser.Name,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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
