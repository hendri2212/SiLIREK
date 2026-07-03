package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func ItemsHandler(db *gorm.DB) *ItemHandler {
	db.AutoMigrate(&models.Item{})
	return &ItemHandler{db: db}
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	var items []models.Item
	query := h.db

	expenditureAccountID := c.Query("expenditure_account_id")
	if expenditureAccountID != "" {
		query = query.Where("expenditure_account_id = ?", expenditureAccountID)
	}

	query.Order("date desc").Find(&items)
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Item
	if err := h.db.Where("code = ?", item.Code).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode item sudah ada"})
		return
	}

	h.db.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Item
	if err := h.db.Where("code = ? AND id != ?", item.Code, id).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode item sudah ada"})
		return
	}

	if err := h.db.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := h.db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := h.db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
