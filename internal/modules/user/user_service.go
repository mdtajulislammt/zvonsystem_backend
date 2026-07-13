package user

import (
	"context"

	"github.com/mdtajulislammt/zvonsystem_backend/internal/model"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAll(ctx)
}
