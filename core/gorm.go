package core

import (
	"blogServer/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() error {
	c := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
	return err
}
