package config

import (
	"aws-sns-local-go/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sns.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Topic{})
	db.AutoMigrate(&domain.Message{})

	return db
}
