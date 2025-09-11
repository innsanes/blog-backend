package handler

import (
	"blog-backend/services/service"
	"blog-backend/structs/req"
	"blog-backend/structs/tod"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// ImageCreate 上传图片
func ImageCreate(ctx *gin.Context) {
	params := &req.ImageCreate{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	md5, err := service.Image.Create(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"md5": md5})
}

// ImageGet 获取图片
func ImageGet(ctx *gin.Context) {
	// 从URL路径中提取MD5
	imagePath := ctx.Param("image_path")

	// 安全检查
	md5Regex := regexp.MustCompile(`^[a-f0-9]{32}\.[a-zA-Z0-9]{1,6}$`)
	if !md5Regex.MatchString(imagePath) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid file format, expected: 32-char-md5.ext"})
		return
	}

	request := &req.ImageGet{
		Path: imagePath,
	}

	imageData, err := service.Image.Get(request)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// 设置响应头
	ctx.Header("Content-Type", "image/webp")
	ctx.Header("Cache-Control", "public, max-age=31536000") // 缓存1年
	ctx.Data(http.StatusOK, "image/webp", imageData)
}

func ImageList(ctx *gin.Context) {
	params := &req.ImageList{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	images, err := service.Image.ListPage(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := tod.ImageList(images)
	ctx.JSON(http.StatusOK, response)
}
