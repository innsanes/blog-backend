package middleware

import (
	"blog-backend/global"
	"github.com/gin-gonic/gin"
)

func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if !global.Token.CheckToken(token) {
			c.JSON(401, gin.H{"error": "token error"})
			c.Abort()
			return
		}
		c.Next()
	}
}
