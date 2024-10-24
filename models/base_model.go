package models

import (
	"time"
)

// BaseModel
type BaseModel struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DataCreated time.Time `json:"data_created"`
	LastUpdated time.Time `json:"last_updated"`
	IsActive    bool      `json:"is_active"`
	JsonMeta    string    `json:"json_meta"`
	Type        string    `json:"type"`
	UUID        string    `json:"uuid"`
}
