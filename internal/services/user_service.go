package services

import (
	"cdn-api/internal/models"
	"errors"
)

type UserService struct {
}

func (s *UserService) CreateUser(name string, age int) (*models.User, error) {
	if age < 18 {
		return nil, errors.New("User must be an adult ")
	}
	user := &models.User{ID: "123", Age: age, Name: name}
	return user, nil
}
