package handler

import (
	"blog-backend/services/service"
	"blog-backend/structs/req"
	"blog-backend/structs/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Echo(ctx *gin.Context) {
	idString := ctx.Param("message")
	request := &req.Echo{
		Message: idString,
	}
	message, err := service.Echo.Echo(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := &resp.Echo{Message: message}
	ctx.JSON(http.StatusOK, response)
}
