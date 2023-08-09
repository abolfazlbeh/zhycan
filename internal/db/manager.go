package db

import (
	"fmt"
	"github.com/abolfazlbeh/zhycan/internal/config"
	"github.com/abolfazlbeh/zhycan/internal/utils"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
	"sync"
)

// Mark: manager

// manager object
type manager struct {
	name                string
	lock                sync.Mutex
	sqliteDbInstances   map[string]*SqlWrapper[Sqlite]
	mysqlDbInstances    map[string]*SqlWrapper[Mysql]
	postgresDbInstances map[string]*SqlWrapper[Postgresql]
	supportedDBs        []string

	isManagerInitialized bool
}

// MARK: Module variables
var managerInstance *manager = nil
var once sync.Once

// Module init function
func init() {
	log.Println("DB Manager Package Initialized...")
}

// init - Manager Constructor - It initializes the manager configuration params
func (m *manager) init() {
	m.name = "db"
	m.isManagerInitialized = false

	m.lock.Lock()
	defer m.lock.Unlock()

	m.supportedDBs = []string{"sqlite", "mysql", "postgresql"}

	// read configs
	connectionsObj, err := config.GetManager().Get(m.name, "connections")
	if err != nil {
		return
	}

	m.sqliteDbInstances = make(map[string]*SqlWrapper[Sqlite])
	m.mysqlDbInstances = make(map[string]*SqlWrapper[Mysql])
	m.postgresDbInstances = make(map[string]*SqlWrapper[Postgresql])

	for _, item := range connectionsObj.([]interface{}) {
		dbInstanceName := item.(string)

		dbTypeKey := fmt.Sprintf("%s.%s", dbInstanceName, "type")
		dbTypeInf, err := config.GetManager().Get(m.name, dbTypeKey)
		if err != nil {
			continue
		}

		//  create a new instance based on type
		dbType := strings.ToLower(dbTypeInf.(string))
		if utils.ArrayContains(&m.supportedDBs, dbType) {
			switch dbType {
			case "sqlite":
				obj, err := NewSqlWrapper[Sqlite](fmt.Sprintf("db/%s", dbInstanceName), dbType)
				if err != nil {
					// TODO: log error here
					return
				}

				m.sqliteDbInstances[dbInstanceName] = reflect.ValueOf(obj).Interface().(*SqlWrapper[Sqlite])
				break
			case "mysql":
				obj, err := NewSqlWrapper[Mysql](fmt.Sprintf("db/%s", dbInstanceName), dbType)
				if err != nil {
					// TODO: log error here
					return
				}

				m.mysqlDbInstances[dbInstanceName] = reflect.ValueOf(obj).Interface().(*SqlWrapper[Mysql])
				break
			case "postgresql":
				obj, err := NewSqlWrapper[Postgresql](fmt.Sprintf("db/%s", dbInstanceName), dbType)
				if err != nil {
					// TODO: log error here
					return
				}

				m.postgresDbInstances[dbInstanceName] = reflect.ValueOf(obj).Interface().(*SqlWrapper[Postgresql])
				break
			}
		}
	}

	m.isManagerInitialized = true
}

// restartOnChangeConfig - subscribe a function for when the config is changed
func (m *manager) restartOnChangeConfig() {
	// Config config server to reload
	wrapper, err := config.GetManager().GetConfigWrapper(m.name)
	if err == nil {
		wrapper.RegisterChangeCallback(func() interface{} {
			if m.isManagerInitialized {
				m.init()
			}
			return nil
		})
	} else {
		// TODO: make some logs
	}
}

// MARK: Public Functions

// GetManager - This function returns singleton instance of Db Manager
func GetManager() *manager {
	// once used for prevent race condition and manage critical section.
	once.Do(func() {
		managerInstance = &manager{}
		managerInstance.init()
		managerInstance.restartOnChangeConfig()
	})
	return managerInstance
}

// GetDb - Get *gorm.DB instance from the underlying interfaces
func (m *manager) GetDb(instanceName string) (*gorm.DB, error) {
	if v, ok := m.sqliteDbInstances[instanceName]; ok {
		return v.GetDb()
	} else if v, ok := m.mysqlDbInstances[instanceName]; ok {
		return v.GetDb()
	} else if v, ok := m.postgresDbInstances[instanceName]; ok {
		return v.GetDb()
	}

	return nil, NewNotExistServiceNameErr(instanceName)
}