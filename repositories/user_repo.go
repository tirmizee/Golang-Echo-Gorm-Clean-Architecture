package repositories

import "clean-architect/repositories/models"

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindAll() ([]models.User, error)
}
