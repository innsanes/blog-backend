package mymodel

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &User{})
	BuildList = append(BuildList, &UserPassword{})
}

type User struct {
	gorm.Model
	Name string `gorm:"column:name;unique;type:VARCHAR(15)"`
}

type UserPassword struct {
	UserID   uint   `gorm:"column:user_id;primaryKey"`
	UserName string `gorm:"column:user_name;primaryKey;type:VARCHAR(15)"`
	Password string `gorm:"column:password;notnull;type:VARCHAR(20)"`
}
