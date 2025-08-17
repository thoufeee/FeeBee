package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup
func Signup(c *gin.Context) {
	var data struct {
		Firstname  string `json:"firstname"`
		Secondname string `json:"secondname"`
		Email      string `json:"email" binding:"required"`
		Phone      string `json:"phone"`
		Password   string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill blanks"})
		return
	}
}
