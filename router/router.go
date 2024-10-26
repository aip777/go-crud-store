package router

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"store/auth"
	"store/handlers"
	"store/middleware"
)

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.Handle("/products", middleware.Authenticate(http.HandlerFunc(handlers.GetProducts(db)))).Methods("GET")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.GetProduct(db)))).Methods("GET")
	router.Handle("/products", middleware.Authenticate(http.HandlerFunc(handlers.CreateProduct(db)))).Methods("POST")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.UpdateProduct(db)))).Methods("PUT")
	router.Handle("/products/{id}", middleware.Authenticate(http.HandlerFunc(handlers.DeleteProduct(db)))).Methods("DELETE")
	return router
}
