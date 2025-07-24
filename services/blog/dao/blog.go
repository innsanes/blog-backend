package dao

import (
	g "blog-backend/global"
	"blog-backend/structs/model"
	"slices"

	"gorm.io/gorm"
)

func Create(m *model.Blog, tags []string) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txError error) {
		mTags, txError := findOrCreateTags(tx, tags)
		if txError != nil {
			return
		}
		m.Tags = nil
		txError = tx.Model(&model.Blog{}).Create(m).Error
		if txError != nil {
			return
		}
		txError = tx.Model(m).Association("Tags").Replace(mTags)
		if txError != nil {
			return
		}
		return
	})
	return
}

func findOrCreateTags(db *gorm.DB, tags []string) (mTags []model.Tag, err error) {
	mTags, err = findTags(db, tags)
	if err != nil {
		return
	}
	if len(tags)-len(mTags) <= 0 {
		return
	}
	cTags := make([]string, 0, len(tags)-len(mTags))
	for _, tag := range tags {
		isContain := slices.ContainsFunc(mTags, func(mTag model.Tag) bool {
			return mTag.Name == tag
		})
		if !isContain {
			cTags = append(cTags, tag)
		}
	}
	appendTags, err := createTags(db, cTags)
	if err != nil {
		return
	}
	mTags = append(mTags, appendTags...)
	return
}

func findTags(db *gorm.DB, tags []string) (mTags []model.Tag, err error) {
	mTags = make([]model.Tag, 0, len(tags))
	err = db.Model(&model.Tag{}).Where("name IN (?)", tags).Find(&mTags).Error
	return mTags, err
}

func createTags(db *gorm.DB, tags []string) (mTags []model.Tag, err error) {
	mTags = make([]model.Tag, 0, len(tags))
	for i := range tags {
		mTags = append(mTags, model.Tag{
			Name: tags[i],
		})
	}
	err = db.Model(&model.Tag{}).Create(&mTags).Error
	return mTags, err
}

func Update(m *model.Blog, tags []string) (err error) {
	err = g.MySQL.Transaction(func(tx *gorm.DB) (txError error) {
		mTags, txError := findOrCreateTags(tx, tags)
		if txError != nil {
			return
		}
		m.Tags = nil
		txError = tx.Model(&model.Blog{}).Where("id = ?", m.ID).Updates(m).Error
		if txError != nil {
			return
		}
		txError = tx.Model(m).Association("Tags").Replace(mTags)
		if txError != nil {
			return
		}
		return
	})
	return
}

func Get(id uint) (t *model.Blog, err error) {
	err = g.MySQL.Model(&model.Blog{}).Where("id = ?", id).Preload("Tags").Find(&t).Error
	return
}

func Count() (count int64, err error) {
	err = g.MySQL.Model(&model.Blog{}).Count(&count).Error
	return
}

func CountWithTag(tag string) (count int64, err error) {
	var tagModel model.Tag
	err = g.MySQL.Where("name = ?", tag).First(&tagModel).Error
	if err != nil {
		return 0, err
	}
	err = g.MySQL.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagModel.ID).Count(&count).Error
	return
}

func ListPage(page int, size int) (t []*model.Blog, err error) {
	offset := page * size
	err = g.MySQL.Model(&model.Blog{}).Offset(offset).Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func ListPageWithTag(page int, size int, tag string) (t []*model.Blog, err error) {
	var tagModel model.Tag
	err = g.MySQL.Where("name = ?", tag).First(&tagModel).Error
	if err != nil {
		return nil, err
	}

	offset := page * size
	err = g.MySQL.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagModel.ID).
		Offset(offset).Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func ListCursorForward(cursor uint, size int) (t []*model.Blog, err error) {
	err = g.MySQL.Model(&model.Blog{}).Where("id > ?", cursor).Order("id").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func ListCursorForwardWithTag(cursor uint, size int, tag string) (t []*model.Blog, err error) {
	var tagModel model.Tag
	err = g.MySQL.Where("name = ?", tag).First(&tagModel).Error
	if err != nil {
		return nil, err
	}

	err = g.MySQL.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagModel.ID).Where("blogs.id > ?", cursor).
		Order("blogs.id ASC").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func ListCursorBackward(cursor uint, size int) (t []*model.Blog, err error) {
	err = g.MySQL.Model(&model.Blog{}).Where("id < ?", cursor).Order("id desc").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func ListCursorBackwardWithTag(cursor uint, size int, tag string) (t []*model.Blog, err error) {
	var tagModel model.Tag
	err = g.MySQL.Where("name = ?", tag).First(&tagModel).Error
	if err != nil {
		return nil, err
	}

	err = g.MySQL.Model(&model.Blog{}).
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Where("blog_tags.tag_id = ?", tagModel.ID).Where("blogs.id < ?", cursor).
		Order("blogs.id DESC").Limit(size).Omit("content").Preload("Tags").Find(&t).Error
	return
}

func Delete(id uint) (err error) {
	blog := model.Blog{Model: gorm.Model{ID: id}}
	err = g.MySQL.Model(&model.Blog{}).Where("id = ?", id).Delete(&blog).Error
	return
}
