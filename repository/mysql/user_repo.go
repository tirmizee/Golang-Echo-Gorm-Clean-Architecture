package mysql

import (
	"clean-architect/repository"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id int) (*repository.User, error) {
	return nil, nil
}

func (r *userRepository) FindAll() ([]repository.User, error) {

	var users []repository.User

	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	} else {
		fmt.Println(result.RowsAffected)
	}

	return users, nil
}
