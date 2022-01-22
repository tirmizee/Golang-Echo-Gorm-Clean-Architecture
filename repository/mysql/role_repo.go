package mysql

import (
	"clean-architect/repository"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindById(id int) (*repository.Role, error) {
	return nil, nil
}

func (r *RoleRepository) FindAll() ([]repository.Role, error) {

	var roles []repository.Role

	if result := r.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
