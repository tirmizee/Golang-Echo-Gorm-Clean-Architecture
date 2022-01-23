package user

import (
	"clean-architect/repository"
	repositoryredis "clean-architect/repository_redis"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	AllUser(c echo.Context) ([]UserRes, error)
	FindByID(c echo.Context, id string) (*UserRes, error)
}

type userService struct {
	redisRepo *repositoryredis.RedisRepository
	userRepo  repository.UserRepository
}

func NewUserService(r *repositoryredis.RedisRepository, u repository.UserRepository) *userService {
	return &userService{
		redisRepo: r,
		userRepo:  u,
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
	for _, user := range users {
		item := UserRes{
			Id:        user.ID,
			Username:  user.Username,
			Password:  user.Password,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		res = append(res, item)
	}

	return res, nil
}
