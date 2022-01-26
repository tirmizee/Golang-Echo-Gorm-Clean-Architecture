package role

import (
	commons "clean-architect/commons/error"
	constants "clean-architect/constant"
	"clean-architect/repository"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoleService interface {
	AllRole(c echo.Context) ([]RoleRes, error)
	FindByID(c echo.Context, id string) (*RoleRes, error)
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) *roleService {
	return &roleService{
		roleRepo: r,
	}
}

func (s *roleService) FindByID(c echo.Context, id string) (*RoleRes, error) {

	role, err := s.roleRepo.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.NewCustomError(constants.ERR001)
	}

	if err != nil {
		return nil, err
	}

	// manual map dto
	var res RoleRes
	copier.Copy(&res, &role)

	return &res, nil
}

func (s *roleService) AllRole(c echo.Context) ([]RoleRes, error) {

	roles, err := s.roleRepo.FindAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.NewCustomError(constants.ERR001)
	}

	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var res []RoleRes = make([]RoleRes, 0)

	// copy data to dto
	copier.Copy(&res, &roles)

	return res, nil
}
