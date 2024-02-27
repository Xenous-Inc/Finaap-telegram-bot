package flags

import (
	"errors"
	"flag"

	"github.com/4kayDev/logger/log"
)

const (
	configPathFlag = "config-path"
	envModeFlag    = "env-mode"
)

type CMDFlags struct {
	ConfigPath string
	EnvMode    string
}

func ParseFlags() (*CMDFlags, error) {
	configPath := flag.String(configPathFlag, "", "Configuration file path")
	envMode := flag.String(envModeFlag, "", "Environment mode")
	flag.Parse()

	if *configPath == "" {
		log.Warn("configPath was not found in application flags", "utils/flags ParseFlags")
		return nil, errors.New("Configuration file path was not found in application flags")
	}

	if *envMode == "" {
		log.Warn("envMode was not found in application flags", "utils/flags ParseFlags")
		return nil, errors.New("Environment mode was not found in application flags")
	}

	return &CMDFlags{ConfigPath: *configPath, EnvMode: *envMode}, nil
}

func MustParseFlags() *CMDFlags {
	flags, err := ParseFlags()
	if err != nil {
		log.Error(err, "Error while ParseFlags", "utils/flags MustParseFlags")
		panic(err)
	}

	return flags
}
