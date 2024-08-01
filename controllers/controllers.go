package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felixrodrigo19/rest-api-go/database"
	"github.com/felixrodrigo19/rest-api-go/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	message, err := fmt.Fprint(w, "Home")
	if err != nil {
		return
	}
	fmt.Print(message)
}

func AllNovels(w http.ResponseWriter, r *http.Request) {
	var novels []models.Novel
	database.DB.Find(&novels)
	err := json.NewEncoder(w).Encode(novels)
	if err != nil {
		return
	}
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
	errEnc := json.NewEncoder(w).Encode(&novel)
	if errEnc != nil {
		return
	}
}

func CreateNewNovel(w http.ResponseWriter, r *http.Request) {
	var novel *models.Novel

	err := json.NewDecoder(r.Body).Decode(&novel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, author := range novel.Authors {
		var authorFounded models.Author

		result := database.DB.First(&authorFounded, author.Id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Record Not Found", http.StatusBadRequest)
			return
		}
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		novel.Authors[i] = &authorFounded
	}

	for i, genre := range novel.Genres {
		var genreFounded models.Genre

		result := database.DB.First(&genreFounded, genre.Id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Record Not Found", http.StatusBadRequest)
			return
		}
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		novel.Genres[i] = &genreFounded
	}

	database.DB.Create(&novel)
	ErrEnc := json.NewEncoder(w).Encode(novel)
	if ErrEnc != nil {
		return
	}
}

func AllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	database.DB.Find(&authors)
	err := json.NewEncoder(w).Encode(authors)
	if err != nil {
		return
	}
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
	errEnc := json.NewEncoder(w).Encode(&author)
	if errEnc != nil {
		return
	}
}

func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {
	var author *models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&author)
	errEnc := json.NewEncoder(w).Encode(author)
	if errEnc != nil {
		return
	}
}

func AllGenres(w http.ResponseWriter, r *http.Request) {
	var genres []models.Genre
	database.DB.Find(&genres)
	err := json.NewEncoder(w).Encode(genres)
	if err != nil {
		return
	}
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
	errEnc := json.NewEncoder(w).Encode(genre)
	if errEnc != nil {
		return
	}
}

func CreateNewGenre(w http.ResponseWriter, r *http.Request) {
	var genre *models.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Create(&genre)
	errEnc := json.NewEncoder(w).Encode(genre)
	if errEnc != nil {
		return
	}
}
