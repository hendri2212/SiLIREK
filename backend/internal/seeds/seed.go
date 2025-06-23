package seeds

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) {
	SeedSuperAdmin(db)
	SeedPositions(db)
	SeedLeaders(db)
}
