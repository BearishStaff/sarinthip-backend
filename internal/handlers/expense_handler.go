package handlers

import (
	expensesvc "sarinthip-backend/internal/pkg/expense-service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseHandlers struct {
	expenseService expensesvc.IExpenseService
}

func NewExpenseHandlers(
	expenseService expensesvc.IExpenseService,
) *ExpenseHandlers {
	return &ExpenseHandlers{
		expenseService: expenseService,
	}
}

func (handler *ExpenseHandlers) GetExpenses(c *gin.Context) {

	branchID := c.Param("branch_id")
	branchIDUUID, err := uuid.Parse(branchID)
	if err != nil {
		c.JSON(400, err)
		return
	}

	expenses, err := handler.expenseService.GetExpenses(branchIDUUID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, expenses)
}
