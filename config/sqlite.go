package config

import (
	"os"

	"github.com/Samuel-Ricardo/jobileu/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Error("Failed to create database directory")
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Error("Failed to create database file")
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database")
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Openig{})
	if err != nil {
		logger.Error("Failed to migrate database")
		return nil, err
	}

	return db, nil
}
