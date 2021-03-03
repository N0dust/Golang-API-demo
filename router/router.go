package router

import (
	"myapp/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter is ...
func SetupRouter() *gin.Engine {
	r := gin.Default()
	controller.GetClient()
	controller.ConnectDB()
	Group1 := r.Group("/user")
	{
		Group1.GET("/", controller.GetUser)
		Group1.POST("/", controller.CreateUser)
		Group1.PUT("/", controller.UpdateUser)
		Group1.DELETE("/", controller.DeleteUser)
	}

	Group2 := r.Group("/group")
	{
		Group2.GET("/", controller.GetGroup)
		Group2.POST("/", controller.CreateGroup)
		Group2.DELETE("/", controller.DeleteGroup)
	}

	return r
}
