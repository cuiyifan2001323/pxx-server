// @Author 齐静春
// @Date 2023-11-14 21:41:00

package dao

import (
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}
type User struct {
	Id int64
	// 密码
	Password string
	// 手机号
	Mobile      string
	UserName    string `gorm:"column:userName"`
	UpdatedTime int64  `gorm:"column:updateTime"`
	CreatedTime int64  `gorm:"column:createTime"`
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (d *UserDao) Insert(context context.Context, user *User) error {
	return d.db.WithContext(context).Table("t_user").Create(user).Error
}
