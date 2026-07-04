package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) error {
	// Ambil rekening belanja "5.1.02.01.001.00052" yang sudah dibuat di expenditure_account_seeder.go
	var acc models.ExpenditureAccount
	err := db.Where("code = ?", "5.1.02.01.001.00052").First(&acc).Error
	if err != nil {
		log.Println("Gagal menemukan rekening belanja '5.1.02.01.001.00052'. Pastikan seeder rekening belanja dijalankan lebih dulu.")
		return err
	}

	date13Mei := time.Date(2026, 5, 13, 0, 0, 0, 0, time.UTC)
	date19Mei := time.Date(2026, 5, 19, 0, 0, 0, 0, time.UTC)

	items := []models.Item{
		{
			Code:                 "935/TBP/KEC.PL.SIGAM/2026",
			Date:                 date13Mei,
			Description:          "Belanja makan minum rapat pertemuan rutin PKK Kelurahan Baharu Selatan, Sub Kegiatan Pemberdayaan Masyarakat di Kelurahan TA. 2026 berdasarkan nota dan data dukung terlampir",
			Credit:               680000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			Code:                 "936/TBP/KEC.PL.SIGAM/2026",
			Date:                 date13Mei,
			Description:          "Belanja makan minum rapat Kantor Kelurahan Baharu Selatan, Sub Kegiatan Pemberdayaan Masyarakat di Kelurahan TA. 2026 berdasarkan nota dan data dukung terlampir",
			Credit:               595000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			Code:                 "962/TBP/KEC.PL.SIGAM/2026",
			Date:                 date19Mei,
			Description:          "Dibayarkan belanja makan dan minum rapat rutin PKK bulan Mei Kelurahan Kotabaru Tengah Sub Kegiatan Pemberdayaan Masyarakat di Kelurahan Tahun Anggaran 2026. berdasarkan nota terlampir",
			Credit:               476000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			Code:                 "967/TBP/KEC.PL.SIGAM/2026",
			Date:                 date19Mei,
			Description:          "Dibayarkan belanja makan dan minum rapat koordinasi bulan Mei Kelurahan Kotabaru Tengah Sub Kegiatan Pemberdayaan Masyarakat di Kelurahan Tahun Anggaran 2026. berdasarkan data dukung terlampir",
			Credit:               2380000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
	}

	for _, item := range items {
		var count int64
		db.Model(&models.Item{}).Where("code = ?", item.Code).Count(&count)
		if count == 0 {
			db.Create(&item)
		}
	}

	return nil
}
