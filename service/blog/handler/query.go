package handler

import (
	"blog-backend/data/model"
	g "blog-backend/global"
	"blog-backend/service/blog/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestGet struct {
	Id uint
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
	UseCursor bool `form:"useCursor"`               // 使用游标
	Page      int  `form:"page"`                    // [分页]: 第几页
	Size      int  `form:"size" binding:"required"` // [分页]/[游标]: 每页大小
	Cursor    uint `form:"cursor"`                  // [游标]: Blog的ID
	Forward   bool `form:"forward"`                 // [游标]: 是向前还是向后
}

type ResponseList struct {
	Data  []ResponseListItem `json:"data"`
	Count int64              `json:"count"`
}

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
	var (
		blogs []model.Blog
		err   error
	)
	count, err := dao.Count()
	if err != nil {
		g.Log.Error("[Blog.ListCursorForward] error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if params.UseCursor {
		if params.Forward {
			blogs, err = dao.ListCursorForward(params.Cursor, params.Size)
			if err != nil {
				g.Log.Error("[Blog.ListCursorForward] error: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			blogs, err = dao.ListCursorBackward(params.Cursor, params.Size)
			if err != nil {
				g.Log.Error("[Blog.ListCursorBackward] error: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	} else {
		blogs, err = dao.ListPage(params.Page, params.Size)
		if err != nil {
			g.Log.Error("[Blog.ListPage] error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	resp := ResponseList{
		Data:  make([]ResponseListItem, 0, len(blogs)),
		Count: count,
	}
	for _, blog := range blogs {
		resp.Data = append(resp.Data, ResponseListItem{
			Id:         blog.ID,
			Name:       blog.Name,
			CreateTime: blog.CreatedAt.UnixMilli(),
			UpdateTime: blog.UpdatedAt.UnixMilli(),
		})
	}
	ctx.JSON(http.StatusOK, resp)

}
