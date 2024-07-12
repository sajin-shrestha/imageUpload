package models

import (
	"time"

	"gorm.io/gorm"
)

type Testimonial struct {
	// gorm.Model
	// Title       string `json:"title"`
	// Description string `json:"description"`
	// Image       string `json:"image"`

	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}
