package seeds

import (
	"log"
	"time"

	"silirek/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUsers membuat akun superadmin, admin, dan user
func SeedUsers(db *gorm.DB) error {
	// Hash default password
	pwd, err := bcrypt.GenerateFromPassword([]byte("Secret@123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Ambil ID Organisasi untuk admin dan user
	var org models.Organization
	orgID := (*uint)(nil)
	if err := db.Where("number = ?", "7.01.0.00.0.00.51.0000").First(&org).Error; err == nil {
		orgID = &org.ID
	} else {
		log.Println("Peringatan: Organisasi tidak ditemukan, admin dan user akan diset tanpa organisasi.")
	}

	photoSuperAdmin := "uploads/photos/20250512085110_199108022024211014.jpeg"
	photoAdmin := "uploads/photos/20250515114056_access bars.jpg"

	users := []models.User{
		{
			FullName:       "Super Admin",
			Email:          "arifin.hendri465@gmail.com",
			Password:       string(pwd),
			Nip:            nil,
			OrganizationID: nil, // Superadmin bebas dari organisasi
			Role:           models.UserRoleSuperadmin,
			Photo:          &photoSuperAdmin,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			FullName:       "Admin Instansi",
			Email:          "admin@gmail.com",
			Password:       string(pwd),
			Nip:            nil,
			OrganizationID: orgID,
			Role:           models.UserRoleAdmin,
			Photo:          &photoAdmin,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			FullName:       "User Pegawai",
			Email:          "user@gmail.com",
			Password:       string(pwd),
			Nip:            nil,
			OrganizationID: orgID,
			Role:           models.UserRoleUser,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	for _, u := range users {
		var count int64
		db.Model(&models.User{}).Where("email = ?", u.Email).Count(&count)
		if count == 0 {
			db.Create(&u)
		}
	}

	return nil
}
