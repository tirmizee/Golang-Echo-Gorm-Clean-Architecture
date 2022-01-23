package mysql

import (
	"clean-architect/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*repository.User, error) {

	// var user repository.User

	return nil, nil
}

func (r *UserRepository) FindAll() ([]repository.User, error) {

	var users []repository.User

	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
