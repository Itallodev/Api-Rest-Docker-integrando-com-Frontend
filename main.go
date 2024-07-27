package main

import (
	"fmt"

	"github.com/devitallo/api-go/database"
	"github.com/devitallo/api-go/routes"
)

func main() {
	database.ConectaDB()
	fmt.Println("Starting GO Server...")
	routes.HandleRequest()
}
