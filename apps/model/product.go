package model

import "time"

// Product represents a product entity
type Product struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
