package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivityHandler struct {
	db *gorm.DB
}

func ActivitiesHandler(db *gorm.DB) *ActivityHandler {
	db.AutoMigrate(&models.Activity{})
	return &ActivityHandler{db: db}
}
func (h *ActivityHandler) GetActivities(c *gin.Context) {
	var activity []models.Activity
	h.db.Find(&activity)

	c.JSON(http.StatusOK, activity)
}
func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Activity
	if err := h.db.Where("code = ? OR name = ?", activity.Code, activity.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode atau nama kegiatan sudah ada"})
		return
	}

	h.db.Create(&activity)
	c.JSON(http.StatusCreated, activity)
}

func (h *ActivityHandler) GetActivityByID(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	if err := h.db.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (h *ActivityHandler) UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	if err := h.db.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Activity
	if err := h.db.Where("(code = ? OR name = ?) AND id != ?", activity.Code, activity.Name, id).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode atau nama kegiatan sudah ada"})
		return
	}

	if err := h.db.Save(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (h *ActivityHandler) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	if err := h.db.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	if err := h.db.Delete(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
