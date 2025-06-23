package seeds

import (
	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedPositions(db *gorm.DB) {
	var count int64
	db.Model(&models.Position{}).Count(&count)
	if count == 0 {
		positions := []models.Position{
			{Name: "Kasubag Perencanaan dan Keuangan"},
			{Name: "Kasubahg Umum dan Kepegawaian"},
			{Name: "Kepala Seksi Pemerintahan"},
			{Name: "Kepala Seksi Ketentraman dan Ketertiban"},
			{Name: "Kepala Seksi Pelayanan"},
			{Name: "Kepala Seksi Pembangunan dan Pemberdayaan Masyarakat Desa/Kelurahan"},
			{Name: "Kepala Seksi Kemasyarakatan dan LH"},
		}
		db.Create(&positions)
	}
}
