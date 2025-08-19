package controllers

import (
	"feebee/db"
	"feebee/model"
	"feebee/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login
func Login(c *gin.Context) {
	var data struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"err": "fill all fileds"})
		return
	}

	var admin model.Admin

	if err := db.DB.Where("email = ?", data.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid email or password"})
		return
	}

	if !utlis.CheckPass(admin.Password, data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid email or pass"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Successfuly logedin"})
}
