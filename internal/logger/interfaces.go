package logger

// Logger interface
type Logger interface {
	Constructor(name string)
	Close()
	Log(obj *LogObject)
	IsInitialized() bool
}
