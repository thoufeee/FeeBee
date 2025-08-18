package model

import "gorm.io/gorm"

// admin
type Admin struct {
	gorm.Model
	Firstname       string `json:"firstname" binding:"required"`
	Secondname      string `json:"secondname" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required"`
	InstitutionName string `json:"institutionname" binding:"required"`
	Photo           string `json:"photo"`
}
