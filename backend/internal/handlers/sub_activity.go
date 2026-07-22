package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubActivityHandler struct {
	db *gorm.DB
}

func SubActivitiesHandler(db *gorm.DB) *SubActivityHandler {
	return &SubActivityHandler{db: db}
}

type SubActivityResponse struct {
	models.SubActivity
	TotalBudget     float64 `json:"total_budget"`
	RemainingBudget float64 `json:"remaining_budget"`
}

func (h *SubActivityHandler) GetSubActivities(c *gin.Context) {
	var subActivities []models.SubActivity
	query := h.db

	// Filter by ActivityID if provided in query params
	activityID := c.Query("activity_id")
	if activityID != "" {
		query = query.Where("activity_id = ?", activityID)
	}

	if err := query.Find(&subActivities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := []SubActivityResponse{}
	for _, sa := range subActivities {
		var totalBudget float64
		h.db.Model(&models.ExpenditureAccount{}).
			Where("sub_activity_id = ?", sa.ID).
			Select("COALESCE(SUM(budget_ceiling), 0)").
			Scan(&totalBudget)

		var totalCredit float64
		h.db.Model(&models.Item{}).
			Joins("JOIN expenditure_accounts ON expenditure_accounts.id = items.expenditure_account_id").
			Where("expenditure_accounts.sub_activity_id = ?", sa.ID).
			Select("COALESCE(SUM(items.credit), 0)").
			Scan(&totalCredit)

		responses = append(responses, SubActivityResponse{
			SubActivity:     sa,
			TotalBudget:     totalBudget,
			RemainingBudget: totalBudget - totalCredit,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (h *SubActivityHandler) CreateSubActivity(c *gin.Context) {
	var subActivity models.SubActivity
	if err := c.ShouldBindJSON(&subActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.SubActivity
	if err := h.db.Where("code = ? OR name = ?", subActivity.Code, subActivity.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode atau nama sub kegiatan sudah ada"})
		return
	}

	if err := h.db.Create(&subActivity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, subActivity)
}

func (h *SubActivityHandler) GetSubActivityByID(c *gin.Context) {
	id := c.Param("id")
	var subActivity models.SubActivity
	if err := h.db.First(&subActivity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub Activity not found"})
		return
	}
	c.JSON(http.StatusOK, subActivity)
}

func (h *SubActivityHandler) UpdateSubActivity(c *gin.Context) {
	id := c.Param("id")
	var subActivity models.SubActivity
	if err := h.db.First(&subActivity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub Activity not found"})
		return
	}

	if err := c.ShouldBindJSON(&subActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.SubActivity
	if err := h.db.Where("(code = ? OR name = ?) AND id != ?", subActivity.Code, subActivity.Name, id).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode atau nama sub kegiatan sudah ada"})
		return
	}

	if err := h.db.Save(&subActivity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subActivity)
}

func (h *SubActivityHandler) DeleteSubActivity(c *gin.Context) {
	id := c.Param("id")
	var subActivity models.SubActivity
	if err := h.db.First(&subActivity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub Activity not found"})
		return
	}

	if err := h.db.Delete(&subActivity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sub Activity deleted successfully"})
}
