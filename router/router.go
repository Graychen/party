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
	apiVersionOne.GET("activities", api.Activities)
	apiVersionOne.GET("activities/:id", api.Activity)
	apiVersionOne.POST("activities", api.CreateActivity)
	apiVersionOne.GET("images", api.Images)
	apiVersionOne.POST("images", api.CreateImages)
	apiVersionOne.DELETE("images", api.DeleteImages)
	apiVersionOne.GET("theoies", api.Theoies)
	apiVersionOne.GET("theoies/:id", api.Theoy)
	apiVersionOne.POST("theoies", api.CreateTheoy)
	apiVersionOne.GET("banners", api.Banners)
	apiVersionOne.GET("miens", api.Miens)
	apiVersionOne.GET("miens/:id", api.Mien)
	apiVersionOne.POST("miens", api.CreateMien)
	apiVersionOne.GET("applies", api.Applies)
	apiVersionOne.POST("applies", api.CreateApply)
	apiVersionOne.GET("dynamics", api.Dynamics)
	apiVersionOne.GET("dynamics/:id", api.Dynamic)
	apiVersionOne.POST("dynamics", api.CreateDynamic)
	apiVersionOne.GET("days", api.Days)
	apiVersionOne.GET("days/:id", api.Day)
	apiVersionOne.POST("days", api.CreateDay)
	return router
}
