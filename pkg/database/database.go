package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/yakuter/centralog/pkg/entry"
)

func New(file string) (*gorm.DB, error) {
	conf := &gorm.Config{
		CreateBatchSize: 250,
		Logger:          logger.Default.LogMode(logger.Error),
	}
	db, err := gorm.Open(sqlite.Open(file), conf)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entry.Entry{})

	return db, nil
}

func Insert(db *gorm.DB, entry *entry.Entry) error {
	return db.Create(entry).Error
}

func List(db *gorm.DB) ([]entry.Entry, error) {
	var entries []entry.Entry
	if err := db.Find(&entries).Error; err != nil {
		return nil, err
	}

	return entries, nil
}
