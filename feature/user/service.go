package user

import (
	"clean-architect/repository"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	AllUser(c echo.Context) ([]UserRes, error)
	FindByID(c echo.Context, id string) (*UserRes, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(u repository.UserRepository) *userService {
	return &userService{
		userRepo: u,
	}
}

func (s *userService) FindByID(c echo.Context, id string) (*UserRes, error) {
	return nil, nil
}

func (s *userService) AllUser(c echo.Context) ([]UserRes, error) {

	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	res := make([]UserRes, 0)
	copier.Copy(&res, &users)

	return res, nil
}
