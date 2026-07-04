package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedActivities(db *gorm.DB) error {
	var program1 models.Program
	if err := db.Where("code = ?", "7.01.01").First(&program1).Error; err != nil {
		log.Println("Gagal menemukan Program '7.01.01'. Pastikan seeder program dijalankan lebih dulu.")
		return err
	}

	var program2 models.Program
	if err := db.Where("code = ?", "7.01.02").First(&program2).Error; err != nil {
		log.Println("Gagal menemukan Program '7.01.02'. Pastikan seeder program dijalankan lebih dulu.")
		return err
	}

	activities := []models.Activity{
		{
			Code:      "7.01.01.2.01",
			Name:      "Perencanaan, Penganggaran, dan Evaluasi Kinerja Perangkat Daerah",
			ProgramID: program1.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Code:      "7.01.02.2.01",
			Name:      "Koordinasi Penyelenggaraan Kegiatan Pemerintahan di Tingkat Kecamatan",
			ProgramID: program2.ID,
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
