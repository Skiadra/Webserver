package migrations

import (
	"log"
	"webserver/internal/models"

	"gorm.io/gorm"
)

func UserMigrations(db *gorm.DB) {
	// Migrate the User model
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed")
}
