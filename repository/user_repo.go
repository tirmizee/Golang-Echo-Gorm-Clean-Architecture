package repository

type UserRepository interface {
	FindById(id string) (*User, error)
	FindAll() ([]User, error)
	Create(user *User) (*User, error)
}
