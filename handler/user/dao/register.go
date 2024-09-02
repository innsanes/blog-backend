package dao

import (
	"blog-backend/global"
	"blog-backend/model/mymodel"
)

func CreateUser(m *mymodel.User) (err error) {
	err = global.MySQL.Model(&mymodel.User{}).Create(m).Error
	return
}

func CreateUserPassword(userId uint, name, password string) (err error) {
	err = global.MySQL.Model(&mymodel.UserPassword{}).Create(&mymodel.UserPassword{
		UserID:   userId,
		UserName: name,
		Password: password,
	}).Error
	return
}
