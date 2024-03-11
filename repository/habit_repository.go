package repository

import (
	"focus/models"

	"gorm.io/gorm"
)

type HabitRepository struct {
	db *gorm.DB
}

func NewHabitRepository(db *gorm.DB) *HabitRepository {
	return &HabitRepository{db: db}
}

// Function to get all habits
func (repo *HabitRepository) GetHabits() ([]models.Habit, error) {
	var habits []models.Habit
	err := repo.db.Find(&habits).Error
	return habits, err
}

// Function to add a new habit
func (repo *HabitRepository) CreateHabit(habit models.Habit) (models.Habit, error) {
	err := repo.db.Create(&habit).Error
	return habit, err
}

// Function to add a new habit completion
func (repo *HabitRepository) AddCompletion(habitCompletion models.HabitCompletion) (models.HabitCompletion, error) {
	err := repo.db.Create(&habitCompletion).Error
	return habitCompletion, err
}
