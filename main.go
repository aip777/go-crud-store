package main

import (
	"fmt"
	"log"
	"net/http"
	"store/models"
	"store/router"
	"store/utils"
)

func main() {
	// Database connection
	utils.Init()
	db := utils.InitDB()
	db.AutoMigrate(&models.Product{})

	// Set up the router
	router := router.InitRouter(db)

	// Start the server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
