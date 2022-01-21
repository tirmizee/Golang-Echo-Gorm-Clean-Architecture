package mysql

import (
	"clean-architect/models"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindById(id int) (*models.Role, error) {
	return nil, nil
}

func (r *roleRepository) FindAll() ([]models.Role, error) {

	var roles []models.Role

	if result := r.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
