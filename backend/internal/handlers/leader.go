package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LeaderHandler struct {
	db *gorm.DB
}

func LeadersHandler(db *gorm.DB) *LeaderHandler {
	db.AutoMigrate(&models.Leader{})
	return &LeaderHandler{db: db}
}
func (h *LeaderHandler) GetLeaders(c *gin.Context) {
	var leaders []models.Leader
	h.db.Find(&leaders)

	c.JSON(http.StatusOK, leaders)
}
func (h *LeaderHandler) CreateLeader(c *gin.Context) {
	var leader models.Leader
	if err := c.ShouldBindJSON(&leader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Create(&leader)

	c.JSON(http.StatusCreated, leader)
}
