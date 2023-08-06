package db

import (
	"reflect"
	"testing"
)

func TestSqlWrapper_Initialize(t *testing.T) {
	wrapper := &SqlWrapper[Sqlite]{name: "test"}
	newWrapper, err := NewSqlWrapper[Sqlite]("test", "sqlite")

	if err != nil {
		t.Errorf("Creating Sql Wrapper --> Expected: %v, but got %v", nil, err)
		return
	}

	if !reflect.DeepEqual(wrapper, newWrapper) {
		t.Errorf("Creating Sql Wrapper --> Expected: %v, but got %v", wrapper, newWrapper)
		return
	}
}

func TestSqlWrapper_SqliteConnection(t *testing.T) {
	makeReadyConfigManager()

	newWrapper, err := NewSqlWrapper[Sqlite]("db/server1", "sqlite")
	if err != nil {
		t.Errorf("Creating Sql Wrapper --> Expected: %v, but got %v", nil, err)
		return
	}

	//db, err := gorm.Open(sqlite.Open("file:file.db?mode=memory&cache=shared&_fk=1"), &gorm.Config{})
	//if err != nil {
	//	t.Errorf("Creating Sql Connection Instance --> Expected: %v, but got %v", nil, err)
	//	return
	//}

	db2, err := newWrapper.GetDb()
	if err != nil {
		t.Errorf("Get database instance --> Expected: %v, but got %v", nil, err)
		return
	}

	tx := db2.Exec("SELECT 1")
	if tx.Error != nil {
		t.Errorf("Query on connected database --> Expected no Error: %v, but got %v", nil, err)
		return
	}
}

func TestSqlWrapper_SqliteConnectionConfiguration(t *testing.T) {

	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	SkipDefaultTransaction:                   false,
	//	NamingStrategy:                           nil,
	//	FullSaveAssociations:                     false,
	//	Logger:                                   nil,
	//	NowFunc:                                  nil,
	//	DryRun:                                   false,
	//	PrepareStmt:                              false,
	//	DisableAutomaticPing:                     false,
	//	DisableForeignKeyConstraintWhenMigrating: false,
	//	IgnoreRelationshipsWhenMigrating:         false,
	//	DisableNestedTransaction:                 false,
	//	AllowGlobalUpdate:                        false,
	//	QueryFields:                              false,
	//	CreateBatchSize:                          0,
	//	TranslateError:                           false,
	//	ClauseBuilders:                           nil,
	//	ConnPool:                                 nil,
	//	Dialector:                                nil,
	//	Plugins:                                  nil,
	//})
	//
	//db.Logger.LogMode()

}
