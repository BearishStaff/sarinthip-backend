package handlers

import (
	"sarinthip-backend/internal/database"
	"sarinthip-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IncomeHandlers struct {
}

func NewIncomeHandlers() *IncomeHandlers {
	return &IncomeHandlers{}
}

func (handler *IncomeHandlers) CreateIncome(c *gin.Context) {
	var income models.Income
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	income.ID = uuid.New()
	if err := database.DB.Create(&income).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save income"})
		return
	}

	c.JSON(201, income)
}
