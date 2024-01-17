package routes

import (
	"log"
	"net/http"

	"github.com/felixrodrigo19/rest-api-go/controllers"
	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/novel", controllers.AllNovels)

	log.Fatal(http.ListenAndServe(":8000", r))
}
