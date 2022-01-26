package user

import (
	commons "clean-architect/commons/error"
	"clean-architect/repository"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	user, err := s.userRepo.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.NewCustomError("ERR001")
	}

	if err != nil {
		return nil, err
	}

	// manual map dto
	var res UserRes
	copier.Copy(&res, &user)

	return &res, nil
}

func (s *userService) AllUser(c echo.Context) ([]UserRes, error) {

	users, err := s.userRepo.FindAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.NewCustomError("ERR001")
	}

	if err != nil {
		return nil, err
	}

	res := make([]UserRes, 0)
	copier.Copy(&res, &users)

	return res, nil
}
