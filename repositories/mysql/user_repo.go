package mysql

import (
	"clean-architect/models"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id int) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {

	var users []models.User

	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	} else {
		fmt.Println(result.RowsAffected)
	}

	return users, nil
}
