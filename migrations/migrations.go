package migrations

import (
	"github.com/Micxxo/artikel-vision-be/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.Post{})
}
