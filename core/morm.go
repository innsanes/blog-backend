package core

import (
	"blog-backend/library/oorm"
	"blog-backend/model/mymodel"
	"github.com/innsanes/serv"
)

type MOrm struct {
	*serv.Service
	conf   Confer
	config *oorm.MySQLConfig
	*oorm.OOrm
}

func NewMOrm(conf Confer) *MOrm {
	ret := &MOrm{
		conf: conf,
	}
	return ret
}

func (s *MOrm) BeforeServe() (err error) {
	s.config = &oorm.MySQLConfig{}
	s.conf.RegisterConfWithName("orm", s.config)
	return
}

func (s *MOrm) Serve() (err error) {
	s.OOrm = oorm.New(s.config)
	err = s.Open()
	if err != nil {
		return
	}
	err = s.AutoMigrate(mymodel.BuildList...)
	if err != nil {
		return
	}
	return
}
