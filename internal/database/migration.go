package database

import (
	"github.com/NeerChayaphon/go-rest-api/internal/todo"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrate our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&todo.Todo{}); result.Error != nil {
		return result.Error
	}
	return nil
}
