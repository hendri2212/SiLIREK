package handlers

import (
	"net/http"
	"silirek/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpenditureAccountHandler struct {
	db *gorm.DB
}

func ExpenditureAccountsHandler(db *gorm.DB) *ExpenditureAccountHandler {
	db.AutoMigrate(&models.ExpenditureAccount{})
	return &ExpenditureAccountHandler{db: db}
}

type ExpenditureAccountResponse struct {
	models.ExpenditureAccount
	TotalCredit     float64 `json:"total_credit"`
	RemainingBudget float64 `json:"remaining_budget"`
}

func (h *ExpenditureAccountHandler) GetExpenditureAccounts(c *gin.Context) {
	var accounts []models.ExpenditureAccount
	query := h.db

	subActivityID := c.Query("sub_activity_id")
	if subActivityID != "" {
		query = query.Where("sub_activity_id = ?", subActivityID)
	}

	query.Find(&accounts)

	var responses []ExpenditureAccountResponse
	for _, account := range accounts {
		var totalCredit float64
		h.db.Model(&models.Item{}).Where("expenditure_account_id = ?", account.ID).Select("COALESCE(SUM(credit), 0)").Scan(&totalCredit)
		
		responses = append(responses, ExpenditureAccountResponse{
			ExpenditureAccount: account,
			TotalCredit:        totalCredit,
			RemainingBudget:    account.BudgetCeiling - totalCredit,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (h *ExpenditureAccountHandler) CreateExpenditureAccount(c *gin.Context) {
	var account models.ExpenditureAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.ExpenditureAccount
	if err := h.db.Where("code = ?", account.Code).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode rekening belanja sudah ada"})
		return
	}

	h.db.Create(&account)
	c.JSON(http.StatusCreated, account)
}

func (h *ExpenditureAccountHandler) GetExpenditureAccountByID(c *gin.Context) {
	id := c.Param("id")
	var account models.ExpenditureAccount
	if err := h.db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rekening Belanja not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *ExpenditureAccountHandler) UpdateExpenditureAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.ExpenditureAccount
	if err := h.db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rekening Belanja not found"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.ExpenditureAccount
	if err := h.db.Where("code = ? AND id != ?", account.Code, id).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kode rekening belanja sudah ada"})
		return
	}

	if err := h.db.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *ExpenditureAccountHandler) DeleteExpenditureAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.ExpenditureAccount
	if err := h.db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rekening Belanja not found"})
		return
	}

	if err := h.db.Delete(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rekening Belanja deleted successfully"})
}
