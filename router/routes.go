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

	// branch
	api.GET("/getbranch", services.GetBranches)
	api.POST("/newbranch", services.NewBranch)
	api.PUT("/updatebranch/:id", services.UpdateBranch)
	api.DELETE("/deletebranch/:id", services.DeleteBranch)

	//  student
	api.GET("/getstudet", services.AllStudents)
	api.POST("/newstudent", services.AddNewStudent)
	api.PUT("/updatestudent/:id", services.UpdateStudent)
	api.DELETE("/deletestudent/:id", services.DeleteStudent)

	// payment
	api.POST("/addpayment", services.AddPayment)
	api.GET("/getpayment", services.Allpayment)

	// admin
	api.GET("/admindetails", services.AdminDetails)
}
