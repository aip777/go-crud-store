package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/auth"
	"store/handlers"
	"store/middleware"
	"store/models"
	"store/utils"
)

func main() {
	// Database connection
	utils.Init()
	db := utils.InitDB()
	db.AutoMigrate(&models.Product{})

	// Set up the router
	router := mux.NewRouter()

	// Define routes and handlers
	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.Handle("/products", middleware.Authenticate(http.HandlerFunc(handlers.GetProducts(db)))).Methods("GET")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.GetProduct(db)))).Methods("GET")
	router.Handle("/products", middleware.Authenticate(http.HandlerFunc(handlers.CreateProduct(db)))).Methods("POST")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.UpdateProduct(db)))).Methods("PUT")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.DeleteProduct(db)))).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
