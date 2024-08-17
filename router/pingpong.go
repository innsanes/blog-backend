package router

import (
	"blog-backend/global"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Gin.RegisterRouter(RegisterPingPong)
}

func RegisterPingPong(group *gin.RouterGroup) {
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
