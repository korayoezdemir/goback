package service

import (
	"github.com/koray.oezdemir/go-back-end/internal/model"
	"github.com/koray.oezdemir/go-back-end/internal/repo"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUser(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) GetUser(id uint) (*model.User, error) {
	return s.userRepo.Read(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}


