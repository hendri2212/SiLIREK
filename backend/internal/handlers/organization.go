package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	db *gorm.DB
}

func OrganizationsHandler(db *gorm.DB) *OrganizationHandler {
	return &OrganizationHandler{db: db}
}

func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	var organizations []models.Organization
	if err := h.db.Preload("Parent").Preload("Children").Find(&organizations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organizations)
}

func (h *OrganizationHandler) GetOrganizationByID(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := h.db.Preload("Parent").Preload("Children").First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Number   string `json:"number" binding:"required"`
		ParentID *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization := models.Organization{
		Name:     input.Name,
		Number:   input.Number,
		ParentID: input.ParentID,
	}

	if err := h.db.Create(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, organization)
}

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := h.db.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	var input struct {
		Name     string `json:"name" binding:"required"`
		Number   string `json:"number" binding:"required"`
		ParentID *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update specific fields to handle NULL for ParentID properly
	updates := map[string]interface{}{
		"name":      input.Name,
		"number":    input.Number,
		"parent_id": input.ParentID, // This will correctly set it to NULL if nil
	}

	if err := h.db.Model(&organization).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch updated
	h.db.Preload("Parent").First(&organization, id)

	c.JSON(http.StatusOK, organization)
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := h.db.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	if err := h.db.Delete(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}
