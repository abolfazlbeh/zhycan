package config

import (
	"reflect"
	"testing"
	"time"
)

func createWrapper(configPath, name, envPrefix string) (*ViperWrapper, error) {
	v := &ViperWrapper{
		ConfigPath:          []string{configPath},
		ConfigName:          name,
		ConfigEnvPrefix:     envPrefix,
		ConfigResourcePlace: "",
	}
	err := v.Load()
	return v, err
}

func TestWrapperCreation(t *testing.T) {
	err := createManager()
	if err == nil {
		time.Sleep(3 * time.Second)

		actualWrapper, err := createWrapper("../../configs/test/",
			"logger",
			"")
		if err != nil {
			t.Errorf("Cannot create wrapper object --> expected %v, but got %v", nil, err)
		}

		expectedWrapper, err := GetManager().GetConfigWrapper("logger")
		if err != nil {
			t.Errorf("Cannot create wrapper object from manager --> expected %v, but got %v", nil, err)
		}

		if !reflect.DeepEqual(actualWrapper, expectedWrapper) {
			t.Errorf("cannot create wrapper object --> Expected %v, but got %v", expectedWrapper, actualWrapper)
		}
	}
}

func TestGetExistValue(t *testing.T) {
	actualWrapper, err := createWrapper("../../configs/test/",
		"logger",
		"")
	if err != nil {
		t.Errorf("Cannot create wrapper object --> expected %v, but got %v", nil, err)
	}

	expectedVal := "zap"
	actualVal, _ := actualWrapper.Get("type", false)
	if expectedVal != actualVal {
		t.Errorf("Get the key for the `type` --> Expected: %v, got %v", expectedVal, actualVal)
	}
}
