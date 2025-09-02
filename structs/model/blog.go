package model

import "gorm.io/gorm"

func init() {
	BuildList = append(BuildList, &Blog{})
	BuildList = append(BuildList, &Category{})
}

type Blog struct {
	gorm.Model
	Name       string      `gorm:"column:name;type:VARCHAR(63)"`
	Summary    string      `gorm:"column:summary;type:VARCHAR(255)"`
	Content    string      `gorm:"column:content;type:LONGTEXT"`
	Categories []*Category `gorm:"many2many:blog_categories;"`
	View       View        `gorm:"polymorphicType:ViewerType;polymorphicId:ViewerID;polymorphicValue:blog"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;type:VARCHAR(30)"`
}
