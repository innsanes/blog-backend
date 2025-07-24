package handler

import (
	g "blog-backend/global"
	"blog-backend/services/service"
	"blog-backend/structs/req"
	"blog-backend/structs/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Echo(ctx *gin.Context) {
	idString := ctx.Param("message")
	request := &req.Echo{
		Message: idString,
	}
	message, err := service.Echo.Echo(request)
	if err != nil {
		g.Log.Error("%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := &resp.Echo{Message: message}
	ctx.JSON(http.StatusOK, response)
}
