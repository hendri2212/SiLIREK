package main

import (
	"log"
	"os"
	"silirek/internal/models"
	"silirek/internal/routes"
	"silirek/internal/seeds"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return db
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	db := InitDB()

	// HATI-HATI: Hanya gunakan ini untuk development saat merombak database!
	// Ini akan menghapus semua tabel dan membuatnya ulang dari awal.
	db.Migrator().DropTable(
		&models.Item{},
		&models.ExpenditureAccount{},
		&models.SubActivity{},
		&models.Activity{},
		&models.Program{},
		&models.User{},
		&models.Organization{},
	)

	// Auto migrate semua tabel
	db.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.Program{},
		&models.Activity{},
		&models.SubActivity{},
		&models.ExpenditureAccount{},
		&models.Item{},
	)

	// Seed data
	seeds.SeedAll(db)

	router := gin.Default()
	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
