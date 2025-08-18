package db

import (
	"feebee/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect database
func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database ", err)
	}

	err = db.AutoMigrate(
		&model.Admin{},
		&model.Branch{},
		&model.Payment{},
		&model.Student{},
	)

	if err != nil {
		log.Fatal("failed to automigrate")
	}

	DB = db
	fmt.Println("database connected successfuly")

}
