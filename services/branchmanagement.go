package services

import (
	"feebee/db"
	"feebee/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// list branch
func GetBranches(c *gin.Context) {
	var data []model.Branch

	if err := db.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to get services"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// add branch

func NewBranch(c *gin.Context) {
	var Data struct {
		AdminId    uint   `json:"adminid"`
		BranchName string `json:"branchname" binding:"required"`
		Location   string `json:"location" binding:"required"`
		Photo      string `json:"photo"`
	}

	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill all blanks"})
		return
	}

	var admin model.Admin

	newbranch := &model.Branch{
		AdminID:    admin.ID,
		BranchName: Data.BranchName,
		Location:   Data.Location,
		Photo:      Data.Photo,
	}

	if err := db.DB.Create(&newbranch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "branch not created"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"res": "New Branch Created"})
}

// update branch

func UpdateBranch(c *gin.Context) {
	id := c.Param("id")

	branch_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "inavlid id"})
		return
	}

	var branch model.Branch

	if err := db.DB.First(&branch, branch_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "branch not found"})
		return
	}

	var data struct {
		BranchName *string `json:"branchname"`
		Location   *string `json:"location"`
		Photo      *string `json:"photo"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid input"})
		return
	}

	if data.BranchName != nil {
		branch.BranchName = *data.BranchName
	}

	if data.Location != nil {
		branch.Location = *data.Location
	}

	if data.Photo != nil {
		branch.Photo = *data.Photo
	}

	if err := db.DB.Save(&branch).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"res": "failed to update branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Successfuly updated Branch"})
}

// delete branch
func DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	branch_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid branch"})
		return
	}

	var branch model.Branch

	if err := db.DB.First(&branch, branch_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "branch not found"})
		return
	}

	if err := db.DB.Unscoped().Delete(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to delete branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Successfuly Deleted Branch"})

}
