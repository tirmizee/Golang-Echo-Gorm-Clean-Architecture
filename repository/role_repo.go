package repository

type RoleRepository interface {
	FindById(id string) (*Role, error)
	FindAll() ([]Role, error)
}
