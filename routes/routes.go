package routes

import (
	"log"
	"net/http"

	"github.com/felixrodrigo19/rest-api-go/controllers"
)

func HandlerRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
