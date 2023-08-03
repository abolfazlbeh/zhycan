package db

import (
	"reflect"
	"testing"
)

func TestSqlWrapper_Initialize(t *testing.T) {
	wrapper := &SqlWrapper[SqliteConfig]{name: "test", dbType: "sqlite"}
	newWrapper, err := NewSqlWrapper("test", "sqlite")

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

	newWrapper, err := NewSqlWrapper("db/server1", "sqlite")
	if err != nil {
		t.Errorf("Creating Sql Wrapper --> Expected: %v, but got %v", nil, err)
		return
	}

	//db, err := gorm.Open(sqlite.Open("file:file.db?mode=memory&cache=shared&_fk=1"), &gorm.Config{})
	//if err != nil {
	//	t.Errorf("Creating Sql Connection Instance --> Expected: %v, but got %v", nil, err)
	//	return
	//}

	db2 := newWrapper.GetDb()
	if db2 == nil {
		t.Errorf("Get Sql Wrapper Instance --> Expected: instance, but got %v", nil)
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

}
