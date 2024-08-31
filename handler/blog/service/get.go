package service

import (
	"blog-backend/handler/blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RequestGet struct {
	Id uint `json:"id" binding:"required"`
}

type ResponseGet struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func Get(ctx *gin.Context) {
	params := ctx.Param("id")
	if params == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog, err := dao.Get(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &ResponseGet{
		Name:    blog.Name,
		Content: blog.Content,
	})
	return
}

type RequestList struct {
}

type ResponseList []ResponseListItem

type ResponseListItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func List(ctx *gin.Context) {
	params := &RequestList{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogs, err := dao.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var res ResponseList
	for _, blog := range blogs {
		res = append(res, ResponseListItem{
			Id:   blog.ID,
			Name: blog.Name,
		})
	}
	ctx.JSON(http.StatusOK, res)
	return
}
