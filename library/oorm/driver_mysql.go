package oorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	User    string `conf:"user,default=root,usage=mysql_user"`
	Pass    string `conf:"pass,default=123456,usage=mysql_pass"`
	Host    string `conf:"host,default=localhost,usage=mysql_host"`
	Port    int    `conf:"port,default=3306,usage=mysql_port"`
	DBName  string `conf:"dbname,default=blog,usage=mysql_dbname"`
	Charset string `conf:"charset,default=utf8mb4,usage=mysql_charset"`
}

func (c *MySQLConfig) Open() gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Pass, c.Host, c.Port, c.DBName, c.Charset)
	return mysql.Open(dsn)
}
