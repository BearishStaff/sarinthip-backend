package main

import (
	"sarinthip-backend/internal/database"
	"sarinthip-backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init Database
	database.Connect()

	// 2. Init Gin Router
	r := gin.Default()

	// 3. Setup CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://https://sarinthip-web.vercel.app"}, // web
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	branchHandler := handlers.NewBranchHandlers()
	// billHandler := handlers.NewBillHandlers()
	// incomeHandler := handlers.NewIncomeHandlers()

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
		// api.POST("/branches", branchHandler.CreateBranch)
		// api.DELETE("/branches/:id", branchHandler.DeleteBranch) // e.g., /api/branches/uuid-string-here

		// Bills & Expenses
		// api.POST("/bills", billHandler.CreateBillGroup)

		// Income
		// api.POST("/income", incomeHandler.CreateIncome)
	}

	r.Run(":8080") // Railway/Render will detect this port
}
