package config

import (
	"errors"

	"github.com/4kayDev/logger/log"
	"github.com/spf13/viper"
)

const (
	ENV_MODE_DEVELOPMENT = iota + 1
	ENV_MODE_PRODUCTION  = iota + 1
	ENV_MODE_STAGE       = iota + 1
)

const (
	envModeDevelopmentStr = "development"
	envModeProductionStr  = "production"
	envModeStageStr       = "stage"
)

type Config struct {
	EnvMode uint8
	Token   string
}

func LoadConfig(envMode, path string) (*Config, error) {
	mode, err := validateEnvMode(envMode)
	if err != nil {
		log.Error(err, "Error while validateEnvMode", "utils/config LoadConfig")
		return nil, err
	}

	config := new(Config)

	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err, "Error while ReadInConfig", "utils/config LoadConfig")
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Error(err, "Error while Unmarshal", "utils/config LoadConfig")
		return nil, err
	}

	config.EnvMode = mode

	return config, nil
}

func MustLoadConfig(envMode, path string) *Config {
	config, err := LoadConfig(envMode, path)
	if err != nil {
		log.Error(err, "Error while LoadConfig", "utils/config MustLoadConfig")
		panic(err)
	}

	return config
}

func validateEnvMode(envMode string) (uint8, error) {
	var mode uint8
	switch envMode {
	case envModeDevelopmentStr:
		mode = ENV_MODE_DEVELOPMENT
	case envModeProductionStr:
		mode = ENV_MODE_PRODUCTION
	case envModeStageStr:
		mode = ENV_MODE_STAGE
	default:
		return mode, errors.New("Unknown environment mode")
	}

	return mode, nil
}
