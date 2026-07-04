package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedPrograms(db *gorm.DB) error {
	// Ambil Organisasi yang akan menjadi parent dari program ini
	var org models.Organization
	if err := db.Where("number = ?", "7.01.0.00.0.00.51.0000").First(&org).Error; err != nil {
		log.Println("Gagal menemukan Organisasi 'Kecamatan Pulau Laut Sigam'. Pastikan seeder organisasi dijalankan lebih dulu.")
		return err
	}

	programs := []models.Program{
		{
			Code:           "7.01.01",
			Name:           "PROGRAM PENUNJANG URUSAN PEMERINTAHAN DAERAH KABUPATEN/KOTA",
			OrganizationID: org.ID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			Code:           "7.01.02",
			Name:           "PROGRAM PENYELENGGARAAN PEMERINTAHAN DAN PELAYANAN PUBLIK",
			OrganizationID: org.ID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			Code:           "7.01.03",
			Name:           "PROGRAM PEMBERDAYAAN MASYARAKAT DESA DAN KELURAHAN",
			OrganizationID: org.ID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	for _, prog := range programs {
		var count int64
		db.Model(&models.Program{}).Where("code = ?", prog.Code).Count(&count)
		if count == 0 {
			db.Create(&prog)
		}
	}

	return nil
}
