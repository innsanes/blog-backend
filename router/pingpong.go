package router

import "github.com/gin-gonic/gin"

func RegisterPingPong(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
