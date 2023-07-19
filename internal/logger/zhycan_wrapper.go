package logger

// Imports needed list
import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
	"zhycan/internal/config"
	"zhycan/internal/utils"
)

type OutputOption struct {
	LevelStr string      `json:"level"`
	Level    LogLevel    `json:"-"`
	Path     string      `json:"path,omitempty"`
	l        *log.Logger `json:"-"`
}

// MARK: LogMeWrapper

// LogMeWrapper structure - implements Logger interface
type LogMeWrapper struct {
	name                  string
	serviceName           string
	ch                    chan LogObject
	initialized           bool
	wg                    sync.WaitGroup
	operationType         string
	supportedOutput       []string
	supportedOutputOption map[string]OutputOption
}

// Constructor - It initializes the logger configuration params
func (l *LogMeWrapper) Constructor(name string) error {
	l.wg.Add(1)
	defer l.wg.Done()

	l.name = name
	l.serviceName = config.GetManager().GetName()
	l.operationType = config.GetManager().GetOperationType()
	l.supportedOutput = []string{"console", "file"}
	l.initialized = false

	channelSize, err := config.GetManager().Get(l.name, "channel_size")
	if err != nil {
		return err
	}

	options, err := config.GetManager().Get(l.name, "options")
	if err != nil {
		return err
	}

	optionArray := make([]string, len(options.([]interface{})))
	for _, v := range options.([]interface{}) {
		optionArray = append(optionArray, v.(string))
	}

	outputs, err := config.GetManager().Get(l.name, "outputs")
	if err != nil {
		return err
	}

	var outputArray []string
	for _, v := range outputs.([]interface{}) {
		outputArray = append(outputArray, v.(string))
	}

	l.ch = make(chan LogObject, int(channelSize.(float64)))

	//if l.operationType == "prod" {
	//} else {
	//}
	l.supportedOutputOption = make(map[string]OutputOption)
	for _, item := range outputArray {
		if utils.ArrayContains(&l.supportedOutput, item) {
			jsonObj, configReadErr := config.GetManager().Get(l.name, item)
			if configReadErr == nil {
				r := OutputOption{}

				jsonStr, jsonErr := json.Marshal(jsonObj.(map[string]interface{}))
				if jsonErr == nil {
					jsonErr = json.Unmarshal(jsonStr, &r)
					if jsonErr == nil {
						// add it to internal map
						r.Level = StringToLogLevel(r.LevelStr)
						if item == "console" {
							r.l = log.New(os.Stdout, "", 0)
						}
						l.supportedOutputOption[item] = r

						// run the instance ...
						go l.runner(item)
					} else {
						log.Printf("Cannot create log instance for: %v - %v", item, jsonErr)
					}
				} else {
					log.Printf("Cannot create log instance for: %v - %v", item, jsonErr)
				}
			} else {
				log.Printf("Cannot create log instance for: %v - %v", item, configReadErr)
			}
		} else {
			log.Printf("Log outout with name `%v` is not supported yet", item)
		}
	}

	l.initialized = true

	return nil
}

// IsInitialized - that returns boolean value whether it's initialized
func (l *LogMeWrapper) IsInitialized() bool {
	return l.initialized
}

// Log - write log object to the channel
func (l *LogMeWrapper) Log(obj *LogObject) {
	l.wg.Wait()

	go func(obj *LogObject) {
		l.ch <- *obj
	}(obj)
}

// Sync - sync all logs to medium
func (l *LogMeWrapper) Sync() {
	l.wg.Wait()
	ch := make(chan bool, 1)
	go func() {
		for {
			if len(l.ch) > 0 {
				time.Sleep(time.Millisecond * 200)
			} else {
				ch <- true
			}
		}
	}()

	<-ch
}

// Close - it closes logger channel
func (l *LogMeWrapper) Close() {
	l.wg.Wait()
	l.Sync()
	defer close(l.ch)
}

// runner - the goroutine that reads from channel and process it
func (l *LogMeWrapper) runner(output string) {
	l.wg.Wait()
	for c := range l.ch {
		if c.Level <= l.supportedOutputOption[output].Level {
			switch c.Level {
			case DEBUG:
				l.debug(&c, output)
			case INFO:
				l.info(&c, output)
			case WARNING:
				l.info(&c, output)
			case ERROR:
				l.info(&c, output)
			}
		}
	}
}

// debug - log with DEBUG level
func (l *LogMeWrapper) debug(object *LogObject, output string) {
	if output == "console" {
		l.supportedOutputOption[output].l.Printf(
			"\\e[37m%v %v >>> %7v >>> (%v/%v)  - %v ... %v\\e[0m\n",
			l.serviceName,
			object.Time,
			object.Level.String(),
			object.LogType,
			object.Module,
			object.Module,
			object.Additional,
		)
	}
}

// info - log with INFO level
func (l *LogMeWrapper) info(object *LogObject, output string) {
	if output == "console" {
		l.supportedOutputOption[output].l.Printf(
			"\\e[32m%v %v >>> %7v >>> (%v/%v)  - %v ... %v\\e[0m\n",
			l.serviceName,
			object.Time,
			object.Level.String(),
			object.LogType,
			object.Module,
			object.Module,
			object.Additional,
		)
	}
}

// warning - log with WARNING level
func (l *LogMeWrapper) warning(object *LogObject, output string) {
	if output == "console" {
		l.supportedOutputOption[output].l.Printf(
			"\\e[33m%v %v >>> %7v >>> (%v/%v)  - %v ... %v\\e[0m\n",
			l.serviceName,
			object.Time,
			object.Level.String(),
			object.LogType,
			object.Module,
			object.Module,
			object.Additional,
		)
	}
}

// error - log with ERROR level
func (l *LogMeWrapper) error(object *LogObject, output string) {
	if output == "console" {
		l.supportedOutputOption[output].l.Printf(
			"\\e[31m%v %v >>> %7v >>> (%v/%v)  - %v ... %v\\e[0m\n",
			l.serviceName,
			object.Time,
			object.Level.String(),
			object.LogType,
			object.Module,
			object.Module,
			object.Additional,
		)
	}
}
