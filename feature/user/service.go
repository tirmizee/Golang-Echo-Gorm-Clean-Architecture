package user

import (
	log "clean-architect/commons/log"
	commons "clean-architect/commons/transfer"
	constants "clean-architect/constant"
	"clean-architect/repository"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService interface {
	AllUser(c echo.Context) ([]UserRes, error)
	FindByID(c echo.Context, id string) (*UserRes, error)
	CreateUser(c echo.Context, req *CreateUserReq) (*CreateUserRes, error)
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
		return nil, commons.NewCustomError(constants.ERR001)
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
		return nil, commons.NewCustomError(constants.ERR001)
	}

	if err != nil {
		return nil, err
	}

	res := make([]UserRes, 0)
	copier.Copy(&res, &users)

	return res, nil
}

func (s *userService) CreateUser(c echo.Context, req *CreateUserReq) (*CreateUserRes, error) {

	fmt.Println(req)

	var entity repository.User
	err := copier.Copy(&entity, req)
	if err != nil {
		log.InfoWithID(c, "step 1.1")
		return nil, err
	}

	persist, err := s.userRepo.Create(&entity)
	if err != nil {
		log.InfoWithID(c, "step 1.2")
		return nil, err
	}

	var res CreateUserRes
	err = copier.Copy(&res, persist)
	if err != nil {
		log.InfoWithID(c, "step 1.3")
		return nil, err
	}

	return &res, nil
}
