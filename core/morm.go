package core

import (
	"blog-backend/data/model"
	"fmt"
	"gorm.io/gorm/logger"

	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MOrm struct {
	*serv.Service
	*gorm.DB
	logger *MOrmLogger
	config *MySQLConfig
}

type MySQLConfig struct {
	User    string `conf:"user,default=root,usage=mysql_user"`
	Pass    string `conf:"pass,default=123456,usage=mysql_pass"`
	Host    string `conf:"host,default=localhost,usage=mysql_host"`
	Port    int    `conf:"port,default=3306,usage=mysql_port"`
	DBName  string `conf:"dbname,default=blog,usage=mysql_dbname"`
	Charset string `conf:"charset,default=utf8mb4,usage=mysql_charset"`
}

func NewMOrm() *MOrm {
	return &MOrm{}
}

func (s *MOrm) BeforeServe() (err error) {
	s.config = &MySQLConfig{}
	conf.RegisterConfWithName("orm", s.config)
	return
}

func (s *MOrm) Serve() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		s.config.User, s.config.Pass, s.config.Host, s.config.Port, s.config.DBName, s.config.Charset)

	s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewMOrmLogger(),
	})
	if err != nil {
		return err
	}
	s.Logger.LogMode(logger.Silent)
	err = s.AutoMigrate(model.BuildList...)
	if err != nil {
		return err
	}
	s.Logger.LogMode(logger.Info)
	return
}
