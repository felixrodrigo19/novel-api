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

// Home godoc
//	@Summary		Home
//	@Description	Prints the home message
//	@ID				home
//	@Produce		json
//	@Success		200	{string}	string	"Home"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/ [get]
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Home")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// AllNovels godoc
//	@Summary		List all novels
//	@Description	Lists all novels registered in the database
//	@ID				all-novels
//	@Produce		json
//	@Success		200	{array}		models.Novel	"List of novels"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/novel [get]
func AllNovels(w http.ResponseWriter, r *http.Request) {
	var novels []models.Novel
	if err := database.DB.Find(&novels).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(novels); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Novel godoc
//	@Summary		Get a novel
//	@Description	Gets a specific novel by its ID
//	@ID				get-novel
//	@Param			id	path	int	true	"Novel ID"
//	@Produce		json
//	@Success		200	{object}	models.Novel	"Details of the novel"
//	@Failure		400	{string}	string			"Invalid novel ID"
//	@Failure		404	{string}	string			"Novel not found"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/novel/{id} [get]
func Novel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	novelID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid novel ID", http.StatusBadRequest)
		return
	}

	var novel models.Novel
	result := database.DB.First(&novel, novelID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Novel not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(novel); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// CreateNewNovel godoc
//	@Summary		Create a new novel
//	@Description	Creates a new novel in the database
//	@ID				create-novel
//	@Accept			json
//	@Produce		json
//	@Param			novel	body		models.Novel	true	"Novel data"
//	@Success		201		{object}	models.Novel	"Created novel"
//	@Failure		400		{string}	string			"Invalid input"
//	@Failure		500		{string}	string			"Internal server error"
//	@Router			/novel [post]
func CreateNewNovel(w http.ResponseWriter, r *http.Request) {
	var novel models.Novel

	if err := json.NewDecoder(r.Body).Decode(&novel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, author := range novel.Authors {
		var authorFounded models.Author
		result := database.DB.First(&authorFounded, author.Id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}
		if result.Error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		novel.Authors[i] = &authorFounded
	}

	for i, genre := range novel.Genres {
		var genreFounded models.Genre
		result := database.DB.First(&genreFounded, genre.Id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Genre not found", http.StatusNotFound)
			return
		}
		if result.Error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		novel.Genres[i] = &genreFounded
	}

	if err := database.DB.Create(&novel).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(novel); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// AllAuthors godoc
//	@Summary		List all authors
//	@Description	Lists all authors registered in the database
//	@ID				all-authors
//	@Produce		json
//	@Success		200	{array}		models.Author	"List of authors"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/author [get]
func AllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	if err := database.DB.Find(&authors).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Author godoc
//	@Summary		Get an author
//	@Description	Gets a specific author by their ID
//	@ID				get-author
//	@Param			id	path	int	true	"Author ID"
//	@Produce		json
//	@Success		200	{object}	models.Author	"Details of the author"
//	@Failure		400	{string}	string			"Invalid author ID"
//	@Failure		404	{string}	string			"Author not found"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/author/{id} [get]
func Author(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	authorID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var author models.Author
	result := database.DB.First(&author, authorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Author not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(author); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// CreateNewAuthor godoc
//	@Summary		Create a new author
//	@Description	Creates a new author in the database
//	@ID				create-author
//	@Accept			json
//	@Produce		json
//	@Param			author	body		models.Author	true	"Author data"
//	@Success		201		{object}	models.Author	"Created author"
//	@Failure		400		{string}	string			"Invalid input"
//	@Failure		500		{string}	string			"Internal server error"
//	@Router			/author [post]
func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&author).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(author); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// AllGenres godoc
//	@Summary		List all genres
//	@Description	Lists all genres registered in the database
//	@ID				all-genres
//	@Produce		json
//	@Success		200	{array}		models.Genre	"List of genres"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/genre [get]
func AllGenres(w http.ResponseWriter, r *http.Request) {
	var genres []models.Genre
	if err := database.DB.Find(&genres).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(genres); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Genre godoc
//	@Summary		Get a genre
//	@Description	Gets a specific genre by its ID
//	@ID				get-genre
//	@Param			id	path	int	true	"Genre ID"
//	@Produce		json
//	@Success		200	{object}	models.Genre	"Details of the genre"
//	@Failure		400	{string}	string			"Invalid genre ID"
//	@Failure		404	{string}	string			"Genre not found"
//	@Failure		500	{string}	string			"Internal server error"
//	@Router			/genre/{id} [get]
func Genre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	genreID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid genre ID", http.StatusBadRequest)
		return
	}

	var genre models.Genre
	result := database.DB.First(&genre, genreID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Genre not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(genre); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// CreateNewGenre godoc
//	@Summary		Create a new genre
//	@Description	Creates a new genre in the database
//	@ID				create-genre
//	@Accept			json
//	@Produce		json
//	@Param			genre	body		models.Genre	true	"Genre data"
//	@Success		201		{object}	models.Genre	"Created genre"
//	@Failure		400		{string}	string			"Invalid input"
//	@Failure		500		{string}	string			"Internal server error"
//	@Router			/genre [post]
func CreateNewGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&genre).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(genre); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
