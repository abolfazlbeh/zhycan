package config

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

// MARK: ViperWrapper

// ViperWrapper object
type ViperWrapper struct {
	Instance            *viper.Viper
	ConfigPath          []string
	ConfigName          string
	ConfigEnvPrefix     string
	ConfigResourcePlace string
	lastModified        time.Time
	wg                  sync.WaitGroup
	lock                sync.Mutex
}

// MARK: Public Methods

// Load - It creates new instance of Viper and load config file base on ConfigName
func (w *ViperWrapper) Load() error {
	w.wg.Add(1)
	defer w.wg.Done()

	w.Instance = viper.New()
	for _, path := range w.ConfigPath {
		w.Instance.AddConfigPath(path)
	}
	w.Instance.SetConfigName(w.ConfigName)
	err := w.Instance.ReadInConfig()
	if err != nil {
		return err
	}

	// Get env variables and bind them if exist in config file
	env, exist := w.Get("env", true)
	if env != nil && exist {
		envArray := make([]string, len(env.([]interface{})))
		for i, v := range env.([]interface{}) {
			envArray[i] = v.(string)
		}

		if len(envArray) > 0 {
			w.Instance.SetEnvPrefix(w.ConfigEnvPrefix)

			for _, e := range envArray {
				_ = w.Instance.BindEnv(e)
			}
		}
	}

	return nil
}

// LoadFromRemote - loads the configs from the remote server
func (w *ViperWrapper) LoadFromRemote(data []byte) error {
	w.wg.Add(1)
	defer w.wg.Done()

	if w.Instance == nil {
		w.Instance = viper.New()
	}

	err := w.Instance.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	// Get env variables and bind them if exist in config file
	env, exist := w.Get("env", true)
	if env != nil && exist {
		envArray := make([]string, len(env.([]interface{})))
		for i, v := range env.([]interface{}) {
			envArray[i] = v.(string)
		}

		if len(envArray) > 0 {
			w.Instance.SetEnvPrefix(w.ConfigEnvPrefix)

			for _, e := range envArray {
				_ = w.Instance.BindEnv(e)
			}
		}
	}

	return nil
}

// RegisterChangeCallback - get function and call it when config file changed
func (w *ViperWrapper) RegisterChangeCallback(fn func() interface{}) {
	w.wg.Wait()

	w.Instance.WatchConfig()
	w.Instance.OnConfigChange(func(e fsnotify.Event) {
		log.Println(w.ConfigName, "Config file changed: ", e.Name)

		if fn != nil {
			fn()
		}
	})
}

// Get method - returns value base on key
func (w *ViperWrapper) Get(key string, bypass bool) (interface{}, bool) {
	if !bypass {
		w.wg.Wait()
	}

	w.lock.Lock()
	defer w.lock.Unlock()

	exist := w.Instance.InConfig(key)
	return w.Instance.Get(key), exist
}

// Set method - set value by given key and write it back to file
func (w *ViperWrapper) Set(key string, value interface{}, bypass bool) error {
	if !bypass {
		w.wg.Wait()
	}

	w.lock.Lock()
	defer w.lock.Unlock()

	w.Instance.Set(key, value)
	return w.Instance.SafeWriteConfig()
}
