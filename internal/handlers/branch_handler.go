package handlers

import (
	"net/http"
	"sarinthip-backend/internal/database"
	"sarinthip-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BranchHandlers struct {
}

func NewBranchHandlers() *BranchHandlers {
	return &BranchHandlers{}
}

// GetBranches fetches all branches for the dropdowns
func (handler *BranchHandlers) GetBranches(c *gin.Context) {
	var branches []models.Branch
	if err := database.DB.Order("name asc").Find(&branches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches"})
		return
	}
	c.JSON(http.StatusOK, branches)
}

// CreateBranch adds a new branch
func (handler *BranchHandlers) CreateBranch(c *gin.Context) {
	var branch models.Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	branch.ID = uuid.New() // Generate new UUID
	if err := database.DB.Create(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create branch"})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

// DeleteBranch removes a branch by its UUID
func (handler *BranchHandlers) DeleteBranch(c *gin.Context) {
	idStr := c.Param("id")
	branchID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID format"})
		return
	}

	// Note: Because of our DB schema 'ON DELETE CASCADE',
	// deleting a branch will also delete its bills, expenses, and income.
	result := database.DB.Delete(&models.Branch{}, branchID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete branch"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
