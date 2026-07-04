package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedExpenditureAccounts(db *gorm.DB) error {
	// Ambil Sub Kegiatan yang akan menjadi parent dari rekening belanja ini
	var subActivity1 models.SubActivity
	if err := db.Where("code = ?", "7.01.02.2.01.0001").First(&subActivity1).Error; err != nil {
		log.Println("Gagal menemukan Sub Kegiatan '7.01.02.2.01.0001'. Pastikan seeder sub kegiatan dijalankan lebih dulu.")
		return err
	}

	var subActivity2 models.SubActivity
	if err := db.Where("code = ?", "7.01.01.2.01.0001").First(&subActivity2).Error; err != nil {
		log.Println("Gagal menemukan Sub Kegiatan '7.01.01.2.01.0001'. Pastikan seeder sub kegiatan dijalankan lebih dulu.")
		return err
	}

	accounts := []models.ExpenditureAccount{
		// RELASI KE 7.01.02.2.01.0001
		{
			Code:          "5.1.02.01.001.00024",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Alat Tulis Kantor",
			BudgetCeiling: 187000.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00025",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Kertas dan Cover",
			BudgetCeiling: 790000.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00026",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Bahan Cetak",
			BudgetCeiling: 958800.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00029",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Bahan Komputer",
			BudgetCeiling: 801000.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00052",
			Description:   "Belanja Makanan dan Minuman Rapat",
			BudgetCeiling: 2720000.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.04.001.00003",
			Description:   "Belanja Perjalanan Dinas Dalam Kota",
			BudgetCeiling: 1295000.00,
			SubActivityID: subActivity1.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		
		// RELASI KE 7.01.01.2.01.0001
		{
			Code:          "5.1.02.01.001.00024",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Alat Tulis Kantor",
			BudgetCeiling: 1584000.00,
			SubActivityID: subActivity2.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00025",
			Description:   "Belanja Alat/Bahan untuk Kegiatan Kantor - Kertas dan Cover",
			BudgetCeiling: 1544000.00,
			SubActivityID: subActivity2.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			Code:          "5.1.02.01.001.00052",
			Description:   "Belanja Makanan dan Minuman Rapat",
			BudgetCeiling: 5440000.00,
			SubActivityID: subActivity2.ID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	for _, acc := range accounts {
		var count int64
		db.Model(&models.ExpenditureAccount{}).
			Where("code = ? AND sub_activity_id = ?", acc.Code, acc.SubActivityID).
			Count(&count)
			
		if count == 0 {
			db.Create(&acc)
		} else {
			// Update jika sudah ada untuk memastikan nilai budget dan parent-nya benar
			db.Model(&models.ExpenditureAccount{}).
				Where("code = ? AND sub_activity_id = ?", acc.Code, acc.SubActivityID).
				Updates(map[string]interface{}{
					"description":    acc.Description,
					"budget_ceiling": acc.BudgetCeiling,
					"updated_at":     time.Now(),
				})
		}
	}

	return nil
}
