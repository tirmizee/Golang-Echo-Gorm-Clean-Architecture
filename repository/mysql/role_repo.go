package mysql

import (
	"clean-architect/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindById(id int) (*repository.Role, error) {
	return nil, nil
}

func (r *roleRepository) FindAll() ([]repository.Role, error) {

	var roles []repository.Role

	if result := r.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
