package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"store/models"
	"store/utils"
)

func CreateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		_ = json.NewDecoder(r.Body).Decode(&product)

		if err := db.Create(&product).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}

func GetProducts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products []models.Product
		response, err := utils.Paginate(db, &products, r)
		if err != nil {
			http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func GetProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var product models.Product

		if err := db.First(&product, "id = ?", params["id"]).Error; err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}

func UpdateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var porduct models.Product
		result := db.First(&porduct, "id = ?", params["id"])
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Product not found", http.StatusNotFound)
			} else {
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			}
			return
		}
		_ = json.NewDecoder(r.Body).Decode(&porduct)
		db.Save(&porduct)
		json.NewEncoder(w).Encode(porduct)
	}
}

func DeleteProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var product models.Product
		result := db.Delete(&product, "id = ?", params["id"])
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Product deleted successfully")
	}
}
