package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/felixrodrigo19/rest-api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func AllNovels(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Novels)
}
