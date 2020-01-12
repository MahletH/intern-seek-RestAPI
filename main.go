package main

import (
	"html/template"
	"net/http"

	"github.com/WebProgrammingAAiT/intern-seek-web-project/handler"
	"github.com/WebProgrammingAAiT/intern-seek-web-project/internship/repository"
	"github.com/WebProgrammingAAiT/intern-seek-web-project/internship/service"

	"github.com/julienschmidt/httprouter"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var tmpl = template.Must(template.ParseGlob("C:/Users/123/go/src/github.com/WebProgrammingAAiT/intern-seek-web-project/ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "internship.new.layout", nil)
}

func main() {

	dbconn, err := gorm.Open("postgres", "user=postgres password=CaputDraconis dbname=interndb sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	internshipRepo := repository.NewInternshipGormRepo(dbconn)
	internshipSrv := service.NewInternshipService(internshipRepo)
	internshipHandler := handler.NewInternshipHandler(internshipSrv)

	router := httprouter.New()

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	router.GET("/internship/", internshipHandler.GetInternships)
	router.POST("/internship/", internshipHandler.PostInternship)
	router.GET("/internship/:id", internshipHandler.GetSingleInternship) //figure out the link the handler handles
	router.DELETE("/internship/:id", internshipHandler.DeleteInternship)

	http.ListenAndServe("localhost:8181", router)

}
