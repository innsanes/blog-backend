package dao

import (
	"blog-backend/structs/model"

	"gorm.io/gorm"
)

var Image IImage = &ImageDao{}

type ImageDao struct{}

type IImage interface {
	Create(db *gorm.DB, m *model.Image) (err error)
	ListPage(db *gorm.DB, page int, size int) (t []*model.Image, err error)
}

func (s *ImageDao) Create(db *gorm.DB, m *model.Image) (err error) {
	err = db.Model(&model.Image{}).Create(m).Error
	return
}

func (s *ImageDao) ListPage(db *gorm.DB, page int, size int) (t []*model.Image, err error) {
	err = db.Model(&model.Image{}).Order("id desc").Offset(page * size).Limit(size).Find(&t).Error
	return
}
