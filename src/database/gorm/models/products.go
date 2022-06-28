package models

import "time"

type Product struct {
	ProductId   uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type Products []Product
