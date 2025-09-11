package service

import (
	"blog-backend/global"
	"blog-backend/services/dao"
	"blog-backend/structs/errc"
	"blog-backend/structs/model"
	"blog-backend/structs/req"
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"gorm.io/gorm"
)

var Image IImage = &ImageService{}

type ImageService struct{}

type IImage interface {
	Create(in *req.ImageCreate) (string, error)
	Get(in *req.ImageGet) ([]byte, error)
	ListPage(in *req.ImageList) (t []*model.Image, err error)
}

func (s *ImageService) Create(in *req.ImageCreate) (msg string, err error) {
	// 获取原始文件
	file, err := in.File.Open()
	if err != nil {
		return "", errc.Handle("[Image.Create] Open file", err)
	}
	defer file.Close()

	// 读取文件名称扩展名
	fileName := in.File.Filename
	ext := filepath.Ext(fileName)

	// 读取文件内容
	originData, err := io.ReadAll(file)
	if err != nil {
		return "", errc.Handle("[Image.Create] Read file", err)
	}

	// 计算origin版本的MD5
	originHasher := md5.New()
	originHasher.Write(originData)
	originMD5 := fmt.Sprintf("%x", originHasher.Sum(nil))

	// 生成compressed版本 (1920*1080, 质量90)
	maxWidth := global.Image.Config.MaxWidth
	maxHeight := global.Image.Config.MaxHeight
	quality := global.Image.Config.Quality
	compressedData, err := s.compress(originData, maxWidth, maxHeight, quality)
	if err != nil {
		return "", errc.Handle("[Image.Create] Convert to compressed", err)
	}

	// 开启事务
	err = global.MySQL.Transaction(func(tx *gorm.DB) (txErr error) {
		// 将name和MD5存入数据库
		imageModel := &model.Image{
			Name:     in.Name,
			FileName: fileName,
			MD5:      originMD5,
		}

		txErr = dao.Image.Create(tx, imageModel)
		if txErr = errc.Handle("[Image.Create] Create image record", txErr); txErr != nil {
			return
		}

		// 保存origin文件 (原始格式)
		originPath := filepath.Join(global.Image.Config.Path, "origin", originMD5+ext)
		txErr = s.saveFile(originData, originPath)
		if txErr = errc.Handle("[Image.Create] Save origin file", txErr); txErr != nil {
			return
		}

		// 保存compressed文件 (WebP格式)
		compressedPath := filepath.Join(global.Image.Config.Path, "compressed", originMD5+".webp")
		txErr = s.saveFile(compressedData, compressedPath)
		if txErr = errc.Handle("[Image.Create] Save compressed file", txErr); txErr != nil {
			return
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return originMD5, nil
}

func (s *ImageService) Get(in *req.ImageGet) (data []byte, err error) {
	// 直接从compressed目录获取图片
	// 路径格式: md5.webp
	filePath := filepath.Join(global.Image.Config.Path, "compressed", in.Path)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errc.Handle("[Image.Get] File not found", fmt.Errorf("image file not found: %s", filePath))
	}

	// 读取文件内容
	data, err = os.ReadFile(filePath)
	if err != nil {
		return nil, errc.Handle("[Image.Get] Read file", err)
	}

	return data, nil
}

// compress 将图片转换为压缩版本
func (s *ImageService) compress(data []byte, maxWidth int, maxHeight int, quality int) ([]byte, error) {
	// 解码图片
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 获取原始图片尺寸
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// 计算缩放比例，保持宽高比
	var newWidth, newHeight int

	// 如果图片尺寸在限制范围内，不需要缩放
	if originalWidth <= maxWidth && originalHeight <= maxHeight {
		newWidth = originalWidth
		newHeight = originalHeight
	} else {
		// 计算缩放比例，取较小的比例以确保图片完全适应限制尺寸
		widthRatio := float64(maxWidth) / float64(originalWidth)
		heightRatio := float64(maxHeight) / float64(originalHeight)
		ratio := widthRatio
		if heightRatio < widthRatio {
			ratio = heightRatio
		}

		newWidth = int(float64(originalWidth) * ratio)
		newHeight = int(float64(originalHeight) * ratio)
	}

	// 如果尺寸有变化，进行缩放
	var resizedImg image.Image = img
	if newWidth != originalWidth || newHeight != originalHeight {
		resizedImg = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
	}

	var buf bytes.Buffer
	options := &webp.Options{
		Lossless: true,
		Quality:  float32(quality),
	}
	if err = webp.Encode(&buf, resizedImg, options); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// saveFile 保存文件到指定路径
func (s *ImageService) saveFile(data []byte, filePath string) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write(data)
	return err
}

func (s *ImageService) ListPage(in *req.ImageList) (t []*model.Image, err error) {
	t, err = dao.Image.ListPage(global.MySQL.DB, in.Page, in.Size)
	if err = errc.Handle("[Image.ListPage] ListPage", err); err != nil {
		return
	}
	return
}
