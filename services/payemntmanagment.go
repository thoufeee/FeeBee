package services

import (
	"feebee/db"
	"feebee/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// add new payment
func AddPayment(c *gin.Context) {
	var data struct {
		StudentID uint   `json:"student_id"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill balnks"})
		return
	}

	new := model.Payment{
		StudentID: data.StudentID,
		Amount:    data.Amount,
		Type:      data.Type,
	}

	if err := db.DB.Create(&new).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "new payment not created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "new payment created"})
}

// get all payments

func Allpayment(c *gin.Context) {
	var data []model.Payment

	if err := db.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to get data"})
		return
	}

	c.JSON(http.StatusOK, data)
}
