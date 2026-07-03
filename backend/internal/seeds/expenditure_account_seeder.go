package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedExpenditureAccounts(db *gorm.DB) error {
	// Ambil Sub Kegiatan yang akan menjadi parent dari rekening belanja ini
	var subActivity models.SubActivity
	if err := db.Where("code = ?", "7.01.01.2.01.0001").First(&subActivity).Error; err != nil {
		log.Println("Gagal menemukan Sub Kegiatan '7.01.01.2.01.0001'. Pastikan seeder sub kegiatan dijalankan lebih dulu.")
		return err
	}

	accounts := []models.ExpenditureAccount{
		{
			Code:          "5.1.02.01.001.00024",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor-Alat Tulis Kantor",
			BudgetCeiling: 1584000.00,
			SubActivityID: subActivity.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00025",
			Description:   "Belanja Alat /Bahan untuk Kegiatan Kantor- Kertas dan Cover",
			BudgetCeiling: 1544000.00,
			SubActivityID: subActivity.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00052",
			Description:   "Belanja Makanan dan Minuman Rapat",
			BudgetCeiling: 5440000.00,
			SubActivityID: subActivity.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	for _, acc := range accounts {
		var count int64
		db.Model(&models.ExpenditureAccount{}).Where("code = ?", acc.Code).Count(&count)
		if count == 0 {
			db.Create(&acc)
		}
	}

	return nil
}
