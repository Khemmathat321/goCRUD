package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// PgConnect function
func PgConnect() {
	var err error

	dsn := "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}
