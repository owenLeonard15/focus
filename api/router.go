package api

import (
	"focus/api/handlers"
	"focus/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// setup router with habitRepository
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())

	habitRepo := repository.NewHabitRepository(db)
	habitHandler := handlers.NewHabitHandler(habitRepo)

	// Routes
	api := router.Group("/api")
	{
		api.GET("/habits", habitHandler.GetHabits)
		api.POST("/habits", habitHandler.CreateHabit)
	}

	return router
}
