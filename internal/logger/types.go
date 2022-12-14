package logger

// Imports needed list
import (
	"time"
)

// Some Constants
const (
	max = 6
)

// LogObject - all methods that want to log must transfer object of this.
type LogObject struct {
	Level      LogLevel
	Module     string
	LogType    string
	Time       int64
	Additional interface{}
	Message    interface{}
}

// NewLogObject - enhance method to create and return reference of LogObject
func NewLogObject(level LogLevel, module string, logType LogType, eventTime time.Time, message interface{}, additional interface{}) *LogObject {
	return &LogObject{
		Level:      level,
		Module:     module,
		LogType:    logType.String(),
		Time:       eventTime.UTC().UnixNano(),
		Message:    message,
		Additional: additional,
	}
}

// LogLevel Object
type LogLevel int

// Some Constants - used with LogLevel
const (
	DEBUG LogLevel = max - iota
	INFO
	WARNING
	ERROR
)

// IsLogLevel - check whether is a true log level
func (l LogLevel) IsLogLevel() bool {
	if l <= DEBUG && l >= ERROR {
		return true
	}

	return false
}

// LogType Object
type LogType struct {
	name string
}

var (
	FuncMaintenanceType = LogType{name: "FUNC_MAINT"}
	DebugType           = LogType{name: "DEBUG_INFORMATION"}
)

//const (
//	TYPE_FUNC_MAINT LogType = iota
//	TYPE_DB_ERROR
//	TYPE_CONVERT_ERROR
//	TYPE_HTTP_ERROR
//	TYPE_NIL_OBJECT
//	TYPE_ENCRYPTION_ERROR
//	TYPE_HTTP_REQUESTED
//	TYPE_HTTP_RESPONSED
//	TYPE_DEBUG_INFORMATION
//)

func (l LogType) String() string {
	return l.name
}
