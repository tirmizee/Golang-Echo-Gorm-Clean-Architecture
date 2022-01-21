package role

import "clean-architect/repositories"

type RoleService interface {
	AllRole() ([]RoleRes, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(r repositories.RoleRepository) *roleService {
	return &roleService{
		roleRepo: r,
	}
}

func (s *roleService) AllRole() ([]RoleRes, error) {
	return nil, nil
}
