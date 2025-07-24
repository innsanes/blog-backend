package dao

import (
	"blog-backend/structs/model"
	"blog-backend/structs/to"

	"gorm.io/gorm"
)

var Tag ITag = &TagDao{}

type TagDao struct{}

type ITag interface {
	GetByName(db *gorm.DB, name string) (t *model.Tag, err error)
	ListByName(db *gorm.DB, tags []string) (mTags []*model.Tag, err error)
	CreateMulti(db *gorm.DB, tags []string) (mTags []*model.Tag, err error)
}

func (s *TagDao) GetByName(db *gorm.DB, name string) (t *model.Tag, err error) {
	err = db.Model(&model.Tag{}).Where("name = ?", name).First(&t).Error
	return
}

func (s *TagDao) ListByName(db *gorm.DB, tags []string) (mTags []*model.Tag, err error) {
	mTags = make([]*model.Tag, 0, len(tags))
	err = db.Model(&model.Tag{}).Where("name IN (?)", tags).Find(&mTags).Error
	return mTags, err
}

func (s *TagDao) CreateMulti(db *gorm.DB, tags []string) (mTags []*model.Tag, err error) {
	mTags = to.Slice(tags, func(elem string) *model.Tag {
		return &model.Tag{Name: elem}
	})
	err = db.Model(&model.Tag{}).Create(&mTags).Error
	return mTags, err
}
