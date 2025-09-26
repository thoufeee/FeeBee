package services

import (
	"feebee/db"
	"feebee/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// admin details
func AdminDetails(c *gin.Context) {
	var data model.Admin

	if err := db.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to get admin data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": data})
}
