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

	novelID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid novel ID", http.StatusBadRequest)
		return
	}

	var novel *models.Novel
	result := database.DB.First(&novel, novelID)
	if result.Error != nil {
		http.Error(w, "Error fetching novel", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&novel)
}

func CreateNewNovel(w http.ResponseWriter, r *http.Request) {
	var novel models.Novel

	err := json.NewDecoder(r.Body).Decode(&novel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&novel)
}

func AllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	database.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func Author(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	authorID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var author *models.Author
	result := database.DB.First(&author, authorID)
	if result.Error != nil {
		http.Error(w, "Error fetching author", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&author)
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

func AllGenres(w http.ResponseWriter, r *http.Request) {
	var genres []models.Genre
	database.DB.Find(&genres)
	json.NewEncoder(w).Encode(genres)
}

func Genre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	genreID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid genre ID", http.StatusBadRequest)
		return
	}
	var genre *models.Genre
	result := database.DB.First(&genre, genreID)
	if result.Error != nil {
		http.Error(w, "Error fetching genre", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(genre)
}

func CreateNewGenre(w http.ResponseWriter, r *http.Request) {
	var genre *models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&genre)
}
