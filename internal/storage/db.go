package storage

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Link struct {
	ID        uint       `gorm:"primaryKey"`
	Code      string     `gorm:"uniqueIndex;not null"`
	URL       string     `gorm:"not null"`
	IsActive  bool       `gorm:"default:true"`
	Clicks    int        `gorm:"default:0"`
	ExpiresAt *time.Time `gorm:"default:null"`
	CreatedAt time.Time
}

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&Link{})

	DB = db
}
