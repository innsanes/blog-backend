package dao

import (
	"blog-backend/global"
	"blog-backend/model/mymodel"
)

func Create(m *mymodel.Blog) (err error) {
	err = global.MySQL.Model(&mymodel.Blog{}).Create(m).Error
	return
}

func Update(m *mymodel.Blog) (err error) {
	err = global.MySQL.Model(&mymodel.Blog{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func Get(id uint) (t *mymodel.Blog, err error) {
	err = global.MySQL.Model(&mymodel.Blog{}).Where("id = ?", id).Find(&t).Error
	return
}

func List() (t []mymodel.Blog, err error) {
	err = global.MySQL.Model(&mymodel.Blog{}).Select("name", "id").Find(&t).Error
	return
}
