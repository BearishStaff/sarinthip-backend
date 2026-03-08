package main

import (
	"sarinthip-backend/internal/database"

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
	// 4. Routes
	api := r.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
		// Branch Management
		// api.GET("/branches", handlers.GetBranches)

		// Expense & Bill Management
		// api.POST("/parse", handlers.ParseRawText)    // Returns JSON for preview
		// api.POST("/bills", handlers.CreateBillGroup) // Saves Bill + many Expenses

		// Income Management
		// api.POST("/income", handlers.CreateIncome) // Saves Income record
	}

	r.Run(":8080") // Railway/Render will detect this port
}
