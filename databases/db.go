package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbPath))
	if err != nil {
		panic(err)
	}

	DB = db

	return db, nil
}
