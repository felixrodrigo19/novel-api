package main

import (
	"fmt"

	"github.com/felixrodrigo19/rest-api-go/routes"
)

func main() {
	fmt.Println("Server On")
	routes.HandlerRequest()
}
