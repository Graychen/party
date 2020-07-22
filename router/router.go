package router

import (
	"danjian/api"
	_ "danjian/docs"
	"danjian/middleware/logger"
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
	router.StaticFS("/logs", http.Dir("logs"))
	apiVersionOne := router.Group("/api/v1/")
	//apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.Use(logger.LoggerToFile())

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
	apiVersionOne.PUT("activities/:id", api.UpdateActivity)
	apiVersionOne.DELETE("activities/:id", api.DeleteActivity)
	apiVersionOne.GET("images", api.Images)
	apiVersionOne.POST("images", api.CreateImages)
	apiVersionOne.DELETE("images", api.DeleteImages)
	apiVersionOne.GET("theoies", api.Theoies)
	apiVersionOne.GET("theoies/:id", api.Theoy)
	apiVersionOne.PUT("theoies/:id", api.UpdateTheoy)
	apiVersionOne.DELETE("theoies/:id", api.DeleteTheoy)
	apiVersionOne.POST("theoies", api.CreateTheoy)
	apiVersionOne.GET("banners", api.Banners)
	apiVersionOne.POST("banners", api.CreateBanner)
	apiVersionOne.DELETE("banners/:id", api.DeleteBanner)
	apiVersionOne.GET("miens", api.Miens)
	apiVersionOne.GET("miens/:id", api.Mien)
	apiVersionOne.POST("miens", api.CreateMien)
	apiVersionOne.PUT("miens/:id", api.UpdateMien)
	apiVersionOne.DELETE("miens/:id", api.DeleteMien)
	apiVersionOne.GET("applies", api.Applies)
	apiVersionOne.POST("applies", api.CreateApply)
	apiVersionOne.GET("dynamics", api.Dynamics)
	apiVersionOne.GET("dynamics/:id", api.Dynamic)
	apiVersionOne.POST("dynamics", api.CreateDynamic)
	apiVersionOne.PUT("dynamics/:id", api.UpdateDynamic)
	apiVersionOne.DELETE("dynamics/:id", api.DeleteDynamic)
	apiVersionOne.GET("days", api.Days)
	apiVersionOne.GET("days/:id", api.Day)
	apiVersionOne.PUT("days/:id", api.UpdateDay)
	apiVersionOne.DELETE("days/:id", api.DeleteDay)
	apiVersionOne.POST("days", api.CreateDay)
	return router
}
