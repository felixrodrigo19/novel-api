package main

import (
	"fmt"

	"github.com/felixrodrigo19/rest-api-go/database"
	"github.com/felixrodrigo19/rest-api-go/routes"
)

func main() {
	database.DBConection()
	fmt.Println("Server On")
	routes.HandlerRequest()
}
