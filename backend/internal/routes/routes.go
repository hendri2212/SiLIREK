package routes

import (
	"net/http"
	"silirek/internal/handlers"
	"silirek/internal/middlewares"

	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.Static("/uploads", "./uploads")

	var allowedOrigins []string
	if gin.Mode() == gin.ReleaseMode {
		allowedOrigins = []string{
			"https://silirek.sipaktusarah.com",
		}
	} else {
		allowedOrigins = []string{
			"http://localhost:5173",
		}
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNoContent)
	})

	userHandler := handlers.UsersHandler(db)
	positionsHandler := handlers.PositionsHandler(db)
	leadersHandler := handlers.LeadersHandler(db)
	activitiesHandler := handlers.ActivitiesHandler(db)
	quaterliesHandler := handlers.QuarterliesHandler(db)
	budgetsHandler := handlers.BudgetsHandler(db)

	api := router.Group("/api")
	{
		api.POST("/login", userHandler.LoginUser)
		api.POST("/logout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
		})

		api.Use(middlewares.AuthMiddleware())

		api.GET("/me", userHandler.Me)
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		api.GET("/positions", positionsHandler.GetPositions)
		api.GET("/leaders", leadersHandler.GetLeaders)
		api.GET("/quarterlies", quaterliesHandler.GetQuarterlies)

		api.GET("/budgets", budgetsHandler.GetBudgets)
		api.POST("/budgets", budgetsHandler.CreateBudget)
		api.GET("/budgets/:id", budgetsHandler.GetBudgetByID)
		api.PUT("/budgets/:id", budgetsHandler.UpdateBudget)
		api.GET("/budgets/user/:user_id", budgetsHandler.GetBudgetsByUserID)
		api.GET("/budgets/user/:user_id/quarterly/:quarterly_id", budgetsHandler.GetBudgetsByUserIDAndQuarterlyID)

		api.GET("/activity", activitiesHandler.GetActivities)
		api.POST("/activity", activitiesHandler.CreateActivity)
	}
}
