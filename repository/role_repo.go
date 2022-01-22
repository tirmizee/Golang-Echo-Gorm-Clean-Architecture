package repository

type RoleRepository interface {
	FindById(id int) (*Role, error)
	FindAll() ([]Role, error)
}
