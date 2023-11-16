// @Author 齐静春
// @Date 2023-11-14 21:46:00

package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:cyf2001323@tcp(127.0.0.1:8001)/pxx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}
