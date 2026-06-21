package services

import (
	"cdn-api/internal/dto"
	"cdn-api/internal/models"
	"cdn-api/internal/repositories"
	"context"
	"errors"
)

type UserService interface {
	RegisterUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// actual business logic is here

	existing, _ := s.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password, // of course this would be hashed on an actual deployment
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
