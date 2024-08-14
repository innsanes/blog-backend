package router

import (
	"blog-backend/handler/image/service"
	"github.com/gin-gonic/gin"
)

func RegisterImage(engine *gin.Engine) {
	group := engine.Group("/image")

	//group.GET("", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "image",
	//	})
	//})
	group.POST("upload", service.Upload)
}
