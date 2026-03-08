package handlers

import (
	"net/http"
	"sarinthip-backend/internal/database"
	"sarinthip-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillHandlers struct {
}

func NewBillHandlers() *BillHandlers {
	return &BillHandlers{}
}

func (handler *BillHandlers) CreateBillGroup(c *gin.Context) {
	var bill models.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// Start Database Transaction
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Create the Bill Header
		bill.ID = uuid.New()
		if err := tx.Create(&bill).Error; err != nil {
			return err // Returns error to rollback
		}

		// 2. Loop and Create Expenses linked to this Bill
		for i := range bill.Expenses {
			bill.Expenses[i].ID = uuid.New()
			bill.Expenses[i].BillID = bill.ID
			bill.Expenses[i].BranchID = bill.BranchID

			if err := tx.Create(&bill.Expenses[i]).Error; err != nil {
				return err // Returns error to rollback
			}
		}

		return nil // Commit transaction
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save bill group"})
		return
	}

	c.JSON(http.StatusCreated, bill)
}
