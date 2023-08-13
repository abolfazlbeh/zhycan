package db

import (
	"context"
	"errors"
	"fmt"
	zlogger "github.com/abolfazlbeh/zhycan/internal/logger"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

// MARK: Variables
var (
	DbLogType      = zlogger.NewLogType("DB_OP")
	DbTraceLogType = zlogger.NewLogType("DB_TRACE_OP")
)

// DbLogger - DB Logger struct
type DbLogger struct {
	logger.Config
}

// NewDbLogger - return instance of DbLogger which implement Interface
func NewDbLogger(config LoggerConfig) logger.Interface {
	logLevel := logger.Silent
	switch strings.ToLower(config.LogLevel) {
	case "error":
		logLevel = logger.Error
		break
	case "warn":
		logLevel = logger.Warn
		break
	case "info":
		logLevel = logger.Info
		break
	}

	return &DbLogger{Config: logger.Config{
		SlowThreshold:             time.Duration(config.SlowThreshold) * time.Millisecond,
		Colorful:                  false,
		IgnoreRecordNotFoundError: config.IgnoreRecordNotFoundError,
		ParameterizedQueries:      config.ParameterizedQueries,
		LogLevel:                  logLevel,
	}}
}

// LogMode - set log mode
func (l *DbLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info - print info
func (l DbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	ll, _ := zlogger.GetManager().GetLogger()
	if l.LogLevel >= logger.Info && ll != nil {
		newMsg := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		ll.Log(zlogger.NewLogObject(
			zlogger.INFO, "db", DbLogType,
			time.Now().UTC(), newMsg, nil,
		))
	}
}

// Warn - print warn messages
func (l DbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	ll, _ := zlogger.GetManager().GetLogger()
	if l.LogLevel >= logger.Warn && ll != nil {
		newMsg := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		ll.Log(zlogger.NewLogObject(
			zlogger.WARNING, "db", DbLogType,
			time.Now().UTC(), newMsg, nil,
		))
	}
}

// Error - print error messages
func (l DbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	ll, _ := zlogger.GetManager().GetLogger()
	if l.LogLevel >= logger.Error && ll != nil {
		newMsg := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		ll.Log(zlogger.NewLogObject(
			zlogger.ERROR, "db", DbLogType,
			time.Now().UTC(), newMsg, nil,
		))
	}
}

// Trace - print sql message
func (l DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	ll, _ := zlogger.GetManager().GetLogger()
	if ll != nil {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error &&
			(!errors.Is(err, logger.ErrRecordNotFound) ||
				!l.IgnoreRecordNotFoundError):
			sql, rows := fc()
			msgLiteral := "%s %s\n[%.3fms] [rows:%v] %s"
			if rows == -1 {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.ERROR, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{err, sql},
				))
			} else {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.ERROR, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{err, sql, rows},
				))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)

			msgLiteral := "%s %s\n[%.3fms] [rows:%v] %s"
			if rows == -1 {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.WARNING, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{slowLog, sql},
				))
			} else {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.WARNING, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{slowLog, sql, rows},
				))
			}
		case l.LogLevel == logger.Info:
			sql, rows := fc()

			msgLiteral := "%s\n[%.3fms] [rows:%v] %s"
			if rows == -1 {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.INFO, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{sql},
				))
			} else {
				msg := fmt.Sprintf(msgLiteral, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
				ll.Log(zlogger.NewLogObject(
					zlogger.INFO, "db", DbTraceLogType,
					time.Now().UTC(), msg, []interface{}{sql, rows},
				))
			}
		}
	}
}
