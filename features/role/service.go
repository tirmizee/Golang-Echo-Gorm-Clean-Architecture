package role

type Service interface {
	AllRole() *RoleRes
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) AllRole() *RoleRes {
	return nil
}
