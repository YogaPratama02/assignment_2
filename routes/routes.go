package routes

import (
	"BootcampHacktiv8/assignment_2/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	e := gin.New()
	e.Use(gin.Recovery())
	v1 := e.Group("/api/v1")
	{
		order := v1.Group("orders")
		{
			order.GET("/", controllers.IndexController)
			order.POST("/", controllers.CreateController)
			order.PUT("/:id", controllers.UpdateController)
			order.DELETE("/:id", controllers.DeleteController)
		}
	}
	e.Run(":8000")
}
