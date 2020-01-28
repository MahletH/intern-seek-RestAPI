package main

import (
	"html/template"
	"net/http"

	// appRep "github.com/MahletH/intern-seek-RestAPI/application/repository"
	// appServ "github.com/MahletH/intern-seek-RestAPI/application/service"
	"github.com/abdimussa87/intern-seek-RestAPI/delivery/http/handler"
	intRep "github.com/abdimussa87/intern-seek-RestAPI/internship/repository"
	intServ "github.com/abdimussa87/intern-seek-RestAPI/internship/service"
	"github.com/abdimussa87/intern-seek-RestAPI/user/repository"
	internRep "github.com/abdimussa87/intern-seek-RestAPI/user/repository"
	userRep "github.com/abdimussa87/intern-seek-RestAPI/user/repository"

	fldRep "github.com/abdimussa87/intern-seek-RestAPI/field/repository"
	fldServ "github.com/abdimussa87/intern-seek-RestAPI/field/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/abdimussa87/intern-seek-RestAPI/user/service"
	internServ "github.com/abdimussa87/intern-seek-RestAPI/user/service"
	userServ "github.com/abdimussa87/intern-seek-RestAPI/user/service"

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

	dbconn, err := gorm.Open("postgres", "user=postgres dbname=gorminterndb password='P@$$wOrDd' sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	// dbconn.DropTableIfExists(&entity.Field{})
	// errs := dbconn.CreateTable(&entity.Internship{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := userRep.NewUserGormRepoImpl(dbconn)
	userServi := userServ.NewUserServiceImpl(userRepo)

	// appRepo := appRep.NewApplicationGormRepoImpl(dbconn)
	// appServi := appServ.NewApplicationServiceImpl(appRepo)

	intRepo := intRep.NewInternshipGormRepo(dbconn)
	intServi := intServ.NewInternshipService(intRepo)

	compRepo := repository.NewCompanyGormRepoImpl(dbconn)
	compServ := service.NewCompanyService(compRepo)

	userHandler := handler.NewUserHandler(userServi)

	compHandler := handler.NewCompanyHandler(compServ)

	userroleRepo := userRep.NewUserRoleGormRepo(dbconn)
	userroleServ := userServ.NewUserRoleService(userroleRepo)

	internRepo := internRep.NewInternGormRepoImpl(dbconn)
	internServi := internServ.NewInternService(internRepo)

	internHandler := handler.NewInternHandler(internServi)
	usroleHandler := handler.NewUserRoleHandler(userroleServ)

	fieldRepo := fldRep.NewFieldGormRepo(dbconn)
	fieldServi := fldServ.NewFieldService(fieldRepo)

	fldHandler := handler.NewFieldHandler(fieldServi)

	// appHandler := handler.NewApplicationHandler(appServi)

	intHandler := handler.NewInternshipHandler(intServi, compServ)

	signUpHandler := handler.NewSignUpHandler(userServi)
	signInHandler := handler.NewSignInHandler(userServi, userroleServ)

	srcHandler := handler.NewSearchHandler(fieldServi)

	router := httprouter.New()

	router.POST("/v1/signup", signUpHandler.SignUp)
	router.POST("/v1/signin", signInHandler.SignIn)

	//Protected route

	router.GET("/v1/field/:id/internship", srcHandler.GetInternshipsbyFieldName)
	router.POST("/v1/userrole", usroleHandler.PostUserRole)

	router.GET("/v1/userrole/:id", usroleHandler.GetSingleUserRole)
	router.DELETE("/v1/userrole/delete/:id", usroleHandler.DeleteUserRole)

	router.GET("/v1/company", compHandler.GetCompanies)
	router.GET("/v1/company/:id", compHandler.GetSingleCompany)
	router.POST("/v1/company", compHandler.PostCompany)
	router.PUT("/v1/company/update/:id", compHandler.PutCompany)
	router.DELETE("/v1/company/delete/:id", compHandler.DeleteCompany)
	router.GET("/v1/users/:id", userHandler.GetSingleUser)
	router.PUT("/v1/user/update/:id", userHandler.PutUser)
	router.GET("/v1/companybyuserid/:id", compHandler.GetSingleCompanyByUserId)

	router.GET("/v1/intern", internHandler.GetInterns)
	router.GET("/v1/intern/:id", internHandler.GetSingleIntern)
	router.POST("/v1/intern", internHandler.PostIntern)
	router.PUT("/v1/intern/update/:id", internHandler.PutIntern)
	router.DELETE("/v1/intern/delete/:id", internHandler.DeleteIntern)
	router.GET("/v1/internbyuser/:id", internHandler.GetSingleInternByUserId)

	router.GET("/v1/field", fldHandler.GetFields)
	router.GET("/v1/field/:id", fldHandler.GetSingleField)
	router.POST("/v1/field", fldHandler.PostField)
	router.PUT("/v1/field/update/:id", fldHandler.PutField)
	router.DELETE("/v1/field/delete/:id", fldHandler.DeleteField)

	router.GET("/v1/internships", intHandler.GetInternships)
	router.GET("/v1/internship/:id", intHandler.GetSingleInternship)
	router.POST("/v1/internship", intHandler.PostInternship)
	router.PUT("/v1/internship/update/:id", intHandler.PutInternship)
	router.DELETE("/v1/internship/delete/:id", intHandler.DeleteInternship)
	router.GET("/v1/companyInternship/:company_id/internships", intHandler.GetCompanyInternships)

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
//Already implemented on client side

// func isAuthorizedCompany(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {
// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return []byte("secret"), nil
// 			})
// 			if err != nil {
// 				fmt.Fprintf(w, err.Error())
// 			}

// 			if token.Valid {
// 				endpoint(w, r)
// 			}

// 		} else {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 	})
// }

//Middleware for checking authorization for viewing a page
//Already implemented on client side

// func isAuthorizedIntern(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {
// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return []byte("secret"), nil
// 			})
// 			if err != nil {
// 				fmt.Fprintf(w, err.Error())
// 			}

// 			if token.Valid {
// 				endpoint(w, r)
// 			}

// 		} else {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 	})
// }
