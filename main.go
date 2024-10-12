package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/francisihe/golang-task-manager-api/routes"
)

func main() {
	// Print Hello World message
	helloWorld()

	// Initialize the router
	router := routes.SetupRouter()

	// Start the HTTP server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func helloWorld() {
	fmt.Println("Hello World\nFrancis is trying out Go!")
}
