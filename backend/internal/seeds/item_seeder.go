package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) error {
	// Karena di gambar ada kode rekening 5.1.02.04.001.00003, kita pastikan rekening ini ada dulu
	var acc models.ExpenditureAccount
	err := db.Where("code = ?", "5.1.02.04.001.00003").First(&acc).Error
	if err != nil {
		// Jika belum ada, kita buatkan dulu agar relasinya tidak error
		var sub models.SubActivity
		if errSub := db.Where("code = ?", "7.01.01.2.01.0001").First(&sub).Error; errSub == nil {
			acc = models.ExpenditureAccount{
				Code:          "5.1.02.04.001.00003",
				Description:   "Belanja Perjalanan Dinas (Otomatis dari Item)",
				BudgetCeiling: 10000000.00, // Dummy pagu
				SubActivityID: sub.ID,
			}
			db.Create(&acc)
		} else {
			log.Println("Gagal menemukan rekening dan gagal membuat karena Sub Kegiatan tidak ada.")
			return err
		}
	}

	date13Mei := time.Date(2026, 5, 13, 0, 0, 0, 0, time.UTC)

	items := []models.Item{
		{
			Code:                 "921/TBP/KEC.PL.SIGAM/2026",
			Date:                 date13Mei,
			Description:          "Dibayarkan biaya transport perjalanan dinas An. Gusti wahidah dalam rangka mengamprah Gaji PNS dan Gaji PPPK Kecamatan Pulaulaut Sigam pada tanggal 04 Mei 2026 di BPKAD Kotabaru. berdasarkan Surat Tugas No.000.1.2.3/231/PLSi tanggal 04 Mei 2026 Sub Kegiatan Penyelenggaraan Rapat Koordinasi dan Konsultasi SKPD",
			Credit:               150000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			Code:                 "922/TBP/KEC.PL.SIGAM/2026",
			Date:                 date13Mei,
			Description:          "Dibayarkan Biaya Transport Perjalanan Dinas an. Gusti Wahidah dalam rangka mengamprah TPP PNS dan TPP PPPK bulan April 2026 pada tanggal 11 Mei 2026 di BPKAD Kotabaru. berdasarkan Surat Tugas No.000.1.2.3/234/PLSi tanggal 11 Mei 2026 Sub Kegiatan Penyelenggaraan Rapat Koordinasi dan Konsultasi SKPD",
			Credit:               150000.00,
			ExpenditureAccountID: acc.ID,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			Code:                 "923/TBP/KEC.PL.SIGAM/2026",
			Date:                 date13Mei,
			Description:          "Dibayarkan Biaya Transport Perjalanan Dinas an. Syamsuddin dalam rangka mengamprah SPM LS Listrik, menyampaikan Laporan Bulanan dan LRA Kecamatan Pulaulaut Sigam pada tanggal 06 Mei 2026 di BPKAD Kotabaru. berdasarkan Surat Tugas No.000.1.2.3/232/PLSi tanggal 06 Mei 2026 Sub Kegiatan Penyelenggaraan Rapat Koordinasi dan Konsultasi SKPD",
			Credit:               150000.00,
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
