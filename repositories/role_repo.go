package repositories

import "clean-architect/repositories/models"

type RoleRepository interface {
	FindById(id int) (*models.Role, error)
	FindAll() ([]models.Role, error)
}
