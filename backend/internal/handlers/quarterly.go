package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuarterlyHandler struct {
	db *gorm.DB
}

func QuarterliesHandler(db *gorm.DB) *QuarterlyHandler {
	db.AutoMigrate(&models.Quarterly{})
	return &QuarterlyHandler{db: db}
}
func (h *QuarterlyHandler) GetQuarterlies(c *gin.Context) {
	var quarterlies []models.Quarterly
	h.db.Order("year asc").Order("quarter asc").Limit(8).Find(&quarterlies)

	c.JSON(http.StatusOK, quarterlies)
}
