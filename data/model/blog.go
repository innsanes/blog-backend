package model

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Blog{})
}

type Blog struct {
	gorm.Model
	Name    string `gorm:"column:name;type:VARCHAR(30)"`
	Content string `gorm:"column:content;type:LONGTEXT"`
}
