package utils

import (
	"gorm.io/gorm"
	"math"
	"net/http"
	"os"
	"strconv"
)

// PaginationResponse is a generic struct for paginated responses
type PaginationResponse struct {
	Count       int         `json:"count"`
	NextOffset  int         `json:"next_offset"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	Results     interface{} `json:"results"`
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
}

// Paginate fetches paginated data and builds a structured response
func Paginate(db *gorm.DB, model interface{}, r *http.Request) (PaginationResponse, error) {
	api_limit := os.Getenv("API_LIMIT")
	defaultLimit, _ := strconv.Atoi(api_limit)

	// Parse limit and page from query parameters
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > defaultLimit {
		limit = defaultLimit
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	// Count total records
	var totalCount int64
	db.Model(model).Count(&totalCount)

	// Query with pagination
	if err := db.Limit(limit).Offset(offset).Find(model).Error; err != nil {
		return PaginationResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return PaginationResponse{
		Count:       int(totalCount),
		NextOffset:  offset + limit,
		CurrentPage: page,
		TotalPages:  totalPages,
		Results:     model,
		Success:     true,
		Message:     "Request processed successfully",
	}, nil
}
