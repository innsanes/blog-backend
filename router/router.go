package router

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	RegisterImage(engine)
	RegisterPingPong(engine)
}
