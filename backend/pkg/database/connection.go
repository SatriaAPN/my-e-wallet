package database

import (
	"fmt"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		db = connectDB()
	}

	return db
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		config.DbHost(),
		config.DbUser(),
		config.DbPassword(),
		config.DbName(),
		config.DbPort(),
		config.DbSslMode(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
