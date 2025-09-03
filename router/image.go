package router

import (
	"blog-backend/services/handler"

	"github.com/gin-gonic/gin"
)

func RegisterImage(group *gin.RouterGroup) {
	imageGroup := group.Group("/image")
	imageGroup.GET("/:image_path", handler.ImageGet)
}

func RegisterImageAuth(group *gin.RouterGroup) {
	imageGroup := group.Group("/image")
	imageGroup.GET(":image_path", handler.ImageGet)
	imageGroup.GET("", handler.ImageList)
	imageGroup.POST("", handler.ImageCreate)
}
