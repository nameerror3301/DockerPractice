package utility

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v3"
)

type Parser interface {
	Unmarshal(data []byte, v interface{})
}

type driver_cfg string

var (
	driverJSON driver_cfg = "json"
	driverTOML driver_cfg = "toml"
	driverYAML driver_cfg = "yaml"

	ENV_DriverConfig = "DRIVER_CONFIG"
	ENV_PathConfig   = "PATH_CONFIG"
)

// CheckConfigDriver get type of driver for load config
func CheckConfigDriver() func([]byte, interface{}) error {
	var wal = os.Getenv(ENV_DriverConfig)
	switch driver_cfg(wal) {
	case driverJSON:
		{
			return json.Unmarshal
		}
	case driverYAML:
		{
			return yaml.Unmarshal
		}
	case driverTOML:
		{
			return toml.Unmarshal
		}
	default:
		{
			return json.Unmarshal
		}
	}
}

// ReadConfig читает конфигурационный файл
//
//	func SomeFunction() {
//		config, err = ReadConfig("configs/config.{json,yaml,toml}")
//		if err != nil {
//			panic(err)
//		}
//	}
func ReadConfig() (*ConfigModel, error) {
	var path string = os.Getenv(ENV_PathConfig)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	driver := CheckConfigDriver()
	var config = &ConfigModel{}

	if err = driver(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
