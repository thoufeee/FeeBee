package controllers

import (
	"feebee/db"
	"feebee/model"
	"feebee/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup
func Signup(c *gin.Context) {
	var data struct {
		Firstname       string `json:"firstname" binding:"required"`
		Secondname      string `json:"secondname" binding:"required"`
		Email           string `json:"email" binding:"required"`
		Phone           string `json:"phone" binding:"required"`
		Password        string `json:"password" binding:"required"`
		InstitutionName string `json:"institutionname" binding:"required"`
		Photo           string `json:"photo"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill blanks"})
		return
	}

	var admin model.Admin

	// email check
	if !utlis.EmailCheck(data.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "email format not valid"})
		return
	}

	if err := db.DB.Where("email = ?", data.Email).First(&admin).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "email already taken"})
		return
	}

	// pass check
	if !utlis.PasswordStrength(data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "increase password strength"})
		return
	}

	// hashing
	hash, err := utlis.GenerateHash(data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to hash pass"})
		return
	}

	newadmin := &model.Admin{
		Firstname:       data.Firstname,
		Secondname:      data.Secondname,
		Email:           data.Email,
		Password:        hash,
		Phone:           data.Phone,
		InstitutionName: data.InstitutionName,
		Photo:           data.Photo,
	}

	if err := db.DB.Create(&newadmin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Successfuly signuped"})

}
