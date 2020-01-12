package main

import (
	"fmt"
	"html/template"
	"net/http"

<<<<<<< HEAD
=======
	"github.com/abdimussa87/Intern-Seek-Version-1/delivery/http/handler"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	userRep "github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/service"
	userServ "github.com/abdimussa87/Intern-Seek-Version-1/user/service"
	"github.com/dgrijalva/jwt-go"
>>>>>>> 93e6acbd0e3224407b062f191fc84d137883d42d
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/nebyubeyene/Intern-Seek-Version-1/delivery/http/handler"
	"github.com/nebyubeyene/Intern-Seek-Version-1/user/repository"
	userRep "github.com/nebyubeyene/Intern-Seek-Version-1/user/repository"
	"github.com/nebyubeyene/Intern-Seek-Version-1/user/service"
	userServ "github.com/nebyubeyene/Intern-Seek-Version-1/user/service"

	_ "github.com/lib/pq"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func main() {

	dbconn, err := gorm.Open("postgres", "user=postgres dbname=gorminterndb password='P@$$w0rDd' sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	// dbconn.DropTableIfExists(&entity.CompanyDetail{}, &entity.User{})
	// errs := dbconn.CreateTable(&entity.UserRole{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := userRep.NewUserGormRepoImpl(dbconn)
	userServi := userServ.NewUserServiceImpl(userRepo)

	compRepo := repository.NewCompanyGormRepoImpl(dbconn)
	compServ := service.NewCompanyService(compRepo)

	userHandler := handler.NewUserHandler(userServi)

	compHandler := handler.NewCompanyHandler(compServ, userServi)

	signUpHandler := handler.NewSignUpHandler(userServi)
	signInHandler := handler.NewSignInHandler(userServi)

	router := httprouter.New()

	router.POST("/v1/signup", signUpHandler.SignUp)
	router.POST("/v1/signin", signInHandler.SignIn)

	//Protected route
	router.GET("/v1/company", compHandler.GetCompanies)
	router.GET("/v1/company/:id", compHandler.GetSingleCompany)
	router.POST("/v1/company", compHandler.PostCompany)
	router.PUT("/v1/company/update/:id", compHandler.PutCompany)
	router.DELETE("/v1/company/delete/:id", compHandler.DeleteCompany)
	router.GET("/v1/users/:id", userHandler.GetSingleUser)
	router.PUT("/v1/user/update/:id", userHandler.PutUser)
	router.GET("/v1/companybyuserid/:id", compHandler.GetSingleCompanyByUserId)

	http.ListenAndServe(":8181", router)

	// mux := http.NewServeMux()
	// fs := http.FileServer(http.Dir("../../ui/assets"))
	// mux.Handle("/assets/", http.StripPrefix("/assets", fs))
	// mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/login", loginHandler)
	// mux.HandleFunc("/signup", userHandler.SignUp)
	// http.ListenAndServe(":8080", mux)
}

//Middleware for checking authorization for viewing a page
func isAuthorizedCompany(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}

//Middleware for checking authorization for viewing a page
func isAuthorizedIntern(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
