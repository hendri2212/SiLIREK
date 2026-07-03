package seeds

import (
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedOrganizations(db *gorm.DB) error {
	orgs := []models.Organization{
		{
			Number:    "7.01.0.00.0.00.51.0000",
			Name:      "Kecamatan Pulau Laut Sigam",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, org := range orgs {
		var count int64
		db.Model(&models.Organization{}).Where("number = ?", org.Number).Count(&count)
		if count == 0 {
			db.Create(&org)
		}
	}

	return nil
}
