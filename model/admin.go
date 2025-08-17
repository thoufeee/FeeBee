package model

// admin
type Admin struct {
	Firstname  string `json:"firstname"`
	Secondname string `json:"secondname"`
	Email      string `json:"email" binding:"required"`
	Phone      string `json:"phone"`
	Password   string `json:"password" binding:"required"`
}
