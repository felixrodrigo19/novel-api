package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felixrodrigo19/rest-api-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllNovels(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Novels)
}

func Novel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, novel := range models.Novels {
		if strconv.Itoa(novel.Id) == id {
			json.NewEncoder(w).Encode(novel)
		}
	}
}
