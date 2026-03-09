package repository

import (
	"sarinthip-backend/internal/database"
	"sarinthip-backend/internal/models"

	"github.com/google/uuid"
)

// GetAllBranches retrieves all branches sorted by name
func GetAllBranches() ([]models.Branch, error) {
	var branches []models.Branch
	// Use database.DB (our global GORM instance)
	err := database.DB.Order("name asc").Find(&branches).Error
	return branches, err
}

// GetBranchByID finds a single branch by its UUID
func GetBranchByID(id uuid.UUID) (models.Branch, error) {
	var branch models.Branch
	err := database.DB.First(&branch, "id = ?", id).Error
	return branch, err
}

// CreateBranch saves a new branch to Supabase
func CreateBranch(branch *models.Branch) error {
	// Ensure the ID is generated if not already present
	if branch.ID == uuid.Nil {
		branch.ID = uuid.New()
	}
	return database.DB.Create(branch).Error
}

// UpdateBranch modifies an existing branch name
func UpdateBranch(branch *models.Branch) error {
	return database.DB.Save(branch).Error
}

// DeleteBranch removes a branch and triggers CASCADE delete for its bills
func DeleteBranch(id uuid.UUID) error {
	// We use an empty model struct to tell GORM which table to target
	return database.DB.Delete(&models.Branch{}, "id = ?", id).Error
}
