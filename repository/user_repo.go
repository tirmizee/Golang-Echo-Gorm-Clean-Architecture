package repository

type UserRepository interface {
	FindById(id int) (*User, error)
	FindAll() ([]User, error)
}
