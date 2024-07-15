package common

import (
	"blogServer/global"
	"blogServer/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := max(0, (option.Page-1)*option.Limit)
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err
}
