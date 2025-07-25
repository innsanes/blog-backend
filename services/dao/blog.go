package dao

import (
	"blog-backend/structs/model"

	"gorm.io/gorm"
)

var Blog IBlog = &BlogDao{}

type BlogDao struct{}

type IBlog interface {
	Create(db *gorm.DB, m *model.Blog) (err error)
	UpdateTags(db *gorm.DB, m *model.Blog, tags []*model.Tag) (err error)
	Update(db *gorm.DB, m *model.Blog) (err error)
	Get(db *gorm.DB, id uint) (t *model.Blog, err error)
	Count(db *gorm.DB) (count int64, err error)
	CountWithTag(db *gorm.DB, tagId uint) (count int64, err error)
	ListPage(db *gorm.DB, page int, size int) (t []*model.Blog, err error)
	ListPageWithTag(db *gorm.DB, page int, size int, tagId uint) (t []*model.Blog, err error)
	ListCursorForward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error)
	ListCursorForwardWithTag(db *gorm.DB, cursor uint, size int, tagId uint) (t []*model.Blog, err error)
	ListCursorBackward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error)
	ListCursorBackwardWithTag(db *gorm.DB, cursor uint, size int, tagId uint) (t []*model.Blog, err error)
	Delete(db *gorm.DB, id uint) (err error)
}

func (s *BlogDao) Create(db *gorm.DB, m *model.Blog) (err error) {
	err = db.Model(&model.Blog{}).Create(m).Error
	return
}

func (s *BlogDao) UpdateTags(db *gorm.DB, m *model.Blog, tags []*model.Tag) (err error) {
	err = db.Model(m).Association("Tags").Replace(tags)
	return
}

func (s *BlogDao) Update(db *gorm.DB, m *model.Blog) (err error) {
	err = db.Model(&model.Blog{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func (s *BlogDao) Get(db *gorm.DB, id uint) (t *model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id = ?", id).Preload("Tags").First(&t).Error
	return
}

func (s *BlogDao) Count(db *gorm.DB) (count int64, err error) {
	err = db.Model(&model.Blog{}).Count(&count).Error
	return
}

func (s *BlogDao) CountWithTag(db *gorm.DB, tagId uint) (count int64, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagId).Count(&count).Error
	return
}

func (s *BlogDao) ListPage(db *gorm.DB, page int, size int) (t []*model.Blog, err error) {
	offset := page * size
	err = db.Model(&model.Blog{}).Offset(offset).Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) ListPageWithTag(db *gorm.DB, page int, size int, tagId uint) (t []*model.Blog, err error) {
	offset := page * size
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagId).
		Offset(offset).Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorForward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id > ?", cursor).Order("id").Limit(size).
		Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorForwardWithTag(db *gorm.DB, cursor uint, size int, tagId uint) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagId).Where("blogs.id > ?", cursor).
		Order("blogs.id ASC").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorBackward(db *gorm.DB, cursor uint, size int) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).Where("id < ?", cursor).Order("id desc").Limit(size).
		Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) ListCursorBackwardWithTag(db *gorm.DB, cursor uint, size int, tagId uint) (t []*model.Blog, err error) {
	err = db.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagId).Where("blogs.id < ?", cursor).
		Order("blogs.id DESC").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func (s *BlogDao) Delete(db *gorm.DB, id uint) (err error) {
	blog := model.Blog{Model: gorm.Model{ID: id}}
	err = db.Model(&model.Blog{}).Where("id = ?", id).Delete(&blog).Error
	return
}
