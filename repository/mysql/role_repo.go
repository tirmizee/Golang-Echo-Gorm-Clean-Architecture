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

func (r *RoleRepository) FindById(id string) (*repository.Role, error) {

	var role repository.Role

	statement := r.db.Table("roles").Where("id = ?", id).First(&role)

	if err := statement.Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindAll() ([]repository.Role, error) {

	var roles []repository.Role

	if result := r.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
