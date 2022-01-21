package user

import (
	"clean-architect/models"
	"clean-architect/repositories"
)

type UserService interface {
	AllUser() ([]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) *userService {
	return &userService{
		userRepo: u,
	}
}

func (s *userService) AllUser() ([]models.User, error) {

	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
