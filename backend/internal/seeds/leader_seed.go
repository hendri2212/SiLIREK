package seeds

import (
	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedLeaders(db *gorm.DB) {
	var count int64
	db.Model(&models.Leader{}).Count(&count)
	if count == 0 {
		leaders := []models.Leader{
			{Name: "PPTK"},
		}
		db.Create(&leaders)
	}
}
