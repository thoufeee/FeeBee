package model

import "gorm.io/gorm"

// student
type Student struct {
	gorm.Model
	BranchID      uint
	FirstName     string `json:"firstname" binding:"requried"`
	SecondName    string `json:"secondname" binding:"required"`
	Age           string `json:"age"`
	PhoneNumber_1 string `json:"phone1" binding:"required"`
	PhoneNumber_2 string `json:"phone2" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
	Address       string `json:"address" binding:"required"`
	GuardianName  string `json:"guardianname" binding:"required"`
	Grade         string `json:"grade" binding:"required"`
	Blood_Group   string `json:"bloodgroup"`
	Photo         string `json:"photo"`
}
