package handler

import (
	"blog-backend/global"
	"blog-backend/service/blog/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = dao.Delete(uint(id))
	if err != nil {
		g.Log.Error("handler.blog.delete error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
