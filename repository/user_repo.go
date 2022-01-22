package repository

import "clean-architect/repository/models"

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindAll() ([]models.User, error)
}
