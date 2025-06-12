package dao

import (
	"blog-backend/data/model"
	"blog-backend/global"
	"gorm.io/gorm"
)

func Create(m *model.Blog) (err error) {
	err = global.MySQL.Model(&model.Blog{}).Create(m).Error
	return
}

func Update(m *model.Blog) (err error) {
	err = global.MySQL.Model(&model.Blog{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func Get(id uint) (t *model.Blog, err error) {
	err = global.MySQL.Model(&model.Blog{}).Where("id = ?", id).Find(&t).Error
	return
}

func List() (t []model.Blog, err error) {
	err = global.MySQL.Model(&model.Blog{}).Omit("content").Find(&t).Error
	return
}

func Delete(id uint) (err error) {
	blog := model.Blog{Model: gorm.Model{ID: id}}
	err = global.MySQL.Model(&model.Blog{}).Where("id = ?", id).Delete(&blog).Error
	return
}
