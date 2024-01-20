package routes

import (
	"log"
	"net/http"

	"github.com/felixrodrigo19/rest-api-go/controllers"
	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/novel", controllers.AllNovels).Methods("GET")
	r.HandleFunc("/novel/{id}", controllers.Novel).Methods("GET")
	r.HandleFunc("/novel", controllers.CreateNewNovel).Methods("POST")

	r.HandleFunc("/author", controllers.AllAuthors).Methods("GET")
	r.HandleFunc("/author/{id}", controllers.Author).Methods("GET")
	r.HandleFunc("/author", controllers.CreateNewAuthor).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
