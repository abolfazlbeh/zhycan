package logger

// Logger interface
type Logger interface {
	Constructor(name string) error
	Close()
	Log(obj *LogObject)
	IsInitialized() bool
	Sync()
}
