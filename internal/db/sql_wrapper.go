package db

import (
	"gorm.io/gorm"
	"strings"
)

// Mark: Definitions

// SqlWrapper struct
type SqlWrapper[T SqliteConfig] struct {
	name     string
	dbType   string
	config   T
	database *gorm.DB
}

// init - SqlWrapper Constructor - It initializes the wrapper
func (s *SqlWrapper[T]) init(name string, dbType string) error {
	s.name = name
	s.dbType = dbType

	// reading config
	//nameParts := strings.Split(s.name, "/")

	s.config = T{
		FileName: "",
		Options:  nil,
	}

	return nil
}

func (s *SqlWrapper[T]) GetDb() *gorm.DB {
	return nil
}

// MARK: Public functions

// NewSqlWrapper - create a new instance of SqlWrapper and return it
func NewSqlWrapper(name string, dbType string) (*SqlWrapper[SqliteConfig], error) {
	if strings.ToLower(dbType) == "sqlite" {
		wrapper := &SqlWrapper[SqliteConfig]{}
		err := wrapper.init(name, dbType)
		if err != nil {
			return nil, NewCreateSqlWrapperErr(err)
		}

		return wrapper, nil
	}

	return nil, NewNotSupportedDbTypeErr(dbType)
}
