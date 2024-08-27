package oorm

import "gorm.io/gorm"

type Driver interface {
	Open() gorm.Dialector
}
