package model

import (
	"github/Khemmathat321/goCRUD/database"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	UserName  string
}

// Migrate func
func (User) Migrate() {
	database.DBConn.AutoMigrate(&User{})
}

func (User) teardown() {
	database.DBConn.Delete(User{})
}
