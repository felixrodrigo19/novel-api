package main

import (
	"fmt"
	"github.com/felixrodrigo19/rest-api-go/database"
	"github.com/felixrodrigo19/rest-api-go/routes"
)

// @title			Novel API Documentation
// @version		2.0.0
// @description	The Novel Backend API is designed to facilitate the management and distribution of novel data.
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			http://localhost:8000
// @BasePath		/
func main() {
	database.DBConnection()
	fmt.Println("Server On")
	routes.HandlerRequest()
}
