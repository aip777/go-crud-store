package models

// Product defines the product model
type Product struct {
	BaseModel
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
