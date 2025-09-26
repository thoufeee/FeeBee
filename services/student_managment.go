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

// get student by id
func GetStudent(c *gin.Context) {
	id := c.Param("id")

	studet_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid id"})
		return
	}

	var data model.Student

	if err := db.DB.Find(&data, studet_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": data})
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

	var data struct {
		BranchID       *uint   `json:"branch_id"`
		Admission_Date *string `json:"admissiondate"`
		FirstName      *string `json:"firstname"`
		SecondName     *string `json:"secondname"`
		Age            *string `json:"age"`
		PhoneNumber_1  *string `json:"phone1"`
		PhoneNumber_2  *string `json:"phone2"`
		Gender         *string `json:"gender"`
		Address        *string `json:"address"`
		GuardianName   *string `json:"guardianname"`
		Grade          *string `json:"grade"`
		Blood_Group    *string `json:"bloodgroup"`
		Photo          *string `json:"photo"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid input"})
		return
	}

	if data.BranchID != nil {
		student.BranchID = *data.BranchID
	}

	if data.Admission_Date != nil {
		student.Admission_Date = *data.Admission_Date
	}

	if data.FirstName != nil {
		student.FirstName = *data.FirstName
	}

	if data.SecondName != nil {
		student.SecondName = *data.SecondName
	}

	if data.Age != nil {
		student.Age = *data.Age
	}

	if data.PhoneNumber_1 != nil {
		student.PhoneNumber_1 = *data.PhoneNumber_1
	}

	if data.PhoneNumber_2 != nil {
		student.PhoneNumber_2 = *data.PhoneNumber_2
	}

	if data.Gender != nil {
		student.Gender = *data.Age
	}

	if data.Address != nil {
		student.Address = *data.Address
	}

	if data.GuardianName != nil {
		student.GuardianName = *data.GuardianName
	}

	if data.Grade != nil {
		student.Grade = *data.Grade
	}

	if data.Blood_Group != nil {
		student.Blood_Group = *data.Blood_Group
	}

	if data.Photo != nil {
		student.Photo = *data.Photo
	}

	if err := db.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "Student profile updated"})
}

// delete student profile
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	student_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "inavlid id"})
		return
	}

	var student model.Student

	if err := db.DB.First(&student, student_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "student not found"})
		return
	}

	if err := db.DB.Unscoped().Delete(student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "successfuly deleted student"})
}
