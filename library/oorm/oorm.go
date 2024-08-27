package oorm

import (
	"gorm.io/gorm"
)

type OOrm struct {
	driver Driver
	*gorm.DB
}

func New(driver Driver) *OOrm {
	ret := &OOrm{
		driver: driver,
	}
	return ret
}

func (o *OOrm) Open() (err error) {
	db, err := gorm.Open(o.driver.Open(), &gorm.Config{})
	if err != nil {
		return err
	}
	o.DB = db
	return
}
