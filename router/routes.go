package router

import (
	"feebee/controllers"

	"github.com/gin-gonic/gin"
)

// routes
func Routes(c *gin.Engine) {
	api := c.Group("/")
	api.POST("signup", controllers.Signup)
}
