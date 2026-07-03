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
	activitiesHandler := handlers.ActivitiesHandler(db)
	organizationsHandler := handlers.OrganizationsHandler(db)
	programsHandler := handlers.ProgramsHandler(db)
	subActivitiesHandler := handlers.SubActivitiesHandler(db)

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


		api.GET("/activity", activitiesHandler.GetActivities)
		api.POST("/activity", activitiesHandler.CreateActivity)
		api.GET("/activity/:id", activitiesHandler.GetActivityByID)
		api.PUT("/activity/:id", activitiesHandler.UpdateActivity)
		api.DELETE("/activity/:id", activitiesHandler.DeleteActivity)

		api.GET("/sub-activity", subActivitiesHandler.GetSubActivities)
		api.POST("/sub-activity", subActivitiesHandler.CreateSubActivity)
		api.GET("/sub-activity/:id", subActivitiesHandler.GetSubActivityByID)
		api.PUT("/sub-activity/:id", subActivitiesHandler.UpdateSubActivity)
		api.DELETE("/sub-activity/:id", subActivitiesHandler.DeleteSubActivity)

		api.GET("/organization", organizationsHandler.GetOrganizations)
		api.POST("/organization", organizationsHandler.CreateOrganization)
		api.GET("/organization/:id", organizationsHandler.GetOrganizationByID)
		api.PUT("/organization/:id", organizationsHandler.UpdateOrganization)
		api.DELETE("/organization/:id", organizationsHandler.DeleteOrganization)

		api.GET("/program", programsHandler.GetPrograms)
		api.POST("/program", programsHandler.CreateProgram)
		api.GET("/program/:id", programsHandler.GetProgramByID)
		api.PUT("/program/:id", programsHandler.UpdateProgram)
		api.DELETE("/program/:id", programsHandler.DeleteProgram)
	}
}
