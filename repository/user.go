// @Author 齐静春
// @Date 2023-11-14 20:51:00
// @Desc
package repository

import (
	"context"
	"pxx-server/domain"
	"pxx-server/repository/dao"
	"strconv"
	"time"
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(userDao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: userDao,
	}
}

func (r *UserRepository) Create(context context.Context, user *domain.User) error {
	return r.dao.Insert(context, &dao.User{
		Password:    user.Password,
		Mobile:      strconv.Itoa(user.Mobile),
		CreatedTime: time.Now().Unix(),
		UserName:    "测试",
	})
}
