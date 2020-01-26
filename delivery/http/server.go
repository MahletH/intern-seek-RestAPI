package main

import (
	"fmt"
	"html/template"
	"net/http"

	appRep "github.com/abdimussa87/Intern-Seek-Version-1/application/repository"
	appServ "github.com/abdimussa87/Intern-Seek-Version-1/application/service"
	"github.com/abdimussa87/Intern-Seek-Version-1/delivery/http/handler"
	intRep "github.com/abdimussa87/Intern-Seek-Version-1/internship/repository"
	intServ "github.com/abdimussa87/Intern-Seek-Version-1/internship/service"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	userRep "github.com/abdimussa87/Intern-Seek-Version-1/user/repository"
	"github.com/abdimussa87/Intern-Seek-Version-1/user/service"
	userServ "github.com/abdimussa87/Intern-Seek-Version-1/user/service"

	// "github.com/abdimussa87/Intern-Seek-Version-1/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

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

	dbconn, err := gorm.Open("postgres", "user=postgres dbname=interndb password='1234' sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	// dbconn.DropTableIfExists(&entity.CompanyDetail{}, &entity.User{}, &entity.Application{}, &entity.PersonalDetails{}, &entity.Internship{})
	// errs := dbconn.CreateTable(&entity.User{}, &entity.CompanyDetail{}, &entity.Application{}, &entity.PersonalDetails{}, &entity.Internship{},&entity.Fields{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := userRep.NewUserGormRepoImpl(dbconn)
	userServi := userServ.NewUserServiceImpl(userRepo)

	appRepo := appRep.NewApplicationGormRepoImpl(dbconn)
	appServi := appServ.NewApplicationServiceImpl(appRepo)

	intRepo := intRep.NewInternshipGormRepo(dbconn)
	intServi := intServ.NewInternshipService(intRepo)

	compRepo := repository.NewCompanyGormRepoImpl(dbconn)
	compServ := service.NewCompanyService(compRepo)

	userHandler := handler.NewUserHandler(userServi)

	compHandler := handler.NewCompanyHandler(compServ, userServi)

	appHandler := handler.NewApplicationHandler(appServi)

	intHandler := handler.NewInternshipHandler(intServi)

	signUpHandler := handler.NewSignUpHandler(userServi)
	signInHandler := handler.NewSignInHandler(userServi)

	router := httprouter.New()

	router.POST("/v1/signup", signUpHandler.SignUp)
	router.POST("/v1/signin", signInHandler.SignIn)

	//Protected route
	router.GET("/v1/companies", compHandler.GetCompanies)
	router.GET("/v1/company/:id", compHandler.GetSingleCompany)
	router.POST("/v1/company", compHandler.PostCompany)
	router.PUT("/v1/company/update/:id", compHandler.PutCompany)
	router.DELETE("/v1/company/delete/:id", compHandler.DeleteCompany)

	router.GET("/v1/users/:id", userHandler.GetSingleUser)
	router.GET("/v1/users/", userHandler.GetUsers)

	router.GET("/v1/application", appHandler.GetApplications)
	router.GET("/v1/application/:id", appHandler.GetSingleApplication)
	router.POST("/v1/application", appHandler.PostApplication)
	router.PUT("/v1/application/update/:id", appHandler.PutApplication)
	router.DELETE("/v1/application/delete/:id", appHandler.DeleteApplication)

	router.GET("/v1/internships", intHandler.GetInternships)
	router.GET("/v1/internship/:id", intHandler.GetSingleInternship)
	router.POST("/v1/internship", intHandler.PostInternship)
	router.PUT("/v1/internship/update/:id", intHandler.PutInternship)
	router.DELETE("/v1/internship/delete/:id", intHandler.DeleteInternship)

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
