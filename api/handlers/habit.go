package handlers

import (
	"focus/models"
	"focus/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	repo *repository.HabitRepository
}

func NewHabitHandler(repo *repository.HabitRepository) *HabitHandler {
	return &HabitHandler{repo: repo}
}

// GetHabits handler
func (h *HabitHandler) GetHabits(c *gin.Context) {
	habits, err := h.repo.GetHabits()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, habits)
}

// // GetHabitByUser handler
// func GetHabitByUser(c *gin.Context) {
// 	// Get the user ID from the URL.
// 	userID := c.Param("userId")

// 	// Loop over the list of habits, looking for
// 	// habits belonging to the requested user.
// 	var userHabits []Habit
// 	for _, a := range habits {
// 		if a.UserID == userID {
// 			userHabits = append(userHabits, a)
// 		}
// 	}
// 	c.IndentedJSON(http.StatusOK, userHabits)
// }

// CreateHabit handler
func (h *HabitHandler) CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit, err := h.repo.CreateHabit(habit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, habit)
}

// AddCompletion handler
func (h *HabitHandler) AddCompletion(c *gin.Context) {
	var habitCompletion models.HabitCompletion
	if err := c.BindJSON(&habitCompletion); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habitCompletion, err := h.repo.AddCompletion(habitCompletion)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, habitCompletion)
}
