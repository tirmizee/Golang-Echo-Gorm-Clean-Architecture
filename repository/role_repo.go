package repository

import "clean-architect/repository/models"

type RoleRepository interface {
	FindById(id int) (*models.Role, error)
	FindAll() ([]models.Role, error)
}
