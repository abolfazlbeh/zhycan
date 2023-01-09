package shortcuts

import (
	"time"
	"zhycan/internal/logger"
)

type LogError logger.Error
type LogObject logger.LogObject
type LogLevel logger.LogLevel
type LogType logger.LogType

// NewLogObject - enhance method to create and return reference of LogObject
func NewLogObject(level LogLevel, module string, logType LogType, eventTime time.Time, message interface{}, additional interface{}) *LogObject {
	return &LogObject{
		Level:      logger.LogLevel(level),
		Module:     module,
		LogType:    logger.LogType(logType).String(),
		Time:       eventTime.UTC().UnixNano(),
		Message:    message,
		Additional: additional,
	}
}

// Log - write log object to the channel
func Log(object *LogObject) *LogError {
	l, err := logger.GetManager().GetLogger()
	if err == nil {
		if l.IsInitialized() {
			p := logger.LogObject(*object)
			l.Log(&p)
		}
		return nil
	}

	p := LogError(*err)
	return &p
}

// Sync - sync all logs to medium
func Sync() *LogError {
	l, err := logger.GetManager().GetLogger()
	if err == nil {
		if l.IsInitialized() {
			l.Sync()
		}
		return nil
	}

	p := LogError(*err)
	return &p
}

// Close - it closes logger channel
func Close() *LogError {
	l, err := logger.GetManager().GetLogger()
	if err == nil {
		if l.IsInitialized() {
			l.Close()
		}
		return nil
	}

	p := LogError(*err)
	return &p
}
