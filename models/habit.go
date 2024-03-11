package models

import (
	"time"
)

type HabitCompletion struct {
	ID      int       `json:"id"`
	HabitID int       `json:"habitId"`
	Date    time.Time `json:"date"`
}

type Habit struct {
	ID          int               `json:"id"`
	UserID      string            `json:"userId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	StartDate   time.Time         `json:"startDate"`
	EndDate     time.Time         `json:"endDate"`
	Completions []HabitCompletion `gorm:"foreignKey:HabitID"`
}
