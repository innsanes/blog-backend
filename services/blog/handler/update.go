package handler

import (
	g "blog-backend/global"
	"blog-backend/services/blog/dao"
	"blog-backend/structs/model"
	"blog-backend/structs/req"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.ParseUint(idString, 10, strconv.IntSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params := &req.BlogUpdate{}
	if err = ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = dao.Update(&model.Blog{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:    params.Name,
		Content: params.Content,
	}, params.Tags)
	if err != nil {
		g.Log.Error("handler.blog.update error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
