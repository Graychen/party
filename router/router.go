package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
		router := gin.New()

		router.GET("/login", Login)
		apiVersionOne := router.Group("/api/v1/")

		apiVersionOne.Use(jwt.Jwt())

		apiVersionOne.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"code": 200,
				"message": "This works",
				"data": nil,
			})
		})
		return router
}