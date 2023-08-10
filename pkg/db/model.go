package db

import (
	"github.com/abolfazlbeh/zhycan/internal/db"
	"gorm.io/gorm"
)

// GetDb - Get *gorm.DB instance from the underlying interfaces
func GetDb(instanceName string) (*gorm.DB, error) {
	return db.GetManager().GetDb(instanceName)
}

// Migrate - migrate models on specific database
func Migrate(instanceName string, models ...interface{}) error {
	return db.GetManager().Migrate(instanceName, models)
}

// AttachMigrationFunc -  attach migration function to be called by end user
func AttachMigrationFunc(instanceName string, f func(migrator gorm.Migrator) error) error {
	return db.GetManager().AttachMigrationFunc(instanceName, f)
}
