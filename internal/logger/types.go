package logger

// Imports needed list
import (
	"strings"
	"time"
)

// Some Constants
const (
	MAX = 6
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
	DEBUG LogLevel = MAX - iota
	INFO
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	}
	return "DEFAULT"
}

// IsLogLevel - check whether is a true log level
func (l LogLevel) IsLogLevel() bool {
	if l <= DEBUG && l >= ERROR {
		return true
	}

	return false
}

func StringToLogLevel(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	}
	return DEBUG
}

// LogType Object
type LogType struct {
	name string
}

func NewLogType(name string) LogType {
	return LogType{name: name}
}

var (
	FuncMaintenanceType = LogType{name: "FUNC_MAINT"}
	DebugType           = LogType{name: "DEBUG_INFORMATION"}
)

func (l LogType) String() string {
	return l.name
}
