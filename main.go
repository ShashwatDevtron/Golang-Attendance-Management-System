package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ShashwatDevtron/attendance/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterAttendanceMangementRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":9010",r))
}
