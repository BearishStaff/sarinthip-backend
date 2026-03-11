package handlers

import (
	branchsvc "sarinthip-backend/internal/pkg/branch-service"

	"github.com/gin-gonic/gin"
)

type BranchHandlers struct {
	branchService branchsvc.IBranchService
}

func NewBranchHandlers(
	branchService branchsvc.IBranchService,
) *BranchHandlers {
	return &BranchHandlers{
		branchService: branchService,
	}
}

// GetBranches fetches all branches for the dropdowns
func (handler *BranchHandlers) GetBranches(c *gin.Context) {
	branches, err := handler.branchService.GetBranches()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, branches)
}
