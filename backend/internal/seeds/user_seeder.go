package seeds

import (
	"time"

	"silirek/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedSuperAdmin membuat user superadmin dengan nip, position_id, leader_id null
func SeedSuperAdmin(db *gorm.DB) error {
	// Cek dulu apakah superadmin sudah ada
	var count int64
	db.Model(&models.User{}).
		Where("role = ?", "superadmin").
		Count(&count)
	if count > 0 {
		return nil // sudah pernah di‐seed
	}

	// Hash default password (ganti dengan env var jika perlu)
	pwd, err := bcrypt.GenerateFromPassword([]byte("Secret@123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	super := models.User{
		FullName:   "Super Admin",
		Email:      "arifin.hendri465@gmail.com",
		Password:   string(pwd),
		Nip:        nil, // pointer nil => NULL in DB
		PositionID: nil, // pointer nil => NULL in DB
		LeaderID:   nil, // pointer nil => NULL in DB
		Role:       models.UserRoleSuperadmin,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return db.Create(&super).Error
}
