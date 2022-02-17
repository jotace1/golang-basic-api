package router

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	albumRouter := router.Group("album")

	albumRouter.GET("/", controllers.GetAlbuns)
	albumRouter.POST("/", controllers.AddAlbum)
	albumRouter.GET("/:id", controllers.FindById)
	albumRouter.PUT("/:id", controllers.UpdateAlbum)
	albumRouter.DELETE("/:id", controllers.DeleteAlbum)

	return router
}
