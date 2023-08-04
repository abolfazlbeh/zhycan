package db

import (
	"fmt"
	"github.com/abolfazlbeh/zhycan/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// Mark: Definitions

// SqlWrapper struct
type SqlWrapper[T SqlConfigurable] struct {
	name             string
	config           T
	databaseInstance *gorm.DB
}

// init - SqlWrapper Constructor - It initializes the wrapper
func (s *SqlWrapper[T]) init(name string) error {
	s.name = name

	// reading config
	nameParts := strings.Split(s.name, "/")

	if reflect.ValueOf(s.config).Type() == reflect.TypeOf(SqliteConfig{}) {
		filenameKey := fmt.Sprintf("%s.%s", nameParts[1], "db")
		filenameStr, err := config.GetManager().Get(nameParts[0], filenameKey)
		if err != nil {
			return err
		}

		optionsKey := fmt.Sprintf("%s.%s", nameParts[1], "options")
		optionsObj, err := config.GetManager().Get(nameParts[0], optionsKey)
		if err != nil {
			return err
		}

		optionsMap := make(map[string]string, len(optionsObj.(map[string]interface{})))
		for key, item := range optionsObj.(map[string]interface{}) {
			optionsMap[key] = item.(string)
		}

		s.config = T(reflect.ValueOf(SqliteConfig{
			FileName: filenameStr.(string),
			Options:  optionsMap,
		}).Interface().(T))
	}

	return nil
}

func (s *SqlWrapper[T]) GetDb() (*gorm.DB, error) {
	if s.databaseInstance == nil {
		if reflect.ValueOf(s.config).Type() == reflect.TypeOf(SqliteConfig{}) {
			optionsQSArr := make([]string, 0)
			config := reflect.ValueOf(s.config).Interface().(SqliteConfig)
			for key, val := range config.Options {
				optionsQSArr = append(optionsQSArr, fmt.Sprintf("%s=%s", key, val))
			}
			optionsQS := strings.Join(optionsQSArr, "&")

			dsn := fmt.Sprintf("file:%s?%s", config.FileName, optionsQS)

			db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			s.databaseInstance = db
		}
	}

	return s.databaseInstance, nil
}

// MARK: Public functions

// NewSqlWrapper - create a new instance of SqlWrapper and return it
func NewSqlWrapper[T SqlConfigurable](name string, dbType string) (*SqlWrapper[T], error) {
	if strings.ToLower(dbType) == "sqlite" {
		wrapper := &SqlWrapper[T]{}
		err := wrapper.init(name)
		if err != nil {
			return nil, NewCreateSqlWrapperErr(err)
		}

		return wrapper, nil
	}

	return nil, NewNotSupportedDbTypeErr(dbType)
}
