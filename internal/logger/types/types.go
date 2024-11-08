package types

// Imports needed list
import (
	"gorm.io/gorm"
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

// NewLogType - Create A New Log Type
func NewLogType(name string) LogType {
	return LogType{name: name}
}

var (
	FuncMaintenanceType = LogType{name: "FUNC_MAINT"}
	DebugType           = LogType{name: "DEBUG_INFORMATION"}
	NilObject           = LogType{name: "NIL_OBJECT"}
)

func (l LogType) String() string {
	return l.name
}

// MARK: DB Record

type ZhycanLog struct {
	gorm.Model

	ServiceName string `gorm:"not null;size:256" json:"service_name"`
	Level       string `gorm:"size:64" json:"level"`
	LogType     string `gorm:"size:256" json:"log_type"`
	Module      string `gorm:"size:1024" json:"module"`
	Message     string `json:"message"`
	Additional  string `json:"additional"`
	LogTime     int64  `json:"logTime"`
}
