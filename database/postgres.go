package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PgConnect function
func PgConnect() (db *gorm.DB, err error) {
	dsn := "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
