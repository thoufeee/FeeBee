package model

import "gorm.io/gorm"

// payment
type Payment struct {
	gorm.Model
	StudentID uint
	Amount    string `json:"amount" binding:"required"`
	Type      string `json:"type" binding:"required"`
}
