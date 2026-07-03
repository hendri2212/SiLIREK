package seeds

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) {
	SeedOrganizations(db)
	SeedPrograms(db)
	SeedActivities(db)
	SeedSubActivities(db)
	SeedSuperAdmin(db)
}
