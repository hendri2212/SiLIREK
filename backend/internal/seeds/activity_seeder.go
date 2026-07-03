package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedActivities(db *gorm.DB) error {
	// Ambil Program yang akan menjadi parent dari kegiatan ini
	var program models.Program
	if err := db.Where("code = ?", "7.01.01").First(&program).Error; err != nil {
		log.Println("Gagal menemukan Program '7.01.01'. Pastikan seeder program dijalankan lebih dulu.")
		return err
	}

	activities := []models.Activity{
		{
			Code:      "7.01.01.2.01",
			Name:      "Perencanaan, Penganggaran, dan Evaluasi Kinerja Perangkat Daerah",
			ProgramID: program.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, act := range activities {
		var count int64
		db.Model(&models.Activity{}).Where("code = ?", act.Code).Count(&count)
		if count == 0 {
			db.Create(&act)
		}
	}

	return nil
}
