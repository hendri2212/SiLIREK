package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProgramHandler struct {
	db *gorm.DB
}

func ProgramsHandler(db *gorm.DB) *ProgramHandler {
	return &ProgramHandler{db: db}
}

func (h *ProgramHandler) GetPrograms(c *gin.Context) {
	var programs []models.Program
	if err := h.db.Preload("Organization").Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, programs)
}

func (h *ProgramHandler) GetProgramByID(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	if err := h.db.Preload("Organization").First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}
	c.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) CreateProgram(c *gin.Context) {
	var input struct {
		Code           string `json:"code" binding:"required"`
		Name           string `json:"name" binding:"required"`
		OrganizationID uint   `json:"organization_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	program := models.Program{
		Code:           input.Code,
		Name:           input.Name,
		OrganizationID: input.OrganizationID,
	}

	if err := h.db.Create(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, program)
}

func (h *ProgramHandler) UpdateProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	if err := h.db.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}

	var input struct {
		Code           string `json:"code" binding:"required"`
		Name           string `json:"name" binding:"required"`
		OrganizationID uint   `json:"organization_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{
		"code":            input.Code,
		"name":            input.Name,
		"organization_id": input.OrganizationID,
	}

	if err := h.db.Model(&program).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.db.Preload("Organization").First(&program, id)
	c.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) DeleteProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	if err := h.db.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}

	if err := h.db.Delete(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Program deleted successfully"})
}
