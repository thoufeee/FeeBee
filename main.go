package main

import (
	"feebee/db"
	"feebee/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// database
	db.Connect()

	r := gin.Default()

	router.Routes(r)
	router.AdminRoute(r)

	r.Run(":8080")
}
