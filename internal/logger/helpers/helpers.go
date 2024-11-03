package helpers

import (
	"github.com/abolfazlbeh/zhycan/internal/db"
	"gorm.io/gorm"
)

func GetSqlDbInstance(instanceName string) (*gorm.DB, error) {
	return db.GetManager().GetDb(instanceName)
}
