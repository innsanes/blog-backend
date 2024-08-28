package dao

import (
	"blog-backend/global"
	"blog-backend/model/mymodel"
)

func QueryUserPasswordByName(name string) (t *mymodel.UserPassword, err error) {
	err = global.MySQL.Where("user_name = ?", name).Find(&t).Error
	return
}
