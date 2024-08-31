package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type RequestUpload struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

func Upload(ctx *gin.Context) {
	params := &RequestUpload{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证文件类型（例如只允许图片类型）
	//fmt.Println("Content-Type", ctx.Request.Header.Get("Content-Type"))
	//if ctx.Request.Header.Get("Content-Type") != "image/png" {
	//	ctx.String(http.StatusBadRequest, "只允许上传png图片")
	//	return
	//}

	src, err := params.File.Open()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "打开文件失败")
		return
	}
	defer src.Close()

	// 读取文件数据, 计算 hash256 值
	hash := sha256.New()
	if _, err = io.Copy(hash, src); err != nil {
		ctx.String(http.StatusInternalServerError, "读取文件失败")
		return
	}
	hashValue := hex.EncodeToString(hash.Sum(nil))
	//ctx.JSON(http.StatusOK, gin.H{"hash": hashValue})

	// 获取文件名称后缀
	fileSuffix := filepath.Ext(params.File.Filename)

	// 保存文件到本地
	filename := hashValue + fileSuffix
	dst, err := os.Create(filepath.Join(config.Path, filename))

	if err != nil {
		ctx.String(http.StatusInternalServerError, "保存文件失败")
		return
	}
	defer dst.Close()

	src.Seek(0, 0)
	if _, err = io.Copy(dst, src); err != nil {
		ctx.String(http.StatusInternalServerError, "保存文件失败")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"url": filename})
}
