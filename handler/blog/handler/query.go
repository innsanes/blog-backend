package handler

import (
	"blog-backend/global"
	"blog-backend/handler/blog/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestGet struct {
	Id uint `json:"id" binding:"required"`
}

type ResponseGet struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
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
		Id:         blog.ID,
		Name:       blog.Name,
		Content:    blog.Content,
		CreateTime: blog.CreatedAt.UnixMilli(),
		UpdateTime: blog.UpdatedAt.UnixMilli(),
	})
}

type RequestList struct {
}

type ResponseList []ResponseListItem

type ResponseListItem struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func List(ctx *gin.Context) {
	params := &RequestList{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blogs, err := dao.List()
	if err != nil {
		global.Log.Error("handler.blog.list error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var res ResponseList
	for _, blog := range blogs {
		res = append(res, ResponseListItem{
			Id:         blog.ID,
			Name:       blog.Name,
			CreateTime: blog.CreatedAt.UnixMilli(),
			UpdateTime: blog.UpdatedAt.UnixMilli(),
		})
	}
	ctx.JSON(http.StatusOK, res)
}
