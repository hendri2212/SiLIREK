package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedSubActivities(db *gorm.DB) error {
	// Ambil Kegiatan yang akan menjadi parent dari sub kegiatan ini
	var activity models.Activity
	if err := db.Where("code = ?", "7.01.01.2.01").First(&activity).Error; err != nil {
		log.Println("Gagal menemukan Kegiatan '7.01.01.2.01'. Pastikan seeder kegiatan dijalankan lebih dulu.")
		return err
	}

	subActivities := []models.SubActivity{
		{
			Code:       "7.01.01.2.01.0001",
			Name:       "Penyusunan Dokumen Perencanaan Perangkat Daerah",
			ActivityID: activity.ID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for _, subAct := range subActivities {
		var count int64
		db.Model(&models.SubActivity{}).Where("code = ?", subAct.Code).Count(&count)
		if count == 0 {
			db.Create(&subAct)
		}
	}

	return nil
}
