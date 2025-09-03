package model

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Image{})
}

type Image struct {
	gorm.Model
	Name string `gorm:"column:name;type:VARCHAR(32)"`
	MD5  string `gorm:"column:md5;type:VARCHAR(32);uniqueIndex"`
}
