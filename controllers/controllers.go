package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felixrodrigo19/rest-api-go/database"
	"github.com/felixrodrigo19/rest-api-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllNovels(w http.ResponseWriter, r *http.Request) {
	var novels []models.Novel
	database.DB.Find(&novels)
	json.NewEncoder(w).Encode(novels)
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
