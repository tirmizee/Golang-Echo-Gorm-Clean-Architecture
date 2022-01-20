package repositories

import "clean-architect/models"

type RoleRepository interface {
	FindById(id int) (*models.Role, error)
	FindAll() ([]models.Role, error)
}
