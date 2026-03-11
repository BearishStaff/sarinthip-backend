package main

import (
	// "sarinthip-backend/internal/database"
	"sarinthip-backend/internal/handlers"
	branchsvc "sarinthip-backend/internal/pkg/branch-service"
	expensesvc "sarinthip-backend/internal/pkg/expense-service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init Database
	// database.Connect()

	// 2. Init Gin Router
	r := gin.Default()

	// 3. Setup CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // web
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	branchService := branchsvc.NewBranchService()
	expenseService := expensesvc.NewExpenseService()

	branchHandler := handlers.NewBranchHandlers(branchService)
	expenseHandler := handlers.NewExpenseHandlers(expenseService)

	// 4. Routes
	api := r.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// Branches
		api.GET("/branches", branchHandler.GetBranches)

		// Expense
		api.GET("/expenses/:branch_id", expenseHandler.GetExpenses)

		// Income

	}

	r.Run(":8080") // Railway/Render will detect this port
}
