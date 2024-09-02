package router

import (
	"blog-backend/global"
	"blog-backend/handler/user/service"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Gin.RegisterRouter(RegisterUser)
}

func RegisterUser(group *gin.RouterGroup) {
	group.POST("/register", service.Register)
	group.POST("/login", service.Login)
}
