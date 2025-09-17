package router

import (
	"feebee/controllers"
	"feebee/services"

	"github.com/gin-gonic/gin"
)

// Public routes
func Routes(c *gin.Engine) {
	api := c.Group("/")
	api.POST("signup", controllers.Signup)
	api.POST("login", controllers.Login)
}

// admin route
func AdminRoute(c *gin.Engine) {
	api := c.Group("/admin")
	api.GET("/getbranch", services.GetBranches)
	api.POST("/newbranch", services.NewBranch)
	api.PUT("/updatebranch/:id", services.UpdateBranch)
	api.DELETE("/deletebranch/:id", services.DeleteBranch)

	api.POST("/addpayment", services.AddPayment)
	api.GET("/getpayment", services.Allpayment)
}
