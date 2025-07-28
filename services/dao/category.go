package dao

import (
	"blog-backend/structs/model"
	"blog-backend/structs/to"

	"gorm.io/gorm"
)

var Category ICategory = &CategoryDao{}

type CategoryDao struct{}

type ICategory interface {
	GetByName(db *gorm.DB, name string) (t *model.Category, err error)
	ListByName(db *gorm.DB, categories []string) (mCategories []*model.Category, err error)
	CreateMulti(db *gorm.DB, categories []string) (mCategories []*model.Category, err error)
}

func (s *CategoryDao) GetByName(db *gorm.DB, name string) (t *model.Category, err error) {
	err = db.Model(&model.Category{}).Where("name = ?", name).First(&t).Error
	return
}

func (s *CategoryDao) ListByName(db *gorm.DB, categories []string) (mCategories []*model.Category, err error) {
	mCategories = make([]*model.Category, 0, len(categories))
	err = db.Model(&model.Category{}).Where("name IN (?)", categories).Find(&mCategories).Error
	return mCategories, err
}

func (s *CategoryDao) CreateMulti(db *gorm.DB, categories []string) (mCategories []*model.Category, err error) {
	mCategories = to.Slice(categories, func(elem string) *model.Category {
		return &model.Category{Name: elem}
	})
	err = db.Model(&model.Category{}).Create(&mCategories).Error
	return mCategories, err
}
