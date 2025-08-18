package model

import "gorm.io/gorm"

// branch
type Branch struct {
	gorm.Model
	AdminID    uint
	BranchName string `json:"branchname" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Photo      string `json:"photo"`
}
