package dao

import (
	"blog-backend/structs/model"

	"gorm.io/gorm"
)

var Blog IBlog = &BlogDao{}

type BlogDao struct{}

type IBlog interface {
	Create(db *gorm.DB, m *model.Blog) (err error)
	UpdateCategories(db *gorm.DB, m *model.Blog, categories []*model.Category) (err error)
	Update(db *gorm.DB, m *model.Blog) (err error)
	Get(db *gorm.DB, id uint) (t *model.Blog, err error)
	Count(db *gorm.DB) (count int64, err error)
	CountWithCategory(db *gorm.DB, categoryId uint) (count int64, err error)
	ListPage(db *gorm.DB, page int, size int) (t []*model.Blog, err error)
	ListPageWithCategory(db *gorm.DB, page int, size int, categoryId uint) (t []*model.Blog, err error)
	ListCursorForward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error)
	ListCursorForwardWithCategory(db *gorm.DB, cursor uint, size int, categoryId uint) (t []*model.Blog, err error)
	ListCursorBackward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error)
	ListCursorBackwardWithCategory(db *gorm.DB, cursor uint, size int, categoryId uint) (t []*model.Blog, err error)
	Delete(db *gorm.DB, id uint) (err error)
}

func (s *BlogDao) Create(db *gorm.DB, m *model.Blog) (err error) {
	err = db.Model(&model.Blog{}).Create(m).Error
	return
}

func (s *BlogDao) UpdateCategories(db *gorm.DB, m *model.Blog, categories []*model.Category) (err error) {
	err = db.Model(m).Association("Categories").Replace(categories)
	return
}

func (s *BlogDao) Update(db *gorm.DB, m *model.Blog) (err error) {
	err = db.Model(&model.Blog{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func (s *BlogDao) Get(db *gorm.DB, id uint) (t *model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id = ?", id).Preload("View").Preload("Categories").First(&t).Error
	return
}

func (s *BlogDao) Count(db *gorm.DB) (count int64, err error) {
	err = db.Model(&model.Blog{}).Count(&count).Error
	return
}

func (s *BlogDao) CountWithCategory(db *gorm.DB, categoryId uint) (count int64, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Where("blog_categories.category_id = ?", categoryId).Count(&count).Error
	return
}

func (s *BlogDao) ListPage(db *gorm.DB, page int, size int) (t []*model.Blog, err error) {
	offset := page * size
	err = db.Model(&model.Blog{}).Offset(offset).Limit(size).Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) ListPageWithCategory(db *gorm.DB, page int, size int, categoryId uint) (t []*model.Blog, err error) {
	offset := page * size
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Where("blog_categories.category_id = ?", categoryId).
		Offset(offset).Limit(size).Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorForward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id > ?", cursor).Order("id").Limit(size).
		Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorForwardWithCategory(db *gorm.DB, cursor uint, size int, categoryId uint) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Where("blog_categories.category_id = ?", categoryId).Where("blogs.id > ?", cursor).
		Order("blogs.id ASC").Limit(size).Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorBackward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id < ?", cursor).Order("id desc").Limit(size).
		Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorBackwardWithCategory(db *gorm.DB, cursor uint, size int, categoryId uint) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Where("blog_categories.category_id = ?", categoryId).Where("blogs.id < ?", cursor).
		Order("blogs.id DESC").Limit(size).Omit("content").Preload("Categories").Find(&t).Error
	return
}

func (s *BlogDao) Delete(db *gorm.DB, id uint) (err error) {
	blog := model.Blog{Model: gorm.Model{ID: id}}
	err = db.Model(&model.Blog{}).Where("id = ?", id).Delete(&blog).Error
	return
}
