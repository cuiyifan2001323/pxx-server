// @Author 齐静春
// @Date 2023-11-14 20:51:00
// @Desc
package service

import (
	"context"
	"pxx-server/domain"
	"pxx-server/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) Signup(ctx context.Context, user *domain.User) error {
	return s.repo.Create(ctx, user)
}
