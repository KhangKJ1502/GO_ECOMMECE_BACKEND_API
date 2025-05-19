package routers

import (
	"go_ecommerce_backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1/2025")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		// v1.PUT("/ping", pong)
	}

	user := router.Group("api")
	{
		user.GET("/user", controller.NewUserController().GetUserByID)
	}

	// v2 := router.Group("/v2/2025") //curl http://localhost:8002/v2/2025/ping2/Khang

	// {
	// 	v2.GET("/ping2/:name", bang)
	// }

	return router
}

func bang(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(200, gin.H{
		"Thong Bao": "Khago ng dep trai" + name,
		"uid":       uid,
	})
}
