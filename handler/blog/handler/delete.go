package handler

import (
	"blog-backend/global"
	"blog-backend/handler/blog/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestDelete struct {
	Id uint `json:"id" binding:"required"`
}

func Delete(ctx *gin.Context) {
	params := &RequestDelete{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.Delete(params.Id)
	if err != nil {
		global.Log.Error("handler.blog.delete error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
