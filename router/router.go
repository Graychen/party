package router

import (
	"danjian/api"
	_ "danjian/docs"
	"net/http"

	//"danjian/middleware/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/login", api.Login)
	router.StaticFS("/uploads", http.Dir("uploads"))
	apiVersionOne := router.Group("/api/v1/")

	//apiVersionOne.Use(jwt.Jwt())

	apiVersionOne.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code":    200,
			"message": "This works",
			"data":    nil,
		})
	})
	//apiVersionOne.GET("offices", api.Offices)
	apiVersionOne.GET("activities", api.Activities)
	apiVersionOne.POST("activities", api.CreateActivity)
	apiVersionOne.GET("images", api.Images)
	apiVersionOne.POST("images", api.CreateImages)
	apiVersionOne.GET("theoies", api.Theoies)
	apiVersionOne.GET("banners", api.Banners)
	apiVersionOne.GET("miens", api.Miens)
	apiVersionOne.POST("miens", api.CreateMien)
	return router
}
