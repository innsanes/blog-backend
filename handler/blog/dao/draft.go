package dao

import (
	"blog-backend/global"
	"blog-backend/model/mymodel"
)

func CreateDraft(m *mymodel.BlogDraft) (err error) {
	err = global.MySQL.Model(&mymodel.BlogDraft{}).Create(m).Error
	return
}

func UpdateDraft(m *mymodel.BlogDraft) (err error) {
	err = global.MySQL.Model(&mymodel.BlogDraft{}).Where("blog_id = ?", m.BlogID).Updates(m).Error
	return
}

func GetBlogDraft(blogId uint) (t *mymodel.BlogDraft, err error) {
	err = global.MySQL.Model(&mymodel.BlogDraft{}).Where("blog_id = ?", blogId).Find(&t).Error
	return
}

func GetDraft(id uint) (t *mymodel.BlogDraft, err error) {
	err = global.MySQL.Model(&mymodel.BlogDraft{}).Where("id = ?", id).Find(&t).Error
	return
}

func ListDraft() (t []mymodel.BlogDraft, err error) {
	err = global.MySQL.Model(&mymodel.BlogDraft{}).Select("name", "blog_id").Find(&t).Error
	return
}
