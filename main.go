package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/handlers"
	"store/models"
	"store/utils"
)

func main() {
	// Database connection
	db := utils.InitDB()
	db.AutoMigrate(&models.Product{})

	// Set up the router
	router := mux.NewRouter()

	// Define routes and handlers
	router.HandleFunc("/products", handlers.GetProducts(db)).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProduct(db)).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct(db)).Methods("POST")

	// Start the server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
