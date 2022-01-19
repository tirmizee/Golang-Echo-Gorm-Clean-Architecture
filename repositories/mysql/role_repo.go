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
