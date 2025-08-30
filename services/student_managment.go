package services

import (
	"feebee/db"
	"feebee/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// all student details
func AllStudents(c *gin.Context) {
	var data model.Student

	if err := db.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to get students data"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// add new student
func AddNewStudent(c *gin.Context) {
	var data struct {
		BranchID       uint   `json:"branch_id" binding:"required"`
		Admission_Date string `json:"admissiondate" binding:"required"`
		FirstName      string `json:"firstname" binding:"requried"`
		SecondName     string `json:"secondname" binding:"required"`
		Age            string `json:"age" binding:"required"`
		PhoneNumber_1  string `json:"phone1" binding:"required"`
		PhoneNumber_2  string `json:"phone2"`
		Gender         string `json:"gender" binding:"required"`
		Address        string `json:"address" binding:"required"`
		GuardianName   string `json:"guardianname" binding:"required"`
		Grade          string `json:"grade" binding:"required"`
		Blood_Group    string `json:"bloodgroup"`
		Photo          string `json:"photo"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill all blanks"})
		return
	}

	new_student := &model.Student{
		BranchID:       data.BranchID,
		Admission_Date: data.Admission_Date,
		FirstName:      data.FirstName,
		SecondName:     data.SecondName,
		Age:            data.Age,
		PhoneNumber_1:  data.PhoneNumber_1,
		PhoneNumber_2:  data.PhoneNumber_2,
		Gender:         data.Gender,
		Address:        data.Address,
		GuardianName:   data.GuardianName,
		Grade:          data.Grade,
		Blood_Group:    data.Blood_Group,
		Photo:          data.Photo,
	}

	if err := db.DB.Create(&new_student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to add new student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"res": "Sucsessfuly added new student"})
}

// update existing student
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	student_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "inavlid id"})
		return
	}

	var student model.Student

	if err := db.DB.Find(&student, student_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "student not found"})
		return
	}
}
