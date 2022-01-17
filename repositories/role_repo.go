package repositories

import "gorm.io/gorm"

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindById(id int) (*Role, error) {
	return nil, nil
}
