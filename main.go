package main

import (
	"fmt"

	"github.com/felixrodrigo19/rest-api-go/models"
	"github.com/felixrodrigo19/rest-api-go/routes"
)

func main() {
	models.Novels = []models.Novel{
		{Id: 1, Title: "Title 1", Description: "Description 1", Language: "Japanese",
			Type: "Novel", Genre: models.Genre{Id: 1, Name: "Romance"},
			Author: models.Author{Id: 1, Name: "Author 1"}, Year: 2020},
	}

	fmt.Println("Server On")
	routes.HandlerRequest()
}
