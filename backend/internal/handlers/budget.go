package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BudgetHandler struct {
	db *gorm.DB
}

func BudgetsHandler(db *gorm.DB) *BudgetHandler {
	db.AutoMigrate(&models.Budget{})
	return &BudgetHandler{db: db}
}
func (h *BudgetHandler) GetBudgets(c *gin.Context) {
	var budgets []models.Budget
	h.db.Preload("User.Leader").Preload("User.Position").Preload("Quarterly").Preload("Activity").Find(&budgets)

	c.JSON(http.StatusOK, budgets)
}

// func (h *BudgetHandler) UsedBudgets(c *gin.Context) {
// 	var budgets []models.Budget
// 	h.db.Find(&budgets)

// 	var response []models.BudgetResponse
// 	for _, b := range budgets {
// 		response = append(response, models.BudgetResponse{
// 			ID: b.ID,
// 			UserID: func() uint {
// 				if b.UserID != nil {
// 					return *b.UserID
// 				} else {
// 					return 0
// 				}
// 			}(),
// 			QuarterlyID:     b.QuarterlyID,
// 			ActivityID:      b.ActivityID,
// 			ParentID:        b.ParentID,
// 			Budget:          b.Budget,
// 			UsedBudget:      b.UsedBudget,
// 			RemainingBudget: b.GetRemainingBudget(),
// 		})
// 	}

//		c.JSON(http.StatusOK, response)
//	}
func (h *BudgetHandler) CreateBudget(c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if budget.SubActivityID != nil {
		// Simpan activity_id original untuk referensi parent
		originalActivityID := budget.ActivityID

		// Ubah activity_id menjadi sub-activity
		budget.ActivityID = *budget.SubActivityID

		// Cari ID dari budget parent berdasarkan activity_id
		var parentBudget models.Budget
		query := h.db.Where("activity_id = ? AND quarterly_id = ?", originalActivityID, budget.QuarterlyID)
		if budget.UserID == nil {
			query = query.Where("user_id IS NULL")
		} else {
			query = query.Where("user_id = ?", *budget.UserID)
		}
		// if err := query.First(&parentBudget).Error; err == nil {
		// 	budget.ParentID = &parentBudget.ID
		// }
		if err := query.First(&parentBudget).Error; err == nil && parentBudget.ID != 0 && parentBudget.ParentID == nil {
			budget.ParentID = &parentBudget.ID
		}
	}

	var existing models.Budget
	query := h.db.Where("quarterly_id = ? AND activity_id = ?", budget.QuarterlyID, budget.ActivityID)
	if budget.UserID == nil {
		query = query.Where("user_id IS NULL")
	} else {
		query = query.Where("user_id = ?", *budget.UserID)
	}
	if budget.ParentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", budget.ParentID)
	}
	if err := query.First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Budget for this activity already exists"})
		return
	}

	h.db.Create(&budget)

	c.JSON(http.StatusCreated, budget)
}
func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	var budget models.Budget
	if err := h.db.First(&budget, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}

	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if budget.UsedBudget > budget.Budget {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Used budget cannot exceed total budget"})
		return
	}

	h.db.Save(&budget)

	c.JSON(http.StatusOK, budget)
}

// func (h *BudgetHandler) DeleteBudget(c *gin.Context) {
// 	id := c.Param("id")
// 	var budget models.Budget
// 	if err := h.db.First(&budget, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
// 		return
// 	}

// 	h.db.Delete(&budget)

//		c.JSON(http.StatusOK, gin.H{"message": "budget deleted"})
//	}
func (h *BudgetHandler) GetBudgetByID(c *gin.Context) {
	id := c.Param("id")
	var budget models.Budget
	if err := h.db.Preload("User").Preload("Quarterly").Preload("Activity").First(&budget, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               budget.ID,
		"user":             budget.User,
		"quarterly":        budget.Quarterly,
		"activity":         budget.Activity,
		"budget":           budget.Budget,
		"used_budget":      budget.UsedBudget,
		"remaining_budget": budget.GetRemainingBudget(),
	})
}

// func (h *BudgetHandler) GetBudgetsByQuarterlyID(c *gin.Context) {
// 	quarterlyID := c.Param("quarterly_id")
// 	var budgets []models.Budget
// 	if err := h.db.Where("quarterly_id = ?", quarterlyID).Find(&budgets).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "budgets not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, budgets)
// }
// func (h *BudgetHandler) GetBudgetsByActivityID(c *gin.Context) {
// 	activityID := c.Param("activity_id")
// 	var budgets []models.Budget
// 	if err := h.db.Where("activity_id = ?", activityID).Find(&budgets).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "budgets not found"})
// 		return
// 	}

//		c.JSON(http.StatusOK, budgets)
//	}
func (h *BudgetHandler) GetBudgetsByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	var budgets []models.Budget
	query := h.db.Preload("User").Preload("Quarterly").Preload("Activity")
	if userID != "1" {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.Find(&budgets).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budgets not found"})
		return
	}

	var response []models.BudgetResponse
	for _, b := range budgets {
		response = append(response, models.BudgetResponse{
			ID:              b.ID,
			Activity:        b.Activity,
			ParentID:        b.ParentID,
			Budget:          b.Budget,
			UsedBudget:      b.UsedBudget,
			RemainingBudget: b.GetRemainingBudget(),
		})
	}

	c.JSON(http.StatusOK, response)
}
func (h *BudgetHandler) GetBudgetsByUserIDAndQuarterlyID(c *gin.Context) {
	userID := c.Param("user_id")
	quarterlyID := c.Param("quarterly_id")

	var budgets []models.Budget
	query := h.db.Preload("User").Preload("Quarterly").Preload("Activity").Where("quarterly_id = ?", quarterlyID)

	if userID != "" && userID != "null" && userID != "0" {
		query = query.Where("user_id = ?", userID)
	} else {
		query = query.Where("user_id IS NULL")
	}

	if err := query.Find(&budgets).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "budgets not found"})
		return
	}

	c.JSON(http.StatusOK, budgets)
}
