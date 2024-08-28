package dao

import (
	"blog-backend/global"
	"blog-backend/model/mymodel"
)

func Create(m *mymodel.Blog) (err error) {
	err = global.MySQL.Create(m).Error
	return
}

func Update(m *mymodel.Blog) (err error) {
	err = global.MySQL.Save(m).Error
	return
}

func Get(id uint) (t *mymodel.Blog, err error) {
	err = global.MySQL.Where("id = ?", id).Find(&t).Error
	return
}

func List() (t []mymodel.Blog, err error) {
	err = global.MySQL.Select("name", "id").Find(&t).Error
	return
}
