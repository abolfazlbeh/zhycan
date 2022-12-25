package logger

// Imports needed list
import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"zhycan/internal/config"
	"zhycan/internal/utils"
)

// Mark: ZapWrapper

// ZapWrapper object - implements Logger interface
type ZapWrapper struct {
	name            string
	serviceName     string
	logger          *zap.Logger
	ch              chan LogObject
	initialized     bool
	wg              sync.WaitGroup
	operationType   string
	supportedOutput []string
}

// Constructor - It initializes the logger configuration params
func (l *ZapWrapper) Constructor(name string) {
	l.wg.Add(1)
	defer l.wg.Done()

	l.name = name
	l.serviceName = config.GetManager().GetName()
	l.operationType = config.GetManager().GetOperationType()
	l.supportedOutput = []string{"console", "file"}
	l.initialized = false

	channelSize, err := config.GetManager().Get(l.name, "channel_size")
	if err != nil {
		return
	}

	options, err := config.GetManager().Get(l.name, "options")
	if err != nil {
		return
	}

	optionArray := make([]string, len(options.([]interface{})))
	for _, v := range options.([]interface{}) {
		optionArray = append(optionArray, v.(string))
	}

	outputs, err := config.GetManager().Get(l.name, "outputs")
	if err != nil {
		return
	}

	outputArray := make([]string, len(outputs.([]interface{})))
	for _, v := range outputs.([]interface{}) {
		outputArray = append(outputArray, v.(string))
	}

	l.ch = make(chan LogObject, int(channelSize.(float64)))

	if l.operationType == "prod" {
		productionEncoderConfig := zap.NewProductionEncoderConfig()
		productionEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		productionEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		var cores []zapcore.Core
		for _, outputItem := range outputArray {
			if utils.ArrayContains(&l.supportedOutput, outputItem) {
				if outputItem == "console" {
					level := zapcore.DebugLevel

					key := fmt.Sprintf("%s.level", outputItem)
					levelStr, err := config.GetManager().Get(l.name, key)
					if err == nil {
						level, err = zapcore.ParseLevel(levelStr.(string))
						if err != nil {
							continue
						}
					}

					consoleEncoder := zapcore.NewConsoleEncoder(productionEncoderConfig)
					c := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
					cores = append(cores, c)
				}
			}
		}

		core := zapcore.NewTee(
			cores...,
		)
		l.logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else { // otherwise we create a development version (`dev`, `test`, ...)
		developmentEncoderConfig := zap.NewDevelopmentEncoderConfig()
		developmentEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		developmentEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		var cores []zapcore.Core
		for _, outputItem := range outputArray {
			if utils.ArrayContains(&l.supportedOutput, outputItem) {
				if outputItem == "console" {
					level := zapcore.DebugLevel

					key := fmt.Sprintf("%s.level", outputItem)
					levelStr, err := config.GetManager().Get(l.name, key)
					if err == nil {
						level, err = zapcore.ParseLevel(levelStr.(string))
						if err != nil {
							continue
						}
					}

					consoleEncoder := zapcore.NewConsoleEncoder(developmentEncoderConfig)
					c := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

					cores = append(cores, c)
				}
			}
		}

		core := zapcore.NewTee(
			cores...,
		)
		l.logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	go l.runner()
	l.initialized = true
}

// Close - it closes logger channel
func (l *ZapWrapper) Close() {
	l.wg.Wait()

	_ = l.logger.Sync()
	defer close(l.ch)
}

// Log - write log object to the channel
func (l *ZapWrapper) Log(obj *LogObject) {
	l.wg.Wait()

	go func(obj *LogObject) {
		l.ch <- *obj
	}(obj)
}

// IsInitialized - that returns boolean value whether it's initialized
func (l *ZapWrapper) IsInitialized() bool {
	return l.initialized
}

// Instance - returns exact logger instance
func (l *ZapWrapper) Instance() *zap.Logger {
	l.wg.Wait()
	return l.logger
}

// ZapWrapper runner - the goroutine that reads from channel and process it
func (l *ZapWrapper) runner() {
	l.wg.Wait()
	for c := range l.ch {
		if c.Level.IsLogLevel() {
			f := []zapcore.Field{
				zap.Any("service", l.serviceName),
				zap.Any("module", c.Module),
				zap.Any("log_type", c.LogType),
				zap.Any("time", c.Time),
				zap.Any("additional", c.Additional),
			}
			switch c.Level {
			case DEBUG:
				l.logger.Debug(fmt.Sprintf("%v", c.Message), f...)
				break
			case INFO:
				l.logger.Info(fmt.Sprintf("%v", c.Message), f...)
				break
			case WARNING:
				l.logger.Warn(fmt.Sprintf("%v", c.Message), f...)
				break
			case ERROR:
				l.logger.Error(fmt.Sprintf("%v", c.Message), f...)
				break
			}
		}
	}
}