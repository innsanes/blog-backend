package router

import (
	"blog-backend/global"
	"blog-backend/handler/image/service"
	"blog-backend/handler/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Gin.RegisterRouter(RegisterImage)
}

func RegisterImage(group *gin.RouterGroup) {
	imageGroup := group.Group("/image")

	imageGroup.GET(":name", service.Get)
	imageGroup.Use(middleware.TokenCheck()).POST("upload", service.Upload)
}
