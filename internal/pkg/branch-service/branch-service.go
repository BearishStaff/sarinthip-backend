package branchsvc

import (
	"sarinthip-backend/internal/models"

	"github.com/google/uuid"
)

type IBranchService interface {
	GetBranches() ([]models.Branch, error)
}

type BranchService struct{}

func NewBranchService() *BranchService {
	return &BranchService{}
}

func (service *BranchService) GetBranches() ([]models.Branch, error) {
	var branches []models.Branch
	// if err := database.DB.Order("name asc").Find(&branches).Error; err != nil {
	// 	return []models.Branch{}, err
	// }

	branches = []models.Branch{
		{ID: uuid.Must(uuid.Parse("9ec944d0-3120-43d3-82e8-bb7adfad6622")), Name: "ทดสอบ 1"},
		{ID: uuid.Must(uuid.NewUUID()), Name: "ทดสอบ 2"},
		{ID: uuid.Must(uuid.NewUUID()), Name: "ทดสอบ 3"},
	}

	return branches, nil
}
