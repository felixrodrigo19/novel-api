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

func CreateNewNovel(w http.ResponseWriter, r *http.Request) {
	var novel models.Novel

	err := json.NewDecoder(r.Body).Decode(&novel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person: %+v", novel)
}

func AllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	database.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func Author(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, author := range models.Authors {
		if strconv.Itoa(author.Id) == id {
			json.NewEncoder(w).Encode(author)
		}
	}
}

func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&author)

}
