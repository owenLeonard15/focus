package repository

import (
	"focus/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Function to get all users
func (repo *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}

// Function to get a user by ID
func (repo *UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := repo.db.First(&user, id).Error
	return user, err
}
